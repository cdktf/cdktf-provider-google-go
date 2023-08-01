package appengineflexibleappversion


type AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilization struct {
	// Target bytes received per second.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/app_engine_flexible_app_version#target_received_bytes_per_second AppEngineFlexibleAppVersion#target_received_bytes_per_second}
	TargetReceivedBytesPerSecond *float64 `field:"optional" json:"targetReceivedBytesPerSecond" yaml:"targetReceivedBytesPerSecond"`
	// Target packets received per second.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/app_engine_flexible_app_version#target_received_packets_per_second AppEngineFlexibleAppVersion#target_received_packets_per_second}
	TargetReceivedPacketsPerSecond *float64 `field:"optional" json:"targetReceivedPacketsPerSecond" yaml:"targetReceivedPacketsPerSecond"`
	// Target bytes sent per second.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/app_engine_flexible_app_version#target_sent_bytes_per_second AppEngineFlexibleAppVersion#target_sent_bytes_per_second}
	TargetSentBytesPerSecond *float64 `field:"optional" json:"targetSentBytesPerSecond" yaml:"targetSentBytesPerSecond"`
	// Target packets sent per second.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/app_engine_flexible_app_version#target_sent_packets_per_second AppEngineFlexibleAppVersion#target_sent_packets_per_second}
	TargetSentPacketsPerSecond *float64 `field:"optional" json:"targetSentPacketsPerSecond" yaml:"targetSentPacketsPerSecond"`
}

