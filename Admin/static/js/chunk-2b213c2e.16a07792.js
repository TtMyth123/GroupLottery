(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-2b213c2e"],{"361e":function(e,t,a){"use strict";a.r(t);var l=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",[a("div",{staticClass:"crumbs"},[a("el-breadcrumb",{attrs:{separator:"/"}},[a("el-breadcrumb-item",[a("i",{staticClass:"el-icon-lx-cascades"}),e._v(" 电影订单\n            ")])],1)],1),a("div",{staticClass:"container"},[a("div",{staticClass:"handle-box"},[a("el-select",{attrs:{placeholder:"请选择"},model:{value:e.query.buyLevel,callback:function(t){e.$set(e.query,"buyLevel",e._n(t))},expression:"query.buyLevel"}},e._l(e.BuyLevelTypeAll,(function(e){return a("el-option",{key:e.value,attrs:{label:e.name,value:e.value}})})),1),a("el-select",{attrs:{placeholder:"请选择状态"},model:{value:e.query.state,callback:function(t){e.$set(e.query,"state",e._n(t))},expression:"query.state"}},e._l(e.StateAll,(function(e){return a("el-option",{key:e.value,attrs:{label:e.name,value:e.value}})})),1),a("el-select",{attrs:{placeholder:"请选择"},model:{value:e.query.userType,callback:function(t){e.$set(e.query,"userType",e._n(t))},expression:"query.userType"}},e._l(e.UserTypeAll,(function(e){return a("el-option",{key:e.value,attrs:{label:e.name,value:e.value}})})),1),a("el-date-picker",{staticClass:"handle-input-min",attrs:{type:"date","value-format":"yyyy-MM-dd",placeholder:"选择开始日期"},model:{value:e.query.beginDay,callback:function(t){e.$set(e.query,"beginDay",t)},expression:"query.beginDay"}}),a("el-date-picker",{staticClass:"handle-input-min",attrs:{type:"date","value-format":"yyyy-MM-dd",placeholder:"选择结束日期"},model:{value:e.query.endDay,callback:function(t){e.$set(e.query,"endDay",t)},expression:"query.endDay"}}),a("el-input",{staticClass:"handle-input mr10",attrs:{placeholder:"电影名称"},model:{value:e.query.filmName,callback:function(t){e.$set(e.query,"filmName",t)},expression:"query.filmName"}}),a("el-input",{staticClass:"handle-input mr10",attrs:{placeholder:"用户名"},model:{value:e.query.userName,callback:function(t){e.$set(e.query,"userName",t)},expression:"query.userName"}}),a("el-button",{attrs:{type:"primary",icon:"el-icon-search"},on:{click:e.handleSearch}},[e._v("搜索")])],1),a("el-table",{ref:"multipleTable",staticClass:"table",attrs:{data:e.tableData,border:"","header-cell-class-name":"table-header","summary-method":e.getSummaries,"show-summary":""}},[a("el-table-column",{attrs:{prop:"Id",label:"ID",width:"55",align:"center"}}),a("el-table-column",{attrs:{prop:"UserName",label:"用户名",width:"100"}}),a("el-table-column",{attrs:{prop:"FullName",label:"姓名"}}),a("el-table-column",{attrs:{label:"用户类别"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(e._s(e._f("UserTypeName")(t.row.UserType)))]}}])}),a("el-table-column",{attrs:{prop:"FilmName",label:"电影名称"}}),a("el-table-column",{attrs:{prop:"SettlementModesName",label:"结算方式"}}),a("el-table-column",{attrs:{label:"级别"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(e._s(e._f("MemberLevelName")(t.row.BuyLevel)))]}}])}),a("el-table-column",{attrs:{label:"状态"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(e._s(e._f("FilmOrderStateName")(t.row.State)))]}}])}),a("el-table-column",{attrs:{label:"投资金额"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(e._s(e._f("money")(t.row.BuyGold)))]}}])}),a("el-table-column",{attrs:{prop:"DailyRate",label:"日化率",width:"150"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(e._s(e._f("money")(100*t.row.DailyRate))+"%+"+e._s(e._f("money4")(100*t.row.AddDailyRate))+"%")]}}])}),a("el-table-column",{attrs:{label:"收益金额"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(e._s(e._f("money")(t.row.CurSumRebateGold)))]}}])}),a("el-table-column",{attrs:{label:"收益次数"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(e._s(t.row.CurCycleCount))]}}])}),a("el-table-column",{attrs:{label:"预估的分红"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(e._s(e._f("money")(t.row.EstimateRebateGold)))]}}])}),a("el-table-column",{attrs:{prop:"CreatedAt",label:"购买时间"}}),a("el-table-column",{attrs:{prop:"EndTime",label:"到期时间"}})],1),a("div",{staticClass:"pagination"},[a("el-pagination",{attrs:{background:"",layout:"total, prev, pager, next","current-page":e.query.pageIndex,"page-size":e.query.pageSize,total:e.pageTotal},on:{"current-change":e.handlePageChange}})],1)],1),a("el-dialog",{attrs:{title:"裁剪图片",visible:e.dialogVisible,width:"30%"},on:{"update:visible":function(t){e.dialogVisible=t}}},[a("vue-cropper",{ref:"cropper",staticStyle:{width:"100%",height:"300px"},attrs:{src:e.imgSrc,ready:e.cropImage,zoom:e.cropImage,cropmove:e.cropImage}}),a("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[a("el-button",{on:{click:e.cancelCrop}},[e._v("取 消")]),a("el-button",{attrs:{type:"primary"},on:{click:function(t){e.dialogVisible=!1}}},[e._v("确 定")])],1)],1)],1)},r=[],s=(a("a481"),a("6762"),a("2fdb"),a("95c3")),n=a.n(s),o=a("751a"),i={name:"FilmOrderListEx",data:function(){return{defaultSrc:a("7159"),fileList:[],imgSrc:"",cropImg:"",dialogVisible:!1,BuyLevelTypeAll:[{name:"全部",value:0},{name:"新手会员",value:1},{name:"白银会员",value:2},{name:"黄金会员",value:3},{name:"钻石会员",value:4},{name:"皇冠会员",value:5}],BuyLevelTypes:[{name:"新手会员",value:1},{name:"白银会员",value:2},{name:"黄金会员",value:3},{name:"钻石会员",value:4},{name:"皇冠会员",value:5}],StateAll:[{name:"全部",value:0},{name:"行进中",value:1},{name:"完成",value:2}],UserTypeAll:[{name:"全部",value:0},{name:"一般玩家",value:1},{name:"业务玩家",value:2}],query:{userType:0,beginDay:"",endDay:"",buyLevel:0,filmName:"",userName:"",pageIndex:1,pageSize:10,state:0},tableData:[],GroupData:{},multipleSelection:[],delList:[],editVisible:!1,pageTotal:0,form:{},idx:-1,id:-1}},created:function(){this.getData()},components:{VueCropper:n.a},methods:{getSummaries:function(){var e=[];return e[1]="合计："+this.GroupData.C+"条",e[8]=this.GroupData.BuyGold,e[10]=this.GroupData.ArriveRebateGold.toFixed(2),e[12]=this.GroupData.EstimateRebateGold.toFixed(2),e},setImage:function(e){var t=this,a=e.target.files[0];if(a.type.includes("image/")){var l=new FileReader;l.onload=function(e){t.dialogVisible=!0,t.imgSrc=e.target.result,t.$refs.cropper&&t.$refs.cropper.replace(e.target.result)},l.readAsDataURL(a)}},cropImage:function(){this.cropImg=this.$refs.cropper.getCroppedCanvas().toDataURL()},cancelCrop:function(){this.dialogVisible=!1,this.cropImg=this.defaultSrc},getData:function(){var e=this;Object(o["a"])({url:"filmbll/getfilmorderlistex",method:"post",data:this.query}).then((function(t){200==t.code&&(console.log(t.obj.ListData),e.tableData=t.obj.ListData,e.pageTotal=t.obj.PageTotal,e.GroupData=t.obj.GroupData)}))},handleSearch:function(){this.$set(this.query,"pageIndex",1),this.getData()},handleDelete:function(e,t){var a=this;this.$confirm("确定要删除吗？","提示",{type:"warning"}).then((function(){var l={id:JSON.stringify(t.Id)};Object(o["a"])({url:"filmbll/delfilminfo",method:"post",data:l}).then((function(t){console.log(t),200==t.code?(a.$message.success("删除成功"),a.tableData.splice(e,1),a.getData()):a.$message.error("删除失败："+t.msg)}))})).catch((function(){}))},handleEdit:function(e,t){this.idx=e,this.form=t,this.editVisible=!0},handleAdd:function(){this.idx=-1;var e={};this.form=e,this.editVisible=!0},saveEdit:function(){var e=this;if(-1===this.idx){var t={jsonData:JSON.stringify(this.form)};Object(o["a"])({url:"filmbll/addfilminfo",method:"post",data:t}).then((function(t){console.log(t),200==t.code?(e.editVisible=!1,e.$message.success("添加成功"),e.getData()):e.$message.error("添加失败："+t.msg)}))}else{t={jsonData:JSON.stringify(this.form)};Object(o["a"])({url:"filmbll/updatefilm",method:"post",data:t}).then((function(t){console.log(t),200==t.code?(e.editVisible=!1,e.$message.success("修改第 ".concat(e.idx+1," 行成功")),e.$set(e.tableData,e.idx,e.form)):e.$message.error("更新失败："+t.msg)}))}},handlePageChange:function(e){this.$set(this.query,"pageIndex",e),this.getData()}}},u=i,c=(a("7a2d"),a("2877")),d=Object(c["a"])(u,l,r,!1,null,"070976ac",null);t["default"]=d.exports},7159:function(e,t,a){e.exports=a.p+"static/img/img.146655c9.jpg"},"7a2d":function(e,t,a){"use strict";var l=a("8a80"),r=a.n(l);r.a},"8a80":function(e,t,a){}}]);
//# sourceMappingURL=chunk-2b213c2e.16a07792.js.map