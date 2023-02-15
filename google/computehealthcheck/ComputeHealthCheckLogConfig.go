package computehealthcheck


type ComputeHealthCheckLogConfig struct {
	// Indicates whether or not to export logs.
	//
	// This is false by default,
	// which means no health check logging will be done.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_health_check#enable ComputeHealthCheck#enable}
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
}

