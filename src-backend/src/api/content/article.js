import request from '@/utils/request'

// 查询标签列表
export function listArticle(query) {
  return request({
    url: '/api/admin/content/article/page',
    method: 'get',
    params: query
  })
}

// 查询标签详细
export function getArticle(id) {
  return request({
    url: '/api/admin/content/article/get',
    method: 'get',
    params: {
      id
    }
  })
}

// 新增标签配置
export function addArticle(data) {
  return request({
    url: '/api/admin/content/article/insert',
    method: 'post',
    data: data
  })
}

// 修改标签配置
export function updateArticle(data) {
  return request({
    url: '/api/admin/content/article/update',
    method: 'put',
    data: data
  })
}

// 删除标签配置
export function delArticle(id) {
  return request({
    url: '/api/admin/content/article/delete',
    method: 'delete',
    params: {
      id
    }
  })
}

// 发布状态修改
export function changeArticleStatus(id, status) {
  const data = {
    id,
    status
  }
  return request({
    url: '/api/admin/content/article/status',
    method: 'put',
    data: data
  })
}

// 显示在Banner
export function changeArticleInBanner(data) {
  return request({
    url: '/api/admin/content/article/inBanner',
    method: 'put',
    data: data
  })
}

