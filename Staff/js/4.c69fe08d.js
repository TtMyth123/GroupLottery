(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([[4],{6796:function(s,t,e){"use strict";e.r(t);var n=function(){var s=this,t=s.$createElement,e=s._self._c||t;return e("view",{staticClass:"cul-wrapper"},[e("scroll-view",{style:{height:s.scrollHeight-70+"px"},attrs:{"scroll-y":"","scroll-into-view":s.scrollIndex,"scroll-with-animation":""}},s._l(s.msgs,(function(t,n){return e("block",{key:n},[e("view",{class:[t.isme?"msg-me":"msg-service"],attrs:{id:"msg-"+n}},[t.isme?s._e():e("image",{staticClass:"avatar",attrs:{src:s.anotherAvatar,"lazy-load":""}}),e("view",{staticClass:"msg-text"},[e("view",{staticClass:"msg-text-content"},[e("text",[s._v(s._s(t.data))])])]),t.isme?e("image",{staticClass:"avatar",attrs:{src:s.meAvatar,"lazy-load":""}}):s._e()])])})),1),e("view",{staticClass:"operation",class:{conversion:!s.conversion}},[e("input",{directives:[{name:"model",rawName:"v-model",value:s.msgInput.msg1,expression:"msgInput.msg1"}],class:[s.conversion?"msg-input1":"msg-input2"],attrs:{type:"text"},domProps:{value:s.msgInput.msg1},on:{tap:function(t){return s.sub1(s.msgInput.msg2)},input:function(t){t.target.composing||s.$set(s.msgInput,"msg1",t.target.value)}}}),e("input",{directives:[{name:"model",rawName:"v-model",value:s.msgInput.msg2,expression:"msgInput.msg2"}],class:[s.conversion?"msg-input2":"msg-input1"],attrs:{type:"text"},domProps:{value:s.msgInput.msg2},on:{tap:function(t){return s.sub2(s.msgInput.msg1)},input:function(t){t.target.composing||s.$set(s.msgInput,"msg2",t.target.value)}}})])],1)},i=[],a={name:"CulChat",props:{scrollHeight:{type:Number,default(){let s=0;return uni.getSystemInfo({success(t){s=t.windowHeight}}),s}},meAvatar:{type:String},anotherAvatar:{type:String}},data(){return{conversion:!0,msgInput:{msg1:"",msg2:""},msgs:[],scrollIndex:""}},methods:{jumpScroll(){this.scrollIndex="",this.$nextTick((function(){this.scrollIndex="msg-"+(this.msgs.length-1)}))},sub1(s){this.conversion||this.sub(s)},sub2(s){this.conversion&&this.sub(s)},sub(s){let t=this;if(s){this.msgInput.msg1="",this.msgInput.msg2="";let e={isme:!0,data:s};this.msgs.push(e),this.conversion=!this.conversion,this.jumpScroll(),this.$emit("submit",{message:s,callback:s=>{let e={isme:!1,data:s};t.msgs.push(e),t.jumpScroll()}})}else this.conversion=!this.conversion,this.jumpScroll()}}},o=a,r=(e("9fce"),e("2877")),m=Object(r["a"])(o,n,i,!1,null,"ee33b64a",null);t["default"]=m.exports},"9fce":function(s,t,e){"use strict";e("d94a")},d94a:function(s,t,e){}}]);