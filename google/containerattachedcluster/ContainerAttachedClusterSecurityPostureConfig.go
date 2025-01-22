// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerattachedcluster


type ContainerAttachedClusterSecurityPostureConfig struct {
	// Sets the mode of the Kubernetes security posture API's workload vulnerability scanning. Possible values: ["VULNERABILITY_DISABLED", "VULNERABILITY_ENTERPRISE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/container_attached_cluster#vulnerability_mode ContainerAttachedCluster#vulnerability_mode}
	VulnerabilityMode *string `field:"required" json:"vulnerabilityMode" yaml:"vulnerabilityMode"`
}

