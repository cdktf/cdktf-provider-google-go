package dataprocworkflowtemplate


type DataprocWorkflowTemplateJobsSparkRJobLoggingConfig struct {
	// The per-package log levels for the driver.
	//
	// This may include "root" package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.63.1/docs/resources/dataproc_workflow_template#driver_log_levels DataprocWorkflowTemplate#driver_log_levels}
	DriverLogLevels *map[string]*string `field:"optional" json:"driverLogLevels" yaml:"driverLogLevels"`
}

