// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sccprojectcustommodule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SccProjectCustomModuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// custom_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.5.0/docs/resources/scc_project_custom_module#custom_config SccProjectCustomModule#custom_config}
	CustomConfig *SccProjectCustomModuleCustomConfig `field:"required" json:"customConfig" yaml:"customConfig"`
	// The display name of the Security Health Analytics custom module.
	//
	// This
	// display name becomes the finding category for all findings that are
	// returned by this custom module. The display name must be between 1 and
	// 128 characters, start with a lowercase letter, and contain alphanumeric
	// characters or underscores only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.5.0/docs/resources/scc_project_custom_module#display_name SccProjectCustomModule#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The enablement state of the custom module. Possible values: ["ENABLED", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.5.0/docs/resources/scc_project_custom_module#enablement_state SccProjectCustomModule#enablement_state}
	EnablementState *string `field:"required" json:"enablementState" yaml:"enablementState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.5.0/docs/resources/scc_project_custom_module#id SccProjectCustomModule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.5.0/docs/resources/scc_project_custom_module#project SccProjectCustomModule#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.5.0/docs/resources/scc_project_custom_module#timeouts SccProjectCustomModule#timeouts}
	Timeouts *SccProjectCustomModuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

