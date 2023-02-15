package dialogflowcxintent


type DialogflowCxIntentTrainingPhrases struct {
	// parts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_cx_intent#parts DialogflowCxIntent#parts}
	Parts interface{} `field:"required" json:"parts" yaml:"parts"`
	// Indicates how many times this example was added to the intent.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_cx_intent#repeat_count DialogflowCxIntent#repeat_count}
	RepeatCount *float64 `field:"optional" json:"repeatCount" yaml:"repeatCount"`
}

