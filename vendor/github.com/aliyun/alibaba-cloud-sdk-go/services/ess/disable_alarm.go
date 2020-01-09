package ess

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

// DisableAlarm invokes the ess.DisableAlarm API synchronously
// api document: https://help.aliyun.com/api/ess/disablealarm.html
func (client *Client) DisableAlarm(request *DisableAlarmRequest) (response *DisableAlarmResponse, err error) {
	response = CreateDisableAlarmResponse()
	err = client.DoAction(request, response)
	return
}

// DisableAlarmWithChan invokes the ess.DisableAlarm API asynchronously
// api document: https://help.aliyun.com/api/ess/disablealarm.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisableAlarmWithChan(request *DisableAlarmRequest) (<-chan *DisableAlarmResponse, <-chan error) {
	responseChan := make(chan *DisableAlarmResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableAlarm(request)
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

// DisableAlarmWithCallback invokes the ess.DisableAlarm API asynchronously
// api document: https://help.aliyun.com/api/ess/disablealarm.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisableAlarmWithCallback(request *DisableAlarmRequest, callback func(response *DisableAlarmResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableAlarmResponse
		var err error
		defer close(result)
		response, err = client.DisableAlarm(request)
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

// DisableAlarmRequest is the request struct for api DisableAlarm
type DisableAlarmRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AlarmTaskId          string           `position:"Query" name:"AlarmTaskId"`
}

// DisableAlarmResponse is the response struct for api DisableAlarm
type DisableAlarmResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDisableAlarmRequest creates a request to invoke DisableAlarm API
func CreateDisableAlarmRequest() (request *DisableAlarmRequest) {
	request = &DisableAlarmRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "DisableAlarm", "ess", "openAPI")
	return
}

// CreateDisableAlarmResponse creates a response to parse from DisableAlarm response
func CreateDisableAlarmResponse() (response *DisableAlarmResponse) {
	response = &DisableAlarmResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
