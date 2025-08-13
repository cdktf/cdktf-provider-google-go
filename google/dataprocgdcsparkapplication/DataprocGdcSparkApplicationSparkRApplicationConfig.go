// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocgdcsparkapplication


type DataprocGdcSparkApplicationSparkRApplicationConfig struct {
	// The HCFS URI of the main R file to use as the driver. Must be a .R file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dataproc_gdc_spark_application#main_r_file_uri DataprocGdcSparkApplication#main_r_file_uri}
	MainRFileUri *string `field:"required" json:"mainRFileUri" yaml:"mainRFileUri"`
	// HCFS URIs of archives to be extracted into the working directory of each executor.
	//
	// Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dataproc_gdc_spark_application#archive_uris DataprocGdcSparkApplication#archive_uris}
	ArchiveUris *[]*string `field:"optional" json:"archiveUris" yaml:"archiveUris"`
	// The arguments to pass to the driver.
	//
	// Do not include arguments, such as '--conf', that can be set as job properties, since a collision may occur that causes an incorrect job submission.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dataproc_gdc_spark_application#args DataprocGdcSparkApplication#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// HCFS URIs of files to be placed in the working directory of each executor. Useful for naively parallel tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/dataproc_gdc_spark_application#file_uris DataprocGdcSparkApplication#file_uris}
	FileUris *[]*string `field:"optional" json:"fileUris" yaml:"fileUris"`
}

