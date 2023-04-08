// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ActionCode string

// Enum values for ActionCode
const (
	ActionCodeArchiveRetrieval   ActionCode = "ArchiveRetrieval"
	ActionCodeInventoryRetrieval ActionCode = "InventoryRetrieval"
	ActionCodeSelect             ActionCode = "Select"
)

// Values returns all known values for ActionCode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ActionCode) Values() []ActionCode {
	return []ActionCode{
		"ArchiveRetrieval",
		"InventoryRetrieval",
		"Select",
	}
}

type CannedACL string

// Enum values for CannedACL
const (
	CannedACLPrivate                CannedACL = "private"
	CannedACLPublicRead             CannedACL = "public-read"
	CannedACLPublicReadWrite        CannedACL = "public-read-write"
	CannedACLAwsExecRead            CannedACL = "aws-exec-read"
	CannedACLAuthenticatedRead      CannedACL = "authenticated-read"
	CannedACLBucketOwnerRead        CannedACL = "bucket-owner-read"
	CannedACLBucketOwnerFullControl CannedACL = "bucket-owner-full-control"
)

// Values returns all known values for CannedACL. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (CannedACL) Values() []CannedACL {
	return []CannedACL{
		"private",
		"public-read",
		"public-read-write",
		"aws-exec-read",
		"authenticated-read",
		"bucket-owner-read",
		"bucket-owner-full-control",
	}
}

type EncryptionType string

// Enum values for EncryptionType
const (
	EncryptionTypeKms EncryptionType = "aws:kms"
	EncryptionTypeS3  EncryptionType = "AES256"
)

// Values returns all known values for EncryptionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EncryptionType) Values() []EncryptionType {
	return []EncryptionType{
		"aws:kms",
		"AES256",
	}
}

type ExpressionType string

// Enum values for ExpressionType
const (
	ExpressionTypeSql ExpressionType = "SQL"
)

// Values returns all known values for ExpressionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExpressionType) Values() []ExpressionType {
	return []ExpressionType{
		"SQL",
	}
}

type FileHeaderInfo string

// Enum values for FileHeaderInfo
const (
	FileHeaderInfoUse    FileHeaderInfo = "USE"
	FileHeaderInfoIgnore FileHeaderInfo = "IGNORE"
	FileHeaderInfoNone   FileHeaderInfo = "NONE"
)

// Values returns all known values for FileHeaderInfo. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FileHeaderInfo) Values() []FileHeaderInfo {
	return []FileHeaderInfo{
		"USE",
		"IGNORE",
		"NONE",
	}
}

type Permission string

// Enum values for Permission
const (
	PermissionFullControl Permission = "FULL_CONTROL"
	PermissionWrite       Permission = "WRITE"
	PermissionWriteAcp    Permission = "WRITE_ACP"
	PermissionRead        Permission = "READ"
	PermissionReadAcp     Permission = "READ_ACP"
)

// Values returns all known values for Permission. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (Permission) Values() []Permission {
	return []Permission{
		"FULL_CONTROL",
		"WRITE",
		"WRITE_ACP",
		"READ",
		"READ_ACP",
	}
}

type QuoteFields string

// Enum values for QuoteFields
const (
	QuoteFieldsAlways   QuoteFields = "ALWAYS"
	QuoteFieldsAsNeeded QuoteFields = "ASNEEDED"
)

// Values returns all known values for QuoteFields. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (QuoteFields) Values() []QuoteFields {
	return []QuoteFields{
		"ALWAYS",
		"ASNEEDED",
	}
}

type StatusCode string

// Enum values for StatusCode
const (
	StatusCodeInProgress StatusCode = "InProgress"
	StatusCodeSucceeded  StatusCode = "Succeeded"
	StatusCodeFailed     StatusCode = "Failed"
)

// Values returns all known values for StatusCode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (StatusCode) Values() []StatusCode {
	return []StatusCode{
		"InProgress",
		"Succeeded",
		"Failed",
	}
}

type StorageClass string

// Enum values for StorageClass
const (
	StorageClassStandard                 StorageClass = "STANDARD"
	StorageClassReducedRedundancy        StorageClass = "REDUCED_REDUNDANCY"
	StorageClassStandardInfrequentAccess StorageClass = "STANDARD_IA"
)

// Values returns all known values for StorageClass. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StorageClass) Values() []StorageClass {
	return []StorageClass{
		"STANDARD",
		"REDUCED_REDUNDANCY",
		"STANDARD_IA",
	}
}

type Type string

// Enum values for Type
const (
	TypeAmazonCustomerByEmail Type = "AmazonCustomerByEmail"
	TypeCanonicalUser         Type = "CanonicalUser"
	TypeGroup                 Type = "Group"
)

// Values returns all known values for Type. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Type) Values() []Type {
	return []Type{
		"AmazonCustomerByEmail",
		"CanonicalUser",
		"Group",
	}
}
