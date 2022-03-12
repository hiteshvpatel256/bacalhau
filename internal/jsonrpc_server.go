package internal

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/rpc"

	"github.com/filecoin-project/bacalhau/internal/logger"
	"github.com/filecoin-project/bacalhau/internal/types"
)

type JobServer struct {
	RequesterNode *RequesterNode
}

func (server *JobServer) List(args *types.ListArgs, reply *types.ListResponse) error {
	res, err := server.RequesterNode.Scheduler.List()
	if err != nil {
		return err
	}
	*reply = res
	return nil
}

func (server *JobServer) Submit(args *types.SubmitArgs, reply *types.Job) error {
	//nolint
	job, err := server.RequesterNode.Scheduler.SubmitJob(args.Spec, args.Deal)
	if err != nil {
		return err
	}
	*reply = *job
	return nil
}

func RunBacalhauJsonRpcServer(
	ctx context.Context,
	host string,
	port int,
	requesterNode *RequesterNode,
) {
	job := &JobServer{
		RequesterNode: requesterNode,
	}

	rpcServer := rpc.NewServer()
	err := rpcServer.Register(job)
	if err != nil {
		log.Fatalf("Format of service Job isn't correct. %s", err)
	}

	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: rpcServer,
	}

	isClosing := false
	go func() {
		err = httpServer.ListenAndServe()
		if err != nil && !isClosing {
			log.Fatalf("http.ListenAndServe failed: %s", err)
		}
	}()

	go func() {
		logger.Debug("Waiting for json rpc context to finish.")
		<-ctx.Done()
		logger.Debug("Closing json rpc server.")
		isClosing = true
		httpServer.Close()
		logger.Debug("Closed json rpc server.")
	}()
}
