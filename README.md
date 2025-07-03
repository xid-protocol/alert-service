# Alert Service

一个用于发送飞书(Lark)通知的服务，支持通过webhook发送消息到飞书群组。

## ✨ 功能特性

- 支持发送文本消息到飞书群组
- 支持Docker容器化部署
- 自动化CI/CD流程（GitHub Actions）
- 支持多平台架构（amd64, arm64）
- 安全扫描和漏洞检测

## 🚀 快速开始

### 本地运行

1. 克隆仓库：
```bash
git clone https://github.com/xidp-protocol/alert-service.git
cd alert-service
```

2. 创建配置文件 `/opt/xidp/conf/alert.yml`：
```yaml
Lark:
  custom_bot: "https://open.larksuite.com/open-apis/bot/v2/hook/your-webhook-url"
```

3. 运行程序：
```bash
go run cmd/main.go
```

### Docker 部署

#### 使用预构建镜像

```bash
# 拉取镜像
docker pull your-username/alert-service:latest

# 运行容器
docker run -d --name alert-service \
  -v /path/to/your/alert.yml:/opt/xidp/conf/alert.yml:ro \
  your-username/alert-service:latest
```

#### 本地构建

```bash
# 构建镜像
docker build -t alert-service .

# 运行容器
docker run -d --name alert-service \
  -v /path/to/your/alert.yml:/opt/xidp/conf/alert.yml:ro \
  alert-service
```

## 🔧 配置说明

### 配置文件格式

配置文件位置：`/opt/xidp/conf/alert.yml`

```yaml
Lark:
  custom_bot: "https://open.larksuite.com/open-apis/bot/v2/hook/xxxxxxxx"
```

### 环境变量

- 无需额外环境变量，所有配置通过配置文件提供

## 📦 GitHub Actions 部署

### 配置 Secrets

在 GitHub 仓库中设置以下 Secrets：

1. 进入仓库 Settings → Secrets and variables → Actions
2. 添加以下 Repository secrets：
   - `DOCKER_USERNAME`: 你的 Docker Hub 用户名
   - `DOCKER_PASSWORD`: 你的 Docker Hub 密码或访问令牌

### 自动部署

#### 自动触发部署

- **推送到 main/master 分支**: 自动构建并推送 `latest` 标签
- **创建版本标签**: 推送形如 `v1.0.0` 的标签会自动部署相应版本

```bash
# 创建和推送版本标签
git tag v1.0.0
git push origin v1.0.0
```

#### 手动部署

1. 进入 GitHub 仓库页面
2. 点击 "Actions" 标签
3. 选择 "Manual Deploy to Docker Hub" 工作流
4. 点击 "Run workflow"
5. 填写参数：
   - **Version**: 版本号（如 v1.0.0）
   - **Push latest**: 是否同时推送 latest 标签
   - **Environment**: 部署环境（production/staging/development）

### 工作流说明

#### `dockerhub.yml` - 自动部署
- 在代码推送或PR时触发
- 支持多平台构建（amd64, arm64）
- 包含安全扫描
- 自动生成标签和元数据

#### `manual-deploy.yml` - 手动部署
- 支持手动指定版本
- 可选择部署环境
- 生产环境部署时自动创建 GitHub Release
- 详细的部署摘要

## 🔐 安全特性

- 使用多阶段构建减少镜像大小
- 非root用户运行
- 集成 Trivy 安全扫描
- SARIF 报告上传到 GitHub Security

## 📋 API 说明

### SendToLark 函数

```go
func SendToLark(text string, webhookURL string)
```

发送文本消息到飞书群组。

**参数:**
- `text`: 要发送的消息内容
- `webhookURL`: 飞书机器人的webhook地址

**消息格式:**
```json
{
  "msg_type": "text",
  "content": {
    "text": "your message here"
  }
}
```

## 🛠️ 开发

### 项目结构

```
alert-service/
├── cmd/
│   └── main.go          # 主程序入口
├── .github/
│   └── workflows/       # GitHub Actions 工作流
├── Dockerfile           # Docker 构建文件
├── .dockerignore        # Docker 忽略文件
├── deploy.sh           # 本地部署脚本
├── go.mod              # Go 模块文件
└── README.md           # 项目说明
```

### 本地部署脚本

使用提供的部署脚本：

```bash
# 给脚本执行权限
chmod +x deploy.sh

# 部署到 Docker Hub
./deploy.sh v1.0.0

# 或使用环境变量
export DOCKER_USERNAME=your-username
export DOCKER_PASSWORD=your-password
./deploy.sh v1.0.0
```

## 📄 许可证

本项目采用 MIT 许可证。详情请查看 [LICENSE](LICENSE) 文件。

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📞 支持

如有问题，请在 GitHub Issues 中提出。