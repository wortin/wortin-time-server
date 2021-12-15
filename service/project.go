package service

import (
	"wortin-time-server/bo"
	"wortin-time-server/db"
	"wortin-time-server/po"
)

func CreateProject(projectName string) error {
	return db.CreateProject(&po.Project{Name: projectName})
}

func QueryProjects(pageNo, pageSize uint) (*bo.QueryProjectsRes, error) {
	return db.QueryProjects(pageNo, pageSize)
}
