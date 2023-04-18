package computerouter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v7/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v7/computerouter/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRouterBgpOutputReference interface {
	cdktf.ComplexObject
	AdvertisedGroups() *[]*string
	SetAdvertisedGroups(val *[]*string)
	AdvertisedGroupsInput() *[]*string
	AdvertisedIpRanges() ComputeRouterBgpAdvertisedIpRangesList
	AdvertisedIpRangesInput() interface{}
	AdvertiseMode() *string
	SetAdvertiseMode(val *string)
	AdvertiseModeInput() *string
	Asn() *float64
	SetAsn(val *float64)
	AsnInput() *float64
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
	InternalValue() *ComputeRouterBgp
	SetInternalValue(val *ComputeRouterBgp)
	KeepaliveInterval() *float64
	SetKeepaliveInterval(val *float64)
	KeepaliveIntervalInput() *float64
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
	PutAdvertisedIpRanges(value interface{})
	ResetAdvertisedGroups()
	ResetAdvertisedIpRanges()
	ResetAdvertiseMode()
	ResetKeepaliveInterval()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeRouterBgpOutputReference
type jsiiProxy_ComputeRouterBgpOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeRouterBgpOutputReference) AdvertisedGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"advertisedGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterBgpOutputReference) AdvertisedGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"advertisedGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterBgpOutputReference) AdvertisedIpRanges() ComputeRouterBgpAdvertisedIpRangesList {
	var returns ComputeRouterBgpAdvertisedIpRangesList
	_jsii_.Get(
		j,
		"advertisedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterBgpOutputReference) AdvertisedIpRangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advertisedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterBgpOutputReference) AdvertiseMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"advertiseMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterBgpOutputReference) AdvertiseModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"advertiseModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterBgpOutputReference) Asn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"asn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterBgpOutputReference) AsnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"asnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterBgpOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterBgpOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterBgpOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterBgpOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterBgpOutputReference) InternalValue() *ComputeRouterBgp {
	var returns *ComputeRouterBgp
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterBgpOutputReference) KeepaliveInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepaliveInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterBgpOutputReference) KeepaliveIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keepaliveIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterBgpOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterBgpOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeRouterBgpOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeRouterBgpOutputReference {
	_init_.Initialize()

	if err := validateNewComputeRouterBgpOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRouterBgpOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRouter.ComputeRouterBgpOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeRouterBgpOutputReference_Override(c ComputeRouterBgpOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRouter.ComputeRouterBgpOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeRouterBgpOutputReference)SetAdvertisedGroups(val *[]*string) {
	if err := j.validateSetAdvertisedGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advertisedGroups",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterBgpOutputReference)SetAdvertiseMode(val *string) {
	if err := j.validateSetAdvertiseModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advertiseMode",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterBgpOutputReference)SetAsn(val *float64) {
	if err := j.validateSetAsnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"asn",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterBgpOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterBgpOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterBgpOutputReference)SetInternalValue(val *ComputeRouterBgp) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterBgpOutputReference)SetKeepaliveInterval(val *float64) {
	if err := j.validateSetKeepaliveIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepaliveInterval",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterBgpOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterBgpOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeRouterBgpOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRouterBgpOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeRouterBgpOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRouterBgpOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRouterBgpOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRouterBgpOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRouterBgpOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRouterBgpOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRouterBgpOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRouterBgpOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRouterBgpOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRouterBgpOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRouterBgpOutputReference) PutAdvertisedIpRanges(value interface{}) {
	if err := c.validatePutAdvertisedIpRangesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAdvertisedIpRanges",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRouterBgpOutputReference) ResetAdvertisedGroups() {
	_jsii_.InvokeVoid(
		c,
		"resetAdvertisedGroups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterBgpOutputReference) ResetAdvertisedIpRanges() {
	_jsii_.InvokeVoid(
		c,
		"resetAdvertisedIpRanges",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterBgpOutputReference) ResetAdvertiseMode() {
	_jsii_.InvokeVoid(
		c,
		"resetAdvertiseMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterBgpOutputReference) ResetKeepaliveInterval() {
	_jsii_.InvokeVoid(
		c,
		"resetKeepaliveInterval",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterBgpOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeRouterBgpOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

