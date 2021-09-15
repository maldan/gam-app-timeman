package timeman

import (
	"embed"
	"flag"
	"fmt"

	"github.com/maldan/gam-app-timeman/internal/app/timeman/api"
	"github.com/maldan/gam-app-timeman/internal/app/timeman/core"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-restserver"
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

	/*m := make([]core.Sleep, 0)
	cmhp_file.ReadJSON("./sleep.json", &m)
	fff := api.SleepApi{}
	for _, xx := range m {
		loc, _ := time.LoadLocation("Europe/Moscow")
		xx.Start = xx.Start.In(loc)
		xx.Stop = xx.Stop.In(loc)
		fff.PostIndex(xx)
	}
	fmt.Println(2)*/

	// Init server
	restserver.Start(fmt.Sprintf("%s:%d", *host, *port), map[string]interface{}{
		"/": restserver.VirtualFs{Root: "frontend/build/", Fs: frontFs},
		"/api": map[string]interface{}{
			"task":  api.TaskApi{},
			"sleep": api.SleepApi{},
		},
		"/system": map[string]interface{}{
			"config": api.ConfigApi{},
		},
	})
}
