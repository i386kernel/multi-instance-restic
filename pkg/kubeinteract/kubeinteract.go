package kubeinteract

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type KubeAccess struct {
	BaseURL string
	EndPoint string
	Token string
	Body string
}

type KubePostData struct {
	EndPoint string
	Body string
}

type HttpClientParams struct {
	InsecureSkipVerify bool
}

func (hc* HttpClientParams)NewHttpClient() http.Client{
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: hc.InsecureSkipVerify}}
	return http.Client{Transport: tr}
}

func (ka *KubeAccess) KubeGet()(int, []byte){

	request, err := http.NewRequest("GET", ka.BaseURL+ka.EndPoint, nil)
	if err != nil {
		fmt.Println(err)
	}
	request.Header = map[string][]string{"Content-type": {"application/json"}, "Authorization": {"Bearer " + ka.Token}}
	clientparams := HttpClientParams{InsecureSkipVerify: true}
	client := clientparams.NewHttpClient()
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.StatusCode)
	if resp.StatusCode > 210 {
		fmt.Println(errors.New("unable to fulfill the get request"))
		fmt.Println(string(body))
	}
	return resp.StatusCode, body

}

func (ka *KubeAccess) KubePost()(int, []byte){

	request, err := http.NewRequest("POST", ka.BaseURL+ka.EndPoint, bytes.NewBuffer([]byte(ka.Body)))
	if err != nil {
		fmt.Println(err)
	}
	request.Header = map[string][]string{"Content-type": {"application/json"}, "Authorization": {"Bearer " + ka.Token}}

	clientparams := HttpClientParams{InsecureSkipVerify: true}
	client := clientparams.NewHttpClient()

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode > 210 {
		fmt.Println(string(body))
	}
	return resp.StatusCode, body
}

func (ka *KubeAccess) KubePatch()(int, []byte){

	request, err := http.NewRequest("PATCH", ka.BaseURL+ka.EndPoint, bytes.NewBuffer([]byte(ka.Body)))
	if err != nil {
		fmt.Println(err)
	}
	request.Header = map[string][]string{"Content-type": {"application/strategic-merge-patch+json"}, "Authorization": {"Bearer " + ka.Token}}
	clientparams := HttpClientParams{InsecureSkipVerify: true}
	client := clientparams.NewHttpClient()
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	return resp.StatusCode, body
}

