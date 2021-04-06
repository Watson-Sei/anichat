package plugin

import (
	"fmt"
	"github.com/mattn/go-pipeline"
	"log"
	"strconv"
	"strings"
	"time"
)

func Reformat(timeDate string, id uint) error {
	log.Println("Reformat: ", timeDate, id)
	// refrence: https://cloud6.net/so/go/187380
	for i, v := range strings.Split(strings.Join(strings.Fields(timeDate), ""), "～") {
		switch i {
		case 0:
			out, err := pipeline.Output(
				[]string{"echo", fmt.Sprintf("cd /go/src/spider && go run main.go -mode 2 -p 1 -id %d", id)},
				[]string{"at", ShapingTime(v)},
			)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(string(out))
			log.Println("start at command register")
		case 1:
			out, err := pipeline.Output(
				[]string{"echo", fmt.Sprintf("cd /go/src/spider && go run main.go -mode 2 -p 2 -id %d", id)},
				[]string{"at", ShapingTime(v)},
			)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(string(out))
			log.Println("finish at command register")
		}
	}
	return nil
}

func ShapingTime(fullTime string) string {
	const layout = "01/02/06"
	day := time.Now()
	vSlice := strings.Split(fullTime, ":")
	vhInt, _ := strconv.Atoi(vSlice[0])
	if vhInt >= 24 {
		log.Println("24時過ぎています")
		return fmt.Sprintf("%d:%s %s", vhInt-24, vSlice[1], day.AddDate(0, 0, 1).Format(layout))
	}
	log.Println("24時前です:")
	return fmt.Sprintf("%s:%s %s", vSlice[0], vSlice[1], day.Format(layout))
}