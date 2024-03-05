// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectdefaultserviceaccounts


type ProjectDefaultServiceAccountsTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/project_default_service_accounts#create ProjectDefaultServiceAccounts#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/project_default_service_accounts#delete ProjectDefaultServiceAccounts#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/project_default_service_accounts#read ProjectDefaultServiceAccounts#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

