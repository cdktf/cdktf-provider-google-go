package dialogflowcxwebhook


type DialogflowCxWebhookServiceDirectoryGenericWebService struct {
	// Whether to use speech adaptation for speech recognition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dialogflow_cx_webhook#uri DialogflowCxWebhook#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// Specifies a list of allowed custom CA certificates (in DER format) for HTTPS verification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dialogflow_cx_webhook#allowed_ca_certs DialogflowCxWebhook#allowed_ca_certs}
	AllowedCaCerts *[]*string `field:"optional" json:"allowedCaCerts" yaml:"allowedCaCerts"`
	// The HTTP request headers to send together with webhook requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dialogflow_cx_webhook#request_headers DialogflowCxWebhook#request_headers}
	RequestHeaders *map[string]*string `field:"optional" json:"requestHeaders" yaml:"requestHeaders"`
}

