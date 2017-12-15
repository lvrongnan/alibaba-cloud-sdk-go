package cdn

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

func (client *Client) DescribeDomainSrcFlowData(request *DescribeDomainSrcFlowDataRequest) (response *DescribeDomainSrcFlowDataResponse, err error) {
	response = CreateDescribeDomainSrcFlowDataResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDomainSrcFlowDataWithChan(request *DescribeDomainSrcFlowDataRequest) (<-chan *DescribeDomainSrcFlowDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainSrcFlowDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainSrcFlowData(request)
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

func (client *Client) DescribeDomainSrcFlowDataWithCallback(request *DescribeDomainSrcFlowDataRequest, callback func(response *DescribeDomainSrcFlowDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainSrcFlowDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainSrcFlowData(request)
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

type DescribeDomainSrcFlowDataRequest struct {
	*requests.RpcRequest
	FixTimeGap    string `position:"Query" name:"FixTimeGap"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	TimeMerge     string `position:"Query" name:"TimeMerge"`
	DomainName    string `position:"Query" name:"DomainName"`
	Action        string `position:"Query" name:"Action"`
	EndTime       string `position:"Query" name:"EndTime"`
	Interval      string `position:"Query" name:"Interval"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	AccessKeyId   string `position:"Query" name:"AccessKeyId"`
}

type DescribeDomainSrcFlowDataResponse struct {
	*responses.BaseResponse
	RequestId              string `json:"RequestId"`
	DomainName             string `json:"DomainName"`
	DataInterval           string `json:"DataInterval"`
	StartTime              string `json:"StartTime"`
	EndTime                string `json:"EndTime"`
	SrcFlowDataPerInterval []struct {
		TimeStamp string `json:"TimeStamp"`
		Value     string `json:"Value"`
	} `json:"SrcFlowDataPerInterval"`
}

func CreateDescribeDomainSrcFlowDataRequest() (request *DescribeDomainSrcFlowDataRequest) {
	request = &DescribeDomainSrcFlowDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainSrcFlowData", "", "")
	return
}

func CreateDescribeDomainSrcFlowDataResponse() (response *DescribeDomainSrcFlowDataResponse) {
	response = &DescribeDomainSrcFlowDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}