package code

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

//IsValid ..
func (all *All) IsValid(arg []string) error {
	err := errors.New("Invalid arg")
	if len(arg) == 0 || len(arg) > 1 {
		return err
	}

	data, err := ioutil.ReadFile(arg[0])
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return errors.New("Invalid map")
	}
	all.Map = string(data)
	if err = all.isValidMap(string(data)); err != nil {
		return err
	}
	return nil
}

func (all *All) isValidMap(data string) error {
	check := strings.Split(data, "\n")
	var start, end bool
	var coords [][]int
	for i, v := range check {
		if i == 0 {
			x, err := strconv.Atoi(v)
			if err != nil || x <= 0 || x > 100000 {
				return errors.New("Invalid number of ants")
			}
			all.Ants = x
		}

		if len(v) == 0 {
			continue
		}

		rooms := strings.Split(v, " ")

		links := strings.Split(v, "-")

		if len(rooms) == 3 {
			var cord []int
			if strings.HasPrefix(v, "#") {
				continue
			}

			if strings.HasPrefix(rooms[0], "L") {
				return errors.New("Invalid room name")
			}

			x, errx := strconv.Atoi(rooms[1])

			y, erry := strconv.Atoi(rooms[2])

			if errx != nil || erry != nil {
				return errors.New("Invalid coords")
			}
			cord = append(cord, x, y)
			all.Rooms = append(all.Rooms, rooms[0])
			coords = append(coords, cord)
		}

		if len(links) == 2 {
			all.Links = append(all.Links, v)
		}

		if v == "##start" && i != len(check)-1 && len(check[i+1]) != 0 {
			room := strings.Split(check[i+1], " ")
			if len(room) != 3 {
				continue
			}
			all.StartRoom = room[0]
			start = true
		} else if v == "##end" && i != len(check)-1 && len(check[i+1]) != 0 {
			room := strings.Split(check[i+1], " ")
			if len(room) != 3 {
				continue
			}
			all.EndRoom = room[0]
			end = true
		}
	}
	if !start {
		return errors.New("No start room")
	}
	if !end {
		return errors.New("No end room")
	}

	if !checkCoord(coords) {
		return errors.New("Invalid coords")
	}
	return nil
}

func checkCoord(ch [][]int) bool {
	for i := 0; i < len(ch); i++ {
		for j := i + 1; j < len(ch); j++ {
			if !checker(ch[i], ch[j]) {
				return false
			}
		}
	}
	return true
}

func checker(x, y []int) bool {
	if x[0] == y[0] && x[1] == y[1] {
		return false
	}
	return true
}
