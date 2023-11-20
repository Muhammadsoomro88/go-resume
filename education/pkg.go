package education

import (
	"strings"

	"github.com/Muammadsoomro88/go-resume/utils"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func AddEducation(m pdf.Maroto, education []utils.Education) {
	m.Row(6, func() {
		m.Col(8, utils.AddHeading(m, "EDUCATION"))
	})

	for _, edu := range education {
		m.Row(5, func() {
			m.Col(10, func() {
				m.Text(edu.Program, props.Text{
					Size:            10,
					Style:           consts.Bold,
					Family:          consts.Arial,
					Extrapolate:     false,
					VerticalPadding: 1.0,
					Color: color.Color{
						Red:   10,
						Green: 20,
						Blue:  30,
					},
				})
			})
		})

		m.Col(3, func() {
			utils.AddCalenderImg(m)
			utils.AddDate(m, edu.Start, edu.End)
		})

		m.Row(5, func() {
			m.Col(10, func() {
				m.Text(edu.Institute, props.Text{
					Size:            9,
					Style:           consts.Bold,
					Family:          consts.Arial,
					Extrapolate:     false,
					VerticalPadding: 1.0,
					Color: color.Color{
						Red:   90,
						Green: 130,
						Blue:  230,
					},
				})
			})
		})

		for _, x := range strings.Split(edu.Summary, "\n") {
			utils.AddNormalText(m, 4, 10, x)
		}
	}
}
