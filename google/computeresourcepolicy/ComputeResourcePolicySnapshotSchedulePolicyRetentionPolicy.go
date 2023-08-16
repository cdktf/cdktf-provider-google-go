package computeresourcepolicy


type ComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy struct {
	// Maximum age of the snapshot that is allowed to be kept.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_resource_policy#max_retention_days ComputeResourcePolicy#max_retention_days}
	MaxRetentionDays *float64 `field:"required" json:"maxRetentionDays" yaml:"maxRetentionDays"`
	// Specifies the behavior to apply to scheduled snapshots when the source disk is deleted.
	//
	// Default value: "KEEP_AUTO_SNAPSHOTS" Possible values: ["KEEP_AUTO_SNAPSHOTS", "APPLY_RETENTION_POLICY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_resource_policy#on_source_disk_delete ComputeResourcePolicy#on_source_disk_delete}
	OnSourceDiskDelete *string `field:"optional" json:"onSourceDiskDelete" yaml:"onSourceDiskDelete"`
}

