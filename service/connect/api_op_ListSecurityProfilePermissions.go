// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API is in preview release for Amazon Connect and is subject to change.
// Lists the permissions granted to a security profile.
func (c *Client) ListSecurityProfilePermissions(ctx context.Context, params *ListSecurityProfilePermissionsInput, optFns ...func(*Options)) (*ListSecurityProfilePermissionsOutput, error) {
	if params == nil {
		params = &ListSecurityProfilePermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSecurityProfilePermissions", params, optFns, c.addOperationListSecurityProfilePermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSecurityProfilePermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSecurityProfilePermissionsInput struct {

	// The identifier of the Amazon Connect instance. You can find the instance ID
	// (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// The identifier for the security profle.
	//
	// This member is required.
	SecurityProfileId *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSecurityProfilePermissionsOutput struct {

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// The permissions granted to the security profile. For a complete list of valid
	// permissions, see List of security profile permissions
	// (https://docs.aws.amazon.com/connect/latest/adminguide/security-profile-list.html).
	Permissions []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSecurityProfilePermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSecurityProfilePermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSecurityProfilePermissions{}, middleware.After)
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
	if err = addOpListSecurityProfilePermissionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSecurityProfilePermissions(options.Region), middleware.Before); err != nil {
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

// ListSecurityProfilePermissionsAPIClient is a client that implements the
// ListSecurityProfilePermissions operation.
type ListSecurityProfilePermissionsAPIClient interface {
	ListSecurityProfilePermissions(context.Context, *ListSecurityProfilePermissionsInput, ...func(*Options)) (*ListSecurityProfilePermissionsOutput, error)
}

var _ ListSecurityProfilePermissionsAPIClient = (*Client)(nil)

// ListSecurityProfilePermissionsPaginatorOptions is the paginator options for
// ListSecurityProfilePermissions
type ListSecurityProfilePermissionsPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSecurityProfilePermissionsPaginator is a paginator for
// ListSecurityProfilePermissions
type ListSecurityProfilePermissionsPaginator struct {
	options   ListSecurityProfilePermissionsPaginatorOptions
	client    ListSecurityProfilePermissionsAPIClient
	params    *ListSecurityProfilePermissionsInput
	nextToken *string
	firstPage bool
}

// NewListSecurityProfilePermissionsPaginator returns a new
// ListSecurityProfilePermissionsPaginator
func NewListSecurityProfilePermissionsPaginator(client ListSecurityProfilePermissionsAPIClient, params *ListSecurityProfilePermissionsInput, optFns ...func(*ListSecurityProfilePermissionsPaginatorOptions)) *ListSecurityProfilePermissionsPaginator {
	if params == nil {
		params = &ListSecurityProfilePermissionsInput{}
	}

	options := ListSecurityProfilePermissionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSecurityProfilePermissionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSecurityProfilePermissionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSecurityProfilePermissions page.
func (p *ListSecurityProfilePermissionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSecurityProfilePermissionsOutput, error) {
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

	result, err := p.client.ListSecurityProfilePermissions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSecurityProfilePermissions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListSecurityProfilePermissions",
	}
}
