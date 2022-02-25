<template>
    <div>
        <div class="container">
            <div class="form-box">
                <el-form ref="form" :model="form" label-width="100px">
                    <el-form-item label="Tên của mã QR">
                        <el-input v-model="form.Alipay"></el-input>
                    </el-form-item>
                    <el-form-item label="Tên của mã QR">
                        <el-input v-model="form.AlipayName"></el-input>
                    </el-form-item>
                    <el-form-item label="Mã QR chuyển tiền Url">
                        <!--                        <el-input readonly v-model="form.AlipayQrcodeUrl"></el-input>-->
                        <el-upload
                                class="avatar-uploader"
                                action="https://jsonplaceholder.typicode.com/posts/"
                                :auto-upload="false"
                                :on-change="handleChangeAlipay"
                                :show-file-list="false">
                            <img v-if="form.AlipayQrcodeUrl" :src="form.AlipayQrcodeUrl" class="avatar">
                            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
                        </el-upload>
                        <button v-if="form.AlipayQrcodeUrl" @click="delAlipayUrl()">Xoá</button>
                    </el-form-item>
                    <el-divider></el-divider>
                    <el-form-item label="Mã QR nhận tiền Url">
                        <!--                        <el-input readonly v-model="form.WXQrcodeUrl"></el-input>-->
                        <el-upload
                                class="avatar-uploader"
                                action="https://jsonplaceholder.typicode.com/posts/"
                                :auto-upload="false"
                                :on-change="handleChangeWX"
                                :show-file-list="false">
                            <img v-if="form.WXQrcodeUrl" :src="form.WXQrcodeUrl" class="avatar">
                            <i v-else class="el-icon-plus avatar-uploader-icon"></i>

                        </el-upload>
                        <button v-if="form.WXQrcodeUrl" @click="delWXUrl()">Xoá</button>
                    </el-form-item>
                    <el-divider></el-divider>
                    <el-form-item label="Số thẻ ngân hàng">
                        <el-input v-model="form.BankCard"></el-input>
                    </el-form-item>
                    <el-form-item label="Hộ tên của bạn">
                        <el-input v-model="form.BankName"></el-input>
                    </el-form-item>
                    <el-form-item label="Địa chị lưu ngân hàng">
                        <el-input v-model="form.BankAddr"></el-input>
                    </el-form-item>
                    <el-form-item label="Họ tên">
                        <el-input v-model="form.BankUser"></el-input>
                    </el-form-item>
                    <el-form-item label="Số di động">
                        <el-input v-model="form.UserMobile"></el-input>
                    </el-form-item>

                    <el-form-item>
                        <el-button type="primary" @click="onSubmit">Nập đơn</el-button>
                        <el-button>Huỷ</el-button>
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
                    Alipay: '',
                    AlipayQrcodeUrl: '',
                    AlipayName: '',
                    WXQrcodeUrl: '',
                    BankCard: '',
                    BankName: '',
                    BankAddr: '',
                    BankUser: '',
                    UserMobile: ''
                },
                AlipayQrcodeUrlFile: {},
                WXQrcodeUrlFile: {}

            };
        },
        created() {
            this.getData();
        },
        methods: {
            delAlipayUrl() {
                this.form.AlipayQrcodeUrl = '';
                this.AlipayQrcodeUrlFile = '';
            },
            delWXUrl() {
                this.form.WXQrcodeUrl = '';
                this.WXQrcodeUrlFile = '';
            },
            handleChangeAlipay(file, fileList) {
                this.form.AlipayQrcodeUrl = URL.createObjectURL(file.raw);
                this.AlipayQrcodeUrlFile = file.raw;
            },
            handleChangeWX(file, fileList) {
                this.form.WXQrcodeUrl = URL.createObjectURL(file.raw);
                this.WXQrcodeUrlFile = file.raw;
            },
            onSubmit() {
                this.form.AlipayQrcodeUrlFile = this.AlipayQrcodeUrlFile;
                this.form.WXQrcodeUrlFile = this.WXQrcodeUrlFile;
                request({ url: 'filmbll/updatefinanceaccount', method: 'file', data: this.form }).then((res) => {
                    if (res.code == 200) {
                        this.$message.success('Gủi thành công！');
                    }
                });
            },
            getData() {
                request({ url: 'filmbll/getfinanceaccount', method: 'post' }).then((res) => {
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
        width: 178px;
        height: 178px;
        line-height: 178px;
        text-align: center;
    }

    .avatar {
        width: 178px;
        height: 178px;
        display: block;
    }
</style>
