package rds

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

func (client *Client) ReleaseReplica(request *ReleaseReplicaRequest) (response *ReleaseReplicaResponse, err error) {
	response = CreateReleaseReplicaResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ReleaseReplicaWithChan(request *ReleaseReplicaRequest) (<-chan *ReleaseReplicaResponse, <-chan error) {
	responseChan := make(chan *ReleaseReplicaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleaseReplica(request)
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

func (client *Client) ReleaseReplicaWithCallback(request *ReleaseReplicaRequest, callback func(response *ReleaseReplicaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleaseReplicaResponse
		var err error
		defer close(result)
		response, err = client.ReleaseReplica(request)
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

type ReleaseReplicaRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ReplicaId            string           `position:"Query" name:"ReplicaId"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
}

type ReleaseReplicaResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateReleaseReplicaRequest() (request *ReleaseReplicaRequest) {
	request = &ReleaseReplicaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ReleaseReplica", "rds", "openAPI")
	return
}

func CreateReleaseReplicaResponse() (response *ReleaseReplicaResponse) {
	response = &ReleaseReplicaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
