package dataproccluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-google-go/google/v3/jsii"

	"github.com/hashicorp/cdktf-provider-google-go/google/v3/dataproccluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference interface {
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
	// Experimental.
	Fqn() *string
	GkeClusterConfig() DataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigOutputReference
	GkeClusterConfigInput() *DataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig
	InternalValue() *DataprocClusterVirtualClusterConfigKubernetesClusterConfig
	SetInternalValue(val *DataprocClusterVirtualClusterConfigKubernetesClusterConfig)
	KubernetesNamespace() *string
	SetKubernetesNamespace(val *string)
	KubernetesNamespaceInput() *string
	KubernetesSoftwareConfig() DataprocClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfigOutputReference
	KubernetesSoftwareConfigInput() *DataprocClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig
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
	PutGkeClusterConfig(value *DataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig)
	PutKubernetesSoftwareConfig(value *DataprocClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig)
	ResetKubernetesNamespace()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference
type jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) GkeClusterConfig() DataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigOutputReference {
	var returns DataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigOutputReference
	_jsii_.Get(
		j,
		"gkeClusterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) GkeClusterConfigInput() *DataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig {
	var returns *DataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig
	_jsii_.Get(
		j,
		"gkeClusterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) InternalValue() *DataprocClusterVirtualClusterConfigKubernetesClusterConfig {
	var returns *DataprocClusterVirtualClusterConfigKubernetesClusterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) KubernetesNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) KubernetesNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) KubernetesSoftwareConfig() DataprocClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfigOutputReference {
	var returns DataprocClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfigOutputReference
	_jsii_.Get(
		j,
		"kubernetesSoftwareConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) KubernetesSoftwareConfigInput() *DataprocClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig {
	var returns *DataprocClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig
	_jsii_.Get(
		j,
		"kubernetesSoftwareConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocCluster.DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference_Override(d DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocCluster.DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference)SetInternalValue(val *DataprocClusterVirtualClusterConfigKubernetesClusterConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference)SetKubernetesNamespace(val *string) {
	if err := j.validateSetKubernetesNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kubernetesNamespace",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) PutGkeClusterConfig(value *DataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig) {
	if err := d.validatePutGkeClusterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGkeClusterConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) PutKubernetesSoftwareConfig(value *DataprocClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig) {
	if err := d.validatePutKubernetesSoftwareConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putKubernetesSoftwareConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) ResetKubernetesNamespace() {
	_jsii_.InvokeVoid(
		d,
		"resetKubernetesNamespace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataprocClusterVirtualClusterConfigKubernetesClusterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

