// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ComputeRegionHealthCheckLogConfig struct {
	// Indicates whether or not to export logs.
	//
	// This is false by default,
	// which means no health check logging will be done.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_health_check#enable ComputeRegionHealthCheck#enable}
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
}

