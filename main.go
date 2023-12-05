package main

import (
	"log"
	"os"

	"github.com/Muammadsoomro88/go-resume/achievements"
	"github.com/Muammadsoomro88/go-resume/education"
	"github.com/Muammadsoomro88/go-resume/experience"
	"github.com/Muammadsoomro88/go-resume/project"
	"github.com/Muammadsoomro88/go-resume/skill"
	"github.com/Muammadsoomro88/go-resume/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func createResume(ctx *fiber.Ctx) error {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 5, 10)

	body := new(utils.Resume)
	err := ctx.BodyParser(body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(err)
		return err
	}

	buildResume(m, body)
	err = m.OutputFileAndClose("pdfs/resume.pdf")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return ctx.SendString("pdf saved successfully.")
}

func buildResume(m pdf.Maroto, info *utils.Resume) {
	m.Row(10, func() {
		m.Col(8, func() {
			m.Text(info.Name, props.Text{
				Size:            23,
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
	})

	m.Row(8, func() {
		m.Col(10, func() {
			m.Text(info.Position, props.Text{
				Size:            12,
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

	// adding social links
	m.Row(5, func() {
		m.Col(8, addLink(m, "assets/GitHub_logo.png", info.Github))
		m.Col(1, addLink(m, "assets/mail_icon.jpg", info.Email))
	})

	m.Row(5, func() {
		m.Col(8, addLink(m, "assets/linkedin-icon-logo-png-transparent.png", info.LinkedIn))
		m.Col(1, addLink(m, "assets/phone-icon-png-3.png", info.Phone))
	})

	m.Row(5, func() {
		m.Col(8, addLink(m, "assets/location_icon.png", info.Location))
	})

	m.Line(5.0, props.Line{
		Color: color.Color{
			Red:   10,
			Green: 20,
			Blue:  30,
		},
		Style: consts.Solid,
		Width: 1.0,
	})

	// Summary
	m.Row(5, func() {
		m.Col(8, utils.AddHeading(m, "SUMMARY"))
	})

	summaryRow := float64(len(info.Summary) / 25)
	m.Row(summaryRow, func() {
		m.Col(12, func() {
			m.Text(info.Summary, props.Text{
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

	// experience
	experience.AddExperience(m, info.Job)

	// Skill
	skill.AddSkills(m, info.Skills)

	// Personal projects
	project.AddPersonalProject(m, info.PersonalProj)

	// education
	education.AddEducation(m, info.Study)

	// Achievements
	achievements.AddAchievements(m, info.Achievements)
}

func addLink(m pdf.Maroto, imgDir string, value string) func() {
	return func() {
		err := m.FileImage(imgDir, props.Rect{
			Percent: 80,
		})
		if err != nil {
			log.Fatal(err)
		}
		m.Text(value, props.Text{
			Size:   8,
			Style:  consts.Bold,
			Family: consts.Arial,
			Left:   5.0,
			Color: color.Color{
				Red:   10,
				Green: 20,
				Blue:  30,
			},
		})
	}
}

func main() {
	app := fiber.New()

	// some necessary imports
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to fiber")
	})
	app.Post("/resume", createResume)

	log.Fatal(app.Listen(":8081"))
}
