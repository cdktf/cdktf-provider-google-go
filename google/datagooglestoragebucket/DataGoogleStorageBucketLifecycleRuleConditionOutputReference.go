package datagooglestoragebucket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v5/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v5/datagooglestoragebucket/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleStorageBucketLifecycleRuleConditionOutputReference interface {
	cdktf.ComplexObject
	Age() *float64
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
	CreatedBefore() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomTimeBefore() *string
	DaysSinceCustomTime() *float64
	DaysSinceNoncurrentTime() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *DataGoogleStorageBucketLifecycleRuleCondition
	SetInternalValue(val *DataGoogleStorageBucketLifecycleRuleCondition)
	MatchesPrefix() *[]*string
	MatchesStorageClass() *[]*string
	MatchesSuffix() *[]*string
	NoncurrentTimeBefore() *string
	NumNewerVersions() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WithState() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGoogleStorageBucketLifecycleRuleConditionOutputReference
type jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) Age() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"age",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) CreatedBefore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) CustomTimeBefore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customTimeBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) DaysSinceCustomTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daysSinceCustomTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) DaysSinceNoncurrentTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daysSinceNoncurrentTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) InternalValue() *DataGoogleStorageBucketLifecycleRuleCondition {
	var returns *DataGoogleStorageBucketLifecycleRuleCondition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) MatchesPrefix() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchesPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) MatchesStorageClass() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchesStorageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) MatchesSuffix() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchesSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) NoncurrentTimeBefore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noncurrentTimeBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) NumNewerVersions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numNewerVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) WithState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"withState",
		&returns,
	)
	return returns
}


func NewDataGoogleStorageBucketLifecycleRuleConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleStorageBucketLifecycleRuleConditionOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleStorageBucketLifecycleRuleConditionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleStorageBucket.DataGoogleStorageBucketLifecycleRuleConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleStorageBucketLifecycleRuleConditionOutputReference_Override(d DataGoogleStorageBucketLifecycleRuleConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleStorageBucket.DataGoogleStorageBucketLifecycleRuleConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference)SetInternalValue(val *DataGoogleStorageBucketLifecycleRuleCondition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleStorageBucketLifecycleRuleConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

