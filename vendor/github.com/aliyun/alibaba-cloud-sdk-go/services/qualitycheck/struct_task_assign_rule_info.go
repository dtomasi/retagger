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

// TaskAssignRuleInfo is a nested struct in qualitycheck response
type TaskAssignRuleInfo struct {
	RuleId         int64                      `json:"RuleId" xml:"RuleId"`
	Enabled        int                        `json:"Enabled" xml:"Enabled"`
	CallType       int                        `json:"CallType" xml:"CallType"`
	Priority       int                        `json:"Priority" xml:"Priority"`
	DurationMin    int                        `json:"DurationMin" xml:"DurationMin"`
	DurationMax    int                        `json:"DurationMax" xml:"DurationMax"`
	CreateTime     string                     `json:"CreateTime" xml:"CreateTime"`
	UpdateTime     string                     `json:"UpdateTime" xml:"UpdateTime"`
	AgentsStr      string                     `json:"AgentsStr" xml:"AgentsStr"`
	SkillGroupsStr string                     `json:"SkillGroupsStr" xml:"SkillGroupsStr"`
	Agents         Agents                     `json:"Agents" xml:"Agents"`
	SkillGroups    SkillGroups                `json:"SkillGroups" xml:"SkillGroups"`
	Reviewers      Reviewers                  `json:"Reviewers" xml:"Reviewers"`
	Rules          RulesInListTaskAssignRules `json:"Rules" xml:"Rules"`
}
