package handler

import (
	"context"
	"github.com/micro/go-log"
	"fmt"
	"encoding/json"

	kubernetes "micro-go/kubernetes/proto/kubernetes"
	"micro-go/kubernetes/model/k8smodel"

	"k8s.io/client-go/tools/clientcmd"
	kubes "k8s.io/client-go/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Kubernetes struct{
	Repo *k8smodel.KubeModel
}

var clientset *kubes.Clientset

// Call is a single request handler called via client.Call or the generated client code
func (e *Kubernetes) Call(ctx context.Context, req *kubernetes.Request, rsp *kubernetes.Response) error {
	log.Log("Received Kubernetes.Call request")
	rsp.Msg = "Hello " + req.Podname
	return nil
}

func (h *Kubernetes) PodGet(ctx context.Context, req *kubernetes.Request, rsp *kubernetes.Response) error {
	kk :=
		`apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUN5akNDQWJLZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQ0FYRFRFNU1EUXpNREEzTXpFME1sb1lEekl4TVRnd05EQTJNRGN6TVRReVdqQVZNUk13RVFZRApWUVFERXdwcmRXSmxjbTVsZEdWek1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBCnhDczBrRXZnSFBHNWNIbm1UdWdrNkQrTmV5WmkraEpXM0h6ZW04NVZuemhhdHRiMmVhb3JZQU9nVUlJUXlIZGMKaVRhNi9tQU1udmxCdWtBODZ3NEZ3dXdrQUYwaFZFZ25IbFFjbUUyeTl1UmIyNUN0ME0wbkVENU8wQUpGdTBkUwpsVkZPT1hmcWR1RGd6cWl0cityMDJXd3Jmamk5ZWh0Y3lwNGpCT2Z5Tzd4cEtPcHJrUStQdnRyUGRST2ZTODFECkhTcVQvMklhT0tBdUpQaGllSmdsL014akFvYUk2M1JDRE9yanc4aUYyTktaby9BaTI3STR0Q0VaQ3ZnSW8vcWMKbCtBdGNDZlZlQVdoR0ZvUFhTRWxMSGQ5MkEwcXhNc3BqWHdRbWJpYWcxUzZnTGJqRm9Bbk95aWJPK1dzL1BCNQptSnpDM1dNZEFLbDc2NDJnZ1p5ZHBRSURBUUFCb3lNd0lUQU9CZ05WSFE4QkFmOEVCQU1DQXFRd0R3WURWUjBUCkFRSC9CQVV3QXdFQi96QU5CZ2txaGtpRzl3MEJBUXNGQUFPQ0FRRUFEeTVPODhKMEp0V3EvYzljRHJnNVZWeXQKUjQwVTZZSGFzN1lyVDlrVU1MWTQ3TkIraXhKK3lYUVRVRTRyRWEvMkxwUzQxRmsrblNCN0dCOHhRZkhyUHBYSApPUWdyNmY5OXN3VVYwQXVJM3ArRWpiZklUdnNCUzJkdmg5Ly8xanU2UHp4aThBb3dBdW1yQlZXK201aXd6VmVxCmZBMEkyZlN1L2VVd25pNGZlSlowS2tyM0plendMWE5Cb0pUaGV6RFhTcStaNzFUOXRwQlhXTEZQcjdwNTdwVzcKTURHN0JzajZyZG11cWZnUWMvcWFLY0QwVkFpTzVYOEFDQ1ZyVnAwYU9nTnFYVFhLQUUwUlM1UVkrMTFrQndPawpZTGt2dmhmQUlrNStmSk8yZUZCMDJOK250dkNLclkxdkVpQ2tTVWFUKzRuVzVGNjhjUitGS29lZXhtQjBqUT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
    server: https://10.100.2.30:6443
  name: kubernetes
contexts:
- context:
    cluster: kubernetes
    user: kubernetes-admin
  name: kubernetes-admin@kubernetes
current-context: kubernetes-admin@kubernetes
kind: Config
preferences: {}
users:
- name: kubernetes-admin
  user:
    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUM5RENDQWR5Z0F3SUJBZ0lJTWxrWEVKNmt0TFV3RFFZSktvWklodmNOQVFFTEJRQXdGVEVUTUJFR0ExVUUKQXhNS2EzVmlaWEp1WlhSbGN6QWdGdzB4T1RBME16QXdOek14TkRKYUdBOHlNVEU0TURRd05qQTNNekUwTlZvdwpOREVYTUJVR0ExVUVDaE1PYzNsemRHVnRPbTFoYzNSbGNuTXhHVEFYQmdOVkJBTVRFR3QxWW1WeWJtVjBaWE10CllXUnRhVzR3Z2dFaU1BMEdDU3FHU0liM0RRRUJBUVVBQTRJQkR3QXdnZ0VLQW9JQkFRREw2dXhGUGM3NUdpQUUKK3ZTaW50eDlmcWNXdkJJRUMwSWhTV1hJMk9Xd01IazE2VCsrdHBJMlVFTWlVWldHSDVuTy9lYnd4VHNXd2RQQgpLdzZmMmMyelBGYnFKTVlMdjJ0a2JOVy9rYlh6aGdIc3htUEo3cEZlellvQkZ3c3Znd0VWNktGOVZaSThFSkNFCi84aUVQUUFqQ2hnWjd3WDBoaEE4THNDeXpvVGlZbEhYVmhHcmUyNVVSR3MycUppOG1LbXdBa1Z3U3gxeDQydHYKM3A5NE1aYnNDZEwwQW1MQWQ0cHl6UEFBTHU3ZXdwMjFudTFzQ1lnTEV3MWoxeEVNalYxSXJlZVM3RC9iaUw0WApoOENVVWdCUktybTYwblk1ZEF3RTluT004UExLekhJMlBmNWcrTlBDeDlYTmljVEd6bjI2bGU5ZlVhNmRRSGxHCnpLTGx3b0s1QWdNQkFBR2pKekFsTUE0R0ExVWREd0VCL3dRRUF3SUZvREFUQmdOVkhTVUVEREFLQmdnckJnRUYKQlFjREFqQU5CZ2txaGtpRzl3MEJBUXNGQUFPQ0FRRUFlVXd2cmZxbXNpSU53YkQ1emx3dVdORTJWdUFOSUkvTQpNRU1ZZ3RMY1BKL2NLR0JGSFlZamUyRDh1bkVrempTbnROQ09vRDF3R2FzMGVEUEpYWTFSREpNTXFZZFdMeXNrCkNUenZiTXd3Qk9sUEdhemVLNE50cHhQeHFzeTgwUml5R3ZlZlZYYlZrNTE4dm1IbFhkaDM4K2d1L3R2cTdMUi8KZ01ZelBMVkNuYkxWSENzY1MydTFxWXBQbzgxZnhmZWtkT1FUU3lyblo4cnNPSHJSa0Z3aHJ4VzlxZXlxd3hnYwpoNnpqazNITXZNSlhhV050a3k3cEx1Sys3S0dUNHNOMDMyQUUxOXNvaTRKQVIzNWVqSnRjMFhXTmtmV2RiaW1WCkd6UEFFWnFVbXd0MWJKZjBLN2h5Z013dlpVR2dIMlIwZmdmcGJxaVhuL0RreWNFR1lheVJJUT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBeStyc1JUM08rUm9nQlByMG9wN2NmWDZuRnJ3U0JBdENJVWxseU5qbHNEQjVOZWsvCnZyYVNObEJESWxHVmhoK1p6djNtOE1VN0ZzSFR3U3NPbjluTnN6eFc2aVRHQzc5clpHelZ2NUcxODRZQjdNWmoKeWU2UlhzMktBUmNMTDRNQkZlaWhmVldTUEJDUWhQL0loRDBBSXdvWUdlOEY5SVlRUEM3QXNzNkU0bUpSMTFZUgpxM3R1VkVSck5xaVl2Smlwc0FKRmNFc2RjZU5yYjk2ZmVER1c3QW5TOUFKaXdIZUtjc3p3QUM3dTNzS2R0Wjd0CmJBbUlDeE1OWTljUkRJMWRTSzNua3V3LzI0aStGNGZBbEZJQVVTcTV1dEoyT1hRTUJQWnpqUER5eXN4eU5qMysKWVBqVHdzZlZ6WW5FeHM1OXVwWHZYMUd1blVCNVJzeWk1Y0tDdVFJREFRQUJBb0lCQUJDVmtKV3BDell6Szd0Nwo5WDFzVWg1YXVKZ2V2NUJZb2c0MisvSXp3YXBzcHM3OW8xT05ZYUxOTUVpUVBncmtjd3ZrbG0ycWZMM0RFY2U2CkNyeHhhRXRWY3ZRai9YOWpHQWJZMnlGelprcmg4VElpaUdjWjRmVEYzcDFzRkZyTzJyMi9aN0xiek1MWnpnY2oKRGxuQzVFbjFQalNOdmZ3L0V2N3lCbllFV2NaWGQybGs0MXd2UlJvVXdFaFRLeDNqL3N2am9NMER3T0hqS29vVwpENmlCaUE3OHhvbGVGbGJ1ZjkzZHc4a2taREx0VDFQd3ZDOWxKVE9qS2JQc1hvMnczenNvQnFkTzh0aTl4QnlvCm5ndVQ5UDJMcGRTaXJoQjBHZ3I5TGtjUmtpWk5veEJ6KzM3NSs3WWN1L1dVdHFsOEk0aVYxZk5RNDhCMG02NmgKL1crMEhnRUNnWUVBMGpWY3RYZFBkTjZKYmNpemdGbHBkbjY0UUxLUVVPM2Zka0txUzVtQy9MenF2UmJjc1lFNApOajFHNnNaSGhIOG56bGhkaE1QWlBPU3JqSnNFb1hVWTdoeGtmZXp6TFRwUVJ2cTBCQkZtbTc5WmIwdE5xZmxSCk5STmQ5eUcvRjZNazl3cC9iYUd2YStYRHE2bCtUY1U4MjhXWlJQZ0hQc0p1N2taa0RFZy9KQmtDZ1lFQStGYSsKb1pxS29MdEk4d0RjRE9mS05UY2czeWFad2xtL0tJSXdUaXpSV1ZHUklhd2diRksxRDJtYkV3UG80MWlGRnBTdAoyc1VpaHo3NmFSbjlRWHJiUlRHQ2UwSmd0cyt2SjQzelp3SXk3aWJHcUhiOHBhMFBhV0lGN3dHaHpiWm80Vi9hCnp6eDZYUUVFY3drRkJHV2YvaVBNYXZMVFhnVGdJeko0dVIyVko2RUNnWUJRR0FzQ0RiclZzZkUxUm5LMnBkcmMKVG41UUVIbVNqSUJIcERFVTZ5SVF3TDdFVzVDdGhhbndhTHE3dTk4R0toajFzNCsyaFpVaGNaTzMyRjBVVS9TeQp6Vml3N25iZHRjbGVzaW1qSHlvMGo5MDQzYjF5MVU5TzVOazV6NzdxOTd1ZThYNEtQTUFGWVorRHFlbzVJYjBGCmxDM0pMS1ErRW9HSUFvUWVkZXl4Q1FLQmdGY0hWMEVjL28zN2RFSXN2L0VKM2ZRdVFLZlRTRGt5NkduU3pnam4KYmVwR1NuMHBiQ2RxQ0tmQkU2WmpUVjgyNjFVUmY2ZldSQlp2ZktRT2hwRkc4OGJZNWVnNlI5aHBDZkhycml3ZwpPaWI1a2M4QisrSTN0dG53anpIR09hVDAwdWFBaUJidUU5ZTUvb3d0MGdOTGRmbERKR1VxZWVaT0FScDBtcGdOCkJQTmhBb0dCQUlORmc1YjA4VmJ1U2EraWhPWG40QW9JazNlZlJVc1NQOXMvdFh1b2RNZVhDL1R6bGhSRENtT0gKNVN5VGFLaytadEJnMS9NYWZScndhOHppenYrMXhaclBQMkU3OXN4cTUyb3ozUlVFWUh4WFBGelNHSzU0cHFRRgpNdnlxSjRJRllRL0d4OWdwVXNmTmVySWJCOUF2alJaWUpCTmhoZGdVRmRPR0pKRWlSZXpMCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==`
	fmt.Println([]byte(kk))
	dd := []byte(kk)
	// getConfig from config_file;
	//k8sconfig := flag.String("k8sconfig", "k8sconfig", "kubernetes config file path")
	//flag.Parse()
	//config, err := clientcmd.BuildConfigFromFlags("", *k8sconfig)
	configs, err := clientcmd.NewClientConfigFromBytes(dd)
	config,_ := configs.ClientConfig()

	if err != nil {
		log.Fatal(err)
	}
	clientset, err = kubes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("connect k8s success")
		fmt.Println("############################################")
	}

	//获取POD
	pods, err := clientset.CoreV1().Pods("devops").List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(pods)

	fmt.Println(pods.Items[1].Name)
	fmt.Println(pods.Items[1].CreationTimestamp)
	fmt.Println(pods.Items[1].Spec.Containers)
	fmt.Println(pods.Items[1].Labels)
	fmt.Println(pods.Items[1].Namespace)
	fmt.Println(pods.Items[1].Status.HostIP)
	fmt.Println(pods.Items[1].Status.PodIP)
	fmt.Println(pods.Items[1].Status.StartTime)
	fmt.Println(pods.Items[1].Status.Phase)
	fmt.Println(pods.Items[1].Status.ContainerStatuses[0].RestartCount) //重启次数
	fmt.Println(pods.Items[1].Status.ContainerStatuses[0].Image)        //获取重启时间

	//获取NODE
	fmt.Println("##################")
	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	mjson, _ := json.Marshal(nodes.Items[0])
	fmt.Println(mjson)

	fmt.Println(nodes.Items[0].Name)
	fmt.Println(nodes.Items[0].CreationTimestamp) //加入集群时间
	fmt.Println(nodes.Items[0].Status.NodeInfo)
	fmt.Println(nodes.Items[0].Status.Conditions[len(nodes.Items[0].Status.Conditions)-1].Type)
	fmt.Println(nodes.Items[0].Status.Allocatable.Memory().String())
	return nil
}


