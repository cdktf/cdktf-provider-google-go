// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagoogleiampolicy


type DataGoogleIamPolicyAuditConfigAuditLogConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/data-sources/iam_policy#log_type DataGoogleIamPolicy#log_type}.
	LogType *string `field:"required" json:"logType" yaml:"logType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/data-sources/iam_policy#exempted_members DataGoogleIamPolicy#exempted_members}.
	ExemptedMembers *[]*string `field:"optional" json:"exemptedMembers" yaml:"exemptedMembers"`
}

