package core

import (
	"io"
	"io/fs"
	"log"

	"encoding/csv"
)

type Room struct {
	ID          string
	Name        string
	Description string
	North       *string
	East        *string
	South       *string
	West        *string
}

type MapGrid [][]*Room

func LoadData(dataFS fs.FS) MapGrid {
	grid := make([][]*Room, 0)

	f, err := dataFS.Open("data/map.csv")
	if err != nil {
		log.Fatal("cannot open map file")
	}
	r := csv.NewReader(f)
	// bom := string('\uFEFF')
	for {
		line, err := r.Read()
		if err == io.EOF {
			f.Close()
			break
		} else if err != nil {
			log.Fatal(err)
		}
		row := make([]*Room, 0)
		for _, s := range line {
			if s == "1" {
				room := Room{}
				row = append(row, &room)
			} else {
				row = append(row, nil)
			}
		}
		grid = append(grid, row)
	}
	return grid
}
