package osconfigospolicyassignment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v4/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v4/osconfigospolicyassignment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference interface {
	cdktf.ComplexObject
	Apt() OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptOutputReference
	AptInput() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt
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
	Goo() OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGooOutputReference
	GooInput() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo
	InternalValue() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository
	SetInternalValue(val *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Yum() OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference
	YumInput() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum
	Zypper() OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypperOutputReference
	ZypperInput() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper
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
	PutApt(value *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt)
	PutGoo(value *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo)
	PutYum(value *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum)
	PutZypper(value *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper)
	ResetApt()
	ResetGoo()
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

// The jsii proxy struct for OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference
type jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) Apt() OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptOutputReference {
	var returns OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryAptOutputReference
	_jsii_.Get(
		j,
		"apt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) AptInput() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt {
	var returns *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt
	_jsii_.Get(
		j,
		"aptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) Goo() OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGooOutputReference {
	var returns OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGooOutputReference
	_jsii_.Get(
		j,
		"goo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) GooInput() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo {
	var returns *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo
	_jsii_.Get(
		j,
		"gooInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) InternalValue() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository {
	var returns *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) Yum() OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference {
	var returns OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYumOutputReference
	_jsii_.Get(
		j,
		"yum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) YumInput() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum {
	var returns *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum
	_jsii_.Get(
		j,
		"yumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) Zypper() OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypperOutputReference {
	var returns OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypperOutputReference
	_jsii_.Get(
		j,
		"zypper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) ZypperInput() *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper {
	var returns *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper
	_jsii_.Get(
		j,
		"zypperInput",
		&returns,
	)
	return returns
}


func NewOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference {
	_init_.Initialize()

	if err := validateNewOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.osConfigOsPolicyAssignment.OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference_Override(o OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.osConfigOsPolicyAssignment.OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference)SetInternalValue(val *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) PutApt(value *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt) {
	if err := o.validatePutAptParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putApt",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) PutGoo(value *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo) {
	if err := o.validatePutGooParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putGoo",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) PutYum(value *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum) {
	if err := o.validatePutYumParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putYum",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) PutZypper(value *OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper) {
	if err := o.validatePutZypperParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putZypper",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) ResetApt() {
	_jsii_.InvokeVoid(
		o,
		"resetApt",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) ResetGoo() {
	_jsii_.InvokeVoid(
		o,
		"resetGoo",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) ResetYum() {
	_jsii_.InvokeVoid(
		o,
		"resetYum",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) ResetZypper() {
	_jsii_.InvokeVoid(
		o,
		"resetZypper",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

