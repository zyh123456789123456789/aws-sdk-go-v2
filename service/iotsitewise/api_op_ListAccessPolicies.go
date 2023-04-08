// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a paginated list of access policies for an identity (an IAM Identity
// Center user, an IAM Identity Center group, or an IAM user) or an IoT SiteWise
// Monitor resource (a portal or project).
func (c *Client) ListAccessPolicies(ctx context.Context, params *ListAccessPoliciesInput, optFns ...func(*Options)) (*ListAccessPoliciesOutput, error) {
	if params == nil {
		params = &ListAccessPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAccessPolicies", params, optFns, c.addOperationListAccessPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAccessPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAccessPoliciesInput struct {

	// The ARN of the IAM user. For more information, see IAM ARNs (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html)
	// in the IAM User Guide. This parameter is required if you specify IAM for
	// identityType .
	IamArn *string

	// The ID of the identity. This parameter is required if you specify USER or GROUP
	// for identityType .
	IdentityId *string

	// The type of identity (IAM Identity Center user, IAM Identity Center group, or
	// IAM user). This parameter is required if you specify identityId .
	IdentityType types.IdentityType

	// The maximum number of results to return for each paginated request. Default: 50
	MaxResults *int32

	// The token to be used for the next set of paginated results.
	NextToken *string

	// The ID of the resource. This parameter is required if you specify resourceType .
	ResourceId *string

	// The type of resource (portal or project). This parameter is required if you
	// specify resourceId .
	ResourceType types.ResourceType

	noSmithyDocumentSerde
}

type ListAccessPoliciesOutput struct {

	// A list that summarizes each access policy.
	//
	// This member is required.
	AccessPolicySummaries []types.AccessPolicySummary

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAccessPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAccessPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAccessPolicies{}, middleware.After)
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
	if err = addEndpointPrefix_opListAccessPoliciesMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAccessPolicies(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListAccessPoliciesMiddleware struct {
}

func (*endpointPrefix_opListAccessPoliciesMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListAccessPoliciesMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "monitor." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opListAccessPoliciesMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opListAccessPoliciesMiddleware{}, `OperationSerializer`, middleware.After)
}

// ListAccessPoliciesAPIClient is a client that implements the ListAccessPolicies
// operation.
type ListAccessPoliciesAPIClient interface {
	ListAccessPolicies(context.Context, *ListAccessPoliciesInput, ...func(*Options)) (*ListAccessPoliciesOutput, error)
}

var _ ListAccessPoliciesAPIClient = (*Client)(nil)

// ListAccessPoliciesPaginatorOptions is the paginator options for
// ListAccessPolicies
type ListAccessPoliciesPaginatorOptions struct {
	// The maximum number of results to return for each paginated request. Default: 50
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAccessPoliciesPaginator is a paginator for ListAccessPolicies
type ListAccessPoliciesPaginator struct {
	options   ListAccessPoliciesPaginatorOptions
	client    ListAccessPoliciesAPIClient
	params    *ListAccessPoliciesInput
	nextToken *string
	firstPage bool
}

// NewListAccessPoliciesPaginator returns a new ListAccessPoliciesPaginator
func NewListAccessPoliciesPaginator(client ListAccessPoliciesAPIClient, params *ListAccessPoliciesInput, optFns ...func(*ListAccessPoliciesPaginatorOptions)) *ListAccessPoliciesPaginator {
	if params == nil {
		params = &ListAccessPoliciesInput{}
	}

	options := ListAccessPoliciesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAccessPoliciesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAccessPoliciesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAccessPolicies page.
func (p *ListAccessPoliciesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAccessPoliciesOutput, error) {
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

	result, err := p.client.ListAccessPolicies(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAccessPolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "ListAccessPolicies",
	}
}
