package mlenginemodel


type MlEngineModelDefaultVersion struct {
	// The name specified for the version when it was created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/ml_engine_model#name MlEngineModel#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

