package computeinstancetemplate


type ComputeInstanceTemplateServiceAccount struct {
	// A list of service scopes.
	//
	// Both OAuth2 URLs and gcloud short names are supported. To allow full access to all Cloud APIs, use the cloud-platform scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_template#scopes ComputeInstanceTemplate#scopes}
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// The service account e-mail address. If not given, the default Google Compute Engine service account is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_template#email ComputeInstanceTemplate#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
}

