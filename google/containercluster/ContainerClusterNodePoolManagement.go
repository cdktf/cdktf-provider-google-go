package containercluster


type ContainerClusterNodePoolManagement struct {
	// Whether the nodes will be automatically repaired.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#auto_repair ContainerCluster#auto_repair}
	AutoRepair interface{} `field:"optional" json:"autoRepair" yaml:"autoRepair"`
	// Whether the nodes will be automatically upgraded.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#auto_upgrade ContainerCluster#auto_upgrade}
	AutoUpgrade interface{} `field:"optional" json:"autoUpgrade" yaml:"autoUpgrade"`
}

