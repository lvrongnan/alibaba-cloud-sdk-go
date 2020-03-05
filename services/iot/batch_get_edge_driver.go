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

// BatchGetEdgeDriver invokes the iot.BatchGetEdgeDriver API synchronously
// api document: https://help.aliyun.com/api/iot/batchgetedgedriver.html
func (client *Client) BatchGetEdgeDriver(request *BatchGetEdgeDriverRequest) (response *BatchGetEdgeDriverResponse, err error) {
	response = CreateBatchGetEdgeDriverResponse()
	err = client.DoAction(request, response)
	return
}

// BatchGetEdgeDriverWithChan invokes the iot.BatchGetEdgeDriver API asynchronously
// api document: https://help.aliyun.com/api/iot/batchgetedgedriver.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchGetEdgeDriverWithChan(request *BatchGetEdgeDriverRequest) (<-chan *BatchGetEdgeDriverResponse, <-chan error) {
	responseChan := make(chan *BatchGetEdgeDriverResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchGetEdgeDriver(request)
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

// BatchGetEdgeDriverWithCallback invokes the iot.BatchGetEdgeDriver API asynchronously
// api document: https://help.aliyun.com/api/iot/batchgetedgedriver.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchGetEdgeDriverWithCallback(request *BatchGetEdgeDriverRequest, callback func(response *BatchGetEdgeDriverResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchGetEdgeDriverResponse
		var err error
		defer close(result)
		response, err = client.BatchGetEdgeDriver(request)
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

// BatchGetEdgeDriverRequest is the request struct for api BatchGetEdgeDriver
type BatchGetEdgeDriverRequest struct {
	*requests.RpcRequest
	DriverIds     *[]string `position:"Query" name:"DriverIds"  type:"Repeated"`
	IotInstanceId string    `position:"Query" name:"IotInstanceId"`
	ApiProduct    string    `position:"Body" name:"ApiProduct"`
	ApiRevision   string    `position:"Body" name:"ApiRevision"`
}

// BatchGetEdgeDriverResponse is the response struct for api BatchGetEdgeDriver
type BatchGetEdgeDriverResponse struct {
	*responses.BaseResponse
	RequestId    string   `json:"RequestId" xml:"RequestId"`
	Success      bool     `json:"Success" xml:"Success"`
	Code         string   `json:"Code" xml:"Code"`
	ErrorMessage string   `json:"ErrorMessage" xml:"ErrorMessage"`
	DriverList   []Driver `json:"DriverList" xml:"DriverList"`
}

// CreateBatchGetEdgeDriverRequest creates a request to invoke BatchGetEdgeDriver API
func CreateBatchGetEdgeDriverRequest() (request *BatchGetEdgeDriverRequest) {
	request = &BatchGetEdgeDriverRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "BatchGetEdgeDriver", "iot", "openAPI")
	return
}

// CreateBatchGetEdgeDriverResponse creates a response to parse from BatchGetEdgeDriver response
func CreateBatchGetEdgeDriverResponse() (response *BatchGetEdgeDriverResponse) {
	response = &BatchGetEdgeDriverResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
