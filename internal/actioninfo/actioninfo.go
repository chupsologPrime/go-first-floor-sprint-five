package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	ActionInfo() (string, error)
	Parse(datastring string) (err error)
}

func Info(dataset []string, dp DataParser) {
	for _, el := range dataset {
		err := dp.Parse(el)
		if err != nil {
			log.Printf("string parsing error\n")
			continue
		}
		str, err := dp.ActionInfo()
		if err != nil {
			log.Printf("string formation error")
		}
		fmt.Println(str)
	}
}
