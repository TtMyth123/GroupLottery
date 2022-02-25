<template>
    <div>
        <el-form ref="form" :inline="true" :model="from" label-width="100px">
            <el-form-item label="用户名" width="100">
                <el-input v-model="from.UserName" readonly></el-input>
            </el-form-item>
            <el-form-item label="昵称" width="100">
                <el-input v-model="from.Nickname" ></el-input>
            </el-form-item>
            <el-form-item label="推荐码">
                <el-input v-model="from.ReferrerCode" readonly></el-input>
            </el-form-item>

            <el-form-item label="用户类型">
                <el-select v-model.number="from.UserType" placeholder="请选择">
                    <el-option v-for="item in UserTypeAll" :label="item.name" :key="item.value"
                               :value="item.value"></el-option>
                </el-select>
            </el-form-item>

          <el-form-item label="是否推手">
            <el-select v-model.number="from.IsReferrer" placeholder="请选择">
              <el-option v-for="item in ReferrerTypeAll" :label="item.name" :key="item.value"
                         :value="item.value"></el-option>
            </el-select>
          </el-form-item>

            <el-divider></el-divider>
            <el-form-item label="余额">
                <el-input v-model="from.Gold" readonly></el-input>
            </el-form-item>
            <el-form-item label="累计充值金币">
                <el-input v-model="from.SumSaveMoney" readonly></el-input>
            </el-form-item>
            <el-form-item label="最大充值金币">
                <el-input v-model="from.MaxSaveMoney" readonly></el-input>
            </el-form-item>
            <el-form-item label="累计提现金币">
                <el-input v-model="from.SumDrawMoney" readonly></el-input>
            </el-form-item>
            <el-form-item label="最大提现金币">
                <el-input v-model="from.MaxDrawMoney" readonly></el-input>
            </el-form-item>

            <el-form-item label="投注金额">
                <el-input v-model="from.SumBet" readonly></el-input>
            </el-form-item>
            <el-form-item label="赢得金额">
                <el-input v-model="from.SumWin" readonly></el-input>
            </el-form-item>
            <el-form-item label="累计上分金额">
                <el-input v-model="from.SumAddMoney" readonly></el-input>
            </el-form-item>
            <el-form-item label="累计下分金额">
                <el-input v-model="from.SumDecMoney" readonly></el-input>
            </el-form-item>

            <el-divider></el-divider>
            <el-form-item label="身份证号">
                <el-input v-model="from.IdentityCard"></el-input>
            </el-form-item>
            <el-form-item label="姓名">
                <el-input v-model="from.FullName"></el-input>
            </el-form-item>
            <el-form-item label="电话">
                <el-input v-model="from.Tel"></el-input>
            </el-form-item>
            <el-form-item label="实名状态">
                <el-input v-model="from.RealNameState"></el-input>
            </el-form-item>


            <el-divider></el-divider>
            <el-form-item label="微信收款码Url">
                <el-input v-model="from.WXSKCodeUrl"></el-input>
            </el-form-item>

            <el-form-item label="银行名">
                <el-input v-model="from.YHName"></el-input>
            </el-form-item>
            <el-form-item label="银行卡号">
                <el-input v-model="from.CardNum"></el-input>
            </el-form-item>
            <el-form-item label="银行用户名">
                <el-input v-model="from.YHUserName"></el-input>
            </el-form-item>
            <el-form-item label="银行预留电话">
                <el-input v-model="from.YHUserTel"></el-input>
            </el-form-item>
            <el-form-item label="银行预留地址">
                <el-input v-model="from.Addr"></el-input>
            </el-form-item>
            <el-form-item label="银行预留信息">
                <el-input v-model="from.Remark"></el-input>
            </el-form-item>
            <el-form-item label="支付宝二维码Url">
                <el-input v-model="from.ZFBSKCodeUrl"></el-input>
            </el-form-item>
            <el-form-item label="支付宝名">
                <el-input v-model="from.ZFBSKName"></el-input>
            </el-form-item>


        </el-form>
        <span slot="footer" class="dialog-footer">
            <el-button @click="handleCancel">取 消</el-button>
            <el-button @click="doModifyUserInfo">确 定</el-button>
        </span>
    </div>
</template>

<script>
    import { request } from '../../utils/http';

    export default {
        name: 'UserDetailInfo',
        data() {
            return {
              UserTypeAll: [
                { name: '一般玩家', value: 1 },
                { name: '业务玩家', value: 2 }
              ],
              ReferrerTypeAll: [
                { name: '是', value: 1 },
                { name: '否', value: 0 }
              ],
                UserStateAll: [
                    { name: '启动', value: 1 },
                    { name: '禁用', value: 2 }
                ],
                MemberLevelAll: [
                    { name: '新手会员', value: 1 },
                    { name: '白银会员', value: 2 },
                    { name: '黄金会员', value: 3 },
                    { name: '钻石会员', value: 4 },
                    { name: '皇冠会员', value: 5 }
                ],
                TeamLevelAll: [
                    { name: '', value: 0 },
                    { name: '一级团长', value: 1 },
                    { name: '二级团长', value: 2 },
                    { name: '三级团长', value: 3 },
                    { name: '四级团长', value: 4 },
                    { name: '五级团长', value: 5 }
                ],
                RealNameStateAll: [
                    { name: '未提交', value: 0 },
                    { name: '待审核', value: 1 },
                    { name: '已审核', value: 2 }
                ],

                from: {},
                query: { UserId: 0 }
            };
        },
        props: {
            userId: {
                type: Number,
                default: () => {
                    return {};
                }
            },
            layerid: {
                type: String,
                default: ''
            }
        },
        methods: {
            getData() {
                request({ url: 'api/getuserdetailinfo', method: 'post', data: this.query }).then((res) => {
                    console.log(res);
                    if (res.code == 200) {
                        this.from = res.obj;
                    }
                });
            },
            handleCancel() {
                this.$layer.close(this.layerid);
            },
            doModifyUserInfo() {
                request({ url: 'api/saveuserdetailinfo', method: 'post', data: this.from }).then((res) => {
                    console.log(res);
                    if (res.code == 200) {
                        this.$message.success('保存成功');
                    } else {
                        this.$message.error('保存失败：' + res.msg);
                    }
                });
            }
        },
        mounted() {

            this.query.UserId = this.userId;
            this.getData();
        }
    };
</script>

<style scoped>

</style>