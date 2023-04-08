// Code generated by smithy-go-codegen DO NOT EDIT.

package tnb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/tnb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists details for a network operation, including when the operation started and
// the status of the operation. A network operation is any operation that is done
// to your network, such as network instance instantiation or termination.
func (c *Client) ListSolNetworkOperations(ctx context.Context, params *ListSolNetworkOperationsInput, optFns ...func(*Options)) (*ListSolNetworkOperationsOutput, error) {
	if params == nil {
		params = &ListSolNetworkOperationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSolNetworkOperations", params, optFns, c.addOperationListSolNetworkOperationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSolNetworkOperationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSolNetworkOperationsInput struct {

	// The maximum number of results to include in the response.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSolNetworkOperationsOutput struct {

	// Lists network operation occurrences. Lifecycle management operations are
	// deploy, update, or delete operations.
	NetworkOperations []types.ListSolNetworkOperationsInfo

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSolNetworkOperationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSolNetworkOperations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSolNetworkOperations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSolNetworkOperations(options.Region), middleware.Before); err != nil {
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

// ListSolNetworkOperationsAPIClient is a client that implements the
// ListSolNetworkOperations operation.
type ListSolNetworkOperationsAPIClient interface {
	ListSolNetworkOperations(context.Context, *ListSolNetworkOperationsInput, ...func(*Options)) (*ListSolNetworkOperationsOutput, error)
}

var _ ListSolNetworkOperationsAPIClient = (*Client)(nil)

// ListSolNetworkOperationsPaginatorOptions is the paginator options for
// ListSolNetworkOperations
type ListSolNetworkOperationsPaginatorOptions struct {
	// The maximum number of results to include in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSolNetworkOperationsPaginator is a paginator for ListSolNetworkOperations
type ListSolNetworkOperationsPaginator struct {
	options   ListSolNetworkOperationsPaginatorOptions
	client    ListSolNetworkOperationsAPIClient
	params    *ListSolNetworkOperationsInput
	nextToken *string
	firstPage bool
}

// NewListSolNetworkOperationsPaginator returns a new
// ListSolNetworkOperationsPaginator
func NewListSolNetworkOperationsPaginator(client ListSolNetworkOperationsAPIClient, params *ListSolNetworkOperationsInput, optFns ...func(*ListSolNetworkOperationsPaginatorOptions)) *ListSolNetworkOperationsPaginator {
	if params == nil {
		params = &ListSolNetworkOperationsInput{}
	}

	options := ListSolNetworkOperationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSolNetworkOperationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSolNetworkOperationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSolNetworkOperations page.
func (p *ListSolNetworkOperationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSolNetworkOperationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListSolNetworkOperations(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListSolNetworkOperations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "tnb",
		OperationName: "ListSolNetworkOperations",
	}
}
