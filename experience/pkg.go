package experience

import (
	"strings"

	"github.com/Muammadsoomro88/go-resume/utils"
	"github.com/johnfercher/maroto/pkg/pdf"
)

func AddExperience(m pdf.Maroto, jobs []utils.Experience) {
	m.Row(6, func() {
		m.Col(8, utils.AddHeading(m, "EXPERIENCE"))
	})

	for _, job := range jobs {
		m.Row(5, func() {
			m.Col(10, utils.AddPosition(m, job.Position))
			m.Col(2, func() {
				utils.AddCalenderImg(m)
				utils.AddDate(m, job.Start, job.End)
			})
		})

		m.Row(5, func() {
			m.Col(10, utils.AddCompany(m, job.Company))
		})

		for i, x := range strings.Split(job.Summary, "\n") {
			if i == 0 {
				utils.AddNormalText(m, float64(len(x)/23), 12, x)
			} else {
				utils.AddNormalText(m, 4, 12, x)
			}
		}
	}
}
