<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryForm" :inline="true" v-show="showSearch" label-width="68px">
      <el-form-item label="作者名称" prop="title">
        <el-input
          v-model="queryParams.name"
          placeholder="请输入作者名称"
          clearable
          size="small"
          style="width: 240px"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="性别" prop="sex">
        <el-select
          v-model="queryParams.sex"
          placeholder="性别"
          clearable
          size="small"
          style="width: 240px"
        >
          <el-option
            v-for="dict in dict.type.sys_user_sex"
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
          v-hasPermi="['content:writer:add']"
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
          v-hasPermi="['content:writer:remove']"
        >删除</el-button>
      </el-col>
      <right-toolbar :showSearch.sync="showSearch" @queryTable="getList"></right-toolbar>
    </el-row>

    <el-table v-loading="loading" :data="writerList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="30" align="center" />
      <el-table-column label="ID" align="center" prop="id" />
      <el-table-column label="头像" align="center" prop="avatar" width="80">
        <template slot-scope="scope">
          <img v-if="scope.row.avatar" :src="parsePreviewSrc(scope.row.avatar)" class="table-avatar">
        </template>
      </el-table-column>
      <el-table-column label="姓名" align="left" prop="name" :show-overflow-tooltip="true" />
      <el-table-column label="性别" align="center" prop="sex">
        <template slot-scope="scope">
          <span>{{ getSex(scope.row.sex) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="备注" align="center" prop="remark" :show-overflow-tooltip="true" />
      <el-table-column label="创建时间" align="center" prop="createdAt" width="180">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.createdAt) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdate(scope.row)"
            v-hasPermi="['content:writer:edit']"
          >修改</el-button>
          <el-button
            size="mini"
            type="text"
            icon="el-icon-delete"
            @click="handleDelete(scope.row)"
            v-hasPermi="['content:writer:remove']"
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

    <!-- 添加或修改作者配置对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="800px" append-to-body>
      <el-form ref="form" :model="form" :rules="rules" label-width="160px">
        <el-form-item label="作者名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入作者名称" />
        </el-form-item>
        <el-form-item label="性别" prop="setx">
          <el-select
            v-model="form.sex"
            placeholder="性别"
            clearable
            size="small"
            style="width: 240px"
          >
            <el-option
              v-for="dict in dict.type.sys_user_sex"
              :key="dict.value"
              :label="dict.label"
              :value="dict.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="头像" prop="avatar">
          <el-upload
            class="img-uploader"
            :action="uploadUrl"
            :file-list="avatarList"
            :multiple="false"
            :on-success="handleAvatarSuccess"
            :on-remove="handleAvatarRemove"
            :headers="headers"
            :before-upload="beforeUpload">
            <div v-if="form.avatar">
              <img :src="parsePreviewSrc(form.avatar)" class="avatar">
            </div>
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>
        </el-form-item>
        <el-form-item label="主页背景" prop="cover">
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
        <el-form-item label="关键字" prop="keywords">
          <el-input v-model="form.keywords" type="textarea" placeholder="请输入作者代码" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="form.description" type="textarea" placeholder="请输入作者值" />
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
  .img-uploader .el-upload {
    border: 2px dashed #d9d9d9;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
  }
  .img-uploader .el-upload:hover {
    border-color: #409EFF;
  }
  .cover-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 300px;
    height: 150px;
    line-height: 150px;
    text-align: center;
  }
  .cover {
    width: 300px;
    height: 150px;
    display: block;
  }
  .avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 60px;
    height: 60px;
    line-height: 60px;
    text-align: center;
  }
  .avatar {
    width: 60px;
    height: 60px;
  }
  .table-avatar {
    width: 60px;
    height: 60px;
  }
</style>
<script>
import { listWriter, getWriter, delWriter, addWriter, updateWriter } from "@/api/content/writer";
import {getToken} from "@/utils/auth";

export default {
  name: "Writer",
  dicts: ['sys_user_sex','sys_normal_yes_no'],
  data() {
    return {
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
      // 作者表格数据
      writerList: [],
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
      // 查询作者
      queryParams: {
        pageNum: 1,
        pageSize: 10,
        title: undefined,
        keywords: undefined,
        description: undefined,
        inMenu: undefined,
        inBanner: undefined,
      },
      // 表单作者
      form: {},
      avatarList: [],
      coverList: [],
      // 表单校验
      rules: {
        name: [
          { required: true, message: "作者姓名不能为空", trigger: "blur" }
        ]
      }
    };
  },
  created() {
    this.getList();
  },
  methods: {
    /** 查询作者列表 */
    getList() {
      this.loading = true;
      listWriter(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
          this.writerList = response.data.rows;
          this.total = response.data.page.total;
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
    handleAvatarSuccess(res,file) {
      this.form.avatar = res.data.url;
      this.avatarList = [{name:res.data.url, url:this.parsePreviewSrc(res.data.url)}];
    },
    handleAvatarRemove() {
      this.form.avatar = "";
      this.avatarList = [];
    },
    handleCoverSuccess(res, file) {
      this.form.cover = res.data.url;
      this.coverList = [{name:res.data.url, url:this.parsePreviewSrc(res.data.url)}];
    },
    handleCoverRemove() {
      this.form.cover = "";
      this.coverList = [];
    },
    getSex(sex) {
      let str = "未知";
      this.dict.type.sys_user_sex.forEach(ele => {
        if (ele.value === sex) {
          str = ele.label;
        }
      })
      return str;
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
        keywords: undefined,
        description: undefined,
        remark: undefined,
        avatar: undefined,
        cover: undefined,
        inBanner: "0",
        inBannerSort: 0,
        inMenu: "0",
        inMenuSort: 0,
      };
      this.avatarList = [];
      this.coverList = [];
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
      this.open = true;
      this.title = "添加作者";
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
      getWriter(id).then(response => {
        this.form = response.data;
        if (this.form.avatar !== undefined && this.form.avatar.length > 0) {
          this.avatarList = [{name:this.form.cover, url: this.parsePreviewSrc(this.form.cover)}];
        }
        if (this.form.cover !== undefined && this.form.cover.length > 0) {
          this.coverList = [{name:this.form.cover, url: this.parsePreviewSrc(this.form.cover)}];
        }
        this.open = true;
        this.title = "修改作者";
      });
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs["form"].validate(valid => {
        if (valid) {
          this.form.sort = parseInt(this.form.sort);
          if (this.form.id !== undefined) {
            updateWriter(this.form).then(response => {
              this.$modal.msgSuccess("修改成功");
              this.open = false;
              this.getList();
            });
          } else {
            addWriter(this.form).then(response => {
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
      this.$modal.confirm('是否确认删除作者ID为"' + ids + '"的数据项？').then(function() {
          return delWriter(ids);
        }).then(() => {
          this.getList();
          this.$modal.msgSuccess("删除成功");
        }).catch(() => {});
    },
  }
};
</script>
