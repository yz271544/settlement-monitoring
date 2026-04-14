package api

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/yz271544/settlement-monitoring/v10/internal/mods/guard/biz"
	"github.com/yz271544/settlement-monitoring/v10/internal/mods/guard/schema"
	"github.com/yz271544/settlement-monitoring/v10/pkg/util"
)

type Plotting struct {
	PlottingBiz *biz.Plotting
}

func (a *Plotting) Query(c *gin.Context) {
	var req schema.PlottingListReq
	if err := util.ParseQuery(c, &req); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.PlottingBiz.Query(c.Request.Context(), &req)
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

func (a *Plotting) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		util.ResError(c, errors.New("id不能为空"))
		return
	}

	result, err := a.PlottingBiz.Get(c.Request.Context(), id)
	if err != nil {
		util.ResError(c, err)
		return
	}

	util.ResSuccess(c, result)
}

func (a *Plotting) Create(c *gin.Context) {
	var req schema.PlottingCreateReq
	if err := util.ParseJSON(c, &req); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.PlottingBiz.Create(c.Request.Context(), &req)
	if err != nil {
		util.ResError(c, err)
		return
	}

	util.ResSuccess(c, result)
}

func (a *Plotting) Update(c *gin.Context) {
	var req schema.PlottingUpdateReq
	if err := util.ParseJSON(c, &req); err != nil {
		util.ResError(c, err)
		return
	}

	if err := a.PlottingBiz.Update(c.Request.Context(), &req); err != nil {
		util.ResError(c, err)
		return
	}

	util.ResOK(c)
}

func (a *Plotting) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		util.ResError(c, errors.New("id不能为空"))
		return
	}

	ids := []string{id}
	if err := a.PlottingBiz.Delete(c.Request.Context(), ids); err != nil {
		util.ResError(c, err)
		return
	}

	util.ResOK(c)
}

func (a *Plotting) Sort(c *gin.Context) {
	var req schema.SortPlottingReq
	if err := util.ParseJSON(c, &req); err != nil {
		util.ResError(c, err)
		return
	}

	if err := a.PlottingBiz.Sort(c.Request.Context(), &req); err != nil {
		util.ResError(c, err)
		return
	}

	util.ResOK(c)
}

func (a *Plotting) Statistics(c *gin.Context) {
	var req schema.PlottingStatisticsReq
	if err := util.ParseJSON(c, &req); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.PlottingBiz.Statistics(c.Request.Context(), &req)
	if err != nil {
		util.ResError(c, err)
		return
	}

	util.ResSuccess(c, result)
}

func (a *Plotting) Calculate(c *gin.Context) {
	var req schema.CalculateShapeReq
	if err := util.ParseJSON(c, &req); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.PlottingBiz.Calculate(c.Request.Context(), &req)
	if err != nil {
		util.ResError(c, err)
		return
	}

	util.ResSuccess(c, result)
}

func (a *Plotting) Import(c *gin.Context) {
	util.ResOK(c)
}

func (a *Plotting) SmokeAnalyse(c *gin.Context) {
	// TODO: 烟雾分析功能
	util.ResOK(c)
}
