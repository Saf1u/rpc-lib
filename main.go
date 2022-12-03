package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/saf1u/rpc-lib/handlers"
)

//protoc -I=$curr_dir --go_out=$curr_dir/handlers $curr_dir/proto/task.proto

func main() {
	task := &handlers.Task{TaskName: "rando-task"}
	data, err := proto.Marshal(task)
	if err != nil {
		log.Fatal("unable to seralize message")
	} else {
		copy := &handlers.Task{}
		err = proto.Unmarshal(data, copy)
		if err != nil {
			panic(err)
		}
		fmt.Println(copy.GetTaskName())
	}
}
