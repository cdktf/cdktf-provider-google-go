package gkehubfeature


type GkeHubFeatureSpecFleetobservabilityLoggingConfigDefaultConfig struct {
	// Specified if fleet logging feature is enabled. Possible values: ["MODE_UNSPECIFIED", "COPY", "MOVE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/gke_hub_feature#mode GkeHubFeature#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

