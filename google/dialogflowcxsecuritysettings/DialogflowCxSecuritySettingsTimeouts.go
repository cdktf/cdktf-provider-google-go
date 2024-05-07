// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxsecuritysettings


type DialogflowCxSecuritySettingsTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.28.0/docs/resources/dialogflow_cx_security_settings#create DialogflowCxSecuritySettings#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.28.0/docs/resources/dialogflow_cx_security_settings#delete DialogflowCxSecuritySettings#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.28.0/docs/resources/dialogflow_cx_security_settings#update DialogflowCxSecuritySettings#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

