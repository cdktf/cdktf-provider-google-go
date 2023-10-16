// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package composerenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v11/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v11/composerenvironment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfig
	SetInternalValue(val *ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfig)
	SnapshotCreationSchedule() *string
	SetSnapshotCreationSchedule(val *string)
	SnapshotCreationScheduleInput() *string
	SnapshotLocation() *string
	SetSnapshotLocation(val *string)
	SnapshotLocationInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
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
	ResetSnapshotCreationSchedule()
	ResetSnapshotLocation()
	ResetTimeZone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference
type jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) InternalValue() *ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfig {
	var returns *ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) SnapshotCreationSchedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotCreationSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) SnapshotCreationScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotCreationScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) SnapshotLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) SnapshotLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}


func NewComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.composerEnvironment.ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference_Override(c ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.composerEnvironment.ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference)SetInternalValue(val *ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference)SetSnapshotCreationSchedule(val *string) {
	if err := j.validateSetSnapshotCreationScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotCreationSchedule",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference)SetSnapshotLocation(val *string) {
	if err := j.validateSetSnapshotLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotLocation",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) ResetSnapshotCreationSchedule() {
	_jsii_.InvokeVoid(
		c,
		"resetSnapshotCreationSchedule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) ResetSnapshotLocation() {
	_jsii_.InvokeVoid(
		c,
		"resetSnapshotLocation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) ResetTimeZone() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComposerEnvironmentConfigRecoveryConfigScheduledSnapshotsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

