<template>
    <div>
        <div class="container">
            <div class="handle-box">
              <el-input v-model="query.userName" placeholder="Tên người dùng" class="handle-input mr10"></el-input>
              UserId:
              <el-input v-model="query.userId" placeholder="Tên người dùng Id" class="handle-input-min mr10"></el-input>

              <el-date-picker class="handle-input-min"
                              v-model="query.beginDay"
                              type="date" value-format="yyyy-MM-dd"
                              placeholder="Chọn ngày bắt đầu">
              </el-date-picker>
              <el-date-picker class="handle-input-min"
                              v-model="query.endDay"
                              type="date" value-format="yyyy-MM-dd"
                              placeholder="Chọn ngày kết thúc">
              </el-date-picker>
              <el-select v-model.number="query.userType" placeholder="Xin vui lòng chọn">
                  <el-option v-for="item in UserTypeAll" :label="item.name" :key="item.value"
                             :value="item.value"></el-option>
              </el-select>
              <el-select v-model.number="query.gameType" placeholder="Xin vui lòng chọn">
                  <el-option v-for="item in GameTypeAll" :label="item.name" :key="item.value"
                             :value="item.value"></el-option>
              </el-select>
              <el-select v-model.number="query.state" placeholder="Xin vui lòng chọn">
                  <el-option v-for="item in BetStateAll" :label="item.name" :key="item.value"
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
                    show-summary>
                <!--                <el-table-column type="selection" width="55" align="center"></el-table-column>-->
                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>
              <el-table-column prop="UserId" label="UserId" width="100"></el-table-column>
                <el-table-column prop="UserName" label="Tên người dùng" width="120"></el-table-column>
                <el-table-column prop="GameName" label="Loại số"></el-table-column>
                <el-table-column prop="StrPeriod" label="Lượt số"></el-table-column>
                <el-table-column prop="BetM" label="Số tiền đặt cược"></el-table-column>
                <el-table-column prop="Win" label="Được thắng">
                    <template slot-scope="scope">{{scope.row.Win|money}}</template>
                </el-table-column>
                <el-table-column prop="CreatedAt" label="Thời gian đặt cược"></el-table-column>
                <el-table-column prop="BetSn" label="Số đơn">
                    <template slot-scope="scope">
                        <el-button
                                type="text"
                                @click="openDetail(scope.$index, scope.row)"
                        >{{scope.row.BetSn}}
                        </el-button>
                    </template>

                </el-table-column>
                <el-table-column prop="BetStr" label="Nội dung đặt cược"></el-table-column>
                <el-table-column prop="Status" label="Trạng thái">
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

        <!-- Chỉnh sửa cửa sổ bật lê -->
        <el-dialog title="Chi tiết" :visible.sync="detailVisible" width="90%">
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
                    <el-table-column prop="UserName" label="Tên người dùng" width="100"></el-table-column>
                  <el-table-column prop="OddsName" label="Nội dung"></el-table-column>
                  <el-table-column prop="OddsDes" label="Nội dung(中)"></el-table-column>
                    <el-table-column prop="BetM" label="Số tiền đặt cược"></el-table-column>
                    <el-table-column prop="Odds" label="Nhân"></el-table-column>
                    <el-table-column prop="Win" label="Thắng được">
                        <template slot-scope="scope">{{scope.row.Win|money}}</template>
                    </el-table-column>
                    <el-table-column prop="CreatedAt" label="Thời gian đặt cược"></el-table-column>
                    <el-table-column prop="GroupBetSn" label="Thắng được">

                    </el-table-column>
                    <el-table-column prop="Status" label="Trạng thái">
                        <template slot-scope="scope">{{scope.row.Status|BetStatusName}}</template>
                    </el-table-column>
                </el-table>
            </el-card>
            <span slot="footer" class="dialog-footer">
                <el-button @click="detailVisible = false">Đóng</el-button>
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
                  userId:0,
                    beginDay: '',
                    endDay: '',
                    userName: '',
                    pageIndex: 1,
                    pageSize: 10,
                    state:0,
                    userType:0,
                    gameType:0,
                },
                tableData: [],
                GroupBetId:0,
                tableDataDetail:[],
                detailVisible: false,
                pageTotal: 0,
                GameTypeAll: [
                  { name: 'Tất cả', value: 0 },
                  { name: 'Hồ chí minh', value: 201 },
                  { name: 'Miền Bắc', value: 202 },
                  { name: 'Keno', value: 203 },

                  { name: '加拿大28', value: 41 },
                  { name: '北京28', value: 42},
                  { name: '新加坡28', value: 43 },
                ],
              GroupData:{BetM:0,Win:0},
                UserTypeAll: [
                    { name: 'Toàn bộ khách', value: 0 },
                    { name: 'Khách thường', value: 1 },
                    { name: 'Khách chào hàng', value: 2 },
                ],

                BetStateAll: [
                    { name: 'Trạng thái đặt cược toàn b', value: 0 },
                    { name: 'Đợi mở thưởng j', value: 1 },
                    { name: 'Đã được đổi', value: 2 },
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
            // Lấy được easy-mock dữ liệu giả
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
            // Kích hoạt nút tìm kiếm
            handleSearch() {
                this.$set(this.query, 'pageIndex', 1);
                this.getData();
            },
            // Chỉnh sửa hoạt động
            handleEdit(index, row) {
                this.idx = index;
                this.form = row;
                this.editVisible = true;
            },
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
