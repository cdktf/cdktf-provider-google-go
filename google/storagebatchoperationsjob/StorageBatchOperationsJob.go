// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagebatchoperationsjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v16/storagebatchoperationsjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/storage_batch_operations_job google_storage_batch_operations_job}.
type StorageBatchOperationsJob interface {
	cdktf.TerraformResource
	BucketList() StorageBatchOperationsJobBucketListStructOutputReference
	BucketListInput() *StorageBatchOperationsJobBucketListStruct
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CompleteTime() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	DeleteObject() StorageBatchOperationsJobDeleteObjectOutputReference
	DeleteObjectInput() *StorageBatchOperationsJobDeleteObject
	DeleteProtection() interface{}
	SetDeleteProtection(val interface{})
	DeleteProtectionInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	JobId() *string
	SetJobId(val *string)
	JobIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PutMetadata() StorageBatchOperationsJobPutMetadataOutputReference
	PutMetadataInput() *StorageBatchOperationsJobPutMetadata
	PutObjectHold() StorageBatchOperationsJobPutObjectHoldOutputReference
	PutObjectHoldInput() *StorageBatchOperationsJobPutObjectHold
	// Experimental.
	RawOverrides() interface{}
	RewriteObject() StorageBatchOperationsJobRewriteObjectOutputReference
	RewriteObjectInput() *StorageBatchOperationsJobRewriteObject
	ScheduleTime() *string
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() StorageBatchOperationsJobTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateTime() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutBucketList(value *StorageBatchOperationsJobBucketListStruct)
	PutDeleteObject(value *StorageBatchOperationsJobDeleteObject)
	PutPutMetadata(value *StorageBatchOperationsJobPutMetadata)
	PutPutObjectHold(value *StorageBatchOperationsJobPutObjectHold)
	PutRewriteObject(value *StorageBatchOperationsJobRewriteObject)
	PutTimeouts(value *StorageBatchOperationsJobTimeouts)
	ResetBucketList()
	ResetDeleteObject()
	ResetDeleteProtection()
	ResetId()
	ResetJobId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetPutMetadata()
	ResetPutObjectHold()
	ResetRewriteObject()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for StorageBatchOperationsJob
type jsiiProxy_StorageBatchOperationsJob struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StorageBatchOperationsJob) BucketList() StorageBatchOperationsJobBucketListStructOutputReference {
	var returns StorageBatchOperationsJobBucketListStructOutputReference
	_jsii_.Get(
		j,
		"bucketList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) BucketListInput() *StorageBatchOperationsJobBucketListStruct {
	var returns *StorageBatchOperationsJobBucketListStruct
	_jsii_.Get(
		j,
		"bucketListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) CompleteTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"completeTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) DeleteObject() StorageBatchOperationsJobDeleteObjectOutputReference {
	var returns StorageBatchOperationsJobDeleteObjectOutputReference
	_jsii_.Get(
		j,
		"deleteObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) DeleteObjectInput() *StorageBatchOperationsJobDeleteObject {
	var returns *StorageBatchOperationsJobDeleteObject
	_jsii_.Get(
		j,
		"deleteObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) DeleteProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) DeleteProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) JobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) JobIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) PutMetadata() StorageBatchOperationsJobPutMetadataOutputReference {
	var returns StorageBatchOperationsJobPutMetadataOutputReference
	_jsii_.Get(
		j,
		"putMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) PutMetadataInput() *StorageBatchOperationsJobPutMetadata {
	var returns *StorageBatchOperationsJobPutMetadata
	_jsii_.Get(
		j,
		"putMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) PutObjectHold() StorageBatchOperationsJobPutObjectHoldOutputReference {
	var returns StorageBatchOperationsJobPutObjectHoldOutputReference
	_jsii_.Get(
		j,
		"putObjectHold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) PutObjectHoldInput() *StorageBatchOperationsJobPutObjectHold {
	var returns *StorageBatchOperationsJobPutObjectHold
	_jsii_.Get(
		j,
		"putObjectHoldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) RewriteObject() StorageBatchOperationsJobRewriteObjectOutputReference {
	var returns StorageBatchOperationsJobRewriteObjectOutputReference
	_jsii_.Get(
		j,
		"rewriteObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) RewriteObjectInput() *StorageBatchOperationsJobRewriteObject {
	var returns *StorageBatchOperationsJobRewriteObject
	_jsii_.Get(
		j,
		"rewriteObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) ScheduleTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) Timeouts() StorageBatchOperationsJobTimeoutsOutputReference {
	var returns StorageBatchOperationsJobTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJob) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/storage_batch_operations_job google_storage_batch_operations_job} Resource.
func NewStorageBatchOperationsJob(scope constructs.Construct, id *string, config *StorageBatchOperationsJobConfig) StorageBatchOperationsJob {
	_init_.Initialize()

	if err := validateNewStorageBatchOperationsJobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageBatchOperationsJob{}

	_jsii_.Create(
		"@cdktf/provider-google.storageBatchOperationsJob.StorageBatchOperationsJob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/storage_batch_operations_job google_storage_batch_operations_job} Resource.
func NewStorageBatchOperationsJob_Override(s StorageBatchOperationsJob, scope constructs.Construct, id *string, config *StorageBatchOperationsJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.storageBatchOperationsJob.StorageBatchOperationsJob",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StorageBatchOperationsJob)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_StorageBatchOperationsJob)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StorageBatchOperationsJob)SetDeleteProtection(val interface{}) {
	if err := j.validateSetDeleteProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteProtection",
		val,
	)
}

func (j *jsiiProxy_StorageBatchOperationsJob)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StorageBatchOperationsJob)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_StorageBatchOperationsJob)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_StorageBatchOperationsJob)SetJobId(val *string) {
	if err := j.validateSetJobIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobId",
		val,
	)
}

func (j *jsiiProxy_StorageBatchOperationsJob)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StorageBatchOperationsJob)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_StorageBatchOperationsJob)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StorageBatchOperationsJob)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a StorageBatchOperationsJob resource upon running "cdktf plan <stack-name>".
func StorageBatchOperationsJob_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateStorageBatchOperationsJob_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.storageBatchOperationsJob.StorageBatchOperationsJob",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func StorageBatchOperationsJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageBatchOperationsJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.storageBatchOperationsJob.StorageBatchOperationsJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StorageBatchOperationsJob_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageBatchOperationsJob_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.storageBatchOperationsJob.StorageBatchOperationsJob",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StorageBatchOperationsJob_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageBatchOperationsJob_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.storageBatchOperationsJob.StorageBatchOperationsJob",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StorageBatchOperationsJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.storageBatchOperationsJob.StorageBatchOperationsJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StorageBatchOperationsJob) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_StorageBatchOperationsJob) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_StorageBatchOperationsJob) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBatchOperationsJob) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBatchOperationsJob) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBatchOperationsJob) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBatchOperationsJob) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBatchOperationsJob) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBatchOperationsJob) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBatchOperationsJob) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBatchOperationsJob) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBatchOperationsJob) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBatchOperationsJob) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_StorageBatchOperationsJob) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBatchOperationsJob) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StorageBatchOperationsJob) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_StorageBatchOperationsJob) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StorageBatchOperationsJob) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StorageBatchOperationsJob) PutBucketList(value *StorageBatchOperationsJobBucketListStruct) {
	if err := s.validatePutBucketListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBucketList",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageBatchOperationsJob) PutDeleteObject(value *StorageBatchOperationsJobDeleteObject) {
	if err := s.validatePutDeleteObjectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDeleteObject",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageBatchOperationsJob) PutPutMetadata(value *StorageBatchOperationsJobPutMetadata) {
	if err := s.validatePutPutMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPutMetadata",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageBatchOperationsJob) PutPutObjectHold(value *StorageBatchOperationsJobPutObjectHold) {
	if err := s.validatePutPutObjectHoldParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPutObjectHold",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageBatchOperationsJob) PutRewriteObject(value *StorageBatchOperationsJobRewriteObject) {
	if err := s.validatePutRewriteObjectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRewriteObject",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageBatchOperationsJob) PutTimeouts(value *StorageBatchOperationsJobTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageBatchOperationsJob) ResetBucketList() {
	_jsii_.InvokeVoid(
		s,
		"resetBucketList",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBatchOperationsJob) ResetDeleteObject() {
	_jsii_.InvokeVoid(
		s,
		"resetDeleteObject",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBatchOperationsJob) ResetDeleteProtection() {
	_jsii_.InvokeVoid(
		s,
		"resetDeleteProtection",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBatchOperationsJob) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBatchOperationsJob) ResetJobId() {
	_jsii_.InvokeVoid(
		s,
		"resetJobId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBatchOperationsJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBatchOperationsJob) ResetProject() {
	_jsii_.InvokeVoid(
		s,
		"resetProject",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBatchOperationsJob) ResetPutMetadata() {
	_jsii_.InvokeVoid(
		s,
		"resetPutMetadata",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBatchOperationsJob) ResetPutObjectHold() {
	_jsii_.InvokeVoid(
		s,
		"resetPutObjectHold",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBatchOperationsJob) ResetRewriteObject() {
	_jsii_.InvokeVoid(
		s,
		"resetRewriteObject",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBatchOperationsJob) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBatchOperationsJob) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBatchOperationsJob) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBatchOperationsJob) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBatchOperationsJob) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBatchOperationsJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBatchOperationsJob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

