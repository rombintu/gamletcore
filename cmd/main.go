package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"

	"github.com/rombintu/gamletcore/core"
	"github.com/rombintu/gamletcore/tools"
)

func main() {
	secret := flag.String("secret", "", "Secret for encode")
	action := flag.String("action", "", "Select action [encode|e; decode|d]")
	pathFile := flag.String("file", "./file.json", "Path to file")

	flag.Parse()
	logger := tools.NewLogger()

	switch *action {
	case "encode", "e":
		if *secret == "" {
			logger.Warn("--secret is nil. Exit with 0")
			os.Exit(0)
		}
		key, gams := core.Encode(*secret)
		data := make(map[string][]rune)
		data["KEY"] = key
		data["G1"] = gams[0]
		data["G2"] = gams[1]
		data["G3"] = gams[2]
		jsonData, err := json.MarshalIndent(&data, "", " ")
		if err != nil {
			logger.Error(err)
		}
		f, err := os.OpenFile(*pathFile, os.O_WRONLY, 0644)
		if err != nil {
			logger.Error(err)
		}
		defer f.Close()
		if _, err := f.Write(jsonData); err != nil {
			logger.Error(err)
		}
		logger.Infof("SUCCESS. Created file: %s", *pathFile)

	case "decode", "d":
		f, err := os.OpenFile(*pathFile, os.O_RDONLY, 0644)
		if err != nil {
			logger.Error(err)
		}
		defer f.Close()
		jsonData, err := ioutil.ReadAll(f)
		if err != nil {
			logger.Error(err)
		}
		data := make(map[string][]rune)
		if err := json.Unmarshal(jsonData, &data); err != nil {
			logger.Error(err)
		}
		secretRunes := core.Decode(data["KEY"], core.Gams{
			data["G1"], data["G2"], data["G3"],
		})
		logger.Infof("SECRET DECODED: %s", core.Convert(secretRunes))
	default:
		logger.Warn("--action is nil, select one [encode, decode]. Exit with 0")
		os.Exit(0)
	}
}
