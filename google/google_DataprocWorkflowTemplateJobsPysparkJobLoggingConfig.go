// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type DataprocWorkflowTemplateJobsPysparkJobLoggingConfig struct {
	// The per-package log levels for the driver.
	//
	// This may include "root" package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_workflow_template#driver_log_levels DataprocWorkflowTemplate#driver_log_levels}
	DriverLogLevels *map[string]*string `field:"optional" json:"driverLogLevels" yaml:"driverLogLevels"`
}

