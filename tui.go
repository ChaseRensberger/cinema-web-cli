package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/huh"
)

func runTUI() error {
	var action string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What would you like to do?").
				Options(
					huh.NewOption("Add Actor", "actor"),
					huh.NewOption("Add Director", "director"),
					huh.NewOption("Add Casting Director", "casting-director"),
					huh.NewOption("Add Project", "project"),
					huh.NewOption("Sync from S3", "sync"),
					huh.NewOption("Upload to S3", "upload"),
					huh.NewOption("Exit", "exit"),
				).
				Value(&action),
		),
	)

	if err := form.Run(); err != nil {
		return err
	}

	switch action {
	case "actor":
		return addActorForm()
	case "director":
		return addDirectorForm()
	case "casting-director":
		return addCastingDirectorForm()
	case "project":
		return addProjectForm()
	case "sync":
		return downloadFromS3(S3Bucket)
	case "upload":
		return uploadToS3(S3Bucket)
	case "exit":
		fmt.Println("Goodbye!")
		return nil
	}

	return nil
}

func addActorForm() error {
	var id, name string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Actor ID").
				Description("Leave empty to auto-generate from name").
				Value(&id),
			huh.NewInput().
				Title("Actor Name").
				Description("Full name of the actor").
				Value(&name).
				Validate(func(s string) error {
					return validateName(s)
				}),
		),
	)

	if err := form.Run(); err != nil {
		return err
	}

	if id == "" {
		id = normalizeID(name)
	} else {
		id = normalizeID(id)
	}

	if err := validateID(id); err != nil {
		return fmt.Errorf("invalid ID: %v", err)
	}

	data, err := loadData()
	if err != nil {
		return fmt.Errorf("error loading data: %v", err)
	}

	for _, actor := range data.Actors {
		if actor.ID == id {
			return fmt.Errorf("actor with ID '%s' already exists", id)
		}
	}

	data.Actors = append(data.Actors, Actor{ID: id, Name: name})

	if err := saveData(data); err != nil {
		return fmt.Errorf("error saving data: %v", err)
	}

	fmt.Printf("✅ Added actor: %s (%s)\n", name, id)
	return nil
}

func addDirectorForm() error {
	var id, name string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Director ID").
				Description("Leave empty to auto-generate from name").
				Value(&id),
			huh.NewInput().
				Title("Director Name").
				Description("Full name of the director").
				Value(&name).
				Validate(func(s string) error {
					return validateName(s)
				}),
		),
	)

	if err := form.Run(); err != nil {
		return err
	}

	if id == "" {
		id = normalizeID(name)
	} else {
		id = normalizeID(id)
	}

	if err := validateID(id); err != nil {
		return fmt.Errorf("invalid ID: %v", err)
	}

	data, err := loadData()
	if err != nil {
		return fmt.Errorf("error loading data: %v", err)
	}

	for _, director := range data.Directors {
		if director.ID == id {
			return fmt.Errorf("director with ID '%s' already exists", id)
		}
	}

	data.Directors = append(data.Directors, Director{ID: id, Name: name})

	if err := saveData(data); err != nil {
		return fmt.Errorf("error saving data: %v", err)
	}

	fmt.Printf("✅ Added director: %s (%s)\n", name, id)
	return nil
}

func addProjectForm() error {
	var id, title, yearStr, director, castingDirector, castStr string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Project ID").
				Description("Leave empty to auto-generate from title").
				Value(&id),
			huh.NewInput().
				Title("Project Title").
				Description("Full title of the project").
				Value(&title).
				Validate(func(s string) error {
					return validateName(s)
				}),
			huh.NewInput().
				Title("Year").
				Description("Release year (e.g., 2025)").
				Value(&yearStr).
				Validate(func(s string) error {
					_, err := strconv.Atoi(s)
					if err != nil {
						return fmt.Errorf("must be a valid year")
					}
					return nil
				}),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Director ID (optional)").
				Description("ID of the director").
				Value(&director),
			huh.NewInput().
				Title("Casting Director ID (optional)").
				Description("ID of the casting director").
				Value(&castingDirector),
			huh.NewInput().
				Title("Cast IDs (optional)").
				Description("Comma-separated list of actor IDs").
				Value(&castStr),
		),
	)

	if err := form.Run(); err != nil {
		return err
	}

	if id == "" {
		id = normalizeID(title)
	} else {
		id = normalizeID(id)
	}

	if err := validateID(id); err != nil {
		return fmt.Errorf("invalid ID: %v", err)
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return fmt.Errorf("invalid year: %v", err)
	}

	data, err := loadData()
	if err != nil {
		return fmt.Errorf("error loading data: %v", err)
	}

	for _, project := range data.Projects {
		if project.ID == id {
			return fmt.Errorf("project with ID '%s' already exists", id)
		}
	}

	project := Project{
		ID:    id,
		Title: title,
		Year:  year,
	}

	if director != "" {
		project.Director = strings.TrimSpace(director)
	}

	if castingDirector != "" {
		project.CastingDirector = strings.TrimSpace(castingDirector)
	}

	if castStr != "" {
		cast := strings.Split(castStr, ",")
		for i, c := range cast {
			cast[i] = strings.TrimSpace(c)
		}
		project.Cast = cast
	}

	data.Projects = append(data.Projects, project)

	if err := saveData(data); err != nil {
		return fmt.Errorf("error saving data: %v", err)
	}

	fmt.Printf("✅ Added project: %s (%s)\n", title, id)
	return nil
}

func addCastingDirectorForm() error {
	var id, name string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Casting Director ID").
				Description("Leave empty to auto-generate from name").
				Value(&id),
			huh.NewInput().
				Title("Casting Director Name").
				Description("Full name of the casting director").
				Value(&name).
				Validate(func(s string) error {
					return validateName(s)
				}),
		),
	)

	if err := form.Run(); err != nil {
		return err
	}

	if id == "" {
		id = normalizeID(name)
	} else {
		id = normalizeID(id)
	}

	if err := validateID(id); err != nil {
		return fmt.Errorf("invalid ID: %v", err)
	}

	data, err := loadData()
	if err != nil {
		return fmt.Errorf("error loading data: %v", err)
	}

	for _, castingDirector := range data.CastingDirectors {
		if castingDirector.ID == id {
			return fmt.Errorf("casting director with ID '%s' already exists", id)
		}
	}

	data.CastingDirectors = append(data.CastingDirectors, CastingDirector{ID: id, Name: name})

	if err := saveData(data); err != nil {
		return fmt.Errorf("error saving data: %v", err)
	}

	fmt.Printf("✅ Added casting director: %s (%s)\n", name, id)
	return nil
}
