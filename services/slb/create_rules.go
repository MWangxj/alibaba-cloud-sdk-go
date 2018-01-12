package slb

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

func (client *Client) CreateRules(request *CreateRulesRequest) (response *CreateRulesResponse, err error) {
	response = CreateCreateRulesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateRulesWithChan(request *CreateRulesRequest) (<-chan *CreateRulesResponse, <-chan error) {
	responseChan := make(chan *CreateRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRules(request)
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

func (client *Client) CreateRulesWithCallback(request *CreateRulesRequest, callback func(response *CreateRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRulesResponse
		var err error
		defer close(result)
		response, err = client.CreateRules(request)
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

type CreateRulesRequest struct {
	*requests.RpcRequest
	Tags                 string           `position:"Query" name:"Tags"`
	ListenerPort         requests.Integer `position:"Query" name:"ListenerPort"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	RuleList             string           `position:"Query" name:"RuleList"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type CreateRulesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Rules     struct {
		Rule []struct {
			RuleId   string `json:"RuleId" xml:"RuleId"`
			RuleName string `json:"RuleName" xml:"RuleName"`
		} `json:"Rule" xml:"Rule"`
	} `json:"Rules" xml:"Rules"`
}

func CreateCreateRulesRequest() (request *CreateRulesRequest) {
	request = &CreateRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "CreateRules", "slb", "openAPI")
	return
}

func CreateCreateRulesResponse() (response *CreateRulesResponse) {
	response = &CreateRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
