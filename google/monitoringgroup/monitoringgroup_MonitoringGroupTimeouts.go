package monitoringgroup


type MonitoringGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_group#create MonitoringGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_group#delete MonitoringGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_group#update MonitoringGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

