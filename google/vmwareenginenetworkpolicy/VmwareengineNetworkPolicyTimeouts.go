// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmwareenginenetworkpolicy


type VmwareengineNetworkPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/vmwareengine_network_policy#create VmwareengineNetworkPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/vmwareengine_network_policy#delete VmwareengineNetworkPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/vmwareengine_network_policy#update VmwareengineNetworkPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

