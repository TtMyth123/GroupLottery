import Vue from 'vue';

Vue.filter('MemberLevelName', function(value) {
    return value === 1 ? '新手会员' : value === 2 ? '白银会员' : value === 3 ? '黄金会员' : value === 4 ? '钻石会员' : value === 5 ? '皇冠会员' : '';
});

Vue.filter('TeamLevelName', function(value) {
    return value === 1 ? '一级团长' : value === 2 ? '二级团长' : value === 3 ? '三级团长' : value === 4 ? '四级团长' : value === 5 ? '五级团长' : '';
});
Vue.filter('UserTypeName', function(value) {
    return value === 1 ? '一般玩家' : value === 2 ? '业务玩家' : '';
});

Vue.filter('FilmStateName', function(value) {
    return value === 1 ? '准备' : value === 2 ? '上架' : '下架';
});

Vue.filter('GoodsStateName', function(value) {
    return value === 1 ? '准备' : value === 2 ? '上架' : '下架';
});

Vue.filter('AwardTypeName', function(value) {
    return value === 1 ? '金币' : value === 2 ? '积分' : '';
});

Vue.filter('BetStatusName', function(value) {
    return value === 1 ? '待开奖' : value === 2 ? '已兑奖' : '';
});



Vue.filter('FilmOrderStateName', function(value) {
    return value === 1 ? '进行中' : value === 2 ? '完成' : '';
});

Vue.filter('SaveStateName', function(value) {
    return value === 1 ? '申请状态' : value === 2 ? '付款中' : value === 3 ? '已付款' : value === 4 ? '已上传凭证' : value === 5 ? '充值成功' : '';
});

Vue.filter('AwardTypeName', function(value) {
    return value === 1 ? '金币' : value === 2 ? '积分' : value === 3 ? '转盘' : '';
});

Vue.filter('StateName', function(value) {
    return value === 1 ? '启用中' : value === 2 ? '禁用中' : '——';
});
Vue.filter('RealNameStateName', function(value) {
    return value === 1 ? '持审核' : value === 2 ? '审核' : '未提交';
});

Vue.filter('money', function(val) {
    return formatMoney(val, 2);
});
Vue.filter('money1', function(val) {
    return formatMoney(val, 1);
});
Vue.filter('money4', function(val) {
    return formatMoney(val, 4);
});

Vue.filter('money3', function(val) {
    return formatMoney(val, 3);
});

function formatMoney(number, decimals = 0, decPoint = '.', thousandsSep = ',') {
    number = (number + '').replace(/[^0-9+-Ee.]/g, '');
    let n = !isFinite(+number) ? 0 : +number;
    let prec = !isFinite(+decimals) ? 0 : Math.abs(decimals);
    let sep = (typeof thousandsSep === 'undefined') ? ',' : thousandsSep;
    let dec = (typeof decPoint === 'undefined') ? '.' : decPoint;
    let s = '';
    let toFixedFix = function(n, prec) {
        let k = Math.pow(10, prec);
        return '' + Math.ceil(n * k) / k;
    };
    s = (prec ? toFixedFix(n, prec) : '' + Math.round(n)).split('.');
    let re = /(-?\d+)(\d{3})/;
    while (re.test(s[0])) {
        s[0] = s[0].replace(re, '$1' + sep + '$2');
    }
    if ((s[1] || '').length < prec) {
        s[1] = s[1] || '';
        s[1] += new Array(prec - s[1].length + 1).join('0');
    }
    return s.join(dec);
}