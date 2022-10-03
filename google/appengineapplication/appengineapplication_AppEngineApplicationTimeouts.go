package appengineapplication


type AppEngineApplicationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_application#create AppEngineApplication#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_application#update AppEngineApplication#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

