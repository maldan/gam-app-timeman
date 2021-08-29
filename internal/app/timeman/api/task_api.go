package api

import (
	"github.com/maldan/gam-app-timeman/internal/app/timeman/core"
	"github.com/maldan/go-cmhp/cmhp_crypto"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_time"
	"github.com/maldan/go-restserver"
)

type TaskApi struct {
}

// Get by id
func (r TaskApi) GetIndex(args core.Task) core.Task {
	// Get file
	var item core.Task
	err := cmhp_file.ReadJSON(core.DataDir+"/task/"+args.Id+".json", &item)
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.NotFound, "id", "Task not found!")
	}

	return item
}

// Create
func (r TaskApi) PostIndex(args core.Task) {
	args.Id = cmhp_crypto.UID(8)
	cmhp_file.WriteJSON(core.DataDir+"/task/"+args.Id+".json", &args)

	// Add id to day list
	var idList = make([]string, 0)
	cmhp_file.ReadJSON(core.DataDir+"/day/"+cmhp_time.Format(args.Start, "YYYY-MM-DD")+".json", &idList)
	idList = append(idList, args.Id)
	cmhp_file.WriteJSON(core.DataDir+"/day/"+cmhp_time.Format(args.Start, "YYYY-MM-DD")+".json", &idList)
}

// Update
func (r TaskApi) PatchIndex(args core.Task) {
	cmhp_file.WriteJSON(core.DataDir+"/task/"+args.Id+".json", &args)
}

// Delete
func (r TaskApi) DeleteIndex(args core.Task) {
	// task := r.GetIndex(args)
	cmhp_file.Delete(core.DataDir + "/task/" + args.Id + ".json")
}
