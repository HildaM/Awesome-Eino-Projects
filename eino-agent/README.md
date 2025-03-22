# Eino Agent

这是一个基于 Eino 框架的智能代理示例项目。

## 配置说明

本项目使用 YAML 格式的配置文件。按照以下步骤进行配置：

1. 在项目根目录复制示例配置文件
   ```bash
   cp config.yaml.example config.yaml
   ```

2. 编辑 `config.yaml` 文件，填入您的 API 密钥和模型 ID

   ```yaml
   DeekSeek:
     api_key: "your-actual-api-key"
     model_id: "deepseek-chat"
     base_url: "https://api.deepseek.com/v1"
   ```

3. 配置文件会自动热加载，修改后无需重启程序

## 启动方式

```bash
go run main.go
```

也可以指定配置文件路径：

```bash
go run main.go --config /path/to/your/config.yaml
```

## 使用体验

```
(base) PS D:\Documents\Codes\ai_projects\eino-examples\eino-agent> go run .\main.go
WARNING:(ast) sonic only supports go1.17~1.23, but your environment is not suitable
2025/03/22 22:20:39 loadConfig debug, success load file from D:\Documents\Codes\ai_projects\eino-examples\eino-agent\config.yaml
欢迎使用员工信息 Agent, 支持用户信息的增删改查，输入 'exit' 退出程序。

请输入操作: show all employee
2025/03/22 22:20:54 姓名: HildaM, 年龄: 20, 部门: IT

请输入操作: add employee 哈基米, and guess how old he is and what job he do. Add infomation for him
2025/03/22 22:21:56 哈基米 added, 年龄: 28, 部门: Software Engineer

请输入操作: get HildaM
2025/03/22 22:22:16 姓名: HildaM, 年龄: 20, 部门: IT

请输入操作: show all
2025/03/22 22:22:28 姓名: HildaM, 年龄: 20, 部门: IT
2025/03/22 22:22:28 姓名: 哈基米, 年龄: 28, 部门: Software Engineer

请输入操作: delete 哈基米
2025/03/22 22:22:45 哈基米 deleted

请输入操作: show all
2025/03/22 22:23:00 姓名: HildaM, 年龄: 20, 部门: IT

请输入操作: exit
欢迎再次使用，再见。
```
