(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-5986648f"],{2986:function(t,e,n){"use strict";n.r(e);var i=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",[n("div",{staticClass:"container"},[n("div",{staticClass:"form-box"},[n("el-form",{ref:"form",attrs:{model:t.ArticleInfo,"label-width":"80px"}},[n("el-form-item",{attrs:{label:"标题"}},[n("el-input",{model:{value:t.ArticleInfo.Title,callback:function(e){t.$set(t.ArticleInfo,"Title",e)},expression:"ArticleInfo.Title"}})],1),n("el-form-item",{attrs:{label:"文本框"}},[n("div",[n("tinymce",{attrs:{height:300},model:{value:t.ArticleInfo.HtmlContentt,callback:function(e){t.$set(t.ArticleInfo,"HtmlContentt",e)},expression:"ArticleInfo.HtmlContentt"}})],1)]),n("el-form-item",[n("el-button",{attrs:{type:"primary"},on:{click:t.onSubmit}},[t._v("表单提交")]),n("el-button",[t._v("取消")])],1)],1)],1)])])},o=[],a=n("751a"),r=n("8256"),l={name:"OnlyArticle14",components:{Tinymce:r["a"]},data:function(){return{ArticleInfo:{},query:{id:0},content:"",editorOption:{placeholder:"Hello World"}}},methods:{onSubmit:function(){var t=this;Object(a["a"])({url:"filmbll/savearticle",method:"post",data:this.ArticleInfo}).then((function(e){200==e.code&&t.$message.success("提交成功！")}))},getData:function(t){var e=this;this.query.id=t,Object(a["a"])({url:"filmbll/getarticle",method:"post",data:this.query}).then((function(t){200==t.code&&(console.log(t.obj),e.ArticleInfo=t.obj)}))}},created:function(){this.getData(14)},mounted:function(){}},s=l,c=n("2877"),u=Object(c["a"])(s,i,o,!1,null,"143caeee",null);e["default"]=u.exports},"3a75":function(t,e,n){},5278:function(t,e,n){"use strict";var i=n("3a75"),o=n.n(i);o.a},8256:function(t,e,n){"use strict";var i=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"tinymce-container",class:{fullscreen:t.fullscreen},style:{width:t.containerWidth}},[n("textarea",{staticClass:"tinymce-textarea",attrs:{id:t.tinymceId}}),n("div",{staticClass:"editor-custom-btn-container"},[n("editorImage",{staticClass:"editor-upload-btn",attrs:{color:"#1890ff"},on:{successCBK:t.imageSuccessCBK}})],1)])},o=[],a=(n("ac6a"),n("c5f6"),function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"upload-container"},[n("el-button",{style:{background:t.color,borderColor:t.color},attrs:{icon:"el-icon-upload",size:"mini",type:"primary"},on:{click:function(e){t.dialogVisible=!0}}},[t._v("\n    upload\n  ")]),n("el-dialog",{attrs:{visible:t.dialogVisible},on:{"update:visible":function(e){t.dialogVisible=e}}},[n("el-upload",{staticClass:"editor-slide-upload",attrs:{multiple:!0,"file-list":t.fileList,"show-file-list":!0,"on-remove":t.handleRemove,"on-success":t.handleSuccess,"before-upload":t.beforeUpload,action:"https://httpbin.org/post","list-type":"picture-card"}},[n("el-button",{attrs:{size:"small",type:"primary"}},[t._v("\n        Click upload\n      ")])],1),n("el-button",{on:{click:function(e){t.dialogVisible=!1}}},[t._v("\n      Cancel\n    ")]),n("el-button",{attrs:{type:"primary"},on:{click:t.handleSubmit}},[t._v("\n      Confirm\n    ")])],1)],1)}),r=[],l=(n("456d"),{name:"EditorSlideUpload",props:{color:{type:String,default:"#1890ff"}},data:function(){return{dialogVisible:!1,listObj:{},fileList:[]}},methods:{checkAllSuccess:function(){var t=this;return Object.keys(this.listObj).every((function(e){return t.listObj[e].hasSuccess}))},handleSubmit:function(){var t=this,e=Object.keys(this.listObj).map((function(e){return t.listObj[e]}));this.checkAllSuccess()?(this.$emit("successCBK",e),this.listObj={},this.fileList=[],this.dialogVisible=!1):this.$message("Please wait for all images to be uploaded successfully. If there is a network problem, please refresh the page and upload again!")},handleSuccess:function(t,e){for(var n=e.uid,i=Object.keys(this.listObj),o=0,a=i.length;o<a;o++)if(this.listObj[i[o]].uid===n)return this.listObj[i[o]].url=t.files.file,void(this.listObj[i[o]].hasSuccess=!0)},handleRemove:function(t){for(var e=t.uid,n=Object.keys(this.listObj),i=0,o=n.length;i<o;i++)if(this.listObj[n[i]].uid===e)return void delete this.listObj[n[i]]},beforeUpload:function(t){var e=this,n=window.URL||window.webkitURL,i=t.uid;return this.listObj[i]={},new Promise((function(o,a){var r=new Image;r.src=n.createObjectURL(t),r.onload=function(){e.listObj[i]={hasSuccess:!1,uid:t.uid,width:this.width,height:this.height}},o(!0)}))}}}),s=l,c=(n("5278"),n("2877")),u=Object(c["a"])(s,a,r,!1,null,"3dae379b",null),d=u.exports,f=["advlist anchor autolink autosave code codesample colorpicker colorpicker contextmenu directionality emoticons fullscreen hr image imagetools insertdatetime link lists media nonbreaking noneditable pagebreak paste preview print save searchreplace spellchecker tabfocus table template textcolor textpattern visualblocks visualchars wordcount"],h=f,m=["searchreplace bold italic underline strikethrough alignleft aligncenter alignright outdent indent  blockquote undo redo removeformat subscript superscript code codesample","hr bullist numlist link image charmap preview anchor pagebreak insertdatetime media table emoticons forecolor backcolor fullscreen"],b=m;n("ac4d"),n("8a81"),n("5df3"),n("1c4c"),n("7f7f"),n("6b54");function p(t,e){var n;if("undefined"===typeof Symbol||null==t[Symbol.iterator]){if(Array.isArray(t)||(n=y(t))||e&&t&&"number"===typeof t.length){n&&(t=n);var i=0,o=function(){};return{s:o,n:function(){return i>=t.length?{done:!0}:{done:!1,value:t[i++]}},e:function(t){throw t},f:o}}throw new TypeError("Invalid attempt to iterate non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}var a,r=!0,l=!1;return{s:function(){n=t[Symbol.iterator]()},n:function(){var t=n.next();return r=t.done,t},e:function(t){l=!0,a=t},f:function(){try{r||null==n.return||n.return()}finally{if(l)throw a}}}}function y(t,e){if(t){if("string"===typeof t)return g(t,e);var n=Object.prototype.toString.call(t).slice(8,-1);return"Object"===n&&t.constructor&&(n=t.constructor.name),"Map"===n||"Set"===n?Array.from(t):"Arguments"===n||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n)?g(t,e):void 0}}function g(t,e){(null==e||e>t.length)&&(e=t.length);for(var n=0,i=new Array(e);n<e;n++)i[n]=t[n];return i}var v=[];function w(){return window.tinymce}var _=function(t,e){var n=document.getElementById(t),i=e||function(){};if(!n){var o=document.createElement("script");o.src=t,o.id=t,document.body.appendChild(o),v.push(i);var a="onload"in o?r:l;a(o)}function r(e){e.onload=function(){this.onerror=this.onload=null;var t,n=p(v);try{for(n.s();!(t=n.n()).done;){var i=t.value;i(null,e)}}catch(o){n.e(o)}finally{n.f()}v=null},e.onerror=function(){this.onerror=this.onload=null,i(new Error("Failed to load "+t),e)}}function l(t){t.onreadystatechange=function(){if("complete"===this.readyState||"loaded"===this.readyState){this.onreadystatechange=null;var e,n=p(v);try{for(n.s();!(e=n.n()).done;){var i=e.value;i(null,t)}}catch(o){n.e(o)}finally{n.f()}v=null}}}n&&i&&(w()?i(null,n):v.push(i))},k=_,C="https://cdn.jsdelivr.net/npm/tinymce-all-in-one@4.9.3/tinymce.min.js",j={name:"Tinymce",components:{editorImage:d},props:{id:{type:String,default:function(){return"vue-tinymce-"+ +new Date+(1e3*Math.random()).toFixed(0)}},value:{type:String,default:""},inivalue:{type:String,default:""},toolbar:{type:Array,required:!1,default:function(){return[]}},menubar:{type:String,default:"file edit insert view format table"},height:{type:[Number,String],required:!1,default:360},width:{type:[Number,String],required:!1,default:"auto"}},data:function(){return{hasChange:!1,hasInit:!1,tinymceId:this.id,fullscreen:!1,languageTypeList:{en:"en",zh:"zh_CN",es:"es_MX",ja:"ja"}}},computed:{containerWidth:function(){var t=this.width;return/^[\d]+(\.[\d]+)?$/.test(t)?"".concat(t,"px"):t}},watch:{value:function(t){var e=this;!this.hasChange&&this.hasInit&&this.$nextTick((function(){return window.tinymce.get(e.tinymceId).setContent(t||"")}))},inivalue:function(t){window.tinymce.get(this.tinymceId).setContent(t||"")}},mounted:function(){this.init()},activated:function(){window.tinymce&&this.initTinymce()},deactivated:function(){this.destroyTinymce()},destroyed:function(){this.destroyTinymce()},methods:{init:function(){var t=this;k(C,(function(e){e?t.$message.error(e.message):t.initTinymce()}))},initTinymce:function(){var t=this,e=this;window.tinymce.init({selector:"#".concat(this.tinymceId),language:this.languageTypeList["zh"],height:this.height,body_class:"panel-body ",object_resizing:!1,toolbar:this.toolbar.length>0?this.toolbar:b,menubar:this.menubar,plugins:h,end_container_on_empty_block:!0,powerpaste_word_import:"clean",code_dialog_height:450,code_dialog_width:1e3,advlist_bullet_styles:"square",advlist_number_styles:"default",imagetools_cors_hosts:["www.tinymce.com","codepen.io"],default_link_target:"_blank",link_title:!1,nonbreaking_force_tab:!0,init_instance_callback:function(n){e.value&&n.setContent(e.value),e.hasInit=!0,n.on("NodeChange Change KeyUp SetContent",(function(){t.hasChange=!0,t.$emit("input",n.getContent())}))},setup:function(t){t.on("FullscreenStateChanged",(function(t){e.fullscreen=t.state}))},convert_urls:!1,images_upload_handler:function(t,e,n,i){var o="data:image/jpeg;base64,"+t.base64();e(o)}})},destroyTinymce:function(){var t=window.tinymce.get(this.tinymceId);this.fullscreen&&t.execCommand("mceFullScreen"),t&&t.destroy()},setContent:function(t){window.tinymce.get(this.tinymceId).setContent(t)},getContent:function(){window.tinymce.get(this.tinymceId).getContent()},imageSuccessCBK:function(t){var e=this;t.forEach((function(t){window.tinymce.get(e.tinymceId).insertContent('<img class="wscnph" src="'.concat(t.url,'" >'))}))}}},S=j,I=(n("d698"),Object(c["a"])(S,i,o,!1,null,"755949c2",null));e["a"]=I.exports},cef9:function(t,e,n){},d698:function(t,e,n){"use strict";var i=n("cef9"),o=n.n(i);o.a}}]);
//# sourceMappingURL=chunk-5986648f.f1ec18e3.js.map