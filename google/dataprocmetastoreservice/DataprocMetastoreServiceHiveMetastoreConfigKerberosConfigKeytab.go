package dataprocmetastoreservice


type DataprocMetastoreServiceHiveMetastoreConfigKerberosConfigKeytab struct {
	// The relative resource name of a Secret Manager secret version, in the following form:.
	//
	// "projects/{projectNumber}/secrets/{secret_id}/versions/{version_id}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataproc_metastore_service#cloud_secret DataprocMetastoreService#cloud_secret}
	CloudSecret *string `field:"required" json:"cloudSecret" yaml:"cloudSecret"`
}

