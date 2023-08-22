package computeinterconnectattachment


type ComputeInterconnectAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_interconnect_attachment#create ComputeInterconnectAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_interconnect_attachment#delete ComputeInterconnectAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_interconnect_attachment#update ComputeInterconnectAttachment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

