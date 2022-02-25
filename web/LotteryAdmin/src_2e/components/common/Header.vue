<template>
    <div class="header">
        <!-- 折叠按钮 -->
        <div class="collapse-btn" @click="collapseChage">
            <i v-if="!collapse" class="el-icon-s-fold"></i>
            <i v-else class="el-icon-s-unfold"></i>
        </div>
        <div class="logo">Hệ thống quản lý</div>
        <div class="header-right">
            <div class="header-user-con">
                <!-- 全屏显示 -->
                <div class="btn-fullscreen" @click="handleFullScreen">
                    <el-tooltip effect="dark" :content="fullscreen?`取消全屏`:`全屏`" placement="bottom">
                        <i class="el-icon-rank"></i>
                    </el-tooltip>
                </div>
                <!-- 消息中心 -->
                <div class="btn-bell">
                    <el-tooltip
                            effect="dark"
                            :content="SaveDrawApply.DrawCount>0?`有${SaveDrawApply.DrawCount}条未处理下分请求`:`下分请求`"
                            placement="bottom"
                    >
                        <router-link to="/DrawMoneyApplyList">
                            <i class="el-icon-bell"></i>
                        </router-link>
                    </el-tooltip>
                    <span class="btn-bell-badge" v-if="SaveDrawApply.DrawCount>0"></span>
                </div>
                <div class="btn-bell">
                    <el-tooltip
                            effect="dark"
                            :content="SaveDrawApply.SaveCount>0?`有${SaveDrawApply.SaveCount}条未处理上分请求`:`上分请求`"
                            placement="bottom"
                    >
                        <router-link to="/SaveMoneyApplyList">
                            <i class="el-icon-bell"></i>
                        </router-link>
                    </el-tooltip>
                    <span class="btn-bell-badge" v-if="SaveDrawApply.SaveCount>0"></span>
                </div>
                <!-- 用户头像 -->
                <div class="user-avator">
                    <img src="../../assets/img/img.jpg"/>
                </div>
                <!-- 用户名下拉菜单 -->
                <el-dropdown class="user-name" trigger="click" @command="handleCommand">
                    <span class="el-dropdown-link">
                        {{username}}
                        <i class="el-icon-caret-bottom"></i>
                    </span>
                    <el-dropdown-menu slot="dropdown">
                        <el-dropdown-item divided command="loginout">退出登录</el-dropdown-item>
                    </el-dropdown-menu>
                </el-dropdown>
            </div>
        </div>
        <div hidden>
            <audio loop="loop" ref="audio" src="../../assets/mp3/b.mp3" controls="controls"></audio>
        </div>
    </div>
</template>
<script>
    import bus from '../common/bus';
    import { GetSaveDrawApply } from '../../api/film';
    // import VueAudio from './VueAudio';

    export default {
        // components:{
        //     VueAudio
        // },
        data() {
            return {
                collapse: false,
                fullscreen: false,
                name: 'linxin',
                message: 2,
                message1: 0,
                SaveDrawApply: { SaveCount: 0, DrawCount: 0, SaveId: 0, DrawId: 0 },
                CurSaveDrawApply: { SaveCount: 0, DrawCount: 0, SaveId: 0, DrawId: 0 },
                playing: 0
            };
        },
        computed: {
            username() {
                let username = localStorage.getItem('ms_username');
                return username ? username : this.name;
            }
        },
        methods: {
            // 用户名下拉菜单选择事件
            handleCommand(command) {
                if (command == 'loginout') {
                    // localStorage.removeItem('ms_username');
                    // this.$router.push('/login');

                    this.$store.dispatch('user/logout')
                        .then(() => {

                        })
                        .catch((e) => {
                            // this.$message.error('账号或密码不正确');
                            // this.$message.error(e);
                            // this.loading = false
                        });
                    this.$router.push('/login');
                }
            },
            // 侧边栏折叠
            collapseChage() {
                this.collapse = !this.collapse;
                bus.$emit('collapse', this.collapse);
            },
            // 全屏事件
            handleFullScreen() {
                let element = document.documentElement;
                if (this.fullscreen) {
                    if (document.exitFullscreen) {
                        document.exitFullscreen();
                    } else if (document.webkitCancelFullScreen) {
                        document.webkitCancelFullScreen();
                    } else if (document.mozCancelFullScreen) {
                        document.mozCancelFullScreen();
                    } else if (document.msExitFullscreen) {
                        document.msExitFullscreen();
                    }
                } else {
                    if (element.requestFullscreen) {
                        element.requestFullscreen();
                    } else if (element.webkitRequestFullScreen) {
                        element.webkitRequestFullScreen();
                    } else if (element.mozRequestFullScreen) {
                        element.mozRequestFullScreen();
                    } else if (element.msRequestFullscreen) {
                        element.msRequestFullscreen();
                    }
                }
                this.fullscreen = !this.fullscreen;
            },
            goGetSaveDrawApply() {
                GetSaveDrawApply().then((res) => {
                    if (res.code == 200) {
                        this.SaveDrawApply = res.obj;
                    }
                });

                if ((this.SaveDrawApply.DrawId != this.CurSaveDrawApply.DrawId)
                    || (this.SaveDrawApply.SaveId != this.CurSaveDrawApply.SaveId)
                    || (this.SaveDrawApply.SaveCount != this.CurSaveDrawApply.SaveCount)
                    || (this.SaveDrawApply.DrawCount != this.CurSaveDrawApply.DrawCount)
                ) {
                    if (this.playing != 2) {
                        this.playing = 2;
                        this.$refs.audio.play();
                    }
                } else {
                    if (this.playing == 2) {
                        this.playing = 1;
                        this.$refs.audio.pause();
                    }
                }
            }
        },
        mounted() {
            if (document.body.clientWidth < 1500) {
                this.collapseChage();
            }
            // if (this.timer) {
            //     clearInterval(this.timer)
            // } else {
            //     this.timer = setInterval(this.goGetSaveDrawApply,3000);
            // }
        },
        created() {
            GetSaveDrawApply().then((res) => {
                if (res.code == 200) {
                    this.SaveDrawApply = res.obj;
                }
            });
            this.goGetSaveDrawApply();
        },
        destroyed() {
            clearInterval(this.timer);
        }
    };
</script>
<style scoped>
    .header {
        position: relative;
        box-sizing: border-box;
        width: 100%;
        height: 70px;
        font-size: 22px;
        color: #fff;
    }

    .collapse-btn {
        float: left;
        padding: 0 21px;
        cursor: pointer;
        line-height: 70px;
    }

    .header .logo {
        float: left;
        width: 250px;
        line-height: 70px;
    }

    .header-right {
        float: right;
        padding-right: 50px;
    }

    .header-user-con {
        display: flex;
        height: 70px;
        align-items: center;
    }

    .btn-fullscreen {
        transform: rotate(45deg);
        margin-right: 5px;
        font-size: 24px;
    }

    .btn-bell,
    .btn-fullscreen {
        position: relative;
        width: 30px;
        height: 30px;
        text-align: center;
        border-radius: 15px;
        cursor: pointer;
    }

    .btn-bell-badge {
        position: absolute;
        right: 0;
        top: -2px;
        width: 8px;
        height: 8px;
        border-radius: 4px;
        background: #f56c6c;
        color: #fff;
    }

    .btn-bell .el-icon-bell {
        color: #fff;
    }

    .user-name {
        margin-left: 10px;
    }

    .user-avator {
        margin-left: 20px;
    }

    .user-avator img {
        display: block;
        width: 40px;
        height: 40px;
        border-radius: 50%;
    }

    .el-dropdown-link {
        color: #fff;
        cursor: pointer;
    }

    .el-dropdown-menu__item {
        text-align: center;
    }
</style>
