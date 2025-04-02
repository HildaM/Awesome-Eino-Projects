# Eino MCP

基于 Eino 框架开发的 MCP (Message Control Protocol) 示例项目。

A MCP (Message Control Protocol) example project built on Eino framework.

## 项目简介 | Introduction

本项目展示了如何使用 Eino 框架实现 MCP 协议，包括：
- MCP 代理实现
- 配置管理
- 工具集成
- 协议扩展示例

This project demonstrates how to implement MCP protocol using Eino framework, including:
- MCP agent implementation
- Configuration management
- Tool integration
- Protocol extension examples

## 项目结构 | Project Structure

```
eino-mcp/
├── agent/          # MCP agent implementation
├── conf/           # Configuration management
├── tools/          # Tool implementations
├── main.go         # Entry point
├── config.yaml     # Configuration file
└── go.mod          # Go module definition
```

## 快速开始 | Quick Start

### 安装依赖 | Install Dependencies

```bash
go mod tidy
```

### 配置设置 | Configuration

1. 复制配置文件示例：
   ```bash
   cp config.yaml.example config.yaml
   ```

2. 编辑 `config.yaml` 文件，配置必要的参数：
   ```yaml
   mcp:
     host: "localhost"
     port: 8080
     # 其他配置项...
   ```

### 运行项目 | Run Project

```bash
go run main.go
```

## 功能特性 | Features

- MCP 协议实现
- 配置热加载
- 工具集成支持
- 代理服务

## 开发计划 | Development Plan

- [ ] 完善错误处理机制
- [ ] 添加更多工具支持
- [ ] 优化配置管理
- [ ] 添加单元测试

## 贡献指南 | Contributing

欢迎提交 Pull Request 或 Issue 来完善项目。

Feel free to submit Pull Requests or Issues to improve the project. 