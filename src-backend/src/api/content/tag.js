import request from '@/utils/request'

export function listTagSimple(query) {
  return request({
    url: '/api/admin/content/tag/listSimple',
    method: 'get',
    params: query
  })
}

// 查询标签列表
export function listTag(query) {
  return request({
    url: '/api/admin/content/tag/page',
    method: 'get',
    params: query
  })
}

// 查询标签详细
export function getTag(id) {
  return request({
    url: '/api/admin/content/tag/get',
    method: 'get',
    params: {
      id
    }
  })
}

// 新增标签配置
export function addTag(data) {
  return request({
    url: '/api/admin/content/tag/insert',
    method: 'post',
    data: data
  })
}

export function fastAddTag(data) {
  return request({
    url: '/api/admin/content/tag/fastInsert',
    method: 'post',
    data: data
  })
}

// 修改标签配置
export function updateTag(data) {
  return request({
    url: '/api/admin/content/tag/update',
    method: 'put',
    data: data
  })
}

// 删除标签配置
export function delTag(id) {
  return request({
    url: '/api/admin/content/tag/delete',
    method: 'delete',
    params: {
      id
    }
  })
}

