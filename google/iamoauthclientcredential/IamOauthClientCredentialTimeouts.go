// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamoauthclientcredential


type IamOauthClientCredentialTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/iam_oauth_client_credential#create IamOauthClientCredential#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/iam_oauth_client_credential#delete IamOauthClientCredential#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/iam_oauth_client_credential#update IamOauthClientCredential#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

