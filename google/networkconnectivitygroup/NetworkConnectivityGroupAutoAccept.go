// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkconnectivitygroup


type NetworkConnectivityGroupAutoAccept struct {
	// A list of project ids or project numbers for which you want to enable auto-accept.
	//
	// The auto-accept setting is applied to spokes being created or updated in these projects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_connectivity_group#auto_accept_projects NetworkConnectivityGroup#auto_accept_projects}
	AutoAcceptProjects *[]*string `field:"required" json:"autoAcceptProjects" yaml:"autoAcceptProjects"`
}

