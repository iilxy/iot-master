import{q as t,o as i,c as e,w as s,a,r as o,e as l,d as n}from"./index.8d4a8305.js";import{_ as r,a as d}from"./uni-list.723f536b.js";import{d as c,_ as u}from"./time.0102cd3d.js";import{_ as m}from"./uni-icons.51822eb1.js";import{r as f}from"./const.9e5d4232.js";import{_ as h}from"./plugin-vue_export-helper.21dcd24c.js";import"./_commonjsHelpers.4e997714.js";var p=h({data:()=>({data:{}}),onLoad(t){this.id=t.id,this.load()},onPullDownRefresh(){this.load()},methods:{format:c,load(){f({url:"interface/"+this.id,success:t=>{this.data=t},complete(){uni.stopPullDownRefresh()}})},remove(){uni.showModal({title:"提示",content:"确定删除？",success:i=>{i.confirm&&(t("log","at pages/interface/detail.vue:77","用户点击确定"),f({url:"interface/"+this.id+"/delete",success:t=>{uni.navigateBack(),uni.showToast({title:"删除成功"})}}))},fail:console.error})}}},[["render",function(t,c,f,h,p,g){const _=o(l("uni-list-item"),r),x=o(l("uni-list"),d),T=o(l("uni-card"),u),b=o(l("uni-icons"),m),j=n;return i(),e(j,null,{default:s((()=>[a(T,{title:p.data.name,subTitle:p.data.id,extra:"在线",note:"Tips",thumbnail:"/static/icons/design.svg"},{default:s((()=>[a(x,{border:!1},{default:s((()=>[a(_,{title:"ID",rightText:p.data.id},null,8,["rightText"]),a(_,{title:"版本",rightText:p.data.version},null,8,["rightText"]),a(_,{title:"创建时间",rightText:p.data.created},null,8,["rightText"])])),_:1})])),_:1},8,["title","subTitle"]),a(x,null,{default:s((()=>[a(_,{title:"设计器",link:"",to:"./design?id="+t.id},{header:s((()=>[a(b,{class:"list-icon",type:"color"})])),_:1},8,["to"]),a(_,{title:"编辑组态",link:"",to:"./edit?id="+t.id},{header:s((()=>[a(b,{class:"list-icon",customPrefix:"iconfont",type:"icon-pen"})])),_:1},8,["to"]),a(_,{title:"删除组态",onClick:g.remove,clickable:!0},{header:s((()=>[a(b,{class:"list-icon",customPrefix:"iconfont",type:"icon-dustbin"})])),_:1},8,["onClick"])])),_:1})])),_:1})}]]);export{p as default};
