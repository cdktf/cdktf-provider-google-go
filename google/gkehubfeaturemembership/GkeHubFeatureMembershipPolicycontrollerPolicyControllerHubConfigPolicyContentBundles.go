// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeaturemembership


type GkeHubFeatureMembershipPolicycontrollerPolicyControllerHubConfigPolicyContentBundles struct {
	// The name for the key in the map for which this object is mapped to in the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.34.0/docs/resources/gke_hub_feature_membership#bundle_name GkeHubFeatureMembership#bundle_name}
	BundleName *string `field:"required" json:"bundleName" yaml:"bundleName"`
	// The set of namespaces to be exempted from the bundle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.34.0/docs/resources/gke_hub_feature_membership#exempted_namespaces GkeHubFeatureMembership#exempted_namespaces}
	ExemptedNamespaces *[]*string `field:"optional" json:"exemptedNamespaces" yaml:"exemptedNamespaces"`
}

