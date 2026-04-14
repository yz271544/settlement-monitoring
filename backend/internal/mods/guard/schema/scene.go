package schema

import "time"

const (
	// SceneType 场景类型
	SceneTypeAdministrative = "01" // 行政边界
	SceneTypeKML            = "02" // KML文件
	SceneTypeResidence      = "04" // 驻地

	// DelFlag 删除标记
	DelFlagFalse = false
	DelFlagTrue  = true
)

// Scene 场景表
type Scene struct {
	ID                    string    `gorm:"column:id;type:varchar(32);primaryKey" json:"id"`
	SceneNo               string    `gorm:"column:scene_no;type:varchar(50)" json:"scene_no"`
	SceneName             string    `gorm:"column:scene_name;type:varchar(200)" json:"scene_name"`
	OrgID                 string    `gorm:"column:org_id;type:varchar(32)" json:"org_id"`
	OrgName               string    `gorm:"column:org_name;type:varchar(255)" json:"org_name"`
	ContactName           string    `gorm:"column:contact_name;type:varchar(100)" json:"contact_name"`
	ContactPhone          string    `gorm:"column:contact_phone;type:varchar(50)" json:"contact_phone"`
	SceneDescription      string    `gorm:"column:scene_description;type:text" json:"scene_description"`
	SceneCover            string    `gorm:"column:scene_cover;type:varchar(500)" json:"scene_cover"`
	SceneType             string    `gorm:"column:scene_type;type:varchar(10)" json:"scene_type"`
	ProvinceCode          string    `gorm:"column:province_code;type:varchar(20)" json:"province_code"`
	ProvinceName          string    `gorm:"column:province_name;type:varchar(100)" json:"province_name"`
	CityCode              string    `gorm:"column:city_code;type:varchar(20)" json:"city_code"`
	CityName              string    `gorm:"column:city_name;type:varchar(100)" json:"city_name"`
	CountyCode            string    `gorm:"column:county_code;type:varchar(20)" json:"county_code"`
	CountyName            string    `gorm:"column:county_name;type:varchar(100)" json:"county_name"`
	TownshipCode          string    `gorm:"column:township_code;type:varchar(20)" json:"township_code"`
	TownshipName          string    `gorm:"column:township_name;type:varchar(100)" json:"township_name"`
	RegionShape           string    `gorm:"column:region_shape;type:text" json:"region_shape"`
	KMLName               string    `gorm:"column:kml_name;type:varchar(255)" json:"kml_name"`
	ServiceCoordinate     string    `gorm:"column:service_coordinate;type:varchar(255)" json:"service_coordinate"`
	Status                string    `gorm:"column:status;type:varchar(20)" json:"status"`
	CenterPoint           string    `gorm:"column:center_point;type:varchar(50)" json:"center_point"`
	Envelope              string    `gorm:"column:envelope;type:text" json:"envelope"`
	Type                  string    `gorm:"column:type;type:varchar(10)" json:"type"`
	NeighborhoodSituation string    `gorm:"column:neighborhood_situation;type:text" json:"neighborhood_situation"`
	DelFlag               bool      `gorm:"column:del_flag;default:false;index" json:"del_flag"`
	CreateBy              string    `gorm:"column:create_by;type:varchar(64);default:''" json:"create_by"`
	CreateTime            time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
	UpdateBy              string    `gorm:"column:update_by;type:varchar(64);default:''" json:"update_by"`
	UpdateTime            time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
	Remark                string    `gorm:"column:remark;type:varchar(255)" json:"remark"`
}

func (Scene) TableName() string {
	return "t_scene"
}

// SceneListReq 场景列表请求
type SceneListReq struct {
	PaginationReq
	QueryOptions
	SceneName    string `json:"scene_name,omitempty" form:"scene_name"`
	SceneType    string `json:"scene_type,omitempty" form:"scene_type"`
	ProvinceCode string `json:"province_code,omitempty" form:"province_code"`
	CityCode     string `json:"city_code,omitempty" form:"city_code"`
	CountyCode   string `json:"county_code,omitempty" form:"county_code"`
	Status       string `json:"status,omitempty" form:"status"`
}

// SceneCreateReq 场景创建请求
type SceneCreateReq struct {
	SceneNo               string `json:"scene_no,omitempty" validate:"max=50"`
	SceneName             string `json:"scene_name" validate:"required,max=200"`
	OrgID                 string `json:"org_id,omitempty" validate:"max=32"`
	OrgName               string `json:"org_name,omitempty" validate:"max=255"`
	ContactName           string `json:"contact_name,omitempty" validate:"max=100"`
	ContactPhone          string `json:"contact_phone,omitempty" validate:"max=50"`
	SceneDescription      string `json:"scene_description,omitempty" validate:"max=65535"`
	SceneCover            string `json:"scene_cover,omitempty" validate:"max=500"`
	SceneType             string `json:"scene_type" validate:"required,max=10"`
	ProvinceCode          string `json:"province_code,omitempty" validate:"max=20"`
	ProvinceName          string `json:"province_name,omitempty" validate:"max=100"`
	CityCode              string `json:"city_code,omitempty" validate:"max=20"`
	CityName              string `json:"city_name,omitempty" validate:"max=100"`
	CountyCode            string `json:"county_code,omitempty" validate:"max=20"`
	CountyName            string `json:"county_name,omitempty" validate:"max=100"`
	TownshipCode          string `json:"township_code,omitempty" validate:"max=20"`
	TownshipName          string `json:"township_name,omitempty" validate:"max=100"`
	RegionShape           string `json:"region_shape,omitempty" validate:"max=65535"`
	KMLName               string `json:"kml_name,omitempty" validate:"max=255"`
	ServiceCoordinate     string `json:"service_coordinate,omitempty" validate:"max=255"`
	Status                string `json:"status,omitempty" validate:"max=20"`
	CenterPoint           string `json:"center_point,omitempty" validate:"max=50"`
	Envelope              string `json:"envelope,omitempty" validate:"max=65535"`
	Type                  string `json:"type,omitempty" validate:"max=10"`
	NeighborhoodSituation string `json:"neighborhood_situation,omitempty" validate:"max=65535"`
	Remark                string `json:"remark,omitempty" validate:"max=255"`
}

// SceneUpdateReq 场景更新请求
type SceneUpdateReq struct {
	ID                    string `json:"id" validate:"required"`
	SceneNo               string `json:"scene_no,omitempty" validate:"max=50"`
	SceneName             string `json:"scene_name" validate:"required,max=200"`
	OrgID                 string `json:"org_id,omitempty" validate:"max=32"`
	OrgName               string `json:"org_name,omitempty" validate:"max=255"`
	ContactName           string `json:"contact_name,omitempty" validate:"max=100"`
	ContactPhone          string `json:"contact_phone,omitempty" validate:"max=50"`
	SceneDescription      string `json:"scene_description,omitempty" validate:"max=65535"`
	SceneCover            string `json:"scene_cover,omitempty" validate:"max=500"`
	SceneType             string `json:"scene_type" validate:"max=10"`
	ProvinceCode          string `json:"province_code,omitempty" validate:"max=20"`
	ProvinceName          string `json:"province_name,omitempty" validate:"max=100"`
	CityCode              string `json:"city_code,omitempty" validate:"max=20"`
	CityName              string `json:"city_name,omitempty" validate:"max=100"`
	CountyCode            string `json:"county_code,omitempty" validate:"max=20"`
	CountyName            string `json:"county_name,omitempty" validate:"max=100"`
	TownshipCode          string `json:"township_code,omitempty" validate:"max=20"`
	TownshipName          string `json:"township_name,omitempty" validate:"max=100"`
	RegionShape           string `json:"region_shape,omitempty" validate:"max=65535"`
	KMLName               string `json:"kml_name,omitempty" validate:"max=255"`
	ServiceCoordinate     string `json:"service_coordinate,omitempty" validate:"max=255"`
	Status                string `json:"status,omitempty" validate:"max=20"`
	CenterPoint           string `json:"center_point,omitempty" validate:"max=50"`
	Envelope              string `json:"envelope,omitempty" validate:"max=65535"`
	Type                  string `json:"type,omitempty" validate:"max=10"`
	NeighborhoodSituation string `json:"neighborhood_situation,omitempty" validate:"max=65535"`
	Remark                string `json:"remark,omitempty" validate:"max=255"`
}

// SceneMergeReq 场景合并请求
type SceneMergeReq struct {
	SourceIDs []string `json:"source_ids" validate:"required,min=1"`
	TargetID  string   `json:"target_id" validate:"required"`
}

// SceneMergeResult 场景合并结果
type SceneMergeResult struct {
	SuccessCount int      `json:"success_count"`
	FailedIDs    []string `json:"failed_ids"`
}
