package cs

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

func (client *Client) DownloadClusterNodeCerts(request *DownloadClusterNodeCertsRequest) (response *DownloadClusterNodeCertsResponse, err error) {
	response = CreateDownloadClusterNodeCertsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DownloadClusterNodeCertsWithChan(request *DownloadClusterNodeCertsRequest) (<-chan *DownloadClusterNodeCertsResponse, <-chan error) {
	responseChan := make(chan *DownloadClusterNodeCertsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DownloadClusterNodeCerts(request)
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

func (client *Client) DownloadClusterNodeCertsWithCallback(request *DownloadClusterNodeCertsRequest, callback func(response *DownloadClusterNodeCertsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DownloadClusterNodeCertsResponse
		var err error
		defer close(result)
		response, err = client.DownloadClusterNodeCerts(request)
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

type DownloadClusterNodeCertsRequest struct {
	*requests.RoaRequest
	Token  string `position:"Path" name:"Token"`
	NodeId string `position:"Path" name:"NodeId"`
}

type DownloadClusterNodeCertsResponse struct {
	*responses.BaseResponse
}

func CreateDownloadClusterNodeCertsRequest() (request *DownloadClusterNodeCertsRequest) {
	request = &DownloadClusterNodeCertsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "DownloadClusterNodeCerts", "/token/[Token]/nodes/[NodeId]/certs", "", "")
	request.Method = requests.GET
	return
}

func CreateDownloadClusterNodeCertsResponse() (response *DownloadClusterNodeCertsResponse) {
	response = &DownloadClusterNodeCertsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
