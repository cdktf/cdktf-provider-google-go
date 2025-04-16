// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigv2policyorchestrator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/osconfigv2policyorchestrator/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference interface {
	cdktf.ComplexObject
	Apt() OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgAptOutputReference
	AptInput() *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgApt
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
	Deb() OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebOutputReference
	DebInput() *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDeb
	DesiredState() *string
	SetDesiredState(val *string)
	DesiredStateInput() *string
	// Experimental.
	Fqn() *string
	Googet() OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgGoogetOutputReference
	GoogetInput() *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgGooget
	InternalValue() *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkg
	SetInternalValue(val *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkg)
	Msi() OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgMsiOutputReference
	MsiInput() *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgMsi
	Rpm() OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgRpmOutputReference
	RpmInput() *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgRpm
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Yum() OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgYumOutputReference
	YumInput() *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgYum
	Zypper() OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgZypperOutputReference
	ZypperInput() *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgZypper
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
	PutApt(value *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgApt)
	PutDeb(value *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDeb)
	PutGooget(value *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgGooget)
	PutMsi(value *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgMsi)
	PutRpm(value *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgRpm)
	PutYum(value *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgYum)
	PutZypper(value *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgZypper)
	ResetApt()
	ResetDeb()
	ResetGooget()
	ResetMsi()
	ResetRpm()
	ResetYum()
	ResetZypper()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference
type jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) Apt() OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgAptOutputReference {
	var returns OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgAptOutputReference
	_jsii_.Get(
		j,
		"apt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) AptInput() *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgApt {
	var returns *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgApt
	_jsii_.Get(
		j,
		"aptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) Deb() OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebOutputReference {
	var returns OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDebOutputReference
	_jsii_.Get(
		j,
		"deb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) DebInput() *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDeb {
	var returns *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDeb
	_jsii_.Get(
		j,
		"debInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) DesiredState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) DesiredStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) Googet() OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgGoogetOutputReference {
	var returns OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgGoogetOutputReference
	_jsii_.Get(
		j,
		"googet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) GoogetInput() *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgGooget {
	var returns *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgGooget
	_jsii_.Get(
		j,
		"googetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) InternalValue() *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkg {
	var returns *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkg
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) Msi() OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgMsiOutputReference {
	var returns OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgMsiOutputReference
	_jsii_.Get(
		j,
		"msi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) MsiInput() *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgMsi {
	var returns *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgMsi
	_jsii_.Get(
		j,
		"msiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) Rpm() OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgRpmOutputReference {
	var returns OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgRpmOutputReference
	_jsii_.Get(
		j,
		"rpm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) RpmInput() *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgRpm {
	var returns *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgRpm
	_jsii_.Get(
		j,
		"rpmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) Yum() OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgYumOutputReference {
	var returns OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgYumOutputReference
	_jsii_.Get(
		j,
		"yum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) YumInput() *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgYum {
	var returns *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgYum
	_jsii_.Get(
		j,
		"yumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) Zypper() OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgZypperOutputReference {
	var returns OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgZypperOutputReference
	_jsii_.Get(
		j,
		"zypper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) ZypperInput() *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgZypper {
	var returns *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgZypper
	_jsii_.Get(
		j,
		"zypperInput",
		&returns,
	)
	return returns
}


func NewOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference {
	_init_.Initialize()

	if err := validateNewOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.osConfigV2PolicyOrchestrator.OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference_Override(o OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.osConfigV2PolicyOrchestrator.OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference)SetDesiredState(val *string) {
	if err := j.validateSetDesiredStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredState",
		val,
	)
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference)SetInternalValue(val *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkg) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) PutApt(value *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgApt) {
	if err := o.validatePutAptParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putApt",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) PutDeb(value *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgDeb) {
	if err := o.validatePutDebParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDeb",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) PutGooget(value *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgGooget) {
	if err := o.validatePutGoogetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putGooget",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) PutMsi(value *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgMsi) {
	if err := o.validatePutMsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putMsi",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) PutRpm(value *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgRpm) {
	if err := o.validatePutRpmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRpm",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) PutYum(value *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgYum) {
	if err := o.validatePutYumParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putYum",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) PutZypper(value *OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgZypper) {
	if err := o.validatePutZypperParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putZypper",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) ResetApt() {
	_jsii_.InvokeVoid(
		o,
		"resetApt",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) ResetDeb() {
	_jsii_.InvokeVoid(
		o,
		"resetDeb",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) ResetGooget() {
	_jsii_.InvokeVoid(
		o,
		"resetGooget",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) ResetMsi() {
	_jsii_.InvokeVoid(
		o,
		"resetMsi",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) ResetRpm() {
	_jsii_.InvokeVoid(
		o,
		"resetRpm",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) ResetYum() {
	_jsii_.InvokeVoid(
		o,
		"resetYum",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) ResetZypper() {
	_jsii_.InvokeVoid(
		o,
		"resetZypper",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesPkgOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

