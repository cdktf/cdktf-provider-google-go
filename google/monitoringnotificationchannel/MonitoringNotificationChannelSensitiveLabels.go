package monitoringnotificationchannel


type MonitoringNotificationChannelSensitiveLabels struct {
	// An authorization token for a notification channel. Channel types that support this field include: slack.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_notification_channel#auth_token MonitoringNotificationChannel#auth_token}
	AuthToken *string `field:"optional" json:"authToken" yaml:"authToken"`
	// An password for a notification channel. Channel types that support this field include: webhook_basicauth.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_notification_channel#password MonitoringNotificationChannel#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// An servicekey token for a notification channel. Channel types that support this field include: pagerduty.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_notification_channel#service_key MonitoringNotificationChannel#service_key}
	ServiceKey *string `field:"optional" json:"serviceKey" yaml:"serviceKey"`
}

