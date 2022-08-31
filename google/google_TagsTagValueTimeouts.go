// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type TagsTagValueTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/tags_tag_value#create TagsTagValue#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/tags_tag_value#delete TagsTagValue#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/tags_tag_value#update TagsTagValue#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

