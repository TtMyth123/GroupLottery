<template>
    <div>
        <div class="container" style="height: 100%">
            <!--            <div class="handle-box">-->
            <!--                <el-input v-model="query.userName" placeholder="用户名" class="handle-input mr10"></el-input>-->
            <!--                <el-button type="primary" icon="el-icon-search" @click="handleSearch">搜索</el-button>-->
            <!--            </div>-->
            <el-table
                    :data="tableData"
                    border
                    class="table"
                    max-height="600"
                    header-cell-class-name="table-header"
                    :summary-method="getSummaries"
                    show-summary
            >
                <!--                <el-table-column type="selection" width="55" align="center"></el-table-column>-->
                <el-table-column prop="Id" label="用户ID" width="80" align="center"></el-table-column>
                <el-table-column prop="UserName" label="用户名" width="120"></el-table-column>
                <el-table-column prop="FullName" label="姓名"></el-table-column>
                <el-table-column label="余额">
                    <template slot-scope="scope">{{scope.row.Gold}}</template>
                </el-table-column>
                <el-table-column label="积分">
                    <template slot-scope="scope">{{scope.row.Point}}</template>
                </el-table-column>
                <el-table-column label="剩余佣金">
                    <template slot-scope="scope">{{scope.row.MemberFilmRemainBrokerage}}</template>
                </el-table-column>
                <el-table-column label="投资金额">
                    <template slot-scope="scope">{{scope.row.AllFilmGold|money}}</template>
                </el-table-column>
                <el-table-column label="提现额度">
                    <template slot-scope="scope">{{scope.row.DrawMoneyQuota|money}}</template>
                </el-table-column>
                <el-table-column label="分红总得">
                    <template slot-scope="scope">{{scope.row.AllBrokerage|money}}</template>
                </el-table-column>

                <el-table-column label="会员级别">
                    <template slot-scope="scope">{{scope.row.MemberLevel|MemberLevelName}}</template>
                </el-table-column>
                <el-table-column label="团队级别">
                    <template slot-scope="scope">{{scope.row.TeamLevel|TeamLevelName}}</template>
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
                <el-table-column label="状态">
                    <template slot-scope="scope">
                        {{scope.row.State|StateName}}
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
    </div>
</template>


<script>
    import { request } from '../../utils/http';
    import JuniorUserListComp from './JuniorUserList.vue';

    export default {
        name: 'JuniorUserList',
        props: {
            pid: {
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
                query: {
                    pid: 0,
                    state: 0,
                    userName: '',
                    pageIndex: 1,
                    pageSize: 10
                },
                tableData: [],
                GroupData: {},
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
        methods: {
            getSummaries() {
                const sums = [];
                // sums[1] = "合计："+ this.GroupData.C+"条";
                // sums[3] = this.GroupData.Gold.toFixed(2);
                // sums[4] = this.GroupData.Point;
                // sums[5] = this.GroupData.MemberFilmRemainBrokerage.toFixed(2);
                //
                // sums[6] = this.GroupData.AllFilmGold.toFixed(2);
                // sums[7] = this.GroupData.DrawMoneyQuota.toFixed(2);
                // sums[8] = this.GroupData.AllBrokerage.toFixed(2);

                return sums;
            },
            onSubmit() {
                // const id = this.$layer.iframe({
                //   content: {
                //     content: formComp,
                //     parent: this,
                //     data: { info: this.info }
                //   },
                //   area: ["500px", "300px"],
                //   title: "这是一个标题这是一个标题这是一个标题这是一个标题",
                //   maxmin: true,
                //   shade: true,
                //   shadeClose: false,
                //   scrollbar: false,
                //   resize: true,
                //   cancel: () => {
                //     alert(2110);
                //   }
                // });
                // this.$layer.full(id);

                this.$layer.msg('提交成功', () => {
                    this.lydata.info.name = this.form.name;
                    this.$layer.close(this.layerid);
                });

            },
            quxiao() {
                this.fn();
                this.$layer.close(this.layerid);
            },
            openJuniorUser(index, row) {
                const id = this.$layer.iframe({
                    content: {
                        content: JuniorUserListComp,
                        parent: this,
                        data: { pid: row.Id }
                    },
                    area: ['100%', '900px'],
                    title: '这是一个标题这是一个标题这是一个标题这是一个标题',
                    maxmin: true,
                    shade: true,
                    shadeClose: false,
                    scrollbar: true,
                    resize: true,
                    cancel: () => {
                        console.log(this.layerid);
                        //this.$layer.close(this.layerid);
                        this.$layer.close(id);
                    }
                }, layerid => {
                    // this.$layer.close(layerid);
                });
            },
            getData() {
                request({ url: 'api/getjuniorgameuserlist', method: 'post', data: this.query }).then((res) => {
                    if (res.code == 200) {
                        this.tableData = res.obj.ListData;
                        this.pageTotal = res.obj.PageTotal;
                        this.GroupData = res.obj.GroupData;
                    }
                });
            },
            handlePageChange(val) {
                this.$set(this.query, 'pageIndex', val);
                this.getData();
            }
        },
        mounted() {
            this.query.pid = this.pid;
            this.getData();
        }
    };
</script>

<style scoped>

</style>