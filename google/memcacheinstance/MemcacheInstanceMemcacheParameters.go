package memcacheinstance


type MemcacheInstanceMemcacheParameters struct {
	// User-defined set of parameters to use in the memcache process.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/memcache_instance#params MemcacheInstance#params}
	Params *map[string]*string `field:"optional" json:"params" yaml:"params"`
}

