// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ComputeBackendServiceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_backend_service#create ComputeBackendService#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_backend_service#delete ComputeBackendService#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_backend_service#update ComputeBackendService#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

