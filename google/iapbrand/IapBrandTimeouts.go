// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iapbrand


type IapBrandTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/iap_brand#create IapBrand#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/iap_brand#delete IapBrand#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

