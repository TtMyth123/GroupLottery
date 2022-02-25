<template>
    <div>
        <el-form ref="form" :inline="true" :model="from" label-width="100px">
            <el-form-item label="Tên người dùng" width="100">
                <el-input v-model="from.UserName" readonly></el-input>
            </el-form-item>
            <el-form-item label="Tên nink" width="100">
                <el-input v-model="from.Nickname" ></el-input>
            </el-form-item>
            <el-form-item label="Mã giới thiệu">
                <el-input v-model="from.ReferrerCode" readonly></el-input>
            </el-form-item>

            <el-form-item label="Thể loại người dùng">
                <el-select v-model.number="from.UserType" placeholder="Vui lòng chọn">
                    <el-option v-for="item in UserTypeAll" :label="item.name" :key="item.value"
                               :value="item.value"></el-option>
                </el-select>
            </el-form-item>

          <el-form-item label="Làm cò">
            <el-select v-model.number="from.IsReferrer" placeholder="Vui lòng chọn">
              <el-option v-for="item in ReferrerTypeAll" :label="item.name" :key="item.value"
                         :value="item.value"></el-option>
            </el-select>
          </el-form-item>

            <el-divider></el-divider>
            <el-form-item label="Số dư">
                <el-input v-model="from.Gold" readonly></el-input>
            </el-form-item>
            <el-form-item label="Tông nạp vào xu vàng">
                <el-input v-model="from.SumSaveMoney" readonly></el-input>
            </el-form-item>
            <el-form-item label="Mưc hạn nạp xu vàng">
                <el-input v-model="from.MaxSaveMoney" readonly></el-input>
            </el-form-item>
            <el-form-item label="Tổng rút ra xu vàng">
                <el-input v-model="from.SumDrawMoney" readonly></el-input>
            </el-form-item>
            <el-form-item label="Mức hạn rút xu vàng">
                <el-input v-model="from.MaxDrawMoney" readonly></el-input>
            </el-form-item>

            <el-form-item label="Đặt cược số điểm">
                <el-input v-model="from.SumBet" readonly></el-input>
            </el-form-item>
            <el-form-item label="Được thắng số điểm">
                <el-input v-model="from.SumWin" readonly></el-input>
            </el-form-item>
            <el-form-item label="Tông sổ lên điểm">
                <el-input v-model="from.SumAddMoney" readonly></el-input>
            </el-form-item>
            <el-form-item label="Tổng sổ xuống điểm">
                <el-input v-model="from.SumDecMoney" readonly></el-input>
            </el-form-item>

            <el-divider></el-divider>
            <el-form-item label="Số chứng mình thư">
                <el-input v-model="from.IdentityCard"></el-input>
            </el-form-item>
            <el-form-item label="Hộ tên">
                <el-input v-model="from.FullName"></el-input>
            </el-form-item>
            <el-form-item label="Số di động">
                <el-input v-model="from.Tel"></el-input>
            </el-form-item>
            <el-form-item label="trạng thái tên thật">
                <el-input v-model="from.RealNameState"></el-input>
            </el-form-item>


            <el-divider></el-divider>
            <el-form-item label="Mã QR nhận tiền Url">
                <el-input v-model="from.WXSKCodeUrl"></el-input>
            </el-form-item>

            <el-form-item label="Tên ngân hàng">
                <el-input v-model="from.YHName"></el-input>
            </el-form-item>
            <el-form-item label="Số thẻ ngân hàng">
                <el-input v-model="from.CardNum"></el-input>
            </el-form-item>
            <el-form-item label="Hộ tên của bạn">
                <el-input v-model="from.YHUserName"></el-input>
            </el-form-item>
            <el-form-item label="Số di động lưu ngân hàng">
                <el-input v-model="from.YHUserTel"></el-input>
            </el-form-item>
            <el-form-item label="Địa chị lưu ngân hàng">
                <el-input v-model="from.Addr"></el-input>
            </el-form-item>
            <el-form-item label="Lưu tin nhắn trông ngân hàng">
                <el-input v-model="from.Remark"></el-input>
            </el-form-item>
            <el-form-item label="Mã QR chuyển tiền Url">
                <el-input v-model="from.ZFBSKCodeUrl"></el-input>
            </el-form-item>
            <el-form-item label="Tên của mã QR">
                <el-input v-model="from.ZFBSKName"></el-input>
            </el-form-item>


        </el-form>
        <span slot="footer" class="dialog-footer">
            <el-button @click="handleCancel">Huỷ</el-button>
            <el-button @click="doModifyUserInfo">Xác nhậ</el-button>
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
                { name: 'Khách thường', value: 1 },
                { name: 'Khách chào hàng', value: 2 }
              ],
              ReferrerTypeAll: [
                { name: 'Phải', value: 1 },
                { name: 'Không', value: 0 }
              ],
                UserStateAll: [
                    { name: 'Khời dùng', value: 1 },
                    { name: 'Cấm dùng', value: 2 }
                ],
                MemberLevelAll: [
                    { name: 'Thành viên mới', value: 1 },
                    { name: 'Thành viên bạc', value: 2 },
                    { name: 'Thành viên vàng', value: 3 },
                    { name: 'Thành viên kim cương', value: 4 },
                    { name: 'Thành viên crown', value: 5 }
                ],
                TeamLevelAll: [
                    { name: '', value: 0 },
                    { name: 'Đoàn trưởng cấp 1', value: 1 },
                    { name: 'Đoàn trưởng cấp 2', value: 2 },
                    { name: 'Đoàn trưởng cấp 3', value: 3 },
                    { name: 'Đoàn trưởng cấp 4', value: 4 },
                    { name: 'Đoàn trưởng cấp 5', value: 5 }
                ],
                RealNameStateAll: [
                    { name: 'Chưa gửi lên', value: 0 },
                    { name: 'Đợi sự lý', value: 1 },
                    { name: 'Đã sự lý', value: 2 }
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
                        this.$message.success('Lưu thành công');
                    } else {
                        this.$message.error('Lưu không thành côn：' + res.msg);
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