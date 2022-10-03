package memcacheinstance


type MemcacheInstanceMemcacheParameters struct {
	// User-defined set of parameters to use in the memcache process.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/memcache_instance#params MemcacheInstance#params}
	Params *map[string]*string `field:"optional" json:"params" yaml:"params"`
}

