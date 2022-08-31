// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type DialogflowCxEnvironmentVersionConfigs struct {
	// Format: projects/{{project}}/locations/{{location}}/agents/{{agent}}/flows/{{flow}}/versions/{{version}}.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_cx_environment#version DialogflowCxEnvironment#version}
	Version *string `field:"required" json:"version" yaml:"version"`
}

