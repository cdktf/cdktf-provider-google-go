// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alloydbcluster


type AlloydbClusterNetworkConfig struct {
	// The name of the allocated IP range for the private IP AlloyDB cluster.
	//
	// For example: "google-managed-services-default".
	// If set, the instance IPs for this cluster will be created in the allocated range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.84.0/docs/resources/alloydb_cluster#allocated_ip_range AlloydbCluster#allocated_ip_range}
	AllocatedIpRange *string `field:"optional" json:"allocatedIpRange" yaml:"allocatedIpRange"`
	// The resource link for the VPC network in which cluster resources are created and from which they are accessible via Private IP.
	//
	// The network must belong to the same project as the cluster.
	// It is specified in the form: "projects/{projectNumber}/global/networks/{network_id}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.84.0/docs/resources/alloydb_cluster#network AlloydbCluster#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
}

