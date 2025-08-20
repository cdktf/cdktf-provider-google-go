// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package contactcenterinsightsanalysisrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContactCenterInsightsAnalysisRuleConfig struct {
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
	// Location of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/contact_center_insights_analysis_rule#location ContactCenterInsightsAnalysisRule#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// If true, apply this rule to conversations. Otherwise, this rule is inactive and saved as a draft.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/contact_center_insights_analysis_rule#active ContactCenterInsightsAnalysisRule#active}
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// Percentage of conversations that we should apply this analysis setting automatically, between [0, 1].
	//
	// For example, 0.1 means 10%. Conversations
	// are sampled in a determenestic way. The original runtime_percentage &
	// upload percentage will be replaced by defining filters on the conversation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/contact_center_insights_analysis_rule#analysis_percentage ContactCenterInsightsAnalysisRule#analysis_percentage}
	AnalysisPercentage *float64 `field:"optional" json:"analysisPercentage" yaml:"analysisPercentage"`
	// annotator_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/contact_center_insights_analysis_rule#annotator_selector ContactCenterInsightsAnalysisRule#annotator_selector}
	AnnotatorSelector *ContactCenterInsightsAnalysisRuleAnnotatorSelector `field:"optional" json:"annotatorSelector" yaml:"annotatorSelector"`
	// Filter for the conversations that should apply this analysis rule.
	//
	// An empty filter means this analysis rule applies to all
	// conversations.
	// Refer to https://cloud.google.com/contact-center/insights/docs/filtering
	// for details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/contact_center_insights_analysis_rule#conversation_filter ContactCenterInsightsAnalysisRule#conversation_filter}
	ConversationFilter *string `field:"optional" json:"conversationFilter" yaml:"conversationFilter"`
	// Display Name of the analysis rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/contact_center_insights_analysis_rule#display_name ContactCenterInsightsAnalysisRule#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/contact_center_insights_analysis_rule#id ContactCenterInsightsAnalysisRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/contact_center_insights_analysis_rule#project ContactCenterInsightsAnalysisRule#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/contact_center_insights_analysis_rule#timeouts ContactCenterInsightsAnalysisRule#timeouts}
	Timeouts *ContactCenterInsightsAnalysisRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

