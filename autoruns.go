package main

import (
	"os"
	"fmt"
	"bytes"
	"path/filepath"
	"encoding/json"
	"github.com/botherder/go-files"
	"github.com/botherder/go-autoruns"
	log "github.com/Sirupsen/logrus"
)

func generateAutoruns() {
	log.Info("Identifying files marked for persistence...")

	// Fetch autoruns.
	autoruns := autoruns.Autoruns()

	// Make backup of autoruns executables.
	for _, autorun := range(autoruns) {
		if _, err := os.Stat(autorun.ImagePath); err == nil {
			copyName := fmt.Sprintf("%s_%s.bin", autorun.MD5, autorun.ImageName)
			copyPath := filepath.Join(acq.Autoruns, copyName)
			files.Copy(autorun.ImagePath, copyPath)
		}
	}

	// Store the json list to file.
	autorunsJsonPath := filepath.Join(acq.Storage, "autoruns.json")
	autorunsJson, err := os.Create(autorunsJsonPath)
	if err != nil {
		log.Error("Unable to create autoruns list: ", err.Error())
		return
	}
	defer autorunsJson.Close()

	// Encoding into json.
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(autoruns)

	autorunsJson.WriteString(buf.String())
	autorunsJson.Sync()

	log.Info("Autoruns collected successfully!")
}