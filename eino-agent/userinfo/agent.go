package userinfo

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/HildaM/eino-examples/eino-agent/config"
	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

type Agent struct {
	runnable compose.Runnable[[]*schema.Message, []*schema.Message]
}

func NewAgent(ctx context.Context) *Agent {
	cfg := config.GetCfg()

	// 检查API密钥是否设置
	if cfg.DeekSeek.APIKey == "" {
		log.Fatalf("NewAgent failed, api_key is empty")
		return nil
	}

	// 创建聊天模型
	// deepseek eino 框架暂不支持tool调用
	// chatModel, err := deepseek.NewChatModel(ctx, &deepseek.ChatModelConfig{
	// 	APIKey:  cfg.DeekSeek.APIKey,
	// 	Model:   cfg.DeekSeek.ModelID,
	// 	BaseURL: cfg.DeekSeek.BaseURL,
	// })
	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		BaseURL: cfg.DeekSeek.BaseURL,
		Model:   cfg.DeekSeek.ModelID,
		APIKey:  cfg.DeekSeek.APIKey,
	})

	if err != nil {
		log.Fatalf("创建聊天模型失败: %v", err)
		return nil
	}

	// 准备工具
	tools, toolInfos := UserInfoTools(ctx)

	// 尝试绑定工具，如果失败则提供有用的错误信息
	err = chatModel.BindTools(toolInfos)
	if err != nil {
		if strings.Contains(err.Error(), "DeepSeek temporarily does not support tool call") {
			log.Fatalf("错误: DeepSeek 目前不支持工具调用功能。\n"+
				"解决方案: 请考虑以下选项:\n"+
				"1. 等待 DeepSeek 更新支持工具调用功能\n"+
				"2. 修改代码使用支持工具调用的其他模型\n"+
				"具体错误: %v", err)
		} else {
			log.Fatalf("绑定工具失败: %v", err)
		}
		return nil
	}

	todoToolsNode, err := compose.NewToolNode(context.Background(), &compose.ToolsNodeConfig{
		Tools: tools,
	})
	if err != nil {
		log.Fatalf("创建工具节点失败: %v", err)
		return nil
	}

	chain := compose.NewChain[[]*schema.Message, []*schema.Message]()
	chain.
		AppendChatModel(chatModel, compose.WithNodeName("chat_model")).
		AppendToolsNode(todoToolsNode, compose.WithNodeName("tools"))

	agent, err := chain.Compile(ctx)
	if err != nil {
		log.Fatalf("编译链失败: %v", err)
		return nil
	}

	return &Agent{
		runnable: agent,
	}
}

func (a *Agent) Invoke(ctx context.Context, content string) {
	_, err := a.runnable.Invoke(ctx, []*schema.Message{
		{
			Role:    schema.User,
			Content: content,
		},
	})
	if err != nil {
		if strings.Contains(err.Error(), "no tool call found in input") {
			fmt.Println("暂不支持该操作")
			return
		}
		log.Printf("agent.Invoke failed, err=%v", err)
	}
}
