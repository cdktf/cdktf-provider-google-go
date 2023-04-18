package tpunode


type TpuNodeSchedulingConfig struct {
	// Defines whether the TPU instance is preemptible.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.62.0/docs/resources/tpu_node#preemptible TpuNode#preemptible}
	Preemptible interface{} `field:"required" json:"preemptible" yaml:"preemptible"`
}

