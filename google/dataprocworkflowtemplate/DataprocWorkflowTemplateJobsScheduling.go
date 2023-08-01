package dataprocworkflowtemplate


type DataprocWorkflowTemplateJobsScheduling struct {
	// Optional.
	//
	// Maximum number of times per hour a driver may be restarted as a result of driver exiting with non-zero code before job is reported failed. A job may be reported as thrashing if driver exits with non-zero code 4 times within 10 minute window. Maximum value is 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataproc_workflow_template#max_failures_per_hour DataprocWorkflowTemplate#max_failures_per_hour}
	MaxFailuresPerHour *float64 `field:"optional" json:"maxFailuresPerHour" yaml:"maxFailuresPerHour"`
	// Optional.
	//
	// Maximum number of times in total a driver may be restarted as a result of driver exiting with non-zero code before job is reported failed. Maximum value is 240.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataproc_workflow_template#max_failures_total DataprocWorkflowTemplate#max_failures_total}
	MaxFailuresTotal *float64 `field:"optional" json:"maxFailuresTotal" yaml:"maxFailuresTotal"`
}

