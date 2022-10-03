package bigtabletable


type BigtableTableColumnFamily struct {
	// The name of the column family.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigtable_table#family BigtableTable#family}
	Family *string `field:"required" json:"family" yaml:"family"`
}

