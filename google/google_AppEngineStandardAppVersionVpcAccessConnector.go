// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type AppEngineStandardAppVersionVpcAccessConnector struct {
	// Full Serverless VPC Access Connector name e.g. /projects/my-project/locations/us-central1/connectors/c1.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_standard_app_version#name AppEngineStandardAppVersion#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

