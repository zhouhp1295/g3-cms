import request from '@/utils/request'

// 查询参数列表
export function listConfig(query) {
  return request({
    url: '/api/admin/system/config/page',
    method: 'get',
    params: query
  })
}

// 查询参数详细
export function getConfig(id) {
  return request({
    url: '/api/admin/system/config/get',
    method: 'get',
    params: {
      id
    }
  })
}

// 根据参数键名查询参数值
export function getConfigByCode(code) {
  return request({
    url: '/api/admin/system/config/getByCode',
    method: 'get',
    params: {
      code
    }
  })
}

// 新增参数配置
export function addConfig(data) {
  return request({
    url: '/api/admin/system/config/insert',
    method: 'post',
    data: data
  })
}

// 修改参数配置
export function updateConfig(data) {
  return request({
    url: '/api/admin/system/config/update',
    method: 'put',
    data: data
  })
}

// 删除参数配置
export function delConfig(id) {
  return request({
    url: '/api/admin/system/config/delete',
    method: 'delete',
    params: {
      id
    }
  })
}

// 刷新参数缓存
export function refreshCache() {
  return request({
    url: '/system/config/refreshCache',
    method: 'delete'
  })
}

// 角色状态修改
export function changeConfigStatus(id, status) {
  const data = {
    id,
    status
  }
  return request({
    url: '/api/admin/system/config/status',
    method: 'put',
    data: data
  })
}
