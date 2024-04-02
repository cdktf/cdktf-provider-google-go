// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package integrationsclient

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IntegrationsClient) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (i *jsiiProxy_IntegrationsClient) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (i *jsiiProxy_IntegrationsClient) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_IntegrationsClient) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_IntegrationsClient) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_IntegrationsClient) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_IntegrationsClient) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_IntegrationsClient) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_IntegrationsClient) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_IntegrationsClient) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_IntegrationsClient) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_IntegrationsClient) validateImportFromParameters(id *string) error {
	return nil
}

func (i *jsiiProxy_IntegrationsClient) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_IntegrationsClient) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (i *jsiiProxy_IntegrationsClient) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (i *jsiiProxy_IntegrationsClient) validateMoveToIdParameters(id *string) error {
	return nil
}

func (i *jsiiProxy_IntegrationsClient) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (i *jsiiProxy_IntegrationsClient) validatePutCloudKmsConfigParameters(value *IntegrationsClientCloudKmsConfig) error {
	return nil
}

func (i *jsiiProxy_IntegrationsClient) validatePutTimeoutsParameters(value *IntegrationsClientTimeouts) error {
	return nil
}

func validateIntegrationsClient_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateIntegrationsClient_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIntegrationsClient_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateIntegrationsClient_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_IntegrationsClient) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IntegrationsClient) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IntegrationsClient) validateSetCreateSampleWorkflowsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IntegrationsClient) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IntegrationsClient) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_IntegrationsClient) validateSetLocationParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IntegrationsClient) validateSetProjectParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IntegrationsClient) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_IntegrationsClient) validateSetProvisionGmekParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IntegrationsClient) validateSetRunAsServiceAccountParameters(val *string) error {
	return nil
}

func validateNewIntegrationsClientParameters(scope constructs.Construct, id *string, config *IntegrationsClientConfig) error {
	return nil
}

