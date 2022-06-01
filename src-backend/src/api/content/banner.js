import request from '@/utils/request'

// 查询标签列表
export function listBanner(query) {
  return request({
    url: '/api/admin/content/banner/page',
    method: 'get',
    params: query
  })
}

// 查询标签详细
export function getBanner(id) {
  return request({
    url: '/api/admin/content/banner/get',
    method: 'get',
    params: {
      id
    }
  })
}

// 新增标签配置
export function addBanner(data) {
  return request({
    url: '/api/admin/content/banner/insert',
    method: 'post',
    data: data
  })
}

// 修改标签配置
export function updateBanner(data) {
  return request({
    url: '/api/admin/content/banner/update',
    method: 'put',
    data: data
  })
}

// 删除标签配置
export function delBanner(id) {
  return request({
    url: '/api/admin/content/banner/delete',
    method: 'delete',
    params: {
      id
    }
  })
}

// 推荐状态修改
export function changeBannerStatus(id, status) {
  const data = {
    id,
    status
  }
  return request({
    url: '/api/admin/content/banner/status',
    method: 'put',
    data: data
  })
}
