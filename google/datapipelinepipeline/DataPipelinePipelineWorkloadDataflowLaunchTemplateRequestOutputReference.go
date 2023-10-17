// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datapipelinepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v12/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v12/datapipelinepipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference interface {
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
	GcsPath() *string
	SetGcsPath(val *string)
	GcsPathInput() *string
	InternalValue() *DataPipelinePipelineWorkloadDataflowLaunchTemplateRequest
	SetInternalValue(val *DataPipelinePipelineWorkloadDataflowLaunchTemplateRequest)
	LaunchParameters() DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersOutputReference
	LaunchParametersInput() *DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParameters
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	ProjectId() *string
	SetProjectId(val *string)
	ProjectIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ValidateOnly() interface{}
	SetValidateOnly(val interface{})
	ValidateOnlyInput() interface{}
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
	PutLaunchParameters(value *DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParameters)
	ResetGcsPath()
	ResetLaunchParameters()
	ResetLocation()
	ResetValidateOnly()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference
type jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) GcsPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) GcsPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) InternalValue() *DataPipelinePipelineWorkloadDataflowLaunchTemplateRequest {
	var returns *DataPipelinePipelineWorkloadDataflowLaunchTemplateRequest
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) LaunchParameters() DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersOutputReference {
	var returns DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersOutputReference
	_jsii_.Get(
		j,
		"launchParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) LaunchParametersInput() *DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParameters {
	var returns *DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParameters
	_jsii_.Get(
		j,
		"launchParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) ProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) ValidateOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) ValidateOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateOnlyInput",
		&returns,
	)
	return returns
}


func NewDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference {
	_init_.Initialize()

	if err := validateNewDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataPipelinePipeline.DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference_Override(d DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataPipelinePipeline.DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference)SetGcsPath(val *string) {
	if err := j.validateSetGcsPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcsPath",
		val,
	)
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference)SetInternalValue(val *DataPipelinePipelineWorkloadDataflowLaunchTemplateRequest) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference)SetProjectId(val *string) {
	if err := j.validateSetProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectId",
		val,
	)
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference)SetValidateOnly(val interface{}) {
	if err := j.validateSetValidateOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validateOnly",
		val,
	)
}

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) PutLaunchParameters(value *DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParameters) {
	if err := d.validatePutLaunchParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLaunchParameters",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) ResetGcsPath() {
	_jsii_.InvokeVoid(
		d,
		"resetGcsPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) ResetLaunchParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetLaunchParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		d,
		"resetLocation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) ResetValidateOnly() {
	_jsii_.InvokeVoid(
		d,
		"resetValidateOnly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

