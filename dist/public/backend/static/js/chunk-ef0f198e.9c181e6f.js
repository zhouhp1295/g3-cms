(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-ef0f198e"],{"8dec":function(e,t,n){"use strict";n.r(t);var r=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"app-container"},[n("el-form",{directives:[{name:"show",rawName:"v-show",value:e.showSearch,expression:"showSearch"}],ref:"queryForm",attrs:{model:e.queryParams,inline:!0,"label-width":"68px"}},[n("el-form-item",{attrs:{label:"分类名称",prop:"title"}},[n("el-input",{staticStyle:{width:"240px"},attrs:{placeholder:"请输入分类名称",clearable:"",size:"small"},nativeOn:{keyup:function(t){return!t.type.indexOf("key")&&e._k(t.keyCode,"enter",13,t.key,"Enter")?null:e.handleQuery(t)}},model:{value:e.queryParams.title,callback:function(t){e.$set(e.queryParams,"title",t)},expression:"queryParams.title"}})],1),n("el-form-item",{attrs:{label:"菜单",prop:"inMenu"}},[n("el-select",{staticStyle:{width:"240px"},attrs:{placeholder:"菜单",clearable:"",size:"small"},model:{value:e.queryParams.inMenu,callback:function(t){e.$set(e.queryParams,"inMenu",t)},expression:"queryParams.inMenu"}},e._l(e.dict.type.sys_normal_yes_no,(function(e){return n("el-option",{key:e.value,attrs:{label:e.label,value:e.value}})})),1)],1),n("el-form-item",{attrs:{label:"Banner",prop:"inBanner"}},[n("el-select",{staticStyle:{width:"240px"},attrs:{placeholder:"Banner",clearable:"",size:"small"},model:{value:e.queryParams.inBanner,callback:function(t){e.$set(e.queryParams,"inBanner",t)},expression:"queryParams.inBanner"}},e._l(e.dict.type.sys_normal_yes_no,(function(e){return n("el-option",{key:e.value,attrs:{label:e.label,value:e.value}})})),1)],1),n("el-form-item",{attrs:{label:"关键字",prop:"keywords"}},[n("el-input",{staticStyle:{width:"240px"},attrs:{placeholder:"请输入关键字",clearable:"",size:"small"},nativeOn:{keyup:function(t){return!t.type.indexOf("key")&&e._k(t.keyCode,"enter",13,t.key,"Enter")?null:e.handleQuery(t)}},model:{value:e.queryParams.keywords,callback:function(t){e.$set(e.queryParams,"keywords",t)},expression:"queryParams.keywords"}})],1),n("el-form-item",{attrs:{label:"描述",prop:"description"}},[n("el-input",{staticStyle:{width:"240px"},attrs:{placeholder:"请输入描述",clearable:"",size:"small"},nativeOn:{keyup:function(t){return!t.type.indexOf("key")&&e._k(t.keyCode,"enter",13,t.key,"Enter")?null:e.handleQuery(t)}},model:{value:e.queryParams.description,callback:function(t){e.$set(e.queryParams,"description",t)},expression:"queryParams.description"}})],1),n("el-form-item",{attrs:{label:"创建时间"}},[n("el-date-picker",{staticStyle:{width:"240px"},attrs:{size:"small","value-format":"yyyy-MM-dd",type:"daterange","range-separator":"-","start-placeholder":"开始日期","end-placeholder":"结束日期"},model:{value:e.dateRange,callback:function(t){e.dateRange=t},expression:"dateRange"}})],1),n("el-form-item",[n("el-button",{attrs:{type:"primary",icon:"el-icon-search",size:"mini"},on:{click:e.handleQuery}},[e._v("搜索")]),n("el-button",{attrs:{icon:"el-icon-refresh",size:"mini"},on:{click:e.resetQuery}},[e._v("重置")])],1)],1),n("el-row",{staticClass:"mb8",attrs:{gutter:10}},[n("el-col",{attrs:{span:1.5}},[n("el-button",{directives:[{name:"hasPermi",rawName:"v-hasPermi",value:["content:category:add"],expression:"['content:category:add']"}],attrs:{type:"primary",plain:"",icon:"el-icon-plus",size:"mini"},on:{click:e.handleAdd}},[e._v("新增")])],1),n("el-col",{attrs:{span:1.5}},[n("el-button",{directives:[{name:"hasPermi",rawName:"v-hasPermi",value:["content:category:remove"],expression:"['content:category:remove']"}],attrs:{type:"danger",plain:"",icon:"el-icon-delete",size:"mini",disabled:e.multiple},on:{click:e.handleDelete}},[e._v("删除")])],1),n("el-col",{attrs:{span:1.5}},[n("el-button",{attrs:{type:"info",plain:"",icon:"el-icon-sort",size:"mini"},on:{click:e.toggleExpandAll}},[e._v("展开/折叠")])],1),n("right-toolbar",{attrs:{showSearch:e.showSearch},on:{"update:showSearch":function(t){e.showSearch=t},"update:show-search":function(t){e.showSearch=t},queryTable:e.getList}})],1),e.refreshTable?n("el-table",{directives:[{name:"loading",rawName:"v-loading",value:e.loading,expression:"loading"}],attrs:{data:e.categoryList,"row-key":"id","default-expand-all":e.isExpandAll,"tree-props":{children:"children",hasChildren:"hasChildren"}},on:{"selection-change":e.handleSelectionChange}},[n("el-table-column",{attrs:{type:"selection",width:"30",align:"center"}}),n("el-table-column",{attrs:{label:"ID",align:"center",prop:"id"}}),n("el-table-column",{attrs:{label:"",align:"center",prop:"cover"},scopedSlots:e._u([{key:"default",fn:function(t){return[t.row.cover?n("img",{staticClass:"table-cover",attrs:{src:e.parsePreviewSrc(t.row.cover)}}):e._e()]}}],null,!1,2273036629)}),n("el-table-column",{attrs:{label:"标题",align:"center",prop:"title","show-overflow-tooltip":!0}}),n("el-table-column",{attrs:{label:"Icon",align:"center",prop:"icon"}}),n("el-table-column",{key:"inMenu",attrs:{label:"显示到菜单",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[n("el-switch",{attrs:{"active-value":"1","inactive-value":"0"},on:{change:function(n){return e.handleInMenuChange(t.row)}},model:{value:t.row.inMenu,callback:function(n){e.$set(t.row,"inMenu",n)},expression:"scope.row.inMenu"}})]}}],null,!1,2593597854)}),n("el-table-column",{key:"inBanner",attrs:{label:"显示到Banner",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[n("el-switch",{attrs:{"active-value":"1","inactive-value":"0"},on:{change:function(n){return e.handleInBannerChange(t.row)}},model:{value:t.row.inBanner,callback:function(n){e.$set(t.row,"inBanner",n)},expression:"scope.row.inBanner"}})]}}],null,!1,162392318)}),n("el-table-column",{attrs:{label:"创建时间",align:"center",prop:"createdAt",width:"180"},scopedSlots:e._u([{key:"default",fn:function(t){return[n("span",[e._v(e._s(e.parseTime(t.row.createdAt)))])]}}],null,!1,147779282)}),n("el-table-column",{attrs:{label:"操作",align:"left","class-name":"small-padding fixed-width"},scopedSlots:e._u([{key:"default",fn:function(t){return[n("el-button",{directives:[{name:"hasPermi",rawName:"v-hasPermi",value:["content:category:edit"],expression:"['content:category:edit']"}],attrs:{size:"mini",type:"text",icon:"el-icon-edit"},on:{click:function(n){return e.handleUpdate(t.row)}}},[e._v("修改")]),n("el-button",{directives:[{name:"hasPermi",rawName:"v-hasPermi",value:["content:category:remove"],expression:"['content:category:remove']"}],attrs:{size:"mini",type:"text",icon:"el-icon-delete"},on:{click:function(n){return e.handleDelete(t.row)}}},[e._v("删除")])]}}],null,!1,3813150474)})],1):e._e(),n("el-dialog",{attrs:{title:e.title,visible:e.open,width:"800px","append-to-body":""},on:{"update:visible":function(t){e.open=t}}},[n("el-form",{ref:"form",attrs:{model:e.form,rules:e.rules,"label-width":"160px"}},[n("el-form-item",{attrs:{label:"上级分类",prop:"pid"}},[n("treeselect",{attrs:{options:e.categoryOptions,normalizer:e.normalizer,placeholder:"选择上级分类"},model:{value:e.form.pid,callback:function(t){e.$set(e.form,"pid",t)},expression:"form.pid"}})],1),n("el-form-item",{attrs:{label:"分类名称",prop:"title"}},[n("el-input",{attrs:{placeholder:"请输入分类名称"},model:{value:e.form.title,callback:function(t){e.$set(e.form,"title",t)},expression:"form.title"}})],1),n("el-form-item",{attrs:{label:"Icon",prop:"icon"}},[n("el-input",{attrs:{placeholder:"请输入icon"},model:{value:e.form.icon,callback:function(t){e.$set(e.form,"icon",t)},expression:"form.icon"}})],1),n("el-form-item",{attrs:{label:"图片",prop:"cover"}},[n("el-upload",{staticClass:"cover-uploader",attrs:{action:e.uploadUrl,multiple:!1,"file-list":e.coverList,"on-success":e.handleCoverSuccess,"on-remove":e.handleCoverRemove,headers:e.headers,"before-upload":e.beforeUpload}},[e.form.cover?n("div",[n("img",{staticClass:"cover",attrs:{src:e.parsePreviewSrc(e.form.cover)}})]):n("i",{staticClass:"el-icon-plus cover-uploader-icon"})])],1),n("el-form-item",{attrs:{label:"Banner图片",prop:"banner"}},[n("el-col",[e._v(" 建议尺寸: 820 x 200 px ")]),n("el-upload",{staticClass:"banner-uploader",attrs:{action:e.uploadUrl,multiple:!1,"file-list":e.bannerList,"on-success":e.handleBannerSuccess,"on-remove":e.handleBannerRemove,headers:e.headers,"before-upload":e.beforeUpload}},[e.form.banner?n("div",[n("img",{staticClass:"banner",attrs:{src:e.parsePreviewSrc(e.form.banner)}})]):n("i",{staticClass:"el-icon-plus banner-uploader-icon"})])],1),n("el-form-item",{attrs:{label:"关键字",prop:"keywords"}},[n("el-input",{attrs:{type:"textarea",placeholder:"请输入分类代码"},model:{value:e.form.keywords,callback:function(t){e.$set(e.form,"keywords",t)},expression:"form.keywords"}})],1),n("el-form-item",{attrs:{label:"描述",prop:"description"}},[n("el-input",{attrs:{type:"textarea",placeholder:"请输入分类值"},model:{value:e.form.description,callback:function(t){e.$set(e.form,"description",t)},expression:"form.description"}})],1),n("el-form-item",{attrs:{label:"在菜单中显示"}},[n("el-radio-group",{model:{value:e.form.inMenu,callback:function(t){e.$set(e.form,"inMenu",t)},expression:"form.inMenu"}},e._l(e.dict.type.sys_normal_yes_no,(function(t){return n("el-radio",{key:t.value,attrs:{label:t.value}},[e._v(e._s(t.label))])})),1)],1),1==e.form.inMenu?n("el-form-item",{attrs:{label:"在菜单中排序"}},[n("el-input",{attrs:{placeholder:"请输入",type:"number"},model:{value:e.form.inMenuSort,callback:function(t){e.$set(e.form,"inMenuSort",t)},expression:"form.inMenuSort"}})],1):e._e(),n("el-form-item",{attrs:{label:"在Banner显示"}},[n("el-radio-group",{model:{value:e.form.inBanner,callback:function(t){e.$set(e.form,"inBanner",t)},expression:"form.inBanner"}},e._l(e.dict.type.sys_normal_yes_no,(function(t){return n("el-radio",{key:t.value,attrs:{label:t.value}},[e._v(e._s(t.label))])})),1)],1),1==e.form.inBanner?n("el-form-item",{attrs:{label:"在Banner中排序"}},[n("el-input",{attrs:{placeholder:"请输入",type:"number"},model:{value:e.form.inBannerSort,callback:function(t){e.$set(e.form,"inBannerSort",t)},expression:"form.inBannerSort"}})],1):e._e(),n("el-form-item",{attrs:{label:"备注",prop:"remark"}},[n("el-input",{attrs:{type:"textarea",placeholder:"请输入内容"},model:{value:e.form.remark,callback:function(t){e.$set(e.form,"remark",t)},expression:"form.remark"}})],1)],1),n("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[n("el-button",{attrs:{type:"primary"},on:{click:e.submitForm}},[e._v("确 定")]),n("el-button",{on:{click:e.cancel}},[e._v("取 消")])],1)],1)],1)},a=[],i=(n("d81d"),n("4e82"),n("90fa")),o=n("ca17"),l=n.n(o),s=(n("542c"),n("5f87")),c={name:"Category",dicts:["sys_normal_yes_no"],components:{Treeselect:l.a},data:function(){return{loading:!0,isExpandAll:!0,refreshTable:!0,ids:[],single:!0,multiple:!0,showSearch:!0,total:0,categoryList:[],categoryOptions:[],title:"",open:!1,dateRange:[],uploadUrl:"/api/common/uploadImage",headers:{Authorization:"Bearer "+Object(s["a"])()},queryParams:{title:void 0,keywords:void 0,description:void 0,inMenu:void 0,inBanner:void 0},form:{},coverList:[],bannerList:[],rules:{title:[{required:!0,message:"分类代码不能为空",trigger:"blur"}]}}},created:function(){this.getList()},methods:{getList:function(){var e=this;this.loading=!0,Object(i["f"])(this.addDateRange(this.queryParams,this.dateRange)).then((function(t){e.categoryList=e.handleTree(t.data.rows,"id","pid"),e.loading=!1}))},beforeUpload:function(e){return"image/jpeg"!==e.type&&"image/png"!==e.type?(this.$message.error("上传图片只能是 JPG、PNG 格式!"),!1):!(e.size/1024/1024>2)||(this.$message.error("上传图片大小不能超过 2MB!"),!1)},handleCoverSuccess:function(e,t){this.form.cover=e.data.url,this.coverList=[{name:e.data.url,url:this.parsePreviewSrc(e.data.url)}]},handleCoverRemove:function(){this.form.cover="",this.coverList=[]},handleBannerSuccess:function(e,t){this.form.banner=e.data.url,this.bannerList=[{name:e.data.url,url:this.parsePreviewSrc(e.data.url)}]},handleBannerRemove:function(){this.form.banner="",this.bannerList=[]},normalizer:function(e){return e.children&&!e.children.length&&delete e.children,{id:e.id,label:e.title,children:e.children}},toggleExpandAll:function(){var e=this;this.refreshTable=!1,this.isExpandAll=!this.isExpandAll,this.$nextTick((function(){e.refreshTable=!0}))},cancel:function(){this.open=!1,this.reset()},reset:function(){this.form={id:void 0,pid:void 0,title:void 0,icon:void 0,keywords:void 0,description:void 0,remark:void 0,cover:void 0,banner:void 0,inBanner:"0",inBannerSort:0,inMenu:"0",inMenuSort:0},this.coverList=[],this.bannerList=[],this.resetForm("form")},handleQuery:function(){this.getList()},resetQuery:function(){this.dateRange=[],this.resetForm("queryForm"),this.handleQuery()},handleAdd:function(){var e=this;this.reset(),this.open=!0,this.title="添加分类",Object(i["f"])().then((function(t){var n=t.data.rows;e.categoryOptions=e.handleTree(n,"id","pid")}))},handleSelectionChange:function(e){this.ids=e.map((function(e){return e.id})),this.single=1!==e.length,this.multiple=!e.length},handleUpdate:function(e){var t=this;this.reset();var n=e.id||this.ids;Object(i["e"])(n).then((function(e){t.form=e.data,void 0!==t.form.cover&&t.form.cover.length>0&&(t.coverList=[{name:t.form.cover,url:t.parsePreviewSrc(t.form.cover)}]),void 0!==t.form.banner&&t.form.banner.length>0&&(t.bannerList=[{name:t.form.banner,url:t.parsePreviewSrc(t.form.banner)}]),t.open=!0,t.title="修改分类"})),Object(i["f"])({excludeId:e.id}).then((function(e){var n=e.data.rows;t.categoryOptions=t.handleTree(n,"id","pid")}))},submitForm:function(){var e=this;this.$refs["form"].validate((function(t){t&&(e.form.inMenuSort=parseInt(e.form.inMenuSort),e.form.inBannerSort=parseInt(e.form.inBannerSort),e.form.sort=parseInt(e.form.sort),void 0!==e.form.id?Object(i["h"])(e.form).then((function(t){e.$modal.msgSuccess("修改成功"),e.open=!1,e.getList()})):Object(i["a"])(e.form).then((function(t){e.$modal.msgSuccess("新增成功"),e.open=!1,e.getList()})))}))},handleDelete:function(e){var t=this,n=e.id||this.ids;this.$modal.confirm('是否确认删除分类ID为"'+n+'"的数据项？').then((function(){return Object(i["d"])(n)})).then((function(){t.getList(),t.$modal.msgSuccess("删除成功")})).catch((function(){}))},handleInMenuChange:function(e){var t=this,n="1"===e.inMenu?"设置显示到菜单":"取消显示到菜单";"1"!==e.inMenu?this.$modal.confirm("确认要将【"+e.title+"】"+n+"吗？").then((function(){return Object(i["c"])({id:e.id,inMenu:e.inMenu,inMenuSort:0})})).then((function(){t.$modal.msgSuccess(n+"成功")})).catch((function(){e.inMenu="0"===e.inMenu?"1":"0"})):this.$prompt("确认要将【"+e.title+"】"+n+"吗？请输入排序").then((function(t){var n=t.value,r=parseInt(n);return Object(i["c"])({id:e.id,inMenu:e.inMenu,inMenuSort:isNaN(r)?e.inMenuSort:r})})).then((function(){t.$modal.msgSuccess(n+"成功")})).catch((function(){e.inMenu="0"===e.inMenu?"1":"0"}))},handleInBannerChange:function(e){var t=this,n="1"===e.inBanner?"设置显示到Banner":"取消显示到Banner";"1"!==e.inBanner?this.$modal.confirm("确认要将【"+e.title+"】"+n+"吗？").then((function(){return Object(i["b"])({id:e.id,inBanner:e.inBanner,inBannerSort:0})})).then((function(){t.$modal.msgSuccess(n+"成功")})).catch((function(){e.inBanner="0"===e.inBanner?"1":"0"})):this.$prompt("确认要将【"+e.title+"】"+n+"吗？请输入排序").then((function(t){var n=t.value,r=parseInt(n);return Object(i["b"])({id:e.id,inBanner:e.inBanner,inBannerSort:isNaN(r)?e.inBannerSort:r})})).then((function(){t.$modal.msgSuccess(n+"成功")})).catch((function(){e.inBanner="0"===e.inBanner?"1":"0"}))}}},u=c,d=(n("e85e"),n("2877")),m=Object(d["a"])(u,r,a,!1,null,null,null);t["default"]=m.exports},"90fa":function(e,t,n){"use strict";n.d(t,"f",(function(){return a})),n.d(t,"e",(function(){return i})),n.d(t,"a",(function(){return o})),n.d(t,"h",(function(){return l})),n.d(t,"d",(function(){return s})),n.d(t,"c",(function(){return c})),n.d(t,"b",(function(){return u}));var r=n("b775");function a(e){return Object(r["a"])({url:"/api/admin/content/category/list",method:"get",params:e})}function i(e){return Object(r["a"])({url:"/api/admin/content/category/get",method:"get",params:{id:e}})}function o(e){return Object(r["a"])({url:"/api/admin/content/category/insert",method:"post",data:e})}function l(e){return Object(r["a"])({url:"/api/admin/content/category/update",method:"put",data:e})}function s(e){return Object(r["a"])({url:"/api/admin/content/category/delete",method:"delete",params:{id:e}})}function c(e){return Object(r["a"])({url:"/api/admin/content/category/inMenu",method:"put",data:e})}function u(e){return Object(r["a"])({url:"/api/admin/content/category/inBanner",method:"put",data:e})}},e6da:function(e,t,n){},e85e:function(e,t,n){"use strict";n("e6da")}}]);