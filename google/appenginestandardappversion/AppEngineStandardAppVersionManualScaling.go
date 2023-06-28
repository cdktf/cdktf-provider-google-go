package appenginestandardappversion


type AppEngineStandardAppVersionManualScaling struct {
	// Number of instances to assign to the service at the start.
	//
	// *Note:** When managing the number of instances at runtime through the App Engine Admin API or the (now deprecated) Python 2
	// Modules API set_num_instances() you must use 'lifecycle.ignore_changes = ["manual_scaling"[0].instances]' to prevent drift detection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.71.0/docs/resources/app_engine_standard_app_version#instances AppEngineStandardAppVersion#instances}
	Instances *float64 `field:"required" json:"instances" yaml:"instances"`
}

