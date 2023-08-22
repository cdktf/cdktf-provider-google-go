package monitoringslo


type MonitoringSloBasicSliAvailability struct {
	// Whether an availability SLI is enabled or not. Must be set to true. Defaults to 'true'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/monitoring_slo#enabled MonitoringSlo#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

