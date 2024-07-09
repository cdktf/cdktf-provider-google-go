// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocmetastoreservice


type DataprocMetastoreServiceNetworkConfigConsumers struct {
	// The subnetwork of the customer project from which an IP address is reserved and used as the Dataproc Metastore service's endpoint.
	//
	// It is accessible to hosts in the subnet and to all hosts in a subnet in the same region and same network.
	// There must be at least one IP address available in the subnet's primary range. The subnet is specified in the following form:
	// 'projects/{projectNumber}/regions/{region_id}/subnetworks/{subnetwork_id}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/dataproc_metastore_service#subnetwork DataprocMetastoreService#subnetwork}
	Subnetwork *string `field:"required" json:"subnetwork" yaml:"subnetwork"`
}

