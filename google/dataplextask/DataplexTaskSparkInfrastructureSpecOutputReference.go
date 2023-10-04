// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplextask

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v10/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v10/dataplextask/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataplexTaskSparkInfrastructureSpecOutputReference interface {
	cdktf.ComplexObject
	Batch() DataplexTaskSparkInfrastructureSpecBatchOutputReference
	BatchInput() *DataplexTaskSparkInfrastructureSpecBatch
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
	ContainerImage() DataplexTaskSparkInfrastructureSpecContainerImageOutputReference
	ContainerImageInput() *DataplexTaskSparkInfrastructureSpecContainerImage
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DataplexTaskSparkInfrastructureSpec
	SetInternalValue(val *DataplexTaskSparkInfrastructureSpec)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpcNetwork() DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference
	VpcNetworkInput() *DataplexTaskSparkInfrastructureSpecVpcNetwork
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
	PutBatch(value *DataplexTaskSparkInfrastructureSpecBatch)
	PutContainerImage(value *DataplexTaskSparkInfrastructureSpecContainerImage)
	PutVpcNetwork(value *DataplexTaskSparkInfrastructureSpecVpcNetwork)
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

// The jsii proxy struct for DataplexTaskSparkInfrastructureSpecOutputReference
type jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) Batch() DataplexTaskSparkInfrastructureSpecBatchOutputReference {
	var returns DataplexTaskSparkInfrastructureSpecBatchOutputReference
	_jsii_.Get(
		j,
		"batch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) BatchInput() *DataplexTaskSparkInfrastructureSpecBatch {
	var returns *DataplexTaskSparkInfrastructureSpecBatch
	_jsii_.Get(
		j,
		"batchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) ContainerImage() DataplexTaskSparkInfrastructureSpecContainerImageOutputReference {
	var returns DataplexTaskSparkInfrastructureSpecContainerImageOutputReference
	_jsii_.Get(
		j,
		"containerImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) ContainerImageInput() *DataplexTaskSparkInfrastructureSpecContainerImage {
	var returns *DataplexTaskSparkInfrastructureSpecContainerImage
	_jsii_.Get(
		j,
		"containerImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) InternalValue() *DataplexTaskSparkInfrastructureSpec {
	var returns *DataplexTaskSparkInfrastructureSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) VpcNetwork() DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference {
	var returns DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference
	_jsii_.Get(
		j,
		"vpcNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) VpcNetworkInput() *DataplexTaskSparkInfrastructureSpecVpcNetwork {
	var returns *DataplexTaskSparkInfrastructureSpecVpcNetwork
	_jsii_.Get(
		j,
		"vpcNetworkInput",
		&returns,
	)
	return returns
}


func NewDataplexTaskSparkInfrastructureSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataplexTaskSparkInfrastructureSpecOutputReference {
	_init_.Initialize()

	if err := validateNewDataplexTaskSparkInfrastructureSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataplexTask.DataplexTaskSparkInfrastructureSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataplexTaskSparkInfrastructureSpecOutputReference_Override(d DataplexTaskSparkInfrastructureSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataplexTask.DataplexTaskSparkInfrastructureSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference)SetInternalValue(val *DataplexTaskSparkInfrastructureSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) PutBatch(value *DataplexTaskSparkInfrastructureSpecBatch) {
	if err := d.validatePutBatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBatch",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) PutContainerImage(value *DataplexTaskSparkInfrastructureSpecContainerImage) {
	if err := d.validatePutContainerImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putContainerImage",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) PutVpcNetwork(value *DataplexTaskSparkInfrastructureSpecVpcNetwork) {
	if err := d.validatePutVpcNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVpcNetwork",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) ResetBatch() {
	_jsii_.InvokeVoid(
		d,
		"resetBatch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) ResetContainerImage() {
	_jsii_.InvokeVoid(
		d,
		"resetContainerImage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) ResetVpcNetwork() {
	_jsii_.InvokeVoid(
		d,
		"resetVpcNetwork",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

