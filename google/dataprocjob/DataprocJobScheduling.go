package dataprocjob


type DataprocJobScheduling struct {
	// Maximum number of times per hour a driver may be restarted as a result of driver exiting with non-zero code before job is reported failed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataproc_job#max_failures_per_hour DataprocJob#max_failures_per_hour}
	MaxFailuresPerHour *float64 `field:"required" json:"maxFailuresPerHour" yaml:"maxFailuresPerHour"`
	// Maximum number of times in total a driver may be restarted as a result of driver exiting with non-zero code before job is reported failed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataproc_job#max_failures_total DataprocJob#max_failures_total}
	MaxFailuresTotal *float64 `field:"required" json:"maxFailuresTotal" yaml:"maxFailuresTotal"`
}

