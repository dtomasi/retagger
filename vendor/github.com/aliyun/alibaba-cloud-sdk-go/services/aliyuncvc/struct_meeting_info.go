package aliyuncvc

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

// MeetingInfo is a nested struct in aliyuncvc response
type MeetingInfo struct {
	MeetingToken  string           `json:"MeetingToken" xml:"MeetingToken"`
	MeetingUUID   string           `json:"MeetingUUID" xml:"MeetingUUID"`
	CreateTime    int64            `json:"CreateTime" xml:"CreateTime"`
	MeetingAppId  string           `json:"MeetingAppId" xml:"MeetingAppId"`
	MeetingCode   string           `json:"MeetingCode" xml:"MeetingCode"`
	MemberUUID    string           `json:"MemberUUID" xml:"MemberUUID"`
	ValidTime     int64            `json:"ValidTime" xml:"ValidTime"`
	ClientAppId   string           `json:"ClientAppId" xml:"ClientAppId"`
	UserId        string           `json:"UserId" xml:"UserId"`
	MeetingName   string           `json:"MeetingName" xml:"MeetingName"`
	UserName      string           `json:"UserName" xml:"UserName"`
	MeetingDomain string           `json:"MeetingDomain" xml:"MeetingDomain"`
	Memo          string           `json:"Memo" xml:"Memo"`
	SlsInfo       SlsInfo          `json:"SlsInfo" xml:"SlsInfo"`
	MemberList    []MemberListItem `json:"MemberList" xml:"MemberList"`
	MemberInfos   []MemberRecord   `json:"MemberInfos" xml:"MemberInfos"`
}
