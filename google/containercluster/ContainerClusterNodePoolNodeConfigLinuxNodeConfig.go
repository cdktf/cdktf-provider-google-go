package containercluster


type ContainerClusterNodePoolNodeConfigLinuxNodeConfig struct {
	// The Linux kernel parameters to be applied to the nodes and all pods running on the nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.63.1/docs/resources/container_cluster#sysctls ContainerCluster#sysctls}
	Sysctls *map[string]*string `field:"required" json:"sysctls" yaml:"sysctls"`
}

