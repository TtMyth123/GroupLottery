(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-63876a20"],{"058a":function(e,a,t){"use strict";var l=t("80d3"),r=t.n(l);r.a},"15eb":function(e,a,t){"use strict";t.r(a);var l=function(){var e=this,a=e.$createElement,t=e._self._c||a;return t("div",[t("div",{staticClass:"crumbs"},[t("el-breadcrumb",{attrs:{separator:"/"}},[t("el-breadcrumb-item",[t("i",{staticClass:"el-icon-lx-cascades"}),e._v(" 电影订单\n            ")])],1)],1),t("div",{staticClass:"container"},[t("div",{staticClass:"handle-box"},[t("el-select",{attrs:{placeholder:"请选择"},model:{value:e.query.userType,callback:function(a){e.$set(e.query,"userType",e._n(a))},expression:"query.userType"}},e._l(e.UserTypeAll,(function(e){return t("el-option",{key:e.value,attrs:{label:e.name,value:e.value}})})),1),t("el-date-picker",{staticClass:"handle-input-min",attrs:{type:"date","value-format":"yyyy-MM-dd",placeholder:"选择开始日期"},model:{value:e.query.beginDay,callback:function(a){e.$set(e.query,"beginDay",a)},expression:"query.beginDay"}}),t("el-date-picker",{staticClass:"handle-input-min",attrs:{type:"date","value-format":"yyyy-MM-dd",placeholder:"选择结束日期"},model:{value:e.query.endDay,callback:function(a){e.$set(e.query,"endDay",a)},expression:"query.endDay"}}),t("el-input",{staticClass:"handle-input mr10",attrs:{placeholder:"电影名称"},model:{value:e.query.filmName,callback:function(a){e.$set(e.query,"filmName",a)},expression:"query.filmName"}}),t("el-input",{staticClass:"handle-input mr10",attrs:{placeholder:"用户名"},model:{value:e.query.userName,callback:function(a){e.$set(e.query,"userName",a)},expression:"query.userName"}}),t("el-button",{attrs:{type:"primary",icon:"el-icon-search"},on:{click:e.handleSearch}},[e._v("搜索")])],1),t("el-table",{ref:"multipleTable",staticClass:"table",attrs:{data:e.tableData,border:"","header-cell-class-name":"table-header","summary-method":e.getSummaries,"show-summary":""}},[t("el-table-column",{attrs:{prop:"Id",label:"ID",width:"55",align:"center"}}),t("el-table-column",{attrs:{prop:"UserName",label:"用户名",width:"100"}}),t("el-table-column",{attrs:{prop:"FullName",label:"姓名"}}),t("el-table-column",{attrs:{label:"用户类别"},scopedSlots:e._u([{key:"default",fn:function(a){return[e._v(e._s(e._f("UserTypeName")(a.row.UserType)))]}}])}),t("el-table-column",{attrs:{prop:"FilmName",label:"电影名称"}}),t("el-table-column",{attrs:{prop:"SettlementModesName",label:"结算方式"}}),t("el-table-column",{attrs:{label:"级别"},scopedSlots:e._u([{key:"default",fn:function(a){return[e._v(e._s(e._f("MemberLevelName")(a.row.BuyLevel)))]}}])}),t("el-table-column",{attrs:{label:"状态"},scopedSlots:e._u([{key:"default",fn:function(a){return[e._v(e._s(e._f("FilmOrderStateName")(a.row.State)))]}}])}),t("el-table-column",{attrs:{label:"投资金额"},scopedSlots:e._u([{key:"default",fn:function(a){return[e._v(e._s(e._f("money")(a.row.BuyGold)))]}}])}),t("el-table-column",{attrs:{prop:"DailyRate",label:"日化率",width:"150"},scopedSlots:e._u([{key:"default",fn:function(a){return[e._v(e._s(e._f("money")(100*a.row.DailyRate))+"%+"+e._s(e._f("money4")(100*a.row.AddDailyRate))+"%")]}}])}),t("el-table-column",{attrs:{label:"收益金额"},scopedSlots:e._u([{key:"default",fn:function(a){return[e._v(e._s(e._f("money")(a.row.CurSumRebateGold)))]}}])}),t("el-table-column",{attrs:{label:"收益次数"},scopedSlots:e._u([{key:"default",fn:function(a){return[e._v(e._s(a.row.CurCycleCount))]}}])}),t("el-table-column",{attrs:{label:"预估的分红"},scopedSlots:e._u([{key:"default",fn:function(a){return[e._v(e._s(e._f("money")(a.row.EstimateRebateGold)))]}}])}),t("el-table-column",{attrs:{prop:"CreatedAt",label:"购买时间"}})],1),t("div",{staticClass:"pagination"},[t("el-pagination",{attrs:{background:"",layout:"total, prev, pager, next","current-page":e.query.pageIndex,"page-size":e.query.pageSize,total:e.pageTotal},on:{"current-change":e.handlePageChange}})],1)],1),t("el-dialog",{attrs:{title:"裁剪图片",visible:e.dialogVisible,width:"30%"},on:{"update:visible":function(a){e.dialogVisible=a}}},[t("vue-cropper",{ref:"cropper",staticStyle:{width:"100%",height:"300px"},attrs:{src:e.imgSrc,ready:e.cropImage,zoom:e.cropImage,cropmove:e.cropImage}}),t("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[t("el-button",{on:{click:e.cancelCrop}},[e._v("取 消")]),t("el-button",{attrs:{type:"primary"},on:{click:function(a){e.dialogVisible=!1}}},[e._v("确 定")])],1)],1)],1)},r=[],n=t("95c3"),o=t.n(n),s=t("751a"),u={name:"FilmOrderList",data:function(){return{defaultSrc:t("7159"),fileList:[],imgSrc:"",cropImg:"",dialogVisible:!1,FilmStateAll:[{name:"全部",value:0},{name:"准备",value:1},{name:"上架",value:2},{name:"下架",value:3}],FilmStates:[{name:"准备",value:1},{name:"上架",value:2},{name:"下架",value:3}],BuyLevelTypeAll:[{name:"全部",value:0},{name:"新手会员",value:1},{name:"白银会员",value:2},{name:"黄金会员",value:3},{name:"钻石会员",value:4},{name:"皇冠会员",value:5}],BuyLevelTypes:[{name:"新手会员",value:1},{name:"白银会员",value:2},{name:"黄金会员",value:3},{name:"钻石会员",value:4},{name:"皇冠会员",value:5}],UserTypeAll:[{name:"全部",value:0},{name:"一般玩家",value:1},{name:"业务玩家",value:2}],query:{userType:0,beginDay:"",endDay:"",buyLevel:2,filmName:"",userName:"",pageIndex:1,pageSize:10},tableData:[],GroupData:{},multipleSelection:[],delList:[],editVisible:!1,pageTotal:0,form:{},idx:-1,id:-1}},created:function(){this.getData()},components:{VueCropper:o.a},methods:{getSummaries:function(){var e=[];return e[1]="合计："+this.GroupData.C+"条",e[8]=this.GroupData.BuyGold,e[10]=this.GroupData.ArriveRebateGold.toFixed(2),e[12]=this.GroupData.EstimateRebateGold.toFixed(2),e},getData:function(){var e=this;Object(s["a"])({url:"filmbll/getfilmorderlist",method:"post",data:this.query}).then((function(a){200==a.code&&(console.log(a.obj.ListData),e.tableData=a.obj.ListData,e.pageTotal=a.obj.PageTotal,e.GroupData=a.obj.GroupData)}))},handleSearch:function(){this.$set(this.query,"pageIndex",1),this.getData()},handlePageChange:function(e){this.$set(this.query,"pageIndex",e),this.getData()}}},i=u,c=(t("058a"),t("2877")),m=Object(c["a"])(i,l,r,!1,null,"07c40cf2",null);a["default"]=m.exports},7159:function(e,a,t){e.exports=t.p+"static/img/img.146655c9.jpg"},"80d3":function(e,a,t){}}]);
//# sourceMappingURL=chunk-63876a20.119a9a77.js.map