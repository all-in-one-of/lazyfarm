package main

import (
	"fmt"
	"errors"
	"net"
	"log"
	"strconv"
	"os"
)

func removeDuplicates(a []int) []int {
        result := []int{}
        seen := map[int]int{}
        for _, val := range a {
                if _, ok := seen[val]; !ok {
                        result = append(result, val)
                        seen[val] = val
                }
        }
        return result
}

type intSlice []int

func (slice intSlice) pos(value int) int {
    for p, v := range slice {
        if (v == value) {
            return p
        }
    }
    return -1
}

func jobToTasks(job *Job) []Task {
	fmt.Println("job to tasks")
	nframes := len(job.Frames)
	tasks := make([]Task, nframes)
	for i := 0 ; i < nframes ; i++ {
		tasks[i].Cmd = job.Cmd
		tasks[i].Frame = job.Frames[i]
	}
	return tasks
}

func localIP() (net.IP, error) {
        tt, err := net.Interfaces()
        if err != nil {
                return nil, err
        }
        for _, t := range tt {
                aa, err := t.Addrs()
                if err != nil {
                        return nil, err
                }
                for _, a := range aa {
                        ipnet, ok := a.(*net.IPNet)
                        if !ok {
                                continue
                        }
                        v4 := ipnet.IP.To4()
                        if v4 == nil || v4[0] == 127 { // loopback address
                                continue
                        }
                        return v4, nil
                }
        }
        return nil, errors.New("cannot find local IP address")
}

func findMyAddress() string {
	ip, err := localIP()
	if err != nil {
		log.Fatal(err)
	}

	port := 8082
	for i :=0; i < 10; i++ {
		address := ip.String()+":"+strconv.Itoa(port)
		ln, err := net.Listen("tcp", address)
		if err != nil {
			port++
			continue
		}
		ln.Close()
		fmt.Printf("my address is %v\n", address)
		return address
	}
	fmt.Printf("cannot find good port")
	os.Exit(1)
	return ""
}
