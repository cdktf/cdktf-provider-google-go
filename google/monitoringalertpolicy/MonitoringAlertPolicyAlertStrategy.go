// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoringalertpolicy


type MonitoringAlertPolicyAlertStrategy struct {
	// If an alert policy that was active has no data for this long, any open incidents will close.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/monitoring_alert_policy#auto_close MonitoringAlertPolicy#auto_close}
	AutoClose *string `field:"optional" json:"autoClose" yaml:"autoClose"`
	// notification_channel_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/monitoring_alert_policy#notification_channel_strategy MonitoringAlertPolicy#notification_channel_strategy}
	NotificationChannelStrategy interface{} `field:"optional" json:"notificationChannelStrategy" yaml:"notificationChannelStrategy"`
	// Control when notifications will be sent out. Possible values: ["NOTIFICATION_PROMPT_UNSPECIFIED", "OPENED", "CLOSED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/monitoring_alert_policy#notification_prompts MonitoringAlertPolicy#notification_prompts}
	NotificationPrompts *[]*string `field:"optional" json:"notificationPrompts" yaml:"notificationPrompts"`
	// notification_rate_limit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/monitoring_alert_policy#notification_rate_limit MonitoringAlertPolicy#notification_rate_limit}
	NotificationRateLimit *MonitoringAlertPolicyAlertStrategyNotificationRateLimit `field:"optional" json:"notificationRateLimit" yaml:"notificationRateLimit"`
}

