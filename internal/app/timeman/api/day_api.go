package api

import (
	"github.com/maldan/gam-app-timeman/internal/app/timeman/core"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_time"
	"github.com/maldan/go-restserver"
)

type DayApi struct {
}

// Get by id
func (r DayApi) GetIndex(args ArgsDate) []core.Task {
	// Get file
	var idList = make([]string, 0)
	err := cmhp_file.ReadJSON(core.DataDir+"/day/"+cmhp_time.Format(args.Date, "YYYY-MM-DD")+".json", &idList)
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.NotFound, "id", "Day not found!")
	}

	var out = make([]core.Task, 0)
	for _, id := range idList {
		t := core.Task{}
		err := cmhp_file.ReadJSON(core.DataDir+"/task/"+id+".json", &t)
		if err == nil {
			out = append(out, t)
		}
	}

	return out
}
