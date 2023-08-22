package beyondcorpappconnector


type BeyondcorpAppConnectorPrincipalInfo struct {
	// service_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/beyondcorp_app_connector#service_account BeyondcorpAppConnector#service_account}
	ServiceAccount *BeyondcorpAppConnectorPrincipalInfoServiceAccount `field:"required" json:"serviceAccount" yaml:"serviceAccount"`
}

