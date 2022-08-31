// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ComputeRouterPeerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_router_peer#create ComputeRouterPeer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_router_peer#delete ComputeRouterPeer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_router_peer#update ComputeRouterPeer#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

