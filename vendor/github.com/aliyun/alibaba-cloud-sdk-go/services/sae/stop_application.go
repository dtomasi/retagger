package sae

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// StopApplication invokes the sae.StopApplication API synchronously
// api document: https://help.aliyun.com/api/sae/stopapplication.html
func (client *Client) StopApplication(request *StopApplicationRequest) (response *StopApplicationResponse, err error) {
	response = CreateStopApplicationResponse()
	err = client.DoAction(request, response)
	return
}

// StopApplicationWithChan invokes the sae.StopApplication API asynchronously
// api document: https://help.aliyun.com/api/sae/stopapplication.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StopApplicationWithChan(request *StopApplicationRequest) (<-chan *StopApplicationResponse, <-chan error) {
	responseChan := make(chan *StopApplicationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopApplication(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// StopApplicationWithCallback invokes the sae.StopApplication API asynchronously
// api document: https://help.aliyun.com/api/sae/stopapplication.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StopApplicationWithCallback(request *StopApplicationRequest, callback func(response *StopApplicationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopApplicationResponse
		var err error
		defer close(result)
		response, err = client.StopApplication(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// StopApplicationRequest is the request struct for api StopApplication
type StopApplicationRequest struct {
	*requests.RoaRequest
	AppId string `position:"Query" name:"AppId"`
}

// StopApplicationResponse is the response struct for api StopApplication
type StopApplicationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateStopApplicationRequest creates a request to invoke StopApplication API
func CreateStopApplicationRequest() (request *StopApplicationRequest) {
	request = &StopApplicationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "StopApplication", "/pop/v1/sam/app/stopApplication", "serverless", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateStopApplicationResponse creates a response to parse from StopApplication response
func CreateStopApplicationResponse() (response *StopApplicationResponse) {
	response = &StopApplicationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
