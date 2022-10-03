package firestoreindex


type FirestoreIndexFields struct {
	// Indicates that this field supports operations on arrayValues. Only one of 'order' and 'arrayConfig' can be specified. Possible values: ["CONTAINS"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/firestore_index#array_config FirestoreIndex#array_config}
	ArrayConfig *string `field:"optional" json:"arrayConfig" yaml:"arrayConfig"`
	// Name of the field.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/firestore_index#field_path FirestoreIndex#field_path}
	FieldPath *string `field:"optional" json:"fieldPath" yaml:"fieldPath"`
	// Indicates that this field supports ordering by the specified order or comparing using =, <, <=, >, >=.
	//
	// Only one of 'order' and 'arrayConfig' can be specified. Possible values: ["ASCENDING", "DESCENDING"]
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/firestore_index#order FirestoreIndex#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
}

