import{q as t,o as i,c as e,w as s,a as o,b as a,r as l,e as n,d as r}from"./index.8d4a8305.js";import{_ as d,a as u}from"./uni-list.723f536b.js";import{d as c,_ as m}from"./time.0102cd3d.js";import{_ as f}from"./uni-icons.51822eb1.js";import{_ as p}from"./uni-section.96f3ba19.js";import{r as h}from"./const.9e5d4232.js";import{_}from"./plugin-vue_export-helper.21dcd24c.js";import"./_commonjsHelpers.4e997714.js";var g=_({data:()=>({data:{}}),onLoad(t){this.id=t.id,this.load()},onPullDownRefresh(){this.load()},methods:{format:c,load(){h({url:"product/"+this.id,success:t=>{this.data=t},complete(){uni.stopPullDownRefresh()}})},remove(){uni.showModal({title:"提示",content:"确定删除？",success:i=>{i.confirm&&(t("log","at pages/product/detail.vue:72","用户点击确定"),h({url:"product/"+this.id+"/delete",success:t=>{uni.navigateBack(),uni.showToast({title:"删除成功"})}}))},fail:console.error})}}},[["render",function(t,c,h,_,g,x){const T=l(n("uni-list-item"),d),b=l(n("uni-list"),u),j=l(n("uni-card"),m),v=l(n("uni-icons"),f),k=l(n("uni-section"),p),w=r;return i(),e(w,null,{default:s((()=>[o(j,{title:g.data.name,subTitle:g.data.id,note:"Tips",thumbnail:"/static/icons/product.svg"},{default:s((()=>[o(b,{border:!1},{default:s((()=>[o(T,{title:"厂商",rightText:g.data.manufacturer},null,8,["rightText"]),o(T,{title:"版本",rightText:g.data.version},null,8,["rightText"]),o(T,{title:"创建时间",rightText:x.format(g.data.created)},null,8,["rightText"])])),_:1})])),_:1},8,["title","subTitle"]),o(b,null,{default:s((()=>[o(T,{title:"编辑产品",link:"",to:"./edit?id="+t.id},{header:s((()=>[o(v,{class:"list-icon",customPrefix:"iconfont",type:"icon-pen"})])),_:1},8,["to"]),o(T,{title:"删除产品",onClick:x.remove,clickable:!0},{header:s((()=>[o(v,{class:"list-icon",customPrefix:"iconfont",type:"icon-dustbin"})])),_:1},8,["onClick"])])),_:1}),o(k,{title:"日志",type:"line"},{default:s((()=>[a(" TODO：Event ")])),_:1})])),_:1})}]]);export{g as default};
