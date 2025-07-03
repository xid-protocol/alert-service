# 部署环境配置指南

本文档详细说明如何在不同环境中配置和部署 Alert Service。

## 🌍 环境类型

### 1. 开发环境 (Development)
- 分支：`dev`, `feature/*`
- 镜像标签：`dev-*`
- 配置：开发测试用的webhook

### 2. 测试环境 (Staging)
- 分支：`staging`
- 镜像标签：`staging-*`
- 配置：测试环境webhook

### 3. 生产环境 (Production)
- 分支：`main`, `master`
- 镜像标签：`latest`, `v*.*.*`
- 配置：生产环境webhook

## 📦 Docker Hub 环境配置

### Repository Settings

在 Docker Hub 仓库页面配置：

1. **Repository Name**: `your-username/alert-service`
2. **Description**: "Alert Service for Lark notifications"
3. **Visibility**: Public/Private
4. **Build Settings**: 
   - Source: GitHub
   - Repository: `your-username/alert-service`
   - Autobuild: Enabled

### Environment Tags

| 环境 | 分支 | Docker标签 | 说明 |
|------|------|------------|------|
| 开发 | `dev` | `dev-latest` | 开发版本 |
| 测试 | `staging` | `staging-latest` | 测试版本 |
| 生产 | `main` | `latest`, `v1.0.0` | 生产版本 |

## 🔧 GitHub Secrets 配置

### 必需的 Secrets

在 GitHub 仓库 → Settings → Secrets and variables → Actions 中添加：

```bash
# Docker Hub 认证
DOCKER_USERNAME=your-docker-hub-username
DOCKER_PASSWORD=your-docker-hub-password-or-token

# 可选：不同环境的配置
DEV_LARK_WEBHOOK=https://open.larksuite.com/open-apis/bot/v2/hook/dev-webhook
STAGING_LARK_WEBHOOK=https://open.larksuite.com/open-apis/bot/v2/hook/staging-webhook
PROD_LARK_WEBHOOK=https://open.larksuite.com/open-apis/bot/v2/hook/prod-webhook
```

### 环境特定 Secrets

你也可以为不同环境创建不同的环境组：

#### Development Environment
```bash
LARK_WEBHOOK=https://open.larksuite.com/open-apis/bot/v2/hook/dev-webhook
```

#### Staging Environment
```bash
LARK_WEBHOOK=https://open.larksuite.com/open-apis/bot/v2/hook/staging-webhook
```

#### Production Environment
```bash
LARK_WEBHOOK=https://open.larksuite.com/open-apis/bot/v2/hook/prod-webhook
```

## 🚀 部署流程

### 自动部署

#### 1. 开发环境部署
```bash
# 推送到 dev 分支
git checkout dev
git commit -m "feat: add new feature"
git push origin dev
```

#### 2. 测试环境部署
```bash
# 推送到 staging 分支
git checkout staging
git merge dev
git push origin staging
```

#### 3. 生产环境部署
```bash
# 方式1: 推送到 main 分支
git checkout main
git merge staging
git push origin main

# 方式2: 创建版本标签
git tag v1.0.0
git push origin v1.0.0
```

### 手动部署

使用 GitHub Actions 手动触发：

1. 进入 GitHub 仓库
2. 点击 Actions → Manual Deploy to Docker Hub
3. 选择参数：
   - **Version**: `v1.0.0`
   - **Environment**: `production`
   - **Push latest**: `true`

## 🐳 Docker 运行配置

### 开发环境
```bash
docker run -d --name alert-service-dev \
  -e LARK_WEBHOOK="https://open.larksuite.com/open-apis/bot/v2/hook/dev-webhook" \
  -v $(pwd)/config/dev-alert.yml:/opt/xidp/conf/alert.yml:ro \
  your-username/alert-service:dev-latest
```

### 测试环境
```bash
docker run -d --name alert-service-staging \
  -e LARK_WEBHOOK="https://open.larksuite.com/open-apis/bot/v2/hook/staging-webhook" \
  -v $(pwd)/config/staging-alert.yml:/opt/xidp/conf/alert.yml:ro \
  your-username/alert-service:staging-latest
```

### 生产环境
```bash
docker run -d --name alert-service-prod \
  -e LARK_WEBHOOK="https://open.larksuite.com/open-apis/bot/v2/hook/prod-webhook" \
  -v $(pwd)/config/prod-alert.yml:/opt/xidp/conf/alert.yml:ro \
  your-username/alert-service:latest
```

## 🔧 Docker Compose 配置

### 多环境 docker-compose

创建不同环境的 docker-compose 文件：

#### `docker-compose.dev.yml`
```yaml
version: '3.8'
services:
  alert-service:
    image: your-username/alert-service:dev-latest
    environment:
      - ENV=development
    volumes:
      - ./config/dev-alert.yml:/opt/xidp/conf/alert.yml:ro
```

#### `docker-compose.staging.yml`
```yaml
version: '3.8'
services:
  alert-service:
    image: your-username/alert-service:staging-latest
    environment:
      - ENV=staging
    volumes:
      - ./config/staging-alert.yml:/opt/xidp/conf/alert.yml:ro
```

#### `docker-compose.prod.yml`
```yaml
version: '3.8'
services:
  alert-service:
    image: your-username/alert-service:latest
    environment:
      - ENV=production
    volumes:
      - ./config/prod-alert.yml:/opt/xidp/conf/alert.yml:ro
```

### 运行命令
```bash
# 开发环境
docker-compose -f docker-compose.dev.yml up -d

# 测试环境
docker-compose -f docker-compose.staging.yml up -d

# 生产环境
docker-compose -f docker-compose.prod.yml up -d
```

## 📊 监控和日志

### 查看容器状态
```bash
# 查看运行状态
docker ps | grep alert-service

# 查看日志
docker logs alert-service-prod

# 实时日志
docker logs -f alert-service-prod
```

### 健康检查
```bash
# 检查容器健康状态
docker inspect alert-service-prod | grep -A5 Health
```

## 🔒 安全最佳实践

1. **使用访问令牌**: 不要使用密码，使用 Docker Hub 访问令牌
2. **环境隔离**: 不同环境使用不同的 webhook 地址
3. **最小权限**: 容器内使用非 root 用户运行
4. **安全扫描**: 自动进行容器安全扫描
5. **配置分离**: 敏感配置通过环境变量或挂载文件提供

## 🚨 故障排除

### 常见问题

1. **Docker Hub 推送失败**
   - 检查 DOCKER_USERNAME 和 DOCKER_PASSWORD
   - 确认 Docker Hub 仓库存在且有推送权限

2. **容器启动失败**
   - 检查配置文件是否正确挂载
   - 查看容器日志 `docker logs container-name`

3. **飞书消息发送失败**
   - 验证 webhook URL 是否正确
   - 检查网络连接和防火墙设置

### 调试命令
```bash
# 进入容器调试
docker exec -it alert-service-prod /bin/sh

# 检查配置文件
docker exec alert-service-prod cat /opt/xidp/conf/alert.yml

# 测试网络连接
docker exec alert-service-prod ping -c 3 open.larksuite.com
``` 