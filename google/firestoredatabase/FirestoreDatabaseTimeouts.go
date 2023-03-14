package firestoredatabase


type FirestoreDatabaseTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/firestore_database#create FirestoreDatabase#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/firestore_database#delete FirestoreDatabase#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/firestore_database#update FirestoreDatabase#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

