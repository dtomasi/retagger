package actiontrail

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

// GetTrailStatus invokes the actiontrail.GetTrailStatus API synchronously
// api document: https://help.aliyun.com/api/actiontrail/gettrailstatus.html
func (client *Client) GetTrailStatus(request *GetTrailStatusRequest) (response *GetTrailStatusResponse, err error) {
	response = CreateGetTrailStatusResponse()
	err = client.DoAction(request, response)
	return
}

// GetTrailStatusWithChan invokes the actiontrail.GetTrailStatus API asynchronously
// api document: https://help.aliyun.com/api/actiontrail/gettrailstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetTrailStatusWithChan(request *GetTrailStatusRequest) (<-chan *GetTrailStatusResponse, <-chan error) {
	responseChan := make(chan *GetTrailStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTrailStatus(request)
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

// GetTrailStatusWithCallback invokes the actiontrail.GetTrailStatus API asynchronously
// api document: https://help.aliyun.com/api/actiontrail/gettrailstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetTrailStatusWithCallback(request *GetTrailStatusRequest, callback func(response *GetTrailStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTrailStatusResponse
		var err error
		defer close(result)
		response, err = client.GetTrailStatus(request)
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

// GetTrailStatusRequest is the request struct for api GetTrailStatus
type GetTrailStatusRequest struct {
	*requests.RpcRequest
	Name string `position:"Query" name:"Name"`
}

// GetTrailStatusResponse is the response struct for api GetTrailStatus
type GetTrailStatusResponse struct {
	*responses.BaseResponse
	RequestId           string `json:"RequestId" xml:"RequestId"`
	IsLogging           bool   `json:"IsLogging" xml:"IsLogging"`
	LatestDeliveryError string `json:"LatestDeliveryError" xml:"LatestDeliveryError"`
	LatestDeliveryTime  string `json:"LatestDeliveryTime" xml:"LatestDeliveryTime"`
	StartLoggingTime    string `json:"StartLoggingTime" xml:"StartLoggingTime"`
	StopLoggingTime     string `json:"StopLoggingTime" xml:"StopLoggingTime"`
}

// CreateGetTrailStatusRequest creates a request to invoke GetTrailStatus API
func CreateGetTrailStatusRequest() (request *GetTrailStatusRequest) {
	request = &GetTrailStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Actiontrail", "2017-12-04", "GetTrailStatus", "actiontrail", "openAPI")
	return
}

// CreateGetTrailStatusResponse creates a response to parse from GetTrailStatus response
func CreateGetTrailStatusResponse() (response *GetTrailStatusResponse) {
	response = &GetTrailStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
