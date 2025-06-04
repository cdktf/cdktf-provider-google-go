// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeaturemembership

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/gkehubfeaturemembership/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference interface {
	cdktf.ComplexObject
	Backends() *[]*string
	SetBackends(val *[]*string)
	BackendsInput() *[]*string
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
	InternalValue() *GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoring
	SetInternalValue(val *GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoring)
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
	ResetBackends()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference
type jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference) Backends() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"backends",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference) BackendsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"backendsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference) InternalValue() *GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoring {
	var returns *GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoring
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference {
	_init_.Initialize()

	if err := validateNewGkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeHubFeatureMembership.GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference_Override(g GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeHubFeatureMembership.GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference)SetBackends(val *[]*string) {
	if err := j.validateSetBackendsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backends",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference)SetInternalValue(val *GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoring) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference) ResetBackends() {
	_jsii_.InvokeVoid(
		g,
		"resetBackends",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoringOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

