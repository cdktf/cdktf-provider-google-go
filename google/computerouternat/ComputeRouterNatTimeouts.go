package computerouternat


type ComputeRouterNatTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_router_nat#create ComputeRouterNat#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_router_nat#delete ComputeRouterNat#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_router_nat#update ComputeRouterNat#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
