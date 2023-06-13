package beyondcorpappconnection


type BeyondcorpAppConnectionGateway struct {
	// AppGateway name in following format: projects/{project_id}/locations/{locationId}/appgateways/{gateway_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/beyondcorp_app_connection#app_gateway BeyondcorpAppConnection#app_gateway}
	AppGateway *string `field:"required" json:"appGateway" yaml:"appGateway"`
	// The type of hosting used by the gateway. Refer to https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#Type_1 for a list of possible values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/beyondcorp_app_connection#type BeyondcorpAppConnection#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

