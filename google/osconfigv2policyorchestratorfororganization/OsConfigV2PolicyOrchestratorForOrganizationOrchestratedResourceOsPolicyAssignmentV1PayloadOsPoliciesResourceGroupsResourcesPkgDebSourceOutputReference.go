// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestratorfororganization

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/osconfigv2policyorchestratorfororganization/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference interface {
	cdktf.ComplexObject
	AllowInsecure() interface{}
	SetAllowInsecure(val interface{})
	AllowInsecureInput() interface{}
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
	Gcs() OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceGcsOutputReference
	GcsInput() *OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceGcs
	InternalValue() *OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSource
	SetInternalValue(val *OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSource)
	LocalPath() *string
	SetLocalPath(val *string)
	LocalPathInput() *string
	Remote() OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference
	RemoteInput() *OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceRemote
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
	PutGcs(value *OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceGcs)
	PutRemote(value *OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceRemote)
	ResetAllowInsecure()
	ResetGcs()
	ResetLocalPath()
	ResetRemote()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference
type jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) AllowInsecure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInsecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) AllowInsecureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInsecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) Gcs() OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceGcsOutputReference {
	var returns OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceGcsOutputReference
	_jsii_.Get(
		j,
		"gcs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) GcsInput() *OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceGcs {
	var returns *OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceGcs
	_jsii_.Get(
		j,
		"gcsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) InternalValue() *OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSource {
	var returns *OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) LocalPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) LocalPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) Remote() OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference {
	var returns OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceRemoteOutputReference
	_jsii_.Get(
		j,
		"remote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) RemoteInput() *OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceRemote {
	var returns *OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceRemote
	_jsii_.Get(
		j,
		"remoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference {
	_init_.Initialize()

	if err := validateNewOsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.osConfigV2PolicyOrchestratorForOrganization.OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference_Override(o OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.osConfigV2PolicyOrchestratorForOrganization.OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference)SetAllowInsecure(val interface{}) {
	if err := j.validateSetAllowInsecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowInsecure",
		val,
	)
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference)SetInternalValue(val *OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference)SetLocalPath(val *string) {
	if err := j.validateSetLocalPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localPath",
		val,
	)
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) PutGcs(value *OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceGcs) {
	if err := o.validatePutGcsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putGcs",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) PutRemote(value *OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceRemote) {
	if err := o.validatePutRemoteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRemote",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) ResetAllowInsecure() {
	_jsii_.InvokeVoid(
		o,
		"resetAllowInsecure",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) ResetGcs() {
	_jsii_.InvokeVoid(
		o,
		"resetGcs",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) ResetLocalPath() {
	_jsii_.InvokeVoid(
		o,
		"resetLocalPath",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) ResetRemote() {
	_jsii_.InvokeVoid(
		o,
		"resetRemote",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

