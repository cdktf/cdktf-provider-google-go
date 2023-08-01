package eventarctrigger


type EventarcTriggerDestinationGke struct {
	// Required.
	//
	// The name of the cluster the GKE service is running in. The cluster must be running in the same project as the trigger being created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/eventarc_trigger#cluster EventarcTrigger#cluster}
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// Required.
	//
	// The name of the Google Compute Engine in which the cluster resides, which can either be compute zone (for example, us-central1-a) for the zonal clusters or region (for example, us-central1) for regional clusters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/eventarc_trigger#location EventarcTrigger#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Required. The namespace the GKE service is running in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/eventarc_trigger#namespace EventarcTrigger#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Required. Name of the GKE service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/eventarc_trigger#service EventarcTrigger#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// Optional.
	//
	// The relative path on the GKE service the events should be sent to. The value must conform to the definition of a URI path segment (section 3.3 of RFC2396). Examples: "/route", "route", "route/subroute".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/eventarc_trigger#path EventarcTrigger#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

