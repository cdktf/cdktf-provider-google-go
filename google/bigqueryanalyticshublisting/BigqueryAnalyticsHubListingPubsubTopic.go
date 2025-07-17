// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigqueryanalyticshublisting


type BigqueryAnalyticsHubListingPubsubTopic struct {
	// Resource name of the Pub/Sub topic source for this listing. e.g. projects/myproject/topics/topicId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/bigquery_analytics_hub_listing#topic BigqueryAnalyticsHubListing#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// Region hint on where the data might be published.
	//
	// Data affinity regions are modifiable.
	// See https://cloud.google.com/about/locations for full listing of possible Cloud regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/bigquery_analytics_hub_listing#data_affinity_regions BigqueryAnalyticsHubListing#data_affinity_regions}
	DataAffinityRegions *[]*string `field:"optional" json:"dataAffinityRegions" yaml:"dataAffinityRegions"`
}

