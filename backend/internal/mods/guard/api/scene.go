package api

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/yz271544/settlement-monitoring/v10/internal/mods/guard/biz"
	"github.com/yz271544/settlement-monitoring/v10/internal/mods/guard/schema"
	"github.com/yz271544/settlement-monitoring/v10/pkg/util"
)

type Scene struct {
	SceneBiz *biz.Scene
}

func (a *Scene) Query(c *gin.Context) {
	var req schema.SceneListReq
	if err := util.ParseQuery(c, &req); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.SceneBiz.Query(c.Request.Context(), &req)
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

func (a *Scene) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		util.ResError(c, errors.New("id不能为空"))
		return
	}

	result, err := a.SceneBiz.Get(c.Request.Context(), id)
	if err != nil {
		util.ResError(c, err)
		return
	}

	util.ResSuccess(c, result)
}

func (a *Scene) Create(c *gin.Context) {
	var req schema.SceneCreateReq
	if err := util.ParseJSON(c, &req); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.SceneBiz.Create(c.Request.Context(), &req)
	if err != nil {
		util.ResError(c, err)
		return
	}

	util.ResSuccess(c, result)
}

func (a *Scene) Update(c *gin.Context) {
	var req schema.SceneUpdateReq
	if err := util.ParseJSON(c, &req); err != nil {
		util.ResError(c, err)
		return
	}

	if err := a.SceneBiz.Update(c.Request.Context(), &req); err != nil {
		util.ResError(c, err)
		return
	}

	util.ResOK(c)
}

func (a *Scene) Delete(c *gin.Context) {
	ids := c.Param("id")
	if ids == "" {
		util.ResError(c, errors.New("id不能为空"))
		return
	}

	if err := a.SceneBiz.Delete(c.Request.Context(), []string{ids}); err != nil {
		util.ResError(c, err)
		return
	}

	util.ResOK(c)
}

func (a *Scene) Merge(c *gin.Context) {
	var req schema.SceneMergeReq
	if err := util.ParseJSON(c, &req); err != nil {
		util.ResError(c, err)
		return
	}

	result, err := a.SceneBiz.Merge(c.Request.Context(), &req)
	if err != nil {
		util.ResError(c, err)
		return
	}

	util.ResSuccess(c, result)
}

func (a *Scene) GetCenterPoint(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		util.ResError(c, errors.New("id不能为空"))
		return
	}

	result, err := a.SceneBiz.GetCenterPoint(c.Request.Context(), id)
	if err != nil {
		util.ResError(c, err)
		return
	}

	util.ResSuccess(c, map[string]interface{}{
		"center_point": result,
	})
}

func (a *Scene) Import(c *gin.Context) {
	util.ResOK(c)
}

func (a *Scene) Export(c *gin.Context) {
	util.ResOK(c)
}

func (a *Scene) Split(c *gin.Context) {
	util.ResOK(c)
}

func (a *Scene) Statistics(c *gin.Context) {
	util.ResOK(c)
}
