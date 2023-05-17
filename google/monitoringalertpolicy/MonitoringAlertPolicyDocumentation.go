package monitoringalertpolicy


type MonitoringAlertPolicyDocumentation struct {
	// The text of the documentation, interpreted according to mimeType.
	//
	// The content may not exceed 8,192 Unicode characters and may not
	// exceed more than 10,240 bytes when encoded in UTF-8 format,
	// whichever is smaller.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/monitoring_alert_policy#content MonitoringAlertPolicy#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// The format of the content field. Presently, only the value "text/markdown" is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/monitoring_alert_policy#mime_type MonitoringAlertPolicy#mime_type}
	MimeType *string `field:"optional" json:"mimeType" yaml:"mimeType"`
}

