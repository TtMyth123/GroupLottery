<template>
    <div>
        <div class="container">
            <div class="handle-box">
              <el-input v-model="query.userName" placeholder="用户名" class="handle-input mr10"></el-input>
              用户ID：
              <el-input v-model.number="query.userId" placeholder="用户Id" class="handle-input-min mr12"></el-input>
              <el-select v-model.number="query.userType" placeholder="请选择">
                    <el-option v-for="item in UserTypeAll" :label="item.name" :key="item.value"
                               :value="item.value"></el-option>
                </el-select>
                <el-button type="primary" icon="el-icon-search" @click="handleSearch">搜索</el-button>
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
                <el-table-column prop="Id" label="用户ID" width="100" align="center"></el-table-column>
                <el-table-column prop="UserName" label="用户名" width="120"></el-table-column>
                <el-table-column label="余额">
                    <template slot-scope="scope">{{scope.row.Gold|money}}</template>
                </el-table-column>

                <el-table-column label="投注金额">
                    <template slot-scope="scope">{{scope.row.SumBet}}</template>
                </el-table-column>

                <el-table-column label="赢得金额">
                    <template slot-scope="scope">{{scope.row.SumWin|money}}</template>
                </el-table-column>

                <el-table-column label="累计充值">
                    <template slot-scope="scope">{{scope.row.SumSaveMoney|money}}</template>
                </el-table-column>
                <el-table-column label="累计提现">
                    <template slot-scope="scope">{{scope.row.SumDrawMoney|money}}</template>
                </el-table-column>


                <el-table-column label="用户类别">
                    <template slot-scope="scope">{{scope.row.UserType|UserTypeName}}</template>
                </el-table-column>

                <el-table-column prop="PUserName" label="上级用户名" width="100">
                    <template slot-scope="scope">
                        <el-button
                                type="text"
                                @click="openPUser(scope.$index, scope.row)"
                        >{{scope.row.PUserName}}
                        </el-button>
                    </template>
                </el-table-column>
                <el-table-column prop="LowerCount" label="下级用户数">
                    <template slot-scope="scope">
                        <el-button
                                type="text"
                                @click="openJuniorUser(scope.$index, scope.row)"
                        >{{scope.row.LowerCount}}
                        </el-button>
                    </template>
                </el-table-column>
                <el-table-column label="注册时间">
                    <template slot-scope="scope">{{scope.row.CreatedAt}}</template>
                </el-table-column>
                <el-table-column label="最后登录时间">
                    <template slot-scope="scope">{{scope.row.LoginTime}}</template>
                </el-table-column>
                <el-table-column prop="FullName" label="姓名"></el-table-column>
                <el-table-column label="状态">
                    <template slot-scope="scope">
                        <el-button
                                type="text"
                                @click="openState(scope.$index, scope.row)"
                        >{{scope.row.State|StateName}}
                        </el-button>
                    </template>
                </el-table-column>

                <el-table-column label="操作" width="140" align="center">
                    <template slot-scope="scope">
                        <el-button
                                type="text"
                                icon="el-icon-edit"
                                @click="openUser(scope.$index, scope.row)"
                        >修改
                        </el-button>

                        <el-button
                                type="text"
                                icon="el-icon-edit"
                                @click="handleSaveMoney(scope.$index, scope.row)"
                        >上分
                        </el-button>
                        <el-button
                                type="text"
                                icon="el-icon-edit"
                                @click="handleDrawSaveMoney(scope.$index, scope.row)"
                        >下分
                        </el-button>
                        <el-button
                                type="text"
                                icon="el-icon-edit"
                                @click="handleModifyPwd(scope.$index, scope.row)"
                        >密码
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

        <el-dialog title="上分" :visible.sync="fromSaveMoneyVisible" width="90%">
            <el-form ref="form" :model="fromSaveMoney" label-width="100px">
                <el-form-item label="用户名" width="100">
                    <el-input v-model="fromSaveMoney.UserName" readonly></el-input>
                </el-form-item>
                <el-form-item label="上分分数">
                    <!--                    <el-input v-model.number="fromSaveMoney.Money"></el-input>-->
                    <el-input-number v-model.number="fromSaveMoney.Money" :step="100" :min="1"></el-input-number>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="fromSaveMoneyVisible = false">取 消</el-button>
                <el-button type="primary" @click="doSaveMoney">确 定</el-button>
            </span>
        </el-dialog>

        <el-dialog title="下分" :visible.sync="fromDrawMoneyVisible" width="90%">
            <el-form ref="form" :model="fromDrawMoney" label-width="100px">
                <el-form-item label="用户名" width="100">
                    <el-input v-model="fromDrawMoney.UserName" readonly></el-input>
                </el-form-item>
                <el-form-item label="下分分数">
                    <!--                    <el-input v-model.number="fromSaveMoney.Money"></el-input>-->
                    <el-input-number v-model.number="fromDrawMoney.Money" :step="100" :min="1"></el-input-number>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="fromDrawMoneyVisible = false">取 消</el-button>
                <el-button type="primary" @click="doDrawMoney">确 定</el-button>
            </span>
        </el-dialog>

        <!-- 编辑弹出框 -->
        <el-dialog title="查看" :visible.sync="editVisible" width="90%">
            <el-form ref="form" :model="form" label-width="100px">
                <el-form-item label="用户名" width="100">
                    <el-input v-model="form.UserName" readonly></el-input>
                </el-form-item>

                <el-form-item label="昵称">
                    <el-input v-model="form.Nickname"></el-input>
                </el-form-item>

                <el-form-item label="状态" prop="State">
                    <el-select v-model.number="form.State" placeholder="请选择">
                        <el-option v-for="item in StateTypes" :label="item.name" :key="item.value"
                                   :value="item.value"></el-option>
                    </el-select>
                </el-form-item>

                <el-form-item label="会员级别" prop="MemberLevel">
                    <el-select v-model.number="form.MemberLevel" placeholder="请选择">
                        <el-option v-for="item in MemberLevelTypes" :label="item.name" :key="item.value"
                                   :value="item.value"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="团队级别" prop="TeamLevel">
                    <el-select v-model.number="form.TeamLevel" placeholder="请选择">
                        <el-option v-for="item in TeamLevelTypes" :label="item.name" :key="item.value"
                                   :value="item.value"></el-option>
                    </el-select>
                </el-form-item>

            </el-form>

            <span slot="footer" class="dialog-footer">
                <el-button @click="editVisible = false">取 消</el-button>
            </span>
        </el-dialog>

        <el-dialog title="修改密码" :visible.sync="fromModifyPwdVisible" width="90%">
            <el-form ref="form" :model="fromModifyPwd" label-width="100px">
                <el-form-item label="用户名" width="100">
                    <el-input v-model="fromModifyPwd.UserName" readonly></el-input>
                </el-form-item>

                <el-form-item label="密码">
                    <el-input v-model.number="fromModifyPwd.Pwd"></el-input>
                </el-form-item>
            </el-form>

            <span slot="footer" class="dialog-footer">
                <el-button @click="handleModifyPwd = false">取 消</el-button>
                <el-button @click="doModifyPwd">确 定</el-button>
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
                    { name: '启用', value: 1 },
                    { name: '禁用', value: 2 }
                ],
                MemberLevelTypes: [
                    { name: '新手会员', value: 1 },
                    { name: '白银会员', value: 2 },
                    { name: '黄金会员', value: 3 },
                    { name: '钻石会员', value: 4 },
                    { name: '皇冠会员', value: 5 }
                ],
                TeamLevelTypes: [
                    { name: '', value: 0 },
                    { name: '一级团长', value: 1 },
                    { name: '二级团长', value: 2 },
                    { name: '三级团长', value: 3 },
                    { name: '四级团长', value: 4 },
                    { name: '五级团长', value: 5 }
                ],
                UserTypeAll: [
                    { name: '全部', value: 0 },
                    { name: '一般玩家', value: 1 },
                    { name: '业务玩家', value: 2 }
                ],
                query: {
                    userType: 0,
                    state: 0,
                    userId:0,
                    userName: '',
                    pageIndex: 1,
                    pageSize: 10
                },
                tableData: [],
                GroupData: {Gold:0,SumBet:0,SumWin:0,SumSaveMoney:0,SumDrawMoney:0},
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
                sums[6] = this.GroupData.SumDrawMoney;

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
                    title: '修改用户信息',
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
                    title: '修改用户信息',
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
                let hint = '确定要启动' + row.UserName + '用户名？';
                let State = 1;
                if (row.State == 1) {
                    hint = '确定要禁用' + row.UserName + '用户名？';
                    State = 2;
                }

                this.$confirm(hint, '提示', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
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
                                message: '设置成功!'
                            });
                            row.State = State;
                        }
                    });
                }).catch(() => {
                    // this.$message({
                    //     type: 'info',
                    //     message: '已取消删除'
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
            // 触发搜索按钮
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
                        this.$message.success(`上分成功`);
                        this.getData();
                        this.fromSaveMoneyVisible = false;
                    } else {
                        this.$message.error('上分失败：' + res.msg);
                    }
                });
            },
            doDrawMoney() {
                request({
                    url: 'api/drawmoney', method: 'post',
                    data: this.fromDrawMoney
                }).then((res) => {
                    if (res.code == 200) {
                        this.$message.success(`下分成功`);
                        this.getData();
                        this.fromDrawMoneyVisible = false;
                    } else {
                        this.$message.error('下分失败：' + res.msg);
                    }
                });
            },
            doModifyPwd() {
                request({
                    url: 'api/updateuserpwd', method: 'post',
                    data: this.fromModifyPwd
                }).then((res) => {
                    if (res.code == 200) {
                        this.$message.success(`修改密码成功`);
                        this.getData();
                        this.fromModifyPwdVisible = false;
                    } else {
                        this.$message.error('下分失败：' + res.msg);
                    }
                });
            },
            // 显示上分窗口
            handleSaveMoney(index, row) {
                console.log(row);
                this.fromSaveMoney.UserId = row.Id;
                this.fromSaveMoney.UserName = row.UserName;
                this.fromSaveMoney.Money = 1000;
                this.fromSaveMoneyVisible = true;
            },
            // 显示下分窗口
            handleDrawSaveMoney(index, row) {
                this.fromDrawMoney.UserId = row.Id;
                this.fromDrawMoney.UserName = row.UserName;
                this.fromDrawMoney.Money = 1000;
                this.fromDrawMoneyVisible = true;
            },
            // 显示编辑操作
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
