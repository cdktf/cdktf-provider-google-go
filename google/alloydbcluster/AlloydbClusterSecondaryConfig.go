// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alloydbcluster


type AlloydbClusterSecondaryConfig struct {
	// Name of the primary cluster must be in the format 'projects/{project}/locations/{location}/clusters/{cluster_id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/alloydb_cluster#primary_cluster_name AlloydbCluster#primary_cluster_name}
	PrimaryClusterName *string `field:"required" json:"primaryClusterName" yaml:"primaryClusterName"`
}

