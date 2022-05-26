import request from '@/utils/request'

// 查询角色列表
export function listRoleOptions() {
  return request({
    url: '/api/admin/system/role/listOptions',
    method: 'get',
  })
}

// 查询角色列表
export function listRole(query) {
  return request({
    url: '/api/admin/system/role/page',
    method: 'get',
    params: query
  })
}

// 查询角色详细
export function getRole(roleId) {
  return request({
    url: '/api/admin/system/role/get',
    method: 'get',
    params: {id : roleId}
  })
}

// 新增角色
export function addRole(data) {
  return request({
    url: '/api/admin/system/role/insert',
    method: 'post',
    data: data
  })
}

// 修改角色
export function updateRole(data) {
  return request({
    url: '/api/admin/system/role/update',
    method: 'put',
    data: data
  })
}

// 角色状态修改
export function changeRoleStatus(id, status) {
  const data = {
    id,
    status
  }
  return request({
    url: '/api/admin/system/role/status',
    method: 'put',
    data: data
  })
}

// 删除角色
export function delRole(roleId) {
  return request({
    url: '/api/admin/system/role/delete',
    method: 'delete',
    params: {id : roleId}
  })
}
