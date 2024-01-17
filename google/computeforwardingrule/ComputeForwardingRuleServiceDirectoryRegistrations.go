// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeforwardingrule


type ComputeForwardingRuleServiceDirectoryRegistrations struct {
	// Service Directory namespace to register the forwarding rule under.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.12.0/docs/resources/compute_forwarding_rule#namespace ComputeForwardingRule#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Service Directory service to register the forwarding rule under.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.12.0/docs/resources/compute_forwarding_rule#service ComputeForwardingRule#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

