package loggingfoldersink

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoggingFolderSinkConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/logging_folder_sink#destination LoggingFolderSink#destination}
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is accepted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/logging_folder_sink#folder LoggingFolderSink#folder}
	Folder *string `field:"required" json:"folder" yaml:"folder"`
	// The name of the logging sink.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/logging_folder_sink#name LoggingFolderSink#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// bigquery_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/logging_folder_sink#bigquery_options LoggingFolderSink#bigquery_options}
	BigqueryOptions *LoggingFolderSinkBigqueryOptions `field:"optional" json:"bigqueryOptions" yaml:"bigqueryOptions"`
	// A description of this sink. The maximum length of the description is 8000 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/logging_folder_sink#description LoggingFolderSink#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If set to True, then this sink is disabled and it does not export any log entries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/logging_folder_sink#disabled LoggingFolderSink#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// exclusions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/logging_folder_sink#exclusions LoggingFolderSink#exclusions}
	Exclusions interface{} `field:"optional" json:"exclusions" yaml:"exclusions"`
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/logging_folder_sink#filter LoggingFolderSink#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/logging_folder_sink#id LoggingFolderSink#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether or not to include children folders in the sink export.
	//
	// If true, logs associated with child projects are also exported; otherwise only logs relating to the provided folder are included.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/logging_folder_sink#include_children LoggingFolderSink#include_children}
	IncludeChildren interface{} `field:"optional" json:"includeChildren" yaml:"includeChildren"`
}

