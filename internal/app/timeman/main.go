package timeman

import (
	"embed"
	"flag"
	"fmt"

	"github.com/maldan/gam-app-timeman/internal/app/timeman/api"
	"github.com/maldan/gam-app-timeman/internal/app/timeman/core"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-rapi"
	"github.com/maldan/go-rapi/rapi_core"
	"github.com/maldan/go-rapi/rapi_rest"
	"github.com/maldan/go-rapi/rapi_vfs"
)

func init() {
	cmhp_file.ReadJSON(core.DataDir+"/config.json", &core.AppConfig)
}

func Start(frontFs embed.FS) {
	// Server
	var host = flag.String("host", "127.0.0.1", "Server Hostname")
	var port = flag.Int("port", 16000, "Server Port")
	_ = flag.Int("clientPort", 8080, "Client Port")

	// Data
	var dataDir = flag.String("dataDir", "db", "Data Directory")
	_ = flag.String("appId", "id", "App id")
	flag.Parse()

	// Set
	core.DataDir = *dataDir

	// Start server
	rapi.Start(rapi.Config{
		Host: fmt.Sprintf("%s:%d", *host, *port),
		Router: map[string]rapi_core.Handler{
			"/": rapi_vfs.VFSHandler{
				Root: "frontend/build",
				Fs:   frontFs,
			},
			"/api": rapi_rest.ApiHandler{
				Controller: map[string]interface{}{
					"task":     api.TaskApi{},
					"sleep":    api.SleepApi{},
					"activity": api.ActivityApi{},
				},
			},
			"/system": rapi_rest.ApiHandler{
				Controller: map[string]interface{}{
					"config": api.ConfigApi{},
				},
			},
		},
		DbPath: core.DataDir,
	})
}
