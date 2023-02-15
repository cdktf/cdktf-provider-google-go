package computeinstancetemplate


type ComputeInstanceTemplateConfidentialInstanceConfig struct {
	// Defines whether the instance should have confidential compute enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance_template#enable_confidential_compute ComputeInstanceTemplate#enable_confidential_compute}
	EnableConfidentialCompute interface{} `field:"required" json:"enableConfidentialCompute" yaml:"enableConfidentialCompute"`
}

