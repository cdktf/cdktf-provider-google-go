// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sccmanagementprojectsecurityhealthanalyticscustommodule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SccManagementProjectSecurityHealthAnalyticsCustomModuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/scc_management_project_security_health_analytics_custom_module#custom_config SccManagementProjectSecurityHealthAnalyticsCustomModule#custom_config}
	CustomConfig *SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfig `field:"optional" json:"customConfig" yaml:"customConfig"`
	// The display name of the Security Health Analytics custom module.
	//
	// This
	// display name becomes the finding category for all findings that are
	// returned by this custom module. The display name must be between 1 and
	// 128 characters, start with a lowercase letter, and contain alphanumeric
	// characters or underscores only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/scc_management_project_security_health_analytics_custom_module#display_name SccManagementProjectSecurityHealthAnalyticsCustomModule#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The enablement state of the custom module. Possible values: ["ENABLED", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/scc_management_project_security_health_analytics_custom_module#enablement_state SccManagementProjectSecurityHealthAnalyticsCustomModule#enablement_state}
	EnablementState *string `field:"optional" json:"enablementState" yaml:"enablementState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/scc_management_project_security_health_analytics_custom_module#id SccManagementProjectSecurityHealthAnalyticsCustomModule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Location ID of the parent organization. If not provided, 'global' will be used as the default location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/scc_management_project_security_health_analytics_custom_module#location SccManagementProjectSecurityHealthAnalyticsCustomModule#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/scc_management_project_security_health_analytics_custom_module#project SccManagementProjectSecurityHealthAnalyticsCustomModule#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/scc_management_project_security_health_analytics_custom_module#timeouts SccManagementProjectSecurityHealthAnalyticsCustomModule#timeouts}
	Timeouts *SccManagementProjectSecurityHealthAnalyticsCustomModuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

