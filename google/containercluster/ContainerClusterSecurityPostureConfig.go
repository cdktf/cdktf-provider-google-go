package containercluster


type ContainerClusterSecurityPostureConfig struct {
	// Sets the mode of the Kubernetes security posture API's off-cluster features. Available options include DISABLED and BASIC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#mode ContainerCluster#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Sets the mode of the Kubernetes security posture API's workload vulnerability scanning. Available options include VULNERABILITY_DISABLED and VULNERABILITY_BASIC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#vulnerability_mode ContainerCluster#vulnerability_mode}
	VulnerabilityMode *string `field:"optional" json:"vulnerabilityMode" yaml:"vulnerabilityMode"`
}

