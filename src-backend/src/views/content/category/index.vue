<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryForm" :inline="true" v-show="showSearch" label-width="68px">
      <el-form-item label="分类名称" prop="title">
        <el-input
          v-model="queryParams.title"
          placeholder="请输入分类名称"
          clearable
          size="small"
          style="width: 240px"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="菜单" prop="inMenu">
        <el-select
          v-model="queryParams.inMenu"
          placeholder="菜单"
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
      <el-form-item label="关键字" prop="keywords">
        <el-input
          v-model="queryParams.keywords"
          placeholder="请输入关键字"
          clearable
          size="small"
          style="width: 240px"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="描述" prop="description">
        <el-input
          v-model="queryParams.description"
          placeholder="请输入描述"
          clearable
          size="small"
          style="width: 240px"
          @keyup.enter.native="handleQuery"
        />
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
          v-hasPermi="['content:category:add']"
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
          v-hasPermi="['content:category:remove']"
        >删除</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          type="info"
          plain
          icon="el-icon-sort"
          size="mini"
          @click="toggleExpandAll"
        >展开/折叠</el-button>
      </el-col>
      <right-toolbar :showSearch.sync="showSearch" @queryTable="getList"></right-toolbar>
    </el-row>

    <el-table
      v-if="refreshTable"
      v-loading="loading"
      :data="categoryList"
      @selection-change="handleSelectionChange"
      row-key="id"
      :default-expand-all="isExpandAll"
      :tree-props="{children: 'children', hasChildren: 'hasChildren'}"
      >
      <el-table-column type="selection" width="30" align="center" />
      <el-table-column label="ID" align="center" prop="id"/>
      <el-table-column label="" align="center" prop="cover">
        <template slot-scope="scope">
          <img v-if="scope.row.cover" :src="parsePreviewSrc(scope.row.cover)" class="table-cover">
        </template>
      </el-table-column>
      <el-table-column label="标题" align="center" prop="title" :show-overflow-tooltip="true" />
      <el-table-column label="Icon" align="center" prop="icon"/>
      <el-table-column label="显示到菜单" align="center" key="inMenu">
        <template slot-scope="scope">
          <el-switch
            v-model="scope.row.inMenu"
            active-value="1"
            inactive-value="0"
            @change="handleInMenuChange(scope.row)"
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
      <el-table-column label="创建时间" align="center" prop="createdAt" width="180">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.createdAt) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="left" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdate(scope.row)"
            v-hasPermi="['content:category:edit']"
          >修改</el-button>
          <el-button
            size="mini"
            type="text"
            icon="el-icon-delete"
            @click="handleDelete(scope.row)"
            v-hasPermi="['content:category:remove']"
          >删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 添加或修改分类配置对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="800px" append-to-body>
      <el-form ref="form" :model="form" :rules="rules" label-width="160px">
        <el-form-item label="上级分类" prop="pid">
          <treeselect v-model="form.pid" :options="categoryOptions" :normalizer="normalizer" placeholder="选择上级分类" />
        </el-form-item>
        <el-form-item label="分类名称" prop="title">
          <el-input v-model="form.title" placeholder="请输入分类名称" />
        </el-form-item>
        <el-form-item label="Icon" prop="icon">
          <el-input v-model="form.icon" placeholder="请输入icon" />
        </el-form-item>
        <el-form-item label="图片" prop="cover">
          <el-upload
            class="cover-uploader"
            :action="uploadUrl"
            :multiple="false"
            :file-list="coverList"
            :on-success="handleCoverSuccess"
            :on-remove="handleCoverRemove"
            :headers="headers"
            :before-upload="beforeUpload">
            <div v-if="form.cover">
              <img :src="parsePreviewSrc(form.cover)" class="cover">
            </div>
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
        <el-form-item label="关键字" prop="keywords">
          <el-input v-model="form.keywords" type="textarea" placeholder="请输入分类代码" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="form.description" type="textarea" placeholder="请输入分类值" />
        </el-form-item>
        <el-form-item label="在菜单中显示">
          <el-radio-group v-model="form.inMenu">
            <el-radio
              v-for="dict in dict.type.sys_normal_yes_no"
              :key="dict.value"
              :label="dict.value"
            >{{dict.label}}</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="在菜单中排序"  v-if="form.inMenu == 1">
          <el-input v-model="form.inMenuSort" placeholder="请输入" type="number"/>
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
        <el-form-item label="备注" prop="remark">
          <el-input v-model="form.remark" type="textarea" placeholder="请输入内容" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="submitForm">确 定</el-button>
        <el-button @click="cancel">取 消</el-button>
      </div>
    </el-dialog>
  </div>
</template>
<style>
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
    width: 80px;
    height: 90px;
    display: block;
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
import { listCategory, getCategory, delCategory, addCategory, updateCategory, changeCategoryInMenu, changeCategoryInBanner } from "@/api/content/category";
import Treeselect from "@riophae/vue-treeselect";
import "@riophae/vue-treeselect/dist/vue-treeselect.css";
import {getToken} from "@/utils/auth";

export default {
  name: "Category",
  dicts: ['sys_normal_yes_no'],
  components: {Treeselect},
  data() {
    return {
      // 遮罩层
      loading: true,
      // 是否展开，默认全部展开
      isExpandAll: true,
      // 重新渲染表格状态
      refreshTable: true,
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
      // 分类表格数据
      categoryList: [],
      // 分类
      categoryOptions: [],
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
      // 查询分类
      queryParams: {
        title: undefined,
        keywords: undefined,
        description: undefined,
        inMenu: undefined,
        inBanner: undefined,
      },
      // 表单分类
      form: {},
      coverList: [],
      bannerList: [],
      // 表单校验
      rules: {
        title: [
          { required: true, message: "分类代码不能为空", trigger: "blur" }
        ]
      }
    };
  },
  created() {
    this.getList();
  },
  methods: {
    /** 查询分类列表 */
    getList() {
      this.loading = true;
      listCategory(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
          this.categoryList = this.handleTree(response.data.rows, "id", "pid");
          this.loading = false;
        }
      );
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
     /** 展开/折叠操作 */
    toggleExpandAll() {
      this.refreshTable = false;
      this.isExpandAll = !this.isExpandAll;
      this.$nextTick(() => {
        this.refreshTable = true;
      });
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
        pid: undefined,
        title: undefined,
        icon: undefined,
        keywords: undefined,
        description: undefined,
        remark: undefined,
        cover: undefined,
        banner: undefined,
        inBanner: "0",
        inBannerSort: 0,
        inMenu: "0",
        inMenuSort: 0,
      };
      this.coverList = [];
      this.bannerList = [];
      this.resetForm("form");
    },
    /** 搜索按钮操作 */
    handleQuery() {
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
      this.open = true;
      this.title = "添加分类";
      listCategory().then(response => {
        let rows = response.data.rows;
        this.categoryOptions = this.handleTree(rows, "id", "pid");
      });
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
      const id = row.id || this.ids
      getCategory(id).then(response => {
        this.form = response.data;
        if (this.form.cover !== undefined && this.form.cover.length > 0) {
          this.coverList = [{name:this.form.cover, url: this.parsePreviewSrc(this.form.cover)}];
        }
        if (this.form.banner !== undefined && this.form.banner.length > 0) {
          this.bannerList = [{name:this.form.banner, url: this.parsePreviewSrc(this.form.banner)}];
        }
        this.open = true;
        this.title = "修改分类";
      });
      listCategory({excludeId:row.id}).then(response => {
        let rows = response.data.rows;
        this.categoryOptions = this.handleTree(rows, "id", "pid");
      });
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs["form"].validate(valid => {
        if (valid) {
          this.form.inMenuSort = parseInt(this.form.inMenuSort);
          this.form.inBannerSort = parseInt(this.form.inBannerSort);
          this.form.sort = parseInt(this.form.sort);
          if (this.form.id !== undefined) {
            updateCategory(this.form).then(response => {
              this.$modal.msgSuccess("修改成功");
              this.open = false;
              this.getList();
            });
          } else {
            addCategory(this.form).then(response => {
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
      this.$modal.confirm('是否确认删除分类ID为"' + ids + '"的数据项？').then(function() {
          return delCategory(ids);
        }).then(() => {
          this.getList();
          this.$modal.msgSuccess("删除成功");
        }).catch(() => {});
    },
    handleInMenuChange(row) {
      let text = row.inMenu === "1" ? "设置显示到菜单" : "取消显示到菜单";
      if (row.inMenu !== "1") {
        this.$modal.confirm('确认要将【' + row.title + '】' + text + '吗？').then(function() {
          return changeCategoryInMenu({id:row.id, inMenu:row.inMenu,inMenuSort: 0});
        }).then(() => {
          this.$modal.msgSuccess(text + "成功");
        }).catch(function() {
          row.inMenu = row.inMenu === "0" ? "1" : "0";
        });
      } else {
        this.$prompt('确认要将【' + row.title + '】' + text + '吗？请输入排序').then(({ value }) => {
          const v = parseInt(value)
          return changeCategoryInMenu({id:row.id,inMenu:row.inMenu, inMenuSort: isNaN(v) ? row.inMenuSort : v});
        }).then(() => {
          this.$modal.msgSuccess(text + "成功");
        }).catch(function() {
          row.inMenu = row.inMenu === "0" ? "1" : "0";
        });
      }
    },
    handleInBannerChange(row) {
      let text = row.inBanner === "1" ? "设置显示到Banner" : "取消显示到Banner";
      if (row.inBanner !== "1") {
        this.$modal.confirm('确认要将【' + row.title + '】' + text + '吗？').then(function() {
          return changeCategoryInBanner({id:row.id, inBanner:row.inBanner,inBannerSort: 0});
        }).then(() => {
          this.$modal.msgSuccess(text + "成功");
        }).catch(function() {
          row.inBanner = row.inBanner === "0" ? "1" : "0";
        });
      } else {
        this.$prompt('确认要将【' + row.title + '】' + text + '吗？请输入排序').then(({ value }) => {
          const v = parseInt(value)
          return changeCategoryInBanner({id:row.id,inBanner:row.inBanner, inBannerSort: isNaN(v) ? row.inBannerSort : v});
        }).then(() => {
          this.$modal.msgSuccess(text + "成功");
        }).catch(function() {
          row.inBanner = row.inBanner === "0" ? "1" : "0";
        });
      }

    },
  }
};
</script>
