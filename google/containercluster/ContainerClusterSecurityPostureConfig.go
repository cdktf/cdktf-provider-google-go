// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterSecurityPostureConfig struct {
	// Sets the mode of the Kubernetes security posture API's off-cluster features. Available options include DISABLED, BASIC, and ENTERPRISE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/container_cluster#mode ContainerCluster#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Sets the mode of the Kubernetes security posture API's workload vulnerability scanning. Available options include VULNERABILITY_DISABLED, VULNERABILITY_BASIC and VULNERABILITY_ENTERPRISE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/container_cluster#vulnerability_mode ContainerCluster#vulnerability_mode}
	VulnerabilityMode *string `field:"optional" json:"vulnerabilityMode" yaml:"vulnerabilityMode"`
}

