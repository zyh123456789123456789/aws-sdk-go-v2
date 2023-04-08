// Code generated by smithy-go-codegen DO NOT EDIT.

// Package organizations provides the API client, operations, and parameter types
// for AWS Organizations.
//
// Organizations is a web service that enables you to consolidate your multiple
// Amazon Web Services accounts into an organization and centrally manage your
// accounts and their resources. This guide provides descriptions of the
// Organizations operations. For more information about using this service, see the
// Organizations User Guide (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_introduction.html)
// . Support and feedback for Organizations We welcome your feedback. Send your
// comments to feedback-awsorganizations@amazon.com (mailto:feedback-awsorganizations@amazon.com)
// or post your feedback and questions in the Organizations support forum (http://forums.aws.amazon.com/forum.jspa?forumID=219)
// . For more information about the Amazon Web Services support forums, see Forums
// Help (http://forums.aws.amazon.com/help.jspa) . Endpoint to call When using the
// CLI or the Amazon Web Services SDK For the current release of Organizations,
// specify the us-east-1 region for all Amazon Web Services API and CLI calls made
// from the commercial Amazon Web Services Regions outside of China. If calling
// from one of the Amazon Web Services Regions in China, then specify
// cn-northwest-1 . You can do this in the CLI by using these parameters and
// commands:
//   - Use the following parameter with each command to specify both the endpoint
//     and its region: --endpoint-url https://organizations.us-east-1.amazonaws.com
//     (from commercial Amazon Web Services Regions outside of China) or
//     --endpoint-url https://organizations.cn-northwest-1.amazonaws.com.cn (from
//     Amazon Web Services Regions in China)
//   - Use the default endpoint, but configure your default region with this
//     command: aws configure set default.region us-east-1 (from commercial Amazon
//     Web Services Regions outside of China) or aws configure set default.region
//     cn-northwest-1 (from Amazon Web Services Regions in China)
//   - Use the following parameter with each command to specify the endpoint:
//     --region us-east-1 (from commercial Amazon Web Services Regions outside of
//     China) or --region cn-northwest-1 (from Amazon Web Services Regions in China)
//
// Recording API Requests Organizations supports CloudTrail, a service that
// records Amazon Web Services API calls for your Amazon Web Services account and
// delivers log files to an Amazon S3 bucket. By using information collected by
// CloudTrail, you can determine which requests the Organizations service received,
// who made the request and when, and so on. For more about Organizations and its
// support for CloudTrail, see Logging Organizations Events with CloudTrail (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_incident-response.html#orgs_cloudtrail-integration)
// in the Organizations User Guide. To learn more about CloudTrail, including how
// to turn it on and find your log files, see the CloudTrail User Guide (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/what_is_cloud_trail_top_level.html)
// .
package organizations
