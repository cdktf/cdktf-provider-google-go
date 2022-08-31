// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type MemcacheInstanceNodeConfig struct {
	// Number of CPUs per node.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/memcache_instance#cpu_count MemcacheInstance#cpu_count}
	CpuCount *float64 `field:"required" json:"cpuCount" yaml:"cpuCount"`
	// Memory size in Mebibytes for each memcache node.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/memcache_instance#memory_size_mb MemcacheInstance#memory_size_mb}
	MemorySizeMb *float64 `field:"required" json:"memorySizeMb" yaml:"memorySizeMb"`
}

