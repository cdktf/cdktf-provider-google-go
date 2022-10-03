package gkehubmembership


type GkeHubMembershipTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/gke_hub_membership#create GkeHubMembership#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/gke_hub_membership#delete GkeHubMembership#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/gke_hub_membership#update GkeHubMembership#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

