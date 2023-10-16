// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datapipelinepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v11/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v11/datapipelinepipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataPipelinePipelineWorkloadOutputReference interface {
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
	DataflowFlexTemplateRequest() DataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference
	DataflowFlexTemplateRequestInput() *DataPipelinePipelineWorkloadDataflowFlexTemplateRequest
	DataflowLaunchTemplateRequest() DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference
	DataflowLaunchTemplateRequestInput() *DataPipelinePipelineWorkloadDataflowLaunchTemplateRequest
	// Experimental.
	Fqn() *string
	InternalValue() *DataPipelinePipelineWorkload
	SetInternalValue(val *DataPipelinePipelineWorkload)
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
	PutDataflowFlexTemplateRequest(value *DataPipelinePipelineWorkloadDataflowFlexTemplateRequest)
	PutDataflowLaunchTemplateRequest(value *DataPipelinePipelineWorkloadDataflowLaunchTemplateRequest)
	ResetDataflowFlexTemplateRequest()
	ResetDataflowLaunchTemplateRequest()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataPipelinePipelineWorkloadOutputReference
type jsiiProxy_DataPipelinePipelineWorkloadOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) DataflowFlexTemplateRequest() DataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference {
	var returns DataPipelinePipelineWorkloadDataflowFlexTemplateRequestOutputReference
	_jsii_.Get(
		j,
		"dataflowFlexTemplateRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) DataflowFlexTemplateRequestInput() *DataPipelinePipelineWorkloadDataflowFlexTemplateRequest {
	var returns *DataPipelinePipelineWorkloadDataflowFlexTemplateRequest
	_jsii_.Get(
		j,
		"dataflowFlexTemplateRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) DataflowLaunchTemplateRequest() DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference {
	var returns DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestOutputReference
	_jsii_.Get(
		j,
		"dataflowLaunchTemplateRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) DataflowLaunchTemplateRequestInput() *DataPipelinePipelineWorkloadDataflowLaunchTemplateRequest {
	var returns *DataPipelinePipelineWorkloadDataflowLaunchTemplateRequest
	_jsii_.Get(
		j,
		"dataflowLaunchTemplateRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) InternalValue() *DataPipelinePipelineWorkload {
	var returns *DataPipelinePipelineWorkload
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataPipelinePipelineWorkloadOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataPipelinePipelineWorkloadOutputReference {
	_init_.Initialize()

	if err := validateNewDataPipelinePipelineWorkloadOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataPipelinePipelineWorkloadOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataPipelinePipeline.DataPipelinePipelineWorkloadOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataPipelinePipelineWorkloadOutputReference_Override(d DataPipelinePipelineWorkloadOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataPipelinePipeline.DataPipelinePipelineWorkloadOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadOutputReference)SetInternalValue(val *DataPipelinePipelineWorkload) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataPipelinePipelineWorkloadOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) PutDataflowFlexTemplateRequest(value *DataPipelinePipelineWorkloadDataflowFlexTemplateRequest) {
	if err := d.validatePutDataflowFlexTemplateRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDataflowFlexTemplateRequest",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) PutDataflowLaunchTemplateRequest(value *DataPipelinePipelineWorkloadDataflowLaunchTemplateRequest) {
	if err := d.validatePutDataflowLaunchTemplateRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDataflowLaunchTemplateRequest",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) ResetDataflowFlexTemplateRequest() {
	_jsii_.InvokeVoid(
		d,
		"resetDataflowFlexTemplateRequest",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) ResetDataflowLaunchTemplateRequest() {
	_jsii_.InvokeVoid(
		d,
		"resetDataflowLaunchTemplateRequest",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataPipelinePipelineWorkloadOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

