<template>
    <div>
        <div class="container">
            <div class="form-box">
                <el-form ref="form" :model="form" label-width="100px">
                  <el-form-item label="收款方式" prop="PayWay">
                    <el-select v-model.number="form.PayWay" placeholder="请收款方式">
                      <el-option v-for="item in PayWays" :label="item.name" :key="item.value"
                                 :value="item.value"></el-option>
                    </el-select>
                  </el-form-item>
                    <el-form-item label="支付宝号">
                        <el-input v-model="form.OnlinePay"></el-input>
                    </el-form-item>
                    <el-form-item label="支付宝用户名">
                        <el-input v-model="form.AlipayName"></el-input>
                    </el-form-item>
<!--                    <el-form-item label="支付宝收款码">-->
<!--                        &lt;!&ndash;                        <el-input readonly v-model="form.AlipayQrcodeUrl"></el-input>&ndash;&gt;-->
<!--                        <el-upload-->
<!--                                class="avatar-uploader"-->
<!--                                action="https://jsonplaceholder.typicode.com/posts/"-->
<!--                                :auto-upload="false"-->
<!--                                :on-change="handleChangeAlipay"-->
<!--                                :show-file-list="false">-->
<!--                            <img v-if="form.AlipayQrcodeUrl" :src="form.AlipayQrcodeUrl" class="avatar">-->
<!--                            <i v-else class="el-icon-plus avatar-uploader-icon"></i>-->
<!--                        </el-upload>-->
<!--                        <button v-if="form.AlipayQrcodeUrl" @click="delAlipayUrl()">删除</button>-->
<!--                    </el-form-item>-->

                  <el-form-item label="支付宝收款码" prop="AlipayUrl">
                    <el-upload
                        class="avatar-uploader"
                        action="kit/uploadpic"
                        :data="AlipayQrcodeUrlFileData"
                        accept="image/jpeg,image/gif,image/png"
                        :before-upload="beforeAvatarUploadAlipayQrcode"
                        :on-change="handlePicUrlAlipayQrcode"
                        :on-success="uploadSuccessAlipayQrcode"
                        :show-file-list="false" >
                      <img v-if="form.AlipayUrl" :src="form.AlipayUrl" class="avatar">
                      <i v-else class="el-icon-plus avatar-uploader-icon"></i>
                    </el-upload>
                  </el-form-item>
                    <el-divider></el-divider>
                  <el-form-item label="微信收款码" prop="WXReceiptUrl">
                    <el-upload
                        class="avatar-uploader"
                        action="kit/uploadpic"
                        :data="WXQrcodeUrlFileData"
                        accept="image/jpeg,image/gif,image/png"
                        :before-upload="beforeAvatarUploadWXQrcode"
                        :on-change="handlePicUrlWXQrcode"
                        :on-success="uploadSuccessWXQrcode"
                        :show-file-list="false" >
                      <img v-if="form.WXReceiptUrl" :src="form.WXReceiptUrl" class="avatar">
                      <i v-else class="el-icon-plus avatar-uploader-icon"></i>
                    </el-upload>
                  </el-form-item>
                    <el-divider></el-divider>
                    <el-form-item label="银行卡号">
                        <el-input v-model="form.BankCard"></el-input>
                    </el-form-item>
                    <el-form-item label="开户行名">
                        <el-input v-model="form.BankName"></el-input>
                    </el-form-item>
                    <el-form-item label="开户地址">
                        <el-input v-model="form.BankAddr"></el-input>
                    </el-form-item>
                    <el-form-item label="姓名">
                        <el-input v-model="form.BankUser"></el-input>
                    </el-form-item>
                    <el-form-item label="手机号">
                        <el-input v-model="form.UserMobile"></el-input>
                    </el-form-item>

                    <el-form-item>
                        <el-button type="primary" @click="onSubmit">表单提交</el-button>
                        <el-button>取消</el-button>
                    </el-form-item>
                </el-form>
            </div>
        </div>
    </div>
</template>

<script>
    import { request } from '../../utils/http';

    export default {
        name: 'FinanceAccount',
        data() {
            return {
                form: {
                  Id: 0,
                  PayWay: 0,

                  AlipayUrl: '',
                  AlipayName: '',
                  OnlinePay:'',

                  WXReceiptUrl: '',
                  AlipayQrcodePicGuid:'',
                  WXQrcodePicGuid:'',

                  BankCard: '',
                  BankName: '',
                  BankAddr: '',
                  BankUser: '',
                  UserMobile: ''
                },
              PayWays: [
                { name: '微信', value: 0 },
                { name: '银行', value: 1 },
                { name: '支付宝', value: 2 }
              ],
              AlipayQrcodePicGuid:'',
              WXQrcodePicGuid:'',
              AlipayQrcodeUrlFileData:{picGuid:''},
              WXQrcodeUrlFileData:{picGuid:''},
                AlipayQrcodeUrlFile: {},
                WXQrcodeUrlFile: {}

            };
        },

        created() {
          this.getGuid();
            this.getData();
        },
        methods: {
          getGuid(){
            request({
              url:'kit/getguid', method:'post',
            }).then((res)=>{
              if (res.code==200){
                this.AlipayQrcodePicGuid = res.obj;
              } else {
                this.$message.error('获取数据失败：'+res.msg);
              }
            });
            request({
              url:'kit/getguid', method:'post',
            }).then((res)=>{
              if (res.code==200){
                this.WXQrcodePicGuid = res.obj;
              } else {
                this.$message.error('获取数据失败：'+res.msg);
              }
            });
          },
          uploadSuccessAlipayQrcode(res,file, fileList){
            console.log("ok" );
          },
          handlePicUrlAlipayQrcode(file, fileList) {
            console.log("handlePicUrlAlipayQrcode" );
            this.form.AlipayUrl = URL.createObjectURL(file.raw);
            console.log(this.form.AlipayUrl);
          },
          beforeAvatarUploadAlipayQrcode(file) {
            const isIMAGE = file.type === 'image/jpeg'||'image/gif'||'image/png';
            if (!isIMAGE) {
              this.$message.error('上传文件只能是图片格式!');
            }
            const isLt2M = file.size / 1024  < 500;
            if (!isLt2M) {
              this.$message.error('上传文件大小不能超过 500bk!');
            }
            if (isLt2M&& isIMAGE) {
              this.AlipayQrcodeUrlFileData.picGuid = this.AlipayQrcodePicGuid;
            }

            return isLt2M&& isIMAGE;
          },
          uploadSuccessWXQrcode(res,file, fileList){
            console.log("ok" );
          },
          handlePicUrlWXQrcode(file, fileList) {
            console.log("handlePicUrlWXQrcode" );
            this.form.WXReceiptUrl = URL.createObjectURL(file.raw);
            console.log(this.form.WXReceiptUrl );
          },
          beforeAvatarUploadWXQrcode(file) {
            const isIMAGE = file.type === 'image/jpeg'||'image/gif'||'image/png';
            if (!isIMAGE) {
              this.$message.error('上传文件只能是图片格式!');
            }
            const isLt2M = file.size / 1024  < 500;
            if (!isLt2M) {
              this.$message.error('上传文件大小不能超过 500bk!');
            }
            if (isLt2M&& isIMAGE) {
              this.WXQrcodeUrlFileData.picGuid = this.WXQrcodePicGuid;
            }

            return isLt2M&& isIMAGE;
          },
          onSubmit() {
            this.form.AlipayQrcodePicGuid = this.AlipayQrcodeUrlFileData.picGuid;
            this.form.WXQrcodePicGuid = this.WXQrcodeUrlFileData.picGuid;
            request({ url: 'api/updatefinanceaccount', method: 'file', data: this.form }).then((res) => {
              if (res.code == 200) {
                this.$message.success('提交成功！');
              }
            });
          },
            getData() {
                request({ url: 'api/getfinanceaccount', method: 'post' }).then((res) => {
                    if (res.code == 200) {
                        this.form = res.obj;
                    }
                });
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
</style>
