package dialogflowfulfillment


type DialogflowFulfillmentFeatures struct {
	// The type of the feature that enabled for fulfillment. SMALLTALK: Fulfillment is enabled for SmallTalk. Possible values: ["SMALLTALK"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_fulfillment#type DialogflowFulfillment#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

