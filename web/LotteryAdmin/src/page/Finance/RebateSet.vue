<template>
    <div>
        <div class="container">
            <div class="form-box">
                <el-form ref="form" :model="form" label-width="100px">
                  <el-form-item label="返佣层级">
                    <el-input-number class="loose-input" v-model="form.Level" :step="1" :min="0"  :max="4"  > </el-input-number>
                  </el-form-item>
                  <el-form-item label="上1级投注佣金比例">
                    <el-input-number class="loose-input" v-model="form.BetRebateRatio" :step="0.1" :min="0"  :precision="4"  > </el-input-number>
                  </el-form-item>
                  <el-form-item label="上2级投注佣金比例">
                    <el-input-number class="loose-input" v-model="form.BetRebateRatio1" :step="0.1" :min="0"  :precision="4"  > </el-input-number>
                  </el-form-item>
                  <el-form-item label="上3级投注佣金比例">
                    <el-input-number class="loose-input" v-model="form.BetRebateRatio2" :step="0.1" :min="0"  :precision="4"  > </el-input-number>
                  </el-form-item>
                  <el-form-item label="上4级投注佣金比例">
                    <el-input-number class="loose-input" v-model="form.BetRebateRatio3" :step="0.1" :min="0"  :precision="4"  > </el-input-number>
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
        name: 'RebateSet',
        data() {
            return {
                form: {
                  Id: 0,
                  Level:0,
                  BetRebateRatio: 0.0,
                  BetRebateRatio1: 0.0,
                  BetRebateRatio2: 0.0,
                  BetRebateRatio3: 0.0,
                }
            };
        },
        created() {
            this.getData();
        },
        methods: {
            onSubmit() {
                request({ url: 'api/savettrebateset', method: 'post', data: this.form }).then((res) => {
                    if (res.code == 200) {
                        this.$message.success('提交成功！');
                    }
                });
            },
            getData() {
                request({ url: 'api/getttrebateset', method: 'post' }).then((res) => {
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
