// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package edgecontainercluster


type EdgecontainerClusterAuthorizationAdminUsers struct {
	// An active Google username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/edgecontainer_cluster#username EdgecontainerCluster#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

