package computenodegroup


type ComputeNodeGroupAutoscalingPolicy struct {
	// Maximum size of the node group.
	//
	// Set to a value less than or equal
	// to 100 and greater than or equal to min-nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_node_group#max_nodes ComputeNodeGroup#max_nodes}
	MaxNodes *float64 `field:"optional" json:"maxNodes" yaml:"maxNodes"`
	// Minimum size of the node group. Must be less than or equal to max-nodes. The default value is 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_node_group#min_nodes ComputeNodeGroup#min_nodes}
	MinNodes *float64 `field:"optional" json:"minNodes" yaml:"minNodes"`
	// The autoscaling mode.
	//
	// Set to one of the following:
	// - OFF: Disables the autoscaler.
	// - ON: Enables scaling in and scaling out.
	// - ONLY_SCALE_OUT: Enables only scaling out.
	// You must use this mode if your node groups are configured to
	// restart their hosted VMs on minimal servers. Possible values: ["OFF", "ON", "ONLY_SCALE_OUT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_node_group#mode ComputeNodeGroup#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

