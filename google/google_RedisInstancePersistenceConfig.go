// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type RedisInstancePersistenceConfig struct {
	// Optional. Available snapshot periods for scheduling.
	//
	// - ONE_HOUR:	Snapshot every 1 hour.
	// - SIX_HOURS:	Snapshot every 6 hours.
	// - TWELVE_HOURS:	Snapshot every 12 hours.
	// - TWENTY_FOUR_HOURS:	Snapshot every 24 horus. Possible values: ["ONE_HOUR", "SIX_HOURS", "TWELVE_HOURS", "TWENTY_FOUR_HOURS"]
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/redis_instance#rdb_snapshot_period RedisInstance#rdb_snapshot_period}
	RdbSnapshotPeriod *string `field:"required" json:"rdbSnapshotPeriod" yaml:"rdbSnapshotPeriod"`
	// Optional. Controls whether Persistence features are enabled. If not provided, the existing value will be used.
	//
	// - DISABLED: 	Persistence is disabled for the instance, and any existing snapshots are deleted.
	// - RDB: RDB based Persistence is enabled. Possible values: ["DISABLED", "RDB"]
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/redis_instance#persistence_mode RedisInstance#persistence_mode}
	PersistenceMode *string `field:"optional" json:"persistenceMode" yaml:"persistenceMode"`
	// Optional.
	//
	// Date and time that the first snapshot was/will be attempted,
	// and to which future snapshots will be aligned. If not provided,
	// the current time will be used.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution
	// and up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/redis_instance#rdb_snapshot_start_time RedisInstance#rdb_snapshot_start_time}
	RdbSnapshotStartTime *string `field:"optional" json:"rdbSnapshotStartTime" yaml:"rdbSnapshotStartTime"`
}
