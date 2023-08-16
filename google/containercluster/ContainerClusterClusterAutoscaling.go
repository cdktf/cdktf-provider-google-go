package containercluster


type ContainerClusterClusterAutoscaling struct {
	// auto_provisioning_defaults block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#auto_provisioning_defaults ContainerCluster#auto_provisioning_defaults}
	AutoProvisioningDefaults *ContainerClusterClusterAutoscalingAutoProvisioningDefaults `field:"optional" json:"autoProvisioningDefaults" yaml:"autoProvisioningDefaults"`
	// Whether node auto-provisioning is enabled. Resource limits for cpu and memory must be defined to enable node auto-provisioning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// resource_limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#resource_limits ContainerCluster#resource_limits}
	ResourceLimits interface{} `field:"optional" json:"resourceLimits" yaml:"resourceLimits"`
}

