package dm

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

// ModifySenderAddressNotification invokes the dm.ModifySenderAddressNotification API synchronously
// api document: https://help.aliyun.com/api/dm/modifysenderaddressnotification.html
func (client *Client) ModifySenderAddressNotification(request *ModifySenderAddressNotificationRequest) (response *ModifySenderAddressNotificationResponse, err error) {
	response = CreateModifySenderAddressNotificationResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySenderAddressNotificationWithChan invokes the dm.ModifySenderAddressNotification API asynchronously
// api document: https://help.aliyun.com/api/dm/modifysenderaddressnotification.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySenderAddressNotificationWithChan(request *ModifySenderAddressNotificationRequest) (<-chan *ModifySenderAddressNotificationResponse, <-chan error) {
	responseChan := make(chan *ModifySenderAddressNotificationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySenderAddressNotification(request)
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

// ModifySenderAddressNotificationWithCallback invokes the dm.ModifySenderAddressNotification API asynchronously
// api document: https://help.aliyun.com/api/dm/modifysenderaddressnotification.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySenderAddressNotificationWithCallback(request *ModifySenderAddressNotificationRequest, callback func(response *ModifySenderAddressNotificationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySenderAddressNotificationResponse
		var err error
		defer close(result)
		response, err = client.ModifySenderAddressNotification(request)
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

// ModifySenderAddressNotificationRequest is the request struct for api ModifySenderAddressNotification
type ModifySenderAddressNotificationRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SenderAddressId      string           `position:"Query" name:"SenderAddressId"`
	SenderAddress        string           `position:"Query" name:"SenderAddress"`
	Region               string           `position:"Query" name:"Region"`
	Status               string           `position:"Query" name:"Status"`
}

// ModifySenderAddressNotificationResponse is the response struct for api ModifySenderAddressNotification
type ModifySenderAddressNotificationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifySenderAddressNotificationRequest creates a request to invoke ModifySenderAddressNotification API
func CreateModifySenderAddressNotificationRequest() (request *ModifySenderAddressNotificationRequest) {
	request = &ModifySenderAddressNotificationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "ModifySenderAddressNotification", "", "")
	return
}

// CreateModifySenderAddressNotificationResponse creates a response to parse from ModifySenderAddressNotification response
func CreateModifySenderAddressNotificationResponse() (response *ModifySenderAddressNotificationResponse) {
	response = &ModifySenderAddressNotificationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
