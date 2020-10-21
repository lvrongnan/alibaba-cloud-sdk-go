package facebody

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

// FaceMakeup invokes the facebody.FaceMakeup API synchronously
func (client *Client) FaceMakeup(request *FaceMakeupRequest) (response *FaceMakeupResponse, err error) {
	response = CreateFaceMakeupResponse()
	err = client.DoAction(request, response)
	return
}

// FaceMakeupWithChan invokes the facebody.FaceMakeup API asynchronously
func (client *Client) FaceMakeupWithChan(request *FaceMakeupRequest) (<-chan *FaceMakeupResponse, <-chan error) {
	responseChan := make(chan *FaceMakeupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FaceMakeup(request)
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

// FaceMakeupWithCallback invokes the facebody.FaceMakeup API asynchronously
func (client *Client) FaceMakeupWithCallback(request *FaceMakeupRequest, callback func(response *FaceMakeupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FaceMakeupResponse
		var err error
		defer close(result)
		response, err = client.FaceMakeup(request)
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

// FaceMakeupRequest is the request struct for api FaceMakeup
type FaceMakeupRequest struct {
	*requests.RpcRequest
	Strength     requests.Float `position:"Body" name:"Strength"`
	MakeupType   string         `position:"Body" name:"MakeupType"`
	ResourceType string         `position:"Body" name:"ResourceType"`
	ImageURL     string         `position:"Body" name:"ImageURL"`
}

// FaceMakeupResponse is the response struct for api FaceMakeup
type FaceMakeupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateFaceMakeupRequest creates a request to invoke FaceMakeup API
func CreateFaceMakeupRequest() (request *FaceMakeupRequest) {
	request = &FaceMakeupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "FaceMakeup", "facebody", "openAPI")
	request.Method = requests.POST
	return
}

// CreateFaceMakeupResponse creates a response to parse from FaceMakeup response
func CreateFaceMakeupResponse() (response *FaceMakeupResponse) {
	response = &FaceMakeupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}