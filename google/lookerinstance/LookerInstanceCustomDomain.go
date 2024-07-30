// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lookerinstance


type LookerInstanceCustomDomain struct {
	// Domain name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/looker_instance#domain LookerInstance#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
}

