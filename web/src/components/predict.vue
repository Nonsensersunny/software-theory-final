<template>
    <div id="predict">
        <el-steps :active="activeStep">
            <el-step title="选取训练集" icon="el-icon-edit"></el-step>
            <el-step title="选取测试集" icon="el-icon-plus"></el-step>
            <el-step title="选取算法" icon="el-icon-edit"></el-step>
            <el-step title="测试" icon="el-icon-edit"></el-step>
        </el-steps>
        <div id="operation-card">
            <el-select v-model="trainId" placeholder="选择训练集" v-if="activeStep === 1">
                <el-option
                        v-for="train in trainDatasets"
                        :key="train.id"
                        :label="train.name"
                        :value="train.id"></el-option>
            </el-select>
            <el-button @click="activeStep += 1" :disabled="trainId === ''" v-if="activeStep === 1">确定</el-button>
            <el-select v-model="testId" placeholder="选择测试集" v-if="activeStep === 2">
                <el-option
                        v-for="test in testDatasets"
                        :key="test.id"
                        :label="test.name"
                        :value="test.id"></el-option>
            </el-select>
            <el-button @click="activeStep += 1" :disabled="testId === ''" v-if="activeStep === 2">确定</el-button>
            <el-select v-model="algorithmId" placeholder="选择算法" v-if="activeStep === 3">
                <el-option
                        v-for="algo in algorithms"
                        :key="algo.id"
                        :label="algo.name"
                        :value="algo.id"></el-option>
            </el-select>
            <el-button @click="activeStep += 1" :disabled="algorithmId === ''" v-if="activeStep === 3">确定</el-button>
            <el-button @click="activeStep -= 1" v-if="activeStep > 1">上一步</el-button>
            <el-button @click="createPrediction" v-if="activeStep === 4">开始预测</el-button>
        </div>
        <div id="display-board">
            <el-card v-if="trainId !== '' && activeStep >= 1">
                <div slot="header">训练集</div>
                <div>
                    <p>名称：{{ getDataById(trainDatasets, trainId).name }}</p>
                    <p>类型：{{ getDataById(trainDatasets, trainId).type | datasetTypeFilter }}</p>
                    <p>上传时间：{{ getDataById(trainDatasets, trainId).time | datetimeFilter }}</p>
                </div>
            </el-card>
            <el-card v-if="testId !== '' && activeStep >= 2">
                <div slot="header">测试集</div>
                <div>
                    <p>名称：{{ getDataById(testDatasets, testId).name }}</p>
                    <p>类型：{{ getDataById(testDatasets, testId).type | datasetTypeFilter }}</p>
                    <p>上传时间：{{ getDataById(testDatasets, testId).time | datetimeFilter }}</p>
                </div>
            </el-card>
            <el-card v-if="algorithmId !== '' && activeStep >= 3">
                <div slot="header">算法</div>
                <p>名称：{{ getDataById(algorithms, algorithmId).name }}</p>
                <p>描述：{{ getDataById(algorithms, algorithmId).description }}</p>
            </el-card>
        </div>
    </div>
</template>

<script>
    export default {
        name: "predict",
        filters: {
            datasetTypeFilter(dataset) {
                if (dataset.type === 'test') {
                    return '测试集';
                } else if (dataset.type === 'train') {
                    return '训练集';
                } else {
                    return '位置类型';
                }
            },
            datetimeFilter(date) {
                return new Date(date).toLocaleString("zh");
            }
        },
        data() {
            return {
                algorithms: [],
                activeStep: 1,
                trainId: '',
                testId: '',
                algorithmId: '',
                trainDatasets: this.datasets.filter((dt) => {
                    return dt.type === 'train'
                }),
                testDatasets: this.datasets.filter((dt) => {
                    return dt.type === 'test'
                }),
            }
        },
        methods: {
            async getAlgorithms() {
                try {
                    let resp = await this.$store.dispatch("getAlgorithms");
                    if (resp.code !== 0) {
                        this.$message.error("无法获取算法信息");
                    } else {
                        this.algorithms = resp.data.algorithms;
                    }
                } catch (e) {
                    console.log(e)
                }
            },
            async createPrediction() {
                try {
                    let trainId = this.trainId;
                    let testId = this.testId;
                    let algorithmId = this.algorithmId;
                    let resp = await this.$store.dispatch("createPrediction", {trainId, testId, algorithmId});
                    if (resp.code === 0) {
                        this.$message.success('预测完成');
                        this.$emit("predict-success");
                    } else {
                        this.$message.error("预测失败");
                    }
                } catch (e) {
                    console.log(e)
                }
            },
            getDataById(data, id) {
                return data.find(d => d.id === id);
            }
        },
        created() {
            this.getAlgorithms();
        },
        props: {
            datasets: {
                type: Array,
                required: true,
            }
        }
    }
</script>

<style scoped>

</style>