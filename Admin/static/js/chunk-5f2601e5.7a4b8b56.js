(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-5f2601e5"],{"55f2":function(a,t,e){"use strict";var l=e("867a"),s=e.n(l);s.a},"867a":function(a,t,e){},f76e:function(a,t,e){"use strict";e.r(t);var l=function(){var a=this,t=a.$createElement,e=a._self._c||t;return e("div",[e("div",{staticClass:"container"},[e("el-tabs",{model:{value:a.activeName,callback:function(t){a.activeName=t},expression:"activeName"}},[e("el-tab-pane",{attrs:{label:"大小单双",name:"tab0"}},[e("el-card",{staticClass:"box-card"},[e("div",[e("el-breadcrumb-item",[a._v("\n                  赔率Key范围（\n                  "),e("el-input-number",{staticClass:"loose-input",attrs:{step:1,min:a.allData[0].BatchForm.Min,max:a.allData[0].BatchForm.Max},model:{value:a.allData[0].BatchForm.MinValue,callback:function(t){a.$set(a.allData[0].BatchForm,"MinValue",t)},expression:"allData[0].BatchForm.MinValue"}}),a._v("\n                  ~\n                  "),e("el-input-number",{staticClass:"loose-input",attrs:{step:1,min:a.allData[0].BatchForm.Min,max:a.allData[0].BatchForm.Max},model:{value:a.allData[0].BatchForm.MaxValue,callback:function(t){a.$set(a.allData[0].BatchForm,"MaxValue",t)},expression:"allData[0].BatchForm.MaxValue"}}),a._v("）\n                  赔率：\n                  "),e("el-input-number",{staticClass:"loose-input",attrs:{precision:2,step:1,min:0},model:{value:a.allData[0].BatchForm.Odds,callback:function(t){a.$set(a.allData[0].BatchForm,"Odds",t)},expression:"allData[0].BatchForm.Odds"}}),e("el-button",{staticClass:"handle-del mr10",attrs:{type:"primary"},on:{click:function(t){return a.batchModifyOdds(0)}}},[a._v("批量修改")])],1)],1)]),e("el-card",{staticClass:"box-card"},[e("el-button",{staticClass:"handle-del mr10",attrs:{type:"primary"},on:{click:function(t){return a.getData(0)}}},[a._v("刷新")]),e("el-table",{ref:"multipleTable",staticClass:"table",attrs:{data:a.allData[0].tableData,border:"","header-cell-class-name":"table-header"}},[e("el-table-column",{attrs:{prop:"OddsType",label:"赔率Key",width:"90",align:"center"}}),e("el-table-column",{attrs:{prop:"OddsDes",label:"赔率名称"}}),e("el-table-column",{attrs:{prop:"Odds",label:"赔率"},scopedSlots:a._u([{key:"default",fn:function(t){return[a._v(a._s(a._f("money")(t.row.Odds)))]}}])}),e("el-table-column",{attrs:{label:"操作",width:"180",align:"center"},scopedSlots:a._u([{key:"default",fn:function(t){return[e("el-button",{attrs:{type:"text",icon:"el-icon-edit"},on:{click:function(e){return a.handleEdit(0,t.$index,t.row)}}},[a._v("编辑\n                    ")])]}}])})],1)],1)],1),e("el-tab-pane",{attrs:{label:"对顺豹子",name:"tab1"}},[e("el-card",{staticClass:"box-card"},[e("div",[e("el-breadcrumb-item",[a._v("\n                  赔率Key范围（\n                  "),e("el-input-number",{staticClass:"loose-input",attrs:{step:1,min:a.allData[1].BatchForm.Min,max:a.allData[1].BatchForm.Max},model:{value:a.allData[1].BatchForm.MinValue,callback:function(t){a.$set(a.allData[1].BatchForm,"MinValue",t)},expression:"allData[1].BatchForm.MinValue"}}),a._v("\n                  ~\n                  "),e("el-input-number",{staticClass:"loose-input",attrs:{step:1,min:a.allData[1].BatchForm.Min,max:a.allData[1].BatchForm.Max},model:{value:a.allData[1].BatchForm.MaxValue,callback:function(t){a.$set(a.allData[1].BatchForm,"MaxValue",t)},expression:"allData[1].BatchForm.MaxValue"}}),a._v("）\n                  赔率：\n                  "),e("el-input-number",{staticClass:"loose-input",attrs:{precision:2,step:1,min:0},model:{value:a.allData[1].BatchForm.Odds,callback:function(t){a.$set(a.allData[1].BatchForm,"Odds",t)},expression:"allData[1].BatchForm.Odds"}}),e("el-button",{staticClass:"handle-del mr10",attrs:{type:"primary"},on:{click:function(t){return a.batchModifyOdds(1)}}},[a._v("批量修改")])],1)],1)]),e("el-card",{staticClass:"box-card"},[e("el-button",{staticClass:"handle-del mr10",attrs:{type:"primary"},on:{click:function(t){return a.getData(1)}}},[a._v("刷新")]),e("el-table",{ref:"multipleTable",staticClass:"table",attrs:{data:a.allData[1].tableData,border:"","header-cell-class-name":"table-header"}},[e("el-table-column",{attrs:{prop:"OddsType",label:"赔率Key",width:"90",align:"center"}}),e("el-table-column",{attrs:{prop:"OddsDes",label:"赔率名称"}}),e("el-table-column",{attrs:{prop:"Odds",label:"赔率"},scopedSlots:a._u([{key:"default",fn:function(t){return[a._v(a._s(a._f("money")(t.row.Odds)))]}}])}),e("el-table-column",{attrs:{label:"操作",width:"180",align:"center"},scopedSlots:a._u([{key:"default",fn:function(t){return[e("el-button",{attrs:{type:"text",icon:"el-icon-edit"},on:{click:function(e){return a.handleEdit(1,t.$index,t.row)}}},[a._v("编辑\n                    ")])]}}])})],1)],1)],1),e("el-tab-pane",{attrs:{label:"波色",name:"tab2"}},[e("el-card",{staticClass:"box-card"},[e("div",[e("el-breadcrumb-item",[a._v("\n                  赔率Key范围（\n                  "),e("el-input-number",{staticClass:"loose-input",attrs:{step:1,min:a.allData[2].BatchForm.Min,max:a.allData[2].BatchForm.Max},model:{value:a.allData[2].BatchForm.MinValue,callback:function(t){a.$set(a.allData[2].BatchForm,"MinValue",t)},expression:"allData[2].BatchForm.MinValue"}}),a._v("\n                  ~\n                  "),e("el-input-number",{staticClass:"loose-input",attrs:{step:1,min:a.allData[2].BatchForm.Min,max:a.allData[2].BatchForm.Max},model:{value:a.allData[2].BatchForm.MaxValue,callback:function(t){a.$set(a.allData[2].BatchForm,"MaxValue",t)},expression:"allData[2].BatchForm.MaxValue"}}),a._v("）\n                  赔率：\n                  "),e("el-input-number",{staticClass:"loose-input",attrs:{precision:2,step:1,min:0},model:{value:a.allData[2].BatchForm.Odds,callback:function(t){a.$set(a.allData[2].BatchForm,"Odds",t)},expression:"allData[2].BatchForm.Odds"}}),e("el-button",{staticClass:"handle-del mr10",attrs:{type:"primary"},on:{click:function(t){return a.batchModifyOdds(2)}}},[a._v("批量修改")])],1)],1)]),e("el-card",{staticClass:"box-card"},[e("el-button",{staticClass:"handle-del mr10",attrs:{type:"primary"},on:{click:function(t){return a.getData(2)}}},[a._v("刷新")]),e("el-table",{ref:"multipleTable",staticClass:"table",attrs:{data:a.allData[2].tableData,border:"","header-cell-class-name":"table-header"}},[e("el-table-column",{attrs:{prop:"OddsType",label:"赔率Key",width:"90",align:"center"}}),e("el-table-column",{attrs:{prop:"OddsDes",label:"赔率名称"}}),e("el-table-column",{attrs:{prop:"Odds",label:"赔率"},scopedSlots:a._u([{key:"default",fn:function(t){return[a._v(a._s(a._f("money")(t.row.Odds)))]}}])}),e("el-table-column",{attrs:{label:"操作",width:"180",align:"center"},scopedSlots:a._u([{key:"default",fn:function(t){return[e("el-button",{attrs:{type:"text",icon:"el-icon-edit"},on:{click:function(e){return a.handleEdit(2,t.$index,t.row)}}},[a._v("编辑\n                    ")])]}}])})],1)],1)],1),e("el-tab-pane",{attrs:{label:"龙虎和",name:"tab3"}},[e("el-card",{staticClass:"box-card"},[e("div",[e("el-breadcrumb-item",[a._v("\n                  赔率Key范围（\n                  "),e("el-input-number",{staticClass:"loose-input",attrs:{step:1,min:a.allData[3].BatchForm.Min,max:a.allData[3].BatchForm.Max},model:{value:a.allData[3].BatchForm.MinValue,callback:function(t){a.$set(a.allData[3].BatchForm,"MinValue",t)},expression:"allData[3].BatchForm.MinValue"}}),a._v("\n                  ~\n                  "),e("el-input-number",{staticClass:"loose-input",attrs:{step:1,min:a.allData[3].BatchForm.Min,max:a.allData[3].BatchForm.Max},model:{value:a.allData[3].BatchForm.MaxValue,callback:function(t){a.$set(a.allData[3].BatchForm,"MaxValue",t)},expression:"allData[3].BatchForm.MaxValue"}}),a._v("）\n                  赔率：\n                  "),e("el-input-number",{staticClass:"loose-input",attrs:{precision:2,step:1,min:0},model:{value:a.allData[3].BatchForm.Odds,callback:function(t){a.$set(a.allData[3].BatchForm,"Odds",t)},expression:"allData[3].BatchForm.Odds"}}),e("el-button",{staticClass:"handle-del mr10",attrs:{type:"primary"},on:{click:function(t){return a.batchModifyOdds(3)}}},[a._v("批量修改")])],1)],1)]),e("el-card",{staticClass:"box-card"},[e("el-button",{staticClass:"handle-del mr10",attrs:{type:"primary"},on:{click:function(t){return a.getData(3)}}},[a._v("刷新")]),e("el-table",{ref:"multipleTable",staticClass:"table",attrs:{data:a.allData[3].tableData,border:"","header-cell-class-name":"table-header"}},[e("el-table-column",{attrs:{prop:"OddsType",label:"赔率Key",width:"90",align:"center"}}),e("el-table-column",{attrs:{prop:"OddsDes",label:"赔率名称"}}),e("el-table-column",{attrs:{prop:"Odds",label:"赔率"},scopedSlots:a._u([{key:"default",fn:function(t){return[a._v(a._s(a._f("money")(t.row.Odds)))]}}])}),e("el-table-column",{attrs:{label:"操作",width:"180",align:"center"},scopedSlots:a._u([{key:"default",fn:function(t){return[e("el-button",{attrs:{type:"text",icon:"el-icon-edit"},on:{click:function(e){return a.handleEdit(3,t.$index,t.row)}}},[a._v("编辑\n                    ")])]}}])})],1)],1)],1),e("el-tab-pane",{attrs:{label:"庄闲",name:"tab4"}},[e("el-card",{staticClass:"box-card"},[e("div",[e("el-breadcrumb-item",[a._v("\n                  赔率Key范围（\n                  "),e("el-input-number",{staticClass:"loose-input",attrs:{step:1,min:a.allData[4].BatchForm.Min,max:a.allData[4].BatchForm.Max},model:{value:a.allData[4].BatchForm.MinValue,callback:function(t){a.$set(a.allData[4].BatchForm,"MinValue",t)},expression:"allData[4].BatchForm.MinValue"}}),a._v("\n                  ~\n                  "),e("el-input-number",{staticClass:"loose-input",attrs:{step:1,min:a.allData[4].BatchForm.Min,max:a.allData[4].BatchForm.Max},model:{value:a.allData[4].BatchForm.MaxValue,callback:function(t){a.$set(a.allData[4].BatchForm,"MaxValue",t)},expression:"allData[4].BatchForm.MaxValue"}}),a._v("）\n                  赔率：\n                  "),e("el-input-number",{staticClass:"loose-input",attrs:{precision:2,step:1,min:0},model:{value:a.allData[4].BatchForm.Odds,callback:function(t){a.$set(a.allData[4].BatchForm,"Odds",t)},expression:"allData[4].BatchForm.Odds"}}),e("el-button",{staticClass:"handle-del mr10",attrs:{type:"primary"},on:{click:function(t){return a.batchModifyOdds(4)}}},[a._v("批量修改")])],1)],1)]),e("el-card",{staticClass:"box-card"},[e("el-button",{staticClass:"handle-del mr10",attrs:{type:"primary"},on:{click:function(t){return a.getData(4)}}},[a._v("刷新")]),e("el-table",{ref:"multipleTable",staticClass:"table",attrs:{data:a.allData[4].tableData,border:"","header-cell-class-name":"table-header"}},[e("el-table-column",{attrs:{prop:"OddsType",label:"赔率Key",width:"90",align:"center"}}),e("el-table-column",{attrs:{prop:"OddsDes",label:"赔率名称"}}),e("el-table-column",{attrs:{prop:"Odds",label:"赔率"},scopedSlots:a._u([{key:"default",fn:function(t){return[a._v(a._s(a._f("money")(t.row.Odds)))]}}])}),e("el-table-column",{attrs:{label:"操作",width:"180",align:"center"},scopedSlots:a._u([{key:"default",fn:function(t){return[e("el-button",{attrs:{type:"text",icon:"el-icon-edit"},on:{click:function(e){return a.handleEdit(4,t.$index,t.row)}}},[a._v("编辑\n                    ")])]}}])})],1)],1)],1),e("el-tab-pane",{attrs:{label:"数值",name:"tab5"}},[e("el-card",{staticClass:"box-card"},[e("div",[e("el-breadcrumb-item",[a._v("\n                  赔率Key范围（\n                  "),e("el-input-number",{staticClass:"loose-input",attrs:{step:1,min:a.allData[5].BatchForm.Min,max:a.allData[5].BatchForm.Max},model:{value:a.allData[5].BatchForm.MinValue,callback:function(t){a.$set(a.allData[5].BatchForm,"MinValue",t)},expression:"allData[5].BatchForm.MinValue"}}),a._v("\n                  ~\n                  "),e("el-input-number",{staticClass:"loose-input",attrs:{step:1,min:a.allData[5].BatchForm.Min,max:a.allData[5].BatchForm.Max},model:{value:a.allData[5].BatchForm.MaxValue,callback:function(t){a.$set(a.allData[5].BatchForm,"MaxValue",t)},expression:"allData[5].BatchForm.MaxValue"}}),a._v("）\n                  赔率：\n                  "),e("el-input-number",{staticClass:"loose-input",attrs:{precision:2,step:1,min:0},model:{value:a.allData[5].BatchForm.Odds,callback:function(t){a.$set(a.allData[5].BatchForm,"Odds",t)},expression:"allData[5].BatchForm.Odds"}}),e("el-button",{staticClass:"handle-del mr10",attrs:{type:"primary"},on:{click:function(t){return a.batchModifyOdds(5)}}},[a._v("批量修改")])],1)],1)]),e("el-card",{staticClass:"box-card"},[e("el-button",{staticClass:"handle-del mr10",attrs:{type:"primary"},on:{click:function(t){return a.getData(5)}}},[a._v("刷新")]),e("el-table",{ref:"multipleTable",staticClass:"table",attrs:{data:a.allData[5].tableData,border:"","header-cell-class-name":"table-header"}},[e("el-table-column",{attrs:{prop:"OddsType",label:"赔率Key",width:"90",align:"center"}}),e("el-table-column",{attrs:{prop:"OddsDes",label:"赔率名称"}}),e("el-table-column",{attrs:{prop:"Odds",label:"赔率"},scopedSlots:a._u([{key:"default",fn:function(t){return[a._v(a._s(a._f("money")(t.row.Odds)))]}}])}),e("el-table-column",{attrs:{label:"操作",width:"180",align:"center"},scopedSlots:a._u([{key:"default",fn:function(t){return[e("el-button",{attrs:{type:"text",icon:"el-icon-edit"},on:{click:function(e){return a.handleEdit(5,t.$index,t.row)}}},[a._v("编辑\n                    ")])]}}])})],1)],1)],1)],1),e("el-dialog",{attrs:{title:"修改赔率",visible:a.editVisible,width:"90%"},on:{"update:visible":function(t){a.editVisible=t}}},[e("el-form",{ref:"form",attrs:{model:a.form,"label-width":"100px"}},[e("el-form-item",{attrs:{label:"赔率Key",prop:"OddsType"}},[e("el-input",{attrs:{disabled:""},model:{value:a.form.OddsType,callback:function(t){a.$set(a.form,"OddsType",t)},expression:"form.OddsType"}})],1),e("el-form-item",{attrs:{label:"赔率名称",prop:"OddsDes"}},[e("el-input",{attrs:{disabled:""},model:{value:a.form.OddsDes,callback:function(t){a.$set(a.form,"OddsDes",t)},expression:"form.OddsDes"}})],1),e("el-form-item",{attrs:{label:"赔率",prop:"Odds"}},[e("el-input-number",{staticClass:"loose-input",attrs:{precision:2,step:1,min:0},model:{value:a.form.Odds,callback:function(t){a.$set(a.form,"Odds",t)},expression:"form.Odds"}})],1),e("el-form-item",{attrs:{label:"一个用户最多",prop:"OneUserMaxBet"}},[e("el-input-number",{staticClass:"loose-input",attrs:{step:1e3,min:0},model:{value:a.form.OneUserMaxBet,callback:function(t){a.$set(a.form,"OneUserMaxBet",t)},expression:"form.OneUserMaxBet"}})],1),e("el-form-item",{attrs:{label:"一个用户最少",prop:"OneUserMinBet"}},[e("el-input-number",{staticClass:"loose-input",attrs:{step:10,min:1},model:{value:a.form.OneUserMinBet,callback:function(t){a.$set(a.form,"OneUserMinBet",t)},expression:"form.OneUserMinBet"}})],1),e("el-form-item",{attrs:{label:"全部用户最多",prop:"AllUserMaxBet"}},[e("el-input-number",{staticClass:"loose-input",attrs:{step:1e3,min:0},model:{value:a.form.AllUserMaxBet,callback:function(t){a.$set(a.form,"AllUserMaxBet",t)},expression:"form.AllUserMaxBet"}})],1)],1),e("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[e("el-button",{on:{click:function(t){a.editVisible=!1}}},[a._v("取 消")]),e("el-button",{attrs:{type:"primary"},on:{click:a.saveOdds}},[a._v("确 定")])],1)],1)],1)])},s=[],n=e("751a"),r={name:"Zg28BjSet",data:function(){return{AwardTypeItem:[{name:"金币",value:1},{name:"积分",value:2}],query:{address:"",name:"",pageIndex:1,pageSize:10},allData:[{I:0,BigTypes:[1],tableData:[],BatchForm:{MinValue:11,MaxValue:20,Odds:1.9,Min:11,Max:20}},{I:1,BigTypes:[2],tableData:[],BatchForm:{MinValue:21,MaxValue:23,Odds:9.9,Min:21,Max:23}},{I:2,BigTypes:[3],tableData:[],BatchForm:{MinValue:24,MaxValue:26,Odds:9.9,Min:24,Max:26}},{I:3,BigTypes:[4],tableData:[],BatchForm:{MinValue:27,MaxValue:29,Odds:9.9,Min:27,Max:29}},{I:4,BigTypes:[5],tableData:[],BatchForm:{MinValue:31,MaxValue:34,Odds:9.9,Min:31,Max:34}},{I:5,BigTypes:[6],tableData:[],BatchForm:{MinValue:100,MaxValue:127,Odds:9.9,Min:100,Max:127}}],activeName:"tab0",multipleSelection:[],delList:[],editVisible:!1,pageTotal:0,form:{Id:0,OddsType:0,Odds:1.9,OddsDes:""},curFormT:0,curFormIdx:-1,GameType:42}},created:function(){this.getData(0),this.getData(1),this.getData(2),this.getData(3),this.getData(4),this.getData(5)},methods:{batchModifyOdds:function(a){var t=this,e={GameType:this.GameType,Odds:this.allData[a].BatchForm.Odds,MinOddType:this.allData[a].BatchForm.MinValue,MaxOddType:this.allData[a].BatchForm.MaxValue};Object(n["a"])({url:"api/savebatchodds",method:"post",data:e}).then((function(e){console.log(e),200==e.code?(t.editVisible=!1,t.$message.success("修改 ".concat(t.form.OddsDes," 成功")),t.getData(a,t.allData[a].BigTypes)):t.$message.error("更新失败："+e.msg)}))},getData:function(a){var t=this,e=this.allData[a].BigTypes,l={GameType:this.GameType,ArrBigType:e},s={jsonData:JSON.stringify(l)};console.log("OK"),Object(n["a"])({url:"api/getoddslist",method:"post",data:s}).then((function(e){console.log(e),200==e.code&&(t.allData[a].tableData=e.obj)}))},saveOdds:function(){var a=this;Object(n["a"])({url:"api/saveoddsinfo",method:"post",data:this.form}).then((function(t){console.log(t),200==t.code?(a.editVisible=!1,a.$message.success("修改 ".concat(a.form.OddsDes," 成功")),a.getData(a.curFormT,a.allData[a.curFormT].BigTypes)):a.$message.error("更新失败："+t.msg)}))},handleEdit:function(a,t,e){this.curFormIdx=t,this.curFormT=a,this.form.Id=e.Id,this.form.OddsType=e.OddsType,this.form.OddsDes=e.OddsDes,this.form.Odds=e.Odds,this.form.OneUserMaxBet=e.OneUserMaxBet,this.form.OneUserMinBet=e.OneUserMinBet,this.form.AllUserMaxBet=e.AllUserMaxBet,this.editVisible=!0}}},o=r,i=(e("55f2"),e("2877")),c=Object(i["a"])(o,l,s,!1,null,"71f59e20",null);t["default"]=c.exports}}]);
//# sourceMappingURL=chunk-5f2601e5.7a4b8b56.js.map