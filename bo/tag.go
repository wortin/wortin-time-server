package bo

type Tag struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Emoji       string `json:"emoji"`
	ActionCount uint   `json:"actionCount"` // 该标签具有的动作数
}

type Tags struct {
	Tags  []Tag `json:"tags"`
	Total uint  `json:"total"`
}
