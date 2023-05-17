package monitoringalertpolicy


type MonitoringAlertPolicyAlertStrategy struct {
	// If an alert policy that was active has no data for this long, any open incidents will close.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/monitoring_alert_policy#auto_close MonitoringAlertPolicy#auto_close}
	AutoClose *string `field:"optional" json:"autoClose" yaml:"autoClose"`
	// notification_rate_limit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/monitoring_alert_policy#notification_rate_limit MonitoringAlertPolicy#notification_rate_limit}
	NotificationRateLimit *MonitoringAlertPolicyAlertStrategyNotificationRateLimit `field:"optional" json:"notificationRateLimit" yaml:"notificationRateLimit"`
}

