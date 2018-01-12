package mts

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

func (client *Client) QueryTagJobList(request *QueryTagJobListRequest) (response *QueryTagJobListResponse, err error) {
	response = CreateQueryTagJobListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryTagJobListWithChan(request *QueryTagJobListRequest) (<-chan *QueryTagJobListResponse, <-chan error) {
	responseChan := make(chan *QueryTagJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTagJobList(request)
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

func (client *Client) QueryTagJobListWithCallback(request *QueryTagJobListRequest, callback func(response *QueryTagJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTagJobListResponse
		var err error
		defer close(result)
		response, err = client.QueryTagJobList(request)
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

type QueryTagJobListRequest struct {
	*requests.RpcRequest
	TagJobIds            string           `position:"Query" name:"TagJobIds"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type QueryTagJobListResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	NonExistIds struct {
		String []string `json:"String" xml:"String"`
	} `json:"NonExistIds" xml:"NonExistIds"`
	TagJobList struct {
		TagJob []struct {
			Id           string `json:"Id" xml:"Id"`
			UserData     string `json:"UserData" xml:"UserData"`
			PipelineId   string `json:"PipelineId" xml:"PipelineId"`
			State        string `json:"State" xml:"State"`
			Code         string `json:"Code" xml:"Code"`
			Message      string `json:"Message" xml:"Message"`
			CreationTime string `json:"CreationTime" xml:"CreationTime"`
			Input        struct {
				Bucket   string `json:"Bucket" xml:"Bucket"`
				Location string `json:"Location" xml:"Location"`
				Object   string `json:"Object" xml:"Object"`
			} `json:"Input" xml:"Input"`
			VideoTagResult struct {
				Details      string `json:"Details" xml:"Details"`
				TagAnResults struct {
					TagAnResult []struct {
						Label string `json:"Label" xml:"Label"`
						Score string `json:"Score" xml:"Score"`
					} `json:"TagAnResult" xml:"TagAnResult"`
				} `json:"TagAnResults" xml:"TagAnResults"`
				TagFrResults struct {
					TagFrResult []struct {
						Time     string `json:"Time" xml:"Time"`
						TagFaces struct {
							TagFace []struct {
								Name   string `json:"Name" xml:"Name"`
								Score  string `json:"Score" xml:"Score"`
								Target string `json:"Target" xml:"Target"`
							} `json:"TagFace" xml:"TagFace"`
						} `json:"TagFaces" xml:"TagFaces"`
					} `json:"TagFrResult" xml:"TagFrResult"`
				} `json:"TagFrResults" xml:"TagFrResults"`
			} `json:"VideoTagResult" xml:"VideoTagResult"`
		} `json:"TagJob" xml:"TagJob"`
	} `json:"TagJobList" xml:"TagJobList"`
}

func CreateQueryTagJobListRequest() (request *QueryTagJobListRequest) {
	request = &QueryTagJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryTagJobList", "mts", "openAPI")
	return
}

func CreateQueryTagJobListResponse() (response *QueryTagJobListResponse) {
	response = &QueryTagJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
