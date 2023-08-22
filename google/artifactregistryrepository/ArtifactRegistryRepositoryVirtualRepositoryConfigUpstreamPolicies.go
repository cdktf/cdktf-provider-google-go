package artifactregistryrepository


type ArtifactRegistryRepositoryVirtualRepositoryConfigUpstreamPolicies struct {
	// The user-provided ID of the upstream policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/artifact_registry_repository#id ArtifactRegistryRepository#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Entries with a greater priority value take precedence in the pull order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/artifact_registry_repository#priority ArtifactRegistryRepository#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// A reference to the repository resource, for example: "projects/p1/locations/us-central1/repository/repo1".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/artifact_registry_repository#repository ArtifactRegistryRepository#repository}
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
}

