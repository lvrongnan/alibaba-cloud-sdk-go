package cassandra

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

// DescribeParameterModificationHistories invokes the cassandra.DescribeParameterModificationHistories API synchronously
// api document: https://help.aliyun.com/api/cassandra/describeparametermodificationhistories.html
func (client *Client) DescribeParameterModificationHistories(request *DescribeParameterModificationHistoriesRequest) (response *DescribeParameterModificationHistoriesResponse, err error) {
	response = CreateDescribeParameterModificationHistoriesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeParameterModificationHistoriesWithChan invokes the cassandra.DescribeParameterModificationHistories API asynchronously
// api document: https://help.aliyun.com/api/cassandra/describeparametermodificationhistories.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeParameterModificationHistoriesWithChan(request *DescribeParameterModificationHistoriesRequest) (<-chan *DescribeParameterModificationHistoriesResponse, <-chan error) {
	responseChan := make(chan *DescribeParameterModificationHistoriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeParameterModificationHistories(request)
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

// DescribeParameterModificationHistoriesWithCallback invokes the cassandra.DescribeParameterModificationHistories API asynchronously
// api document: https://help.aliyun.com/api/cassandra/describeparametermodificationhistories.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeParameterModificationHistoriesWithCallback(request *DescribeParameterModificationHistoriesRequest, callback func(response *DescribeParameterModificationHistoriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeParameterModificationHistoriesResponse
		var err error
		defer close(result)
		response, err = client.DescribeParameterModificationHistories(request)
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

// DescribeParameterModificationHistoriesRequest is the request struct for api DescribeParameterModificationHistories
type DescribeParameterModificationHistoriesRequest struct {
	*requests.RpcRequest
	ClusterId  string           `position:"Query" name:"ClusterId"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeParameterModificationHistoriesResponse is the response struct for api DescribeParameterModificationHistories
type DescribeParameterModificationHistoriesResponse struct {
	*responses.BaseResponse
	RequestId  string                                            `json:"RequestId" xml:"RequestId"`
	PageNumber int                                               `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                                               `json:"PageSize" xml:"PageSize"`
	TotalCount int64                                             `json:"TotalCount" xml:"TotalCount"`
	Histories  HistoriesInDescribeParameterModificationHistories `json:"Histories" xml:"Histories"`
}

// CreateDescribeParameterModificationHistoriesRequest creates a request to invoke DescribeParameterModificationHistories API
func CreateDescribeParameterModificationHistoriesRequest() (request *DescribeParameterModificationHistoriesRequest) {
	request = &DescribeParameterModificationHistoriesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cassandra", "2019-01-01", "DescribeParameterModificationHistories", "Cassandra", "openAPI")
	return
}

// CreateDescribeParameterModificationHistoriesResponse creates a response to parse from DescribeParameterModificationHistories response
func CreateDescribeParameterModificationHistoriesResponse() (response *DescribeParameterModificationHistoriesResponse) {
	response = &DescribeParameterModificationHistoriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}