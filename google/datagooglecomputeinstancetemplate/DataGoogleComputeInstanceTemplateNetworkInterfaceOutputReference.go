package datagooglecomputeinstancetemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v5/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v5/datagooglecomputeinstancetemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference interface {
	cdktf.ComplexObject
	AccessConfig() DataGoogleComputeInstanceTemplateNetworkInterfaceAccessConfigList
	AliasIpRange() DataGoogleComputeInstanceTemplateNetworkInterfaceAliasIpRangeList
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
	InternalValue() *DataGoogleComputeInstanceTemplateNetworkInterface
	SetInternalValue(val *DataGoogleComputeInstanceTemplateNetworkInterface)
	Ipv6AccessConfig() DataGoogleComputeInstanceTemplateNetworkInterfaceIpv6AccessConfigList
	Ipv6AccessType() *string
	Name() *string
	Network() *string
	NetworkIp() *string
	NicType() *string
	QueueCount() *float64
	StackType() *string
	Subnetwork() *string
	SubnetworkProject() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference
type jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) AccessConfig() DataGoogleComputeInstanceTemplateNetworkInterfaceAccessConfigList {
	var returns DataGoogleComputeInstanceTemplateNetworkInterfaceAccessConfigList
	_jsii_.Get(
		j,
		"accessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) AliasIpRange() DataGoogleComputeInstanceTemplateNetworkInterfaceAliasIpRangeList {
	var returns DataGoogleComputeInstanceTemplateNetworkInterfaceAliasIpRangeList
	_jsii_.Get(
		j,
		"aliasIpRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) InternalValue() *DataGoogleComputeInstanceTemplateNetworkInterface {
	var returns *DataGoogleComputeInstanceTemplateNetworkInterface
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) Ipv6AccessConfig() DataGoogleComputeInstanceTemplateNetworkInterfaceIpv6AccessConfigList {
	var returns DataGoogleComputeInstanceTemplateNetworkInterfaceIpv6AccessConfigList
	_jsii_.Get(
		j,
		"ipv6AccessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) Ipv6AccessType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6AccessType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) NetworkIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) NicType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nicType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) QueueCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queueCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) StackType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) Subnetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) SubnetworkProject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleComputeInstanceTemplateNetworkInterfaceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleComputeInstanceTemplate.DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference_Override(d DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleComputeInstanceTemplate.DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference)SetInternalValue(val *DataGoogleComputeInstanceTemplateNetworkInterface) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataGoogleComputeInstanceTemplateNetworkInterfaceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

