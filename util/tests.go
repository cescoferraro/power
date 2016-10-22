package util

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/fatih/color"

	"net/http"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"fmt"
)

func URL() string {
	return "sdflsdfnd"
}

var Server *httptest.Server
var ServerFlag string

type TableTest struct {
	Method       string
	Path         string
	Jwt  		string
	Body         io.Reader
	BodyContains string
	Status       int
	Name         string
	Description  string
}



func SetTestServer(server *httptest.Server) {
	Server = server
	return
}

func SpinSingleTableTests( t *testing.T, test TableTest) (string) {
	NEWLogIfVerbose(color.FgHiBlue, "TEST", "Name: " + test.Name)
	NEWLogIfVerbose(color.FgHiBlue, "TEST", "Description: " + test.Description)


	url := Server.URL + test.Path
	r, err := http.NewRequest(test.Method, url, test.Body)
	assert.NoError(t, err)
	if err != nil {
		NEWLogIfVerbose(color.FgHiBlue, "TEST", err.Error())
	}

	r.Header.Set("Authorization", fmt.Sprintf("Bearer %v", test.Jwt))


	response, err := http.DefaultClient.Do(r)
	if err != nil {
		NEWLogIfVerbose(color.FgHiBlue, "TEST", err.Error())
	}
	assert.NoError(t, err)

	actualBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		NEWLogIfVerbose(color.FgHiBlue, "TEST", err.Error())
	}
	assert.NoError(t, err)

	assert.Contains(t, string(actualBody), test.BodyContains, "body")
	assert.Equal(t, test.Status, response.StatusCode, "status code")


	return string(actualBody)
}
