// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sccmanagementorganizationeventthreatdetectioncustommodule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SccManagementOrganizationEventThreatDetectionCustomModuleConfig struct {
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
	// Numerical ID of the parent organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/scc_management_organization_event_threat_detection_custom_module#organization SccManagementOrganizationEventThreatDetectionCustomModule#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// Config for the module.
	//
	// For the resident module, its config value is defined at this level.
	// For the inherited module, its config value is inherited from the ancestor module.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/scc_management_organization_event_threat_detection_custom_module#config SccManagementOrganizationEventThreatDetectionCustomModule#config}
	Config *string `field:"optional" json:"config" yaml:"config"`
	// The human readable name to be displayed for the module.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/scc_management_organization_event_threat_detection_custom_module#display_name SccManagementOrganizationEventThreatDetectionCustomModule#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The state of enablement for the module at the given level of the hierarchy. Possible values: ["ENABLED", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/scc_management_organization_event_threat_detection_custom_module#enablement_state SccManagementOrganizationEventThreatDetectionCustomModule#enablement_state}
	EnablementState *string `field:"optional" json:"enablementState" yaml:"enablementState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/scc_management_organization_event_threat_detection_custom_module#id SccManagementOrganizationEventThreatDetectionCustomModule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Location ID of the parent organization. Only global is supported at the moment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/scc_management_organization_event_threat_detection_custom_module#location SccManagementOrganizationEventThreatDetectionCustomModule#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/scc_management_organization_event_threat_detection_custom_module#timeouts SccManagementOrganizationEventThreatDetectionCustomModule#timeouts}
	Timeouts *SccManagementOrganizationEventThreatDetectionCustomModuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Immutable. Type for the module. e.g. CONFIGURABLE_BAD_IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/scc_management_organization_event_threat_detection_custom_module#type SccManagementOrganizationEventThreatDetectionCustomModule#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

