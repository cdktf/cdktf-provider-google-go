// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeaturemembership


type GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary struct {
	// Configures the manner in which the template library is installed on the cluster. Possible values: INSTALLATION_UNSPECIFIED, NOT_INSTALLED, ALL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/gke_hub_feature_membership#installation GkeHubFeatureMembership#installation}
	Installation *string `field:"optional" json:"installation" yaml:"installation"`
}

