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

type Staff struct {
	StaffDal *dal.Staff
}

func (a *Staff) Query(ctx context.Context, req *schema.StaffListReq) (*schema.PageResult[*schema.Staff], error) {
	opts := []func(*gorm.DB) *gorm.DB{
		dal.WithName(req.Name),
		dal.WithType(req.Type),
		dal.WithSex(req.Sex),
		dal.WithTel(req.Tel),
		dal.WithCompany(req.Company),
		dal.WithDept(req.Dept),
		dal.WithSortAndOrder(req.Sort, req.Order),
	}

	opts = append(opts, func(db *gorm.DB) *gorm.DB {
		offset := (req.Page - 1) * req.PageSize
		return db.Offset(offset).Limit(req.PageSize)
	})

	items, err := a.StaffDal.Query(ctx, opts...)
	if err != nil {
		return nil, err
	}

	// 获取总数
	countOpts := []func(*gorm.DB) *gorm.DB{
		dal.WithName(req.Name),
		dal.WithType(req.Type),
		dal.WithSex(req.Sex),
		dal.WithTel(req.Tel),
		dal.WithCompany(req.Company),
		dal.WithDept(req.Dept),
	}
	total, err := a.StaffDal.Count(ctx, countOpts...)
	if err != nil {
		return nil, err
	}

	// 转换为指针数组
	result := make([]*schema.Staff, len(items))
	for i := range items {
		result[i] = &items[i]
	}

	return &schema.PageResult[*schema.Staff]{
		Total: total,
		Items: result,
	}, nil
}

func (a *Staff) Get(ctx context.Context, id string) (*schema.Staff, error) {
	if id == "" {
		return nil, errors.New("id不能为空")
	}
	return a.StaffDal.Get(ctx, id)
}

func (a *Staff) Create(ctx context.Context, req *schema.StaffCreateReq) (*schema.Staff, error) {
	// 生成 ID
	id := uuid.New().String()
	if len(id) > 32 {
		id = id[:32]
	}

	// 构建 Staff 对象
	staff := &schema.Staff{
		ID:                  id,
		Name:                req.Name,
		Type:                req.Type,
		Sex:                 req.Sex,
		EducationBackground: req.EducationBackground,
		Post:                req.Post,
		BirthTime:           req.BirthTime,
		Dept:                req.Dept,
		Tel:                 req.Tel,
		Picture:             req.Picture,
		Company:             req.Company,
		CompanyName:         req.CompanyName,
		Duty:                req.Duty,
		Remark:              req.Remark,
		PersonnelID:         req.PersonnelID,
		PersonnelNum:        req.PersonnelNum,
		PoliticsStatus:      req.PoliticsStatus,
		Ethnicity:           req.Ethnicity,
		Nationality:         req.Nationality,
		Birthplace:          req.Birthplace,
		Address:             req.Address,
		Belong:              req.Belong,
		MergeStatus:         "1",
		Deleted:             "0",
	}

	if err := a.StaffDal.Create(ctx, staff); err != nil {
		return nil, err
	}

	return staff, nil
}

func (a *Staff) Update(ctx context.Context, req *schema.StaffUpdateReq) error {
	staff, err := a.StaffDal.Get(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("工作人员不存在")
		}
		return err
	}

	// 更新字段
	staff.Name = req.Name
	staff.Type = req.Type
	staff.Sex = req.Sex
	staff.EducationBackground = req.EducationBackground
	staff.Post = req.Post
	staff.BirthTime = req.BirthTime
	staff.Dept = req.Dept
	staff.Tel = req.Tel
	staff.Picture = req.Picture
	staff.Company = req.Company
	staff.CompanyName = req.CompanyName
	staff.Duty = req.Duty
	staff.Remark = req.Remark
	staff.PersonnelID = req.PersonnelID
	staff.PersonnelNum = req.PersonnelNum
	staff.PoliticsStatus = req.PoliticsStatus
	staff.Ethnicity = req.Ethnicity
	staff.Nationality = req.Nationality
	staff.Birthplace = req.Birthplace
	staff.Address = req.Address
	staff.Belong = req.Belong

	return a.StaffDal.Update(ctx, staff)
}

func (a *Staff) Delete(ctx context.Context, ids []string) error {
	if len(ids) == 0 {
		return errors.New("ids不能为空")
	}
	return a.StaffDal.Delete(ctx, ids)
}
