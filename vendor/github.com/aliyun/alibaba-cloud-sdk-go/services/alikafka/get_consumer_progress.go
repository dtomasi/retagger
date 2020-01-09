package alikafka

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

// GetConsumerProgress invokes the alikafka.GetConsumerProgress API synchronously
// api document: https://help.aliyun.com/api/alikafka/getconsumerprogress.html
func (client *Client) GetConsumerProgress(request *GetConsumerProgressRequest) (response *GetConsumerProgressResponse, err error) {
	response = CreateGetConsumerProgressResponse()
	err = client.DoAction(request, response)
	return
}

// GetConsumerProgressWithChan invokes the alikafka.GetConsumerProgress API asynchronously
// api document: https://help.aliyun.com/api/alikafka/getconsumerprogress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetConsumerProgressWithChan(request *GetConsumerProgressRequest) (<-chan *GetConsumerProgressResponse, <-chan error) {
	responseChan := make(chan *GetConsumerProgressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetConsumerProgress(request)
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

// GetConsumerProgressWithCallback invokes the alikafka.GetConsumerProgress API asynchronously
// api document: https://help.aliyun.com/api/alikafka/getconsumerprogress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetConsumerProgressWithCallback(request *GetConsumerProgressRequest, callback func(response *GetConsumerProgressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetConsumerProgressResponse
		var err error
		defer close(result)
		response, err = client.GetConsumerProgress(request)
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

// GetConsumerProgressRequest is the request struct for api GetConsumerProgress
type GetConsumerProgressRequest struct {
	*requests.RpcRequest
	ConsumerId string `position:"Query" name:"ConsumerId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// GetConsumerProgressResponse is the response struct for api GetConsumerProgress
type GetConsumerProgressResponse struct {
	*responses.BaseResponse
	Success          bool             `json:"Success" xml:"Success"`
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	Code             int              `json:"Code" xml:"Code"`
	Message          string           `json:"Message" xml:"Message"`
	ConsumerProgress ConsumerProgress `json:"ConsumerProgress" xml:"ConsumerProgress"`
}

// CreateGetConsumerProgressRequest creates a request to invoke GetConsumerProgress API
func CreateGetConsumerProgressRequest() (request *GetConsumerProgressRequest) {
	request = &GetConsumerProgressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2019-09-16", "GetConsumerProgress", "alikafka", "openAPI")
	return
}

// CreateGetConsumerProgressResponse creates a response to parse from GetConsumerProgress response
func CreateGetConsumerProgressResponse() (response *GetConsumerProgressResponse) {
	response = &GetConsumerProgressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
