package filestorebackup


type FilestoreBackupTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/filestore_backup#create FilestoreBackup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/filestore_backup#delete FilestoreBackup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/filestore_backup#update FilestoreBackup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

