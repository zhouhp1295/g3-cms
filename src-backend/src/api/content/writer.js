import request from '@/utils/request'

export function listWriterOptions(query) {
  return request({
    url: '/api/admin/content/writer/selectOptions',
    method: 'get',
    params: query
  })
}

// 查询作者列表
export function listWriter(query) {
  return request({
    url: '/api/admin/content/writer/page',
    method: 'get',
    params: query
  })
}

// 查询作者详细
export function getWriter(id) {
  return request({
    url: '/api/admin/content/writer/get',
    method: 'get',
    params: {
      id
    }
  })
}

// 新增作者配置
export function addWriter(data) {
  return request({
    url: '/api/admin/content/writer/insert',
    method: 'post',
    data: data
  })
}

// 修改作者配置
export function updateWriter(data) {
  return request({
    url: '/api/admin/content/writer/update',
    method: 'put',
    data: data
  })
}

// 删除作者配置
export function delWriter(id) {
  return request({
    url: '/api/admin/content/writer/delete',
    method: 'delete',
    params: {
      id
    }
  })
}


