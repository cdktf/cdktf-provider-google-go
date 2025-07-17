// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmwareengineexternalaccessrule


type VmwareengineExternalAccessRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/vmwareengine_external_access_rule#create VmwareengineExternalAccessRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/vmwareengine_external_access_rule#delete VmwareengineExternalAccessRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/vmwareengine_external_access_rule#update VmwareengineExternalAccessRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

