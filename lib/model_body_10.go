/*
 * SendinBlue API
 *
 * SendinBlue provide a RESTFul API that can be used with any languages. With this API, you will be able to :   - Manage your campaigns and get the statistics   - Manage your contacts   - Send transactional Emails and SMS   - and much more...  You can download our wrappers at https://github.com/orgs/sendinblue  **Possible responses**   | Code | Message |   | :-------------: | ------------- |   | 200  | OK. Successful Request  |   | 201  | OK. Successful Creation |   | 202  | OK. Request accepted |   | 204  | OK. Successful Update/Deletion  |   | 400  | Error. Bad Request  |   | 401  | Error. Authentication Needed  |   | 402  | Error. Not enough credit, plan upgrade needed  |   | 403  | Error. Permission denied  |   | 404  | Error. Object does not exist |   | 405  | Error. Method not allowed  |   | 406  | Error. Not Acceptable  |
 *
 * API version: 3.0.0
 * Contact: contact@sendinblue.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package lib

type Body10 struct {
	// visitor’s ID received <a href=\"https://developers.sendinblue.com/docs/conversations-webhooks\">from a webhook</a> or generated by you to <a href=\"https://developers.sendinblue.com/docs/customize-the-widget#identifying-existing-users\">bind existing user account to Conversations</a>
	VisitorId string `json:"visitorId"`
	// message text
	Text string `json:"text"`
	// agent ID. It can be found on agent’s page or received <a href=\"https://developers.sendinblue.com/docs/conversations-webhooks\">from a webhook</a>. Optional if `groupId` is set.
	AgentId string `json:"agentId,omitempty"`
	// group ID. It can be found on group’s page. Optional if `agentId` is set.
	GroupId string `json:"groupId,omitempty"`
}
