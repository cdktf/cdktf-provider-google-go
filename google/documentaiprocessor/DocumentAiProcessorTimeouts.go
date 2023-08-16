package documentaiprocessor


type DocumentAiProcessorTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/document_ai_processor#create DocumentAiProcessor#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/document_ai_processor#delete DocumentAiProcessor#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

