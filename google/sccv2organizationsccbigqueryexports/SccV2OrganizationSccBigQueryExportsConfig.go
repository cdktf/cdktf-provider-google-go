// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sccv2organizationsccbigqueryexports

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SccV2OrganizationSccBigQueryExportsConfig struct {
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
	// This must be unique within the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/scc_v2_organization_scc_big_query_exports#big_query_export_id SccV2OrganizationSccBigQueryExports#big_query_export_id}
	BigQueryExportId *string `field:"required" json:"bigQueryExportId" yaml:"bigQueryExportId"`
	// The organization whose Cloud Security Command Center the Big Query Export Config lives in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/scc_v2_organization_scc_big_query_exports#organization SccV2OrganizationSccBigQueryExports#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// The dataset to write findings' updates to.
	//
	// Its format is "projects/[projectId]/datasets/[bigquery_dataset_id]".
	// BigQuery Dataset unique ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/scc_v2_organization_scc_big_query_exports#dataset SccV2OrganizationSccBigQueryExports#dataset}
	Dataset *string `field:"optional" json:"dataset" yaml:"dataset"`
	// The description of the notification config (max of 1024 characters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/scc_v2_organization_scc_big_query_exports#description SccV2OrganizationSccBigQueryExports#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Expression that defines the filter to apply across create/update events of findings.
	//
	// The
	// expression is a list of zero or more restrictions combined via
	// logical operators AND and OR. Parentheses are supported, and OR
	// has higher precedence than AND.
	//
	// Restrictions have the form <field> <operator> <value> and may have
	// a - character in front of them to indicate negation. The fields
	// map to those defined in the corresponding resource.
	//
	// The supported operators are:
	//
	// * = for all value types.
	// * >, <, >=, <= for integer values.
	// * :, meaning substring matching, for strings.
	//
	// The supported value types are:
	//
	// * string literals in quotes.
	// * integer literals without quotes.
	// * boolean literals true and false without quotes.
	//
	// See
	// [Filtering notifications](https://cloud.google.com/security-command-center/docs/how-to-api-filter-notifications)
	// for information on how to write a filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/scc_v2_organization_scc_big_query_exports#filter SccV2OrganizationSccBigQueryExports#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/scc_v2_organization_scc_big_query_exports#id SccV2OrganizationSccBigQueryExports#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// location Id is provided by organization. If not provided, Use global as default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/scc_v2_organization_scc_big_query_exports#location SccV2OrganizationSccBigQueryExports#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The resource name of this export, in the format 'organizations/{{organization}}/locations/{{location}}/bigQueryExports/{{big_query_export_id}}'.
	//
	// This field is provided in responses, and is ignored when provided in create requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/scc_v2_organization_scc_big_query_exports#name SccV2OrganizationSccBigQueryExports#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/scc_v2_organization_scc_big_query_exports#timeouts SccV2OrganizationSccBigQueryExports#timeouts}
	Timeouts *SccV2OrganizationSccBigQueryExportsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

