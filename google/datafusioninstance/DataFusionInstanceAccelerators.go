package datafusioninstance


type DataFusionInstanceAccelerators struct {
	// The type of an accelator for a CDF instance. Possible values: ["CDC", "HEALTHCARE", "CCAI_INSIGHTS"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_fusion_instance#accelerator_type DataFusionInstance#accelerator_type}
	AcceleratorType *string `field:"required" json:"acceleratorType" yaml:"acceleratorType"`
	// The type of an accelator for a CDF instance. Possible values: ["ENABLED", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_fusion_instance#state DataFusionInstance#state}
	State *string `field:"required" json:"state" yaml:"state"`
}

