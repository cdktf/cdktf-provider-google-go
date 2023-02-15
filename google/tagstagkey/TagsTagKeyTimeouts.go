package tagstagkey


type TagsTagKeyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/tags_tag_key#create TagsTagKey#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/tags_tag_key#delete TagsTagKey#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/tags_tag_key#update TagsTagKey#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

