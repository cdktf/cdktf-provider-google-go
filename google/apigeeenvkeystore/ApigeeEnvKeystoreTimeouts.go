// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeeenvkeystore


type ApigeeEnvKeystoreTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/apigee_env_keystore#create ApigeeEnvKeystore#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/apigee_env_keystore#delete ApigeeEnvKeystore#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

