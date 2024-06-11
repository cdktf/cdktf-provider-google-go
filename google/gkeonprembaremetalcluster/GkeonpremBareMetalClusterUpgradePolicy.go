// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalcluster


type GkeonpremBareMetalClusterUpgradePolicy struct {
	// Specifies which upgrade policy to use. Possible values: ["SERIAL", "CONCURRENT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.33.0/docs/resources/gkeonprem_bare_metal_cluster#policy GkeonpremBareMetalCluster#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

