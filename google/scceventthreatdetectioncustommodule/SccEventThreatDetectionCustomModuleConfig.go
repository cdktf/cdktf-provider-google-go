// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package scceventthreatdetectioncustommodule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SccEventThreatDetectionCustomModuleConfig struct {
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
	// Config for the module.
	//
	// For the resident module, its config value is defined at this level.
	// For the inherited module, its config value is inherited from the ancestor module.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/scc_event_threat_detection_custom_module#config SccEventThreatDetectionCustomModule#config}
	Config *string `field:"required" json:"config" yaml:"config"`
	// The state of enablement for the module at the given level of the hierarchy. Possible values: ["ENABLED", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/scc_event_threat_detection_custom_module#enablement_state SccEventThreatDetectionCustomModule#enablement_state}
	EnablementState *string `field:"required" json:"enablementState" yaml:"enablementState"`
	// Numerical ID of the parent organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/scc_event_threat_detection_custom_module#organization SccEventThreatDetectionCustomModule#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// Immutable. Type for the module. e.g. CONFIGURABLE_BAD_IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/scc_event_threat_detection_custom_module#type SccEventThreatDetectionCustomModule#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The human readable name to be displayed for the module.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/scc_event_threat_detection_custom_module#display_name SccEventThreatDetectionCustomModule#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/scc_event_threat_detection_custom_module#id SccEventThreatDetectionCustomModule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/scc_event_threat_detection_custom_module#timeouts SccEventThreatDetectionCustomModule#timeouts}
	Timeouts *SccEventThreatDetectionCustomModuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

