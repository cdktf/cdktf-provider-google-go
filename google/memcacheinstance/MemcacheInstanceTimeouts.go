package memcacheinstance


type MemcacheInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/memcache_instance#create MemcacheInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/memcache_instance#delete MemcacheInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/memcache_instance#update MemcacheInstance#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
