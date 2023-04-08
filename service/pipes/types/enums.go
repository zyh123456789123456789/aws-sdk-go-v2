// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AssignPublicIp string

// Enum values for AssignPublicIp
const (
	AssignPublicIpEnabled  AssignPublicIp = "ENABLED"
	AssignPublicIpDisabled AssignPublicIp = "DISABLED"
)

// Values returns all known values for AssignPublicIp. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AssignPublicIp) Values() []AssignPublicIp {
	return []AssignPublicIp{
		"ENABLED",
		"DISABLED",
	}
}

type BatchJobDependencyType string

// Enum values for BatchJobDependencyType
const (
	BatchJobDependencyTypeNToN       BatchJobDependencyType = "N_TO_N"
	BatchJobDependencyTypeSequential BatchJobDependencyType = "SEQUENTIAL"
)

// Values returns all known values for BatchJobDependencyType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BatchJobDependencyType) Values() []BatchJobDependencyType {
	return []BatchJobDependencyType{
		"N_TO_N",
		"SEQUENTIAL",
	}
}

type BatchResourceRequirementType string

// Enum values for BatchResourceRequirementType
const (
	BatchResourceRequirementTypeGpu    BatchResourceRequirementType = "GPU"
	BatchResourceRequirementTypeMemory BatchResourceRequirementType = "MEMORY"
	BatchResourceRequirementTypeVcpu   BatchResourceRequirementType = "VCPU"
)

// Values returns all known values for BatchResourceRequirementType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (BatchResourceRequirementType) Values() []BatchResourceRequirementType {
	return []BatchResourceRequirementType{
		"GPU",
		"MEMORY",
		"VCPU",
	}
}

type DynamoDBStreamStartPosition string

// Enum values for DynamoDBStreamStartPosition
const (
	DynamoDBStreamStartPositionTrimHorizon DynamoDBStreamStartPosition = "TRIM_HORIZON"
	DynamoDBStreamStartPositionLatest      DynamoDBStreamStartPosition = "LATEST"
)

// Values returns all known values for DynamoDBStreamStartPosition. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DynamoDBStreamStartPosition) Values() []DynamoDBStreamStartPosition {
	return []DynamoDBStreamStartPosition{
		"TRIM_HORIZON",
		"LATEST",
	}
}

type EcsEnvironmentFileType string

// Enum values for EcsEnvironmentFileType
const (
	EcsEnvironmentFileTypeS3 EcsEnvironmentFileType = "s3"
)

// Values returns all known values for EcsEnvironmentFileType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EcsEnvironmentFileType) Values() []EcsEnvironmentFileType {
	return []EcsEnvironmentFileType{
		"s3",
	}
}

type EcsResourceRequirementType string

// Enum values for EcsResourceRequirementType
const (
	EcsResourceRequirementTypeGpu                  EcsResourceRequirementType = "GPU"
	EcsResourceRequirementTypeInferenceAccelerator EcsResourceRequirementType = "InferenceAccelerator"
)

// Values returns all known values for EcsResourceRequirementType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (EcsResourceRequirementType) Values() []EcsResourceRequirementType {
	return []EcsResourceRequirementType{
		"GPU",
		"InferenceAccelerator",
	}
}

type KinesisStreamStartPosition string

// Enum values for KinesisStreamStartPosition
const (
	KinesisStreamStartPositionTrimHorizon KinesisStreamStartPosition = "TRIM_HORIZON"
	KinesisStreamStartPositionLatest      KinesisStreamStartPosition = "LATEST"
	KinesisStreamStartPositionAtTimestamp KinesisStreamStartPosition = "AT_TIMESTAMP"
)

// Values returns all known values for KinesisStreamStartPosition. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (KinesisStreamStartPosition) Values() []KinesisStreamStartPosition {
	return []KinesisStreamStartPosition{
		"TRIM_HORIZON",
		"LATEST",
		"AT_TIMESTAMP",
	}
}

type LaunchType string

// Enum values for LaunchType
const (
	LaunchTypeEc2      LaunchType = "EC2"
	LaunchTypeFargate  LaunchType = "FARGATE"
	LaunchTypeExternal LaunchType = "EXTERNAL"
)

// Values returns all known values for LaunchType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (LaunchType) Values() []LaunchType {
	return []LaunchType{
		"EC2",
		"FARGATE",
		"EXTERNAL",
	}
}

type MSKStartPosition string

// Enum values for MSKStartPosition
const (
	MSKStartPositionTrimHorizon MSKStartPosition = "TRIM_HORIZON"
	MSKStartPositionLatest      MSKStartPosition = "LATEST"
)

// Values returns all known values for MSKStartPosition. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MSKStartPosition) Values() []MSKStartPosition {
	return []MSKStartPosition{
		"TRIM_HORIZON",
		"LATEST",
	}
}

type OnPartialBatchItemFailureStreams string

// Enum values for OnPartialBatchItemFailureStreams
const (
	OnPartialBatchItemFailureStreamsAutomaticBisect OnPartialBatchItemFailureStreams = "AUTOMATIC_BISECT"
)

// Values returns all known values for OnPartialBatchItemFailureStreams. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (OnPartialBatchItemFailureStreams) Values() []OnPartialBatchItemFailureStreams {
	return []OnPartialBatchItemFailureStreams{
		"AUTOMATIC_BISECT",
	}
}

type PipeState string

// Enum values for PipeState
const (
	PipeStateRunning      PipeState = "RUNNING"
	PipeStateStopped      PipeState = "STOPPED"
	PipeStateCreating     PipeState = "CREATING"
	PipeStateUpdating     PipeState = "UPDATING"
	PipeStateDeleting     PipeState = "DELETING"
	PipeStateStarting     PipeState = "STARTING"
	PipeStateStopping     PipeState = "STOPPING"
	PipeStateCreateFailed PipeState = "CREATE_FAILED"
	PipeStateUpdateFailed PipeState = "UPDATE_FAILED"
	PipeStateStartFailed  PipeState = "START_FAILED"
	PipeStateStopFailed   PipeState = "STOP_FAILED"
)

// Values returns all known values for PipeState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (PipeState) Values() []PipeState {
	return []PipeState{
		"RUNNING",
		"STOPPED",
		"CREATING",
		"UPDATING",
		"DELETING",
		"STARTING",
		"STOPPING",
		"CREATE_FAILED",
		"UPDATE_FAILED",
		"START_FAILED",
		"STOP_FAILED",
	}
}

type PipeTargetInvocationType string

// Enum values for PipeTargetInvocationType
const (
	PipeTargetInvocationTypeRequestResponse PipeTargetInvocationType = "REQUEST_RESPONSE"
	PipeTargetInvocationTypeFireAndForget   PipeTargetInvocationType = "FIRE_AND_FORGET"
)

// Values returns all known values for PipeTargetInvocationType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (PipeTargetInvocationType) Values() []PipeTargetInvocationType {
	return []PipeTargetInvocationType{
		"REQUEST_RESPONSE",
		"FIRE_AND_FORGET",
	}
}

type PlacementConstraintType string

// Enum values for PlacementConstraintType
const (
	PlacementConstraintTypeDistinctInstance PlacementConstraintType = "distinctInstance"
	PlacementConstraintTypeMemberOf         PlacementConstraintType = "memberOf"
)

// Values returns all known values for PlacementConstraintType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PlacementConstraintType) Values() []PlacementConstraintType {
	return []PlacementConstraintType{
		"distinctInstance",
		"memberOf",
	}
}

type PlacementStrategyType string

// Enum values for PlacementStrategyType
const (
	PlacementStrategyTypeRandom  PlacementStrategyType = "random"
	PlacementStrategyTypeSpread  PlacementStrategyType = "spread"
	PlacementStrategyTypeBinpack PlacementStrategyType = "binpack"
)

// Values returns all known values for PlacementStrategyType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PlacementStrategyType) Values() []PlacementStrategyType {
	return []PlacementStrategyType{
		"random",
		"spread",
		"binpack",
	}
}

type PropagateTags string

// Enum values for PropagateTags
const (
	PropagateTagsTaskDefinition PropagateTags = "TASK_DEFINITION"
)

// Values returns all known values for PropagateTags. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PropagateTags) Values() []PropagateTags {
	return []PropagateTags{
		"TASK_DEFINITION",
	}
}

type RequestedPipeState string

// Enum values for RequestedPipeState
const (
	RequestedPipeStateRunning RequestedPipeState = "RUNNING"
	RequestedPipeStateStopped RequestedPipeState = "STOPPED"
)

// Values returns all known values for RequestedPipeState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RequestedPipeState) Values() []RequestedPipeState {
	return []RequestedPipeState{
		"RUNNING",
		"STOPPED",
	}
}

type RequestedPipeStateDescribeResponse string

// Enum values for RequestedPipeStateDescribeResponse
const (
	RequestedPipeStateDescribeResponseRunning RequestedPipeStateDescribeResponse = "RUNNING"
	RequestedPipeStateDescribeResponseStopped RequestedPipeStateDescribeResponse = "STOPPED"
	RequestedPipeStateDescribeResponseDeleted RequestedPipeStateDescribeResponse = "DELETED"
)

// Values returns all known values for RequestedPipeStateDescribeResponse. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (RequestedPipeStateDescribeResponse) Values() []RequestedPipeStateDescribeResponse {
	return []RequestedPipeStateDescribeResponse{
		"RUNNING",
		"STOPPED",
		"DELETED",
	}
}

type SelfManagedKafkaStartPosition string

// Enum values for SelfManagedKafkaStartPosition
const (
	SelfManagedKafkaStartPositionTrimHorizon SelfManagedKafkaStartPosition = "TRIM_HORIZON"
	SelfManagedKafkaStartPositionLatest      SelfManagedKafkaStartPosition = "LATEST"
)

// Values returns all known values for SelfManagedKafkaStartPosition. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (SelfManagedKafkaStartPosition) Values() []SelfManagedKafkaStartPosition {
	return []SelfManagedKafkaStartPosition{
		"TRIM_HORIZON",
		"LATEST",
	}
}
