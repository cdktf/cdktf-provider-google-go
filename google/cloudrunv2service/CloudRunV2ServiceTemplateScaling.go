package cloudrunv2service


type CloudRunV2ServiceTemplateScaling struct {
	// Maximum number of serving instances that this resource should have.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_v2_service#max_instance_count CloudRunV2Service#max_instance_count}
	MaxInstanceCount *float64 `field:"optional" json:"maxInstanceCount" yaml:"maxInstanceCount"`
	// Minimum number of serving instances that this resource should have.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_v2_service#min_instance_count CloudRunV2Service#min_instance_count}
	MinInstanceCount *float64 `field:"optional" json:"minInstanceCount" yaml:"minInstanceCount"`
}

