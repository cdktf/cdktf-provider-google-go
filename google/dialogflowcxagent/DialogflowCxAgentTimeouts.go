package dialogflowcxagent


type DialogflowCxAgentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_cx_agent#create DialogflowCxAgent#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_cx_agent#delete DialogflowCxAgent#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_cx_agent#update DialogflowCxAgent#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

