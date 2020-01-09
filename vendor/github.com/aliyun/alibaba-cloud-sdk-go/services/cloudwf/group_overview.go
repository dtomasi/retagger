package cloudwf

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

// GroupOverview invokes the cloudwf.GroupOverview API synchronously
// api document: https://help.aliyun.com/api/cloudwf/groupoverview.html
func (client *Client) GroupOverview(request *GroupOverviewRequest) (response *GroupOverviewResponse, err error) {
	response = CreateGroupOverviewResponse()
	err = client.DoAction(request, response)
	return
}

// GroupOverviewWithChan invokes the cloudwf.GroupOverview API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/groupoverview.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GroupOverviewWithChan(request *GroupOverviewRequest) (<-chan *GroupOverviewResponse, <-chan error) {
	responseChan := make(chan *GroupOverviewResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GroupOverview(request)
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

// GroupOverviewWithCallback invokes the cloudwf.GroupOverview API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/groupoverview.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GroupOverviewWithCallback(request *GroupOverviewRequest, callback func(response *GroupOverviewResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GroupOverviewResponse
		var err error
		defer close(result)
		response, err = client.GroupOverview(request)
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

// GroupOverviewRequest is the request struct for api GroupOverview
type GroupOverviewRequest struct {
	*requests.RpcRequest
	Gsid requests.Integer `position:"Query" name:"Gsid"`
}

// GroupOverviewResponse is the response struct for api GroupOverview
type GroupOverviewResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateGroupOverviewRequest creates a request to invoke GroupOverview API
func CreateGroupOverviewRequest() (request *GroupOverviewRequest) {
	request = &GroupOverviewRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "GroupOverview", "cloudwf", "openAPI")
	return
}

// CreateGroupOverviewResponse creates a response to parse from GroupOverview response
func CreateGroupOverviewResponse() (response *GroupOverviewResponse) {
	response = &GroupOverviewResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
