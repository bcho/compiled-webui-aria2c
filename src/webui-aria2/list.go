// List tasks.
package main

import (
	"encoding/json"
	"os/exec"
	"strings"
)

type Tasks struct {
	Active  []Task `json:"active"`
	Waiting []Task `json:"waiting"`
	Stopped []Task `json:"stopped"`
}

type Task struct {
	Gid    string `json:"gid"`
	Status string `json:"status"`
}

func (t Task) IsError() bool {
	return t.Status == "error"
}

const ListScript = `
# coding: utf-8

'''
    list
    ~~~~

    List all tasks.

    TODO: This script should be included in the go package...
'''

import json
import xmlrpclib


def connect(host, port):
    server = 'http://{0}:{1}/rpc'.format(host, port)
    return xmlrpclib.ServerProxy(server)


def list_tasks(server):
    ALL_TASKS = 2 << 10

    return {
        'active': server.aria2.tellActive(),
        'waiting': server.aria2.tellWaiting(0, ALL_TASKS),
        'stopped': server.aria2.tellStopped(0, ALL_TASKS)
    }


if __name__ == '__main__':
    import sys

    if len(sys.argv) < 3:
        host, port = 'localhost', '6800'
    else:
        host, port = sys.argv[1], sys.argv[2]

    server = connect(host, port)

    print(json.dumps(list_tasks(server)))
`

func listTasks() (*Tasks, error) {
	script := strings.NewReader(ListScript)
	cmd := exec.Command("python")
	cmd.Stdin = script
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var tasks Tasks
	err = json.Unmarshal(output, &tasks)
	if err != nil {
		return nil, err
	}

	return &tasks, nil
}
