<template>
    <div>
        <div class="container">
            <div class="handle-box">
              <el-input v-model="query.userName" placeholder="用户名" class="handle-input mr10"></el-input>
              用户ID：
              <el-input v-model.number="query.userId" placeholder="用户Id" class="handle-input-min mr12"></el-input>

              <el-date-picker class="handle-input-min"
                                v-model="query.beginDay"
                                type="date" value-format="yyyy-MM-dd"
                                placeholder="选择开始日期">
                </el-date-picker>
                <el-date-picker class="handle-input-min"
                                v-model="query.endDay"
                                type="date" value-format="yyyy-MM-dd"
                                placeholder="选择结束日期">
                </el-date-picker>
                <el-select v-model.number="query.userType" placeholder="请选择">
                    <el-option v-for="item in UserTypeAll" :label="item.name" :key="item.value"
                               :value="item.value"></el-option>
                </el-select>
                <el-select v-model.number="query.gameType" placeholder="请选择">
                    <el-option v-for="item in GameTypeAll" :label="item.name" :key="item.value"
                               :value="item.value"></el-option>
                </el-select>
                <el-select v-model.number="query.state" placeholder="请选择">
                    <el-option v-for="item in BetStateAll" :label="item.name" :key="item.value"
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
                    show-summary>
                <!--                <el-table-column type="selection" width="55" align="center"></el-table-column>-->
                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>
              <el-table-column prop="UserId" label="用户ID" width="100"></el-table-column>
              <el-table-column prop="UserName" label="用户名" width="120"></el-table-column>
              <el-table-column prop="GameName" label="彩种"></el-table-column>
                <el-table-column prop="StrPeriod" label="期号"></el-table-column>
                <el-table-column prop="BetM" label="投注金额"></el-table-column>
                <el-table-column prop="Win" label="赢得">
                    <template slot-scope="scope">{{scope.row.Win|money}}</template>
                </el-table-column>
                <el-table-column prop="CreatedAt" label="投注时间"></el-table-column>
                <el-table-column prop="BetSn" label="单号">
                    <template slot-scope="scope">
                        <el-button
                                type="text"
                                @click="openDetail(scope.$index, scope.row)"
                        >{{scope.row.BetSn}}
                        </el-button>
                    </template>

                </el-table-column>
                <el-table-column prop="BetStr" label="投注内容"></el-table-column>
                <el-table-column prop="Status" label="状态">
                    <template slot-scope="scope">{{scope.row.Status|BetStatusName}}</template>
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

        <!-- 编辑弹出框 -->
        <el-dialog title="明细" :visible.sync="detailVisible" width="90%">
            <el-card class="box-card">

            </el-card>
            <el-card class="box-card">
                <el-table
                        :data="tableDataDetail"
                        border
                        class="table"
                        ref="multipleTable"
                        header-cell-class-name="table-header">
                    <!--                <el-table-column type="selection" width="55" align="center"></el-table-column>-->
                    <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>
                    <el-table-column prop="UserName" label="用户名" width="100"></el-table-column>
                  <el-table-column prop="OddsName" label="内容"></el-table-column>
                  <el-table-column prop="OddsDes" label="内容(中)"></el-table-column>
                    <el-table-column prop="BetM" label="投注金额"></el-table-column>
                    <el-table-column prop="Odds" label="赔率"></el-table-column>
                    <el-table-column prop="Win" label="赢得">
                        <template slot-scope="scope">{{scope.row.Win|money}}</template>
                    </el-table-column>
                    <el-table-column prop="CreatedAt" label="投注时间"></el-table-column>
                    <el-table-column prop="GroupBetSn" label="单号">

                    </el-table-column>
                    <el-table-column prop="Status" label="状态">
                      <template slot-scope="scope">
                        <div v-if="scope.row.Status===1">
                          待开奖
                        </div>
                        <div v-else-if="scope.row.Status===2">
                          已兑奖
                        </div>
                        <div v-else>
                          不开奖退款
                        </div>
                      </template>
                    </el-table-column>
                </el-table>
            </el-card>
            <span slot="footer" class="dialog-footer">
                <el-button @click="detailVisible = false">关 闭</el-button>
            </span>
        </el-dialog>
    </div>
</template>
<script>
    import VueCropper from 'vue-cropperjs';
    import { request } from '../../utils/http';

    export default {
        name: 'GroupBetList',
        data() {
            return {
                defaultSrc: require('../../assets/img/img.jpg'),
                query: {
                  beginDay: '',
                  endDay: '',
                  userName: '',
                  pageIndex: 1,
                  pageSize: 10,
                  state:0,
                  userType:0,
                  userId:0,
                  gameType:0,
                },
                tableData: [],
              GroupData:{BetM:0,Win:0},
                GroupBetId:0,
                tableDataDetail:[],
                detailVisible: false,
                pageTotal: 0,
                GameTypeAll: [
                  { name: '全部彩种', value: 0 },
                  { name: '南部彩', value: 201 },
                  { name: '北部彩', value: 202 },
                  { name: '中部彩', value: 203 },

                  { name: '加拿大28', value: 41 },
                  { name: '北京28', value: 42},
                  { name: '新加坡28', value: 43 },
                ],
                UserTypeAll: [
                    { name: '全部玩家', value: 0 },
                    { name: '一般玩家', value: 1 },
                    { name: '业务玩家', value: 2 },
                ],

                BetStateAll: [
                    { name: '全部投注状态', value: 0 },
                    { name: '待开奖', value: 1 },
                    { name: '已兑奖', value: 2 },
                ],

            };
        },
        created() {
            this.getData();
        },
        components: {
            VueCropper
        },
        methods: {
            getSummaries() {
                const sums = [];
                sums[4] = this.GroupData.BetM;
                sums[5] = this.GroupData.Win.toFixed(2);

                return sums;
            },
            openDetail(index, row){
                this.GroupBetId = row.Id;
                this.getDetailData()
            },
            // 获取 easy-mock 的模拟数据
            getData() {
                request({ url: 'api/getgroupbetlist', method: 'post', data: this.query }).then((res) => {
                    if (res.code == 200) {
                        console.log(res.obj.ListData);
                        this.tableData = res.obj.ListData;
                        this.pageTotal = res.obj.PageTotal;
                    }
                });
            },
            getDetailData() {
                let data = {
                    GroupBetId:this.GroupBetId
                }
                request({ url: 'api/getgroupdetailbetlist', method: 'post', data: data }).then((res) => {
                    if (res.code == 200) {
                        this.detailVisible = true
                        this.tableDataDetail = res.obj
                    }
                });
            },
            // 触发搜索按钮
            handleSearch() {
                this.$set(this.query, 'pageIndex', 1);
                this.getData();
            },
            // 编辑操作
            handleEdit(index, row) {
                this.idx = index;
                this.form = row;
                this.editVisible = true;
            },
            // 分页导航
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
        width: 200px;
        display: inline-block;
    }

    .handle-input-min {
        width: 140px;
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
</style>
