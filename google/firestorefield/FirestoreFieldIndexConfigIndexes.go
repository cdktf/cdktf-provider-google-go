package firestorefield


type FirestoreFieldIndexConfigIndexes struct {
	// Indicates that this field supports operations on arrayValues. Only one of 'order' and 'arrayConfig' can be specified. Possible values: ["CONTAINS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/firestore_field#array_config FirestoreField#array_config}
	ArrayConfig *string `field:"optional" json:"arrayConfig" yaml:"arrayConfig"`
	// Indicates that this field supports ordering by the specified order or comparing using =, <, <=, >, >=, !=. Only one of 'order' and 'arrayConfig' can be specified. Possible values: ["ASCENDING", "DESCENDING"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/firestore_field#order FirestoreField#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// The scope at which a query is run.
	//
	// Collection scoped queries require you specify
	// the collection at query time. Collection group scope allows queries across all
	// collections with the same id. Default value: "COLLECTION" Possible values: ["COLLECTION", "COLLECTION_GROUP"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/firestore_field#query_scope FirestoreField#query_scope}
	QueryScope *string `field:"optional" json:"queryScope" yaml:"queryScope"`
}

