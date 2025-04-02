package main

import (
	"context"
	"log"

	eino_mcp "github.com/cloudwego/eino-ext/components/tool/mcp"

	"github.com/hildam/eino-mcp/agent"
	"github.com/hildam/eino-mcp/conf"
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
)

func main() {
	// 初始化配置
	if err := conf.Init(); err != nil {
		log.Fatal("init config failed")
	}

	// 初始化 mcp client
	ctx := context.Background()
	cli, err := client.NewSSEMCPClient("http://localhost:8080/sse")
	if err != nil {
		log.Fatal("init mcp client failed")
	}
	cli.Start(ctx)
	defer cli.Close()

	// 发送 init request
	initReq := mcp.InitializeRequest{}
	initReq.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initReq.Params.ClientInfo = mcp.Implementation{
		Name:    "current-time",
		Version: "0.0.1",
	}
	initRsp, err := cli.Initialize(ctx, initReq)
	if err != nil {
		log.Printf("send init req failed, req: %+v, err: %+v", initReq, err)
		return
	}
	log.Printf("send init req success, rsp: %+v", initRsp)

	// 查询 mcp tools
	tools, err := eino_mcp.GetTools(ctx, &eino_mcp.Config{Cli: cli})
	if err != nil {
		log.Printf("get tools failed, err: %+v", err)
		return
	}
	log.Printf("get tools success, tools: %+v", tools)

	// 创建 agent 并运行
	myAgent, err := agent.NewAgent(ctx, tools)
	if err != nil {
		log.Printf("create agent failed, err: %+v", err)
		return
	}
	agent.Run(ctx, myAgent)
}
