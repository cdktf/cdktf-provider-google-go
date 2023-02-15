package beyondcorpappconnector


type BeyondcorpAppConnectorPrincipalInfoServiceAccount struct {
	// Email address of the service account.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/beyondcorp_app_connector#email BeyondcorpAppConnector#email}
	Email *string `field:"required" json:"email" yaml:"email"`
}

