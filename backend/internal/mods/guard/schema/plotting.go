package schema

import "time"

const (
	// PlottingType 标绘类型
	PlottingTypePoint   = 0 // 点
	PlottingTypeLine    = 1 // 线
	PlottingTypePolygon = 2 // 面
)

// Plotting 标绘表
type Plotting struct {
	ID                   string    `gorm:"column:id;type:varchar(32);primaryKey" json:"id"`
	SceneID              string    `gorm:"column:scene_id;type:varchar(32);index" json:"scene_id"`
	PlottingType         int       `gorm:"column:plotting_type;type:integer;index" json:"plotting_type"`
	CoverageID           string    `gorm:"column:coverage_id;type:varchar(32);index" json:"coverage_id"`
	CoverageCode         string    `gorm:"column:coverage_code;type:varchar(50)" json:"coverage_code"`
	BasicsPropertiesJson string    `gorm:"column:basics_properties_json;type:text" json:"basics_properties_json"`
	ExtendPropertiesJson string    `gorm:"column:extend_properties_json;type:text" json:"extend_properties_json"`
	Name                 string    `gorm:"column:name;type:varchar(255)" json:"name"`
	LicensePlateNumber   string    `gorm:"column:license_plate_number;type:varchar(100)" json:"license_plate_number"`
	LongitudeLatitude    string    `gorm:"column:longitude_latitude;type:varchar(100)" json:"longitude_latitude"`
	Shape                string    `gorm:"column:shape;type:text" json:"shape"`
	StyleFlag            int       `gorm:"column:style_flag;type:integer;default:0" json:"style_flag"`
	LayerFlag            int       `gorm:"column:layer_flag;type:integer;default:0" json:"layer_flag"`
	StyleInfoJson        string    `gorm:"column:style_info_json;type:text" json:"style_info_json"`
	RealityImagesOne     string    `gorm:"column:reality_images_one;type:varchar(500)" json:"reality_images_one"`
	RealityImagesTwo     string    `gorm:"column:reality_images_two;type:varchar(500)" json:"reality_images_two"`
	RangeImages          string    `gorm:"column:range_images;type:varchar(500)" json:"range_images"`
	Images               string    `gorm:"column:images;type:varchar(500)" json:"images"`
	Sort                 int       `gorm:"column:sort;type:integer;default:0" json:"sort"`
	WatchPerson          string    `gorm:"column:watch_person;type:varchar(500)" json:"watch_person"`
	Equipment            string    `gorm:"column:equipment;type:varchar(500)" json:"equipment"`
	DelFlag              bool      `gorm:"column:del_flag;default:false;index" json:"del_flag"`
	CreateBy             string    `gorm:"column:create_by;type:varchar(64);default:''" json:"create_by"`
	CreateTime           time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
	UpdateBy             string    `gorm:"column:update_by;type:varchar(64);default:''" json:"update_by"`
	UpdateTime           time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
}

func (Plotting) TableName() string {
	return "t_plotting"
}

// PlottingListReq 标绘列表请求
type PlottingListReq struct {
	PaginationReq
	QueryOptions
	SceneID      string `json:"scene_id,omitempty" form:"scene_id"`
	PlottingType int    `json:"plotting_type,omitempty" form:"plotting_type"`
	CoverageCode string `json:"coverage_code,omitempty" form:"coverage_code"`
	Name         string `json:"name,omitempty" form:"name"`
}

// PlottingCreateReq 标绘创建请求
type PlottingCreateReq struct {
	SceneID              string `json:"scene_id" validate:"required,max=32"`
	PlottingType         int    `json:"plotting_type" validate:"required,min=0,max=2"`
	CoverageID           string `json:"coverage_id,omitempty" validate:"max=32"`
	CoverageCode         string `json:"coverage_code,omitempty" validate:"max=50"`
	BasicsPropertiesJson string `json:"basics_properties_json,omitempty"`
	ExtendPropertiesJson string `json:"extend_properties_json,omitempty"`
	Name                 string `json:"name,omitempty" validate:"max=255"`
	LicensePlateNumber   string `json:"license_plate_number,omitempty" validate:"max=100"`
	LongitudeLatitude    string `json:"longitude_latitude,omitempty" validate:"max=100"`
	Shape                string `json:"shape,omitempty" validate:"max=65535"`
	StyleFlag            int    `json:"style_flag" default:"0"`
	LayerFlag            int    `json:"layer_flag" default:"0"`
	StyleInfoJson        string `json:"style_info_json,omitempty"`
	RealityImagesOne     string `json:"reality_images_one,omitempty" validate:"max=500"`
	RealityImagesTwo     string `json:"reality_images_two,omitempty" validate:"max=500"`
	RangeImages          string `json:"range_images,omitempty" validate:"max=500"`
	Images               string `json:"images,omitempty" validate:"max=500"`
	Sort                 int    `json:"sort" default:"0"`
	WatchPerson          string `json:"watch_person,omitempty" validate:"max=500"`
	Equipment            string `json:"equipment,omitempty" validate:"max=500"`
}

// PlottingUpdateReq 标绘更新请求
type PlottingUpdateReq struct {
	ID                   string `json:"id" validate:"required"`
	SceneID              string `json:"scene_id,omitempty" validate:"max=32"`
	PlottingType         int    `json:"plotting_type,omitempty" validate:"min=0,max=2"`
	CoverageID           string `json:"coverage_id,omitempty" validate:"max=32"`
	CoverageCode         string `json:"coverage_code,omitempty" validate:"max=50"`
	BasicsPropertiesJson string `json:"basics_properties_json,omitempty"`
	ExtendPropertiesJson string `json:"extend_properties_json,omitempty"`
	Name                 string `json:"name,omitempty" validate:"max=255"`
	LicensePlateNumber   string `json:"license_plate_number,omitempty" validate:"max=100"`
	LongitudeLatitude    string `json:"longitude_latitude,omitempty" validate:"max=100"`
	Shape                string `json:"shape,omitempty" validate:"max=65535"`
	StyleFlag            int    `json:"style_flag"`
	LayerFlag            int    `json:"layer_flag"`
	StyleInfoJson        string `json:"style_info_json,omitempty"`
	RealityImagesOne     string `json:"reality_images_one,omitempty" validate:"max=500"`
	RealityImagesTwo     string `json:"reality_images_two,omitempty" validate:"max=500"`
	RangeImages          string `json:"range_images,omitempty" validate:"max=500"`
	Images               string `json:"images,omitempty" validate:"max=500"`
	Sort                 int    `json:"sort"`
	WatchPerson          string `json:"watch_person,omitempty" validate:"max=500"`
	Equipment            string `json:"equipment,omitempty" validate:"max=500"`
}

// PlottingStatisticsReq 标绘统计请求
type PlottingStatisticsReq struct {
	SceneID string `json:"scene_id" validate:"required"`
}

// PlottingStatisticsItem 标绘统计项
type PlottingStatisticsItem struct {
	PlottingType int    `json:"plotting_type"`
	Count        int64  `json:"count"`
	TypeName     string `json:"type_name"`
}

// PlottingStatisticsResult 标绘统计结果
type PlottingStatisticsResult struct {
	TotalCount       int64                    `json:"total_count"`
	SceneID          string                   `json:"scene_id"`
	SceneName        string                   `json:"scene_name"`
	ItemCount        int64                    `json:"item_count"`
	StatisticsByType []PlottingStatisticsItem `json:"statistics_by_type"`
}

// CalculateShapeReq 计算图形请求
type CalculateShapeReq struct {
	PlottingType int    `json:"plotting_type" validate:"required,min=0,max=2"`
	Shape        string `json:"shape" validate:"required"`
}

// CalculateShapeResult 计算图形结果
type CalculateShapeResult struct {
	Area      float64 `json:"area"`      // 面积（平方米）
	Perimeter float64 `json:"perimeter"` // 周长（米）
	CenterX   float64 `json:"center_x"`  // 中心点X坐标
	CenterY   float64 `json:"center_y"`  // 中心点Y坐标
}

// SortPlottingReq 排序标绘请求
type SortPlottingReq struct {
	IDs []string `json:"ids" validate:"required,min=1"`
}
