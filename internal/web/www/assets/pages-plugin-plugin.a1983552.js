import{K as a,L as t,M as e,B as s,o,c as i,w as d,a as l,j as n,C as r,F as u,r as m,e as c,d as h,i as p}from"./index.9b1e08a7.js";import{_ as f,a as k}from"./uni-fab.3c991723.js";import{_ as w,a as _}from"./uni-list.23a0ebb9.js";import{r as b}from"./const.19faadc4.js";import{_ as y}from"./plugin-vue_export-helper.21dcd24c.js";import"./uni-icons.cae5d20c.js";var j=y({data:()=>({keyword:"",limit:20,datum:[]}),onPullDownRefresh(){this.datum=[],this.skip=0},onReachBottom(){this.load()},onLoad(){this.load()},methods:{load(){a({}),b({url:"plugin/search",method:"POST",data:{skip:this.datum.length,limit:this.limit,keyword:this.keyword?{id:this.keyword,name:this.keyword}:{}},success:a=>{this.datum=this.datum.concat(a)},complete(){t(),e()}})},onSearch(){this.datum=[],this.load()},create(){s({url:"./edit"})}}},[["render",function(a,t,e,s,b,y){const j=m(c("uni-search-bar"),f),g=m(c("uni-fab"),k),C=p,v=m(c("uni-list-item"),w),F=m(c("uni-list"),_),x=h;return o(),i(x,null,{default:d((()=>[l(j,{onConfirm:y.onSearch,onInput:t[0]||(t[0]=()=>{}),placeholder:"ID 名称",modelValue:b.keyword,"onUpdate:modelValue":t[1]||(t[1]=a=>b.keyword=a)},null,8,["onConfirm","modelValue"]),l(g,{onFabClick:y.create},null,8,["onFabClick"]),l(F,null,{default:d((()=>[(o(!0),n(u,null,r(b.datum,((a,t)=>(o(),i(v,{key:t,title:a.name,note:a.id,link:"",to:"./detail?id="+a.id},{header:d((()=>[l(C,{class:"icon",src:"/assets/plugin.3a38a152.svg",mode:"aspectFit"})])),_:2},1032,["title","note","to"])))),128))])),_:1})])),_:1})}],["__scopeId","data-v-03c85140"]]);export{j as default};
