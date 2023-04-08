// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List all assets ( DASHBOARD , ANALYSIS , and DATASET ) in a folder.
func (c *Client) ListFolderMembers(ctx context.Context, params *ListFolderMembersInput, optFns ...func(*Options)) (*ListFolderMembersOutput, error) {
	if params == nil {
		params = &ListFolderMembersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFolderMembers", params, optFns, c.addOperationListFolderMembersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFolderMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFolderMembersInput struct {

	// The ID for the Amazon Web Services account that contains the folder.
	//
	// This member is required.
	AwsAccountId *string

	// The ID of the folder.
	//
	// This member is required.
	FolderId *string

	// The maximum number of results to be returned per request.
	MaxResults *int32

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListFolderMembersOutput struct {

	// A structure that contains all of the folder members (dashboards, analyses, and
	// datasets) in the folder.
	FolderMemberList []types.MemberIdArnPair

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFolderMembersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListFolderMembers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListFolderMembers{}, middleware.After)
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
	if err = addOpListFolderMembersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFolderMembers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListFolderMembers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "ListFolderMembers",
	}
}
