package rtc

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

// DescribeRtcScaleDetail invokes the rtc.DescribeRtcScaleDetail API synchronously
func (client *Client) DescribeRtcScaleDetail(request *DescribeRtcScaleDetailRequest) (response *DescribeRtcScaleDetailResponse, err error) {
	response = CreateDescribeRtcScaleDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRtcScaleDetailWithChan invokes the rtc.DescribeRtcScaleDetail API asynchronously
func (client *Client) DescribeRtcScaleDetailWithChan(request *DescribeRtcScaleDetailRequest) (<-chan *DescribeRtcScaleDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeRtcScaleDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRtcScaleDetail(request)
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

// DescribeRtcScaleDetailWithCallback invokes the rtc.DescribeRtcScaleDetail API asynchronously
func (client *Client) DescribeRtcScaleDetailWithCallback(request *DescribeRtcScaleDetailRequest, callback func(response *DescribeRtcScaleDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRtcScaleDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeRtcScaleDetail(request)
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

// DescribeRtcScaleDetailRequest is the request struct for api DescribeRtcScaleDetail
type DescribeRtcScaleDetailRequest struct {
	*requests.RpcRequest
	StartTime string           `position:"Query" name:"StartTime"`
	ShowLog   string           `position:"Query" name:"ShowLog"`
	EndTime   string           `position:"Query" name:"EndTime"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
	AppId     string           `position:"Query" name:"AppId"`
}

// DescribeRtcScaleDetailResponse is the response struct for api DescribeRtcScaleDetail
type DescribeRtcScaleDetailResponse struct {
	*responses.BaseResponse
	RequestId string      `json:"RequestId" xml:"RequestId"`
	Scale     []ScaleItem `json:"Scale" xml:"Scale"`
}

// CreateDescribeRtcScaleDetailRequest creates a request to invoke DescribeRtcScaleDetail API
func CreateDescribeRtcScaleDetailRequest() (request *DescribeRtcScaleDetailRequest) {
	request = &DescribeRtcScaleDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "DescribeRtcScaleDetail", "rtc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeRtcScaleDetailResponse creates a response to parse from DescribeRtcScaleDetail response
func CreateDescribeRtcScaleDetailResponse() (response *DescribeRtcScaleDetailResponse) {
	response = &DescribeRtcScaleDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
