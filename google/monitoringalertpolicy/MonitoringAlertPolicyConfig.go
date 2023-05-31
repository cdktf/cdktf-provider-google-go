package monitoringalertpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitoringAlertPolicyConfig struct {
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
	// How to combine the results of multiple conditions to determine if an incident should be opened.
	//
	// Possible values: ["AND", "OR", "AND_WITH_MATCHING_RESOURCE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.67.0/docs/resources/monitoring_alert_policy#combiner MonitoringAlertPolicy#combiner}
	Combiner *string `field:"required" json:"combiner" yaml:"combiner"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.67.0/docs/resources/monitoring_alert_policy#conditions MonitoringAlertPolicy#conditions}
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// A short name or phrase used to identify the policy in dashboards, notifications, and incidents.
	//
	// To avoid confusion, don't use
	// the same display name for multiple policies in the same project. The
	// name is limited to 512 Unicode characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.67.0/docs/resources/monitoring_alert_policy#display_name MonitoringAlertPolicy#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// alert_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.67.0/docs/resources/monitoring_alert_policy#alert_strategy MonitoringAlertPolicy#alert_strategy}
	AlertStrategy *MonitoringAlertPolicyAlertStrategy `field:"optional" json:"alertStrategy" yaml:"alertStrategy"`
	// documentation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.67.0/docs/resources/monitoring_alert_policy#documentation MonitoringAlertPolicy#documentation}
	Documentation *MonitoringAlertPolicyDocumentation `field:"optional" json:"documentation" yaml:"documentation"`
	// Whether or not the policy is enabled. The default is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.67.0/docs/resources/monitoring_alert_policy#enabled MonitoringAlertPolicy#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.67.0/docs/resources/monitoring_alert_policy#id MonitoringAlertPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Identifies the notification channels to which notifications should be sent when incidents are opened or closed or when new violations occur on an already opened incident.
	//
	// Each element of this array corresponds
	// to the name field in each of the NotificationChannel objects that are
	// returned from the notificationChannels.list method. The syntax of the
	// entries in this field is
	// 'projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.67.0/docs/resources/monitoring_alert_policy#notification_channels MonitoringAlertPolicy#notification_channels}
	NotificationChannels *[]*string `field:"optional" json:"notificationChannels" yaml:"notificationChannels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.67.0/docs/resources/monitoring_alert_policy#project MonitoringAlertPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.67.0/docs/resources/monitoring_alert_policy#timeouts MonitoringAlertPolicy#timeouts}
	Timeouts *MonitoringAlertPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// This field is intended to be used for organizing and identifying the AlertPolicy objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.67.0/docs/resources/monitoring_alert_policy#user_labels MonitoringAlertPolicy#user_labels}
	UserLabels *map[string]*string `field:"optional" json:"userLabels" yaml:"userLabels"`
}

