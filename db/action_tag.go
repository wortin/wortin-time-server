package db

import "wortin-time-server/po"

func ActionAddTag(actionID uint, tagID uint) error {
	r := DB.Create(&po.ActionTag{
		ActionID: actionID,
		TagID:    tagID,
	})
	if r.Error != nil {
		return r.Error
	}
	return nil
}
