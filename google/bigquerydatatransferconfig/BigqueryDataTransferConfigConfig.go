package bigquerydatatransferconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigqueryDataTransferConfigConfig struct {
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
	// The data source id. Cannot be changed once the transfer config is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_data_transfer_config#data_source_id BigqueryDataTransferConfig#data_source_id}
	DataSourceId *string `field:"required" json:"dataSourceId" yaml:"dataSourceId"`
	// The user specified display name for the transfer config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_data_transfer_config#display_name BigqueryDataTransferConfig#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Parameters specific to each data source.
	//
	// For more information see the bq tab in the 'Setting up a data transfer'
	// section for each data source. For example the parameters for Cloud Storage transfers are listed here:
	// https://cloud.google.com/bigquery-transfer/docs/cloud-storage-transfer#bq
	//
	// *NOTE** : If you are attempting to update a parameter that cannot be updated (due to api limitations) [please force recreation of the resource](https://www.terraform.io/cli/state/taint#forcing-re-creation-of-resources).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_data_transfer_config#params BigqueryDataTransferConfig#params}
	Params *map[string]*string `field:"required" json:"params" yaml:"params"`
	// The number of days to look back to automatically refresh the data.
	//
	// For example, if dataRefreshWindowDays = 10, then every day BigQuery
	// reingests data for [today-10, today-1], rather than ingesting data for
	// just [today-1]. Only valid if the data source supports the feature.
	// Set the value to 0 to use the default value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_data_transfer_config#data_refresh_window_days BigqueryDataTransferConfig#data_refresh_window_days}
	DataRefreshWindowDays *float64 `field:"optional" json:"dataRefreshWindowDays" yaml:"dataRefreshWindowDays"`
	// The BigQuery target dataset id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_data_transfer_config#destination_dataset_id BigqueryDataTransferConfig#destination_dataset_id}
	DestinationDatasetId *string `field:"optional" json:"destinationDatasetId" yaml:"destinationDatasetId"`
	// When set to true, no runs are scheduled for a given transfer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_data_transfer_config#disabled BigqueryDataTransferConfig#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// email_preferences block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_data_transfer_config#email_preferences BigqueryDataTransferConfig#email_preferences}
	EmailPreferences *BigqueryDataTransferConfigEmailPreferences `field:"optional" json:"emailPreferences" yaml:"emailPreferences"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_data_transfer_config#id BigqueryDataTransferConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The geographic location where the transfer config should reside. Examples: US, EU, asia-northeast1. The default value is US.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_data_transfer_config#location BigqueryDataTransferConfig#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Pub/Sub topic where notifications will be sent after transfer runs associated with this transfer config finish.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_data_transfer_config#notification_pubsub_topic BigqueryDataTransferConfig#notification_pubsub_topic}
	NotificationPubsubTopic *string `field:"optional" json:"notificationPubsubTopic" yaml:"notificationPubsubTopic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_data_transfer_config#project BigqueryDataTransferConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Data transfer schedule.
	//
	// If the data source does not support a custom
	// schedule, this should be empty. If it is empty, the default value for
	// the data source will be used. The specified times are in UTC. Examples
	// of valid format: 1st,3rd monday of month 15:30, every wed,fri of jan,
	// jun 13:15, and first sunday of quarter 00:00. See more explanation
	// about the format here:
	// https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format
	// NOTE: the granularity should be at least 8 hours, or less frequent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_data_transfer_config#schedule BigqueryDataTransferConfig#schedule}
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// schedule_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_data_transfer_config#schedule_options BigqueryDataTransferConfig#schedule_options}
	ScheduleOptions *BigqueryDataTransferConfigScheduleOptions `field:"optional" json:"scheduleOptions" yaml:"scheduleOptions"`
	// sensitive_params block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_data_transfer_config#sensitive_params BigqueryDataTransferConfig#sensitive_params}
	SensitiveParams *BigqueryDataTransferConfigSensitiveParams `field:"optional" json:"sensitiveParams" yaml:"sensitiveParams"`
	// Service account email.
	//
	// If this field is set, transfer config will
	// be created with this service account credentials. It requires that
	// requesting user calling this API has permissions to act as this service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_data_transfer_config#service_account_name BigqueryDataTransferConfig#service_account_name}
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_data_transfer_config#timeouts BigqueryDataTransferConfig#timeouts}
	Timeouts *BigqueryDataTransferConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

