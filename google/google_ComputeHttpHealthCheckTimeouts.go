// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ComputeHttpHealthCheckTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_http_health_check#create ComputeHttpHealthCheck#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_http_health_check#delete ComputeHttpHealthCheck#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_http_health_check#update ComputeHttpHealthCheck#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

