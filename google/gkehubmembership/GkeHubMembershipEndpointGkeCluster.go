package gkehubmembership


type GkeHubMembershipEndpointGkeCluster struct {
	// Self-link of the GCP resource for the GKE cluster.
	//
	// For example: '//container.googleapis.com/projects/my-project/zones/us-west1-a/clusters/my-cluster'.
	// It can be at the most 1000 characters in length. If the cluster is provisioned with Terraform,
	// this can be '"//container.googleapis.com/${google_container_cluster.my-cluster.id}"' or
	// 'google_container_cluster.my-cluster.id'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/gke_hub_membership#resource_link GkeHubMembership#resource_link}
	ResourceLink *string `field:"required" json:"resourceLink" yaml:"resourceLink"`
}

