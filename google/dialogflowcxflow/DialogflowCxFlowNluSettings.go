package dialogflowcxflow


type DialogflowCxFlowNluSettings struct {
	// To filter out false positive results and still get variety in matched natural language inputs for your agent, you can tune the machine learning classification threshold.
	//
	// If the returned score value is less than the threshold value, then a no-match event will be triggered. The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the default of 0.3 is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_cx_flow#classification_threshold DialogflowCxFlow#classification_threshold}
	ClassificationThreshold *float64 `field:"optional" json:"classificationThreshold" yaml:"classificationThreshold"`
	// Indicates NLU model training mode.
	//
	// MODEL_TRAINING_MODE_AUTOMATIC: NLU model training is automatically triggered when a flow gets modified. User can also manually trigger model training in this mode.
	// MODEL_TRAINING_MODE_MANUAL: User needs to manually trigger NLU model training. Best for large flows whose models take long time to train. Possible values: ["MODEL_TRAINING_MODE_AUTOMATIC", "MODEL_TRAINING_MODE_MANUAL"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_cx_flow#model_training_mode DialogflowCxFlow#model_training_mode}
	ModelTrainingMode *string `field:"optional" json:"modelTrainingMode" yaml:"modelTrainingMode"`
	// Indicates the type of NLU model. MODEL_TYPE_STANDARD: Use standard NLU model. MODEL_TYPE_ADVANCED: Use advanced NLU model. Possible values: ["MODEL_TYPE_STANDARD", "MODEL_TYPE_ADVANCED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_cx_flow#model_type DialogflowCxFlow#model_type}
	ModelType *string `field:"optional" json:"modelType" yaml:"modelType"`
}

