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

func (client *Client) UpdateMediaCover(request *UpdateMediaCoverRequest) (response *UpdateMediaCoverResponse, err error) {
	response = CreateUpdateMediaCoverResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) UpdateMediaCoverWithChan(request *UpdateMediaCoverRequest) (<-chan *UpdateMediaCoverResponse, <-chan error) {
	responseChan := make(chan *UpdateMediaCoverResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateMediaCover(request)
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

func (client *Client) UpdateMediaCoverWithCallback(request *UpdateMediaCoverRequest, callback func(response *UpdateMediaCoverResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateMediaCoverResponse
		var err error
		defer close(result)
		response, err = client.UpdateMediaCover(request)
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

type UpdateMediaCoverRequest struct {
	*requests.RpcRequest
	CoverURL             string `position:"Query" name:"CoverURL"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Action               string `position:"Query" name:"Action"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
	AccessKeyId          string `position:"Query" name:"AccessKeyId"`
}

type UpdateMediaCoverResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId"`
}

func CreateUpdateMediaCoverRequest() (request *UpdateMediaCoverRequest) {
	request = &UpdateMediaCoverRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "UpdateMediaCover", "", "")
	return
}

func CreateUpdateMediaCoverResponse() (response *UpdateMediaCoverResponse) {
	response = &UpdateMediaCoverResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
