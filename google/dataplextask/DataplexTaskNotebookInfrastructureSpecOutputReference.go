// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplextask

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/dataplextask/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataplexTaskNotebookInfrastructureSpecOutputReference interface {
	cdktf.ComplexObject
	Batch() DataplexTaskNotebookInfrastructureSpecBatchOutputReference
	BatchInput() *DataplexTaskNotebookInfrastructureSpecBatch
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
	ContainerImage() DataplexTaskNotebookInfrastructureSpecContainerImageOutputReference
	ContainerImageInput() *DataplexTaskNotebookInfrastructureSpecContainerImage
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DataplexTaskNotebookInfrastructureSpec
	SetInternalValue(val *DataplexTaskNotebookInfrastructureSpec)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpcNetwork() DataplexTaskNotebookInfrastructureSpecVpcNetworkOutputReference
	VpcNetworkInput() *DataplexTaskNotebookInfrastructureSpecVpcNetwork
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
	PutBatch(value *DataplexTaskNotebookInfrastructureSpecBatch)
	PutContainerImage(value *DataplexTaskNotebookInfrastructureSpecContainerImage)
	PutVpcNetwork(value *DataplexTaskNotebookInfrastructureSpecVpcNetwork)
	ResetBatch()
	ResetContainerImage()
	ResetVpcNetwork()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataplexTaskNotebookInfrastructureSpecOutputReference
type jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) Batch() DataplexTaskNotebookInfrastructureSpecBatchOutputReference {
	var returns DataplexTaskNotebookInfrastructureSpecBatchOutputReference
	_jsii_.Get(
		j,
		"batch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) BatchInput() *DataplexTaskNotebookInfrastructureSpecBatch {
	var returns *DataplexTaskNotebookInfrastructureSpecBatch
	_jsii_.Get(
		j,
		"batchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) ContainerImage() DataplexTaskNotebookInfrastructureSpecContainerImageOutputReference {
	var returns DataplexTaskNotebookInfrastructureSpecContainerImageOutputReference
	_jsii_.Get(
		j,
		"containerImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) ContainerImageInput() *DataplexTaskNotebookInfrastructureSpecContainerImage {
	var returns *DataplexTaskNotebookInfrastructureSpecContainerImage
	_jsii_.Get(
		j,
		"containerImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) InternalValue() *DataplexTaskNotebookInfrastructureSpec {
	var returns *DataplexTaskNotebookInfrastructureSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) VpcNetwork() DataplexTaskNotebookInfrastructureSpecVpcNetworkOutputReference {
	var returns DataplexTaskNotebookInfrastructureSpecVpcNetworkOutputReference
	_jsii_.Get(
		j,
		"vpcNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) VpcNetworkInput() *DataplexTaskNotebookInfrastructureSpecVpcNetwork {
	var returns *DataplexTaskNotebookInfrastructureSpecVpcNetwork
	_jsii_.Get(
		j,
		"vpcNetworkInput",
		&returns,
	)
	return returns
}


func NewDataplexTaskNotebookInfrastructureSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataplexTaskNotebookInfrastructureSpecOutputReference {
	_init_.Initialize()

	if err := validateNewDataplexTaskNotebookInfrastructureSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataplexTask.DataplexTaskNotebookInfrastructureSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataplexTaskNotebookInfrastructureSpecOutputReference_Override(d DataplexTaskNotebookInfrastructureSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataplexTask.DataplexTaskNotebookInfrastructureSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference)SetInternalValue(val *DataplexTaskNotebookInfrastructureSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) PutBatch(value *DataplexTaskNotebookInfrastructureSpecBatch) {
	if err := d.validatePutBatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBatch",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) PutContainerImage(value *DataplexTaskNotebookInfrastructureSpecContainerImage) {
	if err := d.validatePutContainerImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putContainerImage",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) PutVpcNetwork(value *DataplexTaskNotebookInfrastructureSpecVpcNetwork) {
	if err := d.validatePutVpcNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVpcNetwork",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) ResetBatch() {
	_jsii_.InvokeVoid(
		d,
		"resetBatch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) ResetContainerImage() {
	_jsii_.InvokeVoid(
		d,
		"resetContainerImage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) ResetVpcNetwork() {
	_jsii_.InvokeVoid(
		d,
		"resetVpcNetwork",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataplexTaskNotebookInfrastructureSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

