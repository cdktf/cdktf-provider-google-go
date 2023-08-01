package dialogflowfulfillment


type DialogflowFulfillmentGenericWebService struct {
	// The fulfillment URI for receiving POST requests. It must use https protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_fulfillment#uri DialogflowFulfillment#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// The password for HTTP Basic authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_fulfillment#password DialogflowFulfillment#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The HTTP request headers to send together with fulfillment requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_fulfillment#request_headers DialogflowFulfillment#request_headers}
	RequestHeaders *map[string]*string `field:"optional" json:"requestHeaders" yaml:"requestHeaders"`
	// The user name for HTTP Basic authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_fulfillment#username DialogflowFulfillment#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

