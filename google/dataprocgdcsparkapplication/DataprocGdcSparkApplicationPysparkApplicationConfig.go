// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocgdcsparkapplication


type DataprocGdcSparkApplicationPysparkApplicationConfig struct {
	// The HCFS URI of the main Python file to use as the driver. Must be a .py file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/dataproc_gdc_spark_application#main_python_file_uri DataprocGdcSparkApplication#main_python_file_uri}
	MainPythonFileUri *string `field:"required" json:"mainPythonFileUri" yaml:"mainPythonFileUri"`
	// HCFS URIs of archives to be extracted into the working directory of each executor.
	//
	// Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/dataproc_gdc_spark_application#archive_uris DataprocGdcSparkApplication#archive_uris}
	ArchiveUris *[]*string `field:"optional" json:"archiveUris" yaml:"archiveUris"`
	// The arguments to pass to the driver.
	//
	// Do not include arguments, such as '--conf', that can be set as job properties, since a collision may occur that causes an incorrect job submission.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/dataproc_gdc_spark_application#args DataprocGdcSparkApplication#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// HCFS URIs of files to be placed in the working directory of each executor. Useful for naively parallel tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/dataproc_gdc_spark_application#file_uris DataprocGdcSparkApplication#file_uris}
	FileUris *[]*string `field:"optional" json:"fileUris" yaml:"fileUris"`
	// HCFS URIs of jar files to add to the CLASSPATHs of the Python driver and tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/dataproc_gdc_spark_application#jar_file_uris DataprocGdcSparkApplication#jar_file_uris}
	JarFileUris *[]*string `field:"optional" json:"jarFileUris" yaml:"jarFileUris"`
	// HCFS file URIs of Python files to pass to the PySpark framework. Supported file types: .py, .egg, and .zip.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/dataproc_gdc_spark_application#python_file_uris DataprocGdcSparkApplication#python_file_uris}
	PythonFileUris *[]*string `field:"optional" json:"pythonFileUris" yaml:"pythonFileUris"`
}

