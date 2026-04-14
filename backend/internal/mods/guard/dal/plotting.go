package dal

import (
	"context"

	"github.com/yz271544/settlement-monitoring/v10/internal/mods/guard/schema"
	"gorm.io/gorm"
)

type Plotting struct {
	DB *gorm.DB
}

func (a *Plotting) Query(ctx context.Context, opts ...func(*gorm.DB) *gorm.DB) ([]schema.Plotting, error) {
	var items []schema.Plotting
	db := a.DB.WithContext(ctx)
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.Where("del_flag = ?", schema.DelFlagFalse).Find(&items).Error
	return items, err
}

func (a *Plotting) Count(ctx context.Context, opts ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	db := a.DB.WithContext(ctx)
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.Model(&schema.Plotting{}).Where("del_flag = ?", schema.DelFlagFalse).Count(&count).Error
	return count, err
}

func (a *Plotting) Get(ctx context.Context, id string) (*schema.Plotting, error) {
	var item schema.Plotting
	err := a.DB.WithContext(ctx).Where("id = ? AND del_flag = ?", id, schema.DelFlagFalse).First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (a *Plotting) Create(ctx context.Context, item *schema.Plotting) error {
	return a.DB.WithContext(ctx).Create(item).Error
}

func (a *Plotting) Update(ctx context.Context, item *schema.Plotting) error {
	return a.DB.WithContext(ctx).Save(item).Error
}

func (a *Plotting) Delete(ctx context.Context, ids []string) error {
	return a.DB.WithContext(ctx).Where("id IN ?", ids).Update("del_flag", schema.DelFlagTrue).Error
}

func (a *Plotting) BatchUpdateSort(ctx context.Context, ids []string, sortValues []int) error {
	return a.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for i, id := range ids {
			if err := tx.Model(&schema.Plotting{}).Where("id = ?", id).Update("sort", sortValues[i]).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// QueryOptions 构建查询条件
func WithSceneID(sceneID string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if sceneID != "" {
			db = db.Where("scene_id = ?", sceneID)
		}
		return db
	}
}

func WithPlottingType(plottingType int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if plottingType >= 0 && plottingType <= 2 {
			db = db.Where("plotting_type = ?", plottingType)
		}
		return db
	}
}

func WithCoverageCode(coverageCode string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if coverageCode != "" {
			db = db.Where("coverage_code = ?", coverageCode)
		}
		return db
	}
}

func WithPlottingName(name string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if name != "" {
			db = db.Where("name LIKE ?", "%"+name+"%")
		}
		return db
	}
}
