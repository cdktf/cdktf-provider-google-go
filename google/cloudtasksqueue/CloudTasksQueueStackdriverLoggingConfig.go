package cloudtasksqueue


type CloudTasksQueueStackdriverLoggingConfig struct {
	// Specifies the fraction of operations to write to Stackdriver Logging.
	//
	// This field may contain any value between 0.0 and 1.0, inclusive. 0.0 is the
	// default and means that no operations are logged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_tasks_queue#sampling_ratio CloudTasksQueue#sampling_ratio}
	SamplingRatio *float64 `field:"required" json:"samplingRatio" yaml:"samplingRatio"`
}

