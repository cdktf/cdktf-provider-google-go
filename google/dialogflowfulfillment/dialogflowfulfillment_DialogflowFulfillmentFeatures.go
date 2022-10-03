package dialogflowfulfillment


type DialogflowFulfillmentFeatures struct {
	// The type of the feature that enabled for fulfillment. SMALLTALK: Fulfillment is enabled for SmallTalk. Possible values: ["SMALLTALK"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_fulfillment#type DialogflowFulfillment#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

