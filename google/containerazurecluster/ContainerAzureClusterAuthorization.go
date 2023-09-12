// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerazurecluster


type ContainerAzureClusterAuthorization struct {
	// admin_users block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.82.0/docs/resources/container_azure_cluster#admin_users ContainerAzureCluster#admin_users}
	AdminUsers interface{} `field:"required" json:"adminUsers" yaml:"adminUsers"`
}

