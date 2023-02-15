package alloydbinstance


type AlloydbInstanceMachineConfig struct {
	// The number of CPU's in the VM instance.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/alloydb_instance#cpu_count AlloydbInstance#cpu_count}
	CpuCount *float64 `field:"optional" json:"cpuCount" yaml:"cpuCount"`
}

