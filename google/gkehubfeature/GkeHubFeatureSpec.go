package gkehubfeature


type GkeHubFeatureSpec struct {
	// multiclusteringress block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.72.0/docs/resources/gke_hub_feature#multiclusteringress GkeHubFeature#multiclusteringress}
	Multiclusteringress *GkeHubFeatureSpecMulticlusteringress `field:"optional" json:"multiclusteringress" yaml:"multiclusteringress"`
}

