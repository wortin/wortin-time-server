package db

import (
	"wortin-time-server/bo"
	"wortin-time-server/po"
)

func CreateProject(project *po.Project) error {
	r := DB.Create(project)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func QueryProjects(pageNo, pageSize uint) (*bo.QueryProjectsRes, error) {
	var ps []bo.QPRItemProject
	r := DB.Raw("SELECT projects.id, projects.name, projects.target_id, projects.planed_start_date, projects.planed_end_date, projects.desc, projects.state, " +
		"(SELECT COUNT(actions.id) FROM actions WHERE actions.project_id=projects.id AND actions.state<>1) AS action_count, " +
		"(SELECT COUNT(actions.id) FROM actions WHERE actions.project_id=projects.id AND actions.state=2) AS completed_action_count," +
		"(SELECT name FROM targets WHERE targets.id=projects.target_id) AS target_name" +
		" FROM projects WHERE projects.state<>1" + toLimitSql(pageNo, pageSize)).Scan(&ps)
	if r.Error != nil {
		return nil, r.Error
	}
	var total uint
	r = DB.Raw("SELECT COUNT(id) FROM projects WHERE projects.state<>1").Scan(&total)
	return &bo.QueryProjectsRes{Projects: ps, Total: total}, nil
}
