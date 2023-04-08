// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html)
// . With the latest version, AWS WAF has a single set of endpoints for regional
// and global use. Permanently deletes a RuleGroup . You can't delete a RuleGroup
// if it's still used in any WebACL objects or if it still includes any rules. If
// you just want to remove a RuleGroup from a WebACL , use UpdateWebACL . To
// permanently delete a RuleGroup from AWS WAF, perform the following steps:
//   - Update the RuleGroup to remove rules, if any. For more information, see
//     UpdateRuleGroup .
//   - Use GetChangeToken to get the change token that you provide in the
//     ChangeToken parameter of a DeleteRuleGroup request.
//   - Submit a DeleteRuleGroup request.
func (c *Client) DeleteRuleGroup(ctx context.Context, params *DeleteRuleGroupInput, optFns ...func(*Options)) (*DeleteRuleGroupOutput, error) {
	if params == nil {
		params = &DeleteRuleGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteRuleGroup", params, optFns, c.addOperationDeleteRuleGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteRuleGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRuleGroupInput struct {

	// The value returned by the most recent call to GetChangeToken .
	//
	// This member is required.
	ChangeToken *string

	// The RuleGroupId of the RuleGroup that you want to delete. RuleGroupId is
	// returned by CreateRuleGroup and by ListRuleGroups .
	//
	// This member is required.
	RuleGroupId *string

	noSmithyDocumentSerde
}

type DeleteRuleGroupOutput struct {

	// The ChangeToken that you used to submit the DeleteRuleGroup request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus .
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteRuleGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteRuleGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteRuleGroup{}, middleware.After)
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
	if err = addOpDeleteRuleGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRuleGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteRuleGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "DeleteRuleGroup",
	}
}
