package green

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

// DescribeOssResultItems invokes the green.DescribeOssResultItems API synchronously
func (client *Client) DescribeOssResultItems(request *DescribeOssResultItemsRequest) (response *DescribeOssResultItemsResponse, err error) {
	response = CreateDescribeOssResultItemsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeOssResultItemsWithChan invokes the green.DescribeOssResultItems API asynchronously
func (client *Client) DescribeOssResultItemsWithChan(request *DescribeOssResultItemsRequest) (<-chan *DescribeOssResultItemsResponse, <-chan error) {
	responseChan := make(chan *DescribeOssResultItemsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeOssResultItems(request)
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

// DescribeOssResultItemsWithCallback invokes the green.DescribeOssResultItems API asynchronously
func (client *Client) DescribeOssResultItemsWithCallback(request *DescribeOssResultItemsRequest, callback func(response *DescribeOssResultItemsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeOssResultItemsResponse
		var err error
		defer close(result)
		response, err = client.DescribeOssResultItems(request)
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

// DescribeOssResultItemsRequest is the request struct for api DescribeOssResultItems
type DescribeOssResultItemsRequest struct {
	*requests.RpcRequest
	MinScore     requests.Float   `position:"Query" name:"MinScore"`
	MaxScore     requests.Float   `position:"Query" name:"MaxScore"`
	StockTaskId  requests.Integer `position:"Query" name:"StockTaskId"`
	StartDate    string           `position:"Query" name:"StartDate"`
	Scene        string           `position:"Query" name:"Scene"`
	SourceIp     string           `position:"Query" name:"SourceIp"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	Lang         string           `position:"Query" name:"Lang"`
	Stock        requests.Boolean `position:"Query" name:"Stock"`
	TotalCount   requests.Integer `position:"Query" name:"TotalCount"`
	Suggestion   string           `position:"Query" name:"Suggestion"`
	CurrentPage  requests.Integer `position:"Query" name:"CurrentPage"`
	ResourceType string           `position:"Query" name:"ResourceType"`
	QueryId      string           `position:"Query" name:"QueryId"`
	Bucket       string           `position:"Query" name:"Bucket"`
	EndDate      string           `position:"Query" name:"EndDate"`
	Object       string           `position:"Query" name:"Object"`
}

// DescribeOssResultItemsResponse is the response struct for api DescribeOssResultItems
type DescribeOssResultItemsResponse struct {
	*responses.BaseResponse
	RequestId      string       `json:"RequestId" xml:"RequestId"`
	PageSize       int          `json:"PageSize" xml:"PageSize"`
	CurrentPage    int          `json:"CurrentPage" xml:"CurrentPage"`
	TotalCount     int          `json:"TotalCount" xml:"TotalCount"`
	QueryId        string       `json:"QueryId" xml:"QueryId"`
	ScanResultList []ScanResult `json:"ScanResultList" xml:"ScanResultList"`
}

// CreateDescribeOssResultItemsRequest creates a request to invoke DescribeOssResultItems API
func CreateDescribeOssResultItemsRequest() (request *DescribeOssResultItemsRequest) {
	request = &DescribeOssResultItemsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "DescribeOssResultItems", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeOssResultItemsResponse creates a response to parse from DescribeOssResultItems response
func CreateDescribeOssResultItemsResponse() (response *DescribeOssResultItemsResponse) {
	response = &DescribeOssResultItemsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
