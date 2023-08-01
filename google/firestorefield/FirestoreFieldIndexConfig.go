package firestorefield


type FirestoreFieldIndexConfig struct {
	// indexes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/firestore_field#indexes FirestoreField#indexes}
	Indexes interface{} `field:"optional" json:"indexes" yaml:"indexes"`
}

