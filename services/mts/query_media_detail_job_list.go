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

func (client *Client) QueryMediaDetailJobList(request *QueryMediaDetailJobListRequest) (response *QueryMediaDetailJobListResponse, err error) {
	response = CreateQueryMediaDetailJobListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryMediaDetailJobListWithChan(request *QueryMediaDetailJobListRequest) (<-chan *QueryMediaDetailJobListResponse, <-chan error) {
	responseChan := make(chan *QueryMediaDetailJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMediaDetailJobList(request)
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

func (client *Client) QueryMediaDetailJobListWithCallback(request *QueryMediaDetailJobListRequest, callback func(response *QueryMediaDetailJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMediaDetailJobListResponse
		var err error
		defer close(result)
		response, err = client.QueryMediaDetailJobList(request)
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

type QueryMediaDetailJobListRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string           `position:"Query" name:"JobIds"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type QueryMediaDetailJobListResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	NonExistIds struct {
		String []string `json:"String" xml:"String"`
	} `json:"NonExistIds" xml:"NonExistIds"`
	JobList struct {
		Job []struct {
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
			MediaDetailConfig struct {
				Scenario   string `json:"Scenario" xml:"Scenario"`
				DetailType string `json:"DetailType" xml:"DetailType"`
				OutputFile struct {
					Bucket   string `json:"Bucket" xml:"Bucket"`
					Location string `json:"Location" xml:"Location"`
					Object   string `json:"Object" xml:"Object"`
				} `json:"OutputFile" xml:"OutputFile"`
			} `json:"MediaDetailConfig" xml:"MediaDetailConfig"`
			MediaDetailResult struct {
				Status string `json:"Status" xml:"Status"`
				Tags   struct {
					String []string `json:"String" xml:"String"`
				} `json:"Tags" xml:"Tags"`
				MediaDetailRecgResults struct {
					MediaDetailRecgResult []struct {
						ImageUrl  string `json:"ImageUrl" xml:"ImageUrl"`
						Time      string `json:"Time" xml:"Time"`
						OcrText   string `json:"OcrText" xml:"OcrText"`
						FrameTags struct {
							String []string `json:"String" xml:"String"`
						} `json:"FrameTags" xml:"FrameTags"`
						Celebrities struct {
							Celebrity []struct {
								Name   string `json:"Name" xml:"Name"`
								Score  string `json:"Score" xml:"Score"`
								Target string `json:"Target" xml:"Target"`
							} `json:"Celebrity" xml:"Celebrity"`
						} `json:"Celebrities" xml:"Celebrities"`
						Sensitives struct {
							Sensitive []struct {
								Name   string `json:"Name" xml:"Name"`
								Score  string `json:"Score" xml:"Score"`
								Target string `json:"Target" xml:"Target"`
							} `json:"Sensitive" xml:"Sensitive"`
						} `json:"Sensitives" xml:"Sensitives"`
						Politicians struct {
							Politician []struct {
								Name   string `json:"Name" xml:"Name"`
								Score  string `json:"Score" xml:"Score"`
								Target string `json:"Target" xml:"Target"`
							} `json:"Politician" xml:"Politician"`
						} `json:"Politicians" xml:"Politicians"`
						FrameTagInfos struct {
							FrameTagInfo []struct {
								Tag      string `json:"Tag" xml:"Tag"`
								Score    string `json:"Score" xml:"Score"`
								Category string `json:"Category" xml:"Category"`
							} `json:"FrameTagInfo" xml:"FrameTagInfo"`
						} `json:"FrameTagInfos" xml:"FrameTagInfos"`
					} `json:"MediaDetailRecgResult" xml:"MediaDetailRecgResult"`
				} `json:"MediaDetailRecgResults" xml:"MediaDetailRecgResults"`
			} `json:"MediaDetailResult" xml:"MediaDetailResult"`
		} `json:"Job" xml:"Job"`
	} `json:"JobList" xml:"JobList"`
}

func CreateQueryMediaDetailJobListRequest() (request *QueryMediaDetailJobListRequest) {
	request = &QueryMediaDetailJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryMediaDetailJobList", "mts", "openAPI")
	return
}

func CreateQueryMediaDetailJobListResponse() (response *QueryMediaDetailJobListResponse) {
	response = &QueryMediaDetailJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
