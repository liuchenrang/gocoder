<template>
    <div class="panel">
        <panel-title :title="$route.meta.title">
            <el-button @click.stop="on_refresh" size="small">
                <i class="fa fa-refresh"></i>
            </el-button>
            <router-link :to="{name: '<%.jsName%>Add'}" tag="span">
                <el-button type="primary" icon="plus" size="small">添加数据</el-button>
            </router-link>
        </panel-title>
        <div class="panel-body">
            <el-table
                    :data="table_data"
                    v-loading="load_data"
                    element-loading-text="拼命加载中"
                    border

                    style="width: 100%;">

                <%range $i,$v := .attrs%>
                     <%if ne $v.FormSettings.TYPE "IGNORE"%>
                        <el-table-column prop="<%$v.DBName%>" label="<%$v.FormSettings.COMMENT%>" width="<%$v.FormSettings.SIZE%>"></el-table-column>
                     <%end%>
                <%end%>
                <el-table-column

                        label="添加时间"
                        width="250">
                    <template slot-scope="scope">
                        <span style="margin-left: 10px">{{ scope.row.created_at | dateFormat  }}</span>
                    </template>
                </el-table-column>


                <el-table-column
                        label="操作"
                >
                    <template slot-scope="props">


                        <el-tooltip class="item" effect="dark" content="删除" placement="top-start">
                            <el-button type="danger" size="small" icon="el-icon-delete" @click="delete_data(props.row)">
                            </el-button>
                        </el-tooltip>

                        <router-link :to="{name: '<%.jsName%>Edit', params: {id: props.row.id}}" tag="span">
                            <el-button type="info" size="small" icon="edit">修改</el-button>
                        </router-link>

                    </template>
                </el-table-column>
            </el-table>
            <bottom-tool-bar>

                <div slot="page">
                    <el-pagination
                            @current-change="handleCurrentChange"
                            :current-page="currentPage"
                            :page-size="10"
                            layout="total, prev, pager, next"
                            :total="total">
                    </el-pagination>
                </div>
            </bottom-tool-bar>
        </div>
    </div>
</template>
<script type="text/javascript">
    import {bottomToolBar, panelTitle} from "components";
    import {mapGetters} from "vuex";
    import type from "store/types";
    import Api from "api";
    import EasyDate from "easydate.js";

    export default {
        data() {
            return {
                domain: "https://wx2.qq.com",
                table_data: null,
                //当前页码
                currentPage: 1,
                //数据总条目
                total: 0,
                //每页显示多少条数据
                length: 15,
                //请求时的loading效果
                load_data: true,
                //批量选择数组
                batch_select: []
            };
        },
        components: {
            panelTitle,
            bottomToolBar
        },
        created() {
            this.get_table_data();
        },
        filters: {
            dateFormat: function (value) {
                if (!value) return "";
                value = value.toString();
                return EasyDate(new Date(value)).format("yyyy-M-d H:m:s");
            }
        },
        methods: {
            //刷新
            on_refresh() {
                this.get_table_data();
            },
            ...mapGetters({
                userInfo: type.GET_USER_INFO
            }),
            //获取数据
            get_table_data() {
                let user = this.userInfo();
                this.load_data = true;
                this.indexResource(Api.<%.jsName%>Resource, {
                        page: this.currentPage,
                        length: this.length
                    })
                    .then(({data: {result, page, total}}) => {
                        this.table_data = result;
                        this.currentPage = page;
                        this.total = total;
                        this.load_data = false;
                    })
                    .finally(() => {
                        this.load_data = false;
                    });
            },

            //单个删除
            delete_data(item) {
                let user = this.userInfo();
                this.$confirm("此操作将删除该数据, 是否继续?", "提示", {
                    confirmButtonText: "确定",
                    cancelButtonText: "取消",
                    type: "warning"
                })
                    .then(() => {
                        this.load_data = true;
                        this.deleteResource(
                                Api.<%.jsName%>Resource,item.id
                            )
                            .then(resp => {
                                if(resp.status == 204){
                                    this.$message.success("删除成功！");
                                }
                                this.get_table_data();
                            })
                            .catch(response => {
                                this.$message.success("删除失败");
                            }).finally(()=>{
                                this.load_data = false;
                            });
                    })
                    .catch(() => {
                    });
            },

            //页码选择
            handleCurrentChange(val) {
                this.currentPage = val;
                this.get_table_data();
            },

        }
    };
</script>
