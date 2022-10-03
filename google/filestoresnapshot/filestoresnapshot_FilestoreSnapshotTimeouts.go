package filestoresnapshot


type FilestoreSnapshotTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/filestore_snapshot#create FilestoreSnapshot#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/filestore_snapshot#delete FilestoreSnapshot#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/filestore_snapshot#update FilestoreSnapshot#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

