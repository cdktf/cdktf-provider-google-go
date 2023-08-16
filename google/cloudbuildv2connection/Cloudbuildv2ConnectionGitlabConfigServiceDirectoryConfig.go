package cloudbuildv2connection


type Cloudbuildv2ConnectionGitlabConfigServiceDirectoryConfig struct {
	// Required. The Service Directory service name. Format: projects/{project}/locations/{location}/namespaces/{namespace}/services/{service}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuildv2_connection#service Cloudbuildv2Connection#service}
	Service *string `field:"required" json:"service" yaml:"service"`
}

