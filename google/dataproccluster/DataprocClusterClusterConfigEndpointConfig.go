// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataproccluster


type DataprocClusterClusterConfigEndpointConfig struct {
	// The flag to enable http access to specific ports on the cluster from external sources (aka Component Gateway).
	//
	// Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.16.0/docs/resources/dataproc_cluster#enable_http_port_access DataprocCluster#enable_http_port_access}
	EnableHttpPortAccess interface{} `field:"required" json:"enableHttpPortAccess" yaml:"enableHttpPortAccess"`
}

