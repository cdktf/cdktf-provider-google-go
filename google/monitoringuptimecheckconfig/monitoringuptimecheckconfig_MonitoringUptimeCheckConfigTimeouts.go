package monitoringuptimecheckconfig


type MonitoringUptimeCheckConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_uptime_check_config#create MonitoringUptimeCheckConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_uptime_check_config#delete MonitoringUptimeCheckConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_uptime_check_config#update MonitoringUptimeCheckConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

