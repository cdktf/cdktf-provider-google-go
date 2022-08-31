// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ComputeHttpsHealthCheckTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_https_health_check#create ComputeHttpsHealthCheck#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_https_health_check#delete ComputeHttpsHealthCheck#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_https_health_check#update ComputeHttpsHealthCheck#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

