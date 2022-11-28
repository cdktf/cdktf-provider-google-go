package beyondcorpappconnector


type BeyondcorpAppConnectorPrincipalInfo struct {
	// service_account block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/beyondcorp_app_connector#service_account BeyondcorpAppConnector#service_account}
	ServiceAccount *BeyondcorpAppConnectorPrincipalInfoServiceAccount `field:"required" json:"serviceAccount" yaml:"serviceAccount"`
}

