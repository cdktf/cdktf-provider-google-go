package cloudrunv2job


type CloudRunV2JobTemplateTemplateVolumesCloudSqlInstance struct {
	// The Cloud SQL instance connection names, as can be found in https://console.cloud.google.com/sql/instances. Visit https://cloud.google.com/sql/docs/mysql/connect-run for more information on how to connect Cloud SQL and Cloud Run. Format: {project}:{location}:{instance}.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_run_v2_job#instances CloudRunV2Job#instances}
	Instances *[]*string `field:"optional" json:"instances" yaml:"instances"`
}

