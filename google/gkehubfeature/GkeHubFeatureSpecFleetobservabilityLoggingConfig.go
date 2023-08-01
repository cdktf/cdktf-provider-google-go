package gkehubfeature


type GkeHubFeatureSpecFleetobservabilityLoggingConfig struct {
	// default_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/gke_hub_feature#default_config GkeHubFeature#default_config}
	DefaultConfig *GkeHubFeatureSpecFleetobservabilityLoggingConfigDefaultConfig `field:"optional" json:"defaultConfig" yaml:"defaultConfig"`
	// fleet_scope_logs_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/gke_hub_feature#fleet_scope_logs_config GkeHubFeature#fleet_scope_logs_config}
	FleetScopeLogsConfig *GkeHubFeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfig `field:"optional" json:"fleetScopeLogsConfig" yaml:"fleetScopeLogsConfig"`
}

