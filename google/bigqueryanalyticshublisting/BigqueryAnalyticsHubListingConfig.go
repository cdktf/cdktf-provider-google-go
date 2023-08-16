package bigqueryanalyticshublisting

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigqueryAnalyticsHubListingConfig struct {
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
	// bigquery_dataset block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_analytics_hub_listing#bigquery_dataset BigqueryAnalyticsHubListing#bigquery_dataset}
	BigqueryDataset *BigqueryAnalyticsHubListingBigqueryDataset `field:"required" json:"bigqueryDataset" yaml:"bigqueryDataset"`
	// The ID of the data exchange.
	//
	// Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_analytics_hub_listing#data_exchange_id BigqueryAnalyticsHubListing#data_exchange_id}
	DataExchangeId *string `field:"required" json:"dataExchangeId" yaml:"dataExchangeId"`
	// Human-readable display name of the listing.
	//
	// The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), ampersands (&) and can't start or end with spaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_analytics_hub_listing#display_name BigqueryAnalyticsHubListing#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The ID of the listing.
	//
	// Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_analytics_hub_listing#listing_id BigqueryAnalyticsHubListing#listing_id}
	ListingId *string `field:"required" json:"listingId" yaml:"listingId"`
	// The name of the location this data exchange listing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_analytics_hub_listing#location BigqueryAnalyticsHubListing#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Categories of the listing. Up to two categories are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_analytics_hub_listing#categories BigqueryAnalyticsHubListing#categories}
	Categories *[]*string `field:"optional" json:"categories" yaml:"categories"`
	// data_provider block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_analytics_hub_listing#data_provider BigqueryAnalyticsHubListing#data_provider}
	DataProvider *BigqueryAnalyticsHubListingDataProvider `field:"optional" json:"dataProvider" yaml:"dataProvider"`
	// Short description of the listing.
	//
	// The description must not contain Unicode non-characters and C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_analytics_hub_listing#description BigqueryAnalyticsHubListing#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Documentation describing the listing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_analytics_hub_listing#documentation BigqueryAnalyticsHubListing#documentation}
	Documentation *string `field:"optional" json:"documentation" yaml:"documentation"`
	// Base64 encoded image representing the listing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_analytics_hub_listing#icon BigqueryAnalyticsHubListing#icon}
	Icon *string `field:"optional" json:"icon" yaml:"icon"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_analytics_hub_listing#id BigqueryAnalyticsHubListing#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Email or URL of the primary point of contact of the listing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_analytics_hub_listing#primary_contact BigqueryAnalyticsHubListing#primary_contact}
	PrimaryContact *string `field:"optional" json:"primaryContact" yaml:"primaryContact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_analytics_hub_listing#project BigqueryAnalyticsHubListing#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// publisher block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_analytics_hub_listing#publisher BigqueryAnalyticsHubListing#publisher}
	Publisher *BigqueryAnalyticsHubListingPublisher `field:"optional" json:"publisher" yaml:"publisher"`
	// Email or URL of the request access of the listing. Subscribers can use this reference to request access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_analytics_hub_listing#request_access BigqueryAnalyticsHubListing#request_access}
	RequestAccess *string `field:"optional" json:"requestAccess" yaml:"requestAccess"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_analytics_hub_listing#timeouts BigqueryAnalyticsHubListing#timeouts}
	Timeouts *BigqueryAnalyticsHubListingTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

