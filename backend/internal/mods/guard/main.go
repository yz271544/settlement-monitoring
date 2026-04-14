package guard

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yz271544/settlement-monitoring/v10/internal/mods/guard/api"
	"github.com/yz271544/settlement-monitoring/v10/internal/mods/guard/schema"
	"gorm.io/gorm"
)

type Guard struct {
	DB          *gorm.DB
	StaffAPI    *api.Staff
	SceneAPI    *api.Scene
	PlottingAPI *api.Plotting

	// TODO: 其他API逐步实现
	/*
		WatchkeeperAPI        *api.Watchkeeper
		TargetedIndividualsAPI *api.TargetedIndividuals
		DogTrainerAPI         *api.DogTrainer
		PoliceDogAPI          *api.PoliceDog
		CarAPI                *api.Car
		FirearmAPI            *api.Firearm
		InterphoneAPI         *api.Interphone
		DroneCounterEquipmentAPI    *api.DroneCounterEquipment
		SecurityScreeningEquipmentAPI *api.SecurityScreeningEquipment
		PoliceRecorderAPI     *api.PoliceRecorder
		VantagePointAPI       *api.VantagePoint
		KeyPartAPI            *api.KeyPart
		DangerousPartAPI      *api.DangerousPart
	*/
}

func (a *Guard) AutoMigrate(ctx context.Context) error {
	return a.DB.AutoMigrate(
		new(schema.Staff),
		new(schema.Scene),
		new(schema.Plotting),

		// TODO: 其他表逐步迁移
		/*
			new(schema.Watchkeeper),
			new(schema.TargetedIndividuals),
			new(schema.DogTrainer),
			new(schema.PoliceDog),
			new(schema.Car),
			new(schema.Firearm),
			new(schema.Interphone),
			new(schema.DroneCounterEquipment),
			new(schema.SecurityScreeningEquipment),
			new(schema.PoliceRecorder),
			new(schema.VantagePoint),
			new(schema.KeyPart),
			new(schema.DangerousPart),
			new(schema.Flow),
			new(schema.GisRegion),
			new(schema.Residence),
			new(schema.Camera),
			new(schema.Equipment),
			new(schema.Person),
			new(schema.ThreeDimensionModel),
			new(schema.VrInfo),
			new(schema.Legend),
			new(schema.CoverageTree),
			new(schema.SceneMap),
			new(schema.SimulationRoute),
		*/
	)
}

func (a *Guard) Init(ctx context.Context) error {
	// TODO: 可以在这里从配置文件读取是否自动迁移
	// if config.C.Storage.DB.AutoMigrate {
	//     if err := a.AutoMigrate(ctx); err != nil {
	//         return err
	//     }
	// }
	return nil
}

func (a *Guard) RegisterV1Routers(ctx context.Context, v1 *gin.RouterGroup) error {

	// 防护：接收者或关键子对象为 nil 时返回明确错误，避免空指针 panic
	if a == nil {
		return fmt.Errorf("guard module is not initialized")
	}
	if v1 == nil {
		return fmt.Errorf("router group is nil")
	}
	if a.StaffAPI == nil {
		return fmt.Errorf("guard.StaffAPI is nil")
	}

	// 工作人员管理
	staff := v1.Group("staff")
	{
		staff.GET("", a.StaffAPI.Query)
		staff.GET(":id", a.StaffAPI.Get)
		staff.POST("", a.StaffAPI.Create)
		staff.PUT(":id", a.StaffAPI.Update)
		staff.DELETE(":id", a.StaffAPI.Delete)
		staff.POST("import", a.StaffAPI.Import)
		staff.GET("export", a.StaffAPI.Export)
		staff.GET("export-template", a.StaffAPI.ExportTemplate)
	}

	// 场景管理
	scene := v1.Group("scene")
	{
		scene.GET("", a.SceneAPI.Query)
		scene.GET(":id", a.SceneAPI.Get)
		scene.POST("", a.SceneAPI.Create)
		scene.PUT(":id", a.SceneAPI.Update)
		scene.DELETE(":id", a.SceneAPI.Delete)
		scene.POST("merge", a.SceneAPI.Merge)
		scene.POST("split", a.SceneAPI.Split)
		scene.GET("statistics", a.SceneAPI.Statistics)
		scene.GET("centerpoint/:id", a.SceneAPI.GetCenterPoint)
		scene.POST("import", a.SceneAPI.Import)
		scene.GET("export", a.SceneAPI.Export)
	}

	// 标绘管理
	plotting := v1.Group("plotting")
	{
		plotting.GET("", a.PlottingAPI.Query)
		plotting.GET(":id", a.PlottingAPI.Get)
		plotting.POST("", a.PlottingAPI.Create)
		plotting.PUT(":id", a.PlottingAPI.Update)
		plotting.DELETE(":id", a.PlottingAPI.Delete)
		plotting.POST("sort", a.PlottingAPI.Sort)
		plotting.POST("statistics", a.PlottingAPI.Statistics)
		plotting.POST("calculate", a.PlottingAPI.Calculate)
	}

	// TODO: 其他API逐步实现（Watchkeeper, TargetedIndividuals, DogTrainer等）
	/*
		watchkeeper := v1.Group("watchkeeper") { ... }
		targetedIndividuals := v1.Group("targeted-individuals") { ... }
		policeDog := v1.Group("police-dog") { ... }
		car := v1.Group("car") { ... }
	*/

	return nil
}

func (a *Guard) Release(ctx context.Context) error {
	return nil
}
