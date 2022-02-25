<template>
    <div>
        <div class="container">
            <div class="handle-box">
              <el-input v-model="query.userName" placeholder="Tên người dùng" class="handle-input mr10"></el-input>
              <el-input v-model="query.userId" placeholder="Tên người dùng Id" class="handle-input mr10"></el-input>
              <el-select v-model.number="query.userType" placeholder="Vui lòng chọn">
                    <el-option v-for="item in UserTypeAll" :label="item.name" :key="item.value"
                               :value="item.value"></el-option>
                </el-select>
                <el-button type="primary" icon="el-icon-search" @click="handleSearch">Tìm kiếm</el-button>
            </div>
            <el-table
                    :data="tableData"
                    border
                    class="table"
                    ref="multipleTable"
                    header-cell-class-name="table-header"
                    :summary-method="getSummaries"
                    show-summary
            >
                <!--                <el-table-column type="selection" width="55" align="center"></el-table-column>-->
                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>
                <el-table-column prop="UserName" label="用户名" width="110"></el-table-column>
                <el-table-column label="Số dư">
                    <template slot-scope="scope">{{scope.row.Gold|money}}</template>
                </el-table-column>

                <el-table-column label="Số tiền đặt cược">
                    <template slot-scope="scope">{{scope.row.SumBet}}</template>
                </el-table-column>

                <el-table-column label="Số tiền thắng">
                    <template slot-scope="scope">{{scope.row.SumWin|money}}</template>
                </el-table-column>

                <el-table-column label="Tông nạp tiền">
                    <template slot-scope="scope">{{scope.row.SumSaveMoney|money}}</template>
                </el-table-column>
                <el-table-column label="Tổng ra tiền">
                    <template slot-scope="scope">{{scope.row.SumDrawMoney|money}}</template>
                </el-table-column>


                <el-table-column label="Tổng ra tiền">
                    <template slot-scope="scope">{{scope.row.UserType|UserTypeName}}</template>
                </el-table-column>

                <el-table-column prop="PUserName" label="Cấp trên người dùng" width="100">
                    <template slot-scope="scope">
                        <el-button
                                type="text"
                                @click="openPUser(scope.$index, scope.row)"
                        >{{scope.row.PUserName}}
                        </el-button>
                    </template>
                </el-table-column>
                <el-table-column prop="LowerCount" label="Dân số cấp dưới">
                    <template slot-scope="scope">
                        <el-button
                                type="text"
                                @click="openJuniorUser(scope.$index, scope.row)"
                        >{{scope.row.LowerCount}}
                        </el-button>
                    </template>
                </el-table-column>
                <el-table-column label="Thời gian đăng ký">
                    <template slot-scope="scope">{{scope.row.CreatedAt}}</template>
                </el-table-column>
                <el-table-column label="Lần cuối đăng nhập">
                    <template slot-scope="scope">{{scope.row.LoginTime}}</template>
                </el-table-column>
                <el-table-column prop="FullName" label="Hộ tên"></el-table-column>
                <el-table-column label="Trạng thái">
                    <template slot-scope="scope">
                        <el-button
                                type="text"
                                @click="openState(scope.$index, scope.row)"
                        >{{scope.row.State|StateName}}
                        </el-button>
                    </template>
                </el-table-column>

                <el-table-column label="Điều hành" width="140" align="center">
                    <template slot-scope="scope">
                        <el-button
                                type="text"
                                icon="el-icon-edit"
                                @click="openUser(scope.$index, scope.row)"
                        >Sửa đổi
                        </el-button>

                        <el-button
                                type="text"
                                icon="el-icon-edit"
                                @click="handleSaveMoney(scope.$index, scope.row)"
                        >Lên điểm
                        </el-button>
                        <el-button
                                type="text"
                                icon="el-icon-edit"
                                @click="handleDrawSaveMoney(scope.$index, scope.row)"
                        >Xuống điểm
                        </el-button>
                        <el-button
                                type="text"
                                icon="el-icon-edit"
                                @click="handleModifyPwd(scope.$index, scope.row)"
                        >Mật khẩu
                        </el-button>

                    </template>
                </el-table-column>
            </el-table>
            <div class="pagination">
                <el-pagination
                        background
                        layout="total, prev, pager, next"
                        :current-page="query.pageIndex"
                        :page-size="query.pageSize"
                        :total="pageTotal"
                        @current-change="handlePageChange"
                ></el-pagination>
            </div>
        </div>

        <el-dialog title="Lên điểm" :visible.sync="fromSaveMoneyVisible" width="90%">
            <el-form ref="form" :model="fromSaveMoney" label-width="100px">
                <el-form-item label="Tên người dùng" width="100">
                    <el-input v-model="fromSaveMoney.UserName" readonly></el-input>
                </el-form-item>
                <el-form-item label="Số điểm lên">
                    <!--                    <el-input v-model.number="fromSaveMoney.Money"></el-input>-->
                    <el-input-number v-model.number="fromSaveMoney.Money" :step="100" :min="1"></el-input-number>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="fromSaveMoneyVisible = false">Huỷ</el-button>
                <el-button type="primary" @click="doSaveMoney">Xác nhận</el-button>
            </span>
        </el-dialog>

        <el-dialog title="Xuống điểm" :visible.sync="fromDrawMoneyVisible" width="90%">
            <el-form ref="form" :model="fromDrawMoney" label-width="100px">
                <el-form-item label="用户名" width="100">
                    <el-input v-model="fromDrawMoney.UserName" readonly></el-input>
                </el-form-item>
                <el-form-item label="Số điểm xuống">
                    <!--                    <el-input v-model.number="fromSaveMoney.Money"></el-input>-->
                    <el-input-number v-model.number="fromDrawMoney.Money" :step="100" :min="1"></el-input-number>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="fromDrawMoneyVisible = false">Huỷ</el-button>
                <el-button type="primary" @click="doDrawMoney">Xác nhận</el-button>
            </span>
        </el-dialog>

        <!-- Chỉnh sửa cửa sổ bật lên -->
        <el-dialog title="Lượt xem" :visible.sync="editVisible" width="90%">
            <el-form ref="form" :model="form" label-width="100px">
                <el-form-item label="Tên người dùng" width="100">
                    <el-input v-model="form.UserName" readonly></el-input>
                </el-form-item>

                <el-form-item label="Tên nick">
                    <el-input v-model="form.Nickname"></el-input>
                </el-form-item>

                <el-form-item label="Trạng thái" prop="State">
                    <el-select v-model.number="form.State" placeholder="Vui lòng chọn">
                        <el-option v-for="item in StateTypes" :label="item.name" :key="item.value"
                                   :value="item.value"></el-option>
                    </el-select>
                </el-form-item>

                <el-form-item label="Cấp độ thành viên" prop="MemberLevel">
                    <el-select v-model.number="form.MemberLevel" placeholder="Vui lòng chọn">
                        <el-option v-for="item in MemberLevelTypes" :label="item.name" :key="item.value"
                                   :value="item.value"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="Cấp đọ đoàn" prop="TeamLevel">
                    <el-select v-model.number="form.TeamLevel" placeholder="Vui lòng chọn">
                        <el-option v-for="item in TeamLevelTypes" :label="item.name" :key="item.value"
                                   :value="item.value"></el-option>
                    </el-select>
                </el-form-item>

            </el-form>

            <span slot="footer" class="dialog-footer">
                <el-button @click="editVisible = false">Huỷ</el-button>
            </span>
        </el-dialog>

        <el-dialog title="Sửa đổi mật khẩu" :visible.sync="fromModifyPwdVisible" width="90%">
            <el-form ref="form" :model="fromModifyPwd" label-width="100px">
                <el-form-item label="Tên người dùng" width="100">
                    <el-input v-model="fromModifyPwd.UserName" readonly></el-input>
                </el-form-item>

                <el-form-item label="Mật khẩu">
                    <el-input v-model.number="fromModifyPwd.Pwd"></el-input>
                </el-form-item>
            </el-form>

            <span slot="footer" class="dialog-footer">
                <el-button @click="handleModifyPwd = false">Huỷ</el-button>
                <el-button @click="doModifyPwd">Xác nhận</el-button>
            </span>
        </el-dialog>

    </div>
</template>
<script>
    import VueCropper from 'vue-cropperjs';
    import { request } from '../../utils/http';
    import JuniorUserListComp from './JuniorUserList.vue';
    import UserDetailInfoComp from './UserDetailInfo.vue';

    export default {
        name: 'UserList',
        data() {
            return {
                defaultAvatarSrc: require('../../assets/img/default_avatar.jpg'),
                dialogVisible: false,
                saveMoneyVisible: false,
                StateTypes: [
                    { name: 'Khời dùng', value: 1 },
                    { name: 'Cấm dùng', value: 2 }
                ],
                MemberLevelTypes: [
                    { name: 'Thành viên mới', value: 1 },
                    { name: 'Thành viên bạc', value: 2 },
                    { name: 'Thành viên vàng', value: 3 },
                    { name: 'Thành viên kim cương', value: 4 },
                    { name: 'Thành viên crown', value: 5 }
                ],
                TeamLevelTypes: [
                    { name: '', value: 0 },
                    { name: 'Đoàn trưởng cấp 1', value: 1 },
                    { name: 'Đoàn trưởng cấp 2', value: 2 },
                    { name: 'Đoàn trưởng cấp 3', value: 3 },
                    { name: 'Đoàn trưởng cấp 4', value: 4 },
                    { name: 'Đoàn trưởng cấp 5', value: 5 }
                ],
                UserTypeAll: [
                    { name: 'Toàn bộ', value: 0 },
                    { name: 'Khách thường', value: 1 },
                    { name: 'Khách chào hàng', value: 2 }
                ],
                query: {
                  userId:0,
                    userType: 0,
                    state: 0,
                    userName: '',
                    pageIndex: 1,
                    pageSize: 10
                },
                tableData: [],
                GroupData: {Gold:0,SumBet:0,SumWin:0,SumSaveMoney:0,MaxDrawMoney:0},
                multipleSelection: [],
                delList: [],
                editVisible: false,
                pageTotal: 0,
                form: {},
                idx: -1,
                id: -1,

                fromSaveMoneyVisible: false,
                fromSaveMoney: { UserId: 0, UserName: '', Money: 0 },

                fromDrawMoneyVisible: false,
                fromDrawMoney: { UserId: 0, UserName: '', Money: 0 },

                fromModifyPwdVisible: false,
                fromModifyPwd: { UserId: 0, Pwd: '', UserName: '' }
            };
        },
        created() {
            this.getData();
        },
        mounted() {
            // this.editor = this.$refs.myQuillEditor.quill;

        },
        beforeDestroy() {
            this.form = this.info;
            // this.editor = null;
            // delete this.editor;
        },
        components: {
            VueCropper
        },
        methods: {
            getSummaries() {
                const sums = [];
                sums[2] = this.GroupData.Gold.toFixed(2);
                sums[3] = this.GroupData.SumBet;
                sums[4] = this.GroupData.SumWin.toFixed(2);
                sums[5] = this.GroupData.SumSaveMoney;
                sums[6] = this.GroupData.MaxDrawMoney;

                return sums;
            },
            openUser(index, row) {
                const id = this.$layer.iframe({
                    content: {
                        content: UserDetailInfoComp,
                        parent: this,
                        data: { userId: row.Id }
                    },
                    area: ['100%', '500px'],
                    title: 'Sửa đổi thông tin khách',
                    maxmin: true,
                    shade: true,
                    shadeClose: false,
                    scrollbar: true,
                    resize: true,
                    cancel: () => {
                        this.$layer.close(this.layerid);
                    }
                });
            },
            openPUser(index, row) {
                const id = this.$layer.iframe({
                    content: {
                        content: UserDetailInfoComp,
                        parent: this,
                        data: { userId: row.Pid }
                    },
                    area: ['100%', '500px'],
                    title: '这是一个标题这是一个标题这是一个标题这是一个标题',
                    maxmin: true,
                    shade: true,
                    shadeClose: false,
                    scrollbar: true,
                    resize: true,
                    cancel: () => {
                        this.$layer.close(this.layerid);
                    }
                });
            },
            openJuniorUser(index, row) {
                const id = this.$layer.iframe({
                    content: {
                        content: JuniorUserListComp,
                        parent: this,
                        data: { pid: row.Id }
                    },
                    area: ['100%', '800px'],
                    title: 'Sửa đổi thông tin khách',
                    maxmin: true,
                    shade: true,
                    shadeClose: false,
                    scrollbar: true,
                    resize: true,
                    cancel: () => {
                        console.log(id);
                        //this.$layer.close(this.layerid);
                        this.$layer.close(id);
                    }
                }, layerid => {
                    // this.$layer.close(this.layerid);
                });
            },
            openState(index, row) {
                let hint = 'Chắc chắn đề bắt đầu' + row.UserName + 'Tên người dùng？';
                let State = 1;
                if (row.State == 1) {
                    hint = 'Chắc chắn đề tắt' + row.UserName + 'Tên người dùng？';
                    State = 2;
                }

                this.$confirm(hint, 'Lời nhắc', {
                    confirmButtonText: 'Xác nhận',
                    cancelButtonText: 'Huỷ',
                    type: 'warning'
                }).then(() => {
                    let data = {
                        UserId: row.Id,
                        State: State
                    };
                    request({ url: 'api/setuserstate', method: 'post', data: data }).then((res) => {
                        if (res.code == 200) {
                            this.$message({
                                type: 'success',
                                message: 'Cài đặt thành công!'
                            });
                            row.State = State;
                        }
                    });
                }).catch(() => {
                    // this.$message({
                    //     type: 'info',
                    //     message: 'Không bị xóa'
                    // });
                });
            },

            getData() {
                request({ url: 'api/getgameuserlist', method: 'post', data: this.query }).then((res) => {
                    if (res.code == 200) {
                        this.tableData = res.obj.ListData;
                        this.pageTotal = res.obj.PageTotal;
                        this.GroupData = res.obj.GroupData;
                    }
                });
            },
            // Kích hoạt nút tìm kiếm
            handleSearch() {
                this.$set(this.query, 'pageIndex', 1);
                this.getData();
            },
            doSaveMoney() {
                request({
                    url: 'api/savemoney', method: 'post',
                    data: this.fromSaveMoney
                }).then((res) => {
                    if (res.code == 200) {
                        this.$set(this.tableData, this.idx, res.obj);
                        this.$message.success(`lên điểm thành công`);
                        this.getData();
                        this.fromSaveMoneyVisible = false;
                    } else {
                        this.$message.error('Lên điểm không thành công：' + res.msg);
                    }
                });
            },
            doDrawMoney() {
                request({
                    url: 'api/drawmoney', method: 'post',
                    data: this.fromDrawMoney
                }).then((res) => {
                    if (res.code == 200) {
                        this.$message.success(`Xuống điểm thành công`);
                        this.getData();
                        this.fromDrawMoneyVisible = false;
                    } else {
                        this.$message.error('Xuống điểm không thành công：' + res.msg);
                    }
                });
            },
            doModifyPwd() {
                request({
                    url: 'api/updateuserpwd', method: 'post',
                    data: this.fromModifyPwd
                }).then((res) => {
                    if (res.code == 200) {
                        this.$message.success(`Sửa đổi mật khẩu thành công`);
                        this.getData();
                        this.fromModifyPwdVisible = false;
                    } else {
                        this.$message.error('Xuống điểm không thành côn：' + res.msg);
                    }
                });
            },
            // Hiển thị cửa sổ lên điểm
            handleSaveMoney(index, row) {
                console.log(row);
                this.fromSaveMoney.UserId = row.Id;
                this.fromSaveMoney.UserName = row.UserName;
                this.fromSaveMoney.Money = 1000;
                this.fromSaveMoneyVisible = true;
            },
            // Hiẻn thị cửa sổ xuống điểm
            handleDrawSaveMoney(index, row) {
                this.fromDrawMoney.UserId = row.Id;
                this.fromDrawMoney.UserName = row.UserName;
                this.fromDrawMoney.Money = 1000;
                this.fromDrawMoneyVisible = true;
            },
            // Hiển thị các thao các chỉnh sửa
            handleEdit(index, row) {
                this.idx = index;
                console.log(row);
                this.form = row;
                this.editVisible = true;
            },

            handleModifyPwd(index, row) {
                this.idx = index;
                this.fromModifyPwd.UserId = row.Id;
                this.fromModifyPwd.Pwd = row.Pwd;
                this.fromModifyPwd.UserName = row.UserName;

                this.fromModifyPwdVisible = true;
            },
            handleAdd() {
                this.idx = -1;
                var aFilm = {};
                this.form = aFilm;
                this.editVisible = true;
            },
            // Lưu chỉnh sửa
            // saveEdit() {
            //     if (this.idx===-1) {
            //         var data = {
            //             jsonData:JSON.stringify(this.form)
            //         }
            //         request({
            //             url:'api/addgoodsinfo', method: 'post',
            //             data:data
            //         }).then((res)=>{
            //             console.log(res);
            //             if (res.code==200){
            //                 this.editVisible = false;
            //                 this.$message.success(`Thêm vào thành cống`);
            //                 this.getData();
            //             } else {
            //                 this.$message.error('Thêm vào không thành công：'+res.msg);
            //             }
            //         });
            //     } else {
            //         var data = {
            //             jsonData:JSON.stringify(this.form)
            //         }
            //         request({
            //             url:'api/updategoods', method: 'post',
            //             data:data
            //         }).then((res)=>{
            //             console.log(res);
            //             if (res.code==200){
            //                 this.editVisible = false;
            //                 this.$message.success(`Sửa đổi từ ${this.idx + 1} hàng thành công`);
            //                 this.$set(this.tableData, this.idx, this.form);
            //             } else {
            //                 this.$message.error('Cập Nhật không thành công：'+res.msg);
            //             }
            //         });
            //
            //     }
            // },
            // Điều hướng trang
            handlePageChange(val) {
                this.$set(this.query, 'pageIndex', val);
                this.getData();
            }
        }
    };
</script>

<style scoped>
    .handle-box {
        margin-bottom: 20px;
    }

    .handle-select {
        width: 120px;
    }

    .handle-input {
        width: 300px;
        display: inline-block;
    }

    .table {
        width: 100%;
        font-size: 14px;
    }

    .red {
        color: #ff0000;
    }

    .mr10 {
        margin-right: 10px;
    }

    .table-td-thumb {
        display: block;
        margin: auto;
        width: 40px;
        height: 40px;
    }

    .ivu-upload {
        display: none;
    }

    .ivu-btn {

    }
</style>
