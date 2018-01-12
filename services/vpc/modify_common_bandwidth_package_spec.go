package vpc

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

func (client *Client) ModifyCommonBandwidthPackageSpec(request *ModifyCommonBandwidthPackageSpecRequest) (response *ModifyCommonBandwidthPackageSpecResponse, err error) {
	response = CreateModifyCommonBandwidthPackageSpecResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyCommonBandwidthPackageSpecWithChan(request *ModifyCommonBandwidthPackageSpecRequest) (<-chan *ModifyCommonBandwidthPackageSpecResponse, <-chan error) {
	responseChan := make(chan *ModifyCommonBandwidthPackageSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyCommonBandwidthPackageSpec(request)
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

func (client *Client) ModifyCommonBandwidthPackageSpecWithCallback(request *ModifyCommonBandwidthPackageSpecRequest, callback func(response *ModifyCommonBandwidthPackageSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyCommonBandwidthPackageSpecResponse
		var err error
		defer close(result)
		response, err = client.ModifyCommonBandwidthPackageSpec(request)
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

type ModifyCommonBandwidthPackageSpecRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Bandwidth            string           `position:"Query" name:"Bandwidth"`
	BandwidthPackageId   string           `position:"Query" name:"BandwidthPackageId"`
}

type ModifyCommonBandwidthPackageSpecResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateModifyCommonBandwidthPackageSpecRequest() (request *ModifyCommonBandwidthPackageSpecRequest) {
	request = &ModifyCommonBandwidthPackageSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyCommonBandwidthPackageSpec", "vpc", "openAPI")
	return
}

func CreateModifyCommonBandwidthPackageSpecResponse() (response *ModifyCommonBandwidthPackageSpecResponse) {
	response = &ModifyCommonBandwidthPackageSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
