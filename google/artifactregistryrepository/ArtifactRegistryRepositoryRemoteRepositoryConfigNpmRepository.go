package artifactregistryrepository


type ArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepository struct {
	// Address of the remote repository. Default value: "NPMJS" Possible values: ["NPMJS"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/artifact_registry_repository#public_repository ArtifactRegistryRepository#public_repository}
	PublicRepository *string `field:"optional" json:"publicRepository" yaml:"publicRepository"`
}
