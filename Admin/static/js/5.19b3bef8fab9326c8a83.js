webpackJsonp([5],{BRwL:function(e,t){},r2uM:function(e,t,s){"use strict";Object.defineProperty(t,"__esModule",{value:!0});var n=s("986c"),a=s("RCTC"),o={name:"Wsx3ZbcOddsInfoSet",components:{LoHeader:n.a},data:function(){return{num:10,refreshing:!1,loading:!1,FilterItems:[{label:"赔率名:",Value:"",type:"text",placeholder:"赔率名",icon:"",Name:"OddsDes"},{label:"赔率Key:",Value:"",type:"text",placeholder:"赔率Key",icon:"",Name:"OddsType"},{label:"赔率分类:",Value:0,type:"select",placeholder:"赔率分类",icon:"",Name:"BigType",options:this.Config.Wsx1NBigTypeOptions}],PageTotal:0,tableData:[],GroupData:{},query:{OddsDes:"",pageIndex:1,OddsType:0,BigType:0,pageSize:10,GameType:this.Config.GameType.Wsx_203Zbc,FirstId:0},openFullscreen:!1,formInfo:{OldData:{},OneOddsInfo:{Id:0,OddsDes:"",Odds:1,OneUserMaxBet:99999,OneUserMinBet:0,AllUserMaxBet:999999}}}},methods:{refresh:function(){var e=this;this.refreshing=!0,this.query.pageIndex=1,this.query.FirstId=0,this.getData(function(t){1==t&&(e.query.pageIndex=e.query.pageIndex+1),e.refreshing=!1})},load:function(){var e=this;this.loading=!0,this.getData(function(t){1==t&&(e.query.pageIndex=e.query.pageIndex+1),e.loading=!1})},getData:function(e){var t=this;Object(a.a)(this.query).then(function(s){var n=s.obj,a=s.code,o=s.msg;200==a?(0==t.query.FirstId?t.tableData=n.ListData:t.tableData=t.tableData.concat(n.ListData),t.GroupData=n.GroupData,t.PageTotal=n.PageTotal,t.query.FirstId=n.GroupData.FirstId,e(!0),resolve()):(t.$toast.error(o),e(!1))}).catch(function(t){e(!1)})},openDialog:function(e){this.OldData=e,this.formInfo.OneOddsInfo.Id=e.Id,this.formInfo.OneOddsInfo.OddsDes=e.OddsDes,this.formInfo.OneOddsInfo.Odds=e.Odds,this.formInfo.OneOddsInfo.OneUserMaxBet=e.OneUserMaxBet,this.formInfo.OneOddsInfo.OneUserMinBet=e.OneUserMinBet,this.formInfo.OneOddsInfo.AllUserMaxBet=e.AllUserMaxBet,this.openFullscreen=!0},closeDialog:function(){this.openFullscreen=!1},saveModify:function(){var e=this;Object(a.b)(this.formInfo.OneOddsInfo).then(function(t){t.obj;var s=t.code,n=t.msg;200==s?(e.OldData.OddsDes=e.formInfo.OneOddsInfo.OddsDes,e.OldData.Odds=e.formInfo.OneOddsInfo.Odds,e.OldData.OneUserMaxBet=e.formInfo.OneOddsInfo.OneUserMaxBet,e.OldData.OneUserMinBet=e.formInfo.OneOddsInfo.OneUserMinBet,e.OldData.AllUserMaxBet=e.formInfo.OneOddsInfo.AllUserMaxBet,e.closeDialog(),resolve()):e.$toast.error(n)}).catch(function(e){})},queryData:function(e){var t=this;for(var s in this.query.pageIndex=1,this.query.FirstId=0,this.query.userName="",this.query.userId=0,this.query.beginDay="",this.query.endDay="",e)"OddsDes"===e[s].Name?this.query.OddsDes=e[s].Value:"OddsType"===e[s].Name?this.query.OddsType=e[s].Value:"BigType"===e[s].Name&&(this.query.BigType=e[s].Value);this.getData(function(e){1==e&&(t.query.pageIndex=t.query.pageIndex+1)})}},created:function(){var e=this;this.query.pageIndex=1,this.query.FirstId=0,this.getData(function(t){1==t&&(e.query.pageIndex=e.query.pageIndex+1)})}},r={render:function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("div",[s("LoHeader",{staticClass:"lo_header",attrs:{FilterItems:e.FilterItems,title:"中部彩赔率设置"},on:{query:e.queryData}}),e._v(" "),s("mu-container",{ref:"container",staticClass:"lo_content_box"},[s("mu-load-more",{staticClass:"lo_content",staticStyle:{width:"100%"},attrs:{refreshing:e.refreshing,loading:e.loading},on:{refresh:e.refresh,load:e.load}},[e._l(e.tableData,function(t){return[s("div",[s("div",{staticClass:"container"},[s("div",{staticClass:"desc"},[s("div",{staticClass:"title"},[s("div",{staticClass:"title_1"},[e._v("类型："+e._s(t.OddsDes))]),e._v(" "),s("div",{staticClass:"title_1"},[e._v("赔率："+e._s(t.Odds))]),e._v(" "),s("div",{staticClass:"btn"},[s("div",{staticClass:"btn_01",on:{click:function(s){return e.openDialog(t)}}},[e._v("修改")])])]),e._v(" "),s("div",{staticClass:"priceDiv"},[s("div",[s("span",{staticClass:"label"},[e._v("全上限：")]),e._v(" "),s("span",{staticClass:"price"},[e._v(e._s(t.AllUserMaxBet))])]),e._v(" "),s("div",[s("span",{staticClass:"label"},[e._v("单上限：")]),e._v(" "),s("span",{staticClass:"price"},[e._v(e._s(t.OneUserMaxBet))])]),e._v(" "),s("div",[s("span",{staticClass:"label"},[e._v("单下限：")]),e._v(" "),s("span",{staticClass:"price"},[e._v(e._s(t.OneUserMinBet))])])])])])])]})],2)],1),e._v(" "),s("mu-dialog",{attrs:{width:"360",transition:"slide-bottom",fullscreen:"",open:e.openFullscreen},on:{"update:open":function(t){e.openFullscreen=t}}},[s("mu-appbar",{attrs:{color:"primary",title:"修改赔率"}},[s("mu-button",{attrs:{slot:"left",icon:""},on:{click:e.closeDialog},slot:"left"},[s("mu-icon",{attrs:{value:"close"}})],1),e._v(" "),s("mu-button",{attrs:{slot:"right",flat:""},on:{click:e.saveModify},slot:"right"},[e._v("\n        保存\n      ")])],1),e._v(" "),s("div",{staticStyle:{padding:"24px"}},[s("mu-form",{ref:"form",staticClass:"mu-demo-form",attrs:{model:e.formInfo.OneOddsInfo}},[s("mu-form-item",{attrs:{label:"赔率名称",prop:"OddsDes"}},[s("mu-text-field",{attrs:{disabled:"",prop:"OddsDes"},model:{value:e.formInfo.OneOddsInfo.OddsDes,callback:function(t){e.$set(e.formInfo.OneOddsInfo,"OddsDes",t)},expression:"formInfo.OneOddsInfo.OddsDes"}})],1),e._v(" "),s("mu-form-item",{attrs:{label:"赔率",prop:"Odds"}},[s("mu-text-field",{attrs:{type:"number",prop:"Odds"},model:{value:e.formInfo.OneOddsInfo.Odds,callback:function(t){e.$set(e.formInfo.OneOddsInfo,"Odds",t)},expression:"formInfo.OneOddsInfo.Odds"}})],1),e._v(" "),s("mu-form-item",{attrs:{label:"单上限",prop:"OneUserMaxBet"}},[s("mu-text-field",{attrs:{type:"number",prop:"OneUserMaxBet"},model:{value:e.formInfo.OneOddsInfo.OneUserMaxBet,callback:function(t){e.$set(e.formInfo.OneOddsInfo,"OneUserMaxBet",t)},expression:"formInfo.OneOddsInfo.OneUserMaxBet"}})],1),e._v(" "),s("mu-form-item",{attrs:{label:"单下限",prop:"OneUserMinBet"}},[s("mu-text-field",{attrs:{type:"number",prop:"OneUserMinBet"},model:{value:e.formInfo.OneOddsInfo.OneUserMinBet,callback:function(t){e.$set(e.formInfo.OneOddsInfo,"OneUserMinBet",t)},expression:"formInfo.OneOddsInfo.OneUserMinBet"}})],1),e._v(" "),s("mu-form-item",{attrs:{label:"全下限",prop:"AllUserMaxBet"}},[s("mu-text-field",{attrs:{type:"number",prop:"AllUserMaxBet"},model:{value:e.formInfo.OneOddsInfo.AllUserMaxBet,callback:function(t){e.$set(e.formInfo.OneOddsInfo,"AllUserMaxBet",t)},expression:"formInfo.OneOddsInfo.AllUserMaxBet"}})],1)],1)],1)],1)],1)},staticRenderFns:[]};var d=s("VU/8")(o,r,!1,function(e){s("BRwL")},"data-v-ece96f9e",null);t.default=d.exports}});
//# sourceMappingURL=5.19b3bef8fab9326c8a83.js.map