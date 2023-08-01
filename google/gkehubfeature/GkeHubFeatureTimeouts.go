package gkehubfeature


type GkeHubFeatureTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/gke_hub_feature#create GkeHubFeature#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/gke_hub_feature#delete GkeHubFeature#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/gke_hub_feature#update GkeHubFeature#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

