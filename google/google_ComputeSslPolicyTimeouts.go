// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ComputeSslPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_ssl_policy#create ComputeSslPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_ssl_policy#delete ComputeSslPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_ssl_policy#update ComputeSslPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

