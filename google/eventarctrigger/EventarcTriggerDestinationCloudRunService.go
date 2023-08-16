package eventarctrigger


type EventarcTriggerDestinationCloudRunService struct {
	// Required.
	//
	// The name of the Cloud Run service being addressed. See https://cloud.google.com/run/docs/reference/rest/v1/namespaces.services. Only services located in the same project of the trigger object can be addressed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/eventarc_trigger#service EventarcTrigger#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// Optional.
	//
	// The relative path on the Cloud Run service the events should be sent to. The value must conform to the definition of URI path segment (section 3.3 of RFC2396). Examples: "/route", "route", "route/subroute".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/eventarc_trigger#path EventarcTrigger#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Required. The region the Cloud Run service is deployed in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/eventarc_trigger#region EventarcTrigger#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

