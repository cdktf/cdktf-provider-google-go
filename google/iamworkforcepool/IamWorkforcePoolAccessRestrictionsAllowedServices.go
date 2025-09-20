// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamworkforcepool


type IamWorkforcePoolAccessRestrictionsAllowedServices struct {
	// Domain name of the service. Example: console.cloud.google.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/iam_workforce_pool#domain IamWorkforcePool#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
}

