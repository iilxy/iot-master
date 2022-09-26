import{b as e,_ as a,a as t}from"./uni-forms.3c0cbf84.js";import{o as l,c as n,w as d,a as o,b as i,r as u,e as s,f as c,d as r}from"./index.8d4a8305.js";import{r as m}from"./const.9e5d4232.js";import{_ as p}from"./plugin-vue_export-helper.21dcd24c.js";import"./uni-icons.51822eb1.js";var f=p({data:()=>({id:"",data:{name:"",station:1,tunnel_id:"",product_id:"",interface_id:""}}),onLoad(e){this.id=e.id,this.id&&this.load()},methods:{load(){m({url:"device/"+this.id,success:e=>{this.data=e}})},save(){uni.showLoading({}),m({url:"device/"+(this.id||"create"),method:"POST",data:this.data,success(){uni.navigateBack()},complete(){uni.hideLoading()}})},selectTunnel(){uni.navigateTo({url:"/pages/tunnel/tunnel?select=true",events:{select:e=>{this.data.tunnel_id=e}}})},selectProduct(){uni.navigateTo({url:"/pages/product/product?select=true",events:{select:e=>{this.data.product_id=e}}})},selectInterface(){uni.navigateTo({url:"/pages/interface/interface?select=true",events:{select:e=>{this.data.interface_id=e}}})}}},[["render",function(m,p,f,_,h,V){const v=u(s("uni-easyinput"),e),I=u(s("uni-forms-item"),a),b=u(s("uni-forms"),t),g=c,k=r;return l(),n(k,{class:"p20"},{default:d((()=>[o(b,null,{default:d((()=>[o(I,{label:"名称",name:"name"},{default:d((()=>[o(v,{modelValue:h.data.name,"onUpdate:modelValue":p[0]||(p[0]=e=>h.data.name=e),placeholder:""},null,8,["modelValue"])])),_:1}),o(I,{label:"站号",name:"station"},{default:d((()=>[o(v,{type:"number",modelValue:h.data.station,"onUpdate:modelValue":p[1]||(p[1]=e=>h.data.station=e),placeholder:"从机号"},null,8,["modelValue"])])),_:1}),o(I,{label:"通道",name:"tunnel_id"},{default:d((()=>[o(v,{modelValue:h.data.tunnel_id,"onUpdate:modelValue":p[2]||(p[2]=e=>h.data.tunnel_id=e),placeholder:"点击右边小图标选择",suffixIcon:"plus",onIconClick:V.selectTunnel},null,8,["modelValue","onIconClick"])])),_:1}),o(I,{label:"产品",name:"product_id"},{default:d((()=>[o(v,{modelValue:h.data.product_id,"onUpdate:modelValue":p[3]||(p[3]=e=>h.data.product_id=e),placeholder:"点击右边小图标选择",suffixIcon:"plus",onIconClick:V.selectProduct},null,8,["modelValue","onIconClick"])])),_:1}),o(I,{label:"组态",name:"interface_id"},{default:d((()=>[o(v,{modelValue:h.data.interface_id,"onUpdate:modelValue":p[4]||(p[4]=e=>h.data.interface_id=e),placeholder:"点击右边小图标选择",suffixIcon:"plus",onIconClick:V.selectInterface},null,8,["modelValue","onIconClick"])])),_:1})])),_:1}),o(g,{type:"primary",onClick:V.save},{default:d((()=>[i("保存")])),_:1},8,["onClick"])])),_:1})}]]);export{f as default};
