package beyondcorpappconnector


type BeyondcorpAppConnectorPrincipalInfoServiceAccount struct {
	// Email address of the service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/beyondcorp_app_connector#email BeyondcorpAppConnector#email}
	Email *string `field:"required" json:"email" yaml:"email"`
}

