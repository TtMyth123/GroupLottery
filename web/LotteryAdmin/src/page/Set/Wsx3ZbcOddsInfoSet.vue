<template>
    <div>
        <div class="container">
            <el-tabs v-model="activeName" >
                <el-tab-pane label="赔率" name="tab0">

                    <el-card class="box-card">
                        <el-button type="primary" class="handle-del mr10" @click="getData(0 )" >刷新</el-button>
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
        name: 'Wsx3ZbcOddsInfoSet',
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

                ],

                activeName:"tab0",

                multipleSelection: [],


                delList: [],
                editVisible: false,
                pageTotal: 0,

                form: {Id:0,OddsType:0,Odds:1.9, OddsDes:''},
                curFormT:0,
                curFormIdx: -1,
                GameType:203,
            };
        },
        created() {
          this.getData(0);
        },
        methods: {
            getData(t) {
                request({ url: 'api/getwsxzbcoddslist', method: 'post' ,data:{}}).then((res) => {
                    console.log(res);
                    if (res.code == 200) {
                        this.allData[t].tableData = res.obj;
                    }
                });
            },

            // 保存编辑
            saveOdds() {
                request({
                    url: 'api/savewsxzbcoddsinfo', method: 'post',
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
            // 触发搜索按钮
            handleSearch() {
                this.getData();
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
