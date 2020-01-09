package foas

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

// Package is a nested struct in foas response
type Package struct {
	ModifyTime  int64  `json:"ModifyTime" xml:"ModifyTime"`
	OriginName  string `json:"OriginName" xml:"OriginName"`
	OssOwner    string `json:"OssOwner" xml:"OssOwner"`
	PackageName string `json:"PackageName" xml:"PackageName"`
	CreateTime  int64  `json:"CreateTime" xml:"CreateTime"`
	OssBucket   string `json:"OssBucket" xml:"OssBucket"`
	OssPath     string `json:"OssPath" xml:"OssPath"`
	Tag         string `json:"Tag" xml:"Tag"`
	ProjectName string `json:"ProjectName" xml:"ProjectName"`
	Modifier    string `json:"Modifier" xml:"Modifier"`
	Md5         string `json:"Md5" xml:"Md5"`
	Creator     string `json:"Creator" xml:"Creator"`
	OssEndpoint string `json:"OssEndpoint" xml:"OssEndpoint"`
	Description string `json:"Description" xml:"Description"`
	Type        string `json:"Type" xml:"Type"`
}
