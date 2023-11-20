package utils

import (
	"log"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

type Resume struct {
	Name         string        `json:"name"`
	Position     string        `json:"position"`
	Github       string        `json:"github"`
	Email        string        `json:"email"`
	LinkedIn     string        `json:"linkedin"`
	Phone        string        `json:"phone"`
	Location     string        `json:"location"`
	Summary      string        `json:"summary"`
	Job          []Experience  `json:"job"`
	Study        []Education   `json:"study"`
	PersonalProj []Project     `json:"personalproj"`
	Skills       string        `json:"skills"`
	Achievements []Achievement `json:"achievements"`
}

type Experience struct {
	Position string `json:"position"`
	Company  string `json:"company"`
	Start    string `json:"start"`
	End      string `json:"end"`
	Summary  string `json:"summary"`
}

type Education struct {
	Program   string `json:"program"`
	Start     string `json:"start"`
	End       string `json:"end"`
	Institute string `json:"institution"`
	Summary   string `json:"summary"`
}

type Project struct {
	Title   string `json:"title"`
	Details string `json:"details"`
}

type Achievement struct {
	Title   string `json:"title"`
	Details string `json:"details"`
}

func AddProjectTitle(m pdf.Maroto, title string) {
	m.Row(4, func() {
		m.Col(6, func() {
			m.Text(title, props.Text{
				Size:            9.0,
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
}

func AddHeading(m pdf.Maroto, heading string) func() {
	return func() {
		m.Text(heading, props.Text{
			Size:            11,
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
	}
}

func AddPosition(m pdf.Maroto, post string) func() {
	return func() {
		m.Text(post, props.Text{
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
	}
}

func AddCalenderImg(m pdf.Maroto) {
	err := m.FileImage("assets/calender_icon.jpg", props.Rect{
		Left:    2.0,
		Percent: 60,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func AddDate(m pdf.Maroto, from, to string) {
	m.Text(from+" - "+to, props.Text{
		Size:   7,
		Style:  consts.Normal,
		Family: consts.Arial,
		Left:   5.0,
		Color: color.Color{
			Red:   10,
			Green: 20,
			Blue:  30,
		},
	})
}

func AddCompany(m pdf.Maroto, company string) func() {
	return func() {
		m.Text(company, props.Text{
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
	}
}

func AddNormalText(m pdf.Maroto, rowLength float64, colLength int, text string) {
	m.Row(rowLength, func() {
		m.Col(uint(colLength), func() {
			m.Text(text, props.Text{
				Size:            8,
				Style:           consts.Normal,
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
}
