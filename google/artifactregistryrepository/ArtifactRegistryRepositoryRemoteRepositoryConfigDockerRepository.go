package artifactregistryrepository


type ArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepository struct {
	// Address of the remote repository. Default value: "DOCKER_HUB" Possible values: ["DOCKER_HUB"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/artifact_registry_repository#public_repository ArtifactRegistryRepository#public_repository}
	PublicRepository *string `field:"optional" json:"publicRepository" yaml:"publicRepository"`
}
