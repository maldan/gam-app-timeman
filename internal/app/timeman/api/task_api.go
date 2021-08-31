package api

import (
	"time"

	"github.com/maldan/gam-app-timeman/internal/app/timeman/core"
	"github.com/maldan/go-cmhp/cmhp_crypto"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_time"
	"github.com/maldan/go-restserver"
)

type TaskApi struct {
}

// Add new task
func (r TaskApi) PostIndex(args core.Task) {
	// Set new task id
	args.Id = cmhp_crypto.UID(8)

	// Save task
	cmhp_file.WriteJSON(core.DataDir+"/task/item/"+args.Id+".json", &args)

	// Save id list
	var idList = make([]string, 0)
	cmhp_file.ReadJSON(core.DataDir+"/task/stat/"+cmhp_time.Format(args.Start, "YYYY-MM-DD")+".json", &idList)
	idList = append(idList, args.Id)
	cmhp_file.WriteJSON(core.DataDir+"/task/stat/"+cmhp_time.Format(args.Start, "YYYY-MM-DD")+".json", &idList)
}

// Get by id
func (r TaskApi) GetByDay(args ArgsDate) []core.Task {
	var idList = make([]string, 0)
	cmhp_file.ReadJSON(core.DataDir+"/task/stat/"+cmhp_time.Format(args.Date, "YYYY-MM-DD")+".json", &idList)

	out := make([]core.Task, 0)
	for _, id := range idList {
		task := core.Task{}
		err := cmhp_file.ReadJSON(core.DataDir+"/task/item/"+id+".json", &task)
		if err == nil {
			out = append(out, task)
		}
	}
	return out
}

// Get by id
func (r TaskApi) GetTotalHourByDay(args ArgsDate) int64 {
	list := r.GetByDay(args)
	var out int64
	for _, task := range list {
		out += task.Stop.Unix() - task.Start.Unix()
	}
	return out
}

// Get year calory stat
func (f TaskApi) GetYearMap(args ArgsDate) map[string]interface{} {
	out := map[string]interface{}{}
	t1 := time.Date(args.Date.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)

	for i := 0; i < 366; i++ {
		t2 := t1.AddDate(0, 0, i)
		out[cmhp_time.Format(t2, "YYYY-MM-DD")] = f.GetTotalHourByDay(ArgsDate{Date: t2})
	}

	return out
}

// Get by id
func (r TaskApi) GetIndex(args core.Task) core.Task {
	// Get file
	var item core.Task
	err := cmhp_file.ReadJSON(core.DataDir+"/task/item/"+args.Id+".json", &item)
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.NotFound, "id", "Task not found!")
	}

	return item
}

// Delete
func (r TaskApi) DeleteIndex(args core.Task) {
	// Get task
	item := r.GetIndex(args)

	// Read stat
	var idList = make([]string, 0)
	cmhp_file.ReadJSON(core.DataDir+"/task/stat/"+cmhp_time.Format(item.Start, "YYYY-MM-DD")+".json", &idList)

	// Filter task list
	var outList = make([]string, 0)
	for _, id := range idList {
		if id == item.Id {
			continue
		}
		outList = append(outList, item.Id)
	}

	// Save stat
	cmhp_file.WriteJSON(core.DataDir+"/task/stat/"+cmhp_time.Format(item.Start, "YYYY-MM-DD")+".json", &outList)

	// Delete task file
	cmhp_file.Delete(core.DataDir + "/task/item/" + args.Id + ".json")
}

// Update
func (r TaskApi) PatchIndex(args core.Task) {
	r.DeleteIndex(args)

	// Save task
	cmhp_file.WriteJSON(core.DataDir+"/task/item/"+args.Id+".json", &args)

	// Save id list
	var idList = make([]string, 0)
	cmhp_file.ReadJSON(core.DataDir+"/task/stat/"+cmhp_time.Format(args.Start, "YYYY-MM-DD")+".json", &idList)
	idList = append(idList, args.Id)
	cmhp_file.WriteJSON(core.DataDir+"/task/stat/"+cmhp_time.Format(args.Start, "YYYY-MM-DD")+".json", &idList)
}
