// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// A data type pair that consists of a KeyName and Values list that is used in
// conjunction with the KeyName (https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_SearchProfiles.html#customerprofiles-SearchProfiles-request-KeyName)
// and Values (https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_SearchProfiles.html#customerprofiles-SearchProfiles-request-Values)
// parameters to search for profiles using the SearchProfiles (https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_SearchProfiles.html)
// API.
type AdditionalSearchKey struct {

	// A searchable identifier of a customer profile.
	//
	// This member is required.
	KeyName *string

	// A list of key values.
	//
	// This member is required.
	Values []string

	noSmithyDocumentSerde
}

// A generic address associated with the customer that is not mailing, shipping,
// or billing.
type Address struct {

	// The first line of a customer address.
	Address1 *string

	// The second line of a customer address.
	Address2 *string

	// The third line of a customer address.
	Address3 *string

	// The fourth line of a customer address.
	Address4 *string

	// The city in which a customer lives.
	City *string

	// The country in which a customer lives.
	Country *string

	// The county in which a customer lives.
	County *string

	// The postal code of a customer address.
	PostalCode *string

	// The province in which a customer lives.
	Province *string

	// The state in which a customer lives.
	State *string

	noSmithyDocumentSerde
}

// Details for workflow of type APPFLOW_INTEGRATION .
type AppflowIntegration struct {

	// The configurations that control how Customer Profiles retrieves data from the
	// source, Amazon AppFlow. Customer Profiles uses this information to create an
	// AppFlow flow on behalf of customers.
	//
	// This member is required.
	FlowDefinition *FlowDefinition

	// Batches in workflow of type APPFLOW_INTEGRATION .
	Batches []Batch

	noSmithyDocumentSerde
}

// Structure holding all APPFLOW_INTEGRATION specific workflow attributes.
type AppflowIntegrationWorkflowAttributes struct {

	// The name of the AppFlow connector profile used for ingestion.
	//
	// This member is required.
	ConnectorProfileName *string

	// Specifies the source connector type, such as Salesforce, ServiceNow, and
	// Marketo. Indicates source of ingestion.
	//
	// This member is required.
	SourceConnectorType SourceConnectorType

	// The Amazon Resource Name (ARN) of the IAM role. Customer Profiles assumes this
	// role to create resources on your behalf as part of workflow execution.
	RoleArn *string

	noSmithyDocumentSerde
}

// Workflow specific execution metrics for APPFLOW_INTEGRATION workflow.
type AppflowIntegrationWorkflowMetrics struct {

	// Number of records processed in APPFLOW_INTEGRATION workflow.
	//
	// This member is required.
	RecordsProcessed int64

	// Total steps completed in APPFLOW_INTEGRATION workflow.
	//
	// This member is required.
	StepsCompleted int64

	// Total steps in APPFLOW_INTEGRATION workflow.
	//
	// This member is required.
	TotalSteps int64

	noSmithyDocumentSerde
}

// Workflow step details for APPFLOW_INTEGRATION workflow.
type AppflowIntegrationWorkflowStep struct {

	// End datetime of records pulled in batch during execution of workflow step for
	// APPFLOW_INTEGRATION workflow.
	//
	// This member is required.
	BatchRecordsEndTime *string

	// Start datetime of records pulled in batch during execution of workflow step for
	// APPFLOW_INTEGRATION workflow.
	//
	// This member is required.
	BatchRecordsStartTime *string

	// Creation timestamp of workflow step for APPFLOW_INTEGRATION workflow.
	//
	// This member is required.
	CreatedAt *time.Time

	// Message indicating execution of workflow step for APPFLOW_INTEGRATION workflow.
	//
	// This member is required.
	ExecutionMessage *string

	// Name of the flow created during execution of workflow step. APPFLOW_INTEGRATION
	// workflow type creates an appflow flow during workflow step execution on the
	// customers behalf.
	//
	// This member is required.
	FlowName *string

	// Last updated timestamp for workflow step for APPFLOW_INTEGRATION workflow.
	//
	// This member is required.
	LastUpdatedAt *time.Time

	// Total number of records processed during execution of workflow step for
	// APPFLOW_INTEGRATION workflow.
	//
	// This member is required.
	RecordsProcessed int64

	// Workflow step status for APPFLOW_INTEGRATION workflow.
	//
	// This member is required.
	Status Status

	noSmithyDocumentSerde
}

// Configuration settings for how to perform the auto-merging of profiles.
type AutoMerging struct {

	// The flag that enables the auto-merging of duplicate profiles.
	//
	// This member is required.
	Enabled *bool

	// How the auto-merging process should resolve conflicts between different
	// profiles. For example, if Profile A and Profile B have the same FirstName and
	// LastName (and that is the matching criteria), which EmailAddress should be used?
	ConflictResolution *ConflictResolution

	// A list of matching attributes that represent matching criteria. If two profiles
	// meet at least one of the requirements in the matching attributes list, they will
	// be merged.
	Consolidation *Consolidation

	// A number between 0 and 1 that represents the minimum confidence score required
	// for profiles within a matching group to be merged during the auto-merge process.
	// A higher score means higher similarity required to merge profiles.
	MinAllowedConfidenceScoreForMerging *float64

	noSmithyDocumentSerde
}

// Batch defines the boundaries for ingestion for each step in APPFLOW_INTEGRATION
// workflow. APPFLOW_INTEGRATION workflow splits ingestion based on these
// boundaries.
type Batch struct {

	// End time of batch to split ingestion.
	//
	// This member is required.
	EndTime *time.Time

	// Start time of batch to split ingestion.
	//
	// This member is required.
	StartTime *time.Time

	noSmithyDocumentSerde
}

// How the auto-merging process should resolve conflicts between different
// profiles.
type ConflictResolution struct {

	// How the auto-merging process should resolve conflicts between different
	// profiles.
	//   - RECENCY : Uses the data that was most recently updated.
	//   - SOURCE : Uses the data from a specific source. For example, if a company has
	//   been aquired or two departments have merged, data from the specified source is
	//   used. If two duplicate profiles are from the same source, then RECENCY is used
	//   again.
	//
	// This member is required.
	ConflictResolvingModel ConflictResolvingModel

	// The ObjectType name that is used to resolve profile merging conflicts when
	// choosing SOURCE as the ConflictResolvingModel .
	SourceName *string

	noSmithyDocumentSerde
}

// The operation to be performed on the provided source fields.
type ConnectorOperator struct {

	// The operation to be performed on the provided Marketo source fields.
	Marketo MarketoConnectorOperator

	// The operation to be performed on the provided Amazon S3 source fields.
	S3 S3ConnectorOperator

	// The operation to be performed on the provided Salesforce source fields.
	Salesforce SalesforceConnectorOperator

	// The operation to be performed on the provided ServiceNow source fields.
	ServiceNow ServiceNowConnectorOperator

	// The operation to be performed on the provided Zendesk source fields.
	Zendesk ZendeskConnectorOperator

	noSmithyDocumentSerde
}

// The matching criteria to be used during the auto-merging process.
type Consolidation struct {

	// A list of matching criteria.
	//
	// This member is required.
	MatchingAttributesList [][]string

	noSmithyDocumentSerde
}

// Usage-specific statistics about the domain.
type DomainStats struct {

	// The number of profiles that you are currently paying for in the domain. If you
	// have more than 100 objects associated with a single profile, that profile counts
	// as two profiles. If you have more than 200 objects, that profile counts as
	// three, and so on.
	MeteringProfileCount int64

	// The total number of objects in domain.
	ObjectCount int64

	// The total number of profiles currently in the domain.
	ProfileCount int64

	// The total size, in bytes, of all objects in the domain.
	TotalSize int64

	noSmithyDocumentSerde
}

// Configuration information about the S3 bucket where Identity Resolution Jobs
// writes result files. You need to give Customer Profiles service principal write
// permission to your S3 bucket. Otherwise, you'll get an exception in the API
// response. For an example policy, see Amazon Connect Customer Profiles
// cross-service confused deputy prevention (https://docs.aws.amazon.com/connect/latest/adminguide/cross-service-confused-deputy-prevention.html#customer-profiles-cross-service)
// .
type ExportingConfig struct {

	// The S3 location where Identity Resolution Jobs write result files.
	S3Exporting *S3ExportingConfig

	noSmithyDocumentSerde
}

// The S3 location where Identity Resolution Jobs write result files.
type ExportingLocation struct {

	// Information about the S3 location where Identity Resolution Jobs write result
	// files.
	S3Exporting *S3ExportingLocation

	noSmithyDocumentSerde
}

// A duplicate customer profile that is to be merged into a main profile.
type FieldSourceProfileIds struct {

	// A unique identifier for the account number field to be merged.
	AccountNumber *string

	// A unique identifier for the additional information field to be merged.
	AdditionalInformation *string

	// A unique identifier for the party type field to be merged.
	Address *string

	// A unique identifier for the attributes field to be merged.
	Attributes map[string]string

	// A unique identifier for the billing type field to be merged.
	BillingAddress *string

	// A unique identifier for the birthdate field to be merged.
	BirthDate *string

	// A unique identifier for the party type field to be merged.
	BusinessEmailAddress *string

	// A unique identifier for the business name field to be merged.
	BusinessName *string

	// A unique identifier for the business phone number field to be merged.
	BusinessPhoneNumber *string

	// A unique identifier for the email address field to be merged.
	EmailAddress *string

	// A unique identifier for the first name field to be merged.
	FirstName *string

	// A unique identifier for the gender field to be merged.
	Gender *string

	// A unique identifier for the home phone number field to be merged.
	HomePhoneNumber *string

	// A unique identifier for the last name field to be merged.
	LastName *string

	// A unique identifier for the mailing address field to be merged.
	MailingAddress *string

	// A unique identifier for the middle name field to be merged.
	MiddleName *string

	// A unique identifier for the mobile phone number field to be merged.
	MobilePhoneNumber *string

	// A unique identifier for the party type field to be merged.
	PartyType *string

	// A unique identifier for the personal email address field to be merged.
	PersonalEmailAddress *string

	// A unique identifier for the phone number field to be merged.
	PhoneNumber *string

	// A unique identifier for the shipping address field to be merged.
	ShippingAddress *string

	noSmithyDocumentSerde
}

// The configurations that control how Customer Profiles retrieves data from the
// source, Amazon AppFlow. Customer Profiles uses this information to create an
// AppFlow flow on behalf of customers.
type FlowDefinition struct {

	// The specified name of the flow. Use underscores (_) or hyphens (-) only. Spaces
	// are not allowed.
	//
	// This member is required.
	FlowName *string

	// The Amazon Resource Name of the AWS Key Management Service (KMS) key you
	// provide for encryption.
	//
	// This member is required.
	KmsArn *string

	// The configuration that controls how Customer Profiles retrieves data from the
	// source.
	//
	// This member is required.
	SourceFlowConfig *SourceFlowConfig

	// A list of tasks that Customer Profiles performs while transferring the data in
	// the flow run.
	//
	// This member is required.
	Tasks []Task

	// The trigger settings that determine how and when the flow runs.
	//
	// This member is required.
	TriggerConfig *TriggerConfig

	// A description of the flow you want to create.
	Description *string

	noSmithyDocumentSerde
}

// A data type pair that consists of a KeyName and Values list that were used to
// find a profile returned in response to a SearchProfiles (https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_SearchProfiles.html)
// request.
type FoundByKeyValue struct {

	// A searchable identifier of a customer profile.
	KeyName *string

	// A list of key values.
	Values []string

	noSmithyDocumentSerde
}

// Information about the Identity Resolution Job.
type IdentityResolutionJob struct {

	// The unique name of the domain.
	DomainName *string

	// The S3 location where the Identity Resolution Job writes result files.
	ExportingLocation *ExportingLocation

	// The timestamp of when the job was completed.
	JobEndTime *time.Time

	// The unique identifier of the Identity Resolution Job.
	JobId *string

	// The timestamp of when the job was started or will be started.
	JobStartTime *time.Time

	// Statistics about an Identity Resolution Job.
	JobStats *JobStats

	// The error messages that are generated when the Identity Resolution Job runs.
	Message *string

	// The status of the Identity Resolution Job.
	//   - PENDING : The Identity Resolution Job is scheduled but has not started yet.
	//   If you turn off the Identity Resolution feature in your domain, jobs in the
	//   PENDING state are deleted.
	//   - PREPROCESSING : The Identity Resolution Job is loading your data.
	//   - FIND_MATCHING : The Identity Resolution Job is using the machine learning
	//   model to identify profiles that belong to the same matching group.
	//   - MERGING : The Identity Resolution Job is merging duplicate profiles.
	//   - COMPLETED : The Identity Resolution Job completed successfully.
	//   - PARTIAL_SUCCESS : There's a system error and not all of the data is merged.
	//   The Identity Resolution Job writes a message indicating the source of the
	//   problem.
	//   - FAILED : The Identity Resolution Job did not merge any data. It writes a
	//   message indicating the source of the problem.
	Status IdentityResolutionJobStatus

	noSmithyDocumentSerde
}

// Specifies the configuration used when importing incremental records from the
// source.
type IncrementalPullConfig struct {

	// A field that specifies the date time or timestamp field as the criteria to use
	// when importing incremental records from the source.
	DatetimeTypeFieldName *string

	noSmithyDocumentSerde
}

// Configuration data for integration workflow.
type IntegrationConfig struct {

	// Configuration data for APPFLOW_INTEGRATION workflow type.
	AppflowIntegration *AppflowIntegration

	noSmithyDocumentSerde
}

// The day and time when do you want to start the Identity Resolution Job every
// week.
type JobSchedule struct {

	// The day when the Identity Resolution Job should run every week.
	//
	// This member is required.
	DayOfTheWeek JobScheduleDayOfTheWeek

	// The time when the Identity Resolution Job should run every week.
	//
	// This member is required.
	Time *string

	noSmithyDocumentSerde
}

// Statistics about the Identity Resolution Job.
type JobStats struct {

	// The number of matches found.
	NumberOfMatchesFound int64

	// The number of merges completed.
	NumberOfMergesDone int64

	// The number of profiles reviewed.
	NumberOfProfilesReviewed int64

	noSmithyDocumentSerde
}

// An object in a list that represents a domain.
type ListDomainItem struct {

	// The timestamp of when the domain was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// The timestamp of when the domain was most recently edited.
	//
	// This member is required.
	LastUpdatedAt *time.Time

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

// An integration in list of integrations.
type ListIntegrationItem struct {

	// The timestamp of when the domain was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// The timestamp of when the domain was most recently edited.
	//
	// This member is required.
	LastUpdatedAt *time.Time

	// The URI of the S3 bucket or any other type of data source.
	//
	// This member is required.
	Uri *string

	// Boolean that shows if the Flow that's associated with the Integration is
	// created in Amazon Appflow, or with ObjectTypeName equals _unstructured via
	// API/CLI in flowDefinition.
	IsUnstructured *bool

	// The name of the profile object type.
	ObjectTypeName *string

	// A map in which each key is an event type from an external application such as
	// Segment or Shopify, and each value is an ObjectTypeName (template) used to
	// ingest the event. It supports the following event types: SegmentIdentify ,
	// ShopifyCreateCustomers , ShopifyUpdateCustomers , ShopifyCreateDraftOrders ,
	// ShopifyUpdateDraftOrders , ShopifyCreateOrders , and ShopifyUpdatedOrders .
	ObjectTypeNames map[string]string

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	// Unique identifier for the workflow.
	WorkflowId *string

	noSmithyDocumentSerde
}

// A ProfileObject in a list of ProfileObjects.
type ListProfileObjectsItem struct {

	// A JSON representation of a ProfileObject that belongs to a profile.
	Object *string

	// Specifies the kind of object being added to a profile, such as
	// "Salesforce-Account."
	ObjectTypeName *string

	// The unique identifier of the ProfileObject generated by the service.
	ProfileObjectUniqueKey *string

	noSmithyDocumentSerde
}

// A ProfileObjectType instance.
type ListProfileObjectTypeItem struct {

	// Description of the profile object type.
	//
	// This member is required.
	Description *string

	// The name of the profile object type.
	//
	// This member is required.
	ObjectTypeName *string

	// The timestamp of when the domain was created.
	CreatedAt *time.Time

	// The timestamp of when the domain was most recently edited.
	LastUpdatedAt *time.Time

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

// A ProfileObjectTypeTemplate in a list of ProfileObjectTypeTemplates.
type ListProfileObjectTypeTemplateItem struct {

	// The name of the source of the object template.
	SourceName *string

	// The source of the object template.
	SourceObject *string

	// A unique identifier for the object template.
	TemplateId *string

	noSmithyDocumentSerde
}

// A workflow in list of workflows.
type ListWorkflowsItem struct {

	// Creation timestamp for workflow.
	//
	// This member is required.
	CreatedAt *time.Time

	// Last updated timestamp for workflow.
	//
	// This member is required.
	LastUpdatedAt *time.Time

	// Status of workflow execution.
	//
	// This member is required.
	Status Status

	// Description for workflow execution status.
	//
	// This member is required.
	StatusDescription *string

	// Unique identifier for the workflow.
	//
	// This member is required.
	WorkflowId *string

	// The type of workflow. The only supported value is APPFLOW_INTEGRATION.
	//
	// This member is required.
	WorkflowType WorkflowType

	noSmithyDocumentSerde
}

// The properties that are applied when Marketo is being used as a source.
type MarketoSourceProperties struct {

	// The object specified in the Marketo flow source.
	//
	// This member is required.
	Object *string

	noSmithyDocumentSerde
}

// The flag that enables the matching process of duplicate profiles.
type MatchingRequest struct {

	// The flag that enables the matching process of duplicate profiles.
	//
	// This member is required.
	Enabled *bool

	// Configuration information about the auto-merging process.
	AutoMerging *AutoMerging

	// Configuration information for exporting Identity Resolution results, for
	// example, to an S3 bucket.
	ExportingConfig *ExportingConfig

	// The day and time when do you want to start the Identity Resolution Job every
	// week.
	JobSchedule *JobSchedule

	noSmithyDocumentSerde
}

// The flag that enables the matching process of duplicate profiles.
type MatchingResponse struct {

	// Configuration information about the auto-merging process.
	AutoMerging *AutoMerging

	// The flag that enables the matching process of duplicate profiles.
	Enabled *bool

	// Configuration information for exporting Identity Resolution results, for
	// example, to an S3 bucket.
	ExportingConfig *ExportingConfig

	// The day and time when do you want to start the Identity Resolution Job every
	// week.
	JobSchedule *JobSchedule

	noSmithyDocumentSerde
}

// The Match group object.
type MatchItem struct {

	// A number between 0 and 1, where a higher score means higher similarity.
	// Examining match confidence scores lets you distinguish between groups of similar
	// records in which the system is highly confident (which you may decide to merge),
	// groups of similar records about which the system is uncertain (which you may
	// decide to have reviewed by a human), and groups of similar records that the
	// system deems to be unlikely (which you may decide to reject). Given confidence
	// scores vary as per the data input, it should not be used an absolute measure of
	// matching quality.
	ConfidenceScore *float64

	// The unique identifiers for this group of profiles that match.
	MatchId *string

	// A list of identifiers for profiles that match.
	ProfileIds []string

	noSmithyDocumentSerde
}

// The filter applied to ListProfileObjects response to include profile objects
// with the specified index values. This filter is only supported for
// ObjectTypeName _asset, _case and _order.
type ObjectFilter struct {

	// A searchable identifier of a standard profile object. The predefined keys you
	// can use to search for _asset include: _assetId, _assetName, _serialNumber. The
	// predefined keys you can use to search for _case include: _caseId. The predefined
	// keys you can use to search for _order include: _orderId.
	//
	// This member is required.
	KeyName *string

	// A list of key values.
	//
	// This member is required.
	Values []string

	noSmithyDocumentSerde
}

// Represents a field in a ProfileObjectType.
type ObjectTypeField struct {

	// The content type of the field. Used for determining equality when searching.
	ContentType FieldContentType

	// A field of a ProfileObject. For example: _source.FirstName, where “_source” is
	// a ProfileObjectType of a Zendesk user and “FirstName” is a field in that
	// ObjectType.
	Source *string

	// The location of the data in the standard ProfileObject model. For example:
	// _profile.Address.PostalCode.
	Target *string

	noSmithyDocumentSerde
}

// An object that defines the Key element of a ProfileObject. A Key is a special
// element that can be used to search for a customer profile.
type ObjectTypeKey struct {

	// The reference for the key name of the fields map.
	FieldNames []string

	// The types of keys that a ProfileObject can have. Each ProfileObject can have
	// only 1 UNIQUE key but multiple PROFILE keys. PROFILE, ASSET, CASE, or ORDER
	// means that this key can be used to tie an object to a PROFILE, ASSET, CASE, or
	// ORDER respectively. UNIQUE means that it can be used to uniquely identify an
	// object. If a key a is marked as SECONDARY, it will be used to search for
	// profiles after all other PROFILE keys have been searched. A LOOKUP_ONLY key is
	// only used to match a profile but is not persisted to be used for searching of
	// the profile. A NEW_ONLY key is only used if the profile does not already exist
	// before the object is ingested, otherwise it is only used for matching objects to
	// profiles.
	StandardIdentifiers []StandardIdentifier

	noSmithyDocumentSerde
}

// The standard profile of a customer.
type Profile struct {

	// A unique account number that you have given to the customer.
	AccountNumber *string

	// Any additional information relevant to the customer’s profile.
	AdditionalInformation *string

	// A generic address associated with the customer that is not mailing, shipping,
	// or billing.
	Address *Address

	// A key value pair of attributes of a customer profile.
	Attributes map[string]string

	// The customer’s billing address.
	BillingAddress *Address

	// The customer’s birth date.
	BirthDate *string

	// The customer’s business email address.
	BusinessEmailAddress *string

	// The name of the customer’s business.
	BusinessName *string

	// The customer’s home phone number.
	BusinessPhoneNumber *string

	// The customer’s email address, which has not been specified as a personal or
	// business address.
	EmailAddress *string

	// The customer’s first name.
	FirstName *string

	// A list of items used to find a profile returned in a SearchProfiles (https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_SearchProfiles.html)
	// response. An item is a key-value(s) pair that matches an attribute in the
	// profile. If the optional AdditionalSearchKeys parameter was included in the
	// SearchProfiles (https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_SearchProfiles.html)
	// request, the FoundByItems list should be interpreted based on the
	// LogicalOperator used in the request:
	//   - AND - The profile included in the response matched all of the search keys
	//   specified in the request. The FoundByItems will include all of the
	//   key-value(s) pairs that were specified in the request (as this is a requirement
	//   of AND search logic).
	//   - OR - The profile included in the response matched at least one of the search
	//   keys specified in the request. The FoundByItems will include each of the
	//   key-value(s) pairs that the profile was found by.
	// The OR relationship is the default behavior if the LogicalOperator parameter is
	// not included in the SearchProfiles (https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_SearchProfiles.html)
	// request.
	FoundByItems []FoundByKeyValue

	// The gender with which the customer identifies.
	//
	// Deprecated: This member has been deprecated.
	Gender Gender

	// An alternative to Gender which accepts any string as input.
	GenderString *string

	// The customer’s home phone number.
	HomePhoneNumber *string

	// The customer’s last name.
	LastName *string

	// The customer’s mailing address.
	MailingAddress *Address

	// The customer’s middle name.
	MiddleName *string

	// The customer’s mobile phone number.
	MobilePhoneNumber *string

	// The type of profile used to describe the customer.
	//
	// Deprecated: This member has been deprecated.
	PartyType PartyType

	// An alternative to PartyType which accepts any string as input.
	PartyTypeString *string

	// The customer’s personal email address.
	PersonalEmailAddress *string

	// The customer's phone number, which has not been specified as a mobile, home, or
	// business number.
	PhoneNumber *string

	// The unique identifier of a customer profile.
	ProfileId *string

	// The customer’s shipping address.
	ShippingAddress *Address

	noSmithyDocumentSerde
}

// Configuration information about the S3 bucket where Identity Resolution Jobs
// write result files.
type S3ExportingConfig struct {

	// The name of the S3 bucket where Identity Resolution Jobs write result files.
	//
	// This member is required.
	S3BucketName *string

	// The S3 key name of the location where Identity Resolution Jobs write result
	// files.
	S3KeyName *string

	noSmithyDocumentSerde
}

// The S3 location where Identity Resolution Jobs write result files.
type S3ExportingLocation struct {

	// The name of the S3 bucket name where Identity Resolution Jobs write result
	// files.
	S3BucketName *string

	// The S3 key name of the location where Identity Resolution Jobs write result
	// files.
	S3KeyName *string

	noSmithyDocumentSerde
}

// The properties that are applied when Amazon S3 is being used as the flow source.
type S3SourceProperties struct {

	// The Amazon S3 bucket name where the source files are stored.
	//
	// This member is required.
	BucketName *string

	// The object key for the Amazon S3 bucket in which the source files are stored.
	BucketPrefix *string

	noSmithyDocumentSerde
}

// The properties that are applied when Salesforce is being used as a source.
type SalesforceSourceProperties struct {

	// The object specified in the Salesforce flow source.
	//
	// This member is required.
	Object *string

	// The flag that enables dynamic fetching of new (recently added) fields in the
	// Salesforce objects while running a flow.
	EnableDynamicFieldUpdate bool

	// Indicates whether Amazon AppFlow includes deleted files in the flow run.
	IncludeDeletedRecords bool

	noSmithyDocumentSerde
}

// Specifies the configuration details of a scheduled-trigger flow that you
// define. Currently, these settings only apply to the scheduled-trigger type.
type ScheduledTriggerProperties struct {

	// The scheduling expression that determines the rate at which the schedule will
	// run, for example rate (5 minutes).
	//
	// This member is required.
	ScheduleExpression *string

	// Specifies whether a scheduled flow has an incremental data transfer or a
	// complete data transfer for each flow run.
	DataPullMode DataPullMode

	// Specifies the date range for the records to import from the connector in the
	// first flow run.
	FirstExecutionFrom *time.Time

	// Specifies the scheduled end time for a scheduled-trigger flow.
	ScheduleEndTime *time.Time

	// Specifies the optional offset that is added to the time interval for a
	// schedule-triggered flow.
	ScheduleOffset *int64

	// Specifies the scheduled start time for a scheduled-trigger flow.
	ScheduleStartTime *time.Time

	// Specifies the time zone used when referring to the date and time of a
	// scheduled-triggered flow, such as America/New_York.
	Timezone *string

	noSmithyDocumentSerde
}

// The properties that are applied when ServiceNow is being used as a source.
type ServiceNowSourceProperties struct {

	// The object specified in the ServiceNow flow source.
	//
	// This member is required.
	Object *string

	noSmithyDocumentSerde
}

// Specifies the information that is required to query a particular Amazon AppFlow
// connector. Customer Profiles supports Salesforce, Zendesk, Marketo, ServiceNow
// and Amazon S3.
type SourceConnectorProperties struct {

	// The properties that are applied when Marketo is being used as a source.
	Marketo *MarketoSourceProperties

	// The properties that are applied when Amazon S3 is being used as the flow source.
	S3 *S3SourceProperties

	// The properties that are applied when Salesforce is being used as a source.
	Salesforce *SalesforceSourceProperties

	// The properties that are applied when ServiceNow is being used as a source.
	ServiceNow *ServiceNowSourceProperties

	// The properties that are applied when using Zendesk as a flow source.
	Zendesk *ZendeskSourceProperties

	noSmithyDocumentSerde
}

// Contains information about the configuration of the source connector used in
// the flow.
type SourceFlowConfig struct {

	// The type of connector, such as Salesforce, Marketo, and so on.
	//
	// This member is required.
	ConnectorType SourceConnectorType

	// Specifies the information that is required to query a particular source
	// connector.
	//
	// This member is required.
	SourceConnectorProperties *SourceConnectorProperties

	// The name of the AppFlow connector profile. This name must be unique for each
	// connector profile in the AWS account.
	ConnectorProfileName *string

	// Defines the configuration for a scheduled incremental data pull. If a valid
	// configuration is provided, the fields specified in the configuration are used
	// when querying for the incremental data pull.
	IncrementalPullConfig *IncrementalPullConfig

	noSmithyDocumentSerde
}

// A class for modeling different type of tasks. Task implementation varies based
// on the TaskType.
type Task struct {

	// The source fields to which a particular task is applied.
	//
	// This member is required.
	SourceFields []string

	// Specifies the particular task implementation that Amazon AppFlow performs.
	//
	// This member is required.
	TaskType TaskType

	// The operation to be performed on the provided source fields.
	ConnectorOperator *ConnectorOperator

	// A field in a destination connector, or a field value against which Amazon
	// AppFlow validates a source field.
	DestinationField *string

	// A map used to store task-related information. The service looks for particular
	// information based on the TaskType.
	TaskProperties map[string]string

	noSmithyDocumentSerde
}

// The trigger settings that determine how and when Amazon AppFlow runs the
// specified flow.
type TriggerConfig struct {

	// Specifies the type of flow trigger. It can be OnDemand, Scheduled, or Event.
	//
	// This member is required.
	TriggerType TriggerType

	// Specifies the configuration details of a schedule-triggered flow that you
	// define. Currently, these settings only apply to the Scheduled trigger type.
	TriggerProperties *TriggerProperties

	noSmithyDocumentSerde
}

// Specifies the configuration details that control the trigger for a flow.
// Currently, these settings only apply to the Scheduled trigger type.
type TriggerProperties struct {

	// Specifies the configuration details of a schedule-triggered flow that you
	// define.
	Scheduled *ScheduledTriggerProperties

	noSmithyDocumentSerde
}

// Updates associated with the address properties of a customer profile.
type UpdateAddress struct {

	// The first line of a customer address.
	Address1 *string

	// The second line of a customer address.
	Address2 *string

	// The third line of a customer address.
	Address3 *string

	// The fourth line of a customer address.
	Address4 *string

	// The city in which a customer lives.
	City *string

	// The country in which a customer lives.
	Country *string

	// The county in which a customer lives.
	County *string

	// The postal code of a customer address.
	PostalCode *string

	// The province in which a customer lives.
	Province *string

	// The state in which a customer lives.
	State *string

	noSmithyDocumentSerde
}

// Structure to hold workflow attributes.
type WorkflowAttributes struct {

	// Workflow attributes specific to APPFLOW_INTEGRATION workflow.
	AppflowIntegration *AppflowIntegrationWorkflowAttributes

	noSmithyDocumentSerde
}

// Generic object containing workflow execution metrics.
type WorkflowMetrics struct {

	// Workflow execution metrics for APPFLOW_INTEGRATION workflow.
	AppflowIntegration *AppflowIntegrationWorkflowMetrics

	noSmithyDocumentSerde
}

// List containing steps in workflow.
type WorkflowStepItem struct {

	// Workflow step information specific to APPFLOW_INTEGRATION workflow.
	AppflowIntegration *AppflowIntegrationWorkflowStep

	noSmithyDocumentSerde
}

// The properties that are applied when using Zendesk as a flow source.
type ZendeskSourceProperties struct {

	// The object specified in the Zendesk flow source.
	//
	// This member is required.
	Object *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
