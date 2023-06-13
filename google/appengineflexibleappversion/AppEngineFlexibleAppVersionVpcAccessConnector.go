package appengineflexibleappversion


type AppEngineFlexibleAppVersionVpcAccessConnector struct {
	// Full Serverless VPC Access Connector name e.g. /projects/my-project/locations/us-central1/connectors/c1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/app_engine_flexible_app_version#name AppEngineFlexibleAppVersion#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

