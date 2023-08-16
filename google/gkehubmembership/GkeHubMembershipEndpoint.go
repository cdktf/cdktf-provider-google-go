package gkehubmembership


type GkeHubMembershipEndpoint struct {
	// gke_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/gke_hub_membership#gke_cluster GkeHubMembership#gke_cluster}
	GkeCluster *GkeHubMembershipEndpointGkeCluster `field:"optional" json:"gkeCluster" yaml:"gkeCluster"`
}

