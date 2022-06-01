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
          </el-col>
        </el-tab-pane>
        <el-tab-pane label="站长配置" name="tab02">
          <el-col :span="18">
            <el-form-item label="Baidu站点验证码" prop="baiduSiteVerification">
              <el-input v-model="form.baiduSiteVerification" type="textarea" />
            </el-form-item>
            <el-form-item label="统计代码" prop="trackCode">
              <el-input v-model="form.trackCode" type="textarea" />
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
            <el-form-item label="robots.txt" prop="robots">
              <el-input v-model="form.robots" type="textarea" :autosize="{ minRows: 10, maxRows: 15}" />
            </el-form-item>
          </el-col>
        </el-tab-pane>
        <el-tab-pane label="栏目配置" name="tab03">
          <el-col :span="18">
            <el-form-item label="首页显示分类" prop="indexCategories">
              <el-checkbox-group v-model="form.indexCategories">
                <el-checkbox v-for="item in categoryOptions" :key="item.id" :label="item.id" border>{{item.title}}</el-checkbox>
              </el-checkbox-group>
            </el-form-item>
            <el-form-item label="边栏顶部分类" prop="rightTopCategory" placeholder="请选择">
              <el-select v-model="form.rightTopCategory">
                <el-option
                  v-for="item in categoryOptions"
                  :key="item.id"
                  :label="item.title"
                  :value="item.id+''">
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-tab-pane>
        <el-tab-pane label="文章配置" name="tab04">
          <el-col :span="18">
            <el-form-item label="原创文章版权声明" prop="articleCopyright">
              <el-input v-model="form.articleCopyright" type="textarea" :autosize="{ minRows: 6, maxRows: 10}" />
            </el-form-item>
            <el-form-item label="非原创文章末尾声明" prop="articleSuffix">
              <el-input v-model="form.articleSuffix" type="textarea" :autosize="{ minRows: 6, maxRows: 10}" />
            </el-form-item>
          </el-col>
        </el-tab-pane>
        <el-tab-pane label="版本授权" name="tab05">
          <el-col :span="18">
            <el-form-item label="产品名称">
              <span>{{form.productName}}</span>
            </el-form-item>
            <el-form-item label="产品标识">
              <span>{{form.product}}</span>
            </el-form-item>
            <el-form-item label="当前版本">
              <span>{{form.productVersion}}</span>
            </el-form-item>
            <el-form-item label="授权类型">
              <span v-if="form.grant">正版授权</span>
              <span v-else>试用版<a href="javascript:void(0);" class="gotobuy">获取授权码</a></span>
              <span v-if="form.grantExpired">(已过期)<a href="javascript:void(0);"  class="gotobuy">获取授权码</a></span>
            </el-form-item>
            <el-form-item label="过期时间">
              <span>{{form.grantExpiredAt}}</span>
            </el-form-item>
            <el-form-item label="授权网址">
              <span v-if="form.grantWebsites != undefined && form.grantWebsites.length > 0">{{form.grantWebsites}}</span>
              <span v-else>无</span>
            </el-form-item>
            <el-form-item label="授权码" prop="grantToken">
              <el-input v-model="form.grantToken" :autosize="{ minRows: 6, maxRows: 10}" type="textarea" />
            </el-form-item>
          </el-col>
        </el-tab-pane>
<!--        <el-tab-pane label="微信配置" name="tab06">-->
<!--          <el-form-item label="二维码" prop="wechatQrcode">-->
<!--            <el-upload-->
<!--              class="img-uploader"-->
<!--              :action="uploadUrl"-->
<!--              :show-file-list="false"-->
<!--              :multiple="false"-->
<!--              :on-success="handleWechatSuccess"-->
<!--              :before-upload="beforeUpload">-->
<!--              <div v-if="form.wechatQrcode">-->
<!--                <img :src="parsePreviewSrc(form.wechatQrcode)" class="wechat">-->
<!--              </div>-->
<!--              <i v-else class="el-icon-plus wechat-uploader-icon"></i>-->
<!--            </el-upload>-->
<!--          </el-form-item>-->
<!--        </el-tab-pane>-->
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
import {listSimpleCategory} from "@/api/content/category";

export default {
  name: "CmsConfig",
  data() {
    return {
      activeName: 'tab01',
      // 文件上传路径
      uploadUrl: process.env.VUE_APP_BASE_API + "/api/common/uploadImage",
      // 文章分类
      categoryOptions: [],
      // form
      form: {
        title: undefined,
        host: undefined,
        keywords: undefined,
        description: undefined,
        trackCode: undefined,
        copyright: undefined,
        favicon: undefined,
        logo: undefined,
        baiduSiteVerification: undefined,
        beian: undefined,
        gonganBeian:undefined,
        articleCopyright: undefined,
        articleSuffix: undefined,
        robots: undefined,
        indexCategory: undefined,
        indexCategories: [],
        rightTopCategory: 0,
        wechatQrcode: undefined,
        grantToken: undefined,
      },
      // rules
      rules: {},
    };
  },
  created() {
    this.getCategories().then( () => {
      getConfig().then(res => {
        const data = res.data.row;
        this.form = Object.assign(this.form, data);
        if (data.indexCategory !== undefined) {
          const indexCategories = data.indexCategory.split(",");
          this.form.indexCategories = [];
          indexCategories.forEach(ele => {
            const i = parseInt(ele);
            if (!isNaN(i)) {
              this.form.indexCategories.push(parseInt(ele));
            }
          })
        }
      });
    })
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
          let data = Object.assign({}, this.form);
          delete data["product"];
          delete data["productName"];
          delete data["productVersion"];
          delete data["grant"];
          delete data["grantWebsites"];
          delete data["grantExpiredAt"];
          delete data["grantExpired"];
          data['indexCategory'] = data['indexCategories'].join(",");
          delete data['indexCategories'];
          saveConfig(data).then(() => {
            this.$modal.msgSuccess("更新成功");
          });
        }
      });
    },
    beforeUpload(file) {
      if (file.type != 'image/jpeg' && file.type != 'image/png' && file.type != 'image/x-icon') {
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
    handleWechatSuccess(res) {
      this.form.wechatQrcode = res.data.url;
    },
    getCategories() {
      return listSimpleCategory().then(response => {
        this.categoryOptions = response.data.rows;
      });
    },
  }
}
</script>

<style>
  .gotobuy {
    font-size: 12px;
    margin-left: 10px;
    color: blue;
  }
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
  .config-form .wechat-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 200px;
    height: 200px;
    line-height: 200px;
    text-align: center;
  }
  .config-form .wechat {
    width: 200px;
    height: 200px;
    object-fit:cover;
  }
</style>
