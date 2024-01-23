// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeature


type GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigPolicyContentTemplateLibrary struct {
	// Configures the manner in which the template library is installed on the cluster. Possible values: ["INSTALATION_UNSPECIFIED", "NOT_INSTALLED", "ALL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/gke_hub_feature#installation GkeHubFeature#installation}
	Installation *string `field:"optional" json:"installation" yaml:"installation"`
}

