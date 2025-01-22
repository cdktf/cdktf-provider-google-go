// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secretmanagersecret


type SecretManagerSecretReplication struct {
	// auto block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/secret_manager_secret#auto SecretManagerSecret#auto}
	Auto *SecretManagerSecretReplicationAuto `field:"optional" json:"auto" yaml:"auto"`
	// user_managed block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/secret_manager_secret#user_managed SecretManagerSecret#user_managed}
	UserManaged *SecretManagerSecretReplicationUserManaged `field:"optional" json:"userManaged" yaml:"userManaged"`
}

