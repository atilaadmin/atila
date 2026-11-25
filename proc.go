package main

import (


	"os"
	"path/filepath"
	"strconv"
	"strings"
)




func proc() string {

my_proc :=  "writebackk"





	entries, err := os.ReadDir("/proc")
	if err != nil {
		return ""
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		pidStr := entry.Name()

		if _, err := strconv.Atoi(pidStr); err != nil {
			continue
		}

		basePath := filepath.Join("/proc", pidStr)

		// Verifica o nome do processo.
		commBytes, err := os.ReadFile(
			filepath.Join(basePath, "comm"),
		)
		if err != nil {
			continue
		}

		name := strings.TrimSpace(string(commBytes))
		if name !=   my_proc {
			continue
		}

		// Verifica o estado do processo.
		statBytes, err := os.ReadFile(
			filepath.Join(basePath, "stat"),
		)
		if err != nil {
			continue
		}

		stat := string(statBytes)
		closeParen := strings.LastIndex(stat, ")")

		if closeParen == -1 || closeParen+2 >= len(stat) {
			continue
		}

		state := stat[closeParen+2]

		switch state {
		case 'R', 'S', 'D':
			return "ok"
		default:
			continue
		}
	}


	return "not ok"



}






