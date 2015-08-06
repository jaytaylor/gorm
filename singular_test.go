package gorm_test

import (
	"testing"
)

type Actor struct {
	Id     int64
	Name   string
	Movies []Movie `gorm:"many2many:actor_movie;"`
}

type Movie struct {
	Id     int
	Name   string
	Actors []Actor `gorm:"many2many:actor_movie;"`
}

func TestM2mSingularMode(t *testing.T) {
	defer func() {
		// Always restore original singular-ness mode value.
		DB.SingularTable(false)
	}()

	// Retrieves the unquoted-
	tableName := func(model interface{}) string {
		name := DB.NewScope(model).TableName()
		return name
	}

	for _, table := range []interface{}{&Actor{}, &Movie{}} {
		if err := DB.DropTableIfExists(table).Error; err != nil {
			t.Fatal(err)
		}
		if err := DB.AutoMigrate(table).Error; err != nil {
			t.Fatal(err)
		}
	}

	if actual, expected := tableName(&Actor{}), "actors"; actual != expected {
		t.Fatalf("Expected table name=%v when SingularTable=false but actual=%v", expected, actual)
	}
	if actual, expected := tableName(&Movie{}), "movies"; actual != expected {
		t.Fatalf("Expected table name=%v when SingularTable=false but actual=%v", expected, actual)
	}

	DB.SingularTable(true)

	if actual, expected := tableName(&Actor{}), "actor"; actual != expected {
		t.Fatalf("Expected table name=%v when SingularTable=true but actual=%v", expected, actual)
	}
	if actual, expected := tableName(&Movie{}), "movie"; actual != expected {
		t.Fatalf("Expected table name=%v when SingularTable=true but actual=%v", expected, actual)
	}

	// Verify m2m behavior in singular-mode.
	/*
		jackman := &Actor{
			Name: "Hugh Jackman",
		}
		picard := &Actor{
			Name: "Patrick Stewart",
		}
		actors := []*Actor{
			jackman,
			picard,
		}
		for _, actor := range actors {
			if err := DB.Save(actor).Error; err != nil {
				t.Fatalf("Error saving actor=%+v: %s", *actor, err)
			}
		}

		xMen := &Movie{
			Name: "X-Men",
		}
		chappie := &Movie{
			Name: "Chappie",
		}
		prestige := &Movie{
			Name: "The Prestige",
		}
		starTrekGenerations := &Movie{
			Name: "Star Trek: Generations",
		}
		movies := []*Movie{
			xMen,
			chappie,
			prestige,
			starTrekGenerations,
		}
		for _, movie := range movies {
			if err := DB.Save(movie).Error; err != nil {
				t.Fatalf("Error saving movie=%+v: %s", *movie, err)
			}
		}

		{
			model := DB.Model(jackman)
			if err := model.Error; err != nil {
				t.Fatal(err)
			}
			assoc := model.Association("Movies")
			if err := assoc.Error; err != nil {
				t.Fatal(err)
			}
			if err := assoc.Append(xMen, chappie, prestige).Error; err != nil {
				t.Fatal(err)
			}
		}

		{
			model := DB.Model(picard)
			if err := model.Error; err != nil {
				t.Fatal(err)
			}
			assoc := model.Association("Movies")
			if err := assoc.Error; err != nil {
				t.Fatal(err)
			}
			if err := assoc.Append(xMen, starTrekGenerations).Error; err != nil {
				t.Fatal(err)
			}
		}

		{
			model := DB.Model(jackman)
			if err := model.Error; err != nil {
				t.Fatal(err)
			}
			found := []Movie{}
			if err := model.Related(&found, "Movies").Error; err != nil {
				t.Fatal(err)
			}
			if count := len(found); count != 3 {
				t.Fatal("Expected 3 movies for Hugh Jackman but instead count=%v", count)
			}
		}

		{
			model := DB.Model(picard)
			if err := model.Error; err != nil {
				t.Fatal(err)
			}
			found := []Movie{}
			if err := model.Related(&found, "Movies").Error; err != nil {
				t.Fatal(err)
			}
			if count := len(found); count != 2 {
				t.Fatal("Expected 2 movies for Patrick Stewart but instead count=%v", count)
			}
		}*/
}
