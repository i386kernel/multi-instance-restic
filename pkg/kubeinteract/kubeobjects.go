package kubeinteract

import (
	"encoding/json"
	"fmt"
	"log"
	)

type KubeObject struct {
	Namespace string
	BaseURL string
	Token string
}

type ProtectedVolumes struct {
	ProtectedPVs []string
}

var ProtectedPVPaths []string

func (ko *KubeObject) getPVCsInNamespace() (int, []byte) {
	ka := KubeAccess{
		BaseURL:  ko.BaseURL,
		EndPoint: fmt.Sprintf("/api/v1/namespaces/%s/persistentvolumeclaims/", ko.Namespace),
		Token:    ko.Token,
	}
	status, body := ka.KubeGet()
	return status, body
}


func (ko *KubeObject) getPodDetails() (int, []byte) {
	ka := KubeAccess{
		BaseURL:  ko.BaseURL,
		EndPoint: fmt.Sprintf("/api/v1/namespaces/%s/pods", ko.Namespace),
		Token:    ko.Token,
	}
	status, body := ka.KubeGet()
	return status, body
}


func (ko *KubeObject) getPVs() (int, []byte) {
	ka := KubeAccess{
		BaseURL:  ko.BaseURL,
		EndPoint: "/api/v1/persistentvolumes/",
		Token:    ko.Token,
	}
	status, body := ka.KubeGet()
	return status, body
}


func (ko KubeObject)GetPVNamesFromPVCs() []string {

	// Extract Volume Names from PVC
	type PVCSpecs struct {
		VolumeName string `json:"volumeName"`
	}
	type PVCMeta struct {
		Name string `json:"name"`
	}
	type PVC struct {
		Spec     PVCSpecs `json:"spec"`
		Metadata PVCMeta  `json:"metadata"`
	}
	type PVCList struct {
		Items []PVC `json:"items"`
	}

	_, body := ko.getPVCsInNamespace()

	var volumenames PVCList
	err := json.Unmarshal(body, &volumenames)
	if err != nil {
		log.Println(err)
	}

	var volnames []string

	for _, v := range volumenames.Items {
		volnames = append(volnames, v.Spec.VolumeName)
		fmt.Printf("PVC Name: %s, Attached PV Name: %s\n", v.Metadata.Name, v.Spec.VolumeName)
	}
	return volnames
}


func (ko *KubeObject)GetNFSPathFromPV(volnames []string) {
	type NFS struct {
		Path   string `json:"path"`
		Server string `json:"server"`
	}
	type Spec struct {
		Nfs NFS `json:"nfs"`
	}
	type Metadata struct {
		Name string `json:"name"`
	}
	type PV struct {
		Metadata Metadata `json:"metadata"`
		Spec     Spec     `json:"spec"`
	}
	type PVList struct {
		Items []PV `json:"items"`
	}
	_, pvbody := ko.getPVs()

	volumes := PVList{}
	err := json.Unmarshal(pvbody, &volumes)
	if err != nil {
		log.Println("error unmarshalling to json", err)
	}
	for _, v := range volnames{
		for _, vol := range volumes.Items{
			if v == vol.Metadata.Name{
				ProtectedPVPaths = append(ProtectedPVPaths, vol.Spec.Nfs.Path)
			}
		}
	}
}
