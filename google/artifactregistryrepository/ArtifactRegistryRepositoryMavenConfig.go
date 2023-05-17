package artifactregistryrepository


type ArtifactRegistryRepositoryMavenConfig struct {
	// The repository with this flag will allow publishing the same snapshot versions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/artifact_registry_repository#allow_snapshot_overwrites ArtifactRegistryRepository#allow_snapshot_overwrites}
	AllowSnapshotOverwrites interface{} `field:"optional" json:"allowSnapshotOverwrites" yaml:"allowSnapshotOverwrites"`
	// Version policy defines the versions that the registry will accept. Default value: "VERSION_POLICY_UNSPECIFIED" Possible values: ["VERSION_POLICY_UNSPECIFIED", "RELEASE", "SNAPSHOT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/artifact_registry_repository#version_policy ArtifactRegistryRepository#version_policy}
	VersionPolicy *string `field:"optional" json:"versionPolicy" yaml:"versionPolicy"`
}

