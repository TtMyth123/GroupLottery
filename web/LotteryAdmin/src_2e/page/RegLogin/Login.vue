<template>
    <div class="login-wrap">
        <div class="ms-login">
            <div class="ms-title">Hẹ thống quản lý tằng sau</div>
            <el-form :model="param" :rules="rules" ref="login" label-width="0px" class="ms-content">
                <el-form-item prop="username">
                    <el-input v-model="param.username" placeholder="username" ref="username">
                        <el-button slot="prepend" icon="el-icon-lx-people"></el-button>
                    </el-input>
                </el-form-item>
                <el-form-item prop="password">
                    <el-input
                            ref="password"
                            type="password"
                            placeholder="password"
                            v-model="param.password"
                            @keyup.enter.native="submitForm()"
                    >
                        <el-button slot="prepend" icon="el-icon-lx-lock"></el-button>
                    </el-input>
                </el-form-item>
                <div class="login-btn">
                    <el-button type="primary" @click="submitForm()">Đăng nhập</el-button>
                </div>
                <p class="login-tips">Tips : Tên người dùng và mật khẩu。</p>
            </el-form>
        </div>
    </div>
</template>

<script>
    import { request } from '../../utils/http';

    export default {
        data: function() {
            return {
                param: {
                    username: 'admin',
                    password: ''
                },
                rules: {
                    username: [{ required: true, message: 'Vui long nhập tên người dùng', trigger: 'blur' }],
                    password: [{ required: true, message: 'Nhập mật khẩu', trigger: 'blur' }]
                },
                redirect: undefined,
                otherQuery: {}
            };
        },
        watch: {
            $route: {
                handler: function(route) {
                    const query = route.query;
                    if (query) {
                        this.redirect = query.redirect;
                        this.otherQuery = this.getOtherQuery(query);
                    }
                },
                immediate: true
            }
        },
        mounted() {
            if (this.param.username === '') {
                this.$refs.username.focus();
            } else if (this.param.password === '') {
                this.$refs.password.focus();
            }
        },
        methods: {
            submitForm() {
                this.$refs.login.validate(valid => {
                    if (valid) {
                        this.$store.dispatch('user/login', this.param)
                            .then(() => {
                                this.$router.push({ path: this.redirect || '/', query: this.otherQuery });
                                this.$message.success('Thành công đăng nhập');
                                // this.loading = false
                            })
                            .catch((e) => {
                                // this.$message.error('Tài khoản hoặc mật khẩu không chính xác');
                                this.$message.error(e);
                                // this.loading = false
                            });

                        // request({ url:'loginreg/dologin', method:'post', data:this.param }).then((res)=>{
                        //     if (res.code==200){
                        //         console.log("ccc");
                        //         this.$message.success('Đăng nhập thành công');
                        //         localStorage.setItem("ms_username", res.obj.UserName);
                        //         localStorage.setItem("token", res.obj.CurToken);
                        //         console.log(res.obj.CurToken);
                        //         this.$router.push('/');
                        //     } else {
                        //         console.log("bbbb");
                        //         this.$message.error('Vui lòng đăng nhập tài khoản và mật khẩu');
                        //     }
                        // });
                    } else {
                        this.$message.error('Vui lòng đăng nhập tài khoản và mật khẩu');
                        console.log('error submit!!');
                        return false;
                    }
                });
            },

            getOtherQuery(query) {
                return Object.keys(query).reduce((acc, cur) => {
                    if (cur !== 'redirect') {
                        acc[cur] = query[cur];
                    }
                    return acc;
                }, {});
            }
        }
    };
</script>

<style scoped>
    .login-wrap {
        position: relative;
        width: 100%;
        height: 100%;
        background-image: url(../../assets/img/login-bg.jpg);
        background-size: 100%;
    }

    .ms-title {
        width: 100%;
        line-height: 50px;
        text-align: center;
        font-size: 20px;
        color: #fff;
        border-bottom: 1px solid #ddd;
    }

    .ms-login {
        position: absolute;
        left: 50%;
        top: 50%;
        width: 350px;
        margin: -190px 0 0 -175px;
        border-radius: 5px;
        background: rgba(255, 255, 255, 0.3);
        overflow: hidden;
    }

    .ms-content {
        padding: 30px 30px;
    }

    .login-btn {
        text-align: center;
    }

    .login-btn button {
        width: 100%;
        height: 36px;
        margin-bottom: 10px;
    }

    .login-tips {
        font-size: 12px;
        line-height: 30px;
        color: #fff;
    }
</style>