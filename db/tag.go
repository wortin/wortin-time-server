package db

import (
	"fmt"
	"wortin-time-server/bo"
	"wortin-time-server/po"
)

func QueryTags(pageNo, pageSize uint) (*bo.Tags, error) {
	tags := make([]bo.Tag, 0)
	r := DB.Raw("SELECT tags.id, tags.name, tags.emoji, (SELECT COUNT(action_tags.action_id) FROM action_tags WHERE action_tags.tag_id=tags.id) AS action_count FROM tags" +
		toLimitSql(pageNo, pageSize)).Scan(&tags)
	if r.Error != nil {
		return nil, r.Error
	}

	var total uint
	DB.Raw("SELECT COUNT(id) FROM tags").Scan(&total)

	return &bo.Tags{Tags: tags, Total: total}, nil
}

func InsertDefaultTags() {
	r := DB.Raw("DELETE FROM tags WHERE is_default=true").Scan(nil)
	if r.Error != nil {
		fmt.Println(r.Error.Error())
	}
	tags := []po.Tag{
		{Name: "今天", IsDefault: true},
		{Name: "工作", IsDefault: true},
		{Name: "生活", IsDefault: true},
		{Name: "家", IsDefault: true},
		{Name: "电脑", IsDefault: true},
		{Name: "手机", IsDefault: true},
		{Name: "随时随地", IsDefault: true},
	}
	r = DB.Create(tags)
	if r.Error != nil {
		fmt.Println(r.Error.Error())
	}
}
