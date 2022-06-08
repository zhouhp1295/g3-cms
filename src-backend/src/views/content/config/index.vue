<template>
  <div class="app-container">
    <el-form ref="form" :model="form" :rules="rules" class="config-form" label-width="160px">
      <el-tabs v-model="activeName">
        <el-tab-pane label="基本配置" name="tab01">
          <el-col :span="18">
            <el-form-item label="网站域名" prop="host">
              <el-input v-model="form.host" />
              <el-col class="tip">
                例: https://htoolbox.site
              </el-col>
            </el-form-item>
            <el-form-item label="网站标题" prop="title">
              <el-input v-model="form.title" />
            </el-form-item>
            <el-form-item label="网站图标" prop="favicon">
              <el-upload
                class="img-uploader"
                :action="uploadUrl"
                :show-file-list="false"
                :multiple="false"
                :on-success="handleFaviconSuccess"
                :headers="headers"
                :before-upload="beforeUpload">
                <div v-if="form.favicon">
                  <img :src="parsePreviewSrc(form.favicon)" class="favicon">
                </div>
                <i v-else class="el-icon-plus favicon-uploader-icon"></i>
              </el-upload>
              <el-col class="tip">
                建议尺寸: 64 x 64 px
              </el-col>
            </el-form-item>
            <el-form-item label="网站Logo" prop="logo">
              <el-upload
                class="img-uploader"
                :action="uploadUrl"
                :show-file-list="false"
                :multiple="false"
                :on-success="handleLogoSuccess"
                :headers="headers"
                :before-upload="beforeUpload">
                <div v-if="form.logo">
                  <img :src="parsePreviewSrc(form.logo)" class="logo">
                </div>
                <i v-else class="el-icon-plus logo-uploader-icon"></i>
              </el-upload>
              <el-col class="tip">
                建议尺寸: 150 x 60 px
              </el-col>
            </el-form-item>
            <el-form-item label="网站关键字" prop="pid">
              <el-input type="textarea" v-model="form.keywords" />
            </el-form-item>
            <el-form-item label="网站描述" prop="pid">
              <el-input type="textarea" v-model="form.description" />
            </el-form-item>
            <el-form-item label="Copyright" prop="copyright">
              <el-input v-model="form.copyright" />
            </el-form-item>
            <el-form-item label="备案号" prop="beian">
              <el-input v-model="form.beian" />
            </el-form-item>
            <el-form-item label="公安网备案号" prop="gonganBeian">
              <el-input v-model="form.gonganBeian" />
            </el-form-item>
          </el-col>
        </el-tab-pane>
        <el-tab-pane label="站长配置" name="tab02">
          <el-col :span="18">
            <el-form-item label="Baidu站点验证码" prop="baiduSiteVerification">
              <el-input v-model="form.baiduSiteVerification" type="textarea" />
            </el-form-item>
            <el-form-item label="robots.txt" prop="robots">
              <el-input v-model="form.robots" type="textarea" :autosize="{ minRows: 10, maxRows: 15}" />
            </el-form-item>
          </el-col>
        </el-tab-pane>
        <el-tab-pane label="文章配置" name="tab04">
          <el-col :span="18">
            <el-form-item label="文章末尾声明" prop="articleSuffix">
              <el-input v-model="form.articleSuffix" type="textarea" :autosize="{ minRows: 6, maxRows: 10}" />
            </el-form-item>
          </el-col>
        </el-tab-pane>
      </el-tabs>
    </el-form>
    <el-col style="text-align: center;">
      <el-button type="warn" @click="onRefreshCache">清理缓存</el-button>
      <el-button type="primary" @click="submitForm">保存配置</el-button>
    </el-col>
  </div>
</template>

<script>
import {saveConfig, getConfig, refreshCache} from "@/api/content/config";
import {getToken} from "@/utils/auth";

export default {
  name: "CmsConfig",
  data() {
    return {
      activeName: 'tab01',
      // 文件上传路径
      uploadUrl: process.env.VUE_APP_BASE_API + "/api/common/uploadImage",
      headers: {
        Authorization: "Bearer " + getToken()
      },
      // form
      form: {
        title: undefined,
        host: undefined,
        keywords: undefined,
        description: undefined,
        copyright: undefined,
        favicon: undefined,
        logo: undefined,
        baiduSiteVerification: undefined,
        beian: undefined,
        gonganBeian:undefined,
        articleSuffix: undefined,
        robots: undefined,
      },
      // rules
      rules: {},
    };
  },
  created() {
    getConfig().then(res => {
      const data = res.data;
      this.form = Object.assign(this.form, data);
    });
  },
  methods: {
    onRefreshCache() {
      refreshCache().then(() => {
        this.$modal.msgSuccess("更新成功");
      });
    },
    submitForm() {
      this.$refs["form"].validate(valid => {
        if (valid) {
          saveConfig(this.form).then(() => {
            this.$modal.msgSuccess("更新成功");
          });
        }
      });
    },
    beforeUpload(file) {
      if (file.type !== 'image/jpeg' && file.type !== 'image/png' && file.type !== 'image/x-icon') {
        this.$message.error('上传图片只能是 JPG、PNG、ICO 格式!');
        return false;
      }
      if (file.size / 1024 / 1024 > 2) {
        this.$message.error('上传图片大小不能超过 2MB!');
        return false;
      }
      return true;
    },
    handleFaviconSuccess(res) {
      this.form.favicon = res.data.url;
    },
    handleLogoSuccess(res) {
      this.form.logo = res.data.url;
    },
  }
}
</script>

<style>
  .config-form {
    height: 600px;
    overflow-y: scroll;
  }
  .config-form .tip {
    font-size: 12px;
    color: #3A71A8;
  }
  .config-form .img-uploader .el-upload {
    border: 2px dashed #d9d9d9;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
  }
  .config-form .img-uploader .el-upload:hover {
    border-color: #409EFF;
  }
  .config-form .logo-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 242px;
    height: 60px;
    line-height: 60px;
    text-align: center;
  }
  .config-form .logo {
    width: 242px;
    height: 60px;
    object-fit:cover;
  }
  .config-form .favicon-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 64px;
    height: 64px;
    line-height: 64px;
    text-align: center;
  }
  .config-form .favicon {
    width: 64px;
    height: 64px;
    object-fit:cover;
  }
</style>
