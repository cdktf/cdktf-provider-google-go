package gkehubmembership


type GkeHubMembershipEndpoint struct {
	// gke_cluster block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/gke_hub_membership#gke_cluster GkeHubMembership#gke_cluster}
	GkeCluster *GkeHubMembershipEndpointGkeCluster `field:"optional" json:"gkeCluster" yaml:"gkeCluster"`
}

