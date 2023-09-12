// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamworkforcepoolprovider


type IamWorkforcePoolProviderOidcClientSecretValue struct {
	// The plain text of the client secret value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.82.0/docs/resources/iam_workforce_pool_provider#plain_text IamWorkforcePoolProvider#plain_text}
	PlainText *string `field:"required" json:"plainText" yaml:"plainText"`
}

