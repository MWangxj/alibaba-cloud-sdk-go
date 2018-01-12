package ecs

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

func (client *Client) DescribeSecurityGroupAttribute(request *DescribeSecurityGroupAttributeRequest) (response *DescribeSecurityGroupAttributeResponse, err error) {
	response = CreateDescribeSecurityGroupAttributeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeSecurityGroupAttributeWithChan(request *DescribeSecurityGroupAttributeRequest) (<-chan *DescribeSecurityGroupAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeSecurityGroupAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSecurityGroupAttribute(request)
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

func (client *Client) DescribeSecurityGroupAttributeWithCallback(request *DescribeSecurityGroupAttributeRequest, callback func(response *DescribeSecurityGroupAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSecurityGroupAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeSecurityGroupAttribute(request)
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

type DescribeSecurityGroupAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	SecurityGroupId      string           `position:"Query" name:"SecurityGroupId"`
	Direction            string           `position:"Query" name:"Direction"`
	NicType              string           `position:"Query" name:"NicType"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type DescribeSecurityGroupAttributeResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	RegionId          string `json:"RegionId" xml:"RegionId"`
	SecurityGroupId   string `json:"SecurityGroupId" xml:"SecurityGroupId"`
	Description       string `json:"Description" xml:"Description"`
	SecurityGroupName string `json:"SecurityGroupName" xml:"SecurityGroupName"`
	VpcId             string `json:"VpcId" xml:"VpcId"`
	InnerAccessPolicy string `json:"InnerAccessPolicy" xml:"InnerAccessPolicy"`
	Permissions       struct {
		Permission []struct {
			IpProtocol              string `json:"IpProtocol" xml:"IpProtocol"`
			PortRange               string `json:"PortRange" xml:"PortRange"`
			SourceGroupId           string `json:"SourceGroupId" xml:"SourceGroupId"`
			SourceGroupName         string `json:"SourceGroupName" xml:"SourceGroupName"`
			SourceCidrIp            string `json:"SourceCidrIp" xml:"SourceCidrIp"`
			Policy                  string `json:"Policy" xml:"Policy"`
			NicType                 string `json:"NicType" xml:"NicType"`
			SourceGroupOwnerAccount string `json:"SourceGroupOwnerAccount" xml:"SourceGroupOwnerAccount"`
			DestGroupId             string `json:"DestGroupId" xml:"DestGroupId"`
			DestGroupName           string `json:"DestGroupName" xml:"DestGroupName"`
			DestCidrIp              string `json:"DestCidrIp" xml:"DestCidrIp"`
			DestGroupOwnerAccount   string `json:"DestGroupOwnerAccount" xml:"DestGroupOwnerAccount"`
			Priority                string `json:"Priority" xml:"Priority"`
			Direction               string `json:"Direction" xml:"Direction"`
			Description             string `json:"Description" xml:"Description"`
			CreateTime              string `json:"CreateTime" xml:"CreateTime"`
		} `json:"Permission" xml:"Permission"`
	} `json:"Permissions" xml:"Permissions"`
}

func CreateDescribeSecurityGroupAttributeRequest() (request *DescribeSecurityGroupAttributeRequest) {
	request = &DescribeSecurityGroupAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeSecurityGroupAttribute", "ecs", "openAPI")
	return
}

func CreateDescribeSecurityGroupAttributeResponse() (response *DescribeSecurityGroupAttributeResponse) {
	response = &DescribeSecurityGroupAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
