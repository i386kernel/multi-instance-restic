package main

import (
	"fmt"
	"github.com/i386kernel/resticweb/config"
	"github.com/i386kernel/resticweb/pkg/kubeinteract"
	"github.com/i386kernel/resticweb/pkg/restic"
	"github.com/joho/godotenv"
	"log"
	"os"
)


func init() {
	fmt.Println("Initializing the Init function...")
	fmt.Println("Loading the Environment Variables...")
	if err := godotenv.Load("RESTICENV.env"); err != nil{
		log.Println("Unable to load .env file")
	}
}

var kubeobjectinit = kubeinteract.KubeObject{}

func main() {
	conf := config.New()

	fmt.Println(conf.Restic.ResticPassword)

	kubeobjectinit = kubeinteract.KubeObject{
		Namespace: "wordpress-auto-31379",
		BaseURL:   os.Getenv("KUBE_PR"),
		Token:     os.Getenv("KUBE_PR_TOKEN"),
	}

	// fmt.Println(kubeobjectinit)
	// CreateSnapshots()
	RestoreSnapshots()
}


func CreateSnapshots(){
	// Read the namespace and get the PVC's and PV's.
	pvs := kubeobjectinit.GetPVNamesFromPVCs()

	// Get the NFS Mount Path by going through the mountpath in PV mapped for nfs
	kubeobjectinit.GetNFSPathFromPV(pvs)

	fmt.Println("Protected PV Paths:", kubeinteract.ProtectedPVPaths)
	// Create Snapshots
	restic.CreateConcurrentSnapshots(kubeinteract.ProtectedPVPaths)
}

func RestoreSnapshots(){
	// Snapshots Restore
	snapshotsID := []string{"459673db", "9a02b082"}
	restic.RestoreConcurrentSnapshots(snapshotsID)
}