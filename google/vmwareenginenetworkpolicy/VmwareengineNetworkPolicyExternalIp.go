// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmwareenginenetworkpolicy


type VmwareengineNetworkPolicyExternalIp struct {
	// True if the service is enabled; false otherwise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/vmwareengine_network_policy#enabled VmwareengineNetworkPolicy#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

