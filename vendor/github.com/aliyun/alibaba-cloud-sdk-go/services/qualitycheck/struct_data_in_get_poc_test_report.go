package qualitycheck

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

// DataInGetPocTestReport is a nested struct in qualitycheck response
type DataInGetPocTestReport struct {
	AsrAccuracyRate string `json:"AsrAccuracyRate" xml:"AsrAccuracyRate"`
	LabelNum        int    `json:"LabelNum" xml:"LabelNum"`
	Poc             bool   `json:"Poc" xml:"Poc"`
	AudioBit        string `json:"AudioBit" xml:"AudioBit"`
	AudioNum        int    `json:"AudioNum" xml:"AudioNum"`
	AudioSampleRate string `json:"AudioSampleRate" xml:"AudioSampleRate"`
	AudioTrack      string `json:"AudioTrack" xml:"AudioTrack"`
	CharacterNum    int    `json:"CharacterNum" xml:"CharacterNum"`
	CheckCost       string `json:"CheckCost" xml:"CheckCost"`
	ModelName       string `json:"ModelName" xml:"ModelName"`
	PocTime         string `json:"PocTime" xml:"PocTime"`
	RuleInfo        string `json:"RuleInfo" xml:"RuleInfo"`
	RuleNum         int    `json:"RuleNum" xml:"RuleNum"`
	RuleRealRate    string `json:"RuleRealRate" xml:"RuleRealRate"`
}
