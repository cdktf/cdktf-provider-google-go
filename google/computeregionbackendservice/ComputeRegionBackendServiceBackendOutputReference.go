package computeregionbackendservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v6/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v6/computeregionbackendservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionBackendServiceBackendOutputReference interface {
	cdktf.ComplexObject
	BalancingMode() *string
	SetBalancingMode(val *string)
	BalancingModeInput() *string
	CapacityScaler() *float64
	SetCapacityScaler(val *float64)
	CapacityScalerInput() *float64
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Failover() interface{}
	SetFailover(val interface{})
	FailoverInput() interface{}
	// Experimental.
	Fqn() *string
	Group() *string
	SetGroup(val *string)
	GroupInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxConnections() *float64
	SetMaxConnections(val *float64)
	MaxConnectionsInput() *float64
	MaxConnectionsPerEndpoint() *float64
	SetMaxConnectionsPerEndpoint(val *float64)
	MaxConnectionsPerEndpointInput() *float64
	MaxConnectionsPerInstance() *float64
	SetMaxConnectionsPerInstance(val *float64)
	MaxConnectionsPerInstanceInput() *float64
	MaxRate() *float64
	SetMaxRate(val *float64)
	MaxRateInput() *float64
	MaxRatePerEndpoint() *float64
	SetMaxRatePerEndpoint(val *float64)
	MaxRatePerEndpointInput() *float64
	MaxRatePerInstance() *float64
	SetMaxRatePerInstance(val *float64)
	MaxRatePerInstanceInput() *float64
	MaxUtilization() *float64
	SetMaxUtilization(val *float64)
	MaxUtilizationInput() *float64
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
	ResetBalancingMode()
	ResetCapacityScaler()
	ResetDescription()
	ResetFailover()
	ResetMaxConnections()
	ResetMaxConnectionsPerEndpoint()
	ResetMaxConnectionsPerInstance()
	ResetMaxRate()
	ResetMaxRatePerEndpoint()
	ResetMaxRatePerInstance()
	ResetMaxUtilization()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeRegionBackendServiceBackendOutputReference
type jsiiProxy_ComputeRegionBackendServiceBackendOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) BalancingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"balancingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) BalancingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"balancingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) CapacityScaler() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacityScaler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) CapacityScalerInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacityScalerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) Failover() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) FailoverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failoverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) Group() *string {
	var returns *string
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) GroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) MaxConnections() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) MaxConnectionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) MaxConnectionsPerEndpoint() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionsPerEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) MaxConnectionsPerEndpointInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionsPerEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) MaxConnectionsPerInstance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionsPerInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) MaxConnectionsPerInstanceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionsPerInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) MaxRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) MaxRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) MaxRatePerEndpoint() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRatePerEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) MaxRatePerEndpointInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRatePerEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) MaxRatePerInstance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRatePerInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) MaxRatePerInstanceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRatePerInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) MaxUtilization() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUtilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) MaxUtilizationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUtilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeRegionBackendServiceBackendOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeRegionBackendServiceBackendOutputReference {
	_init_.Initialize()

	if err := validateNewComputeRegionBackendServiceBackendOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionBackendServiceBackendOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionBackendService.ComputeRegionBackendServiceBackendOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeRegionBackendServiceBackendOutputReference_Override(c ComputeRegionBackendServiceBackendOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionBackendService.ComputeRegionBackendServiceBackendOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference)SetBalancingMode(val *string) {
	if err := j.validateSetBalancingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"balancingMode",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference)SetCapacityScaler(val *float64) {
	if err := j.validateSetCapacityScalerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityScaler",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference)SetFailover(val interface{}) {
	if err := j.validateSetFailoverParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failover",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference)SetGroup(val *string) {
	if err := j.validateSetGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"group",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference)SetMaxConnections(val *float64) {
	if err := j.validateSetMaxConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConnections",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference)SetMaxConnectionsPerEndpoint(val *float64) {
	if err := j.validateSetMaxConnectionsPerEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConnectionsPerEndpoint",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference)SetMaxConnectionsPerInstance(val *float64) {
	if err := j.validateSetMaxConnectionsPerInstanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConnectionsPerInstance",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference)SetMaxRate(val *float64) {
	if err := j.validateSetMaxRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRate",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference)SetMaxRatePerEndpoint(val *float64) {
	if err := j.validateSetMaxRatePerEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRatePerEndpoint",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference)SetMaxRatePerInstance(val *float64) {
	if err := j.validateSetMaxRatePerInstanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRatePerInstance",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference)SetMaxUtilization(val *float64) {
	if err := j.validateSetMaxUtilizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUtilization",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) ResetBalancingMode() {
	_jsii_.InvokeVoid(
		c,
		"resetBalancingMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) ResetCapacityScaler() {
	_jsii_.InvokeVoid(
		c,
		"resetCapacityScaler",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) ResetFailover() {
	_jsii_.InvokeVoid(
		c,
		"resetFailover",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) ResetMaxConnections() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxConnections",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) ResetMaxConnectionsPerEndpoint() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxConnectionsPerEndpoint",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) ResetMaxConnectionsPerInstance() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxConnectionsPerInstance",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) ResetMaxRate() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxRate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) ResetMaxRatePerEndpoint() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxRatePerEndpoint",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) ResetMaxRatePerInstance() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxRatePerInstance",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) ResetMaxUtilization() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxUtilization",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeRegionBackendServiceBackendOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

