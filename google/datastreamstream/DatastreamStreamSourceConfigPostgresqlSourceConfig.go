package datastreamstream


type DatastreamStreamSourceConfigPostgresqlSourceConfig struct {
	// The name of the publication that includes the set of all tables that are defined in the stream's include_objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/datastream_stream#publication DatastreamStream#publication}
	Publication *string `field:"required" json:"publication" yaml:"publication"`
	// The name of the logical replication slot that's configured with the pgoutput plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/datastream_stream#replication_slot DatastreamStream#replication_slot}
	ReplicationSlot *string `field:"required" json:"replicationSlot" yaml:"replicationSlot"`
	// exclude_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/datastream_stream#exclude_objects DatastreamStream#exclude_objects}
	ExcludeObjects *DatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjects `field:"optional" json:"excludeObjects" yaml:"excludeObjects"`
	// include_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/datastream_stream#include_objects DatastreamStream#include_objects}
	IncludeObjects *DatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjects `field:"optional" json:"includeObjects" yaml:"includeObjects"`
	// Maximum number of concurrent backfill tasks.
	//
	// The number should be non
	// negative. If not set (or set to 0), the system's default value will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/datastream_stream#max_concurrent_backfill_tasks DatastreamStream#max_concurrent_backfill_tasks}
	MaxConcurrentBackfillTasks *float64 `field:"optional" json:"maxConcurrentBackfillTasks" yaml:"maxConcurrentBackfillTasks"`
}

