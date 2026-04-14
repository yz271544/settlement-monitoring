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

type Plotting struct {
	PlottingDal *dal.Plotting
}

func (a *Plotting) Query(ctx context.Context, req *schema.PlottingListReq) (*schema.PageResult[*schema.Plotting], error) {
	opts := []func(*gorm.DB) *gorm.DB{
		dal.WithSceneID(req.SceneID),
		dal.WithPlottingType(req.PlottingType),
		dal.WithCoverageCode(req.CoverageCode),
		dal.WithPlottingName(req.Name),
	}

	opts = append(opts, func(db *gorm.DB) *gorm.DB {
		offset := (req.Page - 1) * req.PageSize
		return db.Offset(offset).Limit(req.PageSize)
	})

	items, err := a.PlottingDal.Query(ctx, opts...)
	if err != nil {
		return nil, err
	}

	countOpts := []func(*gorm.DB) *gorm.DB{
		dal.WithSceneID(req.SceneID),
		dal.WithPlottingType(req.PlottingType),
		dal.WithCoverageCode(req.CoverageCode),
		dal.WithPlottingName(req.Name),
	}
	total, err := a.PlottingDal.Count(ctx, countOpts...)
	if err != nil {
		return nil, err
	}

	result := make([]*schema.Plotting, len(items))
	for i := range items {
		result[i] = &items[i]
	}

	return &schema.PageResult[*schema.Plotting]{
		Total: total,
		Items: result,
	}, nil
}

func (a *Plotting) Get(ctx context.Context, id string) (*schema.Plotting, error) {
	if id == "" {
		return nil, errors.New("id不能为空")
	}
	return a.PlottingDal.Get(ctx, id)
}

func (a *Plotting) Create(ctx context.Context, req *schema.PlottingCreateReq) (*schema.Plotting, error) {
	id := uuid.New().String()
	if len(id) > 32 {
		id = id[:32]
	}

	plotting := &schema.Plotting{
		ID:                   id,
		SceneID:              req.SceneID,
		PlottingType:         req.PlottingType,
		CoverageID:           req.CoverageID,
		CoverageCode:         req.CoverageCode,
		BasicsPropertiesJson: req.BasicsPropertiesJson,
		ExtendPropertiesJson: req.ExtendPropertiesJson,
		Name:                 req.Name,
		LicensePlateNumber:   req.LicensePlateNumber,
		LongitudeLatitude:    req.LongitudeLatitude,
		Shape:                req.Shape,
		StyleFlag:            req.StyleFlag,
		LayerFlag:            req.LayerFlag,
		StyleInfoJson:        req.StyleInfoJson,
		RealityImagesOne:     req.RealityImagesOne,
		RealityImagesTwo:     req.RealityImagesTwo,
		RangeImages:          req.RangeImages,
		Images:               req.Images,
		Sort:                 req.Sort,
		WatchPerson:          req.WatchPerson,
		Equipment:            req.Equipment,
		DelFlag:              schema.DelFlagFalse,
	}

	if err := a.PlottingDal.Create(ctx, plotting); err != nil {
		return nil, err
	}

	return plotting, nil
}

func (a *Plotting) Update(ctx context.Context, req *schema.PlottingUpdateReq) error {
	plotting, err := a.PlottingDal.Get(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("标绘不存在")
		}
		return err
	}

	plotting.SceneID = req.SceneID
	plotting.PlottingType = req.PlottingType
	plotting.CoverageID = req.CoverageID
	plotting.CoverageCode = req.CoverageCode
	plotting.BasicsPropertiesJson = req.BasicsPropertiesJson
	plotting.ExtendPropertiesJson = req.ExtendPropertiesJson
	plotting.Name = req.Name
	plotting.LicensePlateNumber = req.LicensePlateNumber
	plotting.LongitudeLatitude = req.LongitudeLatitude
	plotting.Shape = req.Shape
	plotting.StyleFlag = req.StyleFlag
	plotting.LayerFlag = req.LayerFlag
	plotting.StyleInfoJson = req.StyleInfoJson
	plotting.RealityImagesOne = req.RealityImagesOne
	plotting.RealityImagesTwo = req.RealityImagesTwo
	plotting.RangeImages = req.RangeImages
	plotting.Images = req.Images
	plotting.Sort = req.Sort
	plotting.WatchPerson = req.WatchPerson
	plotting.Equipment = req.Equipment

	return a.PlottingDal.Update(ctx, plotting)
}

func (a *Plotting) Delete(ctx context.Context, ids []string) error {
	if len(ids) == 0 {
		return errors.New("ids不能为空")
	}
	return a.PlottingDal.Delete(ctx, ids)
}

func (a *Plotting) Sort(ctx context.Context, req *schema.SortPlottingReq) error {
	sortValues := make([]int, len(req.IDs))
	for i := range sortValues {
		sortValues[i] = i
	}
	return a.PlottingDal.BatchUpdateSort(ctx, req.IDs, sortValues)
}

func (a *Plotting) Statistics(ctx context.Context, req *schema.PlottingStatisticsReq) (*schema.PlottingStatisticsResult, error) {
	plottings, err := a.PlottingDal.Query(ctx, dal.WithSceneID(req.SceneID))
	if err != nil {
		return nil, err
	}

	typeCountMap := make(map[int]int64)
	for _, p := range plottings {
		typeCountMap[p.PlottingType]++
	}

	var items []schema.PlottingStatisticsItem
	typeNames := map[int]string{
		0: "点",
		1: "线",
		2: "面",
	}

	for pType, count := range typeCountMap {
		items = append(items, schema.PlottingStatisticsItem{
			PlottingType: pType,
			Count:        count,
			TypeName:     typeNames[pType],
		})
	}

	var totalCount int64
	for _, count := range typeCountMap {
		totalCount += count
	}

	result := &schema.PlottingStatisticsResult{
		TotalCount:       totalCount,
		SceneID:          req.SceneID,
		ItemCount:        int64(len(items)),
		StatisticsByType: items,
	}

	return result, nil
}

func (a *Plotting) Calculate(ctx context.Context, req *schema.CalculateShapeReq) (*schema.CalculateShapeResult, error) {
	result := &schema.CalculateShapeResult{
		Area:      calculateShapeArea(req.Shape, req.PlottingType),
		Perimeter: calculateShapePerimeter(req.Shape, req.PlottingType),
		CenterX:   0,
		CenterY:   0,
	}
	return result, nil
}

func calculateShapeArea(shape string, shapeType int) float64 {
	return 0
}

func calculateShapePerimeter(shape string, shapeType int) float64 {
	return 0
}
