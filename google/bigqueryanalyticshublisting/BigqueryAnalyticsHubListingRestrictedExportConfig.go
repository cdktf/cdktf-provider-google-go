// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigqueryanalyticshublisting


type BigqueryAnalyticsHubListingRestrictedExportConfig struct {
	// If true, enable restricted export.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/bigquery_analytics_hub_listing#enabled BigqueryAnalyticsHubListing#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// If true, restrict export of query result derived from restricted linked dataset table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/bigquery_analytics_hub_listing#restrict_query_result BigqueryAnalyticsHubListing#restrict_query_result}
	RestrictQueryResult interface{} `field:"optional" json:"restrictQueryResult" yaml:"restrictQueryResult"`
}

