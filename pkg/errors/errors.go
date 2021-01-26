package errors

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	errorsHelpers "github.com/pkg/errors"
)

// APIError is the interface that encapsulates all API HTTP Errors
type APIError interface {
	Code() string
	Message() string
}

// New Generates a New Error
func New(message string) error {
	return errorsHelpers.New(message)
}

// Wrap encapsulates an error for forming stacktrace
func Wrap(err error, message string) error {
	return errorsHelpers.Wrap(err, message)
}

//new wrap error

//Fault .
type Fault struct {
	Type        string `json:"type"`
	Code        int    `json:"code"`
	Details     string `json:"details"`
	err         string
	method      string
	endpoint    string
	clientIP    string
	agent       string
	sequence    int64
	userID      string
	programCode string
	trackingID  string
	errorStack  []*ErrorStack
	prueba      []string
}

//ErrorStack .
type ErrorStack struct {
	LOC      int    `json:"loc"`
	File     string `json:"file"`
	Package  string `json:"package"`
	FuncName string `json:"func_name"`
}

//Create .
func Create(code int, err string) *Fault {
	bug := &Fault{
		Code:    code,
		Details: err,
	}
	return bug
}

//GetType return the log type (error|warning|info|debug)
func (e *Fault) GetType() string {
	if e.Type != "" {
		return "Warning"
	}
	return e.Type
}

//GetCode return the http code status to identify
//errors focused on the http request flow
func (e *Fault) GetCode() int {
	return e.Code
}

//GetDetails return more information about the error
func (e *Fault) GetDetails() string {
	return e.Details
}

//Error return more information about the cause
func (e *Fault) Error() string {
	return e.err
}

//GetErrorStack return the current service where
//was produced the error
func (e *Fault) GetErrorStack() []*ErrorStack {

	return e.errorStack
}

//GetMethod return the current request method
//for the current context where was produced the error
func (e *Fault) GetMethod(c *gin.Context) string {
	e.method = c.Request.Method
	return e.method
}

//GetEndpoint return the current request endpoint
//for the current context where was produced the error
func (e *Fault) GetEndpoint(c *gin.Context) string {
	e.endpoint = c.Request.URL.String()
	return e.endpoint
}

//GetClientIP return the current client ip who begins
//the request
func (e *Fault) GetClientIP(c *gin.Context) string {
	e.clientIP = c.Request.Host
	return e.clientIP
}

//GetAgent return the currenct agent (browser/mobile, etc) who begins
//the request
func (e *Fault) GetAgent(c *gin.Context) string {
	return e.agent
}

//GetSequence return a sequence used to identify in what step happens the error
func (e *Fault) GetSequence(c *gin.Context) int64 {
	ISequence, ok := c.Get("Sequence")
	fmt.Sprintf("Type = %s", reflect.TypeOf(ISequence))
	if ok {
		//e.sequence = ISequence.(int)
	} else {
		if c.GetHeader("Sequence") != "" {
			sequence, _ := strconv.ParseInt(c.GetHeader("Sequence"), 10, 64)
			e.sequence = sequence
		}
	}
	return e.sequence
}

//GetUserID return the current user_id who begins the request
func (e *Fault) GetUserID(c *gin.Context) string {
	return e.userID
}

//GetProgramCode return the current program_code of the current request
func (e *Fault) GetProgramCode(c *gin.Context) string {
	return e.programCode
}

//GetTrackingID return a unique identifier used to attach the whole
//request between microservices
func (e *Fault) GetTrackingID(c *gin.Context) string {
	e.trackingID = c.GetHeader("Service-Traceid")
	return e.trackingID
}

//StackTrace return much more information about where
//was generated the current error, information such as name of
//microservice or package, line of code, file and caller's name function.
//This information is getting from execution time.
func (e *Fault) StackTrace() *Fault {
	fpcs := make([]uintptr, 10) // at least 1 entry needed
	n := runtime.Callers(2, fpcs)
	if n == 0 {
		fmt.Println("MSG: NO CALLER")
	}

	caller := runtime.FuncForPC(fpcs[0])
	if caller == nil {
		fmt.Println("MSG CALLER WAS NIL")
	}
	file, line := caller.FileLine(fpcs[0] - 1)

	splitPathName := strings.Split(caller.Name(), "/")
	splitPackageName := strings.Split(splitPathName[len(splitPathName)-1], ".")
	stack := &ErrorStack{
		LOC:      line,
		File:     file,
		Package:  splitPackageName[0],
		FuncName: splitPackageName[1],
	}
	e.errorStack = append(e.errorStack, stack)
	return e
}

//Notify allows to send an email notification when some errors
//it's really important to fix.
func (e *Fault) Notify() *Fault {
	return e
}
