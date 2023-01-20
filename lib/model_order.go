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

type Order struct {
	// Unique ID of the order.
	Id string `json:"id"`
	// Event occurrence UTC date-time (YYYY-MM-DDTHH:mm:ssZ), when order is actually created.
	CreatedAt string `json:"createdAt"`
	// Event updated UTC date-time (YYYY-MM-DDTHH:mm:ssZ), when the status of the order is actually changed/updated.
	UpdatedAt string `json:"updatedAt"`
	// State of the order.
	Status string `json:"status"`
	// Total amount of the order, including all shipping expenses, tax and the price of items.
	Amount   float32         `json:"amount"`
	Products []OrderProducts `json:"products"`
	// Email of the contact, Mandatory if \"phone\" field is not passed in \"billing\" parameter.
	Email   string        `json:"email,omitempty"`
	Billing *OrderBilling `json:"billing,omitempty"`
	// Coupons applied to the order. Stored case insensitive.
	Coupons []string `json:"coupons,omitempty"`
}
