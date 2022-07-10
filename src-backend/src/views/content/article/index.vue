<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryForm" :inline="true" v-show="showSearch" label-width="68px">
      <el-form-item label="文章名称" prop="title">
        <el-input
          v-model="queryParams.title"
          placeholder="请输入文章名称"
          clearable
          size="small"
          style="width: 240px"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="文章分类" prop="title">
        <treeselect
          v-model="queryParams.category"
          clearable
          size="small"
          style="width: 240px"
          :options="categoryOptions"
          :normalizer="normalizer"
          @keyup.enter.native="handleQuery"
          placeholder="请输入文章分类" />
      </el-form-item>
      <el-form-item label="作者" prop="writer">
        <el-select
          v-model="queryParams.writer"
          placeholder="作者"
          clearable
          size="small"
          style="width: 240px"
        >
          <el-option
            v-for="item in writerOptions"
            :key="item.id"
            :label="item.label"
            :value="item.id"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="发布状态" prop="status">
        <el-select
          v-model="queryParams.status"
          placeholder="发布状态"
          clearable
          size="small"
          style="width: 240px"
        >
          <el-option
            v-for="dict in dict.type.sys_normal_yes_no"
            :key="dict.value"
            :label="dict.label"
            :value="dict.value"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="Banner" prop="inBanner">
        <el-select
          v-model="queryParams.inBanner"
          placeholder="Banner"
          clearable
          size="small"
          style="width: 240px"
        >
          <el-option
            v-for="dict in dict.type.sys_normal_yes_no"
            :key="dict.value"
            :label="dict.label"
            :value="dict.value"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="创建时间">
        <el-date-picker
          v-model="dateRange"
          size="small"
          style="width: 240px"
          value-format="yyyy-MM-dd"
          type="daterange"
          range-separator="-"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
        ></el-date-picker>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>

    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button
          type="primary"
          plain
          icon="el-icon-plus"
          size="mini"
          @click="handleAdd"
          v-hasPermi="['content:article:add']"
        >新增</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          type="danger"
          plain
          icon="el-icon-delete"
          size="mini"
          :disabled="multiple"
          @click="handleDelete"
          v-hasPermi="['content:article:remove']"
        >删除</el-button>
      </el-col>
      <right-toolbar :showSearch.sync="showSearch" @queryTable="getList"></right-toolbar>
    </el-row>

    <el-table v-loading="loading" :data="articleList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" align="center" width="30"/>
      <el-table-column label="ID" align="center" prop="id"/>
      <el-table-column label="标题" align="center" prop="title" width="200" :show-overflow-tooltip="true" />
      <el-table-column label="分类" align="center" prop="category">
        <template slot-scope="scope">
          <span>{{ getCategoryName(scope.row.category )}}</span>
        </template>
      </el-table-column>
      <el-table-column label="作者" align="center" prop="writer">
        <template slot-scope="scope">
          <span>{{ getWriterName(scope.row.writer )}}</span>
        </template>
      </el-table-column>
      <el-table-column label="发布" align="center" key="status">
        <template slot-scope="scope">
          <el-switch
            v-model="scope.row.status"
            active-value="1"
            inactive-value="0"
            @change="handleStatusChange(scope.row)"
          ></el-switch>
        </template>
      </el-table-column>
      <el-table-column label="显示到Banner" align="center" key="inBanner">
        <template slot-scope="scope">
          <el-switch
            v-model="scope.row.inBanner"
            active-value="1"
            inactive-value="0"
            @change="handleInBannerChange(scope.row)"
          ></el-switch>
        </template>
      </el-table-column>
      <el-table-column label="时间" align="center" prop="createdAt" width="180">
        <template slot-scope="scope">
          <span v-if="scope.row.status == 1">发布时间:<br/>{{ parseTime(scope.row.publishedAt) }}<br/></span>
          <span>更新时间:<br/>{{ parseTime(scope.row.updatedAt) }}<br/></span>
          <span>创建时间:<br/>{{ parseTime(scope.row.createdAt) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdate(scope.row)"
            v-hasPermi="['content:article:edit']"
          >修改</el-button>
          <el-button
            size="mini"
            type="text"
            icon="el-icon-delete"
            @click="handleDelete(scope.row)"
            v-hasPermi="['content:article:remove']"
          >删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="queryParams.pageNum"
      :limit.sync="queryParams.pageSize"
      @pagination="getList"
    />

    <!-- 添加或修改文章配置对话框 -->
    <el-drawer :title="title" :visible.sync="open" size="100%" direction="rtl" append-to-body>
      <el-col :span="20" :offset="2">
        <el-form ref="form" :model="form" :rules="rules" label-width="120px" style="margin-bottom: 100px;">
          <el-tabs v-model="activeName">
            <el-tab-pane label="基本信息" name="tab01">
              <el-row>
                <el-col :span="24">
                  <el-form-item label="文章名称" prop="title" required>
                    <el-input v-model="form.title" placeholder="请输入文章名称" />
                  </el-form-item>
                </el-col>
              </el-row>
              <el-row>
                <el-col :span="8">
                  <el-form-item label="文章分类" prop="category">
                    <treeselect v-model="form.category" :options="categoryOptions" :normalizer="normalizer" placeholder="选择文章分类" />
                  </el-form-item>
                </el-col>
              </el-row>
              <el-form-item label="内容" prop="content" required>
                <editor v-model="form.content" :min-height="300" :height="600"/>
              </el-form-item>
            </el-tab-pane>
            <el-tab-pane label="SEO设置" name="tab02">
              <el-form-item label="SEO标题" prop="seoTitle">
                <el-input v-model="form.seoTitle" placeholder="请输入文章关键字" />
              </el-form-item>
              <el-form-item label="SEO关键字" prop="keywords">
                <el-input v-model="form.seoKeywords" type="textarea" :autosize="{ minRows: 8, maxRows: 20}" placeholder="请输入文章关键字" />
              </el-form-item>
              <el-form-item label="SEO描述" prop="description">
                <el-input v-model="form.seoDescription" type="textarea" :autosize="{ minRows: 8, maxRows: 20}" placeholder="请输入文章描述" />
              </el-form-item>
            </el-tab-pane>
            <el-tab-pane label="图片" name="tab03">
              <el-form-item label="封面" prop="cover">
                <el-upload
                  class="cover-uploader"
                  :action="uploadUrl"
                  :file-list="coverList"
                  :multiple="false"
                  :on-success="handleCoverSuccess"
                  :on-remove="handleCoverRemove"
                  :headers="headers"
                  :before-upload="beforeUpload">
                  <img :src="parsePreviewSrc(form.cover)"  v-if="form.cover" class="cover">
                  <i v-else class="el-icon-plus cover-uploader-icon"></i>
                </el-upload>
              </el-form-item>
              <el-form-item label="Banner图片" prop="banner">
                <el-col>
                  建议尺寸: 820 x 200 px
                </el-col>
                <el-upload
                  class="banner-uploader"
                  :action="uploadUrl"
                  :multiple="false"
                  :file-list="bannerList"
                  :on-success="handleBannerSuccess"
                  :on-remove="handleBannerRemove"
                  :headers="headers"
                  :before-upload="beforeUpload">
                  <div v-if="form.banner">
                    <img :src="parsePreviewSrc(form.banner)" class="banner">
                  </div>
                  <i v-else class="el-icon-plus banner-uploader-icon"></i>
                </el-upload>
              </el-form-item>
            </el-tab-pane>
            <el-tab-pane label="其他" name="tab04">
              <el-form-item label="作者" prop="writer">
                <el-select
                  v-model="form.writer"
                  placeholder="作者"
                  clearable
                  filterable
                  size="small"
                  style="width: 240px"
                >
                  <el-option
                    v-for="item in writerOptions"
                    :key="item.id"
                    :label="item.label"
                    :value="item.id"
                  />
                </el-select>
              </el-form-item>
              <el-form-item label="标签" prop="tags">
                <el-checkbox-group v-model="form.tags" size="small">
                  <el-checkbox v-for="item in tagOptions" :key="item.id" :label="item.id" border>{{item.label}}</el-checkbox>
                </el-checkbox-group>
                <el-input
                  class="input-new-tag"
                  v-if="newTagInputVisible"
                  v-model="newTagInputValue"
                  ref="saveTagInput"
                  size="small"
                  @keyup.enter.native="handleNewTagInputConfirm"
                  @blur="handleNewTagInputConfirm"
                >
                </el-input>
                <el-button v-else class="button-new-tag" size="small" @click="showNewTagInput">+ 新标签</el-button>
              </el-form-item>
              <el-form-item label="发布">
                <el-radio-group v-model="form.status">
                  <el-radio
                    v-for="dict in dict.type.sys_normal_yes_no"
                    :key="dict.value"
                    :label="dict.value"
                  >{{dict.label}}</el-radio>
                </el-radio-group>
              </el-form-item>
              <el-form-item label="在Banner显示">
                <el-radio-group v-model="form.inBanner">
                  <el-radio
                    v-for="dict in dict.type.sys_normal_yes_no"
                    :key="dict.value"
                    :label="dict.value"
                  >{{dict.label}}</el-radio>
                </el-radio-group>
              </el-form-item>
              <el-form-item label="在Banner中排序" v-if="form.inBanner == 1">
                <el-input v-model="form.inBannerSort" placeholder="请输入" type="number"/>
              </el-form-item>
              <el-form-item label="排序" prop="sort">
                <el-input v-model="form.sort" type="number" placeholder="请输入排序" />
              </el-form-item>
              <el-form-item label="备注" prop="remark">
                <el-input v-model="form.remark" type="textarea" :autosize="{ minRows: 1, maxRows: 3}"  placeholder="请输入备注" />
              </el-form-item>
            </el-tab-pane>
          </el-tabs>
        </el-form>
      </el-col>
      <div style="position: fixed; bottom: 0; width: 100%; background: aliceblue;">
        <el-row >
          <el-col :span="12" :offset="6">
            <div style="text-align: center; display: block; height: 60px; padding-top: 10px;">
              <el-button type="primary" @click="submitForm">确 定</el-button>
              <el-button @click="cancel">取 消</el-button>
            </div>
          </el-col>
        </el-row>
      </div>
    </el-drawer>
  </div>
</template>
<style>
  .el-drawer__header {
    background: aliceblue;
  }
  .cover-uploader .el-upload {
    border: 2px dashed #d9d9d9;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
  }
  .cover-uploader .el-upload:hover {
    border-color: #409EFF;
  }
  .cover-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }
  .cover {
    width: 178px;
    height: 178px;
    display: block;
  }
  .table-cover {
    width: 50px;
    height: 50px;
    display: block;
  }
  .custom-slot .label {
    text-align: right;
  }
  .custom-slot .info {
    text-align: left;
  }
  .button-new-tag {
    height: 32px;
    line-height: 30px;
    padding-top: 0;
    padding-bottom: 0;
  }
  .input-new-tag {
    width: 90px;
    vertical-align: bottom;
  }

  .banner-uploader .el-upload {
    border: 2px dashed #d9d9d9;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
  }
  .banner-uploader .el-upload:hover {
    border-color: #409EFF;
  }
  .banner-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 380px;
    height: 90px;
    line-height: 90px;
    text-align: center;
  }
  .banner {
    width: 380px;
    height: 90px;
    display: block;
  }
</style>
<script>
import {
  addArticle,
  changeArticleInBanner,
  changeArticleStatus,
  delArticle,
  getArticle,
  listArticle,
  updateArticle
} from "@/api/content/article";
import {listCategory} from "@/api/content/category";
import {fastAddTag} from "@/api/content/tag";
import {listWriterOptions} from "@/api/content/writer";
import Treeselect from "@riophae/vue-treeselect";
import "@riophae/vue-treeselect/dist/vue-treeselect.css";
import {getToken} from "@/utils/auth";

export default {
  name: "Article",
  dicts: ['sys_normal_yes_no'],
  components: {Treeselect},
  data() {
    return {
      activeName: 'tab01',
      // 遮罩层
      loading: true,
      // 选中数组
      ids: [],
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      // 显示搜索条件
      showSearch: true,
      // 总条数
      total: 0,
      // 文章表格数据
      articleList: [],
      // 分类
      categoryOptions: [],
      // 标签
      tagOptions: [],
      // 作者
      writerOptions: [],
      // 弹出层标题
      title: "",
      // 是否显示弹出层
      open: false,
      // 日期范围
      dateRange: [],
      // 文件上传路径
      uploadUrl: process.env.VUE_APP_BASE_API + "/api/common/uploadImage",
      headers: {
        Authorization: "Bearer " + getToken()
      },
      // 查询文章
      queryParams: {
        pageNum: 1,
        pageSize: 10,
        title: undefined,
        keywords: undefined,
        description: undefined,
        category: undefined,
      },
      // 表单
      form: {

      },
      coverList: [],
      bannerList: [],
      // 表单校验
      rules: {
        title: [
          { required: true, message: "文章标题不能为空", trigger: "blur" }
        ],
        content: [
          { required: true, message: "文章标题不能为空", trigger: "blur" }
        ]
      },
      // 新建标签数据
      newTagInputVisible: false,
      newTagInputValue: '',
    };
  },
  computed: {
  },
  created() {
    this.getCategories();
    this.getWriters();
    this.getList();
  },
  methods: {
    /** 查询文章列表 */
    getList() {
      this.loading = true;
      listArticle(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
          this.articleList = response.data.rows;
          this.total = response.data.page.total;
          this.loading = false;
        }
      );
    },
    /** 转换分类数据结构 */
    normalizer(node) {
      if (node.children && !node.children.length) {
        delete node.children;
      }
      return {
        id: node.id,
        label: node.title,
        children: node.children
      };
    },
    getWriters() {
      listWriterOptions().then(response => {
        this.writerOptions = response.data.rows;
      });
    },
    getCategoryName(category) {
      let str = "";
      if (category <= 0 ) {
        return str;
      }
      this.categoryOptions.forEach(ele => {
        if (ele.id === category) {
          let pStr = this.getCategoryName(ele.pid);
          if (pStr !== "") {
            str = pStr + ">" + ele.title;
          } else {
            str = ele.title;
          }
        }
      })
      return str;
    },
    getWriterName(writer) {
      let str = "佚名";
      if (writer <= 0 ) {
        return str;
      }
      this.writerOptions.forEach(ele => {
        if (ele.id === writer) {
          str = ele.label;
        }
      })
      return str;
    },
    getCategories() {
      listCategory().then(response => {
        let rows = response.data.rows;
        this.categoryOptions = this.handleTree(rows, "id", "pid");
      });
    },
    beforeUpload(file) {
      if (file.type !== 'image/jpeg' && file.type !== 'image/png') {
        this.$message.error('上传图片只能是 JPG、PNG 格式!');
        return false;
      }
      if (file.size / 1024 / 1024 > 2) {
        this.$message.error('上传图片大小不能超过 2MB!');
        return false;
      }
      return true;
    },
    handleCoverSuccess(res, file) {
      this.form.cover = res.data.url;
      this.coverList = [{name:res.data.url, url:this.parsePreviewSrc(res.data.url)}];
    },
    handleCoverRemove() {
      this.form.cover = "";
      this.coverList = [];
    },
    handleBannerSuccess(res, file) {
      this.form.banner = res.data.url;
      this.bannerList = [{name:res.data.url, url:this.parsePreviewSrc(res.data.url)}];
    },
    handleBannerRemove() {
      this.form.banner = "";
      this.bannerList = [];
    },
    // 取消按钮
    cancel() {
      this.open = false;
      this.reset();
    },
    // 表单重置
    reset() {
      this.form = {
        id: undefined,
        title: undefined,
        seoTitle: undefined,
        seoKeywords: undefined,
        seoDescription: undefined,
        remark: undefined,
        cover: undefined,
        category: undefined,
        content: undefined,
        inBanner: "0",
        inBannerSort: 0,
        tags: [],
        sort: 0,
        status:"0",
      };
      this.coverList = [];
      this.bannerList = [];
      this.resetForm("form");
    },
    /** 搜索按钮操作 */
    handleQuery() {
      this.queryParams.pageNum = 1;
      this.getList();
    },
    /** 重置按钮操作 */
    resetQuery() {
      this.dateRange = [];
      this.resetForm("queryForm");
      this.handleQuery();
    },
    /** 新增按钮操作 */
    handleAdd() {
      this.reset();
      this.getCategories();
      this.open = true;
      this.title = "添加文章";
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.id)
      this.single = selection.length!==1
      this.multiple = !selection.length
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset();
      const id = row.id || this.ids;
      this.getCategories();
      getArticle(id).then(response => {
        this.tagOptions = response.data.tagOptions;
        this.form = response.data;
        if (this.form.cover !== undefined && this.form.cover.length > 0) {
          this.coverList = [{name:this.form.cover, url: this.parsePreviewSrc(this.form.cover)}];
        }
        if (this.form.banner !== undefined && this.form.banner.length > 0) {
          this.bannerList = [{name:this.form.banner, url: this.parsePreviewSrc(this.form.banner)}];
        }
        this.open = true;
        this.title = "修改文章";
      });
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs["form"].validate(valid => {
        if (valid) {
          this.form.inBannerSort = parseInt(this.form.inBannerSort);
          this.form.sort = parseInt(this.form.sort);
          if (this.form.id !== undefined) {
            updateArticle(this.form).then(response => {
              this.$modal.msgSuccess("修改成功");
              this.open = false;
              this.getList();
            });
          } else {
            addArticle(this.form).then(response => {
              this.$modal.msgSuccess("新增成功");
              this.open = false;
              this.getList();
            });
          }
        }
      });
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      const ids = row.id || this.ids;
      this.$modal.confirm('是否确认删除文章ID为"' + ids + '"的数据项？').then(function() {
          return delArticle(ids);
        }).then(() => {
          this.getList();
          this.$modal.msgSuccess("删除成功");
        }).catch(() => {});
    },
    // 发布状态修改
    handleStatusChange(row) {
      let text = row.status === "1" ? "发布" : "取消发布";
      this.$modal.confirm('确认要"' + text + '""' + row.title + '"吗？').then(function() {
        return changeArticleStatus(row.id, row.status);
      }).then(() => {
        this.$modal.msgSuccess(text + "成功");
      }).catch(function() {
        row.status = row.status === "0" ? "1" : "0";
      });
    },
    handleClose(tag) {

    },
    showNewTagInput() {
      this.newTagInputVisible = true;
      this.$nextTick(_ => {
        this.$refs.saveTagInput.$refs.input.focus();
      });
    },
    handleNewTagInputConfirm() {
      let newTagInputValue = this.newTagInputValue;
      if (newTagInputValue) {
        fastAddTag({"title":newTagInputValue}).then(res => {
          this.tagOptions.push({"id":res.data.id, "label":res.data.title});
          this.form.tags.push(res.data.id);
        });
      }
      this.newTagInputVisible = false;
      this.newTagInputValue = '';
    },
    handleInBannerChange(row) {
      let text = row.inBanner === "1" ? "设置显示到Banner" : "取消显示到Banner";
      if (row.inBanner !== "1") {
        this.$modal.confirm('确认要将【' + row.title + '】' + text + '吗？').then(function () {
          return changeArticleInBanner({id: row.id, inBanner: row.inBanner, inBannerSort: 0});
        }).then(() => {
          this.$modal.msgSuccess(text + "成功");
        }).catch(function () {
          row.inBanner = row.inBanner === "0" ? "1" : "0";
        });
      } else {
        this.$prompt('确认要将【' + row.title + '】' + text + '吗？请输入排序').then(({value}) => {
          const v = parseInt(value)
          return changeArticleInBanner({
            id: row.id,
            inBanner: row.inBanner,
            inBannerSort: isNaN(v) ? row.inBannerSort : v
          });
        }).then(() => {
          this.$modal.msgSuccess(text + "成功");
        }).catch(function () {
          row.inBanner = row.inBanner === "0" ? "1" : "0";
        });
      }
    }
  }
};
</script>
