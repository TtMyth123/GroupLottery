<template>
  <div>
    <div class="container">
      <div class="handle-box">
        <el-button
            type="primary"
            icon="el-icon-delete"
            class="handle-del mr10"
            @click="detailVisible=true"
        >添加奖结果</el-button>
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
        <el-table-column prop="CreatedAt" label="CreatedAt"></el-table-column>
        <el-table-column prop="UpdatedAt" label="UpdatedAt"></el-table-column>
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
    <el-dialog title="设置开奖结果" :visible.sync="detailVisible" width="90%">
      <el-form ref="form" :inline="true" :model="tableDataDetail" label-width="100px">
        <el-form-item label="期号" width="100">
          <el-input v-model="tableDataDetail.LotteryStr" ></el-input>
        </el-form-item>
        <el-divider></el-divider>
        <el-form-item label="号码1：">
          <el-input-number v-model.number="tableDataDetail.Num1" :step="1" :min="0" :max="28"></el-input-number>
        </el-form-item>
        <el-form-item label="号码2：">
          <el-input-number v-model.number="tableDataDetail.Num2" :step="1" :min="0" :max="28"></el-input-number>
        </el-form-item>
        <el-form-item label="号码3：">
          <el-input-number v-model.number="tableDataDetail.Num3" :step="1" :min="0" :max="28"></el-input-number>
        </el-form-item>
        <el-divider></el-divider>
      </el-form>
      <span slot="footer" class="dialog-footer">
            <el-button @click="detailVisible=false">取 消</el-button>
            <el-button @click="setAward">确 定</el-button>
        </span>
    </el-dialog>
  </div>
</template>
<script>
import VueCropper from 'vue-cropperjs';
import { request } from '../../utils/http';

export default {
  name: 'SetZg28XgAwardList',
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
      tableDataDetail:{GameName:"",LotteryStr:"",Num1:0,Num2:0,Num3:0},
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
      request({ url: 'api/getsetawardinfolist', method: 'post', data: this.query }).then((res) => {
        if (res.code == 200) {

          this.tableData = res.obj.ListData;
          this.pageTotal = res.obj.PageTotal;
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

    setAward(){
      let LotteryAward = this.tableDataDetail.Num1+","+this.tableDataDetail.Num2+","+this.tableDataDetail.Num3;
      let data = {
        GameType:this.query.GameType,
        LotteryAward:LotteryAward,
        LotteryStr:this.tableDataDetail.LotteryStr,
      }
      request({ url: 'api/setawardinfo', method: 'post', data: data }).then((res) => {
        if (res.code == 200) {
          this.detailVisible = false;
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
