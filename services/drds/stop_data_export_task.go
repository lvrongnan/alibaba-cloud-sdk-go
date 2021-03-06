package drds

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

// StopDataExportTask invokes the drds.StopDataExportTask API synchronously
func (client *Client) StopDataExportTask(request *StopDataExportTaskRequest) (response *StopDataExportTaskResponse, err error) {
	response = CreateStopDataExportTaskResponse()
	err = client.DoAction(request, response)
	return
}

// StopDataExportTaskWithChan invokes the drds.StopDataExportTask API asynchronously
func (client *Client) StopDataExportTaskWithChan(request *StopDataExportTaskRequest) (<-chan *StopDataExportTaskResponse, <-chan error) {
	responseChan := make(chan *StopDataExportTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopDataExportTask(request)
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

// StopDataExportTaskWithCallback invokes the drds.StopDataExportTask API asynchronously
func (client *Client) StopDataExportTaskWithCallback(request *StopDataExportTaskRequest, callback func(response *StopDataExportTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopDataExportTaskResponse
		var err error
		defer close(result)
		response, err = client.StopDataExportTask(request)
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

// StopDataExportTaskRequest is the request struct for api StopDataExportTask
type StopDataExportTaskRequest struct {
	*requests.RpcRequest
	TaskId requests.Integer `position:"Query" name:"TaskId"`
}

// StopDataExportTaskResponse is the response struct for api StopDataExportTask
type StopDataExportTaskResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	Success          bool             `json:"Success" xml:"Success"`
	TaskManageResult TaskManageResult `json:"TaskManageResult" xml:"TaskManageResult"`
}

// CreateStopDataExportTaskRequest creates a request to invoke StopDataExportTask API
func CreateStopDataExportTaskRequest() (request *StopDataExportTaskRequest) {
	request = &StopDataExportTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "StopDataExportTask", "Drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStopDataExportTaskResponse creates a response to parse from StopDataExportTask response
func CreateStopDataExportTaskResponse() (response *StopDataExportTaskResponse) {
	response = &StopDataExportTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
