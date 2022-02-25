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
                                placeholder="Chọn ngày kết thúc">
                </el-date-picker>
              <el-select v-model.number="query.accountType" placeholder="Vui lòng chọn loại">
                <el-option v-for="item in AccountTypeAll" :label="item.name" :key="item.value" :value="item.value"></el-option>
              </el-select>
                <el-button type="primary" icon="el-icon-search" @click="handleSearch">Tìm kiếm</el-button>
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
              <el-table-column prop="UserId" label="UserId" width="100"></el-table-column>
                <el-table-column prop="UserName" label="Tên người dùn" width="120"></el-table-column>
                <el-table-column prop="FullName" label="Họ tên"></el-table-column>
                <el-table-column prop="StrType" label="Thể loại">
                  <template slot-scope="scope">
                    <div v-if="scope.row.AccountType==1">Đoán</div>
                    <div v-else-if="scope.row.AccountType==2">Thắng được</div>
                    <div v-else-if="scope.row.AccountType==3">Nạp điểm</div>
                    <div v-else-if="scope.row.AccountType==4">Rút điểm</div>
                    <div v-else-if="scope.row.AccountType==7">Rút điểm bị từ chối</div>
                    <div v-else-if="scope.row.AccountType==8">Nạp điểm</div>
                    <div v-else-if="scope.row.AccountType==9">Rút điểm</div>
                    <div v-else-if="scope.row.AccountType==5">Khuyễn mãi</div>
                    <div v-else-if="scope.row.AccountType==12">Rửa mã</div>
                      <div v-else-if="scope.row.AccountType==13">Hoa Hồng chuyển đổi</div>
                    <div v-else>未知</div>
                  </template>
                </el-table-column>

                <el-table-column label="Số tiền">
                    <template slot-scope="scope">{{scope.row.Gold|money}}</template>
                </el-table-column>
                <el-table-column label="Tiền số dư">
                    <template slot-scope="scope">{{scope.row.CurUserGold|money}}</template>
                </el-table-column>
                <el-table-column prop="Des" label="miêu tả"></el-table-column>
                <el-table-column label="Thời gian">
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
                {name:'Các loại',value:0},
                {name:'Đoán',value:1},
                {name:'Thắng được',value:2},
                {name:'Nạp điểm',value:3},
                {name:'Ra điểm',value:4},
                {name:'Thắng được',value:5},
                // {name:'Phát phần thưởng',value:6},
                {name:'Rút tiền bị từ chối',value:7},
                {name:'Lên điểm',value:8},
                {name:'Xuống điểm',value:9},
                {name:'Hoa hông chuyển đổi',value:13},
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
            // Kích hoạt nút tìm kiếm
            handleSearch() {
                this.$set(this.query, 'pageIndex', 1);
                this.getData();
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
