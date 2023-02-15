package memcacheinstance


type MemcacheInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime struct {
	// Hours of day in 24 hour format.
	//
	// Should be from 0 to 23.
	// An API may choose to allow the value "24:00:00" for scenarios like business closing time.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/memcache_instance#hours MemcacheInstance#hours}
	Hours *float64 `field:"optional" json:"hours" yaml:"hours"`
	// Minutes of hour of day. Must be from 0 to 59.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/memcache_instance#minutes MemcacheInstance#minutes}
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
	// Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/memcache_instance#nanos MemcacheInstance#nanos}
	Nanos *float64 `field:"optional" json:"nanos" yaml:"nanos"`
	// Seconds of minutes of the time.
	//
	// Must normally be from 0 to 59.
	// An API may allow the value 60 if it allows leap-seconds.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/memcache_instance#seconds MemcacheInstance#seconds}
	Seconds *float64 `field:"optional" json:"seconds" yaml:"seconds"`
}

