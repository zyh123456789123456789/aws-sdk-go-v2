// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Describes a principal for use with Resource Access Manager.
type Principal struct {

	// The date and time when the principal was associated with the resource share.
	CreationTime *time.Time

	// Indicates whether the principal belongs to the same organization in
	// Organizations as the Amazon Web Services account that owns the resource share.
	External *bool

	// The ID of the principal.
	Id *string

	// The date and time when the association was last updated.
	LastUpdatedTime *time.Time

	// The Amazon Resoure Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of a resource share the principal is associated with.
	ResourceShareArn *string

	noSmithyDocumentSerde
}

// Describes a resource associated with a resource share in RAM.
type Resource struct {

	// The Amazon Resoure Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the resource.
	Arn *string

	// The date and time when the resource was associated with the resource share.
	CreationTime *time.Time

	// The date an time when the association was last updated.
	LastUpdatedTime *time.Time

	// The Amazon Resoure Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the resource group. This value is available only if the resource is part of a
	// resource group.
	ResourceGroupArn *string

	// Specifies the scope of visibility of this resource:
	//   - REGIONAL – The resource can be accessed only by using requests that target
	//   the Amazon Web Services Region in which the resource exists.
	//   - GLOBAL – The resource can be accessed from any Amazon Web Services Region.
	ResourceRegionScope ResourceRegionScope

	// The Amazon Resoure Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the resource share this resource is associated with.
	ResourceShareArn *string

	// The current status of the resource.
	Status ResourceStatus

	// A message about the status of the resource.
	StatusMessage *string

	// The resource type. This takes the form of: service-code : resource-code
	Type *string

	noSmithyDocumentSerde
}

// Describes a resource share in RAM.
type ResourceShare struct {

	// Indicates whether principals outside your organization in Organizations can be
	// associated with a resource share.
	AllowExternalPrincipals *bool

	// The date and time when the resource share was created.
	CreationTime *time.Time

	// Indicates how the resource share was created. Possible values include:
	//   - CREATED_FROM_POLICY - Indicates that the resource share was created from an
	//   Identity and Access Management (IAM) resource-based permission policy attached
	//   to the resource. This type of resource share is visible only to the Amazon Web
	//   Services account that created it. You can't modify it in RAM unless you promote
	//   it. For more information, see PromoteResourceShareCreatedFromPolicy .
	//   - PROMOTING_TO_STANDARD - The resource share is in the process of being
	//   promoted. For more information, see PromoteResourceShareCreatedFromPolicy .
	//   - STANDARD - Indicates that the resource share was created in RAM using the
	//   console or APIs. These resource shares are visible to all principals you share
	//   the resource share with. You can modify these resource shares in RAM using the
	//   console or APIs.
	FeatureSet ResourceShareFeatureSet

	// The date and time when the resource share was last updated.
	LastUpdatedTime *time.Time

	// The name of the resource share.
	Name *string

	// The ID of the Amazon Web Services account that owns the resource share.
	OwningAccountId *string

	// The Amazon Resoure Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the resource share
	ResourceShareArn *string

	// The current status of the resource share.
	Status ResourceShareStatus

	// A message about the status of the resource share.
	StatusMessage *string

	// The tag key and value pairs attached to the resource share.
	Tags []Tag

	noSmithyDocumentSerde
}

// Describes an association with a resource share and either a principal or a
// resource.
type ResourceShareAssociation struct {

	// The associated entity. This can be either of the following:
	//   - For a resource association, this is the Amazon Resoure Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	//   of the resource.
	//   - For principal associations, this is one of the following:
	//   - The ID of an Amazon Web Services account
	//   - The Amazon Resoure Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	//   of an organization in Organizations
	//   - The ARN of an organizational unit (OU) in Organizations
	//   - The ARN of an IAM role
	//   - The ARN of an IAM user
	AssociatedEntity *string

	// The type of entity included in this association.
	AssociationType ResourceShareAssociationType

	// The date and time when the association was created.
	CreationTime *time.Time

	// Indicates whether the principal belongs to the same organization in
	// Organizations as the Amazon Web Services account that owns the resource share.
	External *bool

	// The date and time when the association was last updated.
	LastUpdatedTime *time.Time

	// The Amazon Resoure Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the resource share.
	ResourceShareArn *string

	// The name of the resource share.
	ResourceShareName *string

	// The current status of the association.
	Status ResourceShareAssociationStatus

	// A message about the status of the association.
	StatusMessage *string

	noSmithyDocumentSerde
}

// Describes an invitation for an Amazon Web Services account to join a resource
// share.
type ResourceShareInvitation struct {

	// The date and time when the invitation was sent.
	InvitationTimestamp *time.Time

	// The ID of the Amazon Web Services account that received the invitation.
	ReceiverAccountId *string

	// The Amazon Resoure Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the IAM user or role that received the invitation.
	ReceiverArn *string

	// The Amazon Resoure Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the resource share
	ResourceShareArn *string

	// To view the resources associated with a pending resource share invitation, use
	// ListPendingInvitationResources .
	//
	// Deprecated: This member has been deprecated. Use ListPendingInvitationResources.
	ResourceShareAssociations []ResourceShareAssociation

	// The Amazon Resoure Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the invitation.
	ResourceShareInvitationArn *string

	// The name of the resource share.
	ResourceShareName *string

	// The ID of the Amazon Web Services account that sent the invitation.
	SenderAccountId *string

	// The current status of the invitation.
	Status ResourceShareInvitationStatus

	noSmithyDocumentSerde
}

// Information about an RAM permission.
type ResourceSharePermissionDetail struct {

	// The Amazon Resoure Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of this RAM permission.
	Arn *string

	// The date and time when the permission was created.
	CreationTime *time.Time

	// Specifies whether the version of the permission represented in this structure
	// is the default version for this permission.
	DefaultVersion *bool

	// Specifies whether the version of the permission represented in this structure
	// is the default version for all resources of this resource type.
	IsResourceTypeDefault *bool

	// The date and time when the permission was last updated.
	LastUpdatedTime *time.Time

	// The name of this permission.
	Name *string

	// The permission's effect and actions in JSON format. The effect indicates
	// whether the specified actions are allowed or denied. The actions list the
	// operations to which the principal is granted or denied access.
	Permission *string

	// The resource type to which this permission applies.
	ResourceType *string

	// The version of the permission represented in this structure.
	Version *string

	noSmithyDocumentSerde
}

// Information about an RAM permission that is associated with a resource share
// and any of its resources of a specified type.
type ResourceSharePermissionSummary struct {

	// The Amazon Resoure Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the permission you want information about.
	Arn *string

	// The date and time when the permission was created.
	CreationTime *time.Time

	// Specifies whether the version of the permission represented in this structure
	// is the default version for this permission.
	DefaultVersion *bool

	// Specifies whether the version of the permission represented in this structure
	// is the default version for all resources of this resource type.
	IsResourceTypeDefault *bool

	// The date and time when the permission was last updated.
	LastUpdatedTime *time.Time

	// The name of this permission.
	Name *string

	// The type of resource to which this permission applies.
	ResourceType *string

	// The current status of the permission.
	Status *string

	// The version of the permission represented in this structure.
	Version *string

	noSmithyDocumentSerde
}

// Information about a shareable resource type and the Amazon Web Services service
// to which resources of that type belong.
type ServiceNameAndResourceType struct {

	// Specifies the scope of visibility of resources of this type:
	//   - REGIONAL – The resource can be accessed only by using requests that target
	//   the Amazon Web Services Region in which the resource exists.
	//   - GLOBAL – The resource can be accessed from any Amazon Web Services Region.
	ResourceRegionScope ResourceRegionScope

	// The type of the resource.
	ResourceType *string

	// The name of the Amazon Web Services service to which resources of this type
	// belong.
	ServiceName *string

	noSmithyDocumentSerde
}

// A structure containing a tag. A tag is metadata that you can attach to your
// resources to help organize and categorize them. You can also use them to help
// you secure your resources. For more information, see Controlling access to
// Amazon Web Services resources using tags (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html)
// . For more information about tags, see Tagging Amazon Web Services resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
// in the Amazon Web Services General Reference Guide.
type Tag struct {

	// The key, or name, attached to the tag. Every tag must have a key. Key names are
	// case sensitive.
	Key *string

	// The string value attached to the tag. The value can be an empty string. Key
	// values are case sensitive.
	Value *string

	noSmithyDocumentSerde
}

// A tag key and optional list of possible values that you can use to filter
// results for tagged resources.
type TagFilter struct {

	// The tag key. This must have a valid string value and can't be empty.
	TagKey *string

	// A list of zero or more tag values. If no values are provided, then the filter
	// matches any tag with the specified key, regardless of its value.
	TagValues []string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
