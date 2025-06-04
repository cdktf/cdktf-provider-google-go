// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnsrecordset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/dnsrecordset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference interface {
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
	ExternalEndpoints() *[]*string
	SetExternalEndpoints(val *[]*string)
	ExternalEndpointsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalLoadBalancers() DnsRecordSetRoutingPolicyPrimaryBackupPrimaryInternalLoadBalancersList
	InternalLoadBalancersInput() interface{}
	InternalValue() *DnsRecordSetRoutingPolicyPrimaryBackupPrimary
	SetInternalValue(val *DnsRecordSetRoutingPolicyPrimaryBackupPrimary)
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
	PutInternalLoadBalancers(value interface{})
	ResetExternalEndpoints()
	ResetInternalLoadBalancers()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference
type jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) ExternalEndpoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) ExternalEndpointsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalEndpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) InternalLoadBalancers() DnsRecordSetRoutingPolicyPrimaryBackupPrimaryInternalLoadBalancersList {
	var returns DnsRecordSetRoutingPolicyPrimaryBackupPrimaryInternalLoadBalancersList
	_jsii_.Get(
		j,
		"internalLoadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) InternalLoadBalancersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalLoadBalancersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) InternalValue() *DnsRecordSetRoutingPolicyPrimaryBackupPrimary {
	var returns *DnsRecordSetRoutingPolicyPrimaryBackupPrimary
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference {
	_init_.Initialize()

	if err := validateNewDnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dnsRecordSet.DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference_Override(d DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dnsRecordSet.DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference)SetExternalEndpoints(val *[]*string) {
	if err := j.validateSetExternalEndpointsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalEndpoints",
		val,
	)
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference)SetInternalValue(val *DnsRecordSetRoutingPolicyPrimaryBackupPrimary) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) PutInternalLoadBalancers(value interface{}) {
	if err := d.validatePutInternalLoadBalancersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInternalLoadBalancers",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) ResetExternalEndpoints() {
	_jsii_.InvokeVoid(
		d,
		"resetExternalEndpoints",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) ResetInternalLoadBalancers() {
	_jsii_.InvokeVoid(
		d,
		"resetInternalLoadBalancers",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DnsRecordSetRoutingPolicyPrimaryBackupPrimaryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

