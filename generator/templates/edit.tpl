<template>
    <div class="panel">
        <panel-title :title="$route.meta.title"></panel-title>
        <div class="panel-body"
             v-loading="load_data"
             element-loading-text="拼命加载中">
            <el-row>
                <el-col :span="8">
                    <el-form ref="form" :model="form" :rules="rules" label-width="100px">

                                     <%range $i,$v := .attrs%>
                                            <%if ne $v.FormSettings.TYPE "IGNORE"%>
                                                <% $v.Render.RenderElement%>
                                            <%end%>
                                    <%end%>




                        <el-form-item>
                            <el-button type="primary" @click="on_submit_form" :loading="on_submit_loading">立即提交
                            </el-button>
                            <el-button @click="$router.back()">取消</el-button>
                        </el-form-item>
                    </el-form>
                </el-col>
            </el-row>
        </div>
    </div>
</template>
<script type="text/javascript">
    import {panelTitle} from "components";
    import Api from "../../api";

    export default {
        data() {
            return {
                form: {
                        <%range $i,$v := .fields%>
                         <%$v.DBName|js%>: "",
                       <%end%>
                },
                 <%range $i,$v := .attrs%>
                    <%if ne $v.FormSettings.TYPE "IGNORE"%>
                        <% $v.Render.RenderData%>
                    <%end%>
                <%end%>
                route_id: this.$route.params.id,
                load_data: false,
                on_submit_loading: false,
                rules: {}
            };
        },
        created() {
            this.route_id && this.queryResource();
             <%range $i,$v := .attrs%>
                <%if ne $v.FormSettings.TYPE "IGNORE"%>
                    <% $v.Render.RenderInit%>
                <%end%>
            <%end%>
        },
        methods: {
            //获取数据
            queryResource() {
                this.load_data = true;
                this.getResource(Api.<%.jsName%>Resource,this.$route.params.id)
                    .then(response => {
                        this.form = response.data.data;
                        this.load_data = false;
                    })
                    .catch(() => {
                    })
                    .finally(()=>{
                        this.load_data = false;
                    });
            },
            <%range $i,$v := .attrs%>
                <%if ne $v.FormSettings.TYPE "IGNORE"%>
                    <% $v.Render.RenderMethod%>
                <%end%>
            <%end%>

            updateResource(){
                    this.putResource(
                        Api.<%.jsName%>Resource,
                        this.$route.params.id,
                        this.form
                    )
                    .then(response => {
                        if (response.status == 204) {
                            this.$message.success("修改成功!");
                        } else {
                            this.$message.success(response.data.message);
                        }
                        setTimeout(this.$router.back(), 500);
                    })
                    .catch(() => {
                    })
                    .finally(()=>{
                       this.on_submit_loading = false;
                    });
            },
            createResource(){
                this.postResource(
                        Api.<%.jsName%>Resource,
                        this.form
                    )
                    .then(response => {
                        if (response.status == 201) {
                            this.$message.success("创建成功!");
                        } else {
                            this.$message.success(response.data.message);
                        }
                        setTimeout(this.$router.push({name: "<%.jsName%>", params: {}}), 500);
                    }).finally(()=>{
                         this.on_submit_loading = false;
                    });;
            },
            //提交
            on_submit_form() {
                let that = this;
                this.$refs.form.validate(valid => {
                    if (!valid) return false;
                    that.on_submit_loading = true;
                    if (this.route_id) {
                        that.updateResource()
                    } else {
                        that.createResource()
                    }
                });
            }
        },
        components: {
            panelTitle
        }
    };
</script>
