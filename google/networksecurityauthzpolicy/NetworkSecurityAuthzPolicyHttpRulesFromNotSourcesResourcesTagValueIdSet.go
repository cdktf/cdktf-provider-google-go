// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecurityauthzpolicy


type NetworkSecurityAuthzPolicyHttpRulesFromNotSourcesResourcesTagValueIdSet struct {
	// A list of resource tag value permanent IDs to match against the resource manager tags value associated with the source VM of a request.
	//
	// The match follows AND semantics which means all the ids must match.
	// Limited to 5 matches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/network_security_authz_policy#ids NetworkSecurityAuthzPolicy#ids}
	Ids *[]*string `field:"optional" json:"ids" yaml:"ids"`
}

