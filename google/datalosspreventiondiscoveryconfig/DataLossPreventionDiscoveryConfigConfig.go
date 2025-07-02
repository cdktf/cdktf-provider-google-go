// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataLossPreventionDiscoveryConfigConfig struct {
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
	// Location to create the discovery config in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/data_loss_prevention_discovery_config#location DataLossPreventionDiscoveryConfig#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The parent of the discovery config in any of the following formats:.
	//
	// * 'projects/{{project}}/locations/{{location}}'
	// * 'organizations/{{organization_id}}/locations/{{location}}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/data_loss_prevention_discovery_config#parent DataLossPreventionDiscoveryConfig#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/data_loss_prevention_discovery_config#actions DataLossPreventionDiscoveryConfig#actions}
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// Display Name (max 1000 Chars).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/data_loss_prevention_discovery_config#display_name DataLossPreventionDiscoveryConfig#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/data_loss_prevention_discovery_config#id DataLossPreventionDiscoveryConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Detection logic for profile generation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/data_loss_prevention_discovery_config#inspect_templates DataLossPreventionDiscoveryConfig#inspect_templates}
	InspectTemplates *[]*string `field:"optional" json:"inspectTemplates" yaml:"inspectTemplates"`
	// org_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/data_loss_prevention_discovery_config#org_config DataLossPreventionDiscoveryConfig#org_config}
	OrgConfig *DataLossPreventionDiscoveryConfigOrgConfig `field:"optional" json:"orgConfig" yaml:"orgConfig"`
	// Required. A status for this configuration Possible values: ["RUNNING", "PAUSED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/data_loss_prevention_discovery_config#status DataLossPreventionDiscoveryConfig#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/data_loss_prevention_discovery_config#targets DataLossPreventionDiscoveryConfig#targets}
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/data_loss_prevention_discovery_config#timeouts DataLossPreventionDiscoveryConfig#timeouts}
	Timeouts *DataLossPreventionDiscoveryConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

