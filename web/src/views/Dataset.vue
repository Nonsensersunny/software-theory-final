<template>
    <div id="dataset">
        <el-table
            :data="datasets"
            stripe
            style="width: 100%">
            <el-table-column label="上传日期" width="200">
                <template slot-scope="scope">
                    <i class="el-icon-time"></i>
                    <span style="margin-left: 10px">{{ scope.row.time | date }}</span>
                </template>
            </el-table-column>
            <el-table-column label="名称" width="150">
                <template slot-scope="scope">
                    <p>{{ scope.row.name | contentFilter }}</p>
                </template>
            </el-table-column>
            <el-table-column label="类型" width="80">
                <template slot-scope="scope">
                    <el-tag size="mini">{{ scope.row.type | datasetType }}</el-tag>
                </template>
            </el-table-column>
            <el-table-column label="描述" width="200">
                <template slot-scope="scope">
                    <p>{{ scope.row.description | contentFilter }}</p>
                </template>
            </el-table-column>
            <el-table-column width="150">
                <template slot="header" slot-scope="scope">
                    <el-button size="mini" type="primary" @click="uploadDatasetVisible = true"><i class="el-icon-plus"></i></el-button>
                    <el-button size="mini" type="success" @click="predictVisible = true">预测</el-button>
                </template>
                <template slot-scope="scope">
                    <el-button size="mini" type="danger">删除</el-button>
                    <el-button size="mini" type="info">历史</el-button>
                </template>
            </el-table-column>
        </el-table>
        <el-dialog title="上传数据集" :visible.sync="uploadDatasetVisible">
            <el-form :model="dataset">
                <el-form-item label="Name"><el-input v-model="dataset.name" /></el-form-item>
                <el-form-item label="Type">
                    <el-select v-model="dataset.type" placeholder="Select a type">
                        <el-option label="测试集" value="test"></el-option>
                        <el-option label="训练集" value="train"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="Description"><el-input v-model="dataset.description" /></el-form-item>
                <el-form-item label="数据集">
                    <el-upload
                            action="''"
                            :http-request="catchFile"
                            :show-file-list="true"
                            :on-success="handleFileSuccess"
                            :on-progress="handleFileProgress"
                            :before-upload="beforeFileUpload">
                        <i class="el-icon-plus avatar-uploader-icon"></i>
                    </el-upload>
                </el-form-item>
            </el-form>
            <el-button @click="cancelChange">取消</el-button>
            <el-button type="primary" @click="createDataset">确定</el-button>
        </el-dialog>
        <el-dialog title="预测" :visible.sync="predictVisible" width="600">
            <Predict :datasets="datasets" @predict-success="predictSuccess" />
        </el-dialog>
    </div>
</template>

<script>
    import {Dataset} from "@/assets/js/type";
    // import Predict from "@/components/predict.vue";
    import Predict from "../components/predict";

    export default {
        name: "Dataset",
        components: {Predict},
        filters: {
            date(val) {
                return new Date(val).toLocaleString("zh");
            },
            datasetType(cat) {
                if (cat === 'test') {
                    return "测试集";
                } else if (cat === 'train') {
                    return '训练集';
                } else {
                    return '未知类型';
                }
            },
            contentFilter(content) {
                if (content === 'undefined') {
                    return "空";
                }
            }
        },
        data() {
            return {
                datasets: [],
                dataset: new Dataset(),
                datasetFile: '',
                editable: false,
                uploadDatasetVisible: false,
                predictVisible: false,
            }
        },
        methods: {
            cancelChange() {
                this.editable = false;
                window.URL.revokeObjectURL(this.dataset.upload);
                this.datasetFile = '';
                this.uploadDatasetVisible = false;
            },
            catchFile(file) {
                this.datasetFile = file.file;
                let url = window.URL.createObjectURL(file.file);
                this.dataset.upload = url;
            },
            handleFileProgress(ev, rawFile) {
                console.log(URL.createObjectURL(rawFile.raw))
            },
            handleFileSuccess(res, file) {
                this.dataset.upload = URL.createObjectURL(file.raw);
                console.log(URL.createObjectURL(file.raw));
                console.log(this.dataset.upload);
            },
            beforeFileUpload(file) {
                const isCSV1 = file.type === 'text/csv';
                const isCSV2 = file.type === 'text/comma-separated-values';
                const isCSV3 = file.type === '.csv'
                const isCSV4 = file.type === 'application/csv';
                const isLt10M = file.size / 1024 / 1024 < 10;
                const isCSV = isCSV1 || isCSV2 || isCSV3 || isCSV4;

                if (!isCSV1 && !isCSV2 && !isCSV3 && !isCSV4) {
                    this.$message.error("文件类型错误");
                }
                if (!isLt10M) {
                    this.$message.error("文件过大");
                }
                return isCSV && isLt10M;
            },
            async getDatasets() {
                try {
                    let resp = await this.$store.dispatch("getDatasets");
                    if (resp.code !== 0) {
                        this.$notify({
                            title: "Notification",
                            message: resp,
                            duration: 3000
                        })
                    } else {
                        this.datasets = resp.data.datasets;
                    }
                } catch (e) {
                    this.$notify({
                        title: "Notification",
                        message: e,
                        duration: 3000
                    })
                }
            },
            async createDataset() {
                try {
                    this.dataset.upload = this.datasetFile;
                    let resp = await this.$store.dispatch("createDataset", this.dataset);
                    if (resp.code === 0) {
                        this.$notify({
                            title: "提示",
                            type: "success",
                            message: "上传成功"
                        });
                        this.uploadDatasetVisible = false;
                    } else {
                        this.$notify({
                            title: "提示",
                            type: "error",
                            message: "上传失败"
                        });
                    }
                } catch (e) {
                    console.log(e)
                }
            },
            predictSuccess() {
                this.predictVisible = false;
            }
        },
        created() {
            this.getDatasets();
        },
        comments: {
            Predict
        }
    }
</script>

<style scoped>

</style>