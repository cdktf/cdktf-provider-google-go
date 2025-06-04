// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datapipelinepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/datapipelinepipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference interface {
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
	ContainerSpecGcsPath() *string
	SetContainerSpecGcsPath(val *string)
	ContainerSpecGcsPathInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Environment() DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentOutputReference
	EnvironmentInput() *DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironment
	// Experimental.
	Fqn() *string
	InternalValue() *DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameter
	SetInternalValue(val *DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameter)
	JobName() *string
	SetJobName(val *string)
	JobNameInput() *string
	LaunchOptions() *map[string]*string
	SetLaunchOptions(val *map[string]*string)
	LaunchOptionsInput() *map[string]*string
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransformNameMappings() *map[string]*string
	SetTransformNameMappings(val *map[string]*string)
	TransformNameMappingsInput() *map[string]*string
	Update() interface{}
	SetUpdate(val interface{})
	UpdateInput() interface{}
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
	PutEnvironment(value *DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironment)
	ResetContainerSpecGcsPath()
	ResetEnvironment()
	ResetLaunchOptions()
	ResetParameters()
	ResetTransformNameMappings()
	ResetUpdate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference
type jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) ContainerSpecGcsPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerSpecGcsPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) ContainerSpecGcsPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerSpecGcsPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) Environment() DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentOutputReference {
	var returns DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironmentOutputReference
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) EnvironmentInput() *DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironment {
	var returns *DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironment
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) InternalValue() *DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameter {
	var returns *DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) JobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) JobNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) LaunchOptions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"launchOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) LaunchOptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"launchOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) TransformNameMappings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"transformNameMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) TransformNameMappingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"transformNameMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) Update() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) UpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference {
	_init_.Initialize()

	if err := validateNewDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataPipelinePipeline.DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference_Override(d DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataPipelinePipeline.DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference)SetContainerSpecGcsPath(val *string) {
	if err := j.validateSetContainerSpecGcsPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerSpecGcsPath",
		val,
	)
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference)SetInternalValue(val *DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference)SetJobName(val *string) {
	if err := j.validateSetJobNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobName",
		val,
	)
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference)SetLaunchOptions(val *map[string]*string) {
	if err := j.validateSetLaunchOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchOptions",
		val,
	)
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference)SetTransformNameMappings(val *map[string]*string) {
	if err := j.validateSetTransformNameMappingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transformNameMappings",
		val,
	)
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference)SetUpdate(val interface{}) {
	if err := j.validateSetUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) PutEnvironment(value *DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironment) {
	if err := d.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) ResetContainerSpecGcsPath() {
	_jsii_.InvokeVoid(
		d,
		"resetContainerSpecGcsPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		d,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) ResetLaunchOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetLaunchOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) ResetTransformNameMappings() {
	_jsii_.InvokeVoid(
		d,
		"resetTransformNameMappings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

