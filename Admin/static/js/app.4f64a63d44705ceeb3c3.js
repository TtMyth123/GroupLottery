webpackJsonp([26],{"1uuo":function(n,e){},"3pLw":function(n,e,t){"use strict";e.a=function(){var n=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{},e=i.a.create({baseURL:Object({NODE_ENV:"production"}).VUE_APP_API});e.interceptors.request.use(function(n){if("post"===n.method)n.data=c.a.stringify(n.data);else if("file"===n.method&&(n.method="post",n.data&&n.data instanceof Object)){var e=new FormData;for(var t in n.data)e.append(t,n.data[t]);n.data=e}return n},function(n){return a.a.reject(n)}),e.interceptors.response.use(function(n){return console.log(n.data),401==n.data.code&&v.a.push("/login"),n},function(n){return a.a.reject(n)});var t=Object(s.a)();return n.headers={token:t},e.request(n).then(function(n){return n.data})};var o=t("//Fk"),a=t.n(o),r=t("mtWM"),i=t.n(r),s=t("TIfe"),u=t("mw3O"),c=t.n(u),v=t("YaEn")},"4q4r":function(n,e){},Chq7:function(n,e){},E51W:function(n,e){},NHnr:function(n,e,t){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var o=t("7+uW"),a={render:function(){var n=this.$createElement,e=this._self._c||n;return e("div",{attrs:{id:"app"}},[e("router-view")],1)},staticRenderFns:[]};var r=t("VU/8")({name:"App"},a,!1,function(n){t("4q4r")},null,null).exports,i=t("YaEn"),s=t("aFc6"),u=t("NYxO"),c=(t("sax8"),{namespaced:!0,state:{num:0},getters:{reverseOperation:function(n,e,t){return{num:0!=n.num?-1*n.num:n.num}}},actions:{increase:function(n,e){var t=n.commit;n.state;console.log(e),setTimeout(function(){t("INCREASE",3*e)},300)}},mutations:{INCREASE:function(n,e){console.log(e),n.num+=e}}}),v=t("//Fk"),l=t.n(v),m=t("bOdI"),p=t.n(m),h=t("3pLw");var f,d=t("TIfe"),b={namespaced:!0,state:{isLogin:!1,sysUserInfo:{},menuItem:[]},mutations:(f={},p()(f,"SET_SYSUSERINFO",function(n,e){console.log(e),n.sysUserInfo=e}),p()(f,"SET_ISLOGIN",function(n,e){n.isLogin=e}),f),actions:{login:function(n,e){var t=e.username,o=e.password;return new l.a(function(e,a){var r;(r={username:t.trim(),password:o},Object(h.a)({url:"loginreg/dologin",method:"post",data:r})).then(function(t){var o=t.obj,r=t.code,i=t.msg;200==r?(n.commit("SET_SYSUSERINFO",o),n.commit("SET_ISLOGIN",!0),localStorage.setItem("token",o.CurToken),Object(d.c)(o.CurToken),e()):(n.commit("set_IsLogin",!1),a(i))}).catch(function(e){n.commit("set_IsLogin",!1),a(e)})})},logout:function(n,e,t){return new l.a(function(t,o){var a;(a=e.token,Object(h.a)({url:"loginreg/dologout",method:"post",data:a})).then(function(){n.commit("SET_ISLOGIN",!1),Object(d.b)(),t()}).catch(function(n){o(n)})})}}};o.a.use(u.a);var g=new u.a.Store({modules:{a:c,sysUser:b},strict:!1,plugins:[]});o.a.filter("BetStateName",function(n){return 1===n?"待开奖":2===n?"已兑奖":""});t("E51W"),t("uTBe");var _=t("sXio"),B=(t("Chq7"),t("w7Ps")),L={GameType:{Wsx_201Nbc:201,Wsx_202Bbc:202,Wsx_203Zbc:203,ZG28_41Jnd:41,ZG28_42Bj:42,ZG28_43Xjp:43},AccountTypesOptions:[{n:"全部",v:0},{n:"竞猜",v:1},{n:"赢得",v:2},{n:"充值",v:3},{n:"提现",v:4},{n:"提现拒绝",v:7},{n:"上分",v:8},{n:"下分",v:9},{n:"赠送",v:5},{n:"佣金转换",v:13}],UserTypeOptions:[{n:"全部",v:0},{n:"一般玩家",v:1},{n:"业务玩家",v:2},{n:"游客",v:3}],Wsx1NBigTypeOptions:[{n:"全部",v:0},{n:"头特大小",v:1},{n:"头特单双",v:2},{n:"一等奖大小",v:3},{n:"一等奖单双",v:4},{n:"头等-头特",v:5},{n:"B面-头等-头特",v:105},{n:"头等-尾特",v:6},{n:"B面-头等-尾特",v:106},{n:"头等特码",v:7},{n:"B面-头等特码",v:107},{n:"一等特码",v:8},{n:"B面-一等特码",v:108},{n:"二等特码",v:9},{n:"B面-二等特码",v:109},{n:"二连位",v:10},{n:"三连位",v:11},{n:"平码两位区",v:12},{n:"平码三位区",v:13},{n:"头等特码 波色",v:15},{n:"一等特码 波色",v:16},{n:"二等特码 波色",v:17},{n:"一等-头特",v:18},{n:"B面-一等-头特",v:118},{n:"一等-尾特",v:19},{n:"B面-一等-尾特",v:119},{n:"二等-头特",v:20},{n:"B面-二等-头特",v:120},{n:"二等-尾特",v:21},{n:"B面-二等-尾特",v:121}],Wsx2ZBigTypeOptions:[{n:"全部",v:0},{n:"头特大小",v:1},{n:"头特单双",v:2},{n:"一等奖大小",v:3},{n:"一等奖单双",v:4},{n:"头等-头特",v:5},{n:"B面-头等-头特",v:105},{n:"头等-尾特",v:6},{n:"B面-头等-尾特",v:106},{n:"头等特码",v:7},{n:"B面-头等特码",v:107},{n:"一等特码",v:8},{n:"B面-一等特码",v:108},{n:"二等特码",v:9},{n:"B面-二等特码",v:109},{n:"二连位",v:10},{n:"三连位",v:11},{n:"平码两位区",v:12},{n:"平码三位区",v:13},{n:"头等特码 波色",v:15},{n:"一等特码 波色",v:16},{n:"二等特码 波色",v:17},{n:"一等-头特",v:18},{n:"B面-一等-头特",v:118},{n:"一等-尾特",v:19},{n:"B面-一等-尾特",v:119},{n:"二等-头特",v:20},{n:"B面-二等-头特",v:120},{n:"二等-尾特",v:21},{n:"B面-二等-尾特",v:121}],Wsx3BBigTypeOptions:[{n:"全部",v:0},{n:"头特大小",v:1},{n:"头特单双",v:2},{n:"一等奖大小",v:3},{n:"一等奖单双",v:4},{n:"头等-头特",v:5},{n:"B面-头等-头特",v:105},{n:"头等-尾特",v:6},{n:"B面-头等-尾特",v:106},{n:"头等特码",v:7},{n:"B面-头等特码",v:107},{n:"一等特码",v:8},{n:"B面-一等特码",v:108},{n:"二等特码",v:9},{n:"B面-二等特码",v:109},{n:"二连位",v:10},{n:"头等 三连位",v:11},{n:"一等 三连位",v:22},{n:"平码两位区",v:12},{n:"平码三位区",v:13},{n:"头等特码 波色",v:15},{n:"一等特码 波色",v:16},{n:"二等特码 波色",v:17},{n:"一等-头特",v:18},{n:"B面-一等-头特",v:118},{n:"一等-尾特",v:19},{n:"B面-一等-尾特",v:119},{n:"二等-头特",v:20},{n:"B面-二等-头特",v:120},{n:"二等-尾特",v:21},{n:"B面-二等-尾特",v:121}],BetStateOptions:[{n:"全部",v:0},{n:"待开奖",v:1},{n:"已兑奖",v:2}],GameTypeOptions:[{n:"全部",v:0},{n:"南部彩",v:201},{n:"北部彩",v:202},{n:"中部彩",v:203},{n:"加拿大28",v:41},{n:"北京28",v:42},{n:"新加坡28",v:43}],RebateTypeOptions:[{n:"全部",v:0},{n:"竞猜",v:1},{n:"转换金额",v:2}]};o.a.config.productionTip=!1,o.a.use(s.a);o.a.use(_.a,{position:"top",time:2e3,closeIcon:"close",close:!0,successIcon:"check_circle",infoIcon:"info",warningIcon:"priority_high",errorIcon:"warning"}),o.a.use(B.a),o.a.prototype.Config=L,new o.a({el:"#app",router:i.a,store:g,components:{App:r},template:"<App/>"})},TIfe:function(n,e,t){"use strict";e.a=function(){return a.a.get(r)},e.c=function(n){return a.a.set(r,n)},e.b=function(){return a.a.remove(r)};var o=t("lbHh"),a=t.n(o),r="Admin-Token"},YaEn:function(n,e,t){"use strict";var o=t("7+uW"),a=t("/ocq"),r={render:function(){var n=this,e=n.$createElement,t=n._self._c||e;return t("div",{staticClass:"hello"},[t("h1",[n._v(n._s(n.msg))]),n._v(" "),t("h2",[n._v("Essential Links")]),n._v(" "),n._m(0),n._v(" "),t("h2",[n._v("Ecosystem")]),n._v(" "),n._m(1)])},staticRenderFns:[function(){var n=this,e=n.$createElement,t=n._self._c||e;return t("ul",[t("li",[t("a",{attrs:{href:"https://vuejs.org",target:"_blank"}},[n._v("\n        Core Docs\n      ")])]),n._v(" "),t("li",[t("a",{attrs:{href:"https://forum.vuejs.org",target:"_blank"}},[n._v("\n        Forum\n      ")])]),n._v(" "),t("li",[t("a",{attrs:{href:"https://chat.vuejs.org",target:"_blank"}},[n._v("\n        Community Chat\n      ")])]),n._v(" "),t("li",[t("a",{attrs:{href:"https://twitter.com/vuejs",target:"_blank"}},[n._v("\n        Twitter\n      ")])]),n._v(" "),t("br"),n._v(" "),t("li",[t("a",{attrs:{href:"http://vuejs-templates.github.io/webpack/",target:"_blank"}},[n._v("\n        Docs for This Template\n      ")])])])},function(){var n=this.$createElement,e=this._self._c||n;return e("ul",[e("li",[e("a",{attrs:{href:"http://router.vuejs.org/",target:"_blank"}},[this._v("\n        vue-router\n      ")])]),this._v(" "),e("li",[e("a",{attrs:{href:"http://vuex.vuejs.org/",target:"_blank"}},[this._v("\n        vuex\n      ")])]),this._v(" "),e("li",[e("a",{attrs:{href:"http://vue-loader.vuejs.org/",target:"_blank"}},[this._v("\n        vue-loader\n      ")])]),this._v(" "),e("li",[e("a",{attrs:{href:"https://github.com/vuejs/awesome-vue",target:"_blank"}},[this._v("\n        awesome-vue\n      ")])])])}]};t("VU/8")({name:"HelloWorld",data:function(){return{msg:"Welcome to Your Vue.js App"}}},r,!1,function(n){t("1uuo")},"data-v-d8ec41bc",null).exports;o.a.use(a.a);e.a=new a.a({routes:[{path:"/",name:"登录",component:function(){return t.e(1).then(t.bind(null,"nb0W"))}},{path:"/home",name:"home",component:function(){return Promise.all([t.e(0),t.e(2)]).then(t.bind(null,"iwkC"))}},{path:"/Wsx1NbcAwardList",name:"Wsx1NbcAwardList",component:function(){return Promise.all([t.e(0),t.e(19)]).then(t.bind(null,"aMbK"))}},{path:"/Wsx2BbcAwardList",name:"Wsx2BbcAwardList",component:function(){return Promise.all([t.e(0),t.e(13)]).then(t.bind(null,"zn+H"))}},{path:"/Wsx3ZbcAwardList",name:"Wsx3ZbcAwardList",component:function(){return Promise.all([t.e(0),t.e(22)]).then(t.bind(null,"ZJNW"))}},{path:"/Zg28BjAwardList",name:"Zg28BjAwardList",component:function(){return Promise.all([t.e(0),t.e(15)]).then(t.bind(null,"gUm1"))}},{path:"/Zg28JndAwardList",name:"Zg28JndAwardList",component:function(){return Promise.all([t.e(0),t.e(11)]).then(t.bind(null,"LAY+"))}},{path:"/Zg28XjpAwardList",name:"Zg28XjpAwardList",component:function(){return Promise.all([t.e(0),t.e(10)]).then(t.bind(null,"sABS"))}},{path:"/AccountList",name:"AccountList",component:function(){return Promise.all([t.e(0),t.e(23)]).then(t.bind(null,"4jZU"))}},{path:"/RebateList",name:"RebateList",component:function(){return Promise.all([t.e(0),t.e(24)]).then(t.bind(null,"oX++"))}},{path:"/Wsx1NbcOddsInfoSet",name:"Wsx1NbcOddsInfoSet",component:function(){return Promise.all([t.e(0),t.e(20)]).then(t.bind(null,"r2iq"))}},{path:"/Wsx2BbcOddsInfoSet",name:"Wsx2BbcOddsInfoSet",component:function(){return Promise.all([t.e(0),t.e(16)]).then(t.bind(null,"BCxU"))}},{path:"/Wsx3ZbcOddsInfoSet",name:"Wsx3ZbcOddsInfoSet",component:function(){return Promise.all([t.e(0),t.e(5)]).then(t.bind(null,"r2uM"))}},{path:"/Zg28BjOddsInfoSet",name:"Zg28BjOddsInfoSet",component:function(){return Promise.all([t.e(0),t.e(18)]).then(t.bind(null,"gLYM"))}},{path:"/DrawMoneyApplyList",name:"DrawMoneyApplyList",component:function(){return Promise.all([t.e(0),t.e(4)]).then(t.bind(null,"csn1"))}},{path:"/SaveMoneyApplyList",name:"SaveMoneyApplyList",component:function(){return Promise.all([t.e(0),t.e(14)]).then(t.bind(null,"6Qng"))}},{path:"/DrawMoneyList",name:"DrawMoneyList",component:function(){return Promise.all([t.e(0),t.e(12)]).then(t.bind(null,"YXY/"))}},{path:"/SaveMoneyList",name:"SaveMoneyList",component:function(){return Promise.all([t.e(0),t.e(7)]).then(t.bind(null,"Cid2"))}},{path:"/GameUserList",name:"GameUserList",component:function(){return Promise.all([t.e(0),t.e(8)]).then(t.bind(null,"m/1i"))}},{path:"/GroupBetList",name:"GroupBetList",component:function(){return Promise.all([t.e(0),t.e(3)]).then(t.bind(null,"hR0A"))}},{path:"/Login",name:"登录1",component:function(){return t.e(1).then(t.bind(null,"nb0W"))}},{path:"/T1",name:"T1",component:function(){return t.e(17).then(t.bind(null,"vobD"))}},{path:"/T2",name:"T2",component:function(){return Promise.all([t.e(0),t.e(6)]).then(t.bind(null,"7fWg"))}},{path:"/LoFilter",name:"LoFilter",component:function(){return t.e(9).then(t.bind(null,"eii1"))}},{path:"/MyCenter",name:"MyCenter",component:function(){return Promise.all([t.e(0),t.e(21)]).then(t.bind(null,"DhIk"))}}]})},uTBe:function(n,e){}},["NHnr"]);
//# sourceMappingURL=app.4f64a63d44705ceeb3c3.js.map