// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstancetemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/computeinstancetemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInstanceTemplateSchedulingOutputReference interface {
	cdktf.ComplexObject
	AutomaticRestart() interface{}
	SetAutomaticRestart(val interface{})
	AutomaticRestartInput() interface{}
	AvailabilityDomain() *float64
	SetAvailabilityDomain(val *float64)
	AvailabilityDomainInput() *float64
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
	InstanceTerminationAction() *string
	SetInstanceTerminationAction(val *string)
	InstanceTerminationActionInput() *string
	InternalValue() *ComputeInstanceTemplateScheduling
	SetInternalValue(val *ComputeInstanceTemplateScheduling)
	LocalSsdRecoveryTimeout() ComputeInstanceTemplateSchedulingLocalSsdRecoveryTimeoutList
	LocalSsdRecoveryTimeoutInput() interface{}
	MaxRunDuration() ComputeInstanceTemplateSchedulingMaxRunDurationOutputReference
	MaxRunDurationInput() *ComputeInstanceTemplateSchedulingMaxRunDuration
	MinNodeCpus() *float64
	SetMinNodeCpus(val *float64)
	MinNodeCpusInput() *float64
	NodeAffinities() ComputeInstanceTemplateSchedulingNodeAffinitiesList
	NodeAffinitiesInput() interface{}
	OnHostMaintenance() *string
	SetOnHostMaintenance(val *string)
	OnHostMaintenanceInput() *string
	OnInstanceStopAction() ComputeInstanceTemplateSchedulingOnInstanceStopActionOutputReference
	OnInstanceStopActionInput() *ComputeInstanceTemplateSchedulingOnInstanceStopAction
	Preemptible() interface{}
	SetPreemptible(val interface{})
	PreemptibleInput() interface{}
	ProvisioningModel() *string
	SetProvisioningModel(val *string)
	ProvisioningModelInput() *string
	TerminationTime() *string
	SetTerminationTime(val *string)
	TerminationTimeInput() *string
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
	PutLocalSsdRecoveryTimeout(value interface{})
	PutMaxRunDuration(value *ComputeInstanceTemplateSchedulingMaxRunDuration)
	PutNodeAffinities(value interface{})
	PutOnInstanceStopAction(value *ComputeInstanceTemplateSchedulingOnInstanceStopAction)
	ResetAutomaticRestart()
	ResetAvailabilityDomain()
	ResetInstanceTerminationAction()
	ResetLocalSsdRecoveryTimeout()
	ResetMaxRunDuration()
	ResetMinNodeCpus()
	ResetNodeAffinities()
	ResetOnHostMaintenance()
	ResetOnInstanceStopAction()
	ResetPreemptible()
	ResetProvisioningModel()
	ResetTerminationTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeInstanceTemplateSchedulingOutputReference
type jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) AutomaticRestart() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticRestart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) AutomaticRestartInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticRestartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) AvailabilityDomain() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availabilityDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) AvailabilityDomainInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availabilityDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) InstanceTerminationAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTerminationAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) InstanceTerminationActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTerminationActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) InternalValue() *ComputeInstanceTemplateScheduling {
	var returns *ComputeInstanceTemplateScheduling
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) LocalSsdRecoveryTimeout() ComputeInstanceTemplateSchedulingLocalSsdRecoveryTimeoutList {
	var returns ComputeInstanceTemplateSchedulingLocalSsdRecoveryTimeoutList
	_jsii_.Get(
		j,
		"localSsdRecoveryTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) LocalSsdRecoveryTimeoutInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localSsdRecoveryTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) MaxRunDuration() ComputeInstanceTemplateSchedulingMaxRunDurationOutputReference {
	var returns ComputeInstanceTemplateSchedulingMaxRunDurationOutputReference
	_jsii_.Get(
		j,
		"maxRunDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) MaxRunDurationInput() *ComputeInstanceTemplateSchedulingMaxRunDuration {
	var returns *ComputeInstanceTemplateSchedulingMaxRunDuration
	_jsii_.Get(
		j,
		"maxRunDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) MinNodeCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodeCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) MinNodeCpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodeCpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) NodeAffinities() ComputeInstanceTemplateSchedulingNodeAffinitiesList {
	var returns ComputeInstanceTemplateSchedulingNodeAffinitiesList
	_jsii_.Get(
		j,
		"nodeAffinities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) NodeAffinitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodeAffinitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) OnHostMaintenance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onHostMaintenance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) OnHostMaintenanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onHostMaintenanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) OnInstanceStopAction() ComputeInstanceTemplateSchedulingOnInstanceStopActionOutputReference {
	var returns ComputeInstanceTemplateSchedulingOnInstanceStopActionOutputReference
	_jsii_.Get(
		j,
		"onInstanceStopAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) OnInstanceStopActionInput() *ComputeInstanceTemplateSchedulingOnInstanceStopAction {
	var returns *ComputeInstanceTemplateSchedulingOnInstanceStopAction
	_jsii_.Get(
		j,
		"onInstanceStopActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) Preemptible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preemptible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) PreemptibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preemptibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) ProvisioningModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) ProvisioningModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) TerminationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) TerminationTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeInstanceTemplateSchedulingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeInstanceTemplateSchedulingOutputReference {
	_init_.Initialize()

	if err := validateNewComputeInstanceTemplateSchedulingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeInstanceTemplate.ComputeInstanceTemplateSchedulingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeInstanceTemplateSchedulingOutputReference_Override(c ComputeInstanceTemplateSchedulingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeInstanceTemplate.ComputeInstanceTemplateSchedulingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference)SetAutomaticRestart(val interface{}) {
	if err := j.validateSetAutomaticRestartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticRestart",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference)SetAvailabilityDomain(val *float64) {
	if err := j.validateSetAvailabilityDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityDomain",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference)SetInstanceTerminationAction(val *string) {
	if err := j.validateSetInstanceTerminationActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTerminationAction",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference)SetInternalValue(val *ComputeInstanceTemplateScheduling) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference)SetMinNodeCpus(val *float64) {
	if err := j.validateSetMinNodeCpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minNodeCpus",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference)SetOnHostMaintenance(val *string) {
	if err := j.validateSetOnHostMaintenanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onHostMaintenance",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference)SetPreemptible(val interface{}) {
	if err := j.validateSetPreemptibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preemptible",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference)SetProvisioningModel(val *string) {
	if err := j.validateSetProvisioningModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioningModel",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference)SetTerminationTime(val *string) {
	if err := j.validateSetTerminationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationTime",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) PutLocalSsdRecoveryTimeout(value interface{}) {
	if err := c.validatePutLocalSsdRecoveryTimeoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLocalSsdRecoveryTimeout",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) PutMaxRunDuration(value *ComputeInstanceTemplateSchedulingMaxRunDuration) {
	if err := c.validatePutMaxRunDurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMaxRunDuration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) PutNodeAffinities(value interface{}) {
	if err := c.validatePutNodeAffinitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNodeAffinities",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) PutOnInstanceStopAction(value *ComputeInstanceTemplateSchedulingOnInstanceStopAction) {
	if err := c.validatePutOnInstanceStopActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOnInstanceStopAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) ResetAutomaticRestart() {
	_jsii_.InvokeVoid(
		c,
		"resetAutomaticRestart",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) ResetAvailabilityDomain() {
	_jsii_.InvokeVoid(
		c,
		"resetAvailabilityDomain",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) ResetInstanceTerminationAction() {
	_jsii_.InvokeVoid(
		c,
		"resetInstanceTerminationAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) ResetLocalSsdRecoveryTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetLocalSsdRecoveryTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) ResetMaxRunDuration() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxRunDuration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) ResetMinNodeCpus() {
	_jsii_.InvokeVoid(
		c,
		"resetMinNodeCpus",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) ResetNodeAffinities() {
	_jsii_.InvokeVoid(
		c,
		"resetNodeAffinities",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) ResetOnHostMaintenance() {
	_jsii_.InvokeVoid(
		c,
		"resetOnHostMaintenance",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) ResetOnInstanceStopAction() {
	_jsii_.InvokeVoid(
		c,
		"resetOnInstanceStopAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) ResetPreemptible() {
	_jsii_.InvokeVoid(
		c,
		"resetPreemptible",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) ResetProvisioningModel() {
	_jsii_.InvokeVoid(
		c,
		"resetProvisioningModel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) ResetTerminationTime() {
	_jsii_.InvokeVoid(
		c,
		"resetTerminationTime",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeInstanceTemplateSchedulingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

