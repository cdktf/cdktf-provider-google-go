package memcacheinstance


type MemcacheInstanceNodeConfig struct {
	// Number of CPUs per node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/memcache_instance#cpu_count MemcacheInstance#cpu_count}
	CpuCount *float64 `field:"required" json:"cpuCount" yaml:"cpuCount"`
	// Memory size in Mebibytes for each memcache node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/memcache_instance#memory_size_mb MemcacheInstance#memory_size_mb}
	MemorySizeMb *float64 `field:"required" json:"memorySizeMb" yaml:"memorySizeMb"`
}

