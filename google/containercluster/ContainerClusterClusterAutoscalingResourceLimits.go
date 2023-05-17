package containercluster


type ContainerClusterClusterAutoscalingResourceLimits struct {
	// The type of the resource.
	//
	// For example, cpu and memory. See the guide to using Node Auto-Provisioning for a list of types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_cluster#resource_type ContainerCluster#resource_type}
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Maximum amount of the resource in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_cluster#maximum ContainerCluster#maximum}
	Maximum *float64 `field:"optional" json:"maximum" yaml:"maximum"`
	// Minimum amount of the resource in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_cluster#minimum ContainerCluster#minimum}
	Minimum *float64 `field:"optional" json:"minimum" yaml:"minimum"`
}

