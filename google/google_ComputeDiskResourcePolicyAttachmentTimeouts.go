// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ComputeDiskResourcePolicyAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_disk_resource_policy_attachment#create ComputeDiskResourcePolicyAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_disk_resource_policy_attachment#delete ComputeDiskResourcePolicyAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

