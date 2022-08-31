// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type BigqueryTableRangePartitioning struct {
	// The field used to determine how to create a range-based partition.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_table#field BigqueryTable#field}
	Field *string `field:"required" json:"field" yaml:"field"`
	// range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_table#range BigqueryTable#range}
	Range *BigqueryTableRangePartitioningRange `field:"required" json:"range" yaml:"range"`
}

