import Vue from 'vue';
import Router from 'vue-router';

Vue.use(Router);


export default new Router({
    routes: [
        {
            path: '/login',
            component: () => import( '../page/RegLogin/Login.vue'),
            meta: { title: '登录' }
        },
        {
            path: '/',
            redirect: '/home'
        },
        {
            path: '/',
            component: () => import( '../components/common/Home.vue'),
            // component: () => import( '../layout/index.vue'),
            meta: { title: '自述文件' },
            children: [
                {
                    path: '/FinanceAccount',
                    component: () => import( '../page/Finance/FinanceAccount.vue'),
                    meta: { title: '财务收款信息' }
                },
                {
                    path: '/SysUserList',
                    component: () => import( '../page/Sys/SysUserList.vue'),
                    meta: { title: '系统用户管理' }
                },
                {
                    path: '/AccountList',
                    component: () => import( '../page/Finance/AccountList.vue'),
                    meta: { title: '流水列表' }
                },
                {
                    path: '/DrawMoneyApplyList',
                    component: () => import( '../page/Finance/DrawMoneyApplyList.vue'),
                    meta: { title: '提现申请管理' }
                },
                {
                    path: '/DrawMoneyList',
                    component: () => import( '../page/Finance/DrawMoneyList.vue'),
                    meta: { title: '提现列表' }
                },
                {
                    path: '/SaveMoneyApplyList',
                    component: () => import( '../page/Finance/SaveMoneyApplyList.vue'),
                    meta: { title: '充值申请管理' }
                },
                {
                    path: '/SaveMoneyList',
                    component: () => import( '../page/Finance/SaveMoneyList.vue'),
                    meta: { title: '充值列表' }
                },
                {
                    path: '/RebateList',
                    component: () => import( '../page/Finance/RebateList.vue'),
                    meta: { title: '返利流水' }
                },
                {
                    path: '/home',
                    component: () => import( '../page/Dashboard/home.vue'),
                    meta: { title: '首页' }
                },
                {
                    path: '/UserList',
                    component: () => import( '../page/GameUser/UserList.vue'),
                    meta: { title: '用户列表' }
                },
                {
                    path: '/404',
                    component: () => import(/* webpackChunkName: "404" */ '../page/404.vue'),
                    meta: { title: '404' }
                },
                {
                    path: '/403',
                    component: () => import(/* webpackChunkName: "403" */ '../page/403.vue'),
                    meta: { title: '403' }
                },
                {
                    path: '/Wsx1NbcOddsInfoSet',
                    component: () => import( '../page/Set/Wsx1NbcOddsInfoSet.vue'),
                    meta: { title: '南部彩赔率' }
                },
                {
                    path: '/Wsx2BbcOddsInfoSet',
                    component: () => import( '../page/Set/Wsx2BbcOddsInfoSet.vue'),
                    meta: { title: '北部彩赔率' }
                },
                {
                    path: '/Wsx3ZbcOddsInfoSet',
                    component: () => import( '../page/Set/Wsx3ZbcOddsInfoSet.vue'),
                    meta: { title: '中部彩赔率' }
                },///
                {
                    path: '/Usc5Cqssc103OdssInfoSet',
                    component: () => import( '../page/Set/Usc5Cqssc103OdssInfoSet.vue'),
                    meta: { title: '重庆时时彩赔率' }
                },
                {
                    path: '/Usc5Jsssc111OdssInfoSet',
                    component: () => import( '../page/Set/Usc5Jsssc111OdssInfoSet.vue'),
                    meta: { title: '极速时时彩赔率' }
                },
                {
                    path: '/Usc5Ygcyc114OdssInfoSet',
                    component: () => import( '../page/Set/Usc5Ygcyc114OdssInfoSet.vue'),
                    meta: { title: '英国幸运彩赔率' }
                },
                {
                    path: '/Usc5Ygssc120OdssInfoSet',
                    component: () => import( '../page/Set/Usc5Ygssc120OdssInfoSet.vue'),
                    meta: { title: '英国时时彩赔率' }
                },
                {
                    path: '/Usc5Gzxy5116OdssInfoSet',
                    component: () => import( '../page/Set/Usc5Gzxy5116OdssInfoSet.vue'),
                    meta: { title: '澳洲幸运5赔率' }
                },
                {
                    path: '/Usc5Yxssc118OdssInfoSet',
                    component: () => import( '../page/Set/Usc5Yxssc118OdssInfoSet.vue'),
                    meta: { title: '腾讯分分彩赔率' }
                },
                {
                    path: '/Usc8Cqxync107OdssInfoSet',
                    component: () => import( '../page/Set/Usc8Cqxync107OdssInfoSet.vue'),
                    meta: { title: '重庆幸运农场赔率' }
                },
                {
                    path: '/Usc8Gdkl10f102OdssInfoSet',
                    component: () => import( '../page/Set/Usc8Gdkl10f102OdssInfoSet.vue'),
                    meta: { title: '广东快乐十分赔率' }
                },
                {
                    path: '/Usc10Bjsc104OdssInfoSet',
                    component: () => import( '../page/Set/Usc10Bjsc104OdssInfoSet.vue'),
                    meta: { title: '北京赛车赔率' }
                },
                {
                    path: '/Usc10Xyft108OdssInfoSet',
                    component: () => import( '../page/Set/Usc10Xyft108OdssInfoSet.vue'),
                    meta: { title: '幸运飞艇赔率' }
                },
                {
                    path: '/Usc10Jskc109OdssInfoSet',
                    component: () => import( '../page/Set/Usc10Jskc109OdssInfoSet.vue'),
                    meta: { title: '极速快车赔率' }
                },
                {
                    path: '/Usc10Jssc112OdssInfoSet',
                    component: () => import( '../page/Set/Usc10Jssc112OdssInfoSet.vue'),
                    meta: { title: '极速赛车赔率' }
                },
                {
                    path: '/Usc10Espsm113OdssInfoSet',
                    component: () => import( '../page/Set/Usc10Espsm113OdssInfoSet.vue'),
                    meta: { title: 'ESP赛马赔率' }
                },
                {
                    path: '/Usc10Ygxyft115OdssInfoSet',
                    component: () => import( '../page/Set/Usc10Ygxyft115OdssInfoSet.vue'),
                    meta: { title: '英国幸运飞艇赔率' }
                },
                {
                    path: '/Usc10Ygsc119OdssInfoSet',
                    component: () => import( '../page/Set/Usc10Ygsc119OdssInfoSet.vue'),
                    meta: { title: '英国赛车赔率' }
                },
                {
                    path: '/Usc10Gzxy10117OdssInfoSet',
                    component: () => import( '../page/Set/Usc10Gzxy10117OdssInfoSet.vue'),
                    meta: { title: '澳洲幸运10赔率' }
                },
                {
                    path: '/GroupBetList',
                    component: () => import( '../page/Report/GroupBetList.vue'),
                    meta: { title: '主投注记录' }
                },
                {
                    path: '/Wsx1AwardList',
                    component: () => import( '../page/AwardResult/Wsx1AwardList.vue'),
                    meta: { title: '南部彩开奖结果' }
                },
                {
                    path: '/Wsx2AwardList',
                    component: () => import( '../page/AwardResult/Wsx2AwardList.vue'),
                    meta: { title: '北部彩开奖结果' }
                },
                {
                    path: '/Wsx3AwardList',
                    component: () => import( '../page/AwardResult/Wsx3AwardList.vue'),
                    meta: { title: '中部彩开奖结果' }
                },
                {
                    path: '/Zg28JndAwardList',
                    component: () => import( '../page/AwardResult/Zg28JndAwardList.vue'),
                    meta: { title: '加拿大28开奖结果' }
                },
                {
                    path: '/Zg28BjAwardList',
                    component: () => import( '../page/AwardResult/Zg28BjAwardList.vue'),
                    meta: { title: '北京28开奖结果' }
                },
                {
                    path: '/Zg28XjpAwardList',
                    component: () => import( '../page/AwardResult/Zg28XjpAwardList.vue'),
                    meta: { title: '新加坡28开奖结果' }
                },
                {
                    path: '/Zg28XgAwardList',
                    component: () => import( '../page/AwardResult/Zg28XgAwardList.vue'),
                    meta: { title: '香港28开奖结果' }
                },
                {
                    path: '/SetZg28XgAwardList',
                    component: () => import( '../page/AwardResult/SetZg28XgAwardList.vue'),
                    meta: { title: '香港28开奖结果' }
                },
                {
                    path: '/UscCqssc103AwardList',
                    component: () => import( '../page/AwardResult/UscCqssc103AwardList.vue'),
                    meta: { title: '重庆时时彩开奖结果' }
                },//////
                {
                    path: '/UscJsssc111AwardList',
                    component: () => import( '../page/AwardResult/UscJsssc111AwardList.vue'),
                    meta: { title: '极速时时彩开奖结果' }
                },
                {
                    path: '/UscYgcyc114AwardList',
                    component: () => import( '../page/AwardResult/UscYgcyc114AwardList.vue'),
                    meta: { title: '英国幸运彩' }
                },
                {
                    path: '/UscYgssc120AwardList',
                    component: () => import( '../page/AwardResult/UscYgssc120AwardList.vue'),
                    meta: { title: '英国时时彩开奖结果' }
                },
                {
                    path: '/UscGzxy5116AwardList',
                    component: () => import( '../page/AwardResult/UscGzxy5116AwardList.vue'),
                    meta: { title: '澳洲幸运5开奖结果' }
                },
                {
                    path: '/UscYxssc118AwardList',
                    component: () => import( '../page/AwardResult/UscYxssc118AwardList.vue'),
                    meta: { title: '腾讯分分彩开奖结果' }
                },
                {
                    path: '/UscBjsc104AwardList',
                    component: () => import( '../page/AwardResult/UscBjsc104AwardList.vue'),
                    meta: { title: '北京赛车开奖结果' }
                },
                {
                    path: '/UscXyft108AwardList',
                    component: () => import( '../page/AwardResult/UscXyft108AwardList.vue'),
                    meta: { title: '幸运飞艇开奖结果' }
                },
                {
                    path: '/UscJskc109AwardList',
                    component: () => import( '../page/AwardResult/UscJskc109AwardList.vue'),
                    meta: { title: '极速快车开奖结果' }
                },
                {
                    path: '/UscJssc112AwardList',
                    component: () => import( '../page/AwardResult/UscJssc112AwardList.vue'),
                    meta: { title: '极速赛车开奖结果' }
                },
                {
                    path: '/UscEspsm113AwardList',
                    component: () => import( '../page/AwardResult/UscEspsm113AwardList.vue'),
                    meta: { title: 'ESP赛马开奖结果' }
                },
                {
                    path: '/UscYgxyft115AwardList',
                    component: () => import( '../page/AwardResult/UscYgxyft115AwardList.vue'),
                    meta: { title: '英国幸运飞艇开奖结果' }
                },
                {
                    path: '/UscYgsc119AwardList',
                    component: () => import( '../page/AwardResult/UscYgsc119AwardList.vue'),
                    meta: { title: '英国赛车开奖结果' }
                },
                {
                    path: '/UscGdkl10f102AwardList',
                    component: () => import( '../page/AwardResult/UscGdkl10f102AwardList.vue'),
                    meta: { title: '澳洲幸运10开奖结果' }
                },
                {
                    path: '/UscCqxync107AwardList',
                    component: () => import( '../page/AwardResult/UscCqxync107AwardList.vue'),
                    meta: { title: '重庆幸运农场开奖结果' }
                },
                {
                    path: '/UscGdkl10f102AwardList',
                    component: () => import( '../page/AwardResult/UscGdkl10f102AwardList.vue'),
                    meta: { title: '广东快乐十分开奖结果' }
                },
                {
                    path: '/Zg28BjSet',
                    component: () => import( '../page/Set/Zg28BjSet.vue'),
                    meta: { title: '北京28赔率' }
                },
                {
                    path: '/Zg28JndSet',
                    component: () => import( '../page/Set/Zg28JndSet.vue'),
                    meta: { title: '加拿大28赔率' }
                },
                {
                    path: '/Zg28XjpSet',
                    component: () => import( '../page/Set/Zg28XjpSet.vue'),
                    meta: { title: '新加坡28赔率' }
                },
                {
                    path: '/Zg28XgSet',
                    component: () => import( '../page/Set/Zg28XgSet.vue'),
                    meta: { title: '香港28赔率' }
                },

                {
                    path: '/RebateSet',
                    component: () => import( '../page/Finance/RebateSet.vue'),
                    meta: { title: '返佣设置' }
                },
                {
                    path: '/NoticeList',
                    component: () => import( '../page/Article/NoticeList.vue'),
                    meta: { title: '公告列表' }
                },
                {
                    path: '/SlideshowList',
                    component: () => import( '../page/Article/SlideshowList.vue'),
                    meta: { title: '幻灯片列表' }
                },
            ]
        },
        {
            path: '/Root404',
            component: () => import( '../page/Root404.vue'),
            meta: { title: '' }
        },
        {
            path: '*',
            redirect: '/404'
        },
        {
            path: '/LineMarker',
            component: () => import( '../components/Charts/LineMarker.vue'),
            meta: { title: '' }
        },
        {
            path: '/aaa',
            component: () => import( '../mpage/RegLogin/Login.vue'),
            meta: { title: '' }
        }
    ]
});
