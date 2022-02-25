<template>
    <div>
        <div class="container">
            <el-tabs v-model="activeName" >
              <el-tab-pane label="To Nhỏ Lẻ Nhẵn" name="tab0">
                <el-card class="box-card">
                  <div>
                    <el-breadcrumb-item>
                      Nhân Key Phạm vi（
                      <el-input-number class="loose-input" v-model="allData[0].BatchForm.MinValue" :step="1" :min="allData[0].BatchForm.Min" :max="allData[0].BatchForm.Max"> </el-input-number>
                      ~
                      <el-input-number class="loose-input" v-model="allData[0].BatchForm.MaxValue" :step="1" :min="allData[0].BatchForm.Min" :max="allData[0].BatchForm.Max"> </el-input-number>）
                      Nhân：
                      <el-input-number class="loose-input" v-model="allData[0].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                      <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(0)" >Sửa đổi hàng loạt</el-button>
                    </el-breadcrumb-item>
                  </div>
                </el-card>
                <el-card class="box-card">
                  <el-button type="primary" class="handle-del mr10" @click="getData(0)" >Làm mới</el-button>
                  <el-table
                      :data="allData[0].tableData"
                      border
                      class="table"
                      ref="multipleTable"
                      header-cell-class-name="table-header"
                  >
                    <!--                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>-->
                    <el-table-column prop="OddsType" label="Nhân Key" width="90" align="center"></el-table-column>
                    <el-table-column prop="OddsDes" label="Tên nhân"></el-table-column>
                    <el-table-column prop="Odds" label="Nhân">
                      <template slot-scope="scope">{{scope.row.Odds|money}}</template>
                    </el-table-column>
                    <el-table-column prop="OneUserMinBet" label="Tối thiệu 1 người"></el-table-column>
                    <el-table-column prop="OneUserMaxBet" label="Tối đa 1 người"></el-table-column>
                    <el-table-column prop="AllUserMaxBet" label="Tối đa toàn bộ"></el-table-column>

                    <el-table-column label="Điều hành" width="180" align="center">
                      <template slot-scope="scope">
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleEdit(0, scope.$index, scope.row)"
                        >Chỉnh sửa
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-card>
              </el-tab-pane>

              <el-tab-pane label="Đôi bão" name="tab1">
                <el-card class="box-card">
                  <div>
                    <el-breadcrumb-item>
                      Nhân Key Phạm vi（
                      <el-input-number class="loose-input" v-model="allData[1].BatchForm.MinValue" :step="1" :min="allData[1].BatchForm.Min" :max="allData[1].BatchForm.Max"> </el-input-number>
                      ~
                      <el-input-number class="loose-input" v-model="allData[1].BatchForm.MaxValue" :step="1" :min="allData[1].BatchForm.Min" :max="allData[1].BatchForm.Max"> </el-input-number>）
                      Nhân：
                      <el-input-number class="loose-input" v-model="allData[1].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                      <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(1)" >Sửa đổi hàng loạt</el-button>
                    </el-breadcrumb-item>
                  </div>
                </el-card>
                <el-card class="box-card">
                  <el-button type="primary" class="handle-del mr10" @click="getData(1)" >Làm mới</el-button>
                  <el-table
                      :data="allData[1].tableData"
                      border
                      class="table"
                      ref="multipleTable"
                      header-cell-class-name="table-header"
                  >
                    <!--                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>-->
                    <el-table-column prop="OddsType" label="Nhân Key" width="90" align="center"></el-table-column>
                    <el-table-column prop="OddsDes" label="Tên nhân"></el-table-column>
                    <el-table-column prop="Odds" label="Nhân">
                      <template slot-scope="scope">{{scope.row.Odds|money}}</template>
                    </el-table-column>
                    <el-table-column prop="OneUserMinBet" label="Tối thiệu 1 người"></el-table-column>
                    <el-table-column prop="OneUserMaxBet" label="Tối đa 1 người"></el-table-column>
                    <el-table-column prop="AllUserMaxBet" label="Tối đa toàn bộ"></el-table-column>

                    <el-table-column label="Điều hành" width="180" align="center">
                      <template slot-scope="scope">
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleEdit(1,scope.$index, scope.row)"
                        >Chỉnh sửa
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-card>
              </el-tab-pane>

              <el-tab-pane label="Số màu" name="tab2">
                <el-card class="box-card">
                  <div>
                    <el-breadcrumb-item>
                      Nhân Key Phạm vi（
                      <el-input-number class="loose-input" v-model="allData[2].BatchForm.MinValue" :step="1" :min="allData[2].BatchForm.Min" :max="allData[2].BatchForm.Max"> </el-input-number>
                      ~
                      <el-input-number class="loose-input" v-model="allData[2].BatchForm.MaxValue" :step="1" :min="allData[2].BatchForm.Min" :max="allData[2].BatchForm.Max"> </el-input-number>）
                      Nhân：
                      <el-input-number class="loose-input" v-model="allData[2].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                      <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(2)" >Sửa đổi hàng loạt</el-button>
                    </el-breadcrumb-item>
                  </div>
                </el-card>
                <el-card class="box-card">
                  <el-button type="primary" class="handle-del mr10" @click="getData(2)" >Làm mới</el-button>
                  <el-table
                      :data="allData[2].tableData"
                      border
                      class="table"
                      ref="multipleTable"
                      header-cell-class-name="table-header"
                  >
                    <!--                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>-->
                    <el-table-column prop="OddsType" label="Nhân Key" width="90" align="center"></el-table-column>
                    <el-table-column prop="OddsDes" label="Tên nhân"></el-table-column>
                    <el-table-column prop="Odds" label="Nhân">
                      <template slot-scope="scope">{{scope.row.Odds|money}}</template>
                    </el-table-column>
                    <el-table-column prop="OneUserMinBet" label="Tối thiệu 1 người"></el-table-column>
                    <el-table-column prop="OneUserMaxBet" label="Tối đa 1 người"></el-table-column>
                    <el-table-column prop="AllUserMaxBet" label="Tối đa toàn bộ"></el-table-column>

                    <el-table-column label="Điều hành" width="180" align="center">
                      <template slot-scope="scope">
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleEdit(2,scope.$index, scope.row)"
                        >Chỉnh sửa
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-card>
              </el-tab-pane>

              <el-tab-pane label="Long hổ hoà" name="tab3">
                <el-card class="box-card">
                  <div>
                    <el-breadcrumb-item>
                      Nhân Key Phạm vi（
                      <el-input-number class="loose-input" v-model="allData[3].BatchForm.MinValue" :step="1" :min="allData[3].BatchForm.Min" :max="allData[3].BatchForm.Max"> </el-input-number>
                      ~
                      <el-input-number class="loose-input" v-model="allData[3].BatchForm.MaxValue" :step="1" :min="allData[3].BatchForm.Min" :max="allData[3].BatchForm.Max"> </el-input-number>）
                      Nhân：
                      <el-input-number class="loose-input" v-model="allData[3].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                      <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(3)" >Sửa đổi hàng loạt</el-button>
                    </el-breadcrumb-item>
                  </div>
                </el-card>
                <el-card class="box-card">
                  <el-button type="primary" class="handle-del mr10" @click="getData(3)" >Làm mới</el-button>
                  <el-table
                      :data="allData[3].tableData"
                      border
                      class="table"
                      ref="multipleTable"
                      header-cell-class-name="table-header"
                  >
                    <!--                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>-->
                    <el-table-column prop="OddsType" label="Nhân Key" width="90" align="center"></el-table-column>
                    <el-table-column prop="OddsDes" label="Tên nhân"></el-table-column>
                    <el-table-column prop="Odds" label="Nhân">
                      <template slot-scope="scope">{{scope.row.Odds|money}}</template>
                    </el-table-column>

                    <el-table-column prop="OneUserMinBet" label="Tối thiệu 1 người"></el-table-column>
                    <el-table-column prop="OneUserMaxBet" label="Tối đa 1 người"></el-table-column>
                    <el-table-column prop="AllUserMaxBet" label="Tối đa toàn bộ"></el-table-column>

                    <el-table-column label="Điều hành" width="180" align="center">
                      <template slot-scope="scope">
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleEdit(3,scope.$index, scope.row)"
                        >Chỉnh sửa
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-card>
              </el-tab-pane>

              <el-tab-pane label="Xanh đỏ" name="tab4">
                <el-card class="box-card">
                  <div>
                    <el-breadcrumb-item>
                      Nhân Key Phạm vi（
                      <el-input-number class="loose-input" v-model="allData[4].BatchForm.MinValue" :step="1" :min="allData[4].BatchForm.Min" :max="allData[4].BatchForm.Max"> </el-input-number>
                      ~
                      <el-input-number class="loose-input" v-model="allData[4].BatchForm.MaxValue" :step="1" :min="allData[4].BatchForm.Min" :max="allData[4].BatchForm.Max"> </el-input-number>）
                      Nhân：
                      <el-input-number class="loose-input" v-model="allData[4].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                      <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(4)" >Sửa đổi hàng loạt</el-button>
                    </el-breadcrumb-item>
                  </div>
                </el-card>
                <el-card class="box-card">
                  <el-button type="primary" class="handle-del mr10" @click="getData(4)" >Làm mới</el-button>
                  <el-table
                      :data="allData[4].tableData"
                      border
                      class="table"
                      ref="multipleTable"
                      header-cell-class-name="table-header"
                  >
                    <!--                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>-->
                    <el-table-column prop="OddsType" label="Nhân Key" width="90" align="center"></el-table-column>
                    <el-table-column prop="OddsDes" label="Tên nhân"></el-table-column>
                    <el-table-column prop="Odds" label="Nhân">
                      <template slot-scope="scope">{{scope.row.Odds|money}}</template>
                    </el-table-column>

                    <el-table-column prop="OneUserMinBet" label="Tối thiệu 1 người"></el-table-column>
                    <el-table-column prop="OneUserMaxBet" label="Tối đa 1 người"></el-table-column>
                    <el-table-column prop="AllUserMaxBet" label="Tối đa toàn bộ"></el-table-column>

                    <el-table-column label="Điều hành" width="180" align="center">
                      <template slot-scope="scope">
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleEdit(4,scope.$index, scope.row)"
                        >Chỉnh sửa
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-card>
              </el-tab-pane>

              <el-tab-pane label="Số trị" name="tab5">
                <el-card class="box-card">
                  <div>
                    <el-breadcrumb-item>
                      Nhân Key Phạm vi（
                      <el-input-number class="loose-input" v-model="allData[5].BatchForm.MinValue" :step="1" :min="allData[5].BatchForm.Min" :max="allData[5].BatchForm.Max"> </el-input-number>
                      ~
                      <el-input-number class="loose-input" v-model="allData[5].BatchForm.MaxValue" :step="1" :min="allData[5].BatchForm.Min" :max="allData[5].BatchForm.Max"> </el-input-number>）
                      Nhân：
                      <el-input-number class="loose-input" v-model="allData[5].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                      <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(5)" >Sửa đổi hàng loạt</el-button>
                    </el-breadcrumb-item>
                  </div>
                </el-card>
                <el-card class="box-card">
                  <el-button type="primary" class="handle-del mr10" @click="getData(5)" >làm mới</el-button>
                  <el-table
                      :data="allData[5].tableData"
                      border
                      class="table"
                      ref="multipleTable"
                      header-cell-class-name="table-header"
                  >
                    <!--                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>-->
                    <el-table-column prop="OddsType" label="Nhân Key" width="90" align="center"></el-table-column>
                    <el-table-column prop="OddsDes" label="Tên nhâ"></el-table-column>
                    <el-table-column prop="Odds" label="Nhân">
                      <template slot-scope="scope">{{scope.row.Odds|money}}</template>
                    </el-table-column>

                    <el-table-column prop="OneUserMinBet" label="Tối thiệu 1 người"></el-table-column>
                    <el-table-column prop="OneUserMaxBet" label="Tối đa 1 người"></el-table-column>
                    <el-table-column prop="AllUserMaxBet" label="Tối đa toàn bộ"></el-table-column>

                    <el-table-column label="Điều hành" width="180" align="center">
                      <template slot-scope="scope">
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleEdit(5,scope.$index, scope.row)"
                        >Chỉnh sửa
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-card>
              </el-tab-pane>
            </el-tabs>

            <!-- Chỉnh sửa cửa sổ bặt lên -->
            <el-dialog title="Chỉnh sửa số nhân" :visible.sync="editVisible" width="90%">
                <el-form ref="form" :model="form"  label-width="100px">
                    <el-form-item label="Nhân Key" prop="OddsType">
                        <el-input v-model="form.OddsType" disabled></el-input>
                    </el-form-item>
                    <el-form-item label="Tên nhân" prop="OddsDes">
                        <el-input v-model="form.OddsDes" disabled></el-input>
                    </el-form-item>
                    <el-form-item label="Nhân" prop="Odds">
                        <el-input-number class="loose-input" v-model="form.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                    </el-form-item>
                  <el-form-item label="Tối đa 1 người" prop="OneUserMaxBet">
                    <el-input-number class="loose-input" v-model="form.OneUserMaxBet" :step="1000" :min="0"></el-input-number>
                  </el-form-item>
                  <el-form-item label="Tối thiệu 1 người" prop="OneUserMinBet">
                    <el-input-number class="loose-input" v-model="form.OneUserMinBet" :step="10" :min="1"></el-input-number>
                  </el-form-item>
                  <el-form-item label="Tối đa toàn bộ" prop="AllUserMaxBet">
                    <el-input-number class="loose-input" v-model="form.AllUserMaxBet" :step="1000" :min="0"></el-input-number>
                  </el-form-item>
                </el-form>
                <span slot="footer" class="dialog-footer">
                <el-button @click="editVisible = false">Huỷ</el-button>
                <el-button type="primary" @click="saveOdds">Xác nhận</el-button>
            </span>
            </el-dialog>
        </div>
    </div>

</template>

<script>
    import { request } from '../../utils/http';

    export default {
        name: 'Zg28BjSet',
        data() {
            return {
                AwardTypeItem: [{ name: 'Xu vàng', value: 1 }, { name: 'Tích điểm', value: 2 }],
                query: {
                    address: '',
                    name: '',
                    pageIndex: 1,
                    pageSize: 10
                },
                allData:[
                  {I:0, BigTypes: [1], tableData:[] ,BatchForm:{MinValue:11,MaxValue:20,Odds:1.9,Min:11,Max:20}, },
                  {I:1, BigTypes: [2], tableData:[] ,BatchForm:{MinValue:21,MaxValue:23,Odds:9.9,Min:21,Max:23}, },
                  {I:2, BigTypes: [3], tableData:[] ,BatchForm:{MinValue:24,MaxValue:26,Odds:9.9,Min:24,Max:26}, },
                  {I:3, BigTypes: [4], tableData:[] ,BatchForm:{MinValue:27,MaxValue:29,Odds:9.9,Min:27,Max:29}, },
                  {I:4, BigTypes: [5], tableData:[] ,BatchForm:{MinValue:31,MaxValue:34,Odds:9.9,Min:31,Max:34}, },
                  {I:5, BigTypes: [6], tableData:[] ,BatchForm:{MinValue:100,MaxValue:127,Odds:9.9,Min:100,Max:127}, },
                ],
                activeName:"tab0",
                multipleSelection: [],
                delList: [],
                editVisible: false,
                pageTotal: 0,
                form: {Id:0,OddsType:0,Odds:1.9, OddsDes:''},
                curFormT:0,
                curFormIdx: -1,
                GameType:42,
            };
        },
        created() {
          this.getData(0);
          this.getData(1);
          this.getData(2);
          this.getData(3);
          this.getData(4);
          this.getData(5);
          
        },
        methods: {
            batchModifyOdds(t){
                let batchData = {
                    GameType: this.GameType,
                    Odds: this.allData[t].BatchForm.Odds,
                    MinOddType:this.allData[t].BatchForm.MinValue,
                    MaxOddType:this.allData[t].BatchForm.MaxValue,
                };

                request({
                    url: 'api/savebatchodds', method: 'post',
                    data: batchData
                }).then((res) => {
                    console.log(res);
                    if (res.code == 200) {
                        this.editVisible = false;
                        this.$message.success(`Sửa đổ ${this.form.OddsDes} Thành công`);
                        this.getData(t, this.allData[t].BigTypes)
                    } else {
                        this.$message.error('Cập nhật không thành công：' + res.msg);
                    }
                });
            },
            getData(t) {
                let BigTypes = this.allData[t].BigTypes;
                let data = {
                    GameType:this.GameType,
                    ArrBigType:BigTypes,
                };
                let data1 = {
                    jsonData: JSON.stringify(data)
                };
                console.log('OK');
                request({ url: 'api/getoddslist', method: 'post' ,data:data1}).then((res) => {
                    console.log(res);
                    if (res.code == 200) {
                        this.allData[t].tableData = res.obj;
                    }
                });
            },

            // Lưu chỉnh sửa
            saveOdds() {
                request({
                    url: 'api/saveoddsinfo', method: 'post',
                    data: this.form
                }).then((res) => {
                    console.log(res);
                    if (res.code == 200) {
                        this.editVisible = false;
                        this.$message.success(`Sửa đổi ${this.form.OddsDes} Thành công`);
                        this.getData(this.curFormT, this.allData[this.curFormT].BigTypes)
                    } else {
                        this.$message.error('Cập nhật không thành công：' + res.msg);
                    }
                });
            },
            // Chỉnh sửa hoạt động
            handleEdit(t, index, row) {
                this.curFormIdx = index;
                this.curFormT = t
                this.form.Id = row.Id;
                this.form.OddsType = row.OddsType;
                this.form.OddsDes = row.OddsDes;
                this.form.Odds = row.Odds;
              this.form.OneUserMaxBet = row.OneUserMaxBet;
              this.form.OneUserMinBet = row.OneUserMinBet;
              this.form.AllUserMaxBet = row.AllUserMaxBet;

                this.editVisible = true;
            },
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
</style>
