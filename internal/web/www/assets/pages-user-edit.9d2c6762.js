import{b as a,_ as e,a as l}from"./uni-forms.3c0cbf84.js";import{o as d,c as s,w as i,a as t,b as n,r as o,e as u,f as m,d as r}from"./index.8d4a8305.js";import{r as c}from"./const.9e5d4232.js";import{_ as p}from"./plugin-vue_export-helper.21dcd24c.js";import"./uni-icons.51822eb1.js";var f=p({data:()=>({id:0,data:{username:"",nickname:"",email:""}}),onLoad(a){this.id=a.id,this.id&&this.load()},methods:{load(){c({url:"user/"+this.id,success:a=>{this.data=a}})},save(){uni.showLoading({}),c({url:"user/"+(this.id||"create"),method:"POST",data:this.data,success(){uni.navigateBack()},complete(){uni.hideLoading()}})}}},[["render",function(c,p,f,h,V,_){const b=o(u("uni-easyinput"),a),k=o(u("uni-forms-item"),e),j=o(u("uni-forms"),l),v=m,g=r;return d(),s(g,{class:"p20"},{default:i((()=>[t(j,null,{default:i((()=>[t(k,{label:"用户名",name:"username"},{default:i((()=>[t(b,{modelValue:V.data.username,"onUpdate:modelValue":p[0]||(p[0]=a=>V.data.username=a),placeholder:""},null,8,["modelValue"])])),_:1}),t(k,{label:"昵称",name:"nickname"},{default:i((()=>[t(b,{modelValue:V.data.nickname,"onUpdate:modelValue":p[1]||(p[1]=a=>V.data.nickname=a),placeholder:""},null,8,["modelValue"])])),_:1}),t(k,{label:"邮箱",name:"email"},{default:i((()=>[t(b,{modelValue:V.data.email,"onUpdate:modelValue":p[2]||(p[2]=a=>V.data.email=a),placeholder:""},null,8,["modelValue"])])),_:1})])),_:1}),t(v,{type:"primary",onClick:_.save},{default:i((()=>[n("保存")])),_:1},8,["onClick"])])),_:1})}]]);export{f as default};
