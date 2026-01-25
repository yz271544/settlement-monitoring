package api

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/yz271544/settlement-monitoring/v10/internal/mods/guard/biz"
	"github.com/yz271544/settlement-monitoring/v10/internal/mods/guard/schema"
	"github.com/yz271544/settlement-monitoring/v10/pkg/util"
)

type Staff struct {
	StaffBiz *biz.Staff
}

// Query 查询工作人员列表
func (a *Staff) Query(c *gin.Context) {
	var req schema.StaffListReq
	if err := util.ParseQuery(c, &req); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.StaffBiz.Query(c.Request.Context(), &req)
	if err != nil {
		util.ResError(c, err)
		return
	}

	pr := &util.PaginationResult{
		Total:    result.Total,
		Current:  req.Page,
		PageSize: req.PageSize,
	}
	util.ResPage(c, result.Items, pr)
}

// Get 获取工作人员详情
func (a *Staff) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		util.ResError(c, errors.New("id不能为空"))
		return
	}

	result, err := a.StaffBiz.Get(c.Request.Context(), id)
	if err != nil {
		util.ResError(c, err)
		return
	}

	util.ResSuccess(c, result)
}

// Create 创建工作人员
func (a *Staff) Create(c *gin.Context) {
	var req schema.StaffCreateReq
	if err := util.ParseJSON(c, &req); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.StaffBiz.Create(c.Request.Context(), &req)
	if err != nil {
		util.ResError(c, err)
		return
	}

	util.ResSuccess(c, result)
}

// Update 更新工作人员
func (a *Staff) Update(c *gin.Context) {
	id := c.Param("id")
	var req schema.StaffUpdateReq
	if err := util.ParseJSON(c, &req); err != nil {
		util.ResError(c, err)
		return
	}

	req.ID = id
	if err := a.StaffBiz.Update(c.Request.Context(), &req); err != nil {
		util.ResError(c, err)
		return
	}

	util.ResOK(c)
}

// Delete 删除工作人员
func (a *Staff) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		util.ResError(c, errors.New("id不能为空"))
		return
	}

	if err := a.StaffBiz.Delete(c.Request.Context(), []string{id}); err != nil {
		util.ResError(c, err)
		return
	}

	util.ResOK(c)
}

// Import 导入工作人员数据
func (a *Staff) Import(c *gin.Context) {
	// TODO: 实现 Excel 导入功能
	util.ResOK(c)
}

// Export 导出工作人员数据
func (a *Staff) Export(c *gin.Context) {
	// TODO: 实现 Excel 导出功能
	util.ResOK(c)
}

// ExportTemplate 导出模板
func (a *Staff) ExportTemplate(c *gin.Context) {
	// TODO: 实现模板下载功能
	util.ResOK(c)
}
