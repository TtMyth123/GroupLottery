import Vue from 'vue';
import Router from 'vue-router';

Vue.use(Router);


export default new Router({
    routes: [
        {
            path: '/login',
            component: () => import( '../page/RegLogin/Login.vue'),
            meta: { title: 'Đăng nhập' }
        },
        {
            path: '/',
            redirect: '/home'
        },
        {
            path: '/',
            component: () => import( '../components/common/Home.vue'),
            // component: () => import( '../layout/index.vue'),
            meta: { title: 'Tẹp tự báo cáo' },
            children: [
                {
                    path: '/FinanceAccount',
                    component: () => import( '../page/Finance/FinanceAccount.vue'),
                    meta: { title: 'Thông tin thu nhập tài chính' }
                },
                {
                    path: '/SysUserList',
                    component: () => import( '../page/Sys/SysUserList.vue'),
                    meta: { title: 'Quản lý khách hàng hệ thống' }
                },
                {
                    path: '/AccountList',
                    component: () => import( '../page/Finance/AccountList.vue'),
                    meta: { title: 'Danh sách đặt cược' }
                },
                {
                    path: '/DrawMoneyApplyList',
                    component: () => import( '../page/Finance/DrawMoneyApplyList.vue'),
                    meta: { title: 'Quản lý xin phép rút tiền' }
                },
                {
                    path: '/DrawMoneyList',
                    component: () => import( '../page/Finance/DrawMoneyList.vue'),
                    meta: { title: 'Danh sách rút tiền' }
                },
                {
                    path: '/SaveMoneyApplyList',
                    component: () => import( '../page/Finance/SaveMoneyApplyList.vue'),
                    meta: { title: 'Quản lý xin phép nạp tiền' }
                },
                {
                    path: '/SaveMoneyList',
                    component: () => import( '../page/Finance/SaveMoneyList.vue'),
                    meta: { title: 'Danh sách nạp tiền' }
                },
                {
                    path: '/RebateList',
                    component: () => import( '../page/Finance/RebateList.vue'),
                    meta: { title: 'Hoa Hồng đặt cược' }
                },
                {
                    path: '/home',
                    component: () => import( '../page/Dashboard/home.vue'),
                    meta: { title: 'Trang tổng ' }
                },
                {
                    path: '/UserList',
                    component: () => import( '../page/GameUser/UserList.vue'),
                    meta: { title: 'Danh sách khách hàng' }
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
                    meta: { title: 'Số nhân hồ chí minh VIP' }
                },
                {
                    path: '/Wsx2BbcOddsInfoSet',
                    component: () => import( '../page/Set/Wsx2BbcOddsInfoSet.vue'),
                    meta: { title: 'Số nhân miền bắc' }
                },
                {
                    path: '/Wsx3ZbcOddsInfoSet',
                    component: () => import( '../page/Set/Wsx3ZbcOddsInfoSet.vue'),
                    meta: { title: ' Số nhân keno' }
                },
                {
                    path: '/GroupBetList',
                    component: () => import( '../page/Report/GroupBetList.vue'),
                    meta: { title: 'Lịch sử tổng đặt cược' }
                },
                {
                    path: '/Wsx1AwardList',
                    component: () => import( '../page/AwardResult/Wsx1AwardList.vue'),
                    meta: { title: 'Kết quả mở thưởng hồ chí minh VIP' }
                },
                {
                    path: '/Wsx2AwardList',
                    component: () => import( '../page/AwardResult/Wsx2AwardList.vue'),
                    meta: { title: 'Kết quả mở thưởng miền bắc' }
                },
                {
                    path: '/Wsx3AwardList',
                    component: () => import( '../page/AwardResult/Wsx3AwardList.vue'),
                    meta: { title: 'Kết quả mở thưởng keno' }
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
                    path: '/RebateSet',
                    component: () => import( '../page/Finance/RebateSet.vue'),
                    meta: { title: 'Cài đặt hoa hồng' }
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
        },
        {
            path: '/AdminChat',
            component: () => import( '../page/Chat/AdminChat.vue'),
            meta: { title: '' }
        }
    ]
});
