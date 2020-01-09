package ehpc

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

// ListQueues invokes the ehpc.ListQueues API synchronously
// api document: https://help.aliyun.com/api/ehpc/listqueues.html
func (client *Client) ListQueues(request *ListQueuesRequest) (response *ListQueuesResponse, err error) {
	response = CreateListQueuesResponse()
	err = client.DoAction(request, response)
	return
}

// ListQueuesWithChan invokes the ehpc.ListQueues API asynchronously
// api document: https://help.aliyun.com/api/ehpc/listqueues.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListQueuesWithChan(request *ListQueuesRequest) (<-chan *ListQueuesResponse, <-chan error) {
	responseChan := make(chan *ListQueuesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListQueues(request)
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

// ListQueuesWithCallback invokes the ehpc.ListQueues API asynchronously
// api document: https://help.aliyun.com/api/ehpc/listqueues.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListQueuesWithCallback(request *ListQueuesRequest, callback func(response *ListQueuesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListQueuesResponse
		var err error
		defer close(result)
		response, err = client.ListQueues(request)
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

// ListQueuesRequest is the request struct for api ListQueues
type ListQueuesRequest struct {
	*requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
}

// ListQueuesResponse is the response struct for api ListQueues
type ListQueuesResponse struct {
	*responses.BaseResponse
	RequestId string             `json:"RequestId" xml:"RequestId"`
	Queues    QueuesInListQueues `json:"Queues" xml:"Queues"`
}

// CreateListQueuesRequest creates a request to invoke ListQueues API
func CreateListQueuesRequest() (request *ListQueuesRequest) {
	request = &ListQueuesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "ListQueues", "ehs", "openAPI")
	return
}

// CreateListQueuesResponse creates a response to parse from ListQueues response
func CreateListQueuesResponse() (response *ListQueuesResponse) {
	response = &ListQueuesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
