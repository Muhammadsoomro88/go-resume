package skill

import (
	"strings"

	"github.com/Muammadsoomro88/go-resume/utils"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func AddSkills(m pdf.Maroto, skills string) {
	m.Row(5, func() {
		m.Col(8, utils.AddHeading(m, "SKILLS"))
	})

	skillsArr := strings.Split(skills, ",")
	row := len(skillsArr) / 11
	if len(skillsArr)%11 > 0 {
		row += 1
	}
	lower := 0
	upper := 0
	skillColumn := 1
	for i := 1; i <= row; i++ {
		if i*11-len(skillsArr) < 0 {
			upper = i * 11
		} else {
			upper = len(skillsArr)
		}
		//
		m.Row(5, func() {
			for _, skill := range skillsArr[lower:upper] {
				if len(skill) >= 10 {
					skillColumn = 2
				} else {
					skillColumn = 1
				}
				m.Col(uint(skillColumn), func() {
					m.Text("•"+skill, props.Text{
						Size:            8,
						Style:           consts.Bold,
						Family:          consts.Arial,
						Align:           consts.Left,
						Extrapolate:     false,
						VerticalPadding: 1.0,
						Color: color.Color{
							Red:   10,
							Green: 20,
							Blue:  30,
						},
					})
				})
			}
		})
		lower = upper
	}
}
