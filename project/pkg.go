package project

import (
	"github.com/Muammadsoomro88/go-resume/utils"
	"github.com/johnfercher/maroto/pkg/pdf"
)

func AddPersonalProject(m pdf.Maroto, projects []utils.Project) {
	m.Row(6, func() {
		m.Col(8, utils.AddHeading(m, "PERSONAL PROJECTS"))
	})

	for _, project := range projects {
		utils.AddProjectTitle(m, project.Title)
		utils.AddNormalText(m, float64(len(project.Details)/26), 12, project.Details)
		m.Row(1, func() {})
	}
}
