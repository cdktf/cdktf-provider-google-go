// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigospolicyassignment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/osconfigospolicyassignment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference interface {
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
	InternalValue() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote
	SetInternalValue(val *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote)
	Sha256Checksum() *string
	SetSha256Checksum(val *string)
	Sha256ChecksumInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uri() *string
	SetUri(val *string)
	UriInput() *string
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
	ResetSha256Checksum()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference
type jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference) InternalValue() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote {
	var returns *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference) Sha256Checksum() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sha256Checksum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference) Sha256ChecksumInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sha256ChecksumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference) UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}


func NewOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference {
	_init_.Initialize()

	if err := validateNewOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.osConfigOsPolicyAssignment.OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference_Override(o OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.osConfigOsPolicyAssignment.OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference)SetInternalValue(val *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemote) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference)SetSha256Checksum(val *string) {
	if err := j.validateSetSha256ChecksumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sha256Checksum",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference)SetUri(val *string) {
	if err := j.validateSetUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference) ResetSha256Checksum() {
	_jsii_.InvokeVoid(
		o,
		"resetSha256Checksum",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

