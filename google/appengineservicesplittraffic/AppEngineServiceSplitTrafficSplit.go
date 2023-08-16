package appengineservicesplittraffic


type AppEngineServiceSplitTrafficSplit struct {
	// Mapping from version IDs within the service to fractional (0.000, 1] allocations of traffic for that version. Each version can be specified only once, but some versions in the service may not have any traffic allocation. Services that have traffic allocated cannot be deleted until either the service is deleted or their traffic allocation is removed. Allocations must sum to 1. Up to two decimal place precision is supported for IP-based splits and up to three decimal places is supported for cookie-based splits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/app_engine_service_split_traffic#allocations AppEngineServiceSplitTraffic#allocations}
	Allocations *map[string]*string `field:"required" json:"allocations" yaml:"allocations"`
	// Mechanism used to determine which version a request is sent to.
	//
	// The traffic selection algorithm will be stable for either type until allocations are changed. Possible values: ["UNSPECIFIED", "COOKIE", "IP", "RANDOM"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/app_engine_service_split_traffic#shard_by AppEngineServiceSplitTraffic#shard_by}
	ShardBy *string `field:"optional" json:"shardBy" yaml:"shardBy"`
}

