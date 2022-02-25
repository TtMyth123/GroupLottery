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
              <el-table-column prop="UserId" label="UserId" width="100"></el-table-column>
                <el-table-column prop="UserName" label="Tên người dùng" width="120"></el-table-column>
                <el-table-column prop="FullName" label="Họ tên"></el-table-column>
                <el-table-column label="Thể loại người dùng">
                    <template slot-scope="scope">{{scope.row.UserType|UserTypeName}}</template>
                </el-table-column>
                <el-table-column label="Số tiền">
                    <template slot-scope="scope">{{scope.row.Money}}</template>
                </el-table-column>
                <el-table-column label="Số dư">
                    <template slot-scope="scope">{{scope.row.CurGold}}</template>
                </el-table-column>
                <el-table-column label="Thời gian">
                    <template slot-scope="scope">{{scope.row.CreatedAt}}</template>
                </el-table-column>
                <el-table-column label="Điều hành" width="140" align="center">
                    <template slot-scope="scope">
                        <!--                        <el-button-->
                        <!--                                type="text"-->
                        <!--                                icon="el-icon-edit"-->
                        <!--                                @click="handleEdit(scope.$index, scope.row)"-->
                        <!--                        >Lượt xem</el-button>-->
                        <el-button
                                type="text"
                                icon="el-icon-coordinate"
                                @click="handleAgree(scope.$index, scope.row)"
                        >Đồng ý
                        </el-button>
                        <el-button
                                type="text"
                                icon="el-icon-delete"
                                class="red"
                                @click="handleDelete(scope.$index, scope.row)"
                        >Từ chối
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

    </div>
</template>
<script>
    import VueCropper from 'vue-cropperjs';
    import { request } from '../../utils/http';

    export default {
        name: 'SaveMoneyApplyList',
        data() {
            return {
                defaultSrc: require('../../assets/img/img.jpg'),
                fileList: [],
                imgSrc: '',
                cropImg: '',
                dialogVisible: false,
                query: {
                    userId: 0,
                    beginDay: '',
                    endDay: '',
                    userName: '',
                    pageIndex: 1,
                    pageSize: 10
                },
                tableData: [],
                GroupData: {Gold:0,Money:0},
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
            getSummaries() {
                const sums = [];
                sums[1] = 'Thống kê：' + this.GroupData.C + 'Cái';
                sums[3] = this.GroupData.Gold.toFixed(2);
                sums[4] = this.GroupData.Money.toFixed(2);
                return sums;
            },

            setImage(e) {
                const file = e.target.files[0];
                if (!file.type.includes('image/')) {
                    return;
                }
                const reader = new FileReader();
                reader.onload = (event) => {
                    this.dialogVisible = true;
                    this.imgSrc = event.target.result;
                    this.$refs.cropper && this.$refs.cropper.replace(event.target.result);
                };
                reader.readAsDataURL(file);
            },
            cropImage() {
                this.cropImg = this.$refs.cropper.getCroppedCanvas().toDataURL();
            },
            cancelCrop() {
                this.dialogVisible = false;
                this.cropImg = this.defaultSrc;
            },

            // 获取 easy-mock 的模拟数据
            getData() {
                request({ url: 'api/getsavemoneyapplylist', method: 'post', data: this.query }).then((res) => {
                    if (res.code == 200) {
                        console.log(res.obj.ListData);
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
            // Xoá hoạt động
            handleDelete(index, row) {
                // Lần thứ 2 xác nhận xoá
                this.$confirm('Chắc chắn muốn từ chối không？', 'Lời nhắc', {
                    type: 'warning'
                })
                    .then(() => {
                        var data = {
                            Id: row.Id
                        };
                        request({
                            url: 'api/delsavemoney', method: 'post',
                            data: data
                        }).then((res) => {
                            console.log(res);
                            if (res.code == 200) {
                                this.$message.success('Xoá thành công');
                                this.tableData.splice(index, 1);
                                this.getData();
                            } else {
                                this.$message.error('Xoá không thành công：' + res.msg);
                            }
                        });
                    })
                    .catch(() => {
                    });
            },
            handleAgree(index, row) {
                var data = {
                    SaveMoneyId: row.Id,
                    UserId: row.UserId
                };
                request({
                    url: 'api/agreesavemoney', method: 'post',
                    data: data
                }).then((res) => {
                    console.log(res);
                    if (res.code == 200) {
                        this.$message.success('Gửi thành công');
                        this.tableData.splice(index, 1);
                        this.getData();
                    } else {
                        this.$message.error('Gửi không thành công：' + res.msg);
                    }
                });
            },
            // Chỉnh sửa hoạt động
            handleEdit(index, row) {
                this.idx = index;
                this.form = row;
                this.editVisible = true;
            },
            handleAdd() {
                this.idx = -1;
                var aFilm = {};
                this.form = aFilm;
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
