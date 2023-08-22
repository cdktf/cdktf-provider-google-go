package artifactregistryrepository


type ArtifactRegistryRepositoryVirtualRepositoryConfig struct {
	// upstream_policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/artifact_registry_repository#upstream_policies ArtifactRegistryRepository#upstream_policies}
	UpstreamPolicies interface{} `field:"optional" json:"upstreamPolicies" yaml:"upstreamPolicies"`
}

