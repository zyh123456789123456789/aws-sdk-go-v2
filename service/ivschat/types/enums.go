// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ChatTokenCapability string

// Enum values for ChatTokenCapability
const (
	ChatTokenCapabilitySendMessage    ChatTokenCapability = "SEND_MESSAGE"
	ChatTokenCapabilityDisconnectUser ChatTokenCapability = "DISCONNECT_USER"
	ChatTokenCapabilityDeleteMessage  ChatTokenCapability = "DELETE_MESSAGE"
)

// Values returns all known values for ChatTokenCapability. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ChatTokenCapability) Values() []ChatTokenCapability {
	return []ChatTokenCapability{
		"SEND_MESSAGE",
		"DISCONNECT_USER",
		"DELETE_MESSAGE",
	}
}

type CreateLoggingConfigurationState string

// Enum values for CreateLoggingConfigurationState
const (
	CreateLoggingConfigurationStateActive CreateLoggingConfigurationState = "ACTIVE"
)

// Values returns all known values for CreateLoggingConfigurationState. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (CreateLoggingConfigurationState) Values() []CreateLoggingConfigurationState {
	return []CreateLoggingConfigurationState{
		"ACTIVE",
	}
}

type FallbackResult string

// Enum values for FallbackResult
const (
	FallbackResultAllow FallbackResult = "ALLOW"
	FallbackResultDeny  FallbackResult = "DENY"
)

// Values returns all known values for FallbackResult. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FallbackResult) Values() []FallbackResult {
	return []FallbackResult{
		"ALLOW",
		"DENY",
	}
}

type LoggingConfigurationState string

// Enum values for LoggingConfigurationState
const (
	LoggingConfigurationStateCreating     LoggingConfigurationState = "CREATING"
	LoggingConfigurationStateCreateFailed LoggingConfigurationState = "CREATE_FAILED"
	LoggingConfigurationStateDeleting     LoggingConfigurationState = "DELETING"
	LoggingConfigurationStateDeleteFailed LoggingConfigurationState = "DELETE_FAILED"
	LoggingConfigurationStateUpdating     LoggingConfigurationState = "UPDATING"
	LoggingConfigurationStateUpdateFailed LoggingConfigurationState = "UPDATE_FAILED"
	LoggingConfigurationStateActive       LoggingConfigurationState = "ACTIVE"
)

// Values returns all known values for LoggingConfigurationState. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (LoggingConfigurationState) Values() []LoggingConfigurationState {
	return []LoggingConfigurationState{
		"CREATING",
		"CREATE_FAILED",
		"DELETING",
		"DELETE_FAILED",
		"UPDATING",
		"UPDATE_FAILED",
		"ACTIVE",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeRoom ResourceType = "ROOM"
)

// Values returns all known values for ResourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"ROOM",
	}
}

type UpdateLoggingConfigurationState string

// Enum values for UpdateLoggingConfigurationState
const (
	UpdateLoggingConfigurationStateActive UpdateLoggingConfigurationState = "ACTIVE"
)

// Values returns all known values for UpdateLoggingConfigurationState. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (UpdateLoggingConfigurationState) Values() []UpdateLoggingConfigurationState {
	return []UpdateLoggingConfigurationState{
		"ACTIVE",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonUnknownOperation      ValidationExceptionReason = "UNKNOWN_OPERATION"
	ValidationExceptionReasonFieldValidationFailed ValidationExceptionReason = "FIELD_VALIDATION_FAILED"
	ValidationExceptionReasonOther                 ValidationExceptionReason = "OTHER"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"UNKNOWN_OPERATION",
		"FIELD_VALIDATION_FAILED",
		"OTHER",
	}
}
