package computetargettcpproxy


type ComputeTargetTcpProxyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_target_tcp_proxy#create ComputeTargetTcpProxy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_target_tcp_proxy#delete ComputeTargetTcpProxy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_target_tcp_proxy#update ComputeTargetTcpProxy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

