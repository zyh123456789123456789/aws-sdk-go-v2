// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworkscm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of tags that are applied to the specified AWS OpsWorks for Chef
// Automate or AWS OpsWorks for Puppet Enterprise servers or backups.
func (c *Client) ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) {
	if params == nil {
		params = &ListTagsForResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTagsForResource", params, optFns, c.addOperationListTagsForResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTagsForResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTagsForResourceInput struct {

	// The Amazon Resource Number (ARN) of an AWS OpsWorks for Chef Automate or AWS
	// OpsWorks for Puppet Enterprise server for which you want to show applied tags.
	// For example,
	// arn:aws:opsworks-cm:us-west-2:123456789012:server/test-owcm-server/EXAMPLE-66b0-4196-8274-d1a2bEXAMPLE
	// .
	//
	// This member is required.
	ResourceArn *string

	// To receive a paginated response, use this parameter to specify the maximum
	// number of results to be returned with a single call. If the number of available
	// results exceeds this maximum, the response includes a NextToken value that you
	// can assign to the NextToken request parameter to get the next set of results.
	MaxResults *int32

	// NextToken is a string that is returned in some command responses. It indicates
	// that not all entries have been returned, and that you must run at least one more
	// request to get remaining items. To get remaining results, call
	// ListTagsForResource again, and assign the token from the previous results as the
	// value of the nextToken parameter. If there are no more results, the response
	// object's nextToken parameter value is null . Setting a nextToken value that was
	// not returned in your previous results causes an InvalidNextTokenException to
	// occur.
	NextToken *string

	noSmithyDocumentSerde
}

type ListTagsForResourceOutput struct {

	// A token that you can use as the value of NextToken in subsequent calls to the
	// API to show more results.
	NextToken *string

	// Tags that have been applied to the resource.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTagsForResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTagsForResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTagsForResource{}, middleware.After)
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
	if err = addOpListTagsForResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTagsForResource(options.Region), middleware.Before); err != nil {
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

// ListTagsForResourceAPIClient is a client that implements the
// ListTagsForResource operation.
type ListTagsForResourceAPIClient interface {
	ListTagsForResource(context.Context, *ListTagsForResourceInput, ...func(*Options)) (*ListTagsForResourceOutput, error)
}

var _ ListTagsForResourceAPIClient = (*Client)(nil)

// ListTagsForResourcePaginatorOptions is the paginator options for
// ListTagsForResource
type ListTagsForResourcePaginatorOptions struct {
	// To receive a paginated response, use this parameter to specify the maximum
	// number of results to be returned with a single call. If the number of available
	// results exceeds this maximum, the response includes a NextToken value that you
	// can assign to the NextToken request parameter to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTagsForResourcePaginator is a paginator for ListTagsForResource
type ListTagsForResourcePaginator struct {
	options   ListTagsForResourcePaginatorOptions
	client    ListTagsForResourceAPIClient
	params    *ListTagsForResourceInput
	nextToken *string
	firstPage bool
}

// NewListTagsForResourcePaginator returns a new ListTagsForResourcePaginator
func NewListTagsForResourcePaginator(client ListTagsForResourceAPIClient, params *ListTagsForResourceInput, optFns ...func(*ListTagsForResourcePaginatorOptions)) *ListTagsForResourcePaginator {
	if params == nil {
		params = &ListTagsForResourceInput{}
	}

	options := ListTagsForResourcePaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTagsForResourcePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTagsForResourcePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTagsForResource page.
func (p *ListTagsForResourcePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) {
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

	result, err := p.client.ListTagsForResource(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTagsForResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks-cm",
		OperationName: "ListTagsForResource",
	}
}
