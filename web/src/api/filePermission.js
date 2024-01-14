import service from '@/utils/request'

// @Tags FilePermission
// @Summary 创建描述文件权限的信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FilePermission true "创建描述文件权限的信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /FP/createFilePermission [post]
export const createFilePermission = (data) => {
  return service({
    url: '/FP/createFilePermission',
    method: 'post',
    data
  })
}

// @Tags FilePermission
// @Summary 删除描述文件权限的信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FilePermission true "删除描述文件权限的信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /FP/deleteFilePermission [delete]
export const deleteFilePermission = (data) => {
  return service({
    url: '/FP/deleteFilePermission',
    method: 'delete',
    data
  })
}

// @Tags FilePermission
// @Summary 批量删除描述文件权限的信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除描述文件权限的信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /FP/deleteFilePermission [delete]
export const deleteFilePermissionByIds = (data) => {
  return service({
    url: '/FP/deleteFilePermissionByIds',
    method: 'delete',
    data
  })
}

// @Tags FilePermission
// @Summary 更新描述文件权限的信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FilePermission true "更新描述文件权限的信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /FP/updateFilePermission [put]
export const updateFilePermission = (data) => {
  return service({
    url: '/FP/updateFilePermission',
    method: 'put',
    data
  })
}

// @Tags FilePermission
// @Summary 用id查询描述文件权限的信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FilePermission true "用id查询描述文件权限的信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /FP/findFilePermission [get]
export const findFilePermission = (params) => {
  return service({
    url: '/FP/findFilePermission',
    method: 'get',
    params
  })
}

// @Tags FilePermission
// @Summary 分页获取描述文件权限的信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取描述文件权限的信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /FP/getFilePermissionList [get]
export const getFilePermissionList = (params) => {
  return service({
    url: '/FP/getFilePermissionList',
    method: 'get',
    params
  })
}
