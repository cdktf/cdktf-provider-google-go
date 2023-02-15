package clouddeploytarget


type ClouddeployTargetTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/clouddeploy_target#create ClouddeployTarget#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/clouddeploy_target#delete ClouddeployTarget#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/clouddeploy_target#update ClouddeployTarget#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

