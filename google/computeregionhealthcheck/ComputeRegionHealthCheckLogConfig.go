package computeregionhealthcheck


type ComputeRegionHealthCheckLogConfig struct {
	// Indicates whether or not to export logs.
	//
	// This is false by default,
	// which means no health check logging will be done.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_region_health_check#enable ComputeRegionHealthCheck#enable}
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
}

