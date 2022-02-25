<template>
    <div>
        <el-row :gutter="16">
          <el-col :span="12">
            <el-card shadow="hover" class="mgb20" style="height:252px;">
              <div class="user-info">
                <img src="../../assets/img/img.jpg" class="user-avator" alt/>
                <div class="user-info-cont">
                  <div class="user-info-name">{{name}}</div>
                  <div>{{role}}</div>
                </div>
              </div>
              <div class="user-info-list">
                Thời gian đăng nhập lần trước：
                <span>2019-11-01</span>
              </div>
              <div class="user-info-list">
                Địa điểm đăng nhập lần trước：
                <span>东莞</span>
              </div>
            </el-card>
          </el-col>

          <el-col :span="12">
            <el-card shadow="hover" class="mgb20" style="height:252px;">
              <el-button type="primary"   @click="Get7DayBetYieldChart">Làm mới</el-button>
              <schart class="wrapper" canvasId="myCanvas" :options="options"></schart>
            </el-card>
          </el-col>
        </el-row>
    </div>
</template>

<script>
    import Schart from 'vue-schart';
    import { request } from '@/utils/http';
    import resize from '../../components/Charts/mixins/resize';

    export default {
        name: 'home',
        data() {
            return {
              mixins: [resize],
              name: localStorage.getItem('ms_username'),
              options: {
                type: 'bar',
                title: {
                  text: 'Một tuần vừa rồi thành tích'
                },
                labels: ['Thứ', 'Thứ', 'Thứ4', 'Thứ5', 'Thứ 6'],
                datasets: [
                  {
                    label: 'Đồ điện',
                    data: [234, 278, 270, 190, 230]
                  },
                  {
                    label: 'Bách hoá',
                    data: [164, 178, 190, 135, 160]
                  },
                  {
                    label: 'Đồ ăn',
                    data: [144, 198, 150, 235, 120]
                  }
                ]
              }
            };
        },
        components: {
            Schart
        },
        computed: {
            role() {
                return this.name === 'admin' ? 'Người quản lý siêu cấp' : 'Người dùng phổ thông';
            }
        },
        mounted() {

        },
        created() {
             this.Get7DayBetYieldChart();
        },
        destroyed() {
            // clearInterval(this.timer);
        },
        methods: {
          Get7DayBetYieldChart() {
            request({ url: 'api/get7daybetyieldchart', method: 'post', data: {} }).then((res) => {
              if (res.code == 200) {
                console.log(res.obj);
                this.options= res.obj;
              }
            });
          },
            // changeDate() {
            //     const now = new Date().getTime();
            //     this.data.forEach((item, index) => {
            //         const date = new Date(now - (6 - index) * 86400000);
            //         item.name = `${date.getFullYear()}/${date.getMonth() + 1}/${date.getDate()}`;
            //     });
            // }
            // handleListener() {
            //     bus.$on('collapse', this.handleBus);
            //     // 调用renderChart方法对图表进行重新渲染
            //     window.addEventListener('resize', this.renderChart);
            // },
            // handleBus(msg) {
            //     setTimeout(() => {
            //         this.renderChart();
            //     }, 200);
            // },
            // renderChart() {
            //     this.$refs.bar.renderChart();
            //     this.$refs.line.renderChart();
            // }
        }
    };
</script>


<style scoped>
    .el-row {
        margin-bottom: 20px;
    }

    .grid-content {
        display: flex;
        align-items: center;
        height: 100px;
    }

    .grid-cont-right {
        flex: 1;
        text-align: center;
        font-size: 14px;
        color: #999;
    }

    .grid-num {
        font-size: 30px;
        font-weight: bold;
    }

    .grid-con-icon {
        font-size: 50px;
        width: 100px;
        height: 100px;
        text-align: center;
        line-height: 100px;
        color: #fff;
    }

    .grid-con-1 .grid-con-icon {
        background: rgb(45, 140, 240);
    }

    .grid-con-1 .grid-num {
        color: rgb(45, 140, 240);
    }

    .grid-con-2 .grid-con-icon {
        background: rgb(100, 213, 114);
    }

    .grid-con-2 .grid-num {
        color: rgb(45, 140, 240);
    }

    .grid-con-3 .grid-con-icon {
        background: rgb(242, 94, 67);
    }

    .grid-con-3 .grid-num {
        color: rgb(242, 94, 67);
    }

    .user-info {
        display: flex;
        align-items: center;
        padding-bottom: 20px;
        border-bottom: 2px solid #ccc;
        margin-bottom: 20px;
    }

    .user-avator {
        width: 120px;
        height: 120px;
        border-radius: 50%;
    }

    .user-info-cont {
        padding-left: 50px;
        flex: 1;
        font-size: 14px;
        color: #999;
    }

    .user-info-cont div:first-child {
        font-size: 30px;
        color: #222;
    }

    .user-info-list {
        font-size: 14px;
        color: #999;
        line-height: 25px;
    }

    .user-info-list span {
        margin-left: 70px;
    }

    .mgb20 {
        margin-bottom: 20px;
    }

    .todo-item {
        font-size: 14px;
    }

    .todo-item-del {
        text-decoration: line-through;
        color: #999;
    }

    .schart {
        width: 100%;
        height: 300px;
    }
</style>
