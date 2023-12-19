// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterConfidentialNodes struct {
	// Whether Confidential Nodes feature is enabled for all nodes in this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

