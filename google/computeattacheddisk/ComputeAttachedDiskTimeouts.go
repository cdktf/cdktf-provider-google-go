package computeattacheddisk


type ComputeAttachedDiskTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_attached_disk#create ComputeAttachedDisk#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_attached_disk#delete ComputeAttachedDisk#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

