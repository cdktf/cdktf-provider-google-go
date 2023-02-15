package memcacheinstance


type MemcacheInstanceMaintenancePolicy struct {
	// weekly_maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/memcache_instance#weekly_maintenance_window MemcacheInstance#weekly_maintenance_window}
	WeeklyMaintenanceWindow interface{} `field:"required" json:"weeklyMaintenanceWindow" yaml:"weeklyMaintenanceWindow"`
	// Optional. Description of what this policy is for. Create/Update methods return INVALID_ARGUMENT if the length is greater than 512.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/memcache_instance#description MemcacheInstance#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

