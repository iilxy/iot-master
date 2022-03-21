package modbus

import (
	"errors"
	"fmt"
	"github.com/zgwit/iot-master/connect"
	"github.com/zgwit/iot-master/protocol"
	"github.com/zgwit/iot-master/protocol/helper"
	"sync"
	"time"
)

//TCP Modbus-TCP协议
type TCP struct {
	link  connect.Link
	queue chan interface{} //in

	requests  sync.Map
	increment uint16

	slave uint8
}

func newTCP(link connect.Link, opts protocol.Options) protocol.Protocol {
	tcp := &TCP{
		link:      link,
		queue:     make(chan interface{}, 10), //TODO 改成参数
		increment: 0x0A0A,                     //避免首字节为0，有些DTU可能会异常
		slave:     opts["slave"].(uint8),
	}
	link.On("data", func(data []byte) {
		tcp.OnData(data)
	})
	link.On("close", func() {
		close(tcp.queue)
		tcp.requests.Range(func(key, value interface{}) bool {
			close(value.(*request).resp)
			return true
		})
	})
	return tcp
}

func (m *TCP) execute(cmd []byte, immediate bool) ([]byte, error) {
	req := &request{
		cmd:  cmd,
		resp: make(chan response, 1),
	}

	if !immediate {
		//排队等待
		m.queue <- nil
	}

	id := m.increment
	helper.WriteUint16(cmd, id) //写入事务ID
	m.increment++
	if m.increment < 0x0A0A {
		m.increment = 0x0A0A
	}
	m.requests.Store(id, req)

	//下发指令
	err := m.link.Write(cmd)
	if err != nil {
		if !immediate {
			//释放队列
			<-m.queue
		}
		return nil, err
	}

	//等待结果
	select {
	case <-time.After(5 * time.Second):
		if !immediate {
			<-m.queue //清空
		}
		return nil, errors.New("timeout")
	case resp := <-req.resp:
		return resp.buf, resp.err
	}
}

func (m *TCP) OnData(buf []byte) {
	//取出请求，并让出队列，可以开始下一个请示了
	if len(m.queue) > 0 {
		<-m.queue
	}

	//解析数据
	l := len(buf)
	if l < 10 {
		return //长度不够
	}

	id := helper.ParseUint16(buf)
	length := helper.ParseUint16(buf[4:])
	if l < int(length)+6 {
		return //长度不够
	}

	r, ok := m.requests.LoadAndDelete(id)
	if !ok {
		return
	}
	req := r.(*request)

	//slave := buf[6]
	fc := buf[7]
	//解析错误码
	if fc&0x80 > 0 {
		req.resp <- response{err: fmt.Errorf("错误码：%d", buf[2])}
		return
	}

	//解析数据
	length = 4
	count := int(buf[8])
	switch fc {
	case FuncCodeReadDiscreteInputs,
		FuncCodeReadCoils:
		//数组解压
		bb := helper.ExpandBool(buf[9:], count)
		req.resp <- response{buf: bb}
	case FuncCodeReadInputRegisters,
		FuncCodeReadHoldingRegisters,
		FuncCodeReadWriteMultipleRegisters:
		req.resp <- response{buf: helper.Dup(buf[9:])}
	default:
		req.resp <- response{}
	}
}

func (m *TCP) Address(addr string) (protocol.Addr, error) {
	return ParseAddress(addr)
}

func (m *TCP) Read(address protocol.Addr, size uint16) ([]byte, error) {
	addr := address.(*Address)
	b := make([]byte, 12)
	//helper.WriteUint16(b, id)
	helper.WriteUint16(b[2:], 0) //协议版本
	helper.WriteUint16(b[4:], 6) //剩余长度
	b[6] = m.slave
	b[7] = addr.Code
	helper.WriteUint16(b[8:], addr.Offset)
	helper.WriteUint16(b[10:], size)

	return m.execute(b, false)
}

func (m *TCP) Immediate(address protocol.Addr, size uint16) ([]byte, error) {
	addr := address.(*Address)
	b := make([]byte, 12)
	//helper.WriteUint16(b, id)
	helper.WriteUint16(b[2:], 0) //协议版本
	helper.WriteUint16(b[4:], 6) //剩余长度
	b[6] = m.slave
	b[7] = addr.Code
	helper.WriteUint16(b[8:], addr.Offset)
	helper.WriteUint16(b[10:], uint16(size))

	return m.execute(b, true)
}

func (m *TCP) Write(address protocol.Addr, buf []byte) error {
	addr := address.(*Address)
	length := len(buf)
	//如果是线圈，需要Shrink
	code := addr.Code
	switch code {
	case FuncCodeReadCoils:
		if length == 1 {
			code = 5
			//数据 转成 0x0000 0xFF00
			if buf[0] > 0 {
				buf = []byte{0xFF, 0}
			} else {
				buf = []byte{0, 0}
			}
		} else {
			code = 15 //0x0F
			//数组压缩
			b := helper.ShrinkBool(buf)
			count := len(b)
			buf = make([]byte, 3+count)
			helper.WriteUint16(buf, uint16(length))
			buf[2] = uint8(count)
			copy(buf[3:], b)
		}
	case FuncCodeReadHoldingRegisters:
		if length == 2 {
			code = 6
		} else {
			code = 16 //0x10
			b := make([]byte, 3+length)
			helper.WriteUint16(b, uint16(length/2))
			b[2] = uint8(length)
			copy(b[3:], buf)
			buf = b
		}
	default:
		return errors.New("功能码不支持")
	}

	l := 10 + len(buf)
	b := make([]byte, l)
	//helper.WriteUint16(b, id)
	helper.WriteUint16(b[2:], 0) //协议版本
	helper.WriteUint16(b[4:], 6) //剩余长度
	b[6] = m.slave
	b[7] = code
	helper.WriteUint16(b[8:], addr.Offset)
	copy(b[10:], buf)

	_, err := m.execute(b, true)
	return err
}