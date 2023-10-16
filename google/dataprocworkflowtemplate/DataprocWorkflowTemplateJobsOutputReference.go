// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocworkflowtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v11/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v11/dataprocworkflowtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocWorkflowTemplateJobsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	HadoopJob() DataprocWorkflowTemplateJobsHadoopJobOutputReference
	HadoopJobInput() *DataprocWorkflowTemplateJobsHadoopJob
	HiveJob() DataprocWorkflowTemplateJobsHiveJobOutputReference
	HiveJobInput() *DataprocWorkflowTemplateJobsHiveJob
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	PigJob() DataprocWorkflowTemplateJobsPigJobOutputReference
	PigJobInput() *DataprocWorkflowTemplateJobsPigJob
	PrerequisiteStepIds() *[]*string
	SetPrerequisiteStepIds(val *[]*string)
	PrerequisiteStepIdsInput() *[]*string
	PrestoJob() DataprocWorkflowTemplateJobsPrestoJobOutputReference
	PrestoJobInput() *DataprocWorkflowTemplateJobsPrestoJob
	PysparkJob() DataprocWorkflowTemplateJobsPysparkJobOutputReference
	PysparkJobInput() *DataprocWorkflowTemplateJobsPysparkJob
	Scheduling() DataprocWorkflowTemplateJobsSchedulingOutputReference
	SchedulingInput() *DataprocWorkflowTemplateJobsScheduling
	SparkJob() DataprocWorkflowTemplateJobsSparkJobOutputReference
	SparkJobInput() *DataprocWorkflowTemplateJobsSparkJob
	SparkRJob() DataprocWorkflowTemplateJobsSparkRJobOutputReference
	SparkRJobInput() *DataprocWorkflowTemplateJobsSparkRJob
	SparkSqlJob() DataprocWorkflowTemplateJobsSparkSqlJobOutputReference
	SparkSqlJobInput() *DataprocWorkflowTemplateJobsSparkSqlJob
	StepId() *string
	SetStepId(val *string)
	StepIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutHadoopJob(value *DataprocWorkflowTemplateJobsHadoopJob)
	PutHiveJob(value *DataprocWorkflowTemplateJobsHiveJob)
	PutPigJob(value *DataprocWorkflowTemplateJobsPigJob)
	PutPrestoJob(value *DataprocWorkflowTemplateJobsPrestoJob)
	PutPysparkJob(value *DataprocWorkflowTemplateJobsPysparkJob)
	PutScheduling(value *DataprocWorkflowTemplateJobsScheduling)
	PutSparkJob(value *DataprocWorkflowTemplateJobsSparkJob)
	PutSparkRJob(value *DataprocWorkflowTemplateJobsSparkRJob)
	PutSparkSqlJob(value *DataprocWorkflowTemplateJobsSparkSqlJob)
	ResetHadoopJob()
	ResetHiveJob()
	ResetLabels()
	ResetPigJob()
	ResetPrerequisiteStepIds()
	ResetPrestoJob()
	ResetPysparkJob()
	ResetScheduling()
	ResetSparkJob()
	ResetSparkRJob()
	ResetSparkSqlJob()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocWorkflowTemplateJobsOutputReference
type jsiiProxy_DataprocWorkflowTemplateJobsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) HadoopJob() DataprocWorkflowTemplateJobsHadoopJobOutputReference {
	var returns DataprocWorkflowTemplateJobsHadoopJobOutputReference
	_jsii_.Get(
		j,
		"hadoopJob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) HadoopJobInput() *DataprocWorkflowTemplateJobsHadoopJob {
	var returns *DataprocWorkflowTemplateJobsHadoopJob
	_jsii_.Get(
		j,
		"hadoopJobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) HiveJob() DataprocWorkflowTemplateJobsHiveJobOutputReference {
	var returns DataprocWorkflowTemplateJobsHiveJobOutputReference
	_jsii_.Get(
		j,
		"hiveJob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) HiveJobInput() *DataprocWorkflowTemplateJobsHiveJob {
	var returns *DataprocWorkflowTemplateJobsHiveJob
	_jsii_.Get(
		j,
		"hiveJobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) PigJob() DataprocWorkflowTemplateJobsPigJobOutputReference {
	var returns DataprocWorkflowTemplateJobsPigJobOutputReference
	_jsii_.Get(
		j,
		"pigJob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) PigJobInput() *DataprocWorkflowTemplateJobsPigJob {
	var returns *DataprocWorkflowTemplateJobsPigJob
	_jsii_.Get(
		j,
		"pigJobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) PrerequisiteStepIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"prerequisiteStepIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) PrerequisiteStepIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"prerequisiteStepIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) PrestoJob() DataprocWorkflowTemplateJobsPrestoJobOutputReference {
	var returns DataprocWorkflowTemplateJobsPrestoJobOutputReference
	_jsii_.Get(
		j,
		"prestoJob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) PrestoJobInput() *DataprocWorkflowTemplateJobsPrestoJob {
	var returns *DataprocWorkflowTemplateJobsPrestoJob
	_jsii_.Get(
		j,
		"prestoJobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) PysparkJob() DataprocWorkflowTemplateJobsPysparkJobOutputReference {
	var returns DataprocWorkflowTemplateJobsPysparkJobOutputReference
	_jsii_.Get(
		j,
		"pysparkJob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) PysparkJobInput() *DataprocWorkflowTemplateJobsPysparkJob {
	var returns *DataprocWorkflowTemplateJobsPysparkJob
	_jsii_.Get(
		j,
		"pysparkJobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) Scheduling() DataprocWorkflowTemplateJobsSchedulingOutputReference {
	var returns DataprocWorkflowTemplateJobsSchedulingOutputReference
	_jsii_.Get(
		j,
		"scheduling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) SchedulingInput() *DataprocWorkflowTemplateJobsScheduling {
	var returns *DataprocWorkflowTemplateJobsScheduling
	_jsii_.Get(
		j,
		"schedulingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) SparkJob() DataprocWorkflowTemplateJobsSparkJobOutputReference {
	var returns DataprocWorkflowTemplateJobsSparkJobOutputReference
	_jsii_.Get(
		j,
		"sparkJob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) SparkJobInput() *DataprocWorkflowTemplateJobsSparkJob {
	var returns *DataprocWorkflowTemplateJobsSparkJob
	_jsii_.Get(
		j,
		"sparkJobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) SparkRJob() DataprocWorkflowTemplateJobsSparkRJobOutputReference {
	var returns DataprocWorkflowTemplateJobsSparkRJobOutputReference
	_jsii_.Get(
		j,
		"sparkRJob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) SparkRJobInput() *DataprocWorkflowTemplateJobsSparkRJob {
	var returns *DataprocWorkflowTemplateJobsSparkRJob
	_jsii_.Get(
		j,
		"sparkRJobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) SparkSqlJob() DataprocWorkflowTemplateJobsSparkSqlJobOutputReference {
	var returns DataprocWorkflowTemplateJobsSparkSqlJobOutputReference
	_jsii_.Get(
		j,
		"sparkSqlJob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) SparkSqlJobInput() *DataprocWorkflowTemplateJobsSparkSqlJob {
	var returns *DataprocWorkflowTemplateJobsSparkSqlJob
	_jsii_.Get(
		j,
		"sparkSqlJobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) StepId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stepId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) StepIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stepIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataprocWorkflowTemplateJobsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataprocWorkflowTemplateJobsOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocWorkflowTemplateJobsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocWorkflowTemplateJobsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocWorkflowTemplate.DataprocWorkflowTemplateJobsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataprocWorkflowTemplateJobsOutputReference_Override(d DataprocWorkflowTemplateJobsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocWorkflowTemplate.DataprocWorkflowTemplateJobsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference)SetPrerequisiteStepIds(val *[]*string) {
	if err := j.validateSetPrerequisiteStepIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prerequisiteStepIds",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference)SetStepId(val *string) {
	if err := j.validateSetStepIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stepId",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) PutHadoopJob(value *DataprocWorkflowTemplateJobsHadoopJob) {
	if err := d.validatePutHadoopJobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHadoopJob",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) PutHiveJob(value *DataprocWorkflowTemplateJobsHiveJob) {
	if err := d.validatePutHiveJobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHiveJob",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) PutPigJob(value *DataprocWorkflowTemplateJobsPigJob) {
	if err := d.validatePutPigJobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPigJob",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) PutPrestoJob(value *DataprocWorkflowTemplateJobsPrestoJob) {
	if err := d.validatePutPrestoJobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPrestoJob",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) PutPysparkJob(value *DataprocWorkflowTemplateJobsPysparkJob) {
	if err := d.validatePutPysparkJobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPysparkJob",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) PutScheduling(value *DataprocWorkflowTemplateJobsScheduling) {
	if err := d.validatePutSchedulingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putScheduling",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) PutSparkJob(value *DataprocWorkflowTemplateJobsSparkJob) {
	if err := d.validatePutSparkJobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSparkJob",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) PutSparkRJob(value *DataprocWorkflowTemplateJobsSparkRJob) {
	if err := d.validatePutSparkRJobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSparkRJob",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) PutSparkSqlJob(value *DataprocWorkflowTemplateJobsSparkSqlJob) {
	if err := d.validatePutSparkSqlJobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSparkSqlJob",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) ResetHadoopJob() {
	_jsii_.InvokeVoid(
		d,
		"resetHadoopJob",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) ResetHiveJob() {
	_jsii_.InvokeVoid(
		d,
		"resetHiveJob",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		d,
		"resetLabels",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) ResetPigJob() {
	_jsii_.InvokeVoid(
		d,
		"resetPigJob",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) ResetPrerequisiteStepIds() {
	_jsii_.InvokeVoid(
		d,
		"resetPrerequisiteStepIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) ResetPrestoJob() {
	_jsii_.InvokeVoid(
		d,
		"resetPrestoJob",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) ResetPysparkJob() {
	_jsii_.InvokeVoid(
		d,
		"resetPysparkJob",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) ResetScheduling() {
	_jsii_.InvokeVoid(
		d,
		"resetScheduling",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) ResetSparkJob() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkJob",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) ResetSparkRJob() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkRJob",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) ResetSparkSqlJob() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkSqlJob",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

