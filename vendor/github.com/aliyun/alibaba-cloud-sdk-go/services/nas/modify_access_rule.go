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

package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyAccessRule invokes the nas.ModifyAccessRule API synchronously
// api document: https://help.aliyun.com/api/nas/modifyaccessrule.html
func (client *Client) ModifyAccessRule(request *ModifyAccessRuleRequest) (response *ModifyAccessRuleResponse, err error) {
	response = CreateModifyAccessRuleResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyAccessRuleWithChan invokes the nas.ModifyAccessRule API asynchronously
// api document: https://help.aliyun.com/api/nas/modifyaccessrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyAccessRuleWithChan(request *ModifyAccessRuleRequest) (<-chan *ModifyAccessRuleResponse, <-chan error) {
	responseChan := make(chan *ModifyAccessRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyAccessRule(request)
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

// ModifyAccessRuleWithCallback invokes the nas.ModifyAccessRule API asynchronously
// api document: https://help.aliyun.com/api/nas/modifyaccessrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyAccessRuleWithCallback(request *ModifyAccessRuleRequest, callback func(response *ModifyAccessRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyAccessRuleResponse
		var err error
		defer close(result)
		response, err = client.ModifyAccessRule(request)
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

// ModifyAccessRuleRequest is the request struct for api ModifyAccessRule
type ModifyAccessRuleRequest struct {
	*requests.RpcRequest
	AccessGroupName string           `position:"Query" name:"AccessGroupName"`
	AccessRuleId    string           `position:"Query" name:"AccessRuleId"`
	SourceCidrIp    string           `position:"Query" name:"SourceCidrIp"`
	RWAccessType    string           `position:"Query" name:"RWAccessType"`
	UserAccessType  string           `position:"Query" name:"UserAccessType"`
	Priority        requests.Integer `position:"Query" name:"Priority"`
	FileSystemType  string           `position:"Query" name:"FileSystemType"`
}

// ModifyAccessRuleResponse is the response struct for api ModifyAccessRule
type ModifyAccessRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyAccessRuleRequest creates a request to invoke ModifyAccessRule API
func CreateModifyAccessRuleRequest() (request *ModifyAccessRuleRequest) {
	request = &ModifyAccessRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "ModifyAccessRule", "nas", "openAPI")
	return
}

// CreateModifyAccessRuleResponse creates a response to parse from ModifyAccessRule response
func CreateModifyAccessRuleResponse() (response *ModifyAccessRuleResponse) {
	response = &ModifyAccessRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
