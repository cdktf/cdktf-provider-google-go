// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sccv2foldersccbigqueryexport

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SccV2FolderSccBigQueryExportConfig struct {
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
	// It must consist of only lowercase letters,
	// numbers, and hyphens, must start with a letter, must end with either a letter or a number,
	// and must be 63 characters or less.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/scc_v2_folder_scc_big_query_export#big_query_export_id SccV2FolderSccBigQueryExport#big_query_export_id}
	BigQueryExportId *string `field:"required" json:"bigQueryExportId" yaml:"bigQueryExportId"`
	// The folder where Cloud Security Command Center Big Query Export Config lives in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/scc_v2_folder_scc_big_query_export#folder SccV2FolderSccBigQueryExport#folder}
	Folder *string `field:"required" json:"folder" yaml:"folder"`
	// The dataset to write findings' updates to.
	//
	// Its format is "projects/[projectId]/datasets/[bigquery_dataset_id]".
	// BigQuery Dataset unique ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/scc_v2_folder_scc_big_query_export#dataset SccV2FolderSccBigQueryExport#dataset}
	Dataset *string `field:"optional" json:"dataset" yaml:"dataset"`
	// The description of the notification config (max of 1024 characters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/scc_v2_folder_scc_big_query_export#description SccV2FolderSccBigQueryExport#description}
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/scc_v2_folder_scc_big_query_export#filter SccV2FolderSccBigQueryExport#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/scc_v2_folder_scc_big_query_export#id SccV2FolderSccBigQueryExport#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The BigQuery export configuration is stored in this location. If not provided, Use global as default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/scc_v2_folder_scc_big_query_export#location SccV2FolderSccBigQueryExport#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/scc_v2_folder_scc_big_query_export#timeouts SccV2FolderSccBigQueryExport#timeouts}
	Timeouts *SccV2FolderSccBigQueryExportTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

