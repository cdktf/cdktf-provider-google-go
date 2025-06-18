// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamworkforcepoolproviderkey


type IamWorkforcePoolProviderKeyKeyData struct {
	// The specifications for the key. Possible values: ["RSA_2048", "RSA_3072", "RSA_4096"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/iam_workforce_pool_provider_key#key_spec IamWorkforcePoolProviderKey#key_spec}
	KeySpec *string `field:"required" json:"keySpec" yaml:"keySpec"`
}

