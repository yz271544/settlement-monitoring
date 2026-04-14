package dal

import (
	"context"

	"github.com/yz271544/settlement-monitoring/v10/internal/mods/guard/schema"
	"gorm.io/gorm"
)

type Scene struct {
	DB *gorm.DB
}

func (a *Scene) Query(ctx context.Context, opts ...func(*gorm.DB) *gorm.DB) ([]schema.Scene, error) {
	var items []schema.Scene
	db := a.DB.WithContext(ctx)
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.Where("del_flag = ?", schema.DelFlagFalse).Find(&items).Error
	return items, err
}

func (a *Scene) Count(ctx context.Context, opts ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	db := a.DB.WithContext(ctx)
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.Model(&schema.Scene{}).Where("del_flag = ?", schema.DelFlagFalse).Count(&count).Error
	return count, err
}

func (a *Scene) Get(ctx context.Context, id string) (*schema.Scene, error) {
	var item schema.Scene
	err := a.DB.WithContext(ctx).Where("id = ? AND del_flag = ?", id, schema.DelFlagFalse).First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (a *Scene) Create(ctx context.Context, item *schema.Scene) error {
	return a.DB.WithContext(ctx).Create(item).Error
}

func (a *Scene) Update(ctx context.Context, item *schema.Scene) error {
	return a.DB.WithContext(ctx).Save(item).Error
}

func (a *Scene) Delete(ctx context.Context, ids []string) error {
	return a.DB.WithContext(ctx).Where("id IN ?", ids).Update("del_flag", schema.DelFlagTrue).Error
}

// QueryOptions 构建查询条件
func WithSceneName(name string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if name != "" {
			db = db.Where("scene_name LIKE ?", "%"+name+"%")
		}
		return db
	}
}

func WithSceneType(sceneType string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if sceneType != "" {
			db = db.Where("scene_type = ?", sceneType)
		}
		return db
	}
}

func WithStatus(status string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status != "" {
			db = db.Where("status = ?", status)
		}
		return db
	}
}

func WithProvinceCode(provinceCode string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if provinceCode != "" {
			db = db.Where("province_code = ?", provinceCode)
		}
		return db
	}
}

func WithCityCode(cityCode string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if cityCode != "" {
			db = db.Where("city_code = ?", cityCode)
		}
		return db
	}
}

func WithCountyCode(countyCode string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if countyCode != "" {
			db = db.Where("county_code = ?", countyCode)
		}
		return db
	}
}
