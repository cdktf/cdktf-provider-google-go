// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubnamespace


type GkeHubNamespaceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/gke_hub_namespace#create GkeHubNamespace#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/gke_hub_namespace#delete GkeHubNamespace#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/gke_hub_namespace#update GkeHubNamespace#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

