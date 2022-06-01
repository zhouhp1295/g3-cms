import request from '@/utils/request'

// 查询标签列表
export function listFriendLink(query) {
  return request({
    url: '/api/admin/content/friendLink/page',
    method: 'get',
    params: query
  })
}

// 查询标签详细
export function getFriendLink(id) {
  return request({
    url: '/api/admin/content/friendLink/get',
    method: 'get',
    params: {
      id
    }
  })
}

// 新增标签配置
export function addFriendLink(data) {
  return request({
    url: '/api/admin/content/friendLink/insert',
    method: 'post',
    data: data
  })
}

// 修改标签配置
export function updateFriendLink(data) {
  return request({
    url: '/api/admin/content/friendLink/update',
    method: 'put',
    data: data
  })
}

// 删除标签配置
export function delFriendLink(id) {
  return request({
    url: '/api/admin/content/friendLink/delete',
    method: 'delete',
    params: {
      id
    }
  })
}

// 推荐状态修改
export function changeFriendLinkStatus(id, status) {
  const data = {
    id,
    status
  }
  return request({
    url: '/api/admin/content/friendLink/status',
    method: 'put',
    data: data
  })
}
