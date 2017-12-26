package mts

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

func (client *Client) UpdateMediaWorkflow(request *UpdateMediaWorkflowRequest) (response *UpdateMediaWorkflowResponse, err error) {
	response = CreateUpdateMediaWorkflowResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) UpdateMediaWorkflowWithChan(request *UpdateMediaWorkflowRequest) (<-chan *UpdateMediaWorkflowResponse, <-chan error) {
	responseChan := make(chan *UpdateMediaWorkflowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateMediaWorkflow(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) UpdateMediaWorkflowWithCallback(request *UpdateMediaWorkflowRequest, callback func(response *UpdateMediaWorkflowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateMediaWorkflowResponse
		var err error
		defer close(result)
		response, err = client.UpdateMediaWorkflow(request)
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

type UpdateMediaWorkflowRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Topology             string `position:"Query" name:"Topology"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Action               string `position:"Query" name:"Action"`
	MediaWorkflowId      string `position:"Query" name:"MediaWorkflowId"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	AccessKeyId          string `position:"Query" name:"AccessKeyId"`
}

type UpdateMediaWorkflowResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId"`
	MediaWorkflow struct {
		MediaWorkflowId string `json:"MediaWorkflowId"`
		Name            string `json:"Name"`
		Topology        string `json:"Topology"`
		State           string `json:"State"`
		CreationTime    string `json:"CreationTime"`
	} `json:"MediaWorkflow"`
}

func CreateUpdateMediaWorkflowRequest() (request *UpdateMediaWorkflowRequest) {
	request = &UpdateMediaWorkflowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "UpdateMediaWorkflow", "", "")
	return
}

func CreateUpdateMediaWorkflowResponse() (response *UpdateMediaWorkflowResponse) {
	response = &UpdateMediaWorkflowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
