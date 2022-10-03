package appengineflexibleappversion


type AppEngineFlexibleAppVersionResources struct {
	// Number of CPU cores needed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#cpu AppEngineFlexibleAppVersion#cpu}
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// Disk size (GB) needed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#disk_gb AppEngineFlexibleAppVersion#disk_gb}
	DiskGb *float64 `field:"optional" json:"diskGb" yaml:"diskGb"`
	// Memory (GB) needed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#memory_gb AppEngineFlexibleAppVersion#memory_gb}
	MemoryGb *float64 `field:"optional" json:"memoryGb" yaml:"memoryGb"`
	// volumes block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#volumes AppEngineFlexibleAppVersion#volumes}
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}

