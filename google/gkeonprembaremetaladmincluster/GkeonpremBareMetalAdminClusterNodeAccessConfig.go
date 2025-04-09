// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetaladmincluster


type GkeonpremBareMetalAdminClusterNodeAccessConfig struct {
	// LoginUser is the user name used to access node machines. It defaults to "root" if not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/gkeonprem_bare_metal_admin_cluster#login_user GkeonpremBareMetalAdminCluster#login_user}
	LoginUser *string `field:"optional" json:"loginUser" yaml:"loginUser"`
}

