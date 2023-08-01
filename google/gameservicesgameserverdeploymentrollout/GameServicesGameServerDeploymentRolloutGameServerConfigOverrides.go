package gameservicesgameserverdeploymentrollout


type GameServicesGameServerDeploymentRolloutGameServerConfigOverrides struct {
	// Version of the configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/game_services_game_server_deployment_rollout#config_version GameServicesGameServerDeploymentRollout#config_version}
	ConfigVersion *string `field:"optional" json:"configVersion" yaml:"configVersion"`
	// realms_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/game_services_game_server_deployment_rollout#realms_selector GameServicesGameServerDeploymentRollout#realms_selector}
	RealmsSelector *GameServicesGameServerDeploymentRolloutGameServerConfigOverridesRealmsSelector `field:"optional" json:"realmsSelector" yaml:"realmsSelector"`
}

