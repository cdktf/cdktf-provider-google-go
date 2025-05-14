// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package billingbudget


type BillingBudgetAllUpdatesRule struct {
	// Boolean.
	//
	// When set to true, disables default notifications sent
	// when a threshold is exceeded. Default recipients are
	// those with Billing Account Administrators and Billing
	// Account Users IAM roles for the target account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/billing_budget#disable_default_iam_recipients BillingBudget#disable_default_iam_recipients}
	DisableDefaultIamRecipients interface{} `field:"optional" json:"disableDefaultIamRecipients" yaml:"disableDefaultIamRecipients"`
	// When set to true, and when the budget has a single project configured, notifications will be sent to project level recipients of that project.
	//
	// This field will be ignored if the budget has multiple or no project configured.
	//
	// Currently, project level recipients are the users with Owner role on a cloud project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/billing_budget#enable_project_level_recipients BillingBudget#enable_project_level_recipients}
	EnableProjectLevelRecipients interface{} `field:"optional" json:"enableProjectLevelRecipients" yaml:"enableProjectLevelRecipients"`
	// The full resource name of a monitoring notification channel in the form projects/{project_id}/notificationChannels/{channel_id}. A maximum of 5 channels are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/billing_budget#monitoring_notification_channels BillingBudget#monitoring_notification_channels}
	MonitoringNotificationChannels *[]*string `field:"optional" json:"monitoringNotificationChannels" yaml:"monitoringNotificationChannels"`
	// The name of the Cloud Pub/Sub topic where budget related messages will be published, in the form projects/{project_id}/topics/{topic_id}.
	//
	// Updates are sent
	// at regular intervals to the topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/billing_budget#pubsub_topic BillingBudget#pubsub_topic}
	PubsubTopic *string `field:"optional" json:"pubsubTopic" yaml:"pubsubTopic"`
	// The schema version of the notification. Only "1.0" is accepted. It represents the JSON schema as defined in https://cloud.google.com/billing/docs/how-to/budgets#notification_format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/billing_budget#schema_version BillingBudget#schema_version}
	SchemaVersion *string `field:"optional" json:"schemaVersion" yaml:"schemaVersion"`
}

