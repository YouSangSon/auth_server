package response

const (
	MESSAGE string = "message"
	RESULT  string = "result"
	DATA    string = "data"

	OK       string = "ok"
	ACCEPTED string = "accepted"
	ERROR    string = "error"
	INFO     string = "info"

	// ERROR
	INVALID_DATA                 = "Invalid data"
	SERVER_ERROR                 = "Server error."
	INVALID_TOKEN                = "Invalid token."
	BAD_REQUEST                  = "Invalid request."
	USER_NOT_FOUND               = "No users found."
	PASSWD_NOT_MATCH             = "User ID or user password is invalid."
	DEVICE_NOT_FOUND             = "No devices found."
	HOLIDAY_NOT_FOUND            = "No holiday found"
	CAN_NOT_PUBSUB_CREATE_CLINET = "can't create pubsub client"
	CAN_NOT_PUBSUB_PUBLISH       = "can't publish pubsub"
	ERROR_MESSAGE                = "errorMessage"
	NOERROR                      = "NOERROR"
	MISSING                      = "missing"
	NO_AUTHENTIFICATION          = "There is no authentification"
	DB_ERROR                     = "There is database error"
	SITE_UPDATE_SUCCESS          = "The site information has been modified"
)

// type response struct {
// 	Status  interface{} `json:"status,omitempty"`
// 	Data    interface{} `json:"data,omitempty"`
// 	Message interface{} `json:"message,omitempty"`
// }

// func Response(status interface{}, data interface{}, message interface{}) string {
// 	newResponse := response{
// 		Status:  status,
// 		Data:    data,
// 		Message: message,
// 	}

// 	a, err := json.Marshal(newResponse)
// 	if err != nil {
// 		return ""
// 	}

// 	return string(a)
// }

// type Response struct {
// 	Code    string      `json:"code"`
// 	Message string      `json:"message"`
// 	Data    interface{} `json:"data"`
// }
