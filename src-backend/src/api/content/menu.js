import request from '@/utils/request'

// 查询分类列表
export function listMenu(query) {
  return request({
    url: '/api/admin/content/menu/list',
    method: 'get',
    params: query
  })
}

// 查询分类详细
export function getMenu(id) {
  return request({
    url: '/api/admin/content/menu/get',
    method: 'get',
    params: {
      id
    }
  })
}

// 新增分类配置
export function addMenu(data) {
  return request({
    url: '/api/admin/content/menu/insert',
    method: 'post',
    data: data
  })
}

// 修改分类配置
export function updateMenu(data) {
  return request({
    url: '/api/admin/content/menu/update',
    method: 'put',
    data: data
  })
}

// 删除分类配置
export function delMenu(id) {
  return request({
    url: '/api/admin/content/menu/delete',
    method: 'delete',
    params: {
      id
    }
  })
}


// 修改
export function changeMenuIsBlank(id, isBlank) {
  const data = {
    id,
    isBlank
  }
  return request({
    url: '/api/admin/content/menu/isBlank',
    method: 'put',
    data: data
  })
}

// 修改
export function changeMenuStatus(id, status) {
  const data = {
    id,
    status
  }
  return request({
    url: '/api/admin/content/menu/status',
    method: 'put',
    data: data
  })
}
