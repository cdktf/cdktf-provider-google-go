// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package colabschedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/colabschedule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference interface {
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
	DataformRepositorySource() ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference
	DataformRepositorySourceInput() *ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	ExecutionTimeout() *string
	SetExecutionTimeout(val *string)
	ExecutionTimeoutInput() *string
	ExecutionUser() *string
	SetExecutionUser(val *string)
	ExecutionUserInput() *string
	// Experimental.
	Fqn() *string
	GcsNotebookSource() ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSourceOutputReference
	GcsNotebookSourceInput() *ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource
	GcsOutputUri() *string
	SetGcsOutputUri(val *string)
	GcsOutputUriInput() *string
	InternalValue() *ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob
	SetInternalValue(val *ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob)
	NotebookRuntimeTemplateResourceName() *string
	SetNotebookRuntimeTemplateResourceName(val *string)
	NotebookRuntimeTemplateResourceNameInput() *string
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
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
	PutDataformRepositorySource(value *ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource)
	PutGcsNotebookSource(value *ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource)
	ResetDataformRepositorySource()
	ResetExecutionTimeout()
	ResetExecutionUser()
	ResetGcsNotebookSource()
	ResetServiceAccount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference
type jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) DataformRepositorySource() ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference {
	var returns ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySourceOutputReference
	_jsii_.Get(
		j,
		"dataformRepositorySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) DataformRepositorySourceInput() *ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource {
	var returns *ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource
	_jsii_.Get(
		j,
		"dataformRepositorySourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ExecutionTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ExecutionTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ExecutionUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ExecutionUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GcsNotebookSource() ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSourceOutputReference {
	var returns ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSourceOutputReference
	_jsii_.Get(
		j,
		"gcsNotebookSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GcsNotebookSourceInput() *ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource {
	var returns *ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource
	_jsii_.Get(
		j,
		"gcsNotebookSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GcsOutputUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsOutputUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GcsOutputUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsOutputUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) InternalValue() *ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob {
	var returns *ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) NotebookRuntimeTemplateResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebookRuntimeTemplateResourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) NotebookRuntimeTemplateResourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebookRuntimeTemplateResourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference {
	_init_.Initialize()

	if err := validateNewColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.colabSchedule.ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference_Override(c ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.colabSchedule.ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetExecutionTimeout(val *string) {
	if err := j.validateSetExecutionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionTimeout",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetExecutionUser(val *string) {
	if err := j.validateSetExecutionUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionUser",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetGcsOutputUri(val *string) {
	if err := j.validateSetGcsOutputUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcsOutputUri",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetInternalValue(val *ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetNotebookRuntimeTemplateResourceName(val *string) {
	if err := j.validateSetNotebookRuntimeTemplateResourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notebookRuntimeTemplateResourceName",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) PutDataformRepositorySource(value *ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource) {
	if err := c.validatePutDataformRepositorySourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDataformRepositorySource",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) PutGcsNotebookSource(value *ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource) {
	if err := c.validatePutGcsNotebookSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGcsNotebookSource",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetDataformRepositorySource() {
	_jsii_.InvokeVoid(
		c,
		"resetDataformRepositorySource",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetExecutionTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetExecutionTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetExecutionUser() {
	_jsii_.InvokeVoid(
		c,
		"resetExecutionUser",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetGcsNotebookSource() {
	_jsii_.InvokeVoid(
		c,
		"resetGcsNotebookSource",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

