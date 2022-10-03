package monitoringnotificationchannel


type MonitoringNotificationChannelTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_notification_channel#create MonitoringNotificationChannel#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_notification_channel#delete MonitoringNotificationChannel#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_notification_channel#update MonitoringNotificationChannel#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

