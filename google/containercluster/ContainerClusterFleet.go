// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterFleet struct {
	// The Fleet host project of the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/container_cluster#project ContainerCluster#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

