package firestoreindex


type FirestoreIndexTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/firestore_index#create FirestoreIndex#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/firestore_index#delete FirestoreIndex#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

