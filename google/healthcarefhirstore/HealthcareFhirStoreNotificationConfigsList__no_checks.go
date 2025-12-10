// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package healthcarefhirstore

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HealthcareFhirStoreNotificationConfigsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (h *jsiiProxy_HealthcareFhirStoreNotificationConfigsList) validateGetParameters(index *float64) error {
	return nil
}

func (h *jsiiProxy_HealthcareFhirStoreNotificationConfigsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_HealthcareFhirStoreNotificationConfigsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_HealthcareFhirStoreNotificationConfigsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_HealthcareFhirStoreNotificationConfigsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_HealthcareFhirStoreNotificationConfigsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewHealthcareFhirStoreNotificationConfigsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

