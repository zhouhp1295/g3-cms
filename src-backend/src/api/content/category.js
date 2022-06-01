import request from '@/utils/request'

// 查询分类列表
export function listCategory(query) {
  return request({
    url: '/api/admin/content/category/list',
    method: 'get',
    params: query
  })
}

// 查询分类详细
export function getCategory(id) {
  return request({
    url: '/api/admin/content/category/get',
    method: 'get',
    params: {
      id
    }
  })
}

// 新增分类配置
export function addCategory(data) {
  return request({
    url: '/api/admin/content/category/insert',
    method: 'post',
    data: data
  })
}

// 修改分类配置
export function updateCategory(data) {
  return request({
    url: '/api/admin/content/category/update',
    method: 'put',
    data: data
  })
}

// 删除分类配置
export function delCategory(id) {
  return request({
    url: '/api/admin/content/category/delete',
    method: 'delete',
    params: {
      id
    }
  })
}

// 显示在菜单
export function changeCategoryInMenu(data) {
  return request({
    url: '/api/admin/content/category/inMenu',
    method: 'put',
    data: data
  })
}

// 显示在Banner
export function changeCategoryInBanner(data) {
  return request({
    url: '/api/admin/content/category/inBanner',
    method: 'put',
    data: data
  })
}
