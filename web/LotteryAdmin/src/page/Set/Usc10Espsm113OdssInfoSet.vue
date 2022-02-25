<template>
    <div>
        <div class="container">
            <el-tabs v-model="activeName" >
              <el-tab-pane label="冠亚军和" name="tab0">
                <el-card class="box-card">
                  <div>
                    <el-breadcrumb-item>
                      赔率Key范围（
                      <el-input-number class="loose-input" v-model="allData[0].BatchForm.MinValue" :step="1" :min="allData[0].BatchForm.Min" :max="allData[0].BatchForm.Max"> </el-input-number>
                      ~
                      <el-input-number class="loose-input" v-model="allData[0].BatchForm.MaxValue" :step="1" :min="allData[0].BatchForm.Min" :max="allData[0].BatchForm.Max"> </el-input-number>）
                      赔率：
                      <el-input-number class="loose-input" v-model="allData[0].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                      <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(0)" >批量修改</el-button>
                    </el-breadcrumb-item>
                  </div>
                </el-card>
                <el-card class="box-card">
                  <el-button type="primary" class="handle-del mr10" @click="getData(0)" >刷新</el-button>
                  <el-table
                      :data="allData[0].tableData"
                      border
                      class="table"
                      ref="multipleTable"
                      header-cell-class-name="table-header"
                  >
                    <!--                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>-->
                    <el-table-column prop="OddsType" label="赔率Key" width="90" align="center"></el-table-column>
                    <el-table-column prop="OddsDes" label="赔率名称"></el-table-column>
                    <el-table-column prop="Odds" label="赔率">
                      <template slot-scope="scope">{{scope.row.Odds|money}}</template>
                    </el-table-column>
                    <el-table-column prop="OneUserMinBet" label="一个用户最少"></el-table-column>
                    <el-table-column prop="OneUserMaxBet" label="一个用户最多"></el-table-column>
                    <el-table-column prop="AllUserMaxBet" label="全部用户最多"></el-table-column>

                    <el-table-column label="操作" width="180" align="center">
                      <template slot-scope="scope">
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleEdit(0, scope.$index, scope.row)"
                        >编辑
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-card>
              </el-tab-pane>
              <el-tab-pane label="冠军" name="tab1">
                <el-card class="box-card">
                  <div>
                    <el-breadcrumb-item>
                      赔率Key范围（
                      <el-input-number class="loose-input" v-model="allData[1].BatchForm.MinValue" :step="1" :min="allData[1].BatchForm.Min" :max="allData[1].BatchForm.Max"> </el-input-number>
                      ~
                      <el-input-number class="loose-input" v-model="allData[1].BatchForm.MaxValue" :step="1" :min="allData[1].BatchForm.Min" :max="allData[1].BatchForm.Max"> </el-input-number>）
                      赔率：
                      <el-input-number class="loose-input" v-model="allData[1].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                      <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(1)" >批量修改</el-button>
                    </el-breadcrumb-item>
                  </div>
                </el-card>
                <el-card class="box-card">
                  <el-button type="primary" class="handle-del mr10" @click="getData(1)" >刷新</el-button>
                  <el-table
                      :data="allData[1].tableData"
                      border
                      class="table"
                      ref="multipleTable"
                      header-cell-class-name="table-header"
                  >
                    <!--                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>-->
                    <el-table-column prop="OddsType" label="赔率Key" width="90" align="center"></el-table-column>
                    <el-table-column prop="OddsDes" label="赔率名称"></el-table-column>
                    <el-table-column prop="Odds" label="赔率">
                      <template slot-scope="scope">{{scope.row.Odds|money}}</template>
                    </el-table-column>
                    <el-table-column prop="OneUserMinBet" label="一个用户最少"></el-table-column>
                    <el-table-column prop="OneUserMaxBet" label="一个用户最多"></el-table-column>
                    <el-table-column prop="AllUserMaxBet" label="全部用户最多"></el-table-column>

                    <el-table-column label="操作" width="180" align="center">
                      <template slot-scope="scope">
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleEdit(1, scope.$index, scope.row)"
                        >编辑
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-card>
              </el-tab-pane>
              <el-tab-pane label="亚军" name="tab2">
                <el-card class="box-card">
                  <div>
                    <el-breadcrumb-item>
                      赔率Key范围（
                      <el-input-number class="loose-input" v-model="allData[2].BatchForm.MinValue" :step="1" :min="allData[2].BatchForm.Min" :max="allData[2].BatchForm.Max"> </el-input-number>
                      ~
                      <el-input-number class="loose-input" v-model="allData[2].BatchForm.MaxValue" :step="1" :min="allData[2].BatchForm.Min" :max="allData[2].BatchForm.Max"> </el-input-number>）
                      赔率：
                      <el-input-number class="loose-input" v-model="allData[2].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                      <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(2)" >批量修改</el-button>
                    </el-breadcrumb-item>
                  </div>
                </el-card>
                <el-card class="box-card">
                  <el-button type="primary" class="handle-del mr10" @click="getData(2)" >刷新</el-button>
                  <el-table
                      :data="allData[2].tableData"
                      border
                      class="table"
                      ref="multipleTable"
                      header-cell-class-name="table-header"
                  >
                    <!--                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>-->
                    <el-table-column prop="OddsType" label="赔率Key" width="90" align="center"></el-table-column>
                    <el-table-column prop="OddsDes" label="赔率名称"></el-table-column>
                    <el-table-column prop="Odds" label="赔率">
                      <template slot-scope="scope">{{scope.row.Odds|money}}</template>
                    </el-table-column>
                    <el-table-column prop="OneUserMinBet" label="一个用户最少"></el-table-column>
                    <el-table-column prop="OneUserMaxBet" label="一个用户最多"></el-table-column>
                    <el-table-column prop="AllUserMaxBet" label="全部用户最多"></el-table-column>

                    <el-table-column label="操作" width="180" align="center">
                      <template slot-scope="scope">
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleEdit(2, scope.$index, scope.row)"
                        >编辑
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-card>
              </el-tab-pane>
              <el-tab-pane label="第三名" name="tab3">
                <el-card class="box-card">
                  <div>
                    <el-breadcrumb-item>
                      赔率Key范围（
                      <el-input-number class="loose-input" v-model="allData[3].BatchForm.MinValue" :step="1" :min="allData[3].BatchForm.Min" :max="allData[3].BatchForm.Max"> </el-input-number>
                      ~
                      <el-input-number class="loose-input" v-model="allData[3].BatchForm.MaxValue" :step="1" :min="allData[3].BatchForm.Min" :max="allData[3].BatchForm.Max"> </el-input-number>）
                      赔率：
                      <el-input-number class="loose-input" v-model="allData[3].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                      <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(3)" >批量修改</el-button>
                    </el-breadcrumb-item>
                  </div>
                </el-card>
                <el-card class="box-card">
                  <el-button type="primary" class="handle-del mr10" @click="getData(3)" >刷新</el-button>
                  <el-table
                      :data="allData[3].tableData"
                      border
                      class="table"
                      ref="multipleTable"
                      header-cell-class-name="table-header"
                  >
                    <!--                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>-->
                    <el-table-column prop="OddsType" label="赔率Key" width="90" align="center"></el-table-column>
                    <el-table-column prop="OddsDes" label="赔率名称"></el-table-column>
                    <el-table-column prop="Odds" label="赔率">
                      <template slot-scope="scope">{{scope.row.Odds|money}}</template>
                    </el-table-column>
                    <el-table-column prop="OneUserMinBet" label="一个用户最少"></el-table-column>
                    <el-table-column prop="OneUserMaxBet" label="一个用户最多"></el-table-column>
                    <el-table-column prop="AllUserMaxBet" label="全部用户最多"></el-table-column>

                    <el-table-column label="操作" width="180" align="center">
                      <template slot-scope="scope">
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleEdit(3, scope.$index, scope.row)"
                        >编辑
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-card>
              </el-tab-pane>
              <el-tab-pane label="第四名" name="tab4">
                <el-card class="box-card">
                  <div>
                    <el-breadcrumb-item>
                      赔率Key范围（
                      <el-input-number class="loose-input" v-model="allData[4].BatchForm.MinValue" :step="1" :min="allData[4].BatchForm.Min" :max="allData[4].BatchForm.Max"> </el-input-number>
                      ~
                      <el-input-number class="loose-input" v-model="allData[4].BatchForm.MaxValue" :step="1" :min="allData[4].BatchForm.Min" :max="allData[4].BatchForm.Max"> </el-input-number>）
                      赔率：
                      <el-input-number class="loose-input" v-model="allData[4].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                      <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(4)" >批量修改</el-button>
                    </el-breadcrumb-item>
                  </div>
                </el-card>
                <el-card class="box-card">
                  <el-button type="primary" class="handle-del mr10" @click="getData(4)" >刷新</el-button>
                  <el-table
                      :data="allData[4].tableData"
                      border
                      class="table"
                      ref="multipleTable"
                      header-cell-class-name="table-header"
                  >
                    <!--                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>-->
                    <el-table-column prop="OddsType" label="赔率Key" width="90" align="center"></el-table-column>
                    <el-table-column prop="OddsDes" label="赔率名称"></el-table-column>
                    <el-table-column prop="Odds" label="赔率">
                      <template slot-scope="scope">{{scope.row.Odds|money}}</template>
                    </el-table-column>
                    <el-table-column prop="OneUserMinBet" label="一个用户最少"></el-table-column>
                    <el-table-column prop="OneUserMaxBet" label="一个用户最多"></el-table-column>
                    <el-table-column prop="AllUserMaxBet" label="全部用户最多"></el-table-column>

                    <el-table-column label="操作" width="180" align="center">
                      <template slot-scope="scope">
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleEdit(4, scope.$index, scope.row)"
                        >编辑
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-card>
              </el-tab-pane>
              <el-tab-pane label="第五名" name="tab5">
                <el-card class="box-card">
                  <div>
                    <el-breadcrumb-item>
                      赔率Key范围（
                      <el-input-number class="loose-input" v-model="allData[5].BatchForm.MinValue" :step="1" :min="allData[5].BatchForm.Min" :max="allData[5].BatchForm.Max"> </el-input-number>
                      ~
                      <el-input-number class="loose-input" v-model="allData[5].BatchForm.MaxValue" :step="1" :min="allData[5].BatchForm.Min" :max="allData[5].BatchForm.Max"> </el-input-number>）
                      赔率：
                      <el-input-number class="loose-input" v-model="allData[5].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                      <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(5)" >批量修改</el-button>
                    </el-breadcrumb-item>
                  </div>
                </el-card>
                <el-card class="box-card">
                  <el-button type="primary" class="handle-del mr10" @click="getData(5)" >刷新</el-button>
                  <el-table
                      :data="allData[5].tableData"
                      border
                      class="table"
                      ref="multipleTable"
                      header-cell-class-name="table-header"
                  >
                    <!--                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>-->
                    <el-table-column prop="OddsType" label="赔率Key" width="90" align="center"></el-table-column>
                    <el-table-column prop="OddsDes" label="赔率名称"></el-table-column>
                    <el-table-column prop="Odds" label="赔率">
                      <template slot-scope="scope">{{scope.row.Odds|money}}</template>
                    </el-table-column>
                    <el-table-column prop="OneUserMinBet" label="一个用户最少"></el-table-column>
                    <el-table-column prop="OneUserMaxBet" label="一个用户最多"></el-table-column>
                    <el-table-column prop="AllUserMaxBet" label="全部用户最多"></el-table-column>

                    <el-table-column label="操作" width="180" align="center">
                      <template slot-scope="scope">
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleEdit(5, scope.$index, scope.row)"
                        >编辑
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-card>
              </el-tab-pane>
              <el-tab-pane label="第六名" name="tab6">
                <el-card class="box-card">
                  <div>
                    <el-breadcrumb-item>
                      赔率Key范围（
                      <el-input-number class="loose-input" v-model="allData[6].BatchForm.MinValue" :step="1" :min="allData[6].BatchForm.Min" :max="allData[6].BatchForm.Max"> </el-input-number>
                      ~
                      <el-input-number class="loose-input" v-model="allData[6].BatchForm.MaxValue" :step="1" :min="allData[6].BatchForm.Min" :max="allData[6].BatchForm.Max"> </el-input-number>）
                      赔率：
                      <el-input-number class="loose-input" v-model="allData[6].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                      <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(6)" >批量修改</el-button>
                    </el-breadcrumb-item>
                  </div>
                </el-card>
                <el-card class="box-card">
                  <el-button type="primary" class="handle-del mr10" @click="getData(6)" >刷新</el-button>
                  <el-table
                      :data="allData[6].tableData"
                      border
                      class="table"
                      ref="multipleTable"
                      header-cell-class-name="table-header"
                  >
                    <!--                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>-->
                    <el-table-column prop="OddsType" label="赔率Key" width="90" align="center"></el-table-column>
                    <el-table-column prop="OddsDes" label="赔率名称"></el-table-column>
                    <el-table-column prop="Odds" label="赔率">
                      <template slot-scope="scope">{{scope.row.Odds|money}}</template>
                    </el-table-column>
                    <el-table-column prop="OneUserMinBet" label="一个用户最少"></el-table-column>
                    <el-table-column prop="OneUserMaxBet" label="一个用户最多"></el-table-column>
                    <el-table-column prop="AllUserMaxBet" label="全部用户最多"></el-table-column>

                    <el-table-column label="操作" width="180" align="center">
                      <template slot-scope="scope">
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleEdit(6, scope.$index, scope.row)"
                        >编辑
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-card>
              </el-tab-pane>
              <el-tab-pane label="第七名" name="tab7">
                <el-card class="box-card">
                  <div>
                    <el-breadcrumb-item>
                      赔率Key范围（
                      <el-input-number class="loose-input" v-model="allData[7].BatchForm.MinValue" :step="1" :min="allData[7].BatchForm.Min" :max="allData[7].BatchForm.Max"> </el-input-number>
                      ~
                      <el-input-number class="loose-input" v-model="allData[7].BatchForm.MaxValue" :step="1" :min="allData[7].BatchForm.Min" :max="allData[7].BatchForm.Max"> </el-input-number>）
                      赔率：
                      <el-input-number class="loose-input" v-model="allData[7].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                      <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(7)" >批量修改</el-button>
                    </el-breadcrumb-item>
                  </div>
                </el-card>
                <el-card class="box-card">
                  <el-button type="primary" class="handle-del mr10" @click="getData(7)" >刷新</el-button>
                  <el-table
                      :data="allData[7].tableData"
                      border
                      class="table"
                      ref="multipleTable"
                      header-cell-class-name="table-header"
                  >
                    <!--                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>-->
                    <el-table-column prop="OddsType" label="赔率Key" width="90" align="center"></el-table-column>
                    <el-table-column prop="OddsDes" label="赔率名称"></el-table-column>
                    <el-table-column prop="Odds" label="赔率">
                      <template slot-scope="scope">{{scope.row.Odds|money}}</template>
                    </el-table-column>
                    <el-table-column prop="OneUserMinBet" label="一个用户最少"></el-table-column>
                    <el-table-column prop="OneUserMaxBet" label="一个用户最多"></el-table-column>
                    <el-table-column prop="AllUserMaxBet" label="全部用户最多"></el-table-column>

                    <el-table-column label="操作" width="180" align="center">
                      <template slot-scope="scope">
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleEdit(7, scope.$index, scope.row)"
                        >编辑
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-card>
              </el-tab-pane>
              <el-tab-pane label="第八名" name="tab8">
                <el-card class="box-card">
                  <div>
                    <el-breadcrumb-item>
                      赔率Key范围（
                      <el-input-number class="loose-input" v-model="allData[8].BatchForm.MinValue" :step="1" :min="allData[8].BatchForm.Min" :max="allData[8].BatchForm.Max"> </el-input-number>
                      ~
                      <el-input-number class="loose-input" v-model="allData[8].BatchForm.MaxValue" :step="1" :min="allData[8].BatchForm.Min" :max="allData[8].BatchForm.Max"> </el-input-number>）
                      赔率：
                      <el-input-number class="loose-input" v-model="allData[8].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                      <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(8)" >批量修改</el-button>
                    </el-breadcrumb-item>
                  </div>
                </el-card>
                <el-card class="box-card">
                  <el-button type="primary" class="handle-del mr10" @click="getData(8)" >刷新</el-button>
                  <el-table
                      :data="allData[8].tableData"
                      border
                      class="table"
                      ref="multipleTable"
                      header-cell-class-name="table-header"
                  >
                    <!--                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>-->
                    <el-table-column prop="OddsType" label="赔率Key" width="90" align="center"></el-table-column>
                    <el-table-column prop="OddsDes" label="赔率名称"></el-table-column>
                    <el-table-column prop="Odds" label="赔率">
                      <template slot-scope="scope">{{scope.row.Odds|money}}</template>
                    </el-table-column>
                    <el-table-column prop="OneUserMinBet" label="一个用户最少"></el-table-column>
                    <el-table-column prop="OneUserMaxBet" label="一个用户最多"></el-table-column>
                    <el-table-column prop="AllUserMaxBet" label="全部用户最多"></el-table-column>

                    <el-table-column label="操作" width="180" align="center">
                      <template slot-scope="scope">
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleEdit(8, scope.$index, scope.row)"
                        >编辑
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-card>
              </el-tab-pane>
              <el-tab-pane label="第九名" name="tab9">
                <el-card class="box-card">
                  <div>
                    <el-breadcrumb-item>
                      赔率Key范围（
                      <el-input-number class="loose-input" v-model="allData[9].BatchForm.MinValue" :step="1" :min="allData[9].BatchForm.Min" :max="allData[9].BatchForm.Max"> </el-input-number>
                      ~
                      <el-input-number class="loose-input" v-model="allData[9].BatchForm.MaxValue" :step="1" :min="allData[9].BatchForm.Min" :max="allData[9].BatchForm.Max"> </el-input-number>）
                      赔率：
                      <el-input-number class="loose-input" v-model="allData[9].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                      <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(9)" >批量修改</el-button>
                    </el-breadcrumb-item>
                  </div>
                </el-card>
                <el-card class="box-card">
                  <el-button type="primary" class="handle-del mr10" @click="getData(9)" >刷新</el-button>
                  <el-table
                      :data="allData[9].tableData"
                      border
                      class="table"
                      ref="multipleTable"
                      header-cell-class-name="table-header"
                  >
                    <!--                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>-->
                    <el-table-column prop="OddsType" label="赔率Key" width="90" align="center"></el-table-column>
                    <el-table-column prop="OddsDes" label="赔率名称"></el-table-column>
                    <el-table-column prop="Odds" label="赔率">
                      <template slot-scope="scope">{{scope.row.Odds|money}}</template>
                    </el-table-column>
                    <el-table-column prop="OneUserMinBet" label="一个用户最少"></el-table-column>
                    <el-table-column prop="OneUserMaxBet" label="一个用户最多"></el-table-column>
                    <el-table-column prop="AllUserMaxBet" label="全部用户最多"></el-table-column>

                    <el-table-column label="操作" width="180" align="center">
                      <template slot-scope="scope">
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleEdit(9, scope.$index, scope.row)"
                        >编辑
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-card>
              </el-tab-pane>
              <el-tab-pane label="第十名" name="tab10">
                <el-card class="box-card">
                  <div>
                    <el-breadcrumb-item>
                      赔率Key范围（
                      <el-input-number class="loose-input" v-model="allData[10].BatchForm.MinValue" :step="1" :min="allData[10].BatchForm.Min" :max="allData[10].BatchForm.Max"> </el-input-number>
                      ~
                      <el-input-number class="loose-input" v-model="allData[10].BatchForm.MaxValue" :step="1" :min="allData[10].BatchForm.Min" :max="allData[10].BatchForm.Max"> </el-input-number>）
                      赔率：
                      <el-input-number class="loose-input" v-model="allData[10].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                      <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(10)" >批量修改</el-button>
                    </el-breadcrumb-item>
                  </div>
                </el-card>
                <el-card class="box-card">
                  <el-button type="primary" class="handle-del mr10" @click="getData(10)" >刷新</el-button>
                  <el-table
                      :data="allData[10].tableData"
                      border
                      class="table"
                      ref="multipleTable"
                      header-cell-class-name="table-header"
                  >
                    <!--                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>-->
                    <el-table-column prop="OddsType" label="赔率Key" width="90" align="center"></el-table-column>
                    <el-table-column prop="OddsDes" label="赔率名称"></el-table-column>
                    <el-table-column prop="Odds" label="赔率">
                      <template slot-scope="scope">{{scope.row.Odds|money}}</template>
                    </el-table-column>
                    <el-table-column prop="OneUserMinBet" label="一个用户最少"></el-table-column>
                    <el-table-column prop="OneUserMaxBet" label="一个用户最多"></el-table-column>
                    <el-table-column prop="AllUserMaxBet" label="全部用户最多"></el-table-column>

                    <el-table-column label="操作" width="180" align="center">
                      <template slot-scope="scope">
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleEdit(10, scope.$index, scope.row)"
                        >编辑
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-card>
              </el-tab-pane>
              <el-tab-pane label="番摊" name="tab11">
                <el-card class="box-card">
                  <div>
                    <el-breadcrumb-item>
                      赔率Key范围（
                      <el-input-number class="loose-input" v-model="allData[11].BatchForm.MinValue" :step="1" :min="allData[11].BatchForm.Min" :max="allData[11].BatchForm.Max"> </el-input-number>
                      ~
                      <el-input-number class="loose-input" v-model="allData[11].BatchForm.MaxValue" :step="1" :min="allData[11].BatchForm.Min" :max="allData[11].BatchForm.Max"> </el-input-number>）
                      赔率：
                      <el-input-number class="loose-input" v-model="allData[11].BatchForm.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                      <el-button type="primary" class="handle-del mr10" @click="batchModifyOdds(11)" >批量修改</el-button>
                    </el-breadcrumb-item>
                  </div>
                </el-card>
                <el-card class="box-card">
                  <el-button type="primary" class="handle-del mr10" @click="getData(11)" >刷新</el-button>
                  <el-table
                      :data="allData[11].tableData"
                      border
                      class="table"
                      ref="multipleTable"
                      header-cell-class-name="table-header"
                  >
                    <!--                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>-->
                    <el-table-column prop="OddsType" label="赔率Key" width="90" align="center"></el-table-column>
                    <el-table-column prop="OddsDes" label="赔率名称"></el-table-column>
                    <el-table-column prop="Odds" label="赔率">
                      <template slot-scope="scope">{{scope.row.Odds|money}}</template>
                    </el-table-column>
                    <el-table-column prop="OneUserMinBet" label="一个用户最少"></el-table-column>
                    <el-table-column prop="OneUserMaxBet" label="一个用户最多"></el-table-column>
                    <el-table-column prop="AllUserMaxBet" label="全部用户最多"></el-table-column>

                    <el-table-column label="操作" width="180" align="center">
                      <template slot-scope="scope">
                        <el-button
                            type="text"
                            icon="el-icon-edit"
                            @click="handleEdit(11, scope.$index, scope.row)"
                        >编辑
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-card>
              </el-tab-pane>
            </el-tabs>

            <!-- 编辑弹出框 -->
            <el-dialog title="修改赔率" :visible.sync="editVisible" width="90%">
                <el-form ref="form" :model="form"  label-width="100px">
                    <el-form-item label="赔率Key" prop="OddsType">
                        <el-input v-model="form.OddsType" disabled></el-input>
                    </el-form-item>
                    <el-form-item label="赔率名称" prop="OddsDes">
                        <el-input v-model="form.OddsDes" disabled></el-input>
                    </el-form-item>
                    <el-form-item label="赔率" prop="Odds">
                        <el-input-number class="loose-input" v-model="form.Odds" :precision="2" :step="1" :min="0"></el-input-number>
                    </el-form-item>
                  <el-form-item label="一个用户最多" prop="OneUserMaxBet">
                    <el-input-number class="loose-input" v-model="form.OneUserMaxBet" :step="1000" :min="0"></el-input-number>
                  </el-form-item>
                  <el-form-item label="一个用户最少" prop="OneUserMinBet">
                    <el-input-number class="loose-input" v-model="form.OneUserMinBet" :step="10" :min="1"></el-input-number>
                  </el-form-item>
                  <el-form-item label="全部用户最多" prop="AllUserMaxBet">
                    <el-input-number class="loose-input" v-model="form.AllUserMaxBet" :step="1000" :min="0"></el-input-number>
                  </el-form-item>
                </el-form>
                <span slot="footer" class="dialog-footer">
                <el-button @click="editVisible = false">取 消</el-button>
                <el-button type="primary" @click="saveOdds">确 定</el-button>
            </span>
            </el-dialog>
        </div>
    </div>

</template>

<script>
    import { request } from '../../utils/http';

    export default {
        name: 'Usc10Espsm113OdssInfoSet',
        data() {
            return {
                AwardTypeItem: [{ name: '金币', value: 1 }, { name: '积分', value: 2 }],
                query: {
                    address: '',
                    name: '',
                    pageIndex: 1,
                    pageSize: 10
                },
                allData:[
                  {I:0, BigTypes: [1], tableData:[] ,BatchForm:{MinValue:1,MaxValue:21,Odds:9.9,Min:1,Max:21}, },
                  {I:1, BigTypes: [2], tableData:[] ,BatchForm:{MinValue:22,MaxValue:37,Odds:9.9,Min:22,Max:37}, },
                  {I:2, BigTypes: [3], tableData:[] ,BatchForm:{MinValue:38,MaxValue:53,Odds:9.9,Min:38,Max:53}, },
                  {I:3, BigTypes: [4], tableData:[] ,BatchForm:{MinValue:54,MaxValue:69,Odds:9.9,Min:54,Max:69}, },
                  {I:4, BigTypes: [5], tableData:[] ,BatchForm:{MinValue:70,MaxValue:85,Odds:9.9,Min:70,Max:85}, },
                  {I:5, BigTypes: [6], tableData:[] ,BatchForm:{MinValue:86,MaxValue:101,Odds:9.9,Min:86,Max:101}, },
                  {I:6, BigTypes: [7], tableData:[] ,BatchForm:{MinValue:102,MaxValue:115,Odds:9.9,Min:102,Max:115}, },
                  {I:7, BigTypes: [8], tableData:[] ,BatchForm:{MinValue:118,MaxValue:131,Odds:9.9,Min:118,Max:131}, },
                  {I:8, BigTypes: [9], tableData:[] ,BatchForm:{MinValue:134,MaxValue:147,Odds:9.9,Min:134,Max:147}, },
                  {I:9, BigTypes: [10], tableData:[] ,BatchForm:{MinValue:150,MaxValue:163,Odds:9.9,Min:150,Max:163}, },
                  {I:10, BigTypes: [11], tableData:[] ,BatchForm:{MinValue:166,MaxValue:179,Odds:9.9,Min:166,Max:179}, },
                  {I:11, BigTypes: [12], tableData:[] ,BatchForm:{MinValue:182,MaxValue:207,Odds:9.9,Min:182,Max:207}, },

                ],
                activeName:"tab0",
                multipleSelection: [],
                delList: [],
                editVisible: false,
                pageTotal: 0,
                form: {Id:0,OddsType:0,Odds:1.9, OddsDes:''},
                curFormT:0,
                curFormIdx: -1,
                GameType:113,
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
                        this.$message.success(`修改 ${this.form.OddsDes} 成功`);
                        this.getData(t, this.allData[t].BigTypes)
                    } else {
                        this.$message.error('更新失败：' + res.msg);
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

            // 保存编辑
            saveOdds() {
                request({
                    url: 'api/saveoddsinfo', method: 'post',
                    data: this.form
                }).then((res) => {
                    console.log(res);
                    if (res.code == 200) {
                        this.editVisible = false;
                        this.$message.success(`修改 ${this.form.OddsDes} 成功`);
                        this.getData(this.curFormT, this.allData[this.curFormT].BigTypes)
                    } else {
                        this.$message.error('更新失败：' + res.msg);
                    }
                });
            },
            // 编辑操作
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
