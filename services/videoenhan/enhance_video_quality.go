package videoenhan

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

// EnhanceVideoQuality invokes the videoenhan.EnhanceVideoQuality API synchronously
func (client *Client) EnhanceVideoQuality(request *EnhanceVideoQualityRequest) (response *EnhanceVideoQualityResponse, err error) {
	response = CreateEnhanceVideoQualityResponse()
	err = client.DoAction(request, response)
	return
}

// EnhanceVideoQualityWithChan invokes the videoenhan.EnhanceVideoQuality API asynchronously
func (client *Client) EnhanceVideoQualityWithChan(request *EnhanceVideoQualityRequest) (<-chan *EnhanceVideoQualityResponse, <-chan error) {
	responseChan := make(chan *EnhanceVideoQualityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnhanceVideoQuality(request)
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

// EnhanceVideoQualityWithCallback invokes the videoenhan.EnhanceVideoQuality API asynchronously
func (client *Client) EnhanceVideoQualityWithCallback(request *EnhanceVideoQualityRequest, callback func(response *EnhanceVideoQualityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnhanceVideoQualityResponse
		var err error
		defer close(result)
		response, err = client.EnhanceVideoQuality(request)
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

// EnhanceVideoQualityRequest is the request struct for api EnhanceVideoQuality
type EnhanceVideoQualityRequest struct {
	*requests.RpcRequest
	HDRFormat      string           `position:"Body" name:"HDRFormat"`
	FrameRate      requests.Integer `position:"Body" name:"FrameRate"`
	MaxIlluminance requests.Integer `position:"Body" name:"MaxIlluminance"`
	Bitrate        requests.Integer `position:"Body" name:"Bitrate"`
	OutPutWidth    requests.Integer `position:"Body" name:"OutPutWidth"`
	OutPutHeight   requests.Integer `position:"Body" name:"OutPutHeight"`
	Async          requests.Boolean `position:"Body" name:"Async"`
	VideoURL       string           `position:"Body" name:"VideoURL"`
}

// EnhanceVideoQualityResponse is the response struct for api EnhanceVideoQuality
type EnhanceVideoQualityResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateEnhanceVideoQualityRequest creates a request to invoke EnhanceVideoQuality API
func CreateEnhanceVideoQualityRequest() (request *EnhanceVideoQualityRequest) {
	request = &EnhanceVideoQualityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("videoenhan", "2020-03-20", "EnhanceVideoQuality", "videoenhan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnhanceVideoQualityResponse creates a response to parse from EnhanceVideoQuality response
func CreateEnhanceVideoQualityResponse() (response *EnhanceVideoQualityResponse) {
	response = &EnhanceVideoQualityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}