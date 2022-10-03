package eventarctrigger


type EventarcTriggerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/eventarc_trigger#create EventarcTrigger#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/eventarc_trigger#delete EventarcTrigger#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/eventarc_trigger#update EventarcTrigger#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

