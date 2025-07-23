// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chronicleruledeployment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ChronicleRuleDeploymentConfig struct {
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
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_rule_deployment#instance ChronicleRuleDeployment#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// The location of the resource.
	//
	// This is the geographical region where the Chronicle instance resides, such as "us" or "europe-west2".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_rule_deployment#location ChronicleRuleDeployment#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The Rule ID of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_rule_deployment#rule ChronicleRuleDeployment#rule}
	Rule *string `field:"required" json:"rule" yaml:"rule"`
	// Whether detections resulting from this deployment should be considered alerts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_rule_deployment#alerting ChronicleRuleDeployment#alerting}
	Alerting interface{} `field:"optional" json:"alerting" yaml:"alerting"`
	// The archive state of the rule deployment.
	//
	// Cannot be set to true unless enabled is set to false i.e.
	// archiving requires a two-step process: first, disable the rule by
	// setting 'enabled' to false, then set 'archive' to true.
	// If set to true, alerting will automatically be set to false.
	// If currently set to true, enabled, alerting, and run_frequency cannot be
	// updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_rule_deployment#archived ChronicleRuleDeployment#archived}
	Archived interface{} `field:"optional" json:"archived" yaml:"archived"`
	// Whether the rule is currently deployed continuously against incoming data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_rule_deployment#enabled ChronicleRuleDeployment#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_rule_deployment#id ChronicleRuleDeployment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_rule_deployment#project ChronicleRuleDeployment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The run frequency of the rule deployment. Possible values: LIVE HOURLY DAILY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_rule_deployment#run_frequency ChronicleRuleDeployment#run_frequency}
	RunFrequency *string `field:"optional" json:"runFrequency" yaml:"runFrequency"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_rule_deployment#timeouts ChronicleRuleDeployment#timeouts}
	Timeouts *ChronicleRuleDeploymentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

