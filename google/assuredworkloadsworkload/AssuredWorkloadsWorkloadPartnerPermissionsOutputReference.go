// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package assuredworkloadsworkload

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/assuredworkloadsworkload/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AssuredWorkloadsWorkloadPartnerPermissionsOutputReference interface {
	cdktf.ComplexObject
	AssuredWorkloadsMonitoring() interface{}
	SetAssuredWorkloadsMonitoring(val interface{})
	AssuredWorkloadsMonitoringInput() interface{}
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
	DataLogsViewer() interface{}
	SetDataLogsViewer(val interface{})
	DataLogsViewerInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *AssuredWorkloadsWorkloadPartnerPermissions
	SetInternalValue(val *AssuredWorkloadsWorkloadPartnerPermissions)
	ServiceAccessApprover() interface{}
	SetServiceAccessApprover(val interface{})
	ServiceAccessApproverInput() interface{}
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetAssuredWorkloadsMonitoring()
	ResetDataLogsViewer()
	ResetServiceAccessApprover()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AssuredWorkloadsWorkloadPartnerPermissionsOutputReference
type jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) AssuredWorkloadsMonitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assuredWorkloadsMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) AssuredWorkloadsMonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assuredWorkloadsMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) DataLogsViewer() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataLogsViewer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) DataLogsViewerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataLogsViewerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) InternalValue() *AssuredWorkloadsWorkloadPartnerPermissions {
	var returns *AssuredWorkloadsWorkloadPartnerPermissions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) ServiceAccessApprover() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceAccessApprover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) ServiceAccessApproverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceAccessApproverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAssuredWorkloadsWorkloadPartnerPermissionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AssuredWorkloadsWorkloadPartnerPermissionsOutputReference {
	_init_.Initialize()

	if err := validateNewAssuredWorkloadsWorkloadPartnerPermissionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.assuredWorkloadsWorkload.AssuredWorkloadsWorkloadPartnerPermissionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAssuredWorkloadsWorkloadPartnerPermissionsOutputReference_Override(a AssuredWorkloadsWorkloadPartnerPermissionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.assuredWorkloadsWorkload.AssuredWorkloadsWorkloadPartnerPermissionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference)SetAssuredWorkloadsMonitoring(val interface{}) {
	if err := j.validateSetAssuredWorkloadsMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assuredWorkloadsMonitoring",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference)SetDataLogsViewer(val interface{}) {
	if err := j.validateSetDataLogsViewerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataLogsViewer",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference)SetInternalValue(val *AssuredWorkloadsWorkloadPartnerPermissions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference)SetServiceAccessApprover(val interface{}) {
	if err := j.validateSetServiceAccessApproverParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccessApprover",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) ResetAssuredWorkloadsMonitoring() {
	_jsii_.InvokeVoid(
		a,
		"resetAssuredWorkloadsMonitoring",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) ResetDataLogsViewer() {
	_jsii_.InvokeVoid(
		a,
		"resetDataLogsViewer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) ResetServiceAccessApprover() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceAccessApprover",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssuredWorkloadsWorkloadPartnerPermissionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

