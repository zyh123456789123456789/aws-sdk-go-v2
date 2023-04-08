// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Starts an export of DB snapshot or DB cluster data to Amazon S3. The provided
// IAM role must have access to the S3 bucket. You can't export snapshot data from
// RDS Custom DB instances. You can't export cluster data from Multi-AZ DB
// clusters. For more information on exporting DB snapshot data, see Exporting DB
// snapshot data to Amazon S3 (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ExportSnapshot.html)
// in the Amazon RDS User Guide or Exporting DB cluster snapshot data to Amazon S3 (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-export-snapshot.html)
// in the Amazon Aurora User Guide. For more information on exporting DB cluster
// data, see Exporting DB cluster data to Amazon S3 (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/export-cluster-data.html)
// in the Amazon Aurora User Guide.
func (c *Client) StartExportTask(ctx context.Context, params *StartExportTaskInput, optFns ...func(*Options)) (*StartExportTaskOutput, error) {
	if params == nil {
		params = &StartExportTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartExportTask", params, optFns, c.addOperationStartExportTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartExportTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartExportTaskInput struct {

	// A unique identifier for the export task. This ID isn't an identifier for the
	// Amazon S3 bucket where the data is to be exported.
	//
	// This member is required.
	ExportTaskIdentifier *string

	// The name of the IAM role to use for writing to the Amazon S3 bucket when
	// exporting a snapshot or cluster. In the IAM policy attached to your IAM role,
	// include the following required actions to allow the transfer of files from
	// Amazon RDS or Amazon Aurora to an S3 bucket:
	//   - s3:PutObject*
	//   - s3:GetObject*
	//   - s3:ListBucket
	//   - s3:DeleteObject*
	//   - s3:GetBucketLocation
	// In the policy, include the resources to identify the S3 bucket and objects in
	// the bucket. The following list of resources shows the Amazon Resource Name (ARN)
	// format for accessing S3:
	//   - arn:aws:s3:::your-s3-bucket
	//   - arn:aws:s3:::your-s3-bucket/*
	//
	// This member is required.
	IamRoleArn *string

	// The ID of the Amazon Web Services KMS key to use to encrypt the data exported
	// to Amazon S3. The Amazon Web Services KMS key identifier is the key ARN, key ID,
	// alias ARN, or alias name for the KMS key. The caller of this operation must be
	// authorized to run the following operations. These can be set in the Amazon Web
	// Services KMS key policy:
	//   - kms:Encrypt
	//   - kms:Decrypt
	//   - kms:GenerateDataKey
	//   - kms:GenerateDataKeyWithoutPlaintext
	//   - kms:ReEncryptFrom
	//   - kms:ReEncryptTo
	//   - kms:CreateGrant
	//   - kms:DescribeKey
	//   - kms:RetireGrant
	//
	// This member is required.
	KmsKeyId *string

	// The name of the Amazon S3 bucket to export the snapshot or cluster data to.
	//
	// This member is required.
	S3BucketName *string

	// The Amazon Resource Name (ARN) of the snapshot or cluster to export to Amazon
	// S3.
	//
	// This member is required.
	SourceArn *string

	// The data to be exported from the snapshot or cluster. If this parameter is not
	// provided, all of the data is exported. Valid values are the following:
	//   - database - Export all the data from a specified database.
	//   - database.table table-name - Export a table of the snapshot or cluster. This
	//   format is valid only for RDS for MySQL, RDS for MariaDB, and Aurora MySQL.
	//   - database.schema schema-name - Export a database schema of the snapshot or
	//   cluster. This format is valid only for RDS for PostgreSQL and Aurora PostgreSQL.
	//
	//   - database.schema.table table-name - Export a table of the database schema.
	//   This format is valid only for RDS for PostgreSQL and Aurora PostgreSQL.
	ExportOnly []string

	// The Amazon S3 bucket prefix to use as the file name and path of the exported
	// data.
	S3Prefix *string

	noSmithyDocumentSerde
}

// Contains the details of a snapshot or cluster export to Amazon S3. This data
// type is used as a response element in the DescribeExportTasks action.
type StartExportTaskOutput struct {

	// The data exported from the snapshot or cluster. Valid values are the following:
	//   - database - Export all the data from a specified database.
	//   - database.table table-name - Export a table of the snapshot or cluster. This
	//   format is valid only for RDS for MySQL, RDS for MariaDB, and Aurora MySQL.
	//   - database.schema schema-name - Export a database schema of the snapshot or
	//   cluster. This format is valid only for RDS for PostgreSQL and Aurora PostgreSQL.
	//
	//   - database.schema.table table-name - Export a table of the database schema.
	//   This format is valid only for RDS for PostgreSQL and Aurora PostgreSQL.
	ExportOnly []string

	// A unique identifier for the snapshot or cluster export task. This ID isn't an
	// identifier for the Amazon S3 bucket where the data is exported.
	ExportTaskIdentifier *string

	// The reason the export failed, if it failed.
	FailureCause *string

	// The name of the IAM role that is used to write to Amazon S3 when exporting a
	// snapshot or cluster.
	IamRoleArn *string

	// The key identifier of the Amazon Web Services KMS key that is used to encrypt
	// the data when it's exported to Amazon S3. The KMS key identifier is its key ARN,
	// key ID, alias ARN, or alias name. The IAM role used for the export must have
	// encryption and decryption permissions to use this KMS key.
	KmsKeyId *string

	// The progress of the snapshot or cluster export task as a percentage.
	PercentProgress int32

	// The Amazon S3 bucket that the snapshot or cluster is exported to.
	S3Bucket *string

	// The Amazon S3 bucket prefix that is the file name and path of the exported data.
	S3Prefix *string

	// The time that the snapshot was created.
	SnapshotTime *time.Time

	// The Amazon Resource Name (ARN) of the snapshot or cluster exported to Amazon S3.
	SourceArn *string

	// The type of source for the export.
	SourceType types.ExportSourceType

	// The progress status of the export task. The status can be one of the following:
	//   - CANCELED
	//   - CANCELING
	//   - COMPLETE
	//   - FAILED
	//   - IN_PROGRESS
	//   - STARTING
	Status *string

	// The time that the snapshot or cluster export task ended.
	TaskEndTime *time.Time

	// The time that the snapshot or cluster export task started.
	TaskStartTime *time.Time

	// The total amount of data exported, in gigabytes.
	TotalExtractedDataInGB int32

	// A warning about the snapshot or cluster export task.
	WarningMessage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartExportTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpStartExportTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpStartExportTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpStartExportTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartExportTask(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opStartExportTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "StartExportTask",
	}
}
