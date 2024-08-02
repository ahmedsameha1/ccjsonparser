package endtoendtests

import (
	"bytes"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ahmedsameha1/ccjsonparser/cmd/server"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestWebAPI(t *testing.T) {
	rootEndPoint := "http://localhost:8080/"
	t.Run("Error: there is no Content-Type: application/json header", func(t *testing.T) {
		webServer := server.New()
		go webServer.Run()
		WaitWebServerToBeReady()
		webRequest, err := http.NewRequest(http.MethodPost, rootEndPoint,
			nil)
		if err != nil {
			t.Fatal(err)
		}
		httpResponse, err := http.DefaultClient.Do(webRequest)
		if err != nil {
			t.Fatal(err)
		}
		defer httpResponse.Body.Close()
		assert.Equal(t, http.StatusBadRequest, httpResponse.StatusCode)
		responseBody, err := io.ReadAll(httpResponse.Body)
		if err != nil {
			t.Fatal(err)
		}
		var got gin.H
		err = json.Unmarshal(responseBody, &got)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, gin.H{"error": server.ErrNoApplicationJsonHeader.Error()}, got)
	})

	t.Run("Error: there is no JSON in the body to be validated", func(t *testing.T) {
		webServer := server.New()
		go webServer.Run()
		WaitWebServerToBeReady()
		webRequest, err := http.NewRequest(http.MethodPost, rootEndPoint,
			nil)
		if err != nil {
			t.Fatal(err)
		}
		webRequest.Header.Add(server.CONTENT_TYPE_HEADER, server.APPLICATION_JSON)
		httpResponse, err := http.DefaultClient.Do(webRequest)
		if err != nil {
			t.Fatal(err)
		}
		defer httpResponse.Body.Close()
		assert.Equal(t, http.StatusBadRequest, httpResponse.StatusCode)
		responseBody, err := io.ReadAll(httpResponse.Body)
		if err != nil {
			t.Fatal(err)
		}
		var got gin.H
		err = json.Unmarshal(responseBody, &got)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, gin.H{"error": server.ErrWhileReadingBody.Error()}, got)
	})

	t.Run("Error: the body is longer than 1024 character", func(t *testing.T) {
		webServer := server.New()
		go webServer.Run()
		WaitWebServerToBeReady()
		body, err := json.Marshal(strings.Repeat("v", 1025))
		if err != nil {
			t.Fatal(err)
		}
		webRequest, err := http.NewRequest(http.MethodPost, rootEndPoint,
			io.NopCloser(bytes.NewReader(body)))
		if err != nil {
			t.Fatal(err)
		}
		webRequest.Header.Add(server.CONTENT_TYPE_HEADER, server.APPLICATION_JSON)
		httpResponse, err := http.DefaultClient.Do(webRequest)
		if err != nil {
			t.Fatal(err)
		}
		defer httpResponse.Body.Close()
		assert.Equal(t, http.StatusBadRequest, httpResponse.StatusCode)
		responseBody, err := io.ReadAll(httpResponse.Body)
		if err != nil {
			t.Fatal(err)
		}
		var got gin.H
		err = json.Unmarshal(responseBody, &got)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, gin.H{"error": server.ErrBodyIsTooLong.Error()}, got)
	})

	for k, v := range validJSONTests {
		if k < 43 {
			fileContent, err := os.ReadFile("../" + v)
			if err != nil {
				t.Fatal(err)
			}
			t.Run("Good case: a valid JSON", func(t *testing.T) {
				webServer := server.New()
				go webServer.Run()
				WaitWebServerToBeReady()
				body, err := json.Marshal(string(fileContent))
				if err != nil {
					t.Fatal(err)
				}
				webRequest, err := http.NewRequest(http.MethodPost, rootEndPoint,
					io.NopCloser(bytes.NewReader(body)))
				if err != nil {
					t.Fatal(err)
				}
				webRequest.Header.Add(server.CONTENT_TYPE_HEADER, server.APPLICATION_JSON)
				httpResponse, err := http.DefaultClient.Do(webRequest)
				if err != nil {
					t.Fatal(err)
				}
				defer httpResponse.Body.Close()
				assert.Equal(t, http.StatusOK, httpResponse.StatusCode)
				responseBody, err := io.ReadAll(httpResponse.Body)
				if err != nil {
					t.Fatal(err)
				}
				var got gin.H
				err = json.Unmarshal(responseBody, &got)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, gin.H{"result": "This is a valid JSON"}, got)
			})
		}

		for k, v := range invalidJSONTests {
			if k < 115 {
				fileContent, err := os.ReadFile("../" + v.filePath)
				if err != nil {
					t.Fatal(err)
				}
				t.Run("Good case: an invalid JSON", func(t *testing.T) {
					webServer := server.New()
					go webServer.Run()
					WaitWebServerToBeReady()
					body, err := json.Marshal(string(fileContent))
					if err != nil {
						t.Fatal(err)
					}
					webRequest, err := http.NewRequest(http.MethodPost, rootEndPoint,
						io.NopCloser(bytes.NewReader(body)))
					if err != nil {
						t.Fatal(err)
					}
					webRequest.Header.Add(server.CONTENT_TYPE_HEADER, server.APPLICATION_JSON)
					httpResponse, err := http.DefaultClient.Do(webRequest)
					if err != nil {
						t.Fatal(err)
					}
					defer httpResponse.Body.Close()
					assert.Equal(t, http.StatusOK, httpResponse.StatusCode)
					responseBody, err := io.ReadAll(httpResponse.Body)
					if err != nil {
						t.Fatal(err)
					}
					var got gin.H
					err = json.Unmarshal(responseBody, &got)
					if err != nil {
						t.Fatal(err)
					}
					assert.Equal(t, gin.H{"result": v.errResult[:len(v.errResult)-1]}, got)
				})
			}
		}
	}
}

func WaitWebServerToBeReady() {
	for {
		if _, err := net.Dial("tcp", "localhost:8080"); err != nil {
			time.Sleep(time.Millisecond * 10)
		} else {
			break
		}
	}
}
