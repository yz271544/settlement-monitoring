package guard

import (
	"github.com/google/wire"
	"github.com/yz271544/settlement-monitoring/v10/internal/mods/guard/api"
	"github.com/yz271544/settlement-monitoring/v10/internal/mods/guard/biz"
	"github.com/yz271544/settlement-monitoring/v10/internal/mods/guard/dal"
)

// Collection of wire providers
var Set = wire.NewSet(
	wire.Struct(new(Guard), "*"),

	// Staff - 已实现
	wire.Struct(new(dal.Staff), "*"),
	wire.Struct(new(biz.Staff), "*"),
	wire.Struct(new(api.Staff), "*"),

	// TODO: 其他表逐步实现
	/*
		wire.Struct(new(dal.Watchkeeper), "*"),
		wire.Struct(new(biz.Watchkeeper), "*"),
		wire.Struct(new(api.Watchkeeper), "*"),
		wire.Struct(new(dal.TargetedIndividuals), "*"),
		wire.Struct(new(biz.TargetedIndividuals), "*"),
		wire.Struct(new(api.TargetedIndividuals), "*"),
		wire.Struct(new(dal.DogTrainer), "*"),
		wire.Struct(new(biz.DogTrainer), "*"),
		wire.Struct(new(api.DogTrainer), "*"),
		wire.Struct(new(dal.PoliceDog), "*"),
		wire.Struct(new(biz.PoliceDog), "*"),
		wire.Struct(new(api.PoliceDog), "*"),
		wire.Struct(new(dal.Car), "*"),
		wire.Struct(new(biz.Car), "*"),
		wire.Struct(new(api.Car), "*"),
		wire.Struct(new(dal.Firearm), "*"),
		wire.Struct(new(biz.Firearm), "*"),
		wire.Struct(new(api.Firearm), "*"),
		wire.Struct(new(dal.Interphone), "*"),
		wire.Struct(new(biz.Interphone), "*"),
		wire.Struct(new(api.Interphone), "*"),
		wire.Struct(new(dal.DroneCounterEquipment), "*"),
		wire.Struct(new(biz.DroneCounterEquipment), "*"),
		wire.Struct(new(api.DroneCounterEquipment), "*"),
		wire.Struct(new(dal.SecurityScreeningEquipment), "*"),
		wire.Struct(new(biz.SecurityScreeningEquipment), "*"),
		wire.Struct(new(api.SecurityScreeningEquipment), "*"),
		wire.Struct(new(dal.PoliceRecorder), "*"),
		wire.Struct(new(biz.PoliceRecorder), "*"),
		wire.Struct(new(api.PoliceRecorder), "*"),
		wire.Struct(new(dal.VantagePoint), "*"),
		wire.Struct(new(biz.VantagePoint), "*"),
		wire.Struct(new(api.VantagePoint), "*"),
		wire.Struct(new(dal.KeyPart), "*"),
		wire.Struct(new(biz.KeyPart), "*"),
		wire.Struct(new(api.KeyPart), "*"),
		wire.Struct(new(dal.DangerousPart), "*"),
		wire.Struct(new(biz.DangerousPart), "*"),
		wire.Struct(new(api.DangerousPart), "*"),
		wire.Struct(new(dal.Scene), "*"),
		wire.Struct(new(biz.Scene), "*"),
		wire.Struct(new(api.Scene), "*"),
		wire.Struct(new(dal.Plotting), "*"),
		wire.Struct(new(biz.Plotting), "*"),
		wire.Struct(new(api.Plotting), "*"),
	*/
)
