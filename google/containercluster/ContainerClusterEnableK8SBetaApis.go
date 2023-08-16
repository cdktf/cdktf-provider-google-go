package containercluster


type ContainerClusterEnableK8SBetaApis struct {
	// Enabled Kubernetes Beta APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#enabled_apis ContainerCluster#enabled_apis}
	EnabledApis *[]*string `field:"required" json:"enabledApis" yaml:"enabledApis"`
}

