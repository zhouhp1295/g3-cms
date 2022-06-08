import request from '@/utils/request'

// 清空缓存
export function refreshCache() {
  return request({
    url: '/api/admin/content/config/cache/clean',
    method: 'put'
  })
}

// 查询
export function getConfig() {
  return request({
    url: '/api/admin/content/config/web/get',
    method: 'get'
  })
}

// 修改
export function saveConfig(data) {
  return request({
    url: '/api/admin/content/config/web/update',
    method: 'put',
    data: data
  })
}
