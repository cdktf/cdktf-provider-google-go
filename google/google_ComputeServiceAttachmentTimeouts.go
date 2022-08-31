// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ComputeServiceAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_service_attachment#create ComputeServiceAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_service_attachment#delete ComputeServiceAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_service_attachment#update ComputeServiceAttachment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

