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
                <el-button type="primary" icon="el-icon-search" @click="handleSearch">搜索</el-button>
              <el-switch
                  v-model="isAuto"
                  active-color="#13ce66"
                  inactive-color="#ff4949"
                  active-text="自动刷新"
                  inactive-text="">
              </el-switch>
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
              <el-table-column prop="FullName" label="姓名"></el-table-column>
                <el-table-column label="用户类别">
                    <template slot-scope="scope">{{scope.row.UserType|UserTypeName}}</template>
                </el-table-column>
                <el-table-column label="金额">
                    <template slot-scope="scope">{{scope.row.Gold}}</template>
                </el-table-column>
                <el-table-column label="余额">
                    <template slot-scope="scope">{{scope.row.CurGold}}</template>
                </el-table-column>
                <el-table-column label="时间">
                    <template slot-scope="scope">{{scope.row.CreatedAt}}</template>
                </el-table-column>
                <el-table-column label="操作" width="140" align="center">
                    <template slot-scope="scope">
                                                <el-button
                                                        type="text"
                                                        icon="el-icon-edit"
                                                        @click="handleEdit(scope.$index, scope.row)"
                                                >查看信息</el-button>
                        <el-button
                                type="text"
                                icon="el-icon-coordinate"
                                @click="handleAgree(scope.$index, scope.row)"
                        >同意
                        </el-button>
                        <el-button
                                type="text"
                                icon="el-icon-delete"
                                class="red"
                                @click="handleDelete(scope.$index, scope.row)"
                        >拒绝
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
      <!-- 编辑弹出框 -->
      <el-dialog title="用户信息" :visible.sync="editVisible" width="90%">
        <el-form ref="form" :model="form"  label-width="100px">
          <el-form-item label="微信收款码">
              <img v-if="form.WXSKCodeUrl" :src="form.WXSKCodeUrl" class="avatar">
          </el-form-item>
          <el-divider></el-divider>
          <el-form-item label="银行卡号">
            <el-input v-model="form.CardNum"></el-input>
          </el-form-item>
          <el-form-item label="开户行名">
            <el-input v-model="form.YHName"></el-input>
          </el-form-item>
          <el-form-item label="身份证">
            <el-input v-model="form.Cate"></el-input>
          </el-form-item>
          <el-form-item label="银行预留电话">
            <el-input v-model="form.YHUserTel"></el-input>
          </el-form-item>
          <el-form-item label="银行预留地址">
            <el-input v-model="form.Addr"></el-input>
          </el-form-item>
          <el-form-item label="银行预留信息">
            <el-input v-model="form.Remark"></el-input>
          </el-form-item>
          <el-divider></el-divider>
          <el-form-item label="支付宝名">
            <el-input v-model="form.ZFBSKName"></el-input>
          </el-form-item>
          <el-form-item label="支付宝二维码" >
            <img v-if="form.ZFBSKCodeUrl" :src="form.ZFBSKCodeUrl" class="avatar">
          </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
                <el-button @click="editVisible = false">取 消</el-button>
            <el-button @click="handleAgreeEx(form.DrawMoneyId)">同意申请</el-button>
            </span>
      </el-dialog>
    </div>
</template>
<script>
    import VueCropper from 'vue-cropperjs';
    import { request } from '../../utils/http';

    export default {
        name: 'DrawMoneyApplyList',
        data() {
            return {
              isAuto:true,
              dialogTitle:'查看',
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
              GroupData: {C:0,Gold:0,Money:0},
                tableData: [],
                multipleSelection: [],
                delList: [],
                editVisible: false,
                pageTotal: 0,
                form: {DrawMoneyId:0,WXSKCodeUrl:'',YHName:'',CardNum:'',YHUserName:'',Addr:'',Cate:'',Remark:'',ZFBSKCodeUrl:"",ZFBSKName:''},
                idx: -1,
                id: -1
            };
        },
        created() {
            this.getData();
        },
    destroyed() {
      clearInterval(this.timer);
    },
      mounted() {
        if (this.timer) {
          clearInterval(this.timer)
        } else {
          this.timer = setInterval(this.autoGetData,5000);
        }
      },
        methods: {
          autoGetData(){
            if ((this.query.pageIndex==1)&&(this.isAuto)) {
              this.getData();
            }
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

            getSummaries() {
                const sums = [];
                sums[1] = '合计：' + this.GroupData.C + '条';
                sums[4] = this.GroupData.Gold.toFixed(2);
                return sums;
            },
            // 获取 easy-mock 的模拟数据
            getData() {
                request({ url: 'api/getdrawmoneyapplylist', method: 'post', data: this.query }).then((res) => {
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
            // 删除操作
            handleDelete(index, row) {
                this.$prompt('拒绝理由', '确定要拒绝吗', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消'
                }).then(({ value }) => {
                    // this.$message({
                    //     type: 'success',
                    //     message: '拒绝理由: ' + value
                    // });
                    var data = {
                        Id: row.Id,
                        Excuse: value
                    };
                    request({
                        url: 'api/deldrawmoney', method: 'post',
                        data: data
                    }).then((res) => {
                        console.log(res);
                        if (res.code == 200) {
                            this.$message.success('删除成功');
                            this.tableData.splice(index, 1);
                            this.getData();
                        } else {
                            this.$message.error('删除失败：' + res.msg);
                        }
                    });
                }).catch(() => {

                });

            },
            handleAgree(index, row) {
                var data = {
                    id: row.Id
                };
                request({
                    url: 'api/agreedrawmoney', method: 'post',
                    data: data
                }).then((res) => {
                    console.log(res);
                    if (res.code == 200) {
                        this.$message.success('提交成功');
                        this.tableData.splice(index, 1);
                        this.getData();
                    } else {
                        this.$message.error('提交失败：' + res.msg);
                    }
                });
            },
          handleAgreeEx(Id) {
            var data = {
              id: Id
            };
            request({
              url: 'api/agreedrawmoney', method: 'post',
              data: data
            }).then((res) => {
              console.log(res);
              if (res.code == 200) {
                this.editVisible = false;
                this.$message.success('提交成功');
                this.getData();
              } else {
                this.$message.error('提交失败：' + res.msg);
              }
            });
          },
            // 编辑操作
            handleEdit(index, row) {
              request({
                url: 'api/getdrawmoneyaccountinfo', method: 'post',
                data: {DrawMoneyId:row.Id}
              }).then((res) => {
                if (res.code == 200) {
                  this.editVisible = true;
                  this.form = res.obj;
                } else {
                  this.$message.error('获取数据失败：' + res.msg);
                }
              });
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
