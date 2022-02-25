<template>
  <div>
    <div class="container">
      <div class="handle-box">
        <el-button
            type="primary"
            icon="el-icon-delete"
            class="handle-del mr10"
            @click="testAwardVisible=true"
        >Thêm kết quả kiểm tra
        </el-button>
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
          header-cell-class-name="table-header" >
        <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>
        <el-table-column prop="GameName" label="Loại số" width="100"></el-table-column>
        <el-table-column prop="LotteryStr" label="Lượt số"></el-table-column>
        <el-table-column prop="ResultNums" label="Kết quả mở thưởng">
          <template slot-scope="scope">
            <el-button
                type="text"
                @click="openDetail(scope.$index, scope.row)"
            >{{scope.row.ResultNums}}
            </el-button>
          </template>
        </el-table-column>
        <el-table-column prop="CurLotteryTime" label="Thời gian mở thưởng"></el-table-column>
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

    <!-- Chỉnh sửa cửa sổ bật lên -->
    <el-dialog title="Chi tiết" :visible.sync="detailVisible" width="90%">

      <el-form ref="form" :inline="true" :model="tableDataDetail" label-width="100px">
        <el-form-item label="Loại số" width="100">
          <el-input v-model="tableDataDetail.GameName"></el-input>
        </el-form-item>
        <el-form-item label="Lượt số" width="100">
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
<!--            <el-button @click="detailVisible=false">Huỷ</el-button>-->
            <el-button @click="detailVisible=false">Xác nhận</el-button>
        </span>
    </el-dialog>

    <!-- Tạo hộp bật lên kết quả thử nghiệm -->
    <el-dialog title="Tạo kết quả kiểm tra" :visible.sync="testAwardVisible" width="150">
                <span slot="footer" class="dialog-footer">
            <el-button @click="testAwardVisible=false">Huỷ</el-button>
            <el-button @click="addTestAward">Xác nhận</el-button>
        </span>
    </el-dialog>
  </div>
</template>
<script>
import VueCropper from 'vue-cropperjs';
import { request } from '../../utils/http';

export default {
  name: 'Wsx2AwardList',
  data() {
    return {
      defaultSrc: require('../../assets/img/img.jpg'),

      query: {
        beginDay: '',
        endDay: '',
        pageIndex: 1,
        pageSize: 10,
        GameType:202,
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
    },
    // Chỉnh sửa hoạt động
    handleAddTestAward() {
    },

    addTestAward(){
      let data = {
        GameType:this.query.GameType
      }
      request({ url: 'api/addtestaward', method: 'post', data: data }).then((res) => {
        if (res.code == 200) {
          this.testAwardVisible = true
          this.getData();
          this.$message.success('Thêm vào thành công');

        } else {

          this.$message.error('Thêm vào không thành công：' + res.msg);
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
