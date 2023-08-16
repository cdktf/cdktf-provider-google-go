package dialogflowcxintent


type DialogflowCxIntentTrainingPhrasesParts struct {
	// The text for this part.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_cx_intent#text DialogflowCxIntent#text}
	Text *string `field:"required" json:"text" yaml:"text"`
	// The parameter used to annotate this part of the training phrase.
	//
	// This field is required for annotated parts of the training phrase.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_cx_intent#parameter_id DialogflowCxIntent#parameter_id}
	ParameterId *string `field:"optional" json:"parameterId" yaml:"parameterId"`
}

