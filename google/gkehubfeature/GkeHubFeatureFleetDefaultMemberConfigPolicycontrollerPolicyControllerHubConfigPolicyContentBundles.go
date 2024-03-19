// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeature


type GkeHubFeatureFleetDefaultMemberConfigPolicycontrollerPolicyControllerHubConfigPolicyContentBundles struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/gke_hub_feature#bundle GkeHubFeature#bundle}.
	Bundle *string `field:"required" json:"bundle" yaml:"bundle"`
	// The set of namespaces to be exempted from the bundle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/gke_hub_feature#exempted_namespaces GkeHubFeature#exempted_namespaces}
	ExemptedNamespaces *[]*string `field:"optional" json:"exemptedNamespaces" yaml:"exemptedNamespaces"`
}

