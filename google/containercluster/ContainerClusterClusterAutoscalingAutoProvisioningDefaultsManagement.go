package containercluster


type ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagement struct {
	// Specifies whether the node auto-repair is enabled for the node pool.
	//
	// If enabled, the nodes in this node pool will be monitored and, if they fail health checks too many times, an automatic repair action will be triggered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#auto_repair ContainerCluster#auto_repair}
	AutoRepair interface{} `field:"optional" json:"autoRepair" yaml:"autoRepair"`
	// Specifies whether node auto-upgrade is enabled for the node pool.
	//
	// If enabled, node auto-upgrade helps keep the nodes in your node pool up to date with the latest release version of Kubernetes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#auto_upgrade ContainerCluster#auto_upgrade}
	AutoUpgrade interface{} `field:"optional" json:"autoUpgrade" yaml:"autoUpgrade"`
}

