(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-1a2beaad"],{7159:function(t,e,a){t.exports=a.p+"static/img/img.146655c9.jpg"},"7bb8":function(t,e,a){"use strict";var l=a("80d8"),i=a.n(l);i.a},"80d8":function(t,e,a){},c6b5:function(t,e,a){"use strict";a.r(e);var l=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",[a("div",{staticClass:"container"},[a("div",{staticClass:"handle-box"},[a("el-button",{staticClass:"handle-del mr10",attrs:{type:"primary",icon:"el-icon-delete"},on:{click:function(e){t.testAwardVisible=!0}}},[t._v("添加测试结果\n        ")]),a("el-date-picker",{staticClass:"handle-input-min",attrs:{type:"date","value-format":"yyyy-MM-dd",placeholder:"选择开始日期"},model:{value:t.query.beginDay,callback:function(e){t.$set(t.query,"beginDay",e)},expression:"query.beginDay"}}),a("el-date-picker",{staticClass:"handle-input-min",attrs:{type:"date","value-format":"yyyy-MM-dd",placeholder:"选择结束日期"},model:{value:t.query.endDay,callback:function(e){t.$set(t.query,"endDay",e)},expression:"query.endDay"}}),a("el-button",{attrs:{type:"primary",icon:"el-icon-search"},on:{click:t.handleSearch}},[t._v("搜索")])],1),a("el-table",{ref:"multipleTable",staticClass:"table",attrs:{data:t.tableData,border:"","header-cell-class-name":"table-header"}},[a("el-table-column",{attrs:{prop:"Id",label:"ID",width:"55",align:"center"}}),a("el-table-column",{attrs:{prop:"GameName",label:"彩种",width:"100"}}),a("el-table-column",{attrs:{prop:"LotteryStr",label:"期号"}}),a("el-table-column",{attrs:{prop:"ResultNums",label:"开奖结果"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("el-button",{attrs:{type:"text"},on:{click:function(a){return t.openDetail(e.$index,e.row)}}},[t._v(t._s(e.row.ResultNums)+"\n            ")])]}}])}),a("el-table-column",{attrs:{prop:"CurLotteryTime",label:"开期时间"}})],1),a("div",{staticClass:"pagination"},[a("el-pagination",{attrs:{background:"",layout:"total, prev, pager, next","current-page":t.query.pageIndex,"page-size":t.query.pageSize,total:t.pageTotal},on:{"current-change":t.handlePageChange}})],1)],1),a("el-dialog",{attrs:{title:"明细",visible:t.detailVisible,width:"90%"},on:{"update:visible":function(e){t.detailVisible=e}}},[a("el-form",{ref:"form",attrs:{inline:!0,model:t.tableDataDetail,"label-width":"100px"}},[a("el-form-item",{attrs:{label:"采种",width:"100"}},[a("el-input",{model:{value:t.tableDataDetail.GameName,callback:function(e){t.$set(t.tableDataDetail,"GameName",e)},expression:"tableDataDetail.GameName"}})],1),a("el-form-item",{attrs:{label:"期号",width:"100"}},[a("el-input",{model:{value:t.tableDataDetail.LotteryStr,callback:function(e){t.$set(t.tableDataDetail,"LotteryStr",e)},expression:"tableDataDetail.LotteryStr"}})],1),a("el-divider"),a("el-form-item",{attrs:{label:"Jackpots"}},[a("el-input",{model:{value:t.tableDataDetail.Result.jackpots,callback:function(e){t.$set(t.tableDataDetail.Result,"jackpots",e)},expression:"tableDataDetail.Result.jackpots"}})],1),a("el-divider"),a("el-form-item",{attrs:{label:"FirstNum"}},[a("el-input",{model:{value:t.tableDataDetail.Result.firstNum,callback:function(e){t.$set(t.tableDataDetail.Result,"firstNum",e)},expression:"tableDataDetail.Result.firstNum"}})],1),a("el-divider"),a("el-form-item",{attrs:{label:"SecondNum"}},[a("el-input",{model:{value:t.tableDataDetail.Result.secondNum[0],callback:function(e){t.$set(t.tableDataDetail.Result.secondNum,0,e)},expression:"tableDataDetail.Result.secondNum[0]"}})],1),a("el-form-item",{attrs:{label:""}},[a("el-input",{model:{value:t.tableDataDetail.Result.secondNum[1],callback:function(e){t.$set(t.tableDataDetail.Result.secondNum,1,e)},expression:"tableDataDetail.Result.secondNum[1]"}})],1),a("el-divider"),a("el-form-item",{attrs:{label:"ThirdNum"}},[a("el-input",{model:{value:t.tableDataDetail.Result.thirdNum[0],callback:function(e){t.$set(t.tableDataDetail.Result.thirdNum,0,e)},expression:"tableDataDetail.Result.thirdNum[0]"}})],1),a("el-form-item",{attrs:{label:""}},[a("el-input",{model:{value:t.tableDataDetail.Result.thirdNum[1],callback:function(e){t.$set(t.tableDataDetail.Result.thirdNum,1,e)},expression:"tableDataDetail.Result.thirdNum[1]"}})],1),a("el-form-item",{attrs:{label:""}},[a("el-input",{model:{value:t.tableDataDetail.Result.thirdNum[2],callback:function(e){t.$set(t.tableDataDetail.Result.thirdNum,2,e)},expression:"tableDataDetail.Result.thirdNum[2]"}})],1),a("el-form-item",{attrs:{label:""}},[a("el-input",{model:{value:t.tableDataDetail.Result.thirdNum[3],callback:function(e){t.$set(t.tableDataDetail.Result.thirdNum,3,e)},expression:"tableDataDetail.Result.thirdNum[3]"}})],1),a("el-form-item",{attrs:{label:""}},[a("el-input",{model:{value:t.tableDataDetail.Result.thirdNum[4],callback:function(e){t.$set(t.tableDataDetail.Result.thirdNum,4,e)},expression:"tableDataDetail.Result.thirdNum[4]"}})],1),a("el-form-item",{attrs:{label:""}},[a("el-input",{model:{value:t.tableDataDetail.Result.thirdNum[5],callback:function(e){t.$set(t.tableDataDetail.Result.thirdNum,5,e)},expression:"tableDataDetail.Result.thirdNum[5]"}})],1),a("el-divider"),a("el-form-item",{attrs:{label:"forthNum"}},[a("el-input",{model:{value:t.tableDataDetail.Result.forthNum[0],callback:function(e){t.$set(t.tableDataDetail.Result.forthNum,0,e)},expression:"tableDataDetail.Result.forthNum[0]"}})],1),a("el-form-item",{attrs:{label:""}},[a("el-input",{model:{value:t.tableDataDetail.Result.forthNum[1],callback:function(e){t.$set(t.tableDataDetail.Result.forthNum,1,e)},expression:"tableDataDetail.Result.forthNum[1]"}})],1),a("el-form-item",{attrs:{label:""}},[a("el-input",{model:{value:t.tableDataDetail.Result.forthNum[2],callback:function(e){t.$set(t.tableDataDetail.Result.forthNum,2,e)},expression:"tableDataDetail.Result.forthNum[2]"}})],1),a("el-form-item",{attrs:{label:""}},[a("el-input",{model:{value:t.tableDataDetail.Result.forthNum[3],callback:function(e){t.$set(t.tableDataDetail.Result.forthNum,3,e)},expression:"tableDataDetail.Result.forthNum[3]"}})],1),a("el-form-item",{attrs:{label:""}},[a("el-input",{model:{value:t.tableDataDetail.Result.forthNum[4],callback:function(e){t.$set(t.tableDataDetail.Result.forthNum,4,e)},expression:"tableDataDetail.Result.forthNum[4]"}})],1),a("el-divider"),a("el-form-item",{attrs:{label:"fifthNum"}},[a("el-input",{model:{value:t.tableDataDetail.Result.fifthNum[0],callback:function(e){t.$set(t.tableDataDetail.Result.fifthNum,0,e)},expression:"tableDataDetail.Result.fifthNum[0]"}})],1),a("el-form-item",{attrs:{label:""}},[a("el-input",{model:{value:t.tableDataDetail.Result.fifthNum[1],callback:function(e){t.$set(t.tableDataDetail.Result.fifthNum,1,e)},expression:"tableDataDetail.Result.fifthNum[1]"}})],1),a("el-form-item",{attrs:{label:""}},[a("el-input",{model:{value:t.tableDataDetail.Result.fifthNum[2],callback:function(e){t.$set(t.tableDataDetail.Result.fifthNum,2,e)},expression:"tableDataDetail.Result.fifthNum[2]"}})],1),a("el-form-item",{attrs:{label:""}},[a("el-input",{model:{value:t.tableDataDetail.Result.fifthNum[3],callback:function(e){t.$set(t.tableDataDetail.Result.fifthNum,3,e)},expression:"tableDataDetail.Result.fifthNum[3]"}})],1),a("el-form-item",{attrs:{label:""}},[a("el-input",{model:{value:t.tableDataDetail.Result.fifthNum[4],callback:function(e){t.$set(t.tableDataDetail.Result.fifthNum,4,e)},expression:"tableDataDetail.Result.fifthNum[4]"}})],1),a("el-form-item",{attrs:{label:""}},[a("el-input",{model:{value:t.tableDataDetail.Result.fifthNum[5],callback:function(e){t.$set(t.tableDataDetail.Result.fifthNum,5,e)},expression:"tableDataDetail.Result.fifthNum[5]"}})],1),a("el-divider"),a("el-form-item",{attrs:{label:"sixthNum"}},[a("el-input",{model:{value:t.tableDataDetail.Result.sixthNum[0],callback:function(e){t.$set(t.tableDataDetail.Result.sixthNum,0,e)},expression:"tableDataDetail.Result.sixthNum[0]"}})],1),a("el-form-item",{attrs:{label:""}},[a("el-input",{model:{value:t.tableDataDetail.Result.sixthNum[1],callback:function(e){t.$set(t.tableDataDetail.Result.sixthNum,1,e)},expression:"tableDataDetail.Result.sixthNum[1]"}})],1),a("el-form-item",{attrs:{label:""}},[a("el-input",{model:{value:t.tableDataDetail.Result.sixthNum[2],callback:function(e){t.$set(t.tableDataDetail.Result.sixthNum,2,e)},expression:"tableDataDetail.Result.sixthNum[2]"}})],1),a("el-divider"),a("el-form-item",{attrs:{label:"seventhNum"}},[a("el-input",{model:{value:t.tableDataDetail.Result.seventhNum[0],callback:function(e){t.$set(t.tableDataDetail.Result.seventhNum,0,e)},expression:"tableDataDetail.Result.seventhNum[0]"}})],1),a("el-form-item",{attrs:{label:""}},[a("el-input",{model:{value:t.tableDataDetail.Result.seventhNum[1],callback:function(e){t.$set(t.tableDataDetail.Result.seventhNum,1,e)},expression:"tableDataDetail.Result.seventhNum[1]"}})],1),a("el-form-item",{attrs:{label:""}},[a("el-input",{model:{value:t.tableDataDetail.Result.seventhNum[2],callback:function(e){t.$set(t.tableDataDetail.Result.seventhNum,2,e)},expression:"tableDataDetail.Result.seventhNum[2]"}})],1),a("el-form-item",{attrs:{label:""}},[a("el-input",{model:{value:t.tableDataDetail.Result.seventhNum[3],callback:function(e){t.$set(t.tableDataDetail.Result.seventhNum,3,e)},expression:"tableDataDetail.Result.seventhNum[3]"}})],1),a("el-divider")],1),a("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[a("el-button",{on:{click:function(e){t.detailVisible=!1}}},[t._v("确 定")])],1)],1),a("el-dialog",{attrs:{title:"生成测试结果",visible:t.testAwardVisible,width:"150"},on:{"update:visible":function(e){t.testAwardVisible=e}}},[a("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[a("el-button",{on:{click:function(e){t.testAwardVisible=!1}}},[t._v("取 消")]),a("el-button",{on:{click:t.addTestAward}},[t._v("确 定")])],1)])],1)},i=[],s=a("95c3"),u=a.n(s),n=a("751a"),o={name:"Wsx3AwardList",data:function(){return{defaultSrc:a("7159"),query:{beginDay:"",endDay:"",pageIndex:1,pageSize:10,GameType:203},tableData:[],tableDataDetail:{GameName:"",LotteryStr:"",Result:{jackpots:"",firstNum:"",secondNum:["",""],thirdNum:["","","","","",""],forthNum:["","","","",""],fifthNum:["","","","","",""],sixthNum:["","",""],seventhNum:["","","",""]}},detailVisible:!1,pageTotal:0,rowId:0,testAwardVisible:!1}},created:function(){this.getData()},components:{VueCropper:u.a},methods:{getSummaries:function(){var t=[];return t},openDetail:function(t,e){this.rowId=e.Id,this.getDetailData()},getData:function(){var t=this;Object(n["a"])({url:"api/getawardlist",method:"post",data:this.query}).then((function(e){200==e.code&&(t.tableData=e.obj.ListData,t.pageTotal=e.obj.PageTotal)}))},getDetailData:function(){var t=this,e={id:this.rowId};Object(n["a"])({url:"api/getawarddetail",method:"post",data:e}).then((function(e){200==e.code&&(console.log(e.obj),t.detailVisible=!0,t.tableDataDetail=e.obj)}))},handleSearch:function(){this.$set(this.query,"pageIndex",1),this.getData()},handleEdit:function(t,e){this.idx=t,this.form=e,this.editVisible=!0},handlePageChange:function(t){this.$set(this.query,"pageIndex",t),this.getData()},handleAddTestAward:function(){},addTestAward:function(){var t=this,e={GameType:this.query.GameType};Object(n["a"])({url:"api/addtestaward",method:"post",data:e}).then((function(e){200==e.code?(t.testAwardVisible=!1,t.getData(),t.$message.success("添加成功")):t.$message.error("添加失败："+e.msg)}))}}},r=o,m=(a("7bb8"),a("2877")),b=Object(m["a"])(r,l,i,!1,null,"50b574ce",null);e["default"]=b.exports}}]);
//# sourceMappingURL=chunk-1a2beaad.7323a546.js.map