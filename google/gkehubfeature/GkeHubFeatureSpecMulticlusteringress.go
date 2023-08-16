package gkehubfeature


type GkeHubFeatureSpecMulticlusteringress struct {
	// Fully-qualified Membership name which hosts the MultiClusterIngress CRD. Example: 'projects/foo-proj/locations/global/memberships/bar'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/gke_hub_feature#config_membership GkeHubFeature#config_membership}
	ConfigMembership *string `field:"required" json:"configMembership" yaml:"configMembership"`
}

