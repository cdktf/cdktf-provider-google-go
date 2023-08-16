package artifactregistryrepository


type ArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepository struct {
	// Address of the remote repository. Default value: "NPMJS" Possible values: ["NPMJS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/artifact_registry_repository#public_repository ArtifactRegistryRepository#public_repository}
	PublicRepository *string `field:"optional" json:"publicRepository" yaml:"publicRepository"`
}

