// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigqueryconnection


type BigqueryConnectionCloudSpanner struct {
	// Cloud Spanner database in the form 'project/instance/database'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/bigquery_connection#database BigqueryConnection#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Cloud Spanner database role for fine-grained access control.
	//
	// The Cloud Spanner admin should have provisioned the database role with appropriate permissions, such as 'SELECT' and 'INSERT'. Other users should only use roles provided by their Cloud Spanner admins. The database role name must start with a letter, and can only contain letters, numbers, and underscores. For more details, see https://cloud.google.com/spanner/docs/fgac-about.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/bigquery_connection#database_role BigqueryConnection#database_role}
	DatabaseRole *string `field:"optional" json:"databaseRole" yaml:"databaseRole"`
	// Allows setting max parallelism per query when executing on Spanner independent compute resources.
	//
	// If unspecified, default values of parallelism are chosen that are dependent on the Cloud Spanner instance configuration. 'useParallelism' and 'useDataBoost' must be set when setting max parallelism.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/bigquery_connection#max_parallelism BigqueryConnection#max_parallelism}
	MaxParallelism *float64 `field:"optional" json:"maxParallelism" yaml:"maxParallelism"`
	// If set, the request will be executed via Spanner independent compute resources.
	//
	// 'use_parallelism' must be set when using data boost.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/bigquery_connection#use_data_boost BigqueryConnection#use_data_boost}
	UseDataBoost interface{} `field:"optional" json:"useDataBoost" yaml:"useDataBoost"`
	// If parallelism should be used when reading from Cloud Spanner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/bigquery_connection#use_parallelism BigqueryConnection#use_parallelism}
	UseParallelism interface{} `field:"optional" json:"useParallelism" yaml:"useParallelism"`
	// If the serverless analytics service should be used to read data from Cloud Spanner.
	//
	// 'useParallelism' must be set when using serverless analytics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/bigquery_connection#use_serverless_analytics BigqueryConnection#use_serverless_analytics}
	UseServerlessAnalytics interface{} `field:"optional" json:"useServerlessAnalytics" yaml:"useServerlessAnalytics"`
}

