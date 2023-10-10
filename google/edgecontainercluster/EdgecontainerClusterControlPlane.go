// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package edgecontainercluster


type EdgecontainerClusterControlPlane struct {
	// local block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.1.0/docs/resources/edgecontainer_cluster#local EdgecontainerCluster#local}
	Local *EdgecontainerClusterControlPlaneLocal `field:"optional" json:"local" yaml:"local"`
	// remote block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.1.0/docs/resources/edgecontainer_cluster#remote EdgecontainerCluster#remote}
	Remote *EdgecontainerClusterControlPlaneRemote `field:"optional" json:"remote" yaml:"remote"`
}

