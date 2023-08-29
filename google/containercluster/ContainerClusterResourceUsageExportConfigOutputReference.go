// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v9/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v9/containercluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerClusterResourceUsageExportConfigOutputReference interface {
	cdktf.ComplexObject
	BigqueryDestination() ContainerClusterResourceUsageExportConfigBigqueryDestinationOutputReference
	BigqueryDestinationInput() *ContainerClusterResourceUsageExportConfigBigqueryDestination
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
	EnableNetworkEgressMetering() interface{}
	SetEnableNetworkEgressMetering(val interface{})
	EnableNetworkEgressMeteringInput() interface{}
	EnableResourceConsumptionMetering() interface{}
	SetEnableResourceConsumptionMetering(val interface{})
	EnableResourceConsumptionMeteringInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ContainerClusterResourceUsageExportConfig
	SetInternalValue(val *ContainerClusterResourceUsageExportConfig)
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
	PutBigqueryDestination(value *ContainerClusterResourceUsageExportConfigBigqueryDestination)
	ResetEnableNetworkEgressMetering()
	ResetEnableResourceConsumptionMetering()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerClusterResourceUsageExportConfigOutputReference
type jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) BigqueryDestination() ContainerClusterResourceUsageExportConfigBigqueryDestinationOutputReference {
	var returns ContainerClusterResourceUsageExportConfigBigqueryDestinationOutputReference
	_jsii_.Get(
		j,
		"bigqueryDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) BigqueryDestinationInput() *ContainerClusterResourceUsageExportConfigBigqueryDestination {
	var returns *ContainerClusterResourceUsageExportConfigBigqueryDestination
	_jsii_.Get(
		j,
		"bigqueryDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) EnableNetworkEgressMetering() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNetworkEgressMetering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) EnableNetworkEgressMeteringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNetworkEgressMeteringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) EnableResourceConsumptionMetering() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableResourceConsumptionMetering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) EnableResourceConsumptionMeteringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableResourceConsumptionMeteringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) InternalValue() *ContainerClusterResourceUsageExportConfig {
	var returns *ContainerClusterResourceUsageExportConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerClusterResourceUsageExportConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerClusterResourceUsageExportConfigOutputReference {
	_init_.Initialize()

	if err := validateNewContainerClusterResourceUsageExportConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterResourceUsageExportConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerClusterResourceUsageExportConfigOutputReference_Override(c ContainerClusterResourceUsageExportConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterResourceUsageExportConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference)SetEnableNetworkEgressMetering(val interface{}) {
	if err := j.validateSetEnableNetworkEgressMeteringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableNetworkEgressMetering",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference)SetEnableResourceConsumptionMetering(val interface{}) {
	if err := j.validateSetEnableResourceConsumptionMeteringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourceConsumptionMetering",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference)SetInternalValue(val *ContainerClusterResourceUsageExportConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) PutBigqueryDestination(value *ContainerClusterResourceUsageExportConfigBigqueryDestination) {
	if err := c.validatePutBigqueryDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBigqueryDestination",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) ResetEnableNetworkEgressMetering() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableNetworkEgressMetering",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) ResetEnableResourceConsumptionMetering() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableResourceConsumptionMetering",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterResourceUsageExportConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

