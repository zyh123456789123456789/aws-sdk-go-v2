// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Initiates the failover process for a Neptune global database. A failover for a
// Neptune global database promotes one of secondary read-only DB clusters to be
// the primary DB cluster and demotes the primary DB cluster to being a secondary
// (read-only) DB cluster. In other words, the role of the current primary DB
// cluster and the selected target secondary DB cluster are switched. The selected
// secondary DB cluster assumes full read/write capabilities for the Neptune global
// database. This action applies only to Neptune global databases. This action is
// only intended for use on healthy Neptune global databases with healthy Neptune
// DB clusters and no region-wide outages, to test disaster recovery scenarios or
// to reconfigure the global database topology.
func (c *Client) FailoverGlobalCluster(ctx context.Context, params *FailoverGlobalClusterInput, optFns ...func(*Options)) (*FailoverGlobalClusterOutput, error) {
	if params == nil {
		params = &FailoverGlobalClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "FailoverGlobalCluster", params, optFns, c.addOperationFailoverGlobalClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*FailoverGlobalClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type FailoverGlobalClusterInput struct {

	// Identifier of the Neptune global database that should be failed over. The
	// identifier is the unique key assigned by the user when the Neptune global
	// database was created. In other words, it's the name of the global database that
	// you want to fail over. Constraints: Must match the identifier of an existing
	// Neptune global database.
	//
	// This member is required.
	GlobalClusterIdentifier *string

	// The Amazon Resource Name (ARN) of the secondary Neptune DB cluster that you
	// want to promote to primary for the global database.
	//
	// This member is required.
	TargetDbClusterIdentifier *string

	noSmithyDocumentSerde
}

type FailoverGlobalClusterOutput struct {

	// Contains the details of an Amazon Neptune global database. This data type is
	// used as a response element for the CreateGlobalCluster , DescribeGlobalClusters
	// , ModifyGlobalCluster , DeleteGlobalCluster , FailoverGlobalCluster , and
	// RemoveFromGlobalCluster actions.
	GlobalCluster *types.GlobalCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationFailoverGlobalClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpFailoverGlobalCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpFailoverGlobalCluster{}, middleware.After)
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
	if err = addOpFailoverGlobalClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opFailoverGlobalCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opFailoverGlobalCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "FailoverGlobalCluster",
	}
}
