package loggingorganizationsink

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoggingOrganizationSinkConfig struct {
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
	// The destination of the sink (or, in other words, where logs are written to).
	//
	// Can be a Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples: "storage.googleapis.com/[GCS_BUCKET]" "bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]" "pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]" The writer associated with the sink must have access to write to the above resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/logging_organization_sink#destination LoggingOrganizationSink#destination}
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// The name of the logging sink.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/logging_organization_sink#name LoggingOrganizationSink#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The numeric ID of the organization to be exported to the sink.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/logging_organization_sink#org_id LoggingOrganizationSink#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// bigquery_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/logging_organization_sink#bigquery_options LoggingOrganizationSink#bigquery_options}
	BigqueryOptions *LoggingOrganizationSinkBigqueryOptions `field:"optional" json:"bigqueryOptions" yaml:"bigqueryOptions"`
	// A description of this sink. The maximum length of the description is 8000 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/logging_organization_sink#description LoggingOrganizationSink#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If set to True, then this sink is disabled and it does not export any log entries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/logging_organization_sink#disabled LoggingOrganizationSink#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// exclusions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/logging_organization_sink#exclusions LoggingOrganizationSink#exclusions}
	Exclusions interface{} `field:"optional" json:"exclusions" yaml:"exclusions"`
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/logging_organization_sink#filter LoggingOrganizationSink#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/logging_organization_sink#id LoggingOrganizationSink#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether or not to include children organizations in the sink export.
	//
	// If true, logs associated with child projects are also exported; otherwise only logs relating to the provided organization are included.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/logging_organization_sink#include_children LoggingOrganizationSink#include_children}
	IncludeChildren interface{} `field:"optional" json:"includeChildren" yaml:"includeChildren"`
}

