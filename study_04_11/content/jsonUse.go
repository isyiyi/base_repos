package content

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string	`json:"title,omitempty"`
	Year   int	`json:"year"`
	Color  bool
	Actors []string
}

func StructConJson() []byte{
	var movie = []Movie{0: {Title: "战狼",
		Year:   2018,
		Color:  false,
		Actors: []string{"吴京", "张翰"},
	}, 1: {
		Title:  "大话西游",
		Year:   2022,
		Color:  false,
		Actors: []string{"章子怡", "吴青峰"},
	},
	}
	// data, err := json.Marshal(movie)
	data, err := json.MarshalIndent(movie, "", "  ")
	if err != nil {
		fmt.Println("转换失败")
	}
	fmt.Printf("%s", data)
	return data
}

type Phone struct {
	Brand string
}

func JsonConvStruct(res []byte) {
	var movie []Movie
	err := json.Unmarshal([]byte(res), &movie)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(movie)
}