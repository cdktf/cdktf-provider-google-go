// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigqueryanalyticshublisting


type BigqueryAnalyticsHubListingPublisher struct {
	// Name of the listing publisher.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/bigquery_analytics_hub_listing#name BigqueryAnalyticsHubListing#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Email or URL of the listing publisher.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/bigquery_analytics_hub_listing#primary_contact BigqueryAnalyticsHubListing#primary_contact}
	PrimaryContact *string `field:"optional" json:"primaryContact" yaml:"primaryContact"`
}

