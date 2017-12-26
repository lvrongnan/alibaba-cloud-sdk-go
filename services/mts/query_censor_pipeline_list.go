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

func (client *Client) QueryCensorPipelineList(request *QueryCensorPipelineListRequest) (response *QueryCensorPipelineListResponse, err error) {
	response = CreateQueryCensorPipelineListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryCensorPipelineListWithChan(request *QueryCensorPipelineListRequest) (<-chan *QueryCensorPipelineListResponse, <-chan error) {
	responseChan := make(chan *QueryCensorPipelineListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryCensorPipelineList(request)
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

func (client *Client) QueryCensorPipelineListWithCallback(request *QueryCensorPipelineListRequest, callback func(response *QueryCensorPipelineListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryCensorPipelineListResponse
		var err error
		defer close(result)
		response, err = client.QueryCensorPipelineList(request)
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

type QueryCensorPipelineListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PipelineIds          string `position:"Query" name:"PipelineIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Action               string `position:"Query" name:"Action"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	AccessKeyId          string `position:"Query" name:"AccessKeyId"`
}

type QueryCensorPipelineListResponse struct {
	*responses.BaseResponse
	RequestId    string   `json:"RequestId"`
	NonExistIds  []string `json:"NonExistIds"`
	PipelineList []struct {
		Id           string `json:"Id"`
		Name         string `json:"Name"`
		State        string `json:"State"`
		Priority     string `json:"Priority"`
		NotifyConfig struct {
			Topic string `json:"Topic"`
			Queue string `json:"Queue"`
		} `json:"NotifyConfig"`
	} `json:"PipelineList"`
}

func CreateQueryCensorPipelineListRequest() (request *QueryCensorPipelineListRequest) {
	request = &QueryCensorPipelineListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryCensorPipelineList", "", "")
	return
}

func CreateQueryCensorPipelineListResponse() (response *QueryCensorPipelineListResponse) {
	response = &QueryCensorPipelineListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
