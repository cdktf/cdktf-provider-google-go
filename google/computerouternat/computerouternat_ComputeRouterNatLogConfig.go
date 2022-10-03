package computerouternat


type ComputeRouterNatLogConfig struct {
	// Indicates whether or not to export logs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_router_nat#enable ComputeRouterNat#enable}
	Enable interface{} `field:"required" json:"enable" yaml:"enable"`
	// Specifies the desired filtering of logs on this NAT. Possible values: ["ERRORS_ONLY", "TRANSLATIONS_ONLY", "ALL"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_router_nat#filter ComputeRouterNat#filter}
	Filter *string `field:"required" json:"filter" yaml:"filter"`
}

