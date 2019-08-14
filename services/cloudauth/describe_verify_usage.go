package cloudauth

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

// DescribeVerifyUsage invokes the cloudauth.DescribeVerifyUsage API synchronously
// api document: https://help.aliyun.com/api/cloudauth/describeverifyusage.html
func (client *Client) DescribeVerifyUsage(request *DescribeVerifyUsageRequest) (response *DescribeVerifyUsageResponse, err error) {
	response = CreateDescribeVerifyUsageResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVerifyUsageWithChan invokes the cloudauth.DescribeVerifyUsage API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/describeverifyusage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVerifyUsageWithChan(request *DescribeVerifyUsageRequest) (<-chan *DescribeVerifyUsageResponse, <-chan error) {
	responseChan := make(chan *DescribeVerifyUsageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVerifyUsage(request)
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

// DescribeVerifyUsageWithCallback invokes the cloudauth.DescribeVerifyUsage API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/describeverifyusage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVerifyUsageWithCallback(request *DescribeVerifyUsageRequest, callback func(response *DescribeVerifyUsageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVerifyUsageResponse
		var err error
		defer close(result)
		response, err = client.DescribeVerifyUsage(request)
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

// DescribeVerifyUsageRequest is the request struct for api DescribeVerifyUsage
type DescribeVerifyUsageRequest struct {
	*requests.RpcRequest
	BizType   string `position:"Query" name:"BizType"`
	EndDate   string `position:"Query" name:"EndDate"`
	SourceIp  string `position:"Query" name:"SourceIp"`
	StartDate string `position:"Query" name:"StartDate"`
}

// DescribeVerifyUsageResponse is the response struct for api DescribeVerifyUsage
type DescribeVerifyUsageResponse struct {
	*responses.BaseResponse
	RequestId       string        `json:"RequestId" xml:"RequestId"`
	TotalCount      int           `json:"TotalCount" xml:"TotalCount"`
	VerifyUsageList []VerifyUsage `json:"VerifyUsageList" xml:"VerifyUsageList"`
}

// CreateDescribeVerifyUsageRequest creates a request to invoke DescribeVerifyUsage API
func CreateDescribeVerifyUsageRequest() (request *DescribeVerifyUsageRequest) {
	request = &DescribeVerifyUsageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "DescribeVerifyUsage", "cloudauth", "openAPI")
	return
}

// CreateDescribeVerifyUsageResponse creates a response to parse from DescribeVerifyUsage response
func CreateDescribeVerifyUsageResponse() (response *DescribeVerifyUsageResponse) {
	response = &DescribeVerifyUsageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}