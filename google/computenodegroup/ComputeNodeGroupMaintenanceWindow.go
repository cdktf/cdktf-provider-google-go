package computenodegroup


type ComputeNodeGroupMaintenanceWindow struct {
	// instances.start time of the window. This must be in UTC format that resolves to one of 00:00, 04:00, 08:00, 12:00, 16:00, or 20:00. For example, both 13:00-5 and 08:00 are valid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_node_group#start_time ComputeNodeGroup#start_time}
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}

