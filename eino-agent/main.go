package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/HildaM/eino-examples/eino-agent/config"
	"github.com/HildaM/eino-examples/eino-agent/userinfo"
)

func main() {
	// 初始化配置
	if err := config.Init(); err != nil {
		log.Fatalf("初始化配置失败: %v", err)
	}

	ctx := context.Background()
	// 创建 Agent
	agent := userinfo.NewAgent(ctx)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("欢迎使用员工信息 Agent, 支持用户信息的增删改查，输入 'exit' 退出程序。")
	inputTips := "\n请输入操作: "
	for {
		fmt.Print(inputTips)
		if !scanner.Scan() {
			fmt.Println("读取输入失败，程序退出。")
			break
		}

		input := scanner.Text()

		switch strings.ToLower(input) {
		case "exit":
			fmt.Println("欢迎再次使用，再见。")
			return
		default:
			agent.Invoke(ctx, strings.Replace(input, inputTips, "", 1))
		}
	}
}
