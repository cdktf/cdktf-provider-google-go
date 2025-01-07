// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigqueryroutine


type BigqueryRoutineSparkOptions struct {
	// Archive files to be extracted into the working directory of each executor.
	//
	// For more information about Apache Spark, see Apache Spark.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/bigquery_routine#archive_uris BigqueryRoutine#archive_uris}
	ArchiveUris *[]*string `field:"optional" json:"archiveUris" yaml:"archiveUris"`
	// Fully qualified name of the user-provided Spark connection object. Format: "projects/{projectId}/locations/{locationId}/connections/{connectionId}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/bigquery_routine#connection BigqueryRoutine#connection}
	Connection *string `field:"optional" json:"connection" yaml:"connection"`
	// Custom container image for the runtime environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/bigquery_routine#container_image BigqueryRoutine#container_image}
	ContainerImage *string `field:"optional" json:"containerImage" yaml:"containerImage"`
	// Files to be placed in the working directory of each executor.
	//
	// For more information about Apache Spark, see Apache Spark.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/bigquery_routine#file_uris BigqueryRoutine#file_uris}
	FileUris *[]*string `field:"optional" json:"fileUris" yaml:"fileUris"`
	// JARs to include on the driver and executor CLASSPATH. For more information about Apache Spark, see Apache Spark.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/bigquery_routine#jar_uris BigqueryRoutine#jar_uris}
	JarUris *[]*string `field:"optional" json:"jarUris" yaml:"jarUris"`
	// The fully qualified name of a class in jarUris, for example, com.example.wordcount. Exactly one of mainClass and main_jar_uri field should be set for Java/Scala language type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/bigquery_routine#main_class BigqueryRoutine#main_class}
	MainClass *string `field:"optional" json:"mainClass" yaml:"mainClass"`
	// The main file/jar URI of the Spark application.
	//
	// Exactly one of the definitionBody field and the mainFileUri field must be set for Python.
	// Exactly one of mainClass and mainFileUri field should be set for Java/Scala language type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/bigquery_routine#main_file_uri BigqueryRoutine#main_file_uri}
	MainFileUri *string `field:"optional" json:"mainFileUri" yaml:"mainFileUri"`
	// Configuration properties as a set of key/value pairs, which will be passed on to the Spark application.
	//
	// For more information, see Apache Spark and the procedure option list.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/bigquery_routine#properties BigqueryRoutine#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// Python files to be placed on the PYTHONPATH for PySpark application.
	//
	// Supported file types: .py, .egg, and .zip. For more information about Apache Spark, see Apache Spark.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/bigquery_routine#py_file_uris BigqueryRoutine#py_file_uris}
	PyFileUris *[]*string `field:"optional" json:"pyFileUris" yaml:"pyFileUris"`
	// Runtime version. If not specified, the default runtime version is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/bigquery_routine#runtime_version BigqueryRoutine#runtime_version}
	RuntimeVersion *string `field:"optional" json:"runtimeVersion" yaml:"runtimeVersion"`
}

