// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type MonitoringDashboardTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_dashboard#create MonitoringDashboard#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_dashboard#delete MonitoringDashboard#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_dashboard#update MonitoringDashboard#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

