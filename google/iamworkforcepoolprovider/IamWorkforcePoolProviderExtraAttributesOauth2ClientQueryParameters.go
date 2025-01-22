// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamworkforcepoolprovider


type IamWorkforcePoolProviderExtraAttributesOauth2ClientQueryParameters struct {
	// The filter used to request specific records from IdP.
	//
	// In case of attributes type as AZURE_AD_GROUPS_MAIL, it represents the
	// filter used to request specific groups for users from IdP. By default, all of the groups associated with the user are fetched. The
	// groups should be mail enabled and security enabled. See https://learn.microsoft.com/en-us/graph/search-query-parameter for more details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/iam_workforce_pool_provider#filter IamWorkforcePoolProvider#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
}

