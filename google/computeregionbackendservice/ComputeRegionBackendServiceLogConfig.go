package computeregionbackendservice


type ComputeRegionBackendServiceLogConfig struct {
	// Whether to enable logging for the load balancer traffic served by this backend service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_backend_service#enable ComputeRegionBackendService#enable}
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// This field can only be specified if logging is enabled for this backend service.
	//
	// The value of
	// the field must be in [0, 1]. This configures the sampling rate of requests to the load balancer
	// where 1.0 means all logged requests are reported and 0.0 means no logged requests are reported.
	// The default value is 1.0.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_backend_service#sample_rate ComputeRegionBackendService#sample_rate}
	SampleRate *float64 `field:"optional" json:"sampleRate" yaml:"sampleRate"`
}

