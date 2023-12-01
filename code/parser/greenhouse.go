package parser

import (
	"fmt"
	"leet/fx"
	"strings"
	"time"
)

var TagsList = []string{
	"google",
	"ms",
	"At",
	"vero",
	"eos",
	"et",
	"accusamus",
	"iusto",
	"odio",
	"dignissimos",
	"ducimus",
	"qui",
	"blanditiis",
	"praesentium",
	"voluptatum",
	"deleniti",
	"atque",
	"corrupti",
	"quos",
	"dolores",
	"quas",
	"molestias",
	"excepturi",
	"sint",
	"occaecati",
	"cupiditate",
	"non",
	"provident",
	"similique",
	"sunt",
	"in",
	"culpa",
	"officia",
	"deserunt",
	"mollitia",
	"animi",
	"id",
	"est",
	"laborum",
	"dolorum",
	"fuga",
	"harum",
	"quidem",
	"rerum",
	"facilis",
	"expedita",
	"distinctio",
	"Nam",
	"libero",
	"tempore",
	"cum",
	"soluta",
	"nobis",
	"eligendi",
	"optio",
	"cumque",
	"nihil",
	"impedit",
	"quo",
	"minus",
	"quod",
	"maxime",
	"placeat",
	"facere",
	"possimus",
	"omnis",
	"voluptas",
	"assumenda",
	"dolor",
	"repellendus",
	"Temporibus",
	"autem",
	"quibusdam",
	"aut",
	"officiis",
	"debitis",
	"necessitatibus",
	"saepe",
	"eveniet",
	"ut",
	"voluptates",
	"repudiandae",
	"molestiae",
	"recusandae",
	"Itaque",
	"earum",
	"hic",
	"tenetur",
	"a",
	"sapiente",
	"delectus",
	"reiciendis",
	"voluptatibus",
	"maiores",
	"alias",
	"consequatur",
	"perferendis",
	"doloribus",
	"asperiores",
	"repellat",
}

func GreenHouse() {
	url := "https://boards-api.greenhouse.io/v1/boards/tubitv/jobs?content=true"
	jobs := fx.ParseGreenhouse(url)

	startTime := time.Now()

	// Convert TagsList to a map for constant time lookup
	tagsMap := make(map[string]bool)
	for _, tag := range TagsList {
		tagsMap[tag] = true
	}

	count := 0

	for _, i := range jobs.Jobs {
		words := strings.Fields(i.Content)
		for _, word := range words {
			if _, exists := tagsMap[word]; exists {
				count++
			}
		}
	}

	fmt.Println(time.Since(startTime))
	fmt.Println(count)
}
