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
              <el-select v-model.number="query.RebateType" placeholder="请选择类型">
                <el-option v-for="item in RebateTypeAll" :label="item.name" :key="item.value" :value="item.value"></el-option>
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
                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>
              <el-table-column prop="UserId" label="用户ID" width="100"></el-table-column>
              <el-table-column prop="UserName" label="用户名" width="120"></el-table-column>
              <el-table-column prop="StrType" label="类型"></el-table-column>
                <el-table-column label="佣金">
                    <template slot-scope="scope">{{scope.row.Rebate|money}}</template>
                </el-table-column>
                <el-table-column label="剩余佣金">
                    <template slot-scope="scope">{{scope.row.CurUserRebate|money}}</template>
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
        name: 'RebateList',
        data() {
            return {
                defaultSrc: require('../../assets/img/img.jpg'),
              RebateTypeAll: [
                {name:'全部类型',value:0},
                {name:'竞猜',value:1},
                {name:'转换金额',value:2},
              ],
                fileList: [],
                imgSrc: '',
                cropImg: '',
                dialogVisible: false,
                query: {
                    RebateType:0,
                    beginDay: '',
                    endDay: '',
                    userId:0,
                    userName: '',
                    pageIndex: 1,
                    pageSize: 10
                },
                multipleSelection: [],
                delList: [],
                editVisible: false,
                pageTotal: 0,
              tableData: [],
              GroupData: {},
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
            sums[3] = this.GroupData.Rebate;
            return sums;
          },
            // 获取 easy-mock 的模拟数据
            getData() {
                request({ url: 'api/getrebatelist', method: 'post', data: this.query }).then((res) => {
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
