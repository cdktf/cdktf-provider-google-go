// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package composerenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v11/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v11/composerenvironment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference interface {
	cdktf.ComplexObject
	ClusterIpv4CidrBlock() *string
	SetClusterIpv4CidrBlock(val *string)
	ClusterIpv4CidrBlockInput() *string
	ClusterSecondaryRangeName() *string
	SetClusterSecondaryRangeName(val *string)
	ClusterSecondaryRangeNameInput() *string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ServicesIpv4CidrBlock() *string
	SetServicesIpv4CidrBlock(val *string)
	ServicesIpv4CidrBlockInput() *string
	ServicesSecondaryRangeName() *string
	SetServicesSecondaryRangeName(val *string)
	ServicesSecondaryRangeNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseIpAliases() interface{}
	SetUseIpAliases(val interface{})
	UseIpAliasesInput() interface{}
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
	ResetClusterIpv4CidrBlock()
	ResetClusterSecondaryRangeName()
	ResetServicesIpv4CidrBlock()
	ResetServicesSecondaryRangeName()
	ResetUseIpAliases()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference
type jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) ClusterIpv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIpv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) ClusterIpv4CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIpv4CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) ClusterSecondaryRangeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSecondaryRangeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) ClusterSecondaryRangeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSecondaryRangeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) ServicesIpv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicesIpv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) ServicesIpv4CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicesIpv4CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) ServicesSecondaryRangeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicesSecondaryRangeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) ServicesSecondaryRangeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicesSecondaryRangeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) UseIpAliases() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useIpAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) UseIpAliasesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useIpAliasesInput",
		&returns,
	)
	return returns
}


func NewComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.composerEnvironment.ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference_Override(c ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.composerEnvironment.ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference)SetClusterIpv4CidrBlock(val *string) {
	if err := j.validateSetClusterIpv4CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterIpv4CidrBlock",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference)SetClusterSecondaryRangeName(val *string) {
	if err := j.validateSetClusterSecondaryRangeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterSecondaryRangeName",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference)SetServicesIpv4CidrBlock(val *string) {
	if err := j.validateSetServicesIpv4CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicesIpv4CidrBlock",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference)SetServicesSecondaryRangeName(val *string) {
	if err := j.validateSetServicesSecondaryRangeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicesSecondaryRangeName",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference)SetUseIpAliases(val interface{}) {
	if err := j.validateSetUseIpAliasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useIpAliases",
		val,
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) ResetClusterIpv4CidrBlock() {
	_jsii_.InvokeVoid(
		c,
		"resetClusterIpv4CidrBlock",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) ResetClusterSecondaryRangeName() {
	_jsii_.InvokeVoid(
		c,
		"resetClusterSecondaryRangeName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) ResetServicesIpv4CidrBlock() {
	_jsii_.InvokeVoid(
		c,
		"resetServicesIpv4CidrBlock",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) ResetServicesSecondaryRangeName() {
	_jsii_.InvokeVoid(
		c,
		"resetServicesSecondaryRangeName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) ResetUseIpAliases() {
	_jsii_.InvokeVoid(
		c,
		"resetUseIpAliases",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

