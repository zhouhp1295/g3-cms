(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-58bd45b8"],{9429:function(t,e,o){"use strict";o.r(e);var r=function(){var t=this,e=t.$createElement,o=t._self._c||e;return o("div",[o("div",{staticClass:"user-info-head",on:{click:function(e){return t.editCropper()}}},[o("img",{staticClass:"img-circle img-lg",attrs:{src:t.options.img,title:"点击上传头像"}})]),o("el-dialog",{attrs:{title:t.title,visible:t.open,width:"800px","append-to-body":""},on:{"update:visible":function(e){t.open=e},opened:t.modalOpened,close:t.closeDialog}},[o("el-row",[o("el-col",{style:{height:"350px"},attrs:{xs:24,md:12}},[t.visible?o("vue-cropper",{ref:"cropper",attrs:{img:t.options.img,info:!0,autoCrop:t.options.autoCrop,autoCropWidth:t.options.autoCropWidth,autoCropHeight:t.options.autoCropHeight,fixedBox:t.options.fixedBox},on:{realTime:t.realTime}}):t._e()],1),o("el-col",{style:{height:"350px"},attrs:{xs:24,md:12}},[o("div",{staticClass:"avatar-upload-preview"},[o("img",{style:t.previews.img,attrs:{src:t.previews.url}})])])],1),o("br"),o("el-row",[o("el-col",{attrs:{lg:2,md:2}},[o("el-upload",{attrs:{action:"#","http-request":t.requestUpload,"show-file-list":!1,"before-upload":t.beforeUpload}},[o("el-button",{attrs:{size:"small"}},[t._v(" 选择 "),o("i",{staticClass:"el-icon-upload el-icon--right"})])],1)],1),o("el-col",{attrs:{lg:{span:1,offset:2},md:2}},[o("el-button",{attrs:{icon:"el-icon-plus",size:"small"},on:{click:function(e){return t.changeScale(1)}}})],1),o("el-col",{attrs:{lg:{span:1,offset:1},md:2}},[o("el-button",{attrs:{icon:"el-icon-minus",size:"small"},on:{click:function(e){return t.changeScale(-1)}}})],1),o("el-col",{attrs:{lg:{span:1,offset:1},md:2}},[o("el-button",{attrs:{icon:"el-icon-refresh-left",size:"small"},on:{click:function(e){return t.rotateLeft()}}})],1),o("el-col",{attrs:{lg:{span:1,offset:1},md:2}},[o("el-button",{attrs:{icon:"el-icon-refresh-right",size:"small"},on:{click:function(e){return t.rotateRight()}}})],1),o("el-col",{attrs:{lg:{span:2,offset:6},md:2}},[o("el-button",{attrs:{type:"primary",size:"small"},on:{click:function(e){return t.uploadImg()}}},[t._v("提 交")])],1)],1)],1)],1)},i=[],n=o("4360"),a=o("7e79"),s=o("c0c7"),u=o("c38a"),c={components:{VueCropper:a["VueCropper"]},props:{user:{type:Object}},data:function(){return{open:!1,visible:!1,title:"修改头像",options:{img:n["a"].getters.avatar,autoCrop:!0,autoCropWidth:200,autoCropHeight:200,fixedBox:!0},previews:{}}},methods:{editCropper:function(){this.open=!0},modalOpened:function(){this.visible=!0},requestUpload:function(){},rotateLeft:function(){this.$refs.cropper.rotateLeft()},rotateRight:function(){this.$refs.cropper.rotateRight()},changeScale:function(t){t=t||1,this.$refs.cropper.changeScale(t)},beforeUpload:function(t){var e=this;if(-1==t.type.indexOf("image/"))this.$modal.msgError("文件格式错误，请上传图片类型,如：JPG，PNG后缀的文件。");else{var o=new FileReader;o.readAsDataURL(t),o.onload=function(){e.options.img=o.result}}},uploadImg:function(){var t=this;this.$refs.cropper.getCropBlob((function(e){var o=new FormData;o.append("avatar",e),Object(s["k"])(o).then((function(e){t.open=!1,t.options.img=Object(u["e"])(e.data.avatar),n["a"].commit("SET_AVATAR",t.options.img),t.$modal.msgSuccess("修改成功"),t.visible=!1}))}))},realTime:function(t){this.previews=t},closeDialog:function(){this.options.img=n["a"].getters.avatar,this.visible=!1}}},l=c,p=(o("f5e2"),o("2877")),d=Object(p["a"])(l,r,i,!1,null,"7c8969a6",null);e["default"]=d.exports},"9f84":function(t,e,o){},c0c7:function(t,e,o){"use strict";o.d(e,"f",(function(){return n})),o.d(e,"d",(function(){return a})),o.d(e,"a",(function(){return s})),o.d(e,"h",(function(){return u})),o.d(e,"c",(function(){return c})),o.d(e,"g",(function(){return l})),o.d(e,"b",(function(){return p})),o.d(e,"e",(function(){return d})),o.d(e,"i",(function(){return f})),o.d(e,"j",(function(){return m})),o.d(e,"k",(function(){return h}));var r=o("b775"),i=o("c38a");function n(t){return Object(r["a"])({url:"/api/admin/system/user/page",method:"get",params:t})}function a(t){return Object(r["a"])({url:"/api/admin/system/user/get",method:"get",params:{id:Object(i["f"])(t)}})}function s(t){return Object(r["a"])({url:"/api/admin/system/user/insert",method:"post",data:t})}function u(t){return Object(r["a"])({url:"/api/admin/system/user/update",method:"put",data:t})}function c(t){return Object(r["a"])({url:"/api/admin/system/user/delete",method:"delete",params:{id:Object(i["f"])(t)}})}function l(t,e){var o={id:t,password:e};return Object(r["a"])({url:"/system/user/password/reset",method:"put",data:o})}function p(t,e){var o={id:t,status:e};return Object(r["a"])({url:"/api/admin/system/user/status",method:"put",data:o})}function d(){return Object(r["a"])({url:"/api/admin/system/user/profile/get",method:"get"})}function f(t){return Object(r["a"])({url:"/api/admin/system/user/profile/update",method:"put",data:t})}function m(t,e){var o={oldPassword:t,newPassword:e};return Object(r["a"])({url:"/api/admin/system/user/profile/password",method:"put",data:o})}function h(t){return Object(r["a"])({url:"/api/admin/system/user/profile/avatar",method:"post",data:t})}},f5e2:function(t,e,o){"use strict";o("9f84")}}]);