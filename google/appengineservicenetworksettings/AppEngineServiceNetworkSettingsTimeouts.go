package appengineservicenetworksettings


type AppEngineServiceNetworkSettingsTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/app_engine_service_network_settings#create AppEngineServiceNetworkSettings#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/app_engine_service_network_settings#delete AppEngineServiceNetworkSettings#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/app_engine_service_network_settings#update AppEngineServiceNetworkSettings#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

