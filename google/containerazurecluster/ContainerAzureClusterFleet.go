// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerazurecluster


type ContainerAzureClusterFleet struct {
	// The number of the Fleet host project where this cluster will be registered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/container_azure_cluster#project ContainerAzureCluster#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

