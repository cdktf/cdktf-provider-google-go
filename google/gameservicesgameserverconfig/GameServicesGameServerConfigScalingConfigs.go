package gameservicesgameserverconfig


type GameServicesGameServerConfigScalingConfigs struct {
	// Fleet autoscaler spec, which is sent to Agones. Example spec can be found : https://agones.dev/site/docs/reference/fleetautoscaler/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/game_services_game_server_config#fleet_autoscaler_spec GameServicesGameServerConfig#fleet_autoscaler_spec}
	FleetAutoscalerSpec *string `field:"required" json:"fleetAutoscalerSpec" yaml:"fleetAutoscalerSpec"`
	// The name of the ScalingConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/game_services_game_server_config#name GameServicesGameServerConfig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// schedules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/game_services_game_server_config#schedules GameServicesGameServerConfig#schedules}
	Schedules interface{} `field:"optional" json:"schedules" yaml:"schedules"`
	// selectors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/game_services_game_server_config#selectors GameServicesGameServerConfig#selectors}
	Selectors interface{} `field:"optional" json:"selectors" yaml:"selectors"`
}

