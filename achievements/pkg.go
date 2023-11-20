package achievements

import (
	"github.com/Muammadsoomro88/go-resume/utils"
	"github.com/johnfercher/maroto/pkg/pdf"
)

func AddAchievements(m pdf.Maroto, achievements []utils.Achievement) {
	m.Row(5, func() {
		m.Col(8, utils.AddHeading(m, "ACHIEVEMENTS"))
	})

	for _, achivement := range achievements {
		utils.AddNormalText(m, 4, 12, achivement.Title)
		utils.AddNormalText(m, 4, 12, achivement.Details)
	}
}
