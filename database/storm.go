package database

import (
	"github.com/zgwit/storm/v3"
	"os"
	"path/filepath"
)

// Master 基础数据库
var Master *storm.DB

// History 历史数据库
var History *storm.DB

// Error 错误数据库
var Error *storm.DB

//Open 打开数据库
func Open(opts *Options) error {
	if opts == nil {
		opts = DefaultOptions()
	}

	err := os.MkdirAll(opts.Path, os.ModePerm)
	if err != nil {
		return err
	}

	//基础数据
	Master, err = storm.Open(filepath.Join(opts.Path, "master.db"))
	if err != nil {
		return err
	}

	//历史数据
	History, err = storm.Open(filepath.Join(opts.Path, "history.db"))
	if err != nil {
		return err
	}

	//错误日志
	Error, err = storm.Open(filepath.Join(opts.Path, "error.db"))
	if err != nil {
		return err
	}

	return nil
}

//Close 关闭数据库
func Close() error {
	err := Master.Close()
	if err != nil {
		return err
	}

	err = History.Close()
	if err != nil {
		return err
	}

	err = Error.Close()
	if err != nil {
		return err
	}

	return nil
}
