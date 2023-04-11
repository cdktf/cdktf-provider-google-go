package artifactregistryrepository


type ArtifactRegistryRepositoryVirtualRepositoryConfig struct {
	// upstream_policies block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/artifact_registry_repository#upstream_policies ArtifactRegistryRepository#upstream_policies}
	UpstreamPolicies interface{} `field:"optional" json:"upstreamPolicies" yaml:"upstreamPolicies"`
}

