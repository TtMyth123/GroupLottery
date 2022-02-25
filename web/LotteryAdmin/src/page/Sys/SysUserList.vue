<template>
    <div>
        <div class="container">
            <div class="handle-box">
                <el-button
                        type="primary"
                        icon="el-icon-delete"
                        class="handle-del mr10"
                        @click="handleAddUser"
                >添加管理用户
                </el-button>
              <el-button
                  type="primary"
                  icon="el-icon-delete"
                  class="handle-del mr10"
                  @click="handleAddAgentUser"
              >添加代理商
              </el-button>
            </div>
            <el-table
                    :data="tableData"
                    border
                    class="table"
                    header-cell-class-name="table-header">
                <!--                <el-table-column type="selection" width="55" align="center"></el-table-column>-->
                <el-table-column prop="Id" label="ID" width="55" align="center"></el-table-column>
                <!--                <el-table-column prop="Seq" label="序号" width="55" align="center"></el-table-column>-->
                <el-table-column prop="UserName" label="用户名" width="100"></el-table-column>
              <el-table-column prop="UserType" label="用户类型" width="100">
                <template slot-scope="scope">
                  <div  v-if="scope.row['UserType']==1">
                    代理商
                  </div>
                  <div v-else>
                    管理员
                  </div>

                </template>
              </el-table-column>
                <el-table-column label="操作" width="140" align="center">
                    <template slot-scope="scope">
                        <el-button  v-if="scope.row['UserType']==1"
                                type="text"
                                icon="el-icon-edit"
                                @click="handleEdit(scope.$index, scope.row)"
                        >编辑
                        </el-button>
                      <el-button  v-else
                                  type="text"
                                  icon="el-icon-edit"
                                  @click="handleEdit(scope.$index, scope.row)"
                      >编辑
                      </el-button>
                        <el-button
                                type="text"
                                icon="el-icon-delete"
                                class="red"
                                @click="handleDelete(scope.$index, scope.row)"
                        >删除
                        </el-button>
                        <el-button
                                type="text"
                                icon="el-icon-delete"
                                class="red"
                                @click="handleModifyPermission(scope.$index, scope.row)"
                        >权限
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
        <el-dialog title="管理员修改" :visible.sync="editFormVisible" width="90%">
            <el-form ref="form" :model="form" :rules="rules" label-width="100px">
                <el-form-item label="用户名" prop="UserName">
                    <el-input v-model="form.UserName"></el-input>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="editFormVisible = false">取 消</el-button>
                <el-button type="primary" @click="saveEdit">确 定</el-button>
            </span>
        </el-dialog>

      <!-- 编辑弹出框 -->
      <el-dialog title="代理商修改" :visible.sync="editAgentFormVisible" width="90%">
        <el-form ref="form" :model="editAgentForm" :rules="rules" label-width="100px">
          <el-form-item label="用户名" prop="UserName">
            <el-input v-model="editAgentForm.UserName"></el-input>
          </el-form-item>
          <el-form-item label="游戏用户ID" prop="GameId">
            <el-input v-model="editAgentForm.GameId" readonly></el-input>
          </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
                <el-button @click="editAgentFormVisible = false">取 消</el-button>
                <el-button type="primary" @click="saveAgentEdit">确 定</el-button>
            </span>
      </el-dialog>

      <el-dialog title="添加代理" :visible.sync="addAgentFormVisible" width="300">
        <el-form ref="form" :model="formAddAgent" :rules="rules" label-width="100px">
          <el-form-item label="用户名" prop="UserName">
            <el-input v-model="formAddAgent.UserName"></el-input>
          </el-form-item>
          <el-form-item label="密码" prop="Pwd">
            <el-input v-model="formAddAgent.Pwd"></el-input>
          </el-form-item>
          <el-form-item label="代理游戏ID" prop="Pwd">
            <el-input v-model="formAddAgent.GameId"></el-input>
          </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
                <el-button @click="addAgentFormVisible = false">取 消</el-button>
                <el-button type="primary" @click="onAddAgentSysUser">确 定</el-button>
            </span>
      </el-dialog>
      <el-dialog title="添加管理员" :visible.sync="addFormVisible" width="300">
        <el-form ref="form" :model="formAdd" :rules="rules" label-width="100px">
          <el-form-item label="用户名" prop="UserName">
            <el-input v-model="formAdd.UserName"></el-input>
          </el-form-item>
          <el-form-item label="密码" prop="Pwd">
            <el-input v-model="formAdd.Pwd"></el-input>
          </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
                <el-button @click="addFormVisible = false">取 消</el-button>
                <el-button type="primary" @click="onAddManageSysUser">确 定</el-button>
            </span>
      </el-dialog>

      <el-dialog title="权限设置" :visible.sync="permissionFormVisible" width="90%">
            <el-table
                    :data="tableDataMenu"
                    border
                    class="table"
                    header-cell-class-name="table-header">
                <el-table-column prop="TypeName" label="分类" width="100"></el-table-column>
                <el-table-column prop="Title" label="功能" width="180"></el-table-column>
                <el-table-column prop="Power" label="分类" width="250">
                    <template slot-scope="scope">
                        <el-radio-group v-model="scope.row.Power">
                            <el-radio :label="0">禁用</el-radio>
                            <el-radio :label="1">读</el-radio>
                            <el-radio :label="2">写</el-radio>
                        </el-radio-group>
                    </template>

                </el-table-column>
            </el-table>

            <span slot="footer" class="dialog-footer">
                <el-button @click="permissionFormVisible = false">取 消</el-button>
                <el-button type="primary" @click="savePermission">确 定</el-button>
            </span>
        </el-dialog>

    </div>
</template>
<script>
    // import 'quill/dist/quill.core.css';
    import 'quill/dist/quill.snow.css';
    // import 'quill/dist/quill.bubble.css';
    import VueCropper from 'vue-cropperjs';
    import { request } from '../../utils/http';

    export default {
        name: 'SysUserList',
        data() {
            return {
                query: {
                    userId: 0,
                    userName: '',
                    pageIndex: 1,
                    pageSize: 10
                },
                tableData: [],
                tableDataMenu: [],
                pageTotal: 0,
                multipleSelection: [],
              addAgentFormVisible: false,
              formAddAgent: { Id: 0, UserName: '', Pwd: '',GameId:0 },

                addFormVisible: false,
                formAdd: { Id: 0, UserName: '', Pwd: '' },

              editAgentFormVisible:false,
              editAgentForm: { Id: 0, UserName: '' },

                editFormVisible: false,
                form: { Id: 0, UserName: '' },
                idx: -1,
                id: -1,

                curUserId:0,
                permissionFormVisible: false,
                permissionData: [
                ],
                defaultProps: {
                    children: 'children',
                    label: 'title',
                },
                rules: {
                    UserName: [
                        { required: true, message: '请输入用户名', trigger: 'blur' },
                        { min: 1, max: 50, message: '长度在 2 到 50 个字符', trigger: 'blur' }
                    ]
                }

            };
        },
        created() {
            this.getData();
        },
        mounted() {
            // this.editor = this.$refs.myQuillEditor.quill;

        },

        beforeDestroy() {
            // this.editor = null;
            // delete this.editor;
        },
        components: {
            VueCropper
        },
        methods: {
            //鼠标单击的事件
            onClick(e, editor) {
                console.log('Element clicked');
                console.log(e);
                console.log(editor);
            },
            //清空内容
            clear() {
                this.$refs.editor.clear();
            },
            // 准备富文本编辑器
            onEditorReady(editor) {
            },
// 富文本编辑器 失去焦点事件
            onEditorBlur(editor) {
            },
// 富文本编辑器 获得焦点事件
            onEditorFocus(editor) {
            },
// 富文本编辑器 内容改变事件
            onEditorChange(editor) {
            },
            // 富文本编辑器 点击插入图片或者视频
            insertImgClick(e) {
                if (e.target.className.indexOf('icon-pic') != -1) {
                    document.getElementById('insert_image').click();
                } else if (e.target.className.indexOf('icon-video') != -1) {
                    document.getElementById('insert_video').click();
                }
            },
            handleUploadSuccess(response, file, fileList) {
                let quill = this.$refs.myQuillEditor.quill;
                let length = quill.getSelection().index;
                quill.insertEmbed(length, 'image', response);
                quill.setSelection(length + 1);
            },
            // 获取 easy-mock 的模拟数据
            getData() {
                request({ url: 'sys/getsysuserlist', method: 'post', data: this.query }).then((res) => {
                    if (res.code == 200) {
                        this.tableData = res.obj.ListData;
                        this.pageTotal = res.obj.PageTotal;
                    }
                });
            },
            getFunMenuPermission() {
                let data = {
                    UserId:this.curUserId
                }
                request({ url: 'sys/getpermissiontree', method: 'post', data: data }).then((res) => {
                    if (res.code == 200) {
                        this.tableDataMenu = res.obj
                    }
                });
            },
            // 触发搜索按钮
            handleSearch() {
                this.$set(this.query, 'pageIndex', 1);
                this.getData();
            },
            handleModifyPermission(index, row) {
                this.curUserId = row.Id;
                this.getFunMenuPermission()
                this.permissionFormVisible = true;
            },
            savePermission() {
                let data = {
                    UserId:this.curUserId,
                    Data:this.tableDataMenu,
                }
                let data1 ={
                    jsonData: JSON.stringify(data)
                }
                request({ url: 'sys/savepermissiontree', method: 'post', data: data1 }).then((res) => {
                    if (res.code == 200) {
                        this.$message.success('保存成功');
                        this.permissionFormVisible = false
                    }
                });
            },
            // 删除操作
            handleDelete(index, row) {
                // 二次确认删除
                this.$confirm('确定要删除吗？', '提示', {
                    type: 'warning'
                })
                    .then(() => {
                        var data = {
                            id: row.Id
                        };
                        request({
                            url: 'sys/delsysuser', method: 'post',
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

                    })
                    .catch(() => {
                    });
            },
            // 编辑操作
            handleEdit(index, row) {
                this.idx = index;
                console.log(row);
                this.form.Id = row.Id;
                this.form.UserName = row.UserName;
                this.editFormVisible = true;
            },
            handleAddUser() {
                this.formAdd.Id = 0;
                this.formAdd.UserName = '';
                this.formAdd.Pwd = '';
                this.addFormVisible = true;
            },
          handleAddAgentUser() {
            this.formAddAgent.Id = 0;
            this.formAddAgent.UserName = '';
            this.formAddAgent.Pwd = '';
            this.formAddAgent.GameId = 0;
            this.addAgentFormVisible = true;
          },
            //保存 添加用户
            onAddAgentSysUser() {
                request({
                    url: 'sys/addagentsysuser', method: 'post',
                    data: this.formAddAgent
                }).then((res) => {
                    console.log(res);
                    if (res.code == 200) {
                        this.editVisible = false;
                        this.$message.success(`添加成功`);
                        this.getData();
                        this.addAgentFormVisible = false;
                    } else {
                        this.$message.error('添加失败：' + res.msg);
                    }
                });
            },
          onAddManageSysUser() {
            request({
              url: 'sys/addmanagesysuser', method: 'post',
              data: this.formAdd
            }).then((res) => {
              console.log(res);
              if (res.code == 200) {
                this.editVisible = false;
                this.$message.success(`添加成功`);
                this.getData();
                this.addFormVisible = false;
              } else {
                this.$message.error('添加失败：' + res.msg);
              }
            });
          },
            // 保存编辑
            saveEdit() {
                request({
                    url: 'sys/updatesysuser', method: 'post',
                    data: this.form
                }).then((res) => {
                    console.log(res);
                    if (res.code == 200) {
                        this.editFormVisible = false;
                        this.$message.success(`修改第 ${this.idx + 1} 行成功`);
                        this.$set(this.tableData, this.idx, this.form);
                    } else {
                        this.$message.error('更新失败：' + res.msg);
                    }
                });
            },
          saveAgentEdit() {
            request({
              url: 'sys/updateagentsysuser', method: 'post',
              data: this.editAgentForm
            }).then((res) => {
              console.log(res);
              if (res.code == 200) {
                this.editFormVisible = false;
                this.$message.success(`修改第 ${this.idx + 1} 行成功`);
                this.$set(this.tableData, this.idx, this.editAgentForm);
              } else {
                this.$message.error('更新失败：' + res.msg);
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

    .ivu-upload {
        display: none;
    }

    .ivu-btn {

    }

    .custom-tree-node {
        flex: 1;
        display: flex;
        align-items: center;
        justify-content: space-between;
        font-size: 14px;
        padding-right: 8px;
    }
</style>
