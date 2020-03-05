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

// UpdateEdgeDriverVersion invokes the iot.UpdateEdgeDriverVersion API synchronously
// api document: https://help.aliyun.com/api/iot/updateedgedriverversion.html
func (client *Client) UpdateEdgeDriverVersion(request *UpdateEdgeDriverVersionRequest) (response *UpdateEdgeDriverVersionResponse, err error) {
	response = CreateUpdateEdgeDriverVersionResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateEdgeDriverVersionWithChan invokes the iot.UpdateEdgeDriverVersion API asynchronously
// api document: https://help.aliyun.com/api/iot/updateedgedriverversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateEdgeDriverVersionWithChan(request *UpdateEdgeDriverVersionRequest) (<-chan *UpdateEdgeDriverVersionResponse, <-chan error) {
	responseChan := make(chan *UpdateEdgeDriverVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateEdgeDriverVersion(request)
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

// UpdateEdgeDriverVersionWithCallback invokes the iot.UpdateEdgeDriverVersion API asynchronously
// api document: https://help.aliyun.com/api/iot/updateedgedriverversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateEdgeDriverVersionWithCallback(request *UpdateEdgeDriverVersionRequest, callback func(response *UpdateEdgeDriverVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateEdgeDriverVersionResponse
		var err error
		defer close(result)
		response, err = client.UpdateEdgeDriverVersion(request)
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

// UpdateEdgeDriverVersionRequest is the request struct for api UpdateEdgeDriverVersion
type UpdateEdgeDriverVersionRequest struct {
	*requests.RpcRequest
	ConfigCheckRule string `position:"Query" name:"ConfigCheckRule"`
	EdgeVersion     string `position:"Query" name:"EdgeVersion"`
	Description     string `position:"Query" name:"Description"`
	DriverId        string `position:"Query" name:"DriverId"`
	IotInstanceId   string `position:"Query" name:"IotInstanceId"`
	ContainerConfig string `position:"Query" name:"ContainerConfig"`
	DriverVersion   string `position:"Query" name:"DriverVersion"`
	DriverConfig    string `position:"Query" name:"DriverConfig"`
	SourceConfig    string `position:"Query" name:"SourceConfig"`
	ApiProduct      string `position:"Body" name:"ApiProduct"`
	ApiRevision     string `position:"Body" name:"ApiRevision"`
}

// UpdateEdgeDriverVersionResponse is the response struct for api UpdateEdgeDriverVersion
type UpdateEdgeDriverVersionResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateUpdateEdgeDriverVersionRequest creates a request to invoke UpdateEdgeDriverVersion API
func CreateUpdateEdgeDriverVersionRequest() (request *UpdateEdgeDriverVersionRequest) {
	request = &UpdateEdgeDriverVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "UpdateEdgeDriverVersion", "iot", "openAPI")
	return
}

// CreateUpdateEdgeDriverVersionResponse creates a response to parse from UpdateEdgeDriverVersion response
func CreateUpdateEdgeDriverVersionResponse() (response *UpdateEdgeDriverVersionResponse) {
	response = &UpdateEdgeDriverVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
