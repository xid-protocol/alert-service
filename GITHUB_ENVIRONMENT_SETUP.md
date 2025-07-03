# GitHub 环境配置指南

本文档说明如何在 GitHub 仓库中配置 DockerHub 环境，用于安全地管理 Docker Hub 部署凭据。

## 🛠️ 环境设置步骤

### 1. 创建 DockerHub 环境

1. 进入你的 GitHub 仓库
2. 点击 **Settings** 标签
3. 在左侧菜单中找到 **Environments**
4. 点击 **New environment**
5. 输入环境名称：`DockerHub`
6. 点击 **Configure environment**

### 2. 配置环境 Secrets

在 DockerHub 环境中，你需要添加以下 secrets：

#### 必需的 Secrets

| Secret 名称 | 描述 | 示例值 |
|------------|------|--------|
| `DOCKER_USERNAME` | Docker Hub 用户名 | `your-docker-username` |
| `DOCKER_PASSWORD` | Docker Hub 密码或访问令牌 | `dckr_pat_xxx...` (推荐使用访问令牌) |

#### 添加 Secrets 步骤

1. 在 DockerHub 环境配置页面，找到 **Environment secrets**
2. 点击 **Add secret**
3. 添加第一个 secret：
   - **Name**: `DOCKER_USERNAME`
   - **Value**: 你的 Docker Hub 用户名
4. 再次点击 **Add secret**，添加第二个：
   - **Name**: `DOCKER_PASSWORD`
   - **Value**: 你的 Docker Hub 密码或访问令牌

### 3. 配置部署保护规则 (可选)

为了增加安全性，你可以配置部署保护规则：

#### 选项 1: 分支保护
```
Required reviewers: 1
Selected branches:
- main
- master
```

#### 选项 2: 时间延迟
```
Wait timer: 5 minutes
```

#### 选项 3: 管理员审批
```
Required reviewers: 
- repo-admin
- devops-team
```

## 🔐 Docker Hub 访问令牌创建

### 创建访问令牌（推荐）

强烈建议使用访问令牌而不是密码：

1. 登录 [Docker Hub](https://hub.docker.com/)
2. 点击右上角的用户头像 → **Account Settings**
3. 选择 **Security** 标签
4. 点击 **New Access Token**
5. 填写令牌信息：
   - **Access Token Description**: `GitHub Actions for alert-service`
   - **Permissions**: `Read, Write, Delete`
6. 点击 **Generate**
7. **重要**: 复制生成的令牌（只显示一次）
8. 将此令牌作为 `DOCKER_PASSWORD` 添加到 GitHub 环境中

### 令牌权限说明

| 权限 | 用途 | 是否必需 |
|------|------|----------|
| `Read` | 拉取镜像 | ✅ |
| `Write` | 推送镜像 | ✅ |
| `Delete` | 删除镜像 | ❌ |

## 📋 环境配置验证

### 检查配置是否正确

1. 确保环境名称为 `DockerHub`
2. 确保有 2 个 secrets：
   - `DOCKER_USERNAME`
   - `DOCKER_PASSWORD`
3. secrets 值没有多余的空格或特殊字符

### 测试部署

创建一个测试提交来验证配置：

```bash
# 创建一个小的更改来触发 workflow
echo "# Test deployment" >> README.md
git add README.md
git commit -m "test: trigger docker build"
git push origin main
```

然后检查 **Actions** 标签，确保：
1. Workflow 成功运行
2. 能够登录到 Docker Hub
3. 能够构建和推送镜像

## 🚨 常见问题

### 问题 1: 登录失败
```
Error: failed to login to registry docker.io
```

**解决方案**:
- 检查 `DOCKER_USERNAME` 是否正确
- 检查 `DOCKER_PASSWORD` 是否正确
- 确认访问令牌权限包含 `Write`
- 验证 Docker Hub 账户状态

### 问题 2: 推送权限被拒绝
```
Error: denied: requested access to the resource is denied
```

**解决方案**:
- 确认 Docker Hub 仓库存在
- 检查用户名是否与仓库所有者匹配
- 验证访问令牌权限

### 问题 3: 环境未找到
```
Environment 'DockerHub' not found
```

**解决方案**:
- 确认环境名称拼写正确（区分大小写）
- 确认环境已正确创建和保存

## 🔄 更新凭据

### 定期更新访问令牌

1. 在 Docker Hub 中撤销旧令牌
2. 创建新的访问令牌
3. 在 GitHub 环境中更新 `DOCKER_PASSWORD`

### 更新步骤

1. 进入仓库 **Settings** → **Environments** → **DockerHub**
2. 找到 `DOCKER_PASSWORD` secret
3. 点击 **Update**
4. 输入新的访问令牌值
5. 点击 **Update secret**

## 📊 监控和审计

### 查看部署历史

1. 进入 **Settings** → **Environments** → **DockerHub**
2. 查看 **Deployment history**
3. 检查部署状态和时间

### 部署通知

可以在环境设置中配置通知：
- Slack 通知
- 邮件通知
- Webhook 通知

## 🎯 最佳实践

1. **使用访问令牌**: 而不是密码
2. **定期轮换令牌**: 建议每 90 天更新一次
3. **最小权限原则**: 只给予必要的权限
4. **审计日志**: 定期检查部署历史
5. **分支保护**: 为生产部署设置审批流程
6. **备份凭据**: 安全地备份重要的访问令牌

---

配置完成后，你的 GitHub Actions 将能够安全地访问 Docker Hub 并自动部署你的应用！ 