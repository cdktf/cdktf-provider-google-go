package computeprojectmetadata


type ComputeProjectMetadataTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_project_metadata#create ComputeProjectMetadata#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_project_metadata#delete ComputeProjectMetadata#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

