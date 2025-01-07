// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataproccluster


type DataprocClusterClusterConfigGceClusterConfigNodeGroupAffinity struct {
	// The URI of a sole-tenant that the cluster will be created on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/dataproc_cluster#node_group_uri DataprocCluster#node_group_uri}
	NodeGroupUri *string `field:"required" json:"nodeGroupUri" yaml:"nodeGroupUri"`
}

