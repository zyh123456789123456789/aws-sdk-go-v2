// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of Lambda functions, with the version-specific configuration of
// each. Lambda returns up to 50 functions per call. Set FunctionVersion to ALL to
// include all published versions of each function in addition to the unpublished
// version. The ListFunctions operation returns a subset of the
// FunctionConfiguration fields. To get the additional fields (State,
// StateReasonCode, StateReason, LastUpdateStatus, LastUpdateStatusReason,
// LastUpdateStatusReasonCode, RuntimeVersionConfig) for a function or version, use
// GetFunction .
func (c *Client) ListFunctions(ctx context.Context, params *ListFunctionsInput, optFns ...func(*Options)) (*ListFunctionsOutput, error) {
	if params == nil {
		params = &ListFunctionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFunctions", params, optFns, c.addOperationListFunctionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFunctionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFunctionsInput struct {

	// Set to ALL to include entries for all published versions of each function.
	FunctionVersion types.FunctionVersion

	// Specify the pagination token that's returned by a previous request to retrieve
	// the next page of results.
	Marker *string

	// For Lambda@Edge functions, the Amazon Web Services Region of the master
	// function. For example, us-east-1 filters the list of functions to include only
	// Lambda@Edge functions replicated from a master function in US East (N.
	// Virginia). If specified, you must set FunctionVersion to ALL .
	MasterRegion *string

	// The maximum number of functions to return in the response. Note that
	// ListFunctions returns a maximum of 50 items in each response, even if you set
	// the number higher.
	MaxItems *int32

	noSmithyDocumentSerde
}

// A list of Lambda functions.
type ListFunctionsOutput struct {

	// A list of Lambda functions.
	Functions []types.FunctionConfiguration

	// The pagination token that's included if more results are available.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFunctionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListFunctions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListFunctions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFunctions(options.Region), middleware.Before); err != nil {
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

// ListFunctionsAPIClient is a client that implements the ListFunctions operation.
type ListFunctionsAPIClient interface {
	ListFunctions(context.Context, *ListFunctionsInput, ...func(*Options)) (*ListFunctionsOutput, error)
}

var _ ListFunctionsAPIClient = (*Client)(nil)

// ListFunctionsPaginatorOptions is the paginator options for ListFunctions
type ListFunctionsPaginatorOptions struct {
	// The maximum number of functions to return in the response. Note that
	// ListFunctions returns a maximum of 50 items in each response, even if you set
	// the number higher.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFunctionsPaginator is a paginator for ListFunctions
type ListFunctionsPaginator struct {
	options   ListFunctionsPaginatorOptions
	client    ListFunctionsAPIClient
	params    *ListFunctionsInput
	nextToken *string
	firstPage bool
}

// NewListFunctionsPaginator returns a new ListFunctionsPaginator
func NewListFunctionsPaginator(client ListFunctionsAPIClient, params *ListFunctionsInput, optFns ...func(*ListFunctionsPaginatorOptions)) *ListFunctionsPaginator {
	if params == nil {
		params = &ListFunctionsInput{}
	}

	options := ListFunctionsPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFunctionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFunctionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFunctions page.
func (p *ListFunctionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFunctionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.ListFunctions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListFunctions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "ListFunctions",
	}
}
