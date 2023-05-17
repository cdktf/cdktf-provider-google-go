package datagooglemonitoringnotificationchannel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleMonitoringNotificationChannelConfig struct {
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
	// An optional human-readable name for this notification channel.
	//
	// It is recommended that you specify a non-empty and unique name in order to make it easier to identify the channels in your project, though this is not enforced. The display name is limited to 512 Unicode characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/data-sources/monitoring_notification_channel#display_name DataGoogleMonitoringNotificationChannel#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/data-sources/monitoring_notification_channel#id DataGoogleMonitoringNotificationChannel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Configuration fields that define the channel and its behavior.
	//
	// The
	// permissible and required labels are specified in the
	// NotificationChannelDescriptor corresponding to the type field.
	//
	// Labels with sensitive data are obfuscated by the API and therefore Terraform cannot
	// determine if there are upstream changes to these fields. They can also be configured via
	// the sensitive_labels block, but cannot be configured in both places.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/data-sources/monitoring_notification_channel#labels DataGoogleMonitoringNotificationChannel#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/data-sources/monitoring_notification_channel#project DataGoogleMonitoringNotificationChannel#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The type of the notification channel.
	//
	// This field matches the value of the NotificationChannelDescriptor.type field. See https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of valid values such as "email", "slack", etc...
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/data-sources/monitoring_notification_channel#type DataGoogleMonitoringNotificationChannel#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field.
	//
	// This field is intended to be used for organizing and identifying the NotificationChannel objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/data-sources/monitoring_notification_channel#user_labels DataGoogleMonitoringNotificationChannel#user_labels}
	UserLabels *map[string]*string `field:"optional" json:"userLabels" yaml:"userLabels"`
}

