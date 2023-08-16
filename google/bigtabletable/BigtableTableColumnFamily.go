package bigtabletable


type BigtableTableColumnFamily struct {
	// The name of the column family.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigtable_table#family BigtableTable#family}
	Family *string `field:"required" json:"family" yaml:"family"`
}

