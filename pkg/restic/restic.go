package restic

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
	"sync"
)

var Snapmap = make(map[string]string)

func CreateConcurrentSnapshots(dirs []string) {
	var cswg sync.WaitGroup
	fmt.Println("These are the DIRS: ", dirs)
	cswg.Add(len(dirs))
	for _, dir := range dirs {
		go createSnapshot(dir, &cswg)
	}
	cswg.Wait()
}

func RestoreConcurrentSnapshots(snapshotIDs []string) {
	var rswg sync.WaitGroup
	fmt.Println("Restoring Snapshots, Snapshot ID's: ", snapshotIDs)
	rswg.Add(len(snapshotIDs))
	for _, dir := range snapshotIDs {
		go restoreSnapshot(dir, &rswg)
	}
	rswg.Wait()
}

func ListSnapshots() {
	cmd := exec.Command("restic", "-r", "s3:s3.amazonaws.com/resticbucket", "snapshots", "--json")
	bs, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(bs), err)
	}
	fmt.Println(string(bs))
}

func createSnapshot(dir string, cswg *sync.WaitGroup) {
	fmt.Println("IN Create Snsapshot")
	cmd := exec.Command("restic", "-r", "s3:s3.amazonaws.com/resticbucket", "backup", dir, "--json")
	bs, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(bs), err)
	}
	fmt.Println(string(bs))
	snapregex := regexp.MustCompile("\"snapshot_id\":[^}]*")
	snap := snapregex.FindString(string(bs))
	spstr := strings.Split(snap, ":")
	fmt.Println(spstr[1])
	Snapmap[dir]=spstr[1]
	cswg.Done()
}

func restoreSnapshot(snapshotID string, rswg *sync.WaitGroup) bool {
	cmd := exec.Command("restic", "-r", "s3:s3.amazonaws.com/resticbucket", "restore", snapshotID, "--target", "/", "--json")
	bs, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(bs), err)
		return false
	}
	rswg.Done()
	fmt.Println(string(bs))
	return true
}



