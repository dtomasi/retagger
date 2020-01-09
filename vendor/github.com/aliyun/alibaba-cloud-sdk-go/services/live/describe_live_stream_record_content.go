package live

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

// DescribeLiveStreamRecordContent invokes the live.DescribeLiveStreamRecordContent API synchronously
// api document: https://help.aliyun.com/api/live/describelivestreamrecordcontent.html
func (client *Client) DescribeLiveStreamRecordContent(request *DescribeLiveStreamRecordContentRequest) (response *DescribeLiveStreamRecordContentResponse, err error) {
	response = CreateDescribeLiveStreamRecordContentResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveStreamRecordContentWithChan invokes the live.DescribeLiveStreamRecordContent API asynchronously
// api document: https://help.aliyun.com/api/live/describelivestreamrecordcontent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveStreamRecordContentWithChan(request *DescribeLiveStreamRecordContentRequest) (<-chan *DescribeLiveStreamRecordContentResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamRecordContentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamRecordContent(request)
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

// DescribeLiveStreamRecordContentWithCallback invokes the live.DescribeLiveStreamRecordContent API asynchronously
// api document: https://help.aliyun.com/api/live/describelivestreamrecordcontent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveStreamRecordContentWithCallback(request *DescribeLiveStreamRecordContentRequest, callback func(response *DescribeLiveStreamRecordContentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamRecordContentResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamRecordContent(request)
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

// DescribeLiveStreamRecordContentRequest is the request struct for api DescribeLiveStreamRecordContent
type DescribeLiveStreamRecordContentRequest struct {
	*requests.RpcRequest
	StartTime     string           `position:"Query" name:"StartTime"`
	AppName       string           `position:"Query" name:"AppName"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	StreamName    string           `position:"Query" name:"StreamName"`
	DomainName    string           `position:"Query" name:"DomainName"`
	EndTime       string           `position:"Query" name:"EndTime"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveStreamRecordContentResponse is the response struct for api DescribeLiveStreamRecordContent
type DescribeLiveStreamRecordContentResponse struct {
	*responses.BaseResponse
	RequestId             string                `json:"RequestId" xml:"RequestId"`
	RecordContentInfoList RecordContentInfoList `json:"RecordContentInfoList" xml:"RecordContentInfoList"`
}

// CreateDescribeLiveStreamRecordContentRequest creates a request to invoke DescribeLiveStreamRecordContent API
func CreateDescribeLiveStreamRecordContentRequest() (request *DescribeLiveStreamRecordContentRequest) {
	request = &DescribeLiveStreamRecordContentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveStreamRecordContent", "live", "openAPI")
	return
}

// CreateDescribeLiveStreamRecordContentResponse creates a response to parse from DescribeLiveStreamRecordContent response
func CreateDescribeLiveStreamRecordContentResponse() (response *DescribeLiveStreamRecordContentResponse) {
	response = &DescribeLiveStreamRecordContentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
