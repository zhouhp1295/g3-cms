<template>
  <div id="app">
    <img src="./logo.png">
    <h3>欢迎使用 g3-cms</h3>
    <el-steps :active="activeStep" align-center>
      <el-step title="步骤1" description="数据库设置"></el-step>
      <el-step title="步骤2" description="超级管理员设置"></el-step>
      <el-step title="步骤3" description="安装"></el-step>
    </el-steps>
    <div class="form-body">
      <div class="form-steps">
        <div v-if="activeStep == 1">
          <el-form ref="step1form" :model="form" label-width="120px" :rules="step1rules" style="text-align: left;">
            <div class="form-step-title">数据库设置</div>
            <el-form-item label="数据库类型" style="text-align: left;" prop="databaseType" required>
              <el-select v-model="form.databaseType" placeholder="请选择">
                <el-option
                    v-for="item in databaseOptions"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value">
                </el-option>
              </el-select>
            </el-form-item>
            <div v-if="form.databaseType != 'sqlite3'">
              <el-form-item label="数据库主机" prop="databaseHost" required>
                <el-input v-model="form.databaseHost"></el-input>
              </el-form-item>
              <el-form-item label="数据库用户" prop="databaseUser" required>
                <el-input v-model="form.databaseUser"></el-input>
              </el-form-item>
              <el-form-item label="数据库密码" prop="databasePwd" required>
                <el-input v-model="form.databasePwd"></el-input>
              </el-form-item>
            </div>
            <el-form-item label="数据库名称" prop="databaseName" required>
              <el-input v-model="form.databaseName"></el-input>
            </el-form-item>
<!--            <el-form-item label="安装测试数据" prop="installDemo">-->
<!--              <el-checkbox v-model="form.installDemo">是</el-checkbox>-->
<!--            </el-form-item>-->
          </el-form>
        </div>
        <div v-if="activeStep == 2">
          <el-form ref="step2form" :model="form" label-width="120px" :rules="step2rules">
            <div class="form-step-title">超级管理员设置</div>
            <el-form-item label="账户" prop="username" required>
              <el-input v-model="form.username"></el-input>
            </el-form-item>
            <el-form-item label="密码" prop="password" required>
              <el-input v-model="form.password"></el-input>
            </el-form-item>
          </el-form>
        </div>
        <div v-if="activeStep == maxStep">
        </div>
      </div>
      <div class="form-footer">
        <el-button type="warn" @click="handleTestConn"  v-loading.fullscreen.lock="fullscreenLoading" v-if="activeStep == 1">测试连接</el-button>
        <el-button type="primary" @click="handlePre" v-if="activeStep != 1">上一步</el-button>
        <el-button type="primary" @click="handleNext" v-if="activeStep < maxStep-1">下一步</el-button>
        <el-button type="primary" @click="handleSubmit" v-loading.fullscreen.lock="fullscreenLoading" v-if="activeStep == maxStep-1">开始安装</el-button>
      </div>
    </div>

  </div>
</template>

<script>
import axios from 'axios'

axios.defaults.headers['Content-Type'] = 'application/json;charset=utf-8'
// 创建axios实例
const request = axios.create({
  // axios中请求配置有baseURL选项，表示请求URL公共部分
  baseURL: process.env.VUE_APP_BASE_API,
  // 超时
  timeout: 10000
})

export default {
  name: 'App',
  components: {},
  data() {
    return {
      activeStep: 1,
      maxStep: 3,
      fullscreenLoading: false,
      databaseOptions: [
        {
          "label": "MySQL",
          "value": "mysql",
        }
      ],
      form: {
        databaseType: "mysql",
        databaseHost: undefined,
        databaseUser: undefined,
        databasePwd: undefined,
        databaseName: undefined,
        username: undefined,
        password: undefined,
        installDemo: true,
      },
      step1rules: {
        databaseType: [
          { required: true, message: '请选择数据库类型', trigger: 'blur' }
        ],
        databaseHost: [
          { required: true, message: '请输入数据库主机', trigger: 'blur' }
        ],
        databaseUser: [
          { required: true, message: '请输入数据库用户', trigger: 'blur' }
        ],
        databasePwd: [
          { required: true, message: '请输入数据库密码', trigger: 'blur' }
        ],
        databaseName: [
          { required: true, message: '请输入数据库名称', trigger: 'blur' }
        ],
      },
      step2rules: {
        username: [
          { required: true, message: '请输入管理员账户', trigger: 'blur' },
          { min: 4, max: 20, message: '长度在 4 到 20 个字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入管理员账户', trigger: 'blur' },
          { min: 6, max: 30, message: '长度在 6 到 30 个字符', trigger: 'blur' }
        ],
      },
    }
  },
  methods: {
    handlePre() {
      this.activeStep --;
    },
    handleNext() {
      if (this.activeStep === 1) {
        this.handleStep1FormSubmit(true);
      } else {
        this.activeStep ++;
      }
    },
    handleStep1FormSubmit(next) {
      this.$refs.step1form.validate((valid) => {
        if (valid) {
          this.fullscreenLoading = true;
          request({
            url: '/api/common/testConn',
            method: 'post',
            data: this.form,
          }).then(res => {
            if (res.data.code === 200) {
              if (next) {
                this.activeStep ++;
              } else {
                this.$message({
                  message: res.data.msg,
                  type: 'success'
                });
              }
            } else {
              this.$message({
                message: res.data.msg,
                type: 'error'
              });
            }
            this.fullscreenLoading = false;
          })
        } else {
          return false;
        }
      });
    },
    handleStep2FormSubmit() {
      this.$refs.step2form.validate((valid) => {
        if (valid) {
          this.fullscreenLoading = true;
          //提交
          request({
            url: '/api/common/install',
            method: 'post',
            data: this.form,
          }).then(res => {
            if (res.data.code === 200) {
              this.$alert('欢迎使用 g3-cms', '提示', {
                confirmButtonText: '确定',
                callback: () => {
                  window.location.href = "/backend/";
                }
              });
            } else {
              this.$message({
                message: res.data.msg,
                type: 'error'
              });
            }
            this.fullscreenLoading = false;
          })
        } else {
          return false;
        }
      });
    },
    handleSubmit() {
      this.handleStep2FormSubmit();
    },
    handleTestConn() {
      this.handleStep1FormSubmit(false);
    },
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 20px;
}
.form-body {
  min-height: 300px;
  margin-top: 20px;
  width: 600px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04);
  text-align: center;
  display: inline-block;
  padding: 0 50px 50px 50px;
}
.form-steps {
  min-height: 300px;
}
.form-step-title {
  font-size: 18px;
  font-weight: bold;
  margin: 20px;
}
.form-footer {
}
</style>
