import request from '@/utils/request'

// 查询
export function refreshCache() {
  return request({
    url: '/api/admin/content/config/refreshCache',
    method: 'get'
  })
}

// 查询
export function getConfig() {
  return request({
    url: '/api/admin/content/config/crud',
    method: 'get'
  })
}

// 修改
export function saveConfig(data) {
  return request({
    url: '/api/admin/content/config/crud',
    method: 'put',
    data: data
  })
}
