<template>
    <div class="sidebar">
        <el-menu
                class="sidebar-el-menu"
                :default-active="onRoutes"
                :collapse="collapse"
                background-color="#324157"
                text-color="#bfcbd9"
                active-text-color="#20a0ff"
                :unique-opened="false"
                @select="handleOpen"
                router
        >
            <!--            :collapse="collapse"-->
            <template v-for="item in items">
                <template v-if="item.children">
                    <el-submenu :index="item.path" :key="item.Id">
                        <template slot="title">
                            <i :class="item.icon"></i>
                            <span slot="title">{{ item.title }}</span>
                        </template>
                        <template v-for="subItem in item.children">
                            <el-submenu
                                    v-if="subItem.children"
                                    :index="subItem.path"
                                    :key="subItem.Id"
                            >
                                <template slot="title">{{ subItem.title }}</template>
                                <el-menu-item
                                        v-for="(threeItem,i) in subItem.children"
                                        :key="threeItem.Id"
                                        :index="threeItem.path"

                                >{{ threeItem.title }}
                                </el-menu-item>
                            </el-submenu>
                            <el-menu-item
                                    v-else
                                    :index="subItem.path"
                                    :key="subItem.Id"

                            >{{ subItem.title }}
                            </el-menu-item>
                        </template>
                    </el-submenu>
                </template>
                <template v-else>
                    <el-menu-item :index="item.path" :key="item.Id">
                        <i :class="item.icon"></i>
                        <span slot="title">{{ item.title }}</span>
                    </el-menu-item>
                </template>
            </template>
        </el-menu>
    </div>
</template>

<script>
    import bus from '../common/bus';
    import { request } from '../../utils/http';

    export default {
        data() {
            return {
                collapse: false,
                items: [],
                curTitle: ''
            };
        },
        computed: {
            onRoutes() {
                // this.$route.meta.title = this.curTitle;
                return this.$route.path.replace('/', '');
            }
        },
        methods: {
            handleOpen(key, keyPath) {

                console.log(key, keyPath);
            },
            handleClose(key, keyPath) {
                console.log(key, keyPath);
            }
        },
        created() {
            request({ url: 'sys/getmenulist', method: 'get' }).then((res) => {
                console.log(res);
                if (res.code == 200) {
                    this.items = res.obj;
                }
            });
            // 通过 Event Bus 进行组件间通信，来折叠侧边栏
            bus.$on('collapse', msg => {
                this.collapse = msg;
                bus.$emit('collapse-content', msg);
            });
        }
    };
</script>

<style scoped>
    .sidebar {
        display: block;
        position: absolute;
        left: 0;
        top: 70px;
        bottom: 0;
        overflow-y: scroll;
    }

    .sidebar::-webkit-scrollbar {
        width: 0;
    }

    .sidebar-el-menu:not(.el-menu--collapse) {
        width: 250px;
    }

    .sidebar > ul {
        height: 100%;
    }
</style>
