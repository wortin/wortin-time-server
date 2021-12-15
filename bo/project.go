package bo

type QueryProjectsRes struct {
	Total    uint             `json:"total"`
	Projects []QPRItemProject `json:"projects"`
}

type QPRItemProject struct {
	ID                   uint   `json:"id"`
	Name                 string `json:"name"`
	TargetID             uint   `json:"targetID"`
	PlanedStartDate      string `json:"planedStartDate"`
	PlanedEndDate        string `json:"planedEndDate"`
	Desc                 string `json:"desc"`
	State                int    `json:"state"` // 1 删除在回收站; 2 已完成;
	ActionCount          uint   `json:"actionCount"`
	CompletedActionCount uint   `json:"completedActionCount"`
	TargetName           string `json:"targetName"`
}
