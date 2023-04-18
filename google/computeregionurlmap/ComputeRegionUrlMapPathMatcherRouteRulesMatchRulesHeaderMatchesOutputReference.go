package computeregionurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v7/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v7/computeregionurlmap/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference interface {
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
	ExactMatch() *string
	SetExactMatch(val *string)
	ExactMatchInput() *string
	// Experimental.
	Fqn() *string
	HeaderName() *string
	SetHeaderName(val *string)
	HeaderNameInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	InvertMatch() interface{}
	SetInvertMatch(val interface{})
	InvertMatchInput() interface{}
	PrefixMatch() *string
	SetPrefixMatch(val *string)
	PrefixMatchInput() *string
	PresentMatch() interface{}
	SetPresentMatch(val interface{})
	PresentMatchInput() interface{}
	RangeMatch() ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference
	RangeMatchInput() *ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatch
	RegexMatch() *string
	SetRegexMatch(val *string)
	RegexMatchInput() *string
	SuffixMatch() *string
	SetSuffixMatch(val *string)
	SuffixMatchInput() *string
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
	PutRangeMatch(value *ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatch)
	ResetExactMatch()
	ResetInvertMatch()
	ResetPrefixMatch()
	ResetPresentMatch()
	ResetRangeMatch()
	ResetRegexMatch()
	ResetSuffixMatch()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference
type jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) ExactMatch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exactMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) ExactMatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exactMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) HeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) HeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) InvertMatch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invertMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) InvertMatchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invertMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) PrefixMatch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) PrefixMatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) PresentMatch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"presentMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) PresentMatchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"presentMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) RangeMatch() ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference {
	var returns ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference
	_jsii_.Get(
		j,
		"rangeMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) RangeMatchInput() *ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatch {
	var returns *ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatch
	_jsii_.Get(
		j,
		"rangeMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) RegexMatch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regexMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) RegexMatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regexMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) SuffixMatch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suffixMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) SuffixMatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suffixMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference {
	_init_.Initialize()

	if err := validateNewComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference_Override(c ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference)SetExactMatch(val *string) {
	if err := j.validateSetExactMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exactMatch",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference)SetHeaderName(val *string) {
	if err := j.validateSetHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headerName",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference)SetInvertMatch(val interface{}) {
	if err := j.validateSetInvertMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"invertMatch",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference)SetPrefixMatch(val *string) {
	if err := j.validateSetPrefixMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefixMatch",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference)SetPresentMatch(val interface{}) {
	if err := j.validateSetPresentMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"presentMatch",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference)SetRegexMatch(val *string) {
	if err := j.validateSetRegexMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regexMatch",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference)SetSuffixMatch(val *string) {
	if err := j.validateSetSuffixMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suffixMatch",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) PutRangeMatch(value *ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatch) {
	if err := c.validatePutRangeMatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRangeMatch",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) ResetExactMatch() {
	_jsii_.InvokeVoid(
		c,
		"resetExactMatch",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) ResetInvertMatch() {
	_jsii_.InvokeVoid(
		c,
		"resetInvertMatch",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) ResetPrefixMatch() {
	_jsii_.InvokeVoid(
		c,
		"resetPrefixMatch",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) ResetPresentMatch() {
	_jsii_.InvokeVoid(
		c,
		"resetPresentMatch",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) ResetRangeMatch() {
	_jsii_.InvokeVoid(
		c,
		"resetRangeMatch",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) ResetRegexMatch() {
	_jsii_.InvokeVoid(
		c,
		"resetRegexMatch",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) ResetSuffixMatch() {
	_jsii_.InvokeVoid(
		c,
		"resetSuffixMatch",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

