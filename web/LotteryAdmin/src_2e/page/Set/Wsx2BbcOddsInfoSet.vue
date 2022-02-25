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
                        <el-button type="primary" class="handle-del mr10" @click="getData(0 )" >Làm mới</el-button>
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
              <el-tab-pane label="Giải đặc biệt-Số đầu đặc biệt" name="tab1">
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

              <el-tab-pane label="Giải đặc biệt-Số đuôi đặc biệt" name="tab2">
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

              <el-tab-pane label="Giải nhất-Số đầu đặc biệt" name="tab11">
                <el-card class="box-card">
                  <div>
                    <el-breadcrumb-item>
                      Nhân Key Phạm vi（
                      <el-input-number class="loose-input" v-model="allData[11].BatchForm.MinValue" :step="1" :min="allData[11].BatchForm.Min" :max="allData[11].BatchForm.Max"> </el-input-number>
                      ~
                      <el-input-number class="loose-input" v-model="allData[11].BatchForm.MaxValue" :step="1" :min="allData[11].BatchForm.Min" :max="allData[11].BatchForm.Max"> </el-input-number>）
                      Nhân：
                      <el-input-number class="loose-input" v-model="allData[11].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                      <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(11)" >Sửa đổi hàng loạt</el-button>
                    </el-breadcrumb-item>
                  </div>
                </el-card>
                <el-card class="box-card">
                  <el-button type="primary" class="handle-del mr10" @click="getData(11)" >Làm mới</el-button>
                  <el-table
                      :data="allData[11].tableData"
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
                            @click="handleEdit(11,scope.$index, scope.row)"
                        >Chỉnh sửa
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-card>
              </el-tab-pane>
              <el-tab-pane label="Giải nhất-Số đuôi đặc biệt" name="tab12">
                <el-card class="box-card">
                  <div>
                    <el-breadcrumb-item>
                      Nhân Key Phạm vi（
                      <el-input-number class="loose-input" v-model="allData[12].BatchForm.MinValue" :step="1" :min="allData[12].BatchForm.Min" :max="allData[12].BatchForm.Max"> </el-input-number>
                      ~
                      <el-input-number class="loose-input" v-model="allData[12].BatchForm.MaxValue" :step="1" :min="allData[12].BatchForm.Min" :max="allData[12].BatchForm.Max"> </el-input-number>）
                      Nhân：
                      <el-input-number class="loose-input" v-model="allData[12].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                      <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(12)" >Sửa đổi hàng loạt</el-button>
                    </el-breadcrumb-item>
                  </div>
                </el-card>
                <el-card class="box-card">
                  <el-button type="primary" class="handle-del mr10" @click="getData(12)" >Làm mới</el-button>
                  <el-table
                      :data="allData[12].tableData"
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
                            @click="handleEdit(12,scope.$index, scope.row)"
                        >Chỉnh sửa
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-card>
              </el-tab-pane>

              <el-tab-pane label="Đặc biệt giải đặc biệt mặt A" name="tab3">
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
                        <el-button type="primary" class="handle-del mr10" @click="getData(3 )" >Làm mới</el-button>
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
                          <el-table-column prop="AllUserMaxBet" label="oàn bộ"></el-table-column>

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
              <el-tab-pane label="Đặc biệt giải đặc biệt mặt B" name="tab15">
                <el-card class="box-card">
                  <div>
                    <el-breadcrumb-item>
                      Nhân Key Phạm vi（
                      <el-input-number class="loose-input" v-model="allData[15].BatchForm.MinValue" :step="1" :min="allData[15].BatchForm.Min" :max="allData[15].BatchForm.Max"> </el-input-number>
                      ~
                      <el-input-number class="loose-input" v-model="allData[15].BatchForm.MaxValue" :step="1" :min="allData[15].BatchForm.Min" :max="allData[15].BatchForm.Max"> </el-input-number>）
                      Nhân：
                      <el-input-number class="loose-input" v-model="allData[15].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                      <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(15)" >Sửa đổi hàng loạt</el-button>
                    </el-breadcrumb-item>
                  </div>
                </el-card>
                <el-card class="box-card">
                  <el-button type="primary" class="handle-del mr10" @click="getData(15)" >Làm mới</el-button>
                  <el-table
                      :data="allData[15].tableData"
                      border
                      class="table"
                      ref="multipleTable"
                      header-cell-class-name="table-header"
                  >
                    <!--                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>-->
                    <el-table-column prop="OddsType" label="NhânKey" width="90" align="center"></el-table-column>
                    <el-table-column prop="OddsDes" label="Tên nhân"></el-table-column>
                    <el-table-column prop="Odds" label="Nhân">
                      <template slot-scope="scope">{{scope.row.Odds|money}}</template>
                    </el-table-column>

                    <el-table-column prop="OneUserMinBet" label="Tối thiệu 1 người"></el-table-column>
                    <el-table-column prop="OneUserMaxBet" label="Tối đa 1 người"></el-table-column>
                    <el-table-column prop="AllUserMaxBet" label="oàn bộ"></el-table-column>

                    <el-table-column label="Điều hành" width="180" align="center">
                      <template slot-scope="scope">
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleEdit(15,scope.$index, scope.row)"
                        >Chỉnh sửa
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-card>
              </el-tab-pane>

              <el-tab-pane label="Đặc biệt giải nhất mặt A" name="tab4">
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
                        <el-button type="primary" class="handle-del mr10" @click="getData(4 )" >Làm mới</el-button>

                        <el-table
                                :data="allData[4].tableData"
                                border
                                class="table"
                                ref="multipleTable"
                                header-cell-class-name="table-header"
                        >
                            <!--                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>-->
                            <el-table-column prop="OddsType" label="NhânKey" width="90" align="center"></el-table-column>
                            <el-table-column prop="OddsDes" label="Tên nhân"></el-table-column>
                            <el-table-column prop="Odds" label="Nhân">
                                <template slot-scope="scope">{{scope.row.Odds|money}}</template>
                            </el-table-column>

                          <el-table-column prop="OneUserMinBet" label="Tối thiệu 1 người"></el-table-column>
                          <el-table-column prop="OneUserMaxBet" label="Tối đa 1 người"></el-table-column>
                          <el-table-column prop="AllUserMaxBet" label="oàn bộ"></el-table-column>

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
              <el-tab-pane label="Đặc biệt giải nhát  mặt B" name="tab16">
                <el-card class="box-card">
                  <div>
                    <el-breadcrumb-item>
                      NhânKey Phạm vi（
                      <el-input-number class="loose-input" v-model="allData[16].BatchForm.MinValue" :step="1" :min="allData[16].BatchForm.Min" :max="allData[16].BatchForm.Max"> </el-input-number>
                      ~
                      <el-input-number class="loose-input" v-model="allData[16].BatchForm.MaxValue" :step="1" :min="allData[16].BatchForm.Min" :max="allData[16].BatchForm.Max"> </el-input-number>）
                      Nhân：
                      <el-input-number class="loose-input" v-model="allData[16].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                      <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(16)" >Sửa đổi hàng loạt</el-button>
                    </el-breadcrumb-item>
                  </div>
                </el-card>
                <el-card class="box-card">
                  <el-button type="primary" class="handle-del mr10" @click="getData(16)" >Làm mới</el-button>

                  <el-table
                      :data="allData[16].tableData"
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
                    <el-table-column prop="AllUserMaxBet" label="oàn bộ"></el-table-column>

                    <el-table-column label="Điều hành" width="180" align="center">
                      <template slot-scope="scope">
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleEdit(16,scope.$index, scope.row)"
                        >Chỉnh sửa
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-card>
              </el-tab-pane>

              <el-tab-pane label="Đặc biệt giải nhì mặt A" name="tab5">
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
                      <el-button type="primary" class="handle-del mr10" @click="getData(5 )" >Làm mới</el-button>
                      <el-table
                              :data="allData[5].tableData"
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
                        <el-table-column prop="AllUserMaxBet" label="oàn bộ"></el-table-column>

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
              <el-tab-pane label="Đặc biệt giải nhì mặt B" name="tab17">
                <el-card class="box-card">
                  <div>
                    <el-breadcrumb-item>
                      Nhân Key Phạm vi（
                      <el-input-number class="loose-input" v-model="allData[17].BatchForm.MinValue" :step="1" :min="allData[17].BatchForm.Min" :max="allData[17].BatchForm.Max"> </el-input-number>
                      ~
                      <el-input-number class="loose-input" v-model="allData[17].BatchForm.MaxValue" :step="1" :min="allData[17].BatchForm.Min" :max="allData[17].BatchForm.Max"> </el-input-number>）
                      Nhân：
                      <el-input-number class="loose-input" v-model="allData[17].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                      <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(17)" >Sửa đổi hàng loạt</el-button>
                    </el-breadcrumb-item>
                  </div>
                </el-card>
                <el-card class="box-card">
                  <el-button type="primary" class="handle-del mr10" @click="getData(17)" >Làm mới</el-button>
                  <el-table
                      :data="allData[17].tableData"
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
                    <el-table-column prop="AllUserMaxBet" label="oàn bộ"></el-table-column>

                    <el-table-column label="Điều hành" width="180" align="center">
                      <template slot-scope="scope">
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleEdit(17,scope.$index, scope.row)"
                        >Chỉnh sửa
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-card>
              </el-tab-pane>


              <el-tab-pane label="Lô" name="tab6">
                    <el-card class="box-card">
                        <div>
                            <el-breadcrumb-item>
                                Nhân Key Phạm vi（
                                <el-input-number class="loose-input" v-model="allData[6].BatchForm.MinValue" :step="1" :min="allData[6].BatchForm.Min" :max="allData[6].BatchForm.Max"> </el-input-number>
                                ~
                                <el-input-number class="loose-input" v-model="allData[6].BatchForm.MaxValue" :step="1" :min="allData[6].BatchForm.Min" :max="allData[6].BatchForm.Max"> </el-input-number>）
                                Nhân：
                                <el-input-number class="loose-input" v-model="allData[6].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                                <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(6)" >Sửa đổi hàng loạt</el-button>
                            </el-breadcrumb-item>
                        </div>
                    </el-card>
                    <el-card class="box-card">
                        <el-button type="primary" class="handle-del mr10" @click="getData(6 )" >Làm mới</el-button>
                        <el-table
                                :data="allData[6].tableData"
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
                          <el-table-column prop="AllUserMaxBet" label="oàn bộ"></el-table-column>

                          <el-table-column label="Điều hành" width="180" align="center">
                                <template slot-scope="scope">
                                    <el-button
                                            type="text"
                                            icon="el-icon-edit"
                                            @click="handleEdit(6,scope.$index, scope.row)"
                                    >Chỉnh sửa
                                    </el-button>
                                </template>
                            </el-table-column>
                        </el-table>
                    </el-card>
                </el-tab-pane>

                <el-tab-pane label="Giải đặc biệt 3 Càng" name="tab7">
                    <el-card class="box-card">
                        <div>
                            <el-breadcrumb-item>
                                Nhân Key Phạm vi（
                                <el-input-number class="loose-input" v-model="allData[7].BatchForm.MinValue" :step="1" :min="allData[7].BatchForm.Min" :max="allData[7].BatchForm.Max"> </el-input-number>
                                ~
                                <el-input-number class="loose-input" v-model="allData[7].BatchForm.MaxValue" :step="1" :min="allData[7].BatchForm.Min" :max="allData[7].BatchForm.Max"> </el-input-number>）
                                Nhân：
                                <el-input-number class="loose-input" v-model="allData[7].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                                <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(7)" >Sửa đổi hàng loạt</el-button>
                            </el-breadcrumb-item>
                        </div>
                    </el-card>
                    <el-card class="box-card">
                        <el-button type="primary" class="handle-del mr10" @click="getData(7 )" >Làm mới</el-button>
                        <el-table
                                :data="allData[7].tableData"
                                border
                                class="table"
                                ref="multipleTable"
                                header-cell-class-name="table-header"
                        >
                            <!--                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>-->
                            <el-table-column prop="OddsType" label="NhânKey" width="90" align="center"></el-table-column>
                            <el-table-column prop="OddsDes" label="Tên nhân"></el-table-column>
                            <el-table-column prop="Odds" label="Nhân">
                                <template slot-scope="scope">{{scope.row.Odds|money}}</template>
                            </el-table-column>

                          <el-table-column prop="OneUserMinBet" label="Tối thiệu 1 người"></el-table-column>
                          <el-table-column prop="OneUserMaxBet" label="Tối đa 1 người"></el-table-column>
                          <el-table-column prop="AllUserMaxBet" label="oàn bộ"></el-table-column>

                          <el-table-column label="Điều hành" width="180" align="center">
                                <template slot-scope="scope">
                                    <el-button
                                            type="text"
                                            icon="el-icon-edit"
                                            @click="handleEdit(7,scope.$index, scope.row)"
                                    >Chỉnh sửa
                                    </el-button>
                                </template>
                            </el-table-column>
                        </el-table>
                    </el-card>
                </el-tab-pane>

              <el-tab-pane label="Giải nhất Xiên 2" name="tab18">
                <el-card class="box-card">
                  <div>
                    <el-breadcrumb-item>
                      Nhân Key Phạm vi（
                      <el-input-number class="loose-input" v-model="allData[24].BatchForm.MinValue" :step="1" :min="allData[24].BatchForm.Min" :max="allData[24].BatchForm.Max"> </el-input-number>
                      ~
                      <el-input-number class="loose-input" v-model="allData[24].BatchForm.MaxValue" :step="1" :min="allData[24].BatchForm.Min" :max="allData[24].BatchForm.Max"> </el-input-number>）
                      Nhân：
                      <el-input-number class="loose-input" v-model="allData[24].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                      <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(24)" >Sửa đổi hàng loạt</el-button>
                    </el-breadcrumb-item>
                  </div>
                </el-card>
                <el-card class="box-card">
                  <el-button type="primary" class="handle-del mr10" @click="getData(24 )" >Làm mới</el-button>
                  <el-table
                      :data="allData[24].tableData"
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
                    <el-table-column prop="AllUserMaxBet" label="oàn bộ"></el-table-column>

                    <el-table-column label="Điều hành" width="180" align="center">
                      <template slot-scope="scope">
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleEdit(24,scope.$index, scope.row)"
                        >Chỉnh sửa
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-card>
              </el-tab-pane>

              <el-tab-pane label="Xiên 2" name="tab8">
                    <el-card class="box-card">
                        <div>
                            <el-breadcrumb-item>
                                Nhân Key Phạm vi（
                                <el-input-number class="loose-input" v-model="allData[8].BatchForm.MinValue" :step="1" :min="allData[8].BatchForm.Min" :max="allData[8].BatchForm.Max"> </el-input-number>
                                ~
                                <el-input-number class="loose-input" v-model="allData[8].BatchForm.MaxValue" :step="1" :min="allData[8].BatchForm.Min" :max="allData[8].BatchForm.Max"> </el-input-number>）
                                Nhân：
                                <el-input-number class="loose-input" v-model="allData[8].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                                <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(8)" >Sửa đổi hàng loạ</el-button>
                            </el-breadcrumb-item>
                        </div>
                    </el-card>
                    <el-card class="box-card">
                        <el-button type="primary" class="handle-del mr10" @click="getData(8 )" >Làm mới</el-button>
                        <el-table
                                :data="allData[8].tableData"
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
                          <el-table-column prop="AllUserMaxBet" label="oàn bộ"></el-table-column>

                          <el-table-column label="Điều hành" width="180" align="center">
                                <template slot-scope="scope">
                                    <el-button
                                            type="text"
                                            icon="el-icon-edit"
                                            @click="handleEdit(8,scope.$index, scope.row)"
                                    >Chỉnh sửa
                                    </el-button>
                                </template>
                            </el-table-column>
                        </el-table>
                    </el-card>
                </el-tab-pane>
                <el-tab-pane label="Xiên 3" name="tab9">
                    <el-card class="box-card">
                        <div>
                            <el-breadcrumb-item>
                                Nhân Key Phạm vi（
                                <el-input-number class="loose-input" v-model="allData[9].BatchForm.MinValue" :step="1" :min="allData[9].BatchForm.Min" :max="allData[9].BatchForm.Max"> </el-input-number>
                                ~
                                <el-input-number class="loose-input" v-model="allData[9].BatchForm.MaxValue" :step="1" :min="allData[9].BatchForm.Min" :max="allData[9].BatchForm.Max"> </el-input-number>）
                                Nhân：
                                <el-input-number class="loose-input" v-model="allData[9].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                                <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(9)" >Sửa đổi hàng loạ</el-button>
                            </el-breadcrumb-item>
                        </div>
                    </el-card>
                    <el-card class="box-card">
                        <el-button type="primary" class="handle-del mr10" @click="getData(9 )" >Làm mới</el-button>
                        <el-table
                                :data="allData[9].tableData"
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
                          <el-table-column prop="AllUserMaxBet" label="oàn bộ"></el-table-column>

                          <el-table-column label="Điều hành" width="180" align="center">
                                <template slot-scope="scope">
                                    <el-button
                                            type="text"
                                            icon="el-icon-edit"
                                            @click="handleEdit(9,scope.$index, scope.row)"
                                    >Chỉnh sửa
                                    </el-button>
                                </template>
                            </el-table-column>
                        </el-table>
                    </el-card>
                </el-tab-pane>
                <el-tab-pane label="Số màu" name="tab10">
                    <el-card class="box-card">
                        <div>
                            <el-breadcrumb-item>
                                Nhân Key Phạm vi（
                                <el-input-number class="loose-input" v-model="allData[10].BatchForm.MinValue" :step="1" :min="allData[10].BatchForm.Min" :max="allData[10].BatchForm.Max"> </el-input-number>
                                ~
                                <el-input-number class="loose-input" v-model="allData[10].BatchForm.MaxValue" :step="1" :min="allData[10].BatchForm.Min" :max="allData[10].BatchForm.Max"> </el-input-number>）
                                Nhân：
                                <el-input-number class="loose-input" v-model="allData[10].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                                <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(10)" >Sửa đổi hàng loạ</el-button>
                            </el-breadcrumb-item>
                        </div>
                    </el-card>
                    <el-card class="box-card">
                        <el-button type="primary" class="handle-del mr10" @click="getData(10 )" >Làm mới</el-button>
                        <el-table
                                :data="allData[10].tableData"
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
                                            @click="handleEdit(10,scope.$index, scope.row)"
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
        name: 'Wsx2BbcOddsInfoSet',
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
                    {I:0, T: 'DXDS', BigTypes: [1,2,3,4], tableData: [] ,BatchForm:{MinValue:1,MaxValue:8,Odds:1.9,Min:1,Max:8},              },
                    {I:1, T: 'TT', BigTypes: [5], tableData: [] 				,BatchForm:{MinValue:10,MaxValue:19,Odds:9.9,Min:10,Max:19},          },
                    {I:2, T: 'WT', BigTypes: [6], tableData: [] 				,BatchForm:{MinValue:20,MaxValue:29,Odds:9.9,Min:20,Max:29},          },
                    {I:3, T: 'TDTM', BigTypes: [7], tableData: [] 			,BatchForm:{MinValue:100,MaxValue:199,Odds:99.9,Min:100,Max:199},     },
                    {I:4, T: '1DTM', BigTypes: [8], tableData: [] 			,BatchForm:{MinValue:200,MaxValue:299,Odds:99.9,Min:200,Max:299},     },
                    {I:5, T: '2DTM', BigTypes: [9], tableData: [] 			,BatchForm:{MinValue:300,MaxValue:399,Odds:99.9,Min:300,Max:399},     },
                    {I:6, T: '2LW', BigTypes: [10], tableData: [] 			,BatchForm:{MinValue:400,MaxValue:499,Odds:99.99,Min:400,Max:499},    },
                    {I:7, T: '3LW', BigTypes: [11], tableData: [] 			,BatchForm:{MinValue:1000,MaxValue:1999,Odds:999.9,Min:1000,Max:1999},},
                    {I:8, T: 'PM2WQ', BigTypes: [12], tableData: [] 		,BatchForm:{MinValue:60,MaxValue:60,Odds:99.9,Min:60,Max:60},     },
                    {I:9, T: 'PM3WQ', BigTypes: [13], tableData: [] 		,BatchForm:{MinValue:61,MaxValue:61,Odds:999.9,Min:61,Max:61},},
                    {I:10, T: 'BS', BigTypes: [15,16,17], tableData: [] 			,BatchForm:{MinValue:30,MaxValue:53,Odds:99.9,Min:30,Max:53},     },

                    {I:11, T: '1TT', BigTypes: [18], tableData: [] 				,BatchForm:{MinValue:610,MaxValue:619,Odds:9.9,Min:610,Max:619},          },
                    {I:12, T: '1WT', BigTypes: [19], tableData: [] 				,BatchForm:{MinValue:620,MaxValue:629,Odds:9.9,Min:620,Max:629},          },

                  {I:13, T: '2TT', BigTypes: [20], tableData: [] 				,BatchForm:{MinValue:630,MaxValue:639,Odds:9.9,Min:630,Max:639},          },
                  {I:14, T: '2WT', BigTypes: [21], tableData: [] 				,BatchForm:{MinValue:640,MaxValue:649,Odds:9.9,Min:640,Max:649},          },

                  {I:15, T: 'TDTM_B', BigTypes: [107], tableData: [] 			,BatchForm:{MinValue:20100,MaxValue:20199,Odds:99.9,Min:20100,Max:20199},     },
                  {I:16, T: '1DTM_B', BigTypes: [108], tableData: [] 			,BatchForm:{MinValue:20200,MaxValue:20299,Odds:99.9,Min:20200,Max:20299},     },
                  {I:17, T: '2DTM_B', BigTypes: [109], tableData: [] 			,BatchForm:{MinValue:20300,MaxValue:20399,Odds:99.9,Min:20300,Max:20399},     },

                  {I:18, T: 'TT_B', BigTypes: [105], tableData: [] 				,BatchForm:{MinValue:20010,MaxValue:20019,Odds:9.9,Min:20010,Max:20019},          },
                  {I:19, T: 'WT_B', BigTypes: [106], tableData: [] 				,BatchForm:{MinValue:20020,MaxValue:20029,Odds:9.9,Min:20020,Max:20029},          },
                  {I:20, T: '1TT_B', BigTypes: [118], tableData: [] 				,BatchForm:{MinValue:20610,MaxValue:20619,Odds:9.9,Min:20610,Max:20619},          },
                  {I:21, T: '1WT_B', BigTypes: [119], tableData: [] 				,BatchForm:{MinValue:20620,MaxValue:20629,Odds:9.9,Min:20620,Max:20629},          },
                  {I:22, T: '2TT_B', BigTypes: [120], tableData: [] 				,BatchForm:{MinValue:20630,MaxValue:20639,Odds:9.9,Min:20630,Max:20639},          },
                  {I:23, T: '2WT_B', BigTypes: [121], tableData: [] 				,BatchForm:{MinValue:20640,MaxValue:20649,Odds:9.9,Min:20640,Max:20649},          },
                  {I:24, T: 'T3LW', BigTypes: [22], tableData: [] 				,BatchForm:{MinValue:2000,MaxValue:2999,Odds:999.9,Min:2000,Max:2999},          },

                ],

                activeName:"tab0",

                multipleSelection: [],


                delList: [],
                editVisible: false,
                pageTotal: 0,

                form: {Id:0,OddsType:0,Odds:1.9, OddsDes:''},
                curFormT:0,
                curFormIdx: -1,
                GameType:202,
            };
        },
        created() {
          this.getData(0);
          this.getData(1);
          this.getData(2);
          this.getData(3);
          this.getData(4);
          this.getData(5);
          this.getData(6);
          this.getData(7);
          this.getData(8);
          this.getData(9);
          this.getData(10);
          this.getData(11);
          this.getData(12);
          // this.getData(13);
          // this.getData(14);
          this.getData(15);
          this.getData(16);
          this.getData(17);

          // this.getData(18);
          // this.getData(19);
          // this.getData(20);
          // this.getData(21);
          // this.getData(22);
          // this.getData(23);
          this.getData(24);


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
                        this.$message.success(`Sửa đổi ${this.form.OddsDes} Thành công`);
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
            // Kích hoạt nút tìm kiếm
            handleSearch() {
                this.getData();
            },
            // ửa hoChỉnh sạt độn
            handleEdit(t, index, row) {
                this.curFormIdx = index;
                this.curFormT = t
                this.form.Id = row.Id;
                this.form.OddsType = row.OddsType;
                this.form.OddsDes = row.OddsDes;
                this.form.Odds = row.Odds;
              this.form.OneUserMaxBet = row.OneUserMaxBet;
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
