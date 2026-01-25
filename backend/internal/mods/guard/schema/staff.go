package schema

import (
	"time"
)

// Staff 工作人员表
type Staff struct {
	ID                  string    `gorm:"column:id;size:32;primaryKey" json:"id"`
	Name                string    `gorm:"column:name;size:20;not null" json:"name" validate:"required,max=20"`
	Type                string    `gorm:"column:type;size:10;not null" json:"type" validate:"required,max=10"`
	Sex                 string    `gorm:"column:sex;size:1;not null" json:"sex" validate:"required,oneof=0 1"`
	EducationBackground string    `gorm:"column:education_background;size:10" json:"education_background"`
	Post                string    `gorm:"column:post;size:20" json:"post"`
	BirthTime           time.Time `gorm:"column:birth_time" json:"birth_time"`
	Dept                string    `gorm:"column:dept;size:255" json:"dept"`
	Tel                 string    `gorm:"column:tel;size:20;not null" json:"tel" validate:"required,max=20"`
	Picture             string    `gorm:"column:picture;size:100" json:"picture"`
	Company             string    `gorm:"column:company;size:100" json:"company"`
	CompanyName         string    `gorm:"column:company_name;size:255" json:"company_name"`
	Duty                string    `gorm:"column:duty;size:255" json:"duty"`
	Remark              string    `gorm:"column:remark;size:255" json:"remark"`
	PersonnelID         string    `gorm:"column:personnel_id;size:100;not null;index" json:"personnel_id" validate:"required,max=100"`
	PersonnelNum        string    `gorm:"column:personnel_num;size:100" json:"personnel_num"`
	PoliticsStatus      string    `gorm:"column:politics_status;size:10" json:"politics_status"`
	Ethnicity           string    `gorm:"column:ethnicity;size:255" json:"ethnicity"`
	Nationality         string    `gorm:"column:nationality;size:255" json:"nationality"`
	Birthplace          string    `gorm:"column:birthplace;size:255" json:"birthplace"`
	Address             string    `gorm:"column:address;size:255" json:"address"`
	Belong              string    `gorm:"column:belong;size:100" json:"belong"`
	ProcInsID           string    `gorm:"column:proc_ins_id;size:32" json:"proc_ins_id"`
	MergeStatus         string    `gorm:"column:merge_status;size:1;default:'1'" json:"merge_status"`
	Deleted             string    `gorm:"column:deleted;size:1;default:'0';index" json:"deleted"`
	CreateBy            string    `gorm:"column:create_by;size:64;default:''" json:"create_by"`
	CreateTime          time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
	UpdateBy            string    `gorm:"column:update_by;size:64;default:''" json:"update_by"`
	UpdateTime          time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
}

func (Staff) TableName() string {
	return "t_staff"
}

// StaffListReq 工作人员列表请求
type StaffListReq struct {
	PaginationReq
	QueryOptions
	Name        string `json:"name" form:"name"`
	Type        string `json:"type" form:"type"`
	Sex         string `json:"sex" form:"sex"`
	Tel         string `json:"tel" form:"tel"`
	PersonnelID string `json:"personnel_id" form:"personnel_id"`
	Company     string `json:"company" form:"company"`
	Dept        string `json:"dept" form:"dept"`
}

// StaffCreateReq 工作人员创建请求
type StaffCreateReq struct {
	MergeStatus         string    `json:"merge_status" validate:"omitempty,max=1"`
	Deleted             string    `json:"deleted" validate:"omitempty,max=1"`
	Name                string    `json:"name" validate:"required,max=20"`
	Type                string    `json:"type" validate:"required,max=10"`
	Sex                 string    `json:"sex" validate:"required,oneof=0 1"`
	EducationBackground string    `json:"education_background" validate:"max=10"`
	Post                string    `json:"post" validate:"max=20"`
	BirthTime           time.Time `json:"birth_time"`
	Dept                string    `json:"dept"`
	Tel                 string    `json:"tel" validate:"required,max=20"`
	Picture             string    `json:"picture" validate:"max=100"`
	Company             string    `json:"company" validate:"max=100"`
	CompanyName         string    `json:"company_name"`
	Duty                string    `json:"duty"`
	Remark              string    `json:"remark"`
	PersonnelID         string    `json:"personnel_id" validate:"required,max=100"`
	PersonnelNum        string    `json:"personnel_num" validate:"max=100"`
	PoliticsStatus      string    `json:"politics_status" validate:"max=10"`
	Ethnicity           string    `json:"ethnicity"`
	Nationality         string    `json:"nationality"`
	Birthplace          string    `json:"birthplace"`
	Address             string    `json:"address"`
	Belong              string    `json:"belong"`
}

func (a *StaffCreateReq) SetDefaults() {
	if a.MergeStatus == "" {
		a.MergeStatus = "1"
	}
	if a.Deleted == "" {
		a.Deleted = "0"
	}
}

// StaffUpdateReq 工作人员更新请求
type StaffUpdateReq struct {
	ID                  string    `json:"id" validate:"required"`
	Name                string    `json:"name" validate:"required,max=20"`
	Type                string    `json:"type" validate:"required,max=10"`
	Sex                 string    `json:"sex" validate:"required,oneof=0 1"`
	EducationBackground string    `json:"education_background" validate:"max=10"`
	Post                string    `json:"post" validate:"max=20"`
	BirthTime           time.Time `json:"birth_time"`
	Dept                string    `json:"dept"`
	Tel                 string    `json:"tel" validate:"required,max=20"`
	Picture             string    `json:"picture" validate:"max=100"`
	Company             string    `json:"company" validate:"max=100"`
	CompanyName         string    `json:"company_name"`
	Duty                string    `json:"duty"`
	Remark              string    `json:"remark"`
	PersonnelID         string    `json:"personnel_id" validate:"required,max=100"`
	PersonnelNum        string    `json:"personnel_num" validate:"max=100"`
	PoliticsStatus      string    `json:"politics_status" validate:"max=10"`
	Ethnicity           string    `json:"ethnicity"`
	Nationality         string    `json:"nationality"`
	Birthplace          string    `json:"birthplace"`
	Address             string    `json:"address"`
	Belong              string    `json:"belong"`
}

// StaffExportReq 工作人员导出请求
type StaffExportReq struct {
	IDs []string `json:"ids" validate:"required"`
}

// ImportResult 导入结果
type ImportResult struct {
	SuccessCount int      `json:"success_count"`
	FailCount    int      `json:"fail_count"`
	FailRecords  []string `json:"fail_records"`
}

// SetDefaults 设置默认值
func (v *ImportResult) SetDefaults() {
	if v.FailRecords == nil {
		v.FailRecords = []string{}
	}
}
