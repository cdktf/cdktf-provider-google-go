package mlenginemodel


type MlEngineModelTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/ml_engine_model#create MlEngineModel#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/ml_engine_model#delete MlEngineModel#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

