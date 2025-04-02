package agent

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
	"github.com/hildam/eino-mcp/conf"
)

// NewAgent 新建 agent
// mcp tools 与 eino agent 绑定
func NewAgent(ctx context.Context, tools []tool.BaseTool) (*react.Agent, error) {
	// 初始化 chatModel
	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		BaseURL: conf.GetCfg().DeekSeek.BaseURL,
		Model:   conf.GetCfg().DeekSeek.ModelID,
		APIKey:  conf.GetCfg().DeekSeek.APIKey,
	})
	if err != nil {
		log.Printf("NewAgent failed, init chat model failed, err: %+v", err)
		return nil, err
	}

	// 新建 agent
	agent, err := react.NewAgent(ctx, &react.AgentConfig{
		Model:       chatModel,
		ToolsConfig: compose.ToolsNodeConfig{Tools: tools},
	})
	if err != nil {
		log.Printf("NewAgent failed, init react agent failed, err: %+v", err)
		return nil, err
	}
	return agent, nil
}

// Run 运行 agent
func Run(ctx context.Context, agent *react.Agent) {
	scanner := bufio.NewScanner(os.Stdin)
	log.Printf("Agent is running...")
	inputTips := "Please input: "
	for {
		fmt.Print(inputTips)

		if !scanner.Scan() {
			log.Printf("Read input failed, err: %+v", scanner.Err())
			return
		}

		input := scanner.Text()
		switch strings.ToLower(input) {
		case "exit":
			log.Printf("Agent is exiting...")
			return
		default:
			// 调用 agent
			rsp, err := agent.Generate(ctx, []*schema.Message{
				{
					Role:    schema.User,
					Content: strings.Replace(input, inputTips, "", 1),
				},
			})
			if err != nil {
				log.Fatalf("Agent generate failed, err: %+v", err)
			}
			fmt.Println(rsp.Content)
		}
	}
}
