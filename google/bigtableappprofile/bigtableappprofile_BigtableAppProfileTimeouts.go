package bigtableappprofile


type BigtableAppProfileTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_app_profile#create BigtableAppProfile#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_app_profile#delete BigtableAppProfile#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_app_profile#update BigtableAppProfile#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

