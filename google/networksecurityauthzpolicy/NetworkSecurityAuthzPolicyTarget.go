// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecurityauthzpolicy


type NetworkSecurityAuthzPolicyTarget struct {
	// All gateways and forwarding rules referenced by this policy and extensions must share the same load balancing scheme.
	//
	// For more information, refer to [Backend services overview](https://cloud.google.com/load-balancing/docs/backend-service). Possible values: ["INTERNAL_MANAGED", "EXTERNAL_MANAGED", "INTERNAL_SELF_MANAGED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/network_security_authz_policy#load_balancing_scheme NetworkSecurityAuthzPolicy#load_balancing_scheme}
	LoadBalancingScheme *string `field:"required" json:"loadBalancingScheme" yaml:"loadBalancingScheme"`
	// A list of references to the Forwarding Rules on which this policy will be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/network_security_authz_policy#resources NetworkSecurityAuthzPolicy#resources}
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
}

