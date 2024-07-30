// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigqueryanalyticshubdataexchangeiambinding

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigqueryAnalyticsHubDataExchangeIamBindingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/bigquery_analytics_hub_data_exchange_iam_binding#data_exchange_id BigqueryAnalyticsHubDataExchangeIamBinding#data_exchange_id}.
	DataExchangeId *string `field:"required" json:"dataExchangeId" yaml:"dataExchangeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/bigquery_analytics_hub_data_exchange_iam_binding#members BigqueryAnalyticsHubDataExchangeIamBinding#members}.
	Members *[]*string `field:"required" json:"members" yaml:"members"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/bigquery_analytics_hub_data_exchange_iam_binding#role BigqueryAnalyticsHubDataExchangeIamBinding#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/bigquery_analytics_hub_data_exchange_iam_binding#condition BigqueryAnalyticsHubDataExchangeIamBinding#condition}
	Condition *BigqueryAnalyticsHubDataExchangeIamBindingCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/bigquery_analytics_hub_data_exchange_iam_binding#id BigqueryAnalyticsHubDataExchangeIamBinding#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/bigquery_analytics_hub_data_exchange_iam_binding#location BigqueryAnalyticsHubDataExchangeIamBinding#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/bigquery_analytics_hub_data_exchange_iam_binding#project BigqueryAnalyticsHubDataExchangeIamBinding#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

