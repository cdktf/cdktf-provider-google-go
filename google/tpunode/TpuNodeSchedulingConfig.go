package tpunode


type TpuNodeSchedulingConfig struct {
	// Defines whether the TPU instance is preemptible.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/tpu_node#preemptible TpuNode#preemptible}
	Preemptible interface{} `field:"required" json:"preemptible" yaml:"preemptible"`
}

