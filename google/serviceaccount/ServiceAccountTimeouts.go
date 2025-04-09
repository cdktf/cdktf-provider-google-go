// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceaccount


type ServiceAccountTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/service_account#create ServiceAccount#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

