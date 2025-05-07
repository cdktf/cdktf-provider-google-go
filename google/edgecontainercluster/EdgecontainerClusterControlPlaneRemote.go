// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package edgecontainercluster


type EdgecontainerClusterControlPlaneRemote struct {
	// Name of the Google Distributed Cloud Edge zones where this node pool will be created. For example: 'us-central1-edge-customer-a'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/edgecontainer_cluster#node_location EdgecontainerCluster#node_location}
	NodeLocation *string `field:"optional" json:"nodeLocation" yaml:"nodeLocation"`
}

