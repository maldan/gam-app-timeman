package api

import (
	"github.com/maldan/gam-app-timeman/internal/app/timeman/core"
	"github.com/maldan/go-cmhp/cmhp_file"
)

type ActivityApi struct {
}

// Get config
func (r ActivityApi) GetList() []core.Activity {
	config := core.Config{}
	cmhp_file.ReadJSON(core.DataDir+"/config.json", &config)

	out := make([]core.Activity, 0)
	for activity, color := range config.ActivityColor {
		out = append(out, core.Activity{
			Name:  activity,
			Color: color,
		})
	}

	return out
}
