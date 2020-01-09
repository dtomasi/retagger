package iot

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

// QueryDeviceStatistics invokes the iot.QueryDeviceStatistics API synchronously
// api document: https://help.aliyun.com/api/iot/querydevicestatistics.html
func (client *Client) QueryDeviceStatistics(request *QueryDeviceStatisticsRequest) (response *QueryDeviceStatisticsResponse, err error) {
	response = CreateQueryDeviceStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDeviceStatisticsWithChan invokes the iot.QueryDeviceStatistics API asynchronously
// api document: https://help.aliyun.com/api/iot/querydevicestatistics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDeviceStatisticsWithChan(request *QueryDeviceStatisticsRequest) (<-chan *QueryDeviceStatisticsResponse, <-chan error) {
	responseChan := make(chan *QueryDeviceStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDeviceStatistics(request)
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

// QueryDeviceStatisticsWithCallback invokes the iot.QueryDeviceStatistics API asynchronously
// api document: https://help.aliyun.com/api/iot/querydevicestatistics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDeviceStatisticsWithCallback(request *QueryDeviceStatisticsRequest, callback func(response *QueryDeviceStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDeviceStatisticsResponse
		var err error
		defer close(result)
		response, err = client.QueryDeviceStatistics(request)
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

// QueryDeviceStatisticsRequest is the request struct for api QueryDeviceStatistics
type QueryDeviceStatisticsRequest struct {
	*requests.RpcRequest
	ProductKey    string `position:"Query" name:"ProductKey"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
}

// QueryDeviceStatisticsResponse is the response struct for api QueryDeviceStatistics
type QueryDeviceStatisticsResponse struct {
	*responses.BaseResponse
	RequestId    string                      `json:"RequestId" xml:"RequestId"`
	Success      bool                        `json:"Success" xml:"Success"`
	Code         string                      `json:"Code" xml:"Code"`
	ErrorMessage string                      `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQueryDeviceStatistics `json:"Data" xml:"Data"`
}

// CreateQueryDeviceStatisticsRequest creates a request to invoke QueryDeviceStatistics API
func CreateQueryDeviceStatisticsRequest() (request *QueryDeviceStatisticsRequest) {
	request = &QueryDeviceStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryDeviceStatistics", "iot", "openAPI")
	return
}

// CreateQueryDeviceStatisticsResponse creates a response to parse from QueryDeviceStatistics response
func CreateQueryDeviceStatisticsResponse() (response *QueryDeviceStatisticsResponse) {
	response = &QueryDeviceStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
