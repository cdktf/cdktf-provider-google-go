package computeinstancegroupmanager


type ComputeInstanceGroupManagerAutoHealingPolicies struct {
	// The health check resource that signals autohealing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_instance_group_manager#health_check ComputeInstanceGroupManager#health_check}
	HealthCheck *string `field:"required" json:"healthCheck" yaml:"healthCheck"`
	// The number of seconds that the managed instance group waits before it applies autohealing policies to new instances or recently recreated instances.
	//
	// Between 0 and 3600.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_instance_group_manager#initial_delay_sec ComputeInstanceGroupManager#initial_delay_sec}
	InitialDelaySec *float64 `field:"required" json:"initialDelaySec" yaml:"initialDelaySec"`
}

