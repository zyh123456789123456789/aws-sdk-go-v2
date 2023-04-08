// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type InstanceAccessControlAttributeConfigurationStatus string

// Enum values for InstanceAccessControlAttributeConfigurationStatus
const (
	InstanceAccessControlAttributeConfigurationStatusEnabled            InstanceAccessControlAttributeConfigurationStatus = "ENABLED"
	InstanceAccessControlAttributeConfigurationStatusCreationInProgress InstanceAccessControlAttributeConfigurationStatus = "CREATION_IN_PROGRESS"
	InstanceAccessControlAttributeConfigurationStatusCreationFailed     InstanceAccessControlAttributeConfigurationStatus = "CREATION_FAILED"
)

// Values returns all known values for
// InstanceAccessControlAttributeConfigurationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InstanceAccessControlAttributeConfigurationStatus) Values() []InstanceAccessControlAttributeConfigurationStatus {
	return []InstanceAccessControlAttributeConfigurationStatus{
		"ENABLED",
		"CREATION_IN_PROGRESS",
		"CREATION_FAILED",
	}
}

type PrincipalType string

// Enum values for PrincipalType
const (
	PrincipalTypeUser  PrincipalType = "USER"
	PrincipalTypeGroup PrincipalType = "GROUP"
)

// Values returns all known values for PrincipalType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PrincipalType) Values() []PrincipalType {
	return []PrincipalType{
		"USER",
		"GROUP",
	}
}

type ProvisioningStatus string

// Enum values for ProvisioningStatus
const (
	ProvisioningStatusLatestPermissionSetProvisioned    ProvisioningStatus = "LATEST_PERMISSION_SET_PROVISIONED"
	ProvisioningStatusLatestPermissionSetNotProvisioned ProvisioningStatus = "LATEST_PERMISSION_SET_NOT_PROVISIONED"
)

// Values returns all known values for ProvisioningStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ProvisioningStatus) Values() []ProvisioningStatus {
	return []ProvisioningStatus{
		"LATEST_PERMISSION_SET_PROVISIONED",
		"LATEST_PERMISSION_SET_NOT_PROVISIONED",
	}
}

type ProvisionTargetType string

// Enum values for ProvisionTargetType
const (
	ProvisionTargetTypeAwsAccount             ProvisionTargetType = "AWS_ACCOUNT"
	ProvisionTargetTypeAllProvisionedAccounts ProvisionTargetType = "ALL_PROVISIONED_ACCOUNTS"
)

// Values returns all known values for ProvisionTargetType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ProvisionTargetType) Values() []ProvisionTargetType {
	return []ProvisionTargetType{
		"AWS_ACCOUNT",
		"ALL_PROVISIONED_ACCOUNTS",
	}
}

type StatusValues string

// Enum values for StatusValues
const (
	StatusValuesInProgress StatusValues = "IN_PROGRESS"
	StatusValuesFailed     StatusValues = "FAILED"
	StatusValuesSucceeded  StatusValues = "SUCCEEDED"
)

// Values returns all known values for StatusValues. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StatusValues) Values() []StatusValues {
	return []StatusValues{
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
	}
}

type TargetType string

// Enum values for TargetType
const (
	TargetTypeAwsAccount TargetType = "AWS_ACCOUNT"
)

// Values returns all known values for TargetType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TargetType) Values() []TargetType {
	return []TargetType{
		"AWS_ACCOUNT",
	}
}
