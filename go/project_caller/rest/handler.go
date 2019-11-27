package rest

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/http/httptrace"
	"os"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

type transport struct {
	current *http.Request
}

// RoundTrip wraps http.DefaultTransport.RoundTrip to keep track
// of the current request.
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.current = req
	return http.DefaultTransport.RoundTrip(req)
}

// GotConn prints whether the connection has been used previously
// for the current request.
func (t *transport) GotConn(info httptrace.GotConnInfo) {
	fmt.Printf("Connection reused for %v? %v\n", t.current.URL, info.Reused)
}

func (handler *Handler) HelloWorld(responseWriter http.ResponseWriter, request *http.Request) {
	//httpClient := http.DefaultClient
	//t := &transport{}
	//trace := &httptrace.ClientTrace{
	//	DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
	//		fmt.Printf("DNS Info: %+v\n Time [%v]\n", dnsInfo, time.Now())
	//	},
	//	GotConn: func(connInfo httptrace.GotConnInfo) {
	//		fmt.Printf("Got Conn: %+v\n Time [%v]\n", connInfo, time.Now())
	//	},
	//}
	//req, err := http.NewRequest(http.MethodGet, os.Getenv("PATH_HELLO_WORLD"), nil)
	//req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	client := resty.New()

	response, err := client.R().
		EnableTrace().
		Get(os.Getenv("PATH_HELLO_WORLD"))
	//
	//client := &http.Client{Transport: t}
	//if err != nil {
	//	WriteErrorResponse(responseWriter, http.StatusInternalServerError, "InternalError", "Cannot create request to call hello world")
	//	return
	//}

	//response, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error [%v]", err)
		WriteErrorResponse(responseWriter, http.StatusInternalServerError, "InternalError", "Failed to call hello world resource")
		return
	}
	//defer response.Body.func() {}():Close()

	if response.StatusCode() != http.StatusOK {
		WriteErrorResponse(responseWriter, http.StatusInternalServerError, "InternalError", fmt.Sprint("Bad status return: [%i] ", response.StatusCode))
		return
	}

	//var result string

	//body, err := ioutil.ReadAll(response.Body())
	//if err != nil {
	//	WriteErrorResponse(responseWriter, http.StatusInternalServerError, "InternalError", "Cannot decode body response from hello world resource")
	//	return
	//}
	//if err := json.NewDecoder(response.Body()).Decode(&result); err != nil {
	//	WriteErrorResponse(responseWriter, http.StatusInternalServerError, "InternalError", "Cannot decode body response from hello world resource")
	//	return
	//}

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("Error      :", err)
	fmt.Println("Status Code:", response.StatusCode())
	fmt.Println("Status     :", response.Status())
	fmt.Println("Time       :", response.Time())
	fmt.Println("Received At:", response.ReceivedAt())
	fmt.Println("Body       :\n", response)
	fmt.Println()

	// Explore trace info
	fmt.Println("Request Trace Info:")
	ti := response.Request.TraceInfo()
	fmt.Println("DNSLookup    :", ti.DNSLookup)
	fmt.Println("ConnTime     :", ti.ConnTime)
	fmt.Println("TLSHandshake :", ti.TLSHandshake)
	fmt.Println("ServerTime   :", ti.ServerTime)
	fmt.Println("ResponseTime :", ti.ResponseTime)
	fmt.Println("TotalTime    :", ti.TotalTime)
	fmt.Println("IsConnReused :", ti.IsConnReused)
	fmt.Println("IsConnWasIdle:", ti.IsConnWasIdle)
	fmt.Println("ConnIdleTime :", ti.ConnIdleTime)

	WriteOkResponse(responseWriter, "Hey I got: "+string(response.Body()))
}

type ErrorResponseBody struct {
	Code    string
	Message string
}

func WriteErrorResponse(writer http.ResponseWriter, status int, code string, message string) {
	writer.Header().Set("Content-Type", "application/json;charset=utf-8")
	writer.WriteHeader(status)
	errorResponseBody := ErrorResponseBody{
		Code:    code,
		Message: message,
	}
	if encodingError := json.NewEncoder(writer).Encode(errorResponseBody); encodingError != nil {
		fmt.Printf("Aie")
	}
}

func WriteOkResponse(writer http.ResponseWriter, body interface{}) {
	writeSuccessWithContent(writer, body, http.StatusOK)
}

func writeSuccessWithContent(writer http.ResponseWriter, body interface{}, status int) {
	writer.Header().Add("Content-Type", "application/json;charset=utf-8")
	writer.WriteHeader(status)
	err := json.NewEncoder(writer).Encode(body)
	if err != nil {
		WriteErrorResponse(writer, http.StatusInternalServerError, "InternalError", "fail to encode body object")
	}
}
