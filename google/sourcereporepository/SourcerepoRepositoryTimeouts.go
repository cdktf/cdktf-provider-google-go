package sourcereporepository


type SourcerepoRepositoryTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sourcerepo_repository#create SourcerepoRepository#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sourcerepo_repository#delete SourcerepoRepository#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sourcerepo_repository#update SourcerepoRepository#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

