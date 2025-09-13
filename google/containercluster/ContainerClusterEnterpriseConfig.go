// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterEnterpriseConfig struct {
	// Indicates the desired cluster tier. Available options include STANDARD and ENTERPRISE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_cluster#desired_tier ContainerCluster#desired_tier}
	DesiredTier *string `field:"optional" json:"desiredTier" yaml:"desiredTier"`
}

