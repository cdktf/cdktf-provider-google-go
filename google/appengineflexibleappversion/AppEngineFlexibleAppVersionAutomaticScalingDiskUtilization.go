package appengineflexibleappversion


type AppEngineFlexibleAppVersionAutomaticScalingDiskUtilization struct {
	// Target bytes read per second.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#target_read_bytes_per_second AppEngineFlexibleAppVersion#target_read_bytes_per_second}
	TargetReadBytesPerSecond *float64 `field:"optional" json:"targetReadBytesPerSecond" yaml:"targetReadBytesPerSecond"`
	// Target ops read per seconds.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#target_read_ops_per_second AppEngineFlexibleAppVersion#target_read_ops_per_second}
	TargetReadOpsPerSecond *float64 `field:"optional" json:"targetReadOpsPerSecond" yaml:"targetReadOpsPerSecond"`
	// Target bytes written per second.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#target_write_bytes_per_second AppEngineFlexibleAppVersion#target_write_bytes_per_second}
	TargetWriteBytesPerSecond *float64 `field:"optional" json:"targetWriteBytesPerSecond" yaml:"targetWriteBytesPerSecond"`
	// Target ops written per second.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#target_write_ops_per_second AppEngineFlexibleAppVersion#target_write_ops_per_second}
	TargetWriteOpsPerSecond *float64 `field:"optional" json:"targetWriteOpsPerSecond" yaml:"targetWriteOpsPerSecond"`
}

