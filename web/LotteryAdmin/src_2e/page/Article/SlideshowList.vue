<template>
  <div>
    <div class="container">
      <div class="handle-box">
        <el-select v-model.number="query.state" placeholder="状态">
          <el-option v-for="item in StateTypeAll" :label="item.name" :key="item.value"
                     :value="item.value"></el-option>
        </el-select>
        <el-button type="primary" icon="el-icon-search" @click="handleSearch">搜索</el-button>
        <el-button
            type="primary"
            icon="el-icon-delete"
            class="handle-del mr10"
            @click="handleAdd"
        >添加</el-button>
      </div>
      <el-table
          :data="tableData"
          border
          class="table"
          ref="multipleTable"
          header-cell-class-name="table-header"  >
        <!--                <el-table-column type="selection" width="55" align="center"></el-table-column>-->
        <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>
        <el-table-column prop="Seq" label="顺序" width="100"></el-table-column>
        <el-table-column prop="Title" label="标题" width="100"></el-table-column>
        <el-table-column label="标题图片">
          <template slot-scope="scope">
            <el-image  style="width: 200px; height: 100px"  :src="scope.row.TitleImgUrl" lazy></el-image>
          </template>
        </el-table-column>
        <el-table-column prop="UpdatedAt" label="更新时间"></el-table-column>
        <el-table-column prop="State" label="状态">
          <template slot-scope="scope">{{scope.row.State|ArticleStateName}}</template>
        </el-table-column>
        <el-table-column label="操作" width="140" align="center">
          <template slot-scope="scope">
            <el-button
                type="text"
                icon="el-icon-edit"
                @click="handleEdit(scope.$index, scope.row)"
            >编辑</el-button>
            <el-button
                type="text"
                icon="el-icon-delete"
                class="red"
                @click="handleDelete(scope.$index, scope.row)"
            >删除</el-button>
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
    <el-dialog :title="dialogTitle" :visible.sync="editVisible" width="90%">
      <el-form ref="form" :model="form"  label-width="100px">
        <el-form-item label="顺序" prop="Seq">
          <el-input-number class="loose-input" v-model="form.Seq" :step="1" :min="0"></el-input-number>
        </el-form-item>
        <el-form-item label="标题" prop="Title">
          <el-input v-model="form.Title"></el-input>
        </el-form-item>
        <el-form-item label="图片" prop="TitleImgUrl">
          <el-upload
              class="avatar-uploader"
              action="kit/uploadpic"
              :data="uploadPicData"
              accept="image/jpeg,image/gif,image/png"
              :before-upload="beforeAvatarUpload"
              :on-change="handlePicUrl"
              :on-success="uploadSuccess"
              :show-file-list="false" >
            <img v-if="form.TitleImgUrl" :src="form.TitleImgUrl" class="avatar">
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>
<!--          <button v-if="form.TitleImgUrl" @click="delPicUrl()"  >删除</button>-->

        </el-form-item>

        <el-form-item label="状态" prop="OddsType">
          <el-select v-model.number="form.State" placeholder="请选择状态">
            <el-option v-for="item in StateTypes" :label="item.name" :key="item.value"
                       :value="item.value"></el-option>
          </el-select>
        </el-form-item>

      </el-form>
      <span slot="footer" class="dialog-footer">
                <el-button @click="editVisible = false">取 消</el-button>
                <el-button type="primary" @click="SaveNotice">确 定</el-button>
            </span>
    </el-dialog>
  </div>
</template>
<script>
import VueCropper from 'vue-cropperjs';
import { request } from '../../utils/http';

export default {
  name: 'SlideshowList',
  data() {
    return {
      defaultSrc: require('../../assets/img/img.jpg'),
      query: {
        articleType:5,
        state:0,
        pageIndex: 1,
        pageSize: 10
      },
      tableData: [],
      GroupBetId:0,
      tableDataDetail:[],
      editVisible: false,
      pageTotal: 0,
      dialogTitle:'修改',
      picGuid:'',
      uploadPicData:{picGuid:''},
      StateTypeAll: [
        { name: '全部', value: 0 },
        { name: '新建', value: 1 },
        { name: '启动', value: 2 },
        { name: '禁用', value: 3 },
      ],
      StateTypes: [
        { name: '启动', value: 2 },
        { name: '禁用', value: 3 },
      ],
      form:{Id:0,Seq:1,Title:'',Des:'',State:2,TitleImgUrl:''},
      TitleImgUrlFile:{},
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
      request({ url: 'api/getarticlelist', method: 'post', data: this.query }).then((res) => {
        if (res.code == 200) {
          console.log(res.obj.ListData);
          this.tableData = res.obj.ListData;
          this.pageTotal = res.obj.PageTotal;
        }
      });
    },
    SaveNotice() {
      this.form.ArticleType = this.query.articleType;
      this.form.PicGuid = this.uploadPicData.picGuid
      request({ url: 'api/savearticlehastmppic', method: 'post', data: this.form }).then((res) => {
        if (res.code == 200) {
          this.getData();
          this.editVisible = false;
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
      this.form.Id = row.Id;
      this.form.Seq = row.Seq;
      this.form.Title = row.Title;
      this.form.Des = row.Des;
      this.form.State = row.State;
      this.form.TitleImgUrl = row.TitleImgUrl;
      this.uploadPicData.picGuid = '';
      request({
        url:'kit/getguid', method:'post',
      }).then((res)=>{
        if (res.code==200){
          this.picGuid = res.obj;
          this.editVisible = true;
        } else {
          this.$message.error('获取数据失败：'+res.msg);
        }
      });

    },
    handleDelete(index, row) {
      // 二次确认删除
      this.$confirm('确定要删除吗？', '提示', {
        type: 'warning'
      })
          .then(() => {
            var data = {
              Id:row.Id
            }
            request({
              url:'api/delarticle', method:'post',
              data:data
            }).then((res)=>{
              console.log(res);
              if (res.code==200){
                this.$message.success('删除成功');
                this.getData();
                this.editVisible = false;
              } else {
                this.$message.error('删除失败：'+res.msg);
              }
            });

          })
          .catch(() => {});
    },
    handleAdd() {
      this.idx =-1;

      this.form.Id = 0;
      this.form.Seq = 1;
      this.form.Title = '';
      this.form.Des = '';
      this.form.State = 2;
      this.form.TitleImgUrl = '';
      this.uploadPicData.picGuid = '';
      request({
        url:'kit/getguid', method:'post',
      }).then((res)=>{
        if (res.code==200){
          this.picGuid = res.obj;
          this.editVisible = true;
        } else {
          this.$message.error('获取数据失败：'+res.msg);
        }
      });
    },
    uploadSuccess(res,file, fileList){
      console.log("ok" );
    },
    handlePicUrl(file, fileList) {
      console.log("aaa" );
      this.form.TitleImgUrl = URL.createObjectURL(file.raw);
      console.log(this.form.TitleImgUrl );
    },
    clear(){
      this.form.TitleImgUrl = "";
    },
    beforeAvatarUpload(file) {
      const isIMAGE = file.type === 'image/jpeg'||'image/gif'||'image/png';
      if (!isIMAGE) {
        this.$message.error('上传文件只能是图片格式!');
      }

      const isLt2M = file.size / 1024  < 300;
      if (!isLt2M) {
        this.$message.error('上传文件大小不能超过 300bk!');
      }
      if (isLt2M&& isIMAGE) {
        this.uploadPicData.picGuid = this.picGuid;
      }

      return isLt2M&& isIMAGE;
    },

    delPicUrl(){
      this.form.TitleImgUrl = "";
      this.TitleImgUrlFile = '';
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

.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-uploader .el-upload:hover {
  border-color: #409EFF;
}
.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 111px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}
.avatar {
  width: 111px;
  height: 178px;
  display: block;
}

.avatar-actor-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-actor-uploader .el-upload:hover {
  border-color: #409EFF;
}
.avatar-actor-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 80px;
  height: 80px;
  line-height: 80px;
  text-align: center;
}
.avatar-actor {
  width: 80px;
  height: 80px;
  display: block;
}
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
