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
              <el-select v-model.number="query.accountType" placeholder="请选择类型">
                <el-option v-for="item in AccountTypeAll" :label="item.name" :key="item.value" :value="item.value"></el-option>
              </el-select>
                <el-button type="primary" icon="el-icon-search" @click="handleSearch">搜索</el-button>
            </div>
            <el-table
                    :data="tableData"
                    border
                    class="table"
                    ref="multipleTable"
                    header-cell-class-name="table-header"
            >
                <!--                <el-table-column type="selection" width="55" align="center"></el-table-column>-->
                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>
              <el-table-column prop="UserId" label="用户ID"  width="100"></el-table-column>
              <el-table-column prop="UserName" label="用户名" width="120"></el-table-column>
                <el-table-column prop="FullName" label="姓名"></el-table-column>
                <el-table-column prop="StrType" label="类型"></el-table-column>
                <el-table-column label="金额">
                    <template slot-scope="scope">{{scope.row.Gold|money}}</template>
                </el-table-column>
                <el-table-column label="剩余金额">
                    <template slot-scope="scope">{{scope.row.CurUserGold|money}}</template>
                </el-table-column>
                <el-table-column prop="Des" label="描述"></el-table-column>
                <el-table-column label="时间">
                    <template slot-scope="scope">{{scope.row.CreatedAt}}</template>
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
    import VueCropper from 'vue-cropperjs';
    import { request } from '../../utils/http';

    export default {
        name: 'AccountList',
        data() {
            return {
                defaultSrc: require('../../assets/img/img.jpg'),
              AccountTypeAll: [
                {name:'全部类型',value:0},
                {name:'竞猜',value:1},
                {name:'赢得',value:2},
                {name:'充值',value:3},
                {name:'充值_上级减',value:103},
                {name:'提现',value:4},
                {name:'提现_上级加',value:104},
                // {name:'赢得',value:5},
                {name:'赠送',value:5},
                {name:'赠送_上级减',value:105},

                {name:'提现拒绝',value:7},
                {name:'提现拒绝_上级减',value:107},
                {name:'上分',value:8},
                {name:'上分_上级减',value:108},
                {name:'下分',value:9},
                {name:'下分_上级加',value:109},
                {name:'佣金转换',value:13},
                {name:'佣金转换_上级减',value:113},
              ],
                imgSrc: '',
                cropImg: '',
                dialogVisible: false,
                query: {
                  userId: 0,
                  accountType: 0,
                    beginDay: '',
                    endDay: '',
                    userName: '',
                    pageIndex: 1,
                    pageSize: 10
                },
                tableData: [],
                multipleSelection: [],
                delList: [],
                editVisible: false,
                pageTotal: 0,
                form: {},
                idx: -1,
                id: -1
            };
        },
        created() {
            this.getData();
        },
        components: {
            VueCropper
        },
        methods: {
            // 获取 easy-mock 的模拟数据
            getData() {
                request({ url: 'api/getaccountlist', method: 'post', data: this.query }).then((res) => {
                    if (res.code == 200) {
                        console.log(res.obj.ListData);
                        this.tableData = res.obj.ListData;
                        this.pageTotal = res.obj.PageTotal;
                    }
                });
            },
            // 触发搜索按钮
            handleSearch() {
                this.$set(this.query, 'pageIndex', 1);
                this.getData();
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
