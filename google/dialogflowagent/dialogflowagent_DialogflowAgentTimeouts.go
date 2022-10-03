package dialogflowagent


type DialogflowAgentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_agent#create DialogflowAgent#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_agent#delete DialogflowAgent#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_agent#update DialogflowAgent#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

