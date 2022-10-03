package monitoringcustomservice


type MonitoringCustomServiceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_custom_service#create MonitoringCustomService#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_custom_service#delete MonitoringCustomService#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_custom_service#update MonitoringCustomService#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

