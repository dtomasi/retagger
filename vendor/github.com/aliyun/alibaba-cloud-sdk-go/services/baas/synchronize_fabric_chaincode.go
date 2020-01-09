package baas

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

// SynchronizeFabricChaincode invokes the baas.SynchronizeFabricChaincode API synchronously
// api document: https://help.aliyun.com/api/baas/synchronizefabricchaincode.html
func (client *Client) SynchronizeFabricChaincode(request *SynchronizeFabricChaincodeRequest) (response *SynchronizeFabricChaincodeResponse, err error) {
	response = CreateSynchronizeFabricChaincodeResponse()
	err = client.DoAction(request, response)
	return
}

// SynchronizeFabricChaincodeWithChan invokes the baas.SynchronizeFabricChaincode API asynchronously
// api document: https://help.aliyun.com/api/baas/synchronizefabricchaincode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SynchronizeFabricChaincodeWithChan(request *SynchronizeFabricChaincodeRequest) (<-chan *SynchronizeFabricChaincodeResponse, <-chan error) {
	responseChan := make(chan *SynchronizeFabricChaincodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SynchronizeFabricChaincode(request)
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

// SynchronizeFabricChaincodeWithCallback invokes the baas.SynchronizeFabricChaincode API asynchronously
// api document: https://help.aliyun.com/api/baas/synchronizefabricchaincode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SynchronizeFabricChaincodeWithCallback(request *SynchronizeFabricChaincodeRequest, callback func(response *SynchronizeFabricChaincodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SynchronizeFabricChaincodeResponse
		var err error
		defer close(result)
		response, err = client.SynchronizeFabricChaincode(request)
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

// SynchronizeFabricChaincodeRequest is the request struct for api SynchronizeFabricChaincode
type SynchronizeFabricChaincodeRequest struct {
	*requests.RpcRequest
	OrganizationId string `position:"Body" name:"OrganizationId"`
	ChaincodeId    string `position:"Body" name:"ChaincodeId"`
}

// SynchronizeFabricChaincodeResponse is the response struct for api SynchronizeFabricChaincode
type SynchronizeFabricChaincodeResponse struct {
	*responses.BaseResponse
	RequestId string                             `json:"RequestId" xml:"RequestId"`
	Success   bool                               `json:"Success" xml:"Success"`
	ErrorCode int                                `json:"ErrorCode" xml:"ErrorCode"`
	Result    ResultInSynchronizeFabricChaincode `json:"Result" xml:"Result"`
}

// CreateSynchronizeFabricChaincodeRequest creates a request to invoke SynchronizeFabricChaincode API
func CreateSynchronizeFabricChaincodeRequest() (request *SynchronizeFabricChaincodeRequest) {
	request = &SynchronizeFabricChaincodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "SynchronizeFabricChaincode", "baas", "openAPI")
	return
}

// CreateSynchronizeFabricChaincodeResponse creates a response to parse from SynchronizeFabricChaincode response
func CreateSynchronizeFabricChaincodeResponse() (response *SynchronizeFabricChaincodeResponse) {
	response = &SynchronizeFabricChaincodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
