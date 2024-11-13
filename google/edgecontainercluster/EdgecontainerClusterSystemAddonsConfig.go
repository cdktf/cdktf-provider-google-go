// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package edgecontainercluster


type EdgecontainerClusterSystemAddonsConfig struct {
	// ingress block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/edgecontainer_cluster#ingress EdgecontainerCluster#ingress}
	Ingress *EdgecontainerClusterSystemAddonsConfigIngress `field:"optional" json:"ingress" yaml:"ingress"`
}

