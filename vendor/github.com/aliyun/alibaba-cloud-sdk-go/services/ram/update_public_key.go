package ram

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

// UpdatePublicKey invokes the ram.UpdatePublicKey API synchronously
// api document: https://help.aliyun.com/api/ram/updatepublickey.html
func (client *Client) UpdatePublicKey(request *UpdatePublicKeyRequest) (response *UpdatePublicKeyResponse, err error) {
	response = CreateUpdatePublicKeyResponse()
	err = client.DoAction(request, response)
	return
}

// UpdatePublicKeyWithChan invokes the ram.UpdatePublicKey API asynchronously
// api document: https://help.aliyun.com/api/ram/updatepublickey.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdatePublicKeyWithChan(request *UpdatePublicKeyRequest) (<-chan *UpdatePublicKeyResponse, <-chan error) {
	responseChan := make(chan *UpdatePublicKeyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdatePublicKey(request)
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

// UpdatePublicKeyWithCallback invokes the ram.UpdatePublicKey API asynchronously
// api document: https://help.aliyun.com/api/ram/updatepublickey.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdatePublicKeyWithCallback(request *UpdatePublicKeyRequest, callback func(response *UpdatePublicKeyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdatePublicKeyResponse
		var err error
		defer close(result)
		response, err = client.UpdatePublicKey(request)
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

// UpdatePublicKeyRequest is the request struct for api UpdatePublicKey
type UpdatePublicKeyRequest struct {
	*requests.RpcRequest
	UserPublicKeyId string `position:"Query" name:"UserPublicKeyId"`
	UserName        string `position:"Query" name:"UserName"`
	Status          string `position:"Query" name:"Status"`
}

// UpdatePublicKeyResponse is the response struct for api UpdatePublicKey
type UpdatePublicKeyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdatePublicKeyRequest creates a request to invoke UpdatePublicKey API
func CreateUpdatePublicKeyRequest() (request *UpdatePublicKeyRequest) {
	request = &UpdatePublicKeyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "UpdatePublicKey", "", "")
	return
}

// CreateUpdatePublicKeyResponse creates a response to parse from UpdatePublicKey response
func CreateUpdatePublicKeyResponse() (response *UpdatePublicKeyResponse) {
	response = &UpdatePublicKeyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
