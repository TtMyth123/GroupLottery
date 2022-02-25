<template>
  <div>
    <div class="container">
      <div class="handle-box">
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
      </div>
      <el-table
          :data="tableData"
          border
          class="table"
          ref="multipleTable"
          header-cell-class-name="table-header" >
        <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>
        <el-table-column prop="GameName" label="彩种" width="100"></el-table-column>
        <el-table-column prop="LotteryStr" label="期号"></el-table-column>
        <el-table-column prop="ResultNums" label="开奖结果">
          <template slot-scope="scope">
            {{scope.row.ResultNums}}

          </template>
        </el-table-column>
        <el-table-column prop="CurLotteryTime" label="开期时间"></el-table-column>
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
    <el-dialog title="明细" :visible.sync="detailVisible" width="90%">

      <el-form ref="form" :inline="true" :model="tableDataDetail" label-width="100px">
        <el-form-item label="采种" width="100">
          <el-input v-model="tableDataDetail.GameName"></el-input>
        </el-form-item>
        <el-form-item label="期号" width="100">
          <el-input v-model="tableDataDetail.LotteryStr" ></el-input>
        </el-form-item>

        <el-divider></el-divider>

        <el-form-item label="Jackpots">
          <el-input v-model="tableDataDetail.Result.jackpots"></el-input>
        </el-form-item>
        <el-divider></el-divider>

        <el-form-item label="FirstNum">
          <el-input v-model="tableDataDetail.Result.firstNum"></el-input>
        </el-form-item>
        <el-divider></el-divider>

        <el-form-item label="SecondNum">
          <el-input v-model="tableDataDetail.Result.secondNum[0]"></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="tableDataDetail.Result.secondNum[1]"></el-input>
        </el-form-item>
        <el-divider></el-divider>

        <el-form-item label="ThirdNum">
          <el-input v-model="tableDataDetail.Result.thirdNum[0]"></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="tableDataDetail.Result.thirdNum[1]"></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="tableDataDetail.Result.thirdNum[2]"></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="tableDataDetail.Result.thirdNum[3]"></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="tableDataDetail.Result.thirdNum[4]"></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="tableDataDetail.Result.thirdNum[5]"></el-input>
        </el-form-item>
        <el-divider></el-divider>

        <el-form-item label="forthNum">
          <el-input v-model="tableDataDetail.Result.forthNum[0]"></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="tableDataDetail.Result.forthNum[1]"></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="tableDataDetail.Result.forthNum[2]"></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="tableDataDetail.Result.forthNum[3]"></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="tableDataDetail.Result.forthNum[4]"></el-input>
        </el-form-item>
        <el-divider></el-divider>

        <el-form-item label="fifthNum">
          <el-input v-model="tableDataDetail.Result.fifthNum[0]"></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="tableDataDetail.Result.fifthNum[1]"></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="tableDataDetail.Result.fifthNum[2]"></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="tableDataDetail.Result.fifthNum[3]"></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="tableDataDetail.Result.fifthNum[4]"></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="tableDataDetail.Result.fifthNum[5]"></el-input>
        </el-form-item>
        <el-divider></el-divider>

        <el-form-item label="sixthNum">
          <el-input v-model="tableDataDetail.Result.sixthNum[0]"></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="tableDataDetail.Result.sixthNum[1]"></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="tableDataDetail.Result.sixthNum[2]"></el-input>
        </el-form-item>
        <el-divider></el-divider>

        <el-form-item label="seventhNum">
          <el-input v-model="tableDataDetail.Result.seventhNum[0]"></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="tableDataDetail.Result.seventhNum[1]"></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="tableDataDetail.Result.seventhNum[2]"></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="tableDataDetail.Result.seventhNum[3]"></el-input>
        </el-form-item>
        <el-divider></el-divider>



      </el-form>
      <span slot="footer" class="dialog-footer">
<!--            <el-button @click="detailVisible=false">取 消</el-button>-->
            <el-button @click="detailVisible=false">确 定</el-button>
        </span>
    </el-dialog>

    <!-- 生成测试结果弹出框 -->
    <el-dialog title="生成测试结果" :visible.sync="testAwardVisible" width="150">
                <span slot="footer" class="dialog-footer">
            <el-button @click="testAwardVisible=false">取 消</el-button>
            <el-button @click="addTestAward">确 定</el-button>
        </span>
    </el-dialog>
  </div>
</template>
<script>
import VueCropper from 'vue-cropperjs';
import { request } from '../../utils/http';

export default {
  name: 'Zg28XgAwardList',
  data() {
    return {
      defaultSrc: require('../../assets/img/img.jpg'),

      query: {
        beginDay: '',
        endDay: '',
        pageIndex: 1,
        pageSize: 10,
        GameType:44,
      },
      tableData: [],
      tableDataDetail:{GameName:"",LotteryStr:"",Result:{
          jackpots:"",firstNum:"",
          secondNum:["",""],thirdNum:["","","","","",""],
          forthNum:["","","","",""],fifthNum:["","","","","",""],
          sixthNum:["","",""],seventhNum:["","","",""]}},
      detailVisible: false,
      pageTotal: 0,
      rowId:0,
      testAwardVisible: false,

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

      return sums;
    },
    openDetail(index, row){
      this.rowId = row.Id;
      this.getDetailData()
    },
    // 获取 easy-mock 的模拟数据
    getData() {
      request({ url: 'api/getawardlist', method: 'post', data: this.query }).then((res) => {
        if (res.code == 200) {

          this.tableData = res.obj.ListData;
          this.pageTotal = res.obj.PageTotal;
        }
      });
    },
    getDetailData() {
      let data = {
        id:this.rowId
      }
      request({ url: 'api/getawarddetail', method: 'post', data: data }).then((res) => {
        if (res.code == 200) {
          console.log(res.obj);
          this.detailVisible = true
          this.tableDataDetail = res.obj
        }
      });
    },
    // 触发搜索按钮
    handleSearch() {
      this.$set(this.query, 'pageIndex', 1);
      this.getData();
    },
    // 编辑操作
    handleEdit(index, row) {
      this.idx = index;
      this.form = row;
      this.editVisible = true;
    },
    // 分页导航
    handlePageChange(val) {
      this.$set(this.query, 'pageIndex', val);
      this.getData();
    },
    // 编辑操作
    handleAddTestAward() {
    },

    addTestAward(){
      let data = {
        GameType:this.query.GameType
      }
      request({ url: 'api/addtestaward', method: 'post', data: data }).then((res) => {
        if (res.code == 200) {
          this.testAwardVisible = false;
          this.getData();
          this.$message.success('添加成功');
        } else {
          this.$message.error('添加失败：' + res.msg);
        }
      });
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
