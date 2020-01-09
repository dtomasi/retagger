package vod

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

// DetachAppPolicyFromIdentity invokes the vod.DetachAppPolicyFromIdentity API synchronously
// api document: https://help.aliyun.com/api/vod/detachapppolicyfromidentity.html
func (client *Client) DetachAppPolicyFromIdentity(request *DetachAppPolicyFromIdentityRequest) (response *DetachAppPolicyFromIdentityResponse, err error) {
	response = CreateDetachAppPolicyFromIdentityResponse()
	err = client.DoAction(request, response)
	return
}

// DetachAppPolicyFromIdentityWithChan invokes the vod.DetachAppPolicyFromIdentity API asynchronously
// api document: https://help.aliyun.com/api/vod/detachapppolicyfromidentity.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DetachAppPolicyFromIdentityWithChan(request *DetachAppPolicyFromIdentityRequest) (<-chan *DetachAppPolicyFromIdentityResponse, <-chan error) {
	responseChan := make(chan *DetachAppPolicyFromIdentityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetachAppPolicyFromIdentity(request)
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

// DetachAppPolicyFromIdentityWithCallback invokes the vod.DetachAppPolicyFromIdentity API asynchronously
// api document: https://help.aliyun.com/api/vod/detachapppolicyfromidentity.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DetachAppPolicyFromIdentityWithCallback(request *DetachAppPolicyFromIdentityRequest, callback func(response *DetachAppPolicyFromIdentityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetachAppPolicyFromIdentityResponse
		var err error
		defer close(result)
		response, err = client.DetachAppPolicyFromIdentity(request)
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

// DetachAppPolicyFromIdentityRequest is the request struct for api DetachAppPolicyFromIdentity
type DetachAppPolicyFromIdentityRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PolicyNames          string           `position:"Query" name:"PolicyNames"`
	IdentityName         string           `position:"Query" name:"IdentityName"`
	IdentityType         string           `position:"Query" name:"IdentityType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AppId                string           `position:"Query" name:"AppId"`
}

// DetachAppPolicyFromIdentityResponse is the response struct for api DetachAppPolicyFromIdentity
type DetachAppPolicyFromIdentityResponse struct {
	*responses.BaseResponse
	RequestId           string   `json:"RequestId" xml:"RequestId"`
	NonExistPolicyNames []string `json:"NonExistPolicyNames" xml:"NonExistPolicyNames"`
	FailedPolicyNames   []string `json:"FailedPolicyNames" xml:"FailedPolicyNames"`
}

// CreateDetachAppPolicyFromIdentityRequest creates a request to invoke DetachAppPolicyFromIdentity API
func CreateDetachAppPolicyFromIdentityRequest() (request *DetachAppPolicyFromIdentityRequest) {
	request = &DetachAppPolicyFromIdentityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DetachAppPolicyFromIdentity", "vod", "openAPI")
	return
}

// CreateDetachAppPolicyFromIdentityResponse creates a response to parse from DetachAppPolicyFromIdentity response
func CreateDetachAppPolicyFromIdentityResponse() (response *DetachAppPolicyFromIdentityResponse) {
	response = &DetachAppPolicyFromIdentityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
