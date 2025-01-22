// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecurityauthzpolicy


type NetworkSecurityAuthzPolicyHttpRulesFromSourcesResources struct {
	// iam_service_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/network_security_authz_policy#iam_service_account NetworkSecurityAuthzPolicy#iam_service_account}
	IamServiceAccount *NetworkSecurityAuthzPolicyHttpRulesFromSourcesResourcesIamServiceAccount `field:"optional" json:"iamServiceAccount" yaml:"iamServiceAccount"`
	// tag_value_id_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/network_security_authz_policy#tag_value_id_set NetworkSecurityAuthzPolicy#tag_value_id_set}
	TagValueIdSet *NetworkSecurityAuthzPolicyHttpRulesFromSourcesResourcesTagValueIdSet `field:"optional" json:"tagValueIdSet" yaml:"tagValueIdSet"`
}

