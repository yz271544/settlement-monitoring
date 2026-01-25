package dal

import (
	"context"
	"github.com/yz271544/settlement-monitoring/v10/internal/mods/guard/schema"
	"gorm.io/gorm"
)

type Staff struct {
	DB *gorm.DB
}

func (a *Staff) Query(ctx context.Context, opts ...func(*gorm.DB) *gorm.DB) ([]schema.Staff, error) {
	var items []schema.Staff
	db := a.DB.WithContext(ctx)
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.Where("deleted = '0'").Find(&items).Error
	return items, err
}

func (a *Staff) Count(ctx context.Context, opts ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	db := a.DB.WithContext(ctx)
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.Model(&schema.Staff{}).Where("deleted = '0'").Count(&count).Error
	return count, err
}

func (a *Staff) Get(ctx context.Context, id string) (*schema.Staff, error) {
	var item schema.Staff
	err := a.DB.WithContext(ctx).Where("id = ? AND deleted = '0'", id).First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (a *Staff) Create(ctx context.Context, item *schema.Staff) error {
	return a.DB.WithContext(ctx).Create(item).Error
}

func (a *Staff) Update(ctx context.Context, item *schema.Staff) error {
	return a.DB.WithContext(ctx).Save(item).Error
}

func (a *Staff) Delete(ctx context.Context, ids []string) error {
	return a.DB.WithContext(ctx).Where("id IN ?", ids).Updates(map[string]interface{}{
		"deleted": "2",
	}).Error
}

// QueryOptions 构建查询条件
func WithName(name string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if name != "" {
			db = db.Where("name LIKE ?", "%"+name+"%")
		}
		return db
	}
}

func WithType(typ string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if typ != "" {
			db = db.Where("type = ?", typ)
		}
		return db
	}
}

func WithSex(sex string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if sex != "" {
			db = db.Where("sex = ?", sex)
		}
		return db
	}
}

func WithTel(tel string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if tel != "" {
			db = db.Where("tel LIKE ?", "%"+tel+"%")
		}
		return db
	}
}

func WithCompany(company string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if company != "" {
			db = db.Where("company LIKE ?", "%"+company+"%")
		}
		return db
	}
}

func WithDept(dept string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if dept != "" {
			db = db.Where("dept LIKE ?", "%"+dept+"%")
		}
		return db
	}
}

func WithSortAndOrder(sort, order string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if sort != "" {
			orderBy := sort
			if order == "desc" {
				orderBy += " DESC"
			}
			db = db.Order(orderBy)
		} else {
			db = db.Order("create_time DESC")
		}
		return db
	}
}
