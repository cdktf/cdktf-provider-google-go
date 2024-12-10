// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamworkloadidentitypoolprovider


type IamWorkloadIdentityPoolProviderX509TrustStore struct {
	// trust_anchors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/iam_workload_identity_pool_provider#trust_anchors IamWorkloadIdentityPoolProvider#trust_anchors}
	TrustAnchors interface{} `field:"required" json:"trustAnchors" yaml:"trustAnchors"`
	// intermediate_cas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/iam_workload_identity_pool_provider#intermediate_cas IamWorkloadIdentityPoolProvider#intermediate_cas}
	IntermediateCas interface{} `field:"optional" json:"intermediateCas" yaml:"intermediateCas"`
}

