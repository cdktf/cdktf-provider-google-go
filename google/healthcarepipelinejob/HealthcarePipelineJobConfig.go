// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcarepipelinejob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HealthcarePipelineJobConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Healthcare Dataset under which the Pipeline Job is to run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/healthcare_pipeline_job#dataset HealthcarePipelineJob#dataset}
	Dataset *string `field:"required" json:"dataset" yaml:"dataset"`
	// Location where the Pipeline Job is to run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/healthcare_pipeline_job#location HealthcarePipelineJob#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Specifies the name of the pipeline job. This field is user-assigned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/healthcare_pipeline_job#name HealthcarePipelineJob#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// backfill_pipeline_job block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/healthcare_pipeline_job#backfill_pipeline_job HealthcarePipelineJob#backfill_pipeline_job}
	BackfillPipelineJob *HealthcarePipelineJobBackfillPipelineJob `field:"optional" json:"backfillPipelineJob" yaml:"backfillPipelineJob"`
	// If true, disables writing lineage for the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/healthcare_pipeline_job#disable_lineage HealthcarePipelineJob#disable_lineage}
	DisableLineage interface{} `field:"optional" json:"disableLineage" yaml:"disableLineage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/healthcare_pipeline_job#id HealthcarePipelineJob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// User-supplied key-value pairs used to organize Pipeline Jobs.
	//
	// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of
	// maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	// Label values are optional, must be between 1 and 63 characters long, have a
	// UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE
	// regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	// No more than 64 labels can be associated with a given pipeline.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/healthcare_pipeline_job#labels HealthcarePipelineJob#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// mapping_pipeline_job block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/healthcare_pipeline_job#mapping_pipeline_job HealthcarePipelineJob#mapping_pipeline_job}
	MappingPipelineJob *HealthcarePipelineJobMappingPipelineJob `field:"optional" json:"mappingPipelineJob" yaml:"mappingPipelineJob"`
	// reconciliation_pipeline_job block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/healthcare_pipeline_job#reconciliation_pipeline_job HealthcarePipelineJob#reconciliation_pipeline_job}
	ReconciliationPipelineJob *HealthcarePipelineJobReconciliationPipelineJob `field:"optional" json:"reconciliationPipelineJob" yaml:"reconciliationPipelineJob"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/healthcare_pipeline_job#timeouts HealthcarePipelineJob#timeouts}
	Timeouts *HealthcarePipelineJobTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

