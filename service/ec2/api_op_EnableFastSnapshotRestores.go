// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables fast snapshot restores for the specified snapshots in the specified
// Availability Zones. You get the full benefit of fast snapshot restores after
// they enter the enabled state. To get the current state of fast snapshot
// restores, use DescribeFastSnapshotRestores . To disable fast snapshot restores,
// use DisableFastSnapshotRestores . For more information, see Amazon EBS fast
// snapshot restore (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-fast-snapshot-restore.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) EnableFastSnapshotRestores(ctx context.Context, params *EnableFastSnapshotRestoresInput, optFns ...func(*Options)) (*EnableFastSnapshotRestoresOutput, error) {
	if params == nil {
		params = &EnableFastSnapshotRestoresInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableFastSnapshotRestores", params, optFns, c.addOperationEnableFastSnapshotRestoresMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableFastSnapshotRestoresOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableFastSnapshotRestoresInput struct {

	// One or more Availability Zones. For example, us-east-2a .
	//
	// This member is required.
	AvailabilityZones []string

	// The IDs of one or more snapshots. For example, snap-1234567890abcdef0 . You can
	// specify a snapshot that was shared with you from another Amazon Web Services
	// account.
	//
	// This member is required.
	SourceSnapshotIds []string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

type EnableFastSnapshotRestoresOutput struct {

	// Information about the snapshots for which fast snapshot restores were
	// successfully enabled.
	Successful []types.EnableFastSnapshotRestoreSuccessItem

	// Information about the snapshots for which fast snapshot restores could not be
	// enabled.
	Unsuccessful []types.EnableFastSnapshotRestoreErrorItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnableFastSnapshotRestoresMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpEnableFastSnapshotRestores{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpEnableFastSnapshotRestores{}, middleware.After)
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
	if err = addOpEnableFastSnapshotRestoresValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableFastSnapshotRestores(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnableFastSnapshotRestores(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "EnableFastSnapshotRestores",
	}
}
