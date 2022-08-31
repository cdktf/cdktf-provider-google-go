// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ComputeNodeTemplateTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_node_template#create ComputeNodeTemplate#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_node_template#delete ComputeNodeTemplate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

