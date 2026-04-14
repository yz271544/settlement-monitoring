package biz

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/yz271544/settlement-monitoring/v10/internal/mods/guard/dal"
	"github.com/yz271544/settlement-monitoring/v10/internal/mods/guard/schema"
	"gorm.io/gorm"
)

type Scene struct {
	SceneDal *dal.Scene
}

func (a *Scene) Query(ctx context.Context, req *schema.SceneListReq) (*schema.PageResult[*schema.Scene], error) {
	opts := []func(*gorm.DB) *gorm.DB{
		dal.WithSceneName(req.SceneName),
		dal.WithSceneType(req.SceneType),
		dal.WithStatus(req.Status),
		dal.WithProvinceCode(req.ProvinceCode),
		dal.WithCityCode(req.CityCode),
		dal.WithCountyCode(req.CountyCode),
	}

	opts = append(opts, func(db *gorm.DB) *gorm.DB {
		offset := (req.Page - 1) * req.PageSize
		return db.Offset(offset).Limit(req.PageSize)
	})

	items, err := a.SceneDal.Query(ctx, opts...)
	if err != nil {
		return nil, err
	}

	// 获取总数
	countOpts := []func(*gorm.DB) *gorm.DB{
		dal.WithSceneName(req.SceneName),
		dal.WithSceneType(req.SceneType),
		dal.WithStatus(req.Status),
		dal.WithProvinceCode(req.ProvinceCode),
		dal.WithCityCode(req.CityCode),
		dal.WithCountyCode(req.CountyCode),
	}
	total, err := a.SceneDal.Count(ctx, countOpts...)
	if err != nil {
		return nil, err
	}

	// 转换为指针数组
	result := make([]*schema.Scene, len(items))
	for i := range items {
		result[i] = &items[i]
	}

	return &schema.PageResult[*schema.Scene]{
		Total: total,
		Items: result,
	}, nil
}

func (a *Scene) Get(ctx context.Context, id string) (*schema.Scene, error) {
	if id == "" {
		return nil, errors.New("id不能为空")
	}
	return a.SceneDal.Get(ctx, id)
}

func (a *Scene) Create(ctx context.Context, req *schema.SceneCreateReq) (*schema.Scene, error) {
	// 生成 ID
	id := uuid.New().String()
	if len(id) > 32 {
		id = id[:32]
	}

	// 构建 Scene 对象
	scene := &schema.Scene{
		ID:                    id,
		SceneNo:               req.SceneNo,
		SceneName:             req.SceneName,
		OrgID:                 req.OrgID,
		OrgName:               req.OrgName,
		ContactName:           req.ContactName,
		ContactPhone:          req.ContactPhone,
		SceneDescription:      req.SceneDescription,
		SceneCover:            req.SceneCover,
		SceneType:             req.SceneType,
		ProvinceCode:          req.ProvinceCode,
		ProvinceName:          req.ProvinceName,
		CityCode:              req.CityCode,
		CityName:              req.CityName,
		CountyCode:            req.CountyCode,
		CountyName:            req.CountyName,
		TownshipCode:          req.TownshipCode,
		TownshipName:          req.TownshipName,
		RegionShape:           req.RegionShape,
		KMLName:               req.KMLName,
		ServiceCoordinate:     req.ServiceCoordinate,
		Status:                req.Status,
		CenterPoint:           req.CenterPoint,
		Envelope:              req.Envelope,
		Type:                  req.Type,
		NeighborhoodSituation: req.NeighborhoodSituation,
		DelFlag:               schema.DelFlagFalse,
		Remark:                req.Remark,
	}

	if err := a.SceneDal.Create(ctx, scene); err != nil {
		return nil, err
	}

	return scene, nil
}

func (a *Scene) Update(ctx context.Context, req *schema.SceneUpdateReq) error {
	scene, err := a.SceneDal.Get(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("场景不存在")
		}
		return err
	}

	// 更新字段
	scene.SceneNo = req.SceneNo
	scene.SceneName = req.SceneName
	scene.OrgID = req.OrgID
	scene.OrgName = req.OrgName
	scene.ContactName = req.ContactName
	scene.ContactPhone = req.ContactPhone
	scene.SceneDescription = req.SceneDescription
	scene.SceneCover = req.SceneCover
	scene.SceneType = req.SceneType
	scene.ProvinceCode = req.ProvinceCode
	scene.ProvinceName = req.ProvinceName
	scene.CityCode = req.CityCode
	scene.CityName = req.CityName
	scene.CountyCode = req.CountyCode
	scene.CountyName = req.CountyName
	scene.TownshipCode = req.TownshipCode
	scene.TownshipName = req.TownshipName
	scene.RegionShape = req.RegionShape
	scene.KMLName = req.KMLName
	scene.ServiceCoordinate = req.ServiceCoordinate
	scene.Status = req.Status
	scene.CenterPoint = req.CenterPoint
	scene.Envelope = req.Envelope
	scene.Type = req.Type
	scene.NeighborhoodSituation = req.NeighborhoodSituation
	scene.Remark = req.Remark

	return a.SceneDal.Update(ctx, scene)
}

func (a *Scene) Delete(ctx context.Context, ids []string) error {
	if len(ids) == 0 {
		return errors.New("ids不能为空")
	}
	return a.SceneDal.Delete(ctx, ids)
}

func (a *Scene) Merge(ctx context.Context, req *schema.SceneMergeReq) (*schema.SceneMergeResult, error) {
	if len(req.SourceIDs) == 0 {
		return nil, errors.New("源场景ID列表不能为空")
	}

	target, err := a.SceneDal.Get(ctx, req.TargetID)
	if err != nil {
		return nil, fmt.Errorf("目标场景不存在: %w", err)
	}

	var failedIDs []string
	for _, sourceID := range req.SourceIDs {
		if sourceID == req.TargetID {
			failedIDs = append(failedIDs, sourceID)
			continue
		}

		source, err := a.SceneDal.Get(ctx, sourceID)
		if err != nil {
			failedIDs = append(failedIDs, sourceID)
			continue
		}

		// 合并区域形状
		if source.RegionShape != "" && target.RegionShape != "" {
			target.RegionShape = mergeGeoJSON(target.RegionShape, source.RegionShape)
		}
	}

	// 更新目标场景
	if err := a.SceneDal.Update(ctx, target); err != nil {
		return nil, err
	}

	// 删除源场景
	if len(failedIDs) == 0 {
		_ = a.SceneDal.Delete(ctx, req.SourceIDs)
	}

	return &schema.SceneMergeResult{
		SuccessCount: len(req.SourceIDs) - len(failedIDs),
		FailedIDs:    failedIDs,
	}, nil
}

func (a *Scene) GetCenterPoint(ctx context.Context, id string) (string, error) {
	scene, err := a.SceneDal.Get(ctx, id)
	if err != nil {
		return "", err
	}
	return scene.CenterPoint, nil
}

// mergeGeoJSON 简化的 GeoJSON 合并
func mergeGeoJSON(target, source string) string {
	// TODO: 实现真正的 GeoJSON 合并逻辑
	// 这里需要使用 PostGIS 的 ST_Union 或类似的GIS库
	// 暂时返回目标
	return target
}
