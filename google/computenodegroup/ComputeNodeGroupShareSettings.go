package computenodegroup


type ComputeNodeGroupShareSettings struct {
	// Node group sharing type. Possible values: ["ORGANIZATION", "SPECIFIC_PROJECTS", "LOCAL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_node_group#share_type ComputeNodeGroup#share_type}
	ShareType *string `field:"required" json:"shareType" yaml:"shareType"`
	// project_map block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_node_group#project_map ComputeNodeGroup#project_map}
	ProjectMap interface{} `field:"optional" json:"projectMap" yaml:"projectMap"`
}

