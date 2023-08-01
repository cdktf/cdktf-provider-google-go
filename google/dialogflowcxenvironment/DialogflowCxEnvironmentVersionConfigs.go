package dialogflowcxenvironment


type DialogflowCxEnvironmentVersionConfigs struct {
	// Format: projects/{{project}}/locations/{{location}}/agents/{{agent}}/flows/{{flow}}/versions/{{version}}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_cx_environment#version DialogflowCxEnvironment#version}
	Version *string `field:"required" json:"version" yaml:"version"`
}

