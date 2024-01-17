// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamworkforcepool


type IamWorkforcePoolAccessRestrictions struct {
	// allowed_services block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.12.0/docs/resources/iam_workforce_pool#allowed_services IamWorkforcePool#allowed_services}
	AllowedServices interface{} `field:"optional" json:"allowedServices" yaml:"allowedServices"`
	// Disable programmatic sign-in by disabling token issue via the Security Token API endpoint. See [Security Token Service API](https://cloud.google.com/iam/docs/reference/sts/rest).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.12.0/docs/resources/iam_workforce_pool#disable_programmatic_signin IamWorkforcePool#disable_programmatic_signin}
	DisableProgrammaticSignin interface{} `field:"optional" json:"disableProgrammaticSignin" yaml:"disableProgrammaticSignin"`
}

