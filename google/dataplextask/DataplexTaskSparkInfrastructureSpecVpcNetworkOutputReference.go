// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplextask

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/dataplextask/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference interface {
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
	InternalValue() *DataplexTaskSparkInfrastructureSpecVpcNetwork
	SetInternalValue(val *DataplexTaskSparkInfrastructureSpecVpcNetwork)
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	NetworkTags() *[]*string
	SetNetworkTags(val *[]*string)
	NetworkTagsInput() *[]*string
	SubNetwork() *string
	SetSubNetwork(val *string)
	SubNetworkInput() *string
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
	ResetNetwork()
	ResetNetworkTags()
	ResetSubNetwork()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference
type jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) InternalValue() *DataplexTaskSparkInfrastructureSpecVpcNetwork {
	var returns *DataplexTaskSparkInfrastructureSpecVpcNetwork
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) NetworkTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) NetworkTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) SubNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) SubNetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference {
	_init_.Initialize()

	if err := validateNewDataplexTaskSparkInfrastructureSpecVpcNetworkOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataplexTask.DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference_Override(d DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataplexTask.DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference)SetInternalValue(val *DataplexTaskSparkInfrastructureSpecVpcNetwork) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference)SetNetworkTags(val *[]*string) {
	if err := j.validateSetNetworkTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkTags",
		val,
	)
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference)SetSubNetwork(val *string) {
	if err := j.validateSetSubNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subNetwork",
		val,
	)
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) ResetNetwork() {
	_jsii_.InvokeVoid(
		d,
		"resetNetwork",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) ResetNetworkTags() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) ResetSubNetwork() {
	_jsii_.InvokeVoid(
		d,
		"resetSubNetwork",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataplexTaskSparkInfrastructureSpecVpcNetworkOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

