package cloudfunctionsfunction


type CloudfunctionsFunctionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions_function#create CloudfunctionsFunction#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions_function#delete CloudfunctionsFunction#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions_function#read CloudfunctionsFunction#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions_function#update CloudfunctionsFunction#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

