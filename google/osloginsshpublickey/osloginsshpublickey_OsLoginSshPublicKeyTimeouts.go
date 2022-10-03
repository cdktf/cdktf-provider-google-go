package osloginsshpublickey


type OsLoginSshPublicKeyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_login_ssh_public_key#create OsLoginSshPublicKey#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_login_ssh_public_key#delete OsLoginSshPublicKey#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/os_login_ssh_public_key#update OsLoginSshPublicKey#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

