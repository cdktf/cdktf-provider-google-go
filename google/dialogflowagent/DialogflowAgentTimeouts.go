package dialogflowagent


type DialogflowAgentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_agent#create DialogflowAgent#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_agent#delete DialogflowAgent#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_agent#update DialogflowAgent#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

