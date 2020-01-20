package iot

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

// CreateProduct invokes the iot.CreateProduct API synchronously
// api document: https://help.aliyun.com/api/iot/createproduct.html
func (client *Client) CreateProduct(request *CreateProductRequest) (response *CreateProductResponse, err error) {
	response = CreateCreateProductResponse()
	err = client.DoAction(request, response)
	return
}

// CreateProductWithChan invokes the iot.CreateProduct API asynchronously
// api document: https://help.aliyun.com/api/iot/createproduct.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateProductWithChan(request *CreateProductRequest) (<-chan *CreateProductResponse, <-chan error) {
	responseChan := make(chan *CreateProductResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateProduct(request)
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

// CreateProductWithCallback invokes the iot.CreateProduct API asynchronously
// api document: https://help.aliyun.com/api/iot/createproduct.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateProductWithCallback(request *CreateProductRequest, callback func(response *CreateProductResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateProductResponse
		var err error
		defer close(result)
		response, err = client.CreateProduct(request)
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

// CreateProductRequest is the request struct for api CreateProduct
type CreateProductRequest struct {
	*requests.RpcRequest
	NodeType            requests.Integer `position:"Query" name:"NodeType"`
	Description         string           `position:"Query" name:"Description"`
	CategoryKey         string           `position:"Query" name:"CategoryKey"`
	JoinPermissionId    string           `position:"Query" name:"JoinPermissionId"`
	AuthType            string           `position:"Query" name:"AuthType"`
	ResourceGroupId     string           `position:"Query" name:"ResourceGroupId"`
	IotInstanceId       string           `position:"Query" name:"IotInstanceId"`
	ProductName         string           `position:"Query" name:"ProductName"`
	AliyunCommodityCode string           `position:"Query" name:"AliyunCommodityCode"`
	PublishAuto         requests.Boolean `position:"Query" name:"PublishAuto"`
	CategoryId          requests.Integer `position:"Query" name:"CategoryId"`
	DataFormat          requests.Integer `position:"Query" name:"DataFormat"`
	Id2                 requests.Boolean `position:"Query" name:"Id2"`
	NetType             string           `position:"Query" name:"NetType"`
	ProtocolType        string           `position:"Query" name:"ProtocolType"`
}

// CreateProductResponse is the response struct for api CreateProduct
type CreateProductResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	ProductKey   string `json:"ProductKey" xml:"ProductKey"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateCreateProductRequest creates a request to invoke CreateProduct API
func CreateCreateProductRequest() (request *CreateProductRequest) {
	request = &CreateProductRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "CreateProduct", "iot", "openAPI")
	return
}

// CreateCreateProductResponse creates a response to parse from CreateProduct response
func CreateCreateProductResponse() (response *CreateProductResponse) {
	response = &CreateProductResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
