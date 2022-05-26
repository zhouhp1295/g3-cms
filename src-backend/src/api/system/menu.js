import request from '@/utils/request'

// 查询菜单下拉树结构
export function treeselect() {
  return request({
    url: '/api/admin/system/menu/treeOptions',
    method: 'get'
  })
}
