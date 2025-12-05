# 📋 Clipboard Bridge - 跨设备剪贴板同步服务

一个轻量级的桌面剪贴板 HTTP 服务，通过简单的 REST API 实现电脑与手机之间的剪贴板同步。支持 Windows、Linux 和 macOS 系统。

## ✨ 主要特性

- 🌐 **HTTP API** - 简单的 REST 接口，任何设备都能访问
- 🔄 **实时监听** - Windows 系统级监听，Linux/macOS 轮询监听
- 🔒 **Token 认证** - 可选的访问令牌保护
- 🚀 **开机自启** - Windows 支持自动配置
- 📊 **系统托盘** - 友好的托盘图标和菜单管理
- 📝 **日志记录** - 支持 debug/info/error 三级日志
- ⚡ **轻量高效** - CPU ~0.1%，内存 ~15MB

## 🖥️ 支持的平台

| 平台 | 架构 | 剪贴板监听 | 开机自启 |
|------|------|------------|----------|
| Windows | x64 | ✅ 系统级（实时） | ✅ 自动配置 |
| Linux | x64 | ⚡ 轮询（1秒） | 📝 手动配置 |
| macOS | Apple Silicon | ⚡ 轮询（1秒） | 📝 手动配置 |

> **注意**: macOS Intel 用户可使用 Rosetta 2 运行 ARM64 版本

## 📥 安装和使用

### 1. 下载程序

访问 [Releases 页面](https://github.com/copypasteengine/clipboard-bridge/releases) 下载对应平台的文件：

**桌面版本：**
- **Windows**: `clipboard-bridge-windows-amd64.zip`
- **Linux**: `clipboard-bridge-linux-amd64.tar.gz`
- **macOS**: `clipboard-bridge-macos-arm64.tar.gz`

**Android 版本：**
- **Android**: `clipboard-bridge-android-*-debug.apk`

### 2. 安装运行

**Windows:**
```powershell
# 1. 解压 zip 文件
# 2. 双击 clipboard-bridge.exe
# 3. 程序会在系统托盘显示图标
# 4. 服务自动启动在 5678 端口
```

**Linux / macOS:**
```bash
# 1. 解压
tar -xzf clipboard-bridge-*.tar.gz

# 2. 添加执行权限
chmod +x clipboard-bridge

# 3. 运行
./clipboard-bridge

# 4. (可选) 移动到系统路径
sudo mv clipboard-bridge /usr/local/bin/
```

### 3. 配置文件

程序首次运行会创建 `config.json`：

```json
{
  "port": 5678,
  "token": "",
  "auto_start": true,
  "auto_firewall": true,
  "log_level": "info"
}
```

**配置说明：**

| 配置项 | 类型 | 默认值 | 说明 |
|--------|------|--------|------|
| `port` | 整数 | 5678 | 服务监听端口（1024-65535） |
| `token` | 字符串 | "" | API 访问令牌，空则不验证 |
| `auto_start` | 布尔 | true | 是否开机自启 |
| `auto_firewall` | 布尔 | true | 是否自动配置防火墙（仅 Windows） |
| `log_level` | 字符串 | "info" | 日志级别：debug/info/error |

修改配置后需要重启程序。

## 🔌 API 接口

服务启动后，可通过 HTTP 接口访问剪贴板。

### 获取剪贴板内容

```http
GET http://电脑IP:5678/pull
X-Auth-Token: your-token
```

**响应：**
```
Hello World
```

### 设置剪贴板内容

**方式 1：表单提交**
```http
POST http://电脑IP:5678/push
X-Auth-Token: your-token
Content-Type: application/x-www-form-urlencoded

text=Hello World
```

**方式 2：直接提交**
```http
POST http://电脑IP:5678/push
X-Auth-Token: your-token

Hello World
```

**响应：**
```
OK
```

### 获取剪贴板元数据

```http
GET http://电脑IP:5678/meta
X-Auth-Token: your-token
```

**响应：**
```json
{
  "text": "Hello World",
  "updated": 1733400000
}
```

### 健康检查

```http
GET http://电脑IP:5678/ping
X-Auth-Token: your-token
```

**响应：**
```
PONG
```

### Token 认证

如果设置了 `token`，请求需要携带认证：

**方式 1：HTTP 头**
```http
X-Auth-Token: your-token
```

**方式 2：URL 参数**
```http
http://电脑IP:5678/pull?token=your-token
```

## 📱 手机端集成

### Android - 原生应用（推荐）

我们提供了完整的 Android 原生应用！

**功能特性：**
- ✅ Material Design 3 现代化 UI
- ✅ 智能同步（自动判断方向）
- ✅ 实时显示双端剪贴板内容
- ✅ 自动保存服务器配置
- ✅ 支持深色模式

**使用方法：**

1. **从源码编译**（推荐）
   ```bash
   cd android-app
   # 使用 Android Studio 打开，或命令行构建：
   ./gradlew assembleDebug
   ```
   详见 [Android App 文档](./android-app/README.md)

2. **使用 HTTP Shortcuts**
   
   无需编程，5 分钟配置完成。详见 [Android 集成指南](./ANDROID.md)

### iOS - 快捷指令

#### 基本配置

1. 打开 iOS"快捷指令" App
2. 创建新快捷指令
3. 添加以下动作：

**从电脑获取剪贴板：**
```
1. "获取 URL 的内容"
   - URL: http://你的电脑IP:5678/pull
   - 方法: GET
   - 添加标头: X-Auth-Token = your-token

2. "设定剪贴板"
   - 内容: [上一步的结果]

3. "显示通知"
   - 内容: ✓ 已从电脑同步
```

**发送到电脑剪贴板：**
```
1. "获取剪贴板"

2. "获取 URL 的内容"
   - URL: http://你的电脑IP:5678/push
   - 方法: POST
   - 请求体: 表单
   - 字段: text = [剪贴板内容]
   - 添加标头: X-Auth-Token = your-token

3. "显示通知"
   - 内容: ✓ 已发送到电脑
```

#### 智能同步快捷指令

创建一个更智能的版本，自动判断同步方向：

```
1. 获取电脑剪贴板（/meta 接口）
2. 获取 iOS 剪贴板
3. 比较两者内容：
   - 相同 → 提示"已同步"
   - iOS 为空 → 从电脑同步
   - 电脑为空 → 发送到电脑
   - 都有但不同 → 弹出菜单选择
```

详细配置步骤参见 [iOS 快捷指令配置指南](https://github.com/copypasteengine/clipboard-bridge/wiki)

### Android - HTTP Shortcuts 或 Tasker

#### 使用 HTTP Shortcuts App

1. 安装 [HTTP Shortcuts](https://play.google.com/store/apps/details?id=ch.rmy.android.http_shortcuts)
2. 创建快捷方式：

**获取剪贴板：**
```
Name: 从电脑获取
URL: http://你的电脑IP:5678/pull
Method: GET
Headers: X-Auth-Token: your-token
Post-execution Actions:
  - Copy Response to Clipboard
  - Show Toast: "✓ 已同步"
```

**发送剪贴板：**
```
Name: 发送到电脑
URL: http://你的电脑IP:5678/push
Method: POST
Headers: X-Auth-Token: your-token
Request Body: {clipboard}
Post-execution Actions:
  - Show Toast: "✓ 已发送"
```

#### 使用 Tasker

1. 安装 [Tasker](https://play.google.com/store/apps/details?id=net.dinglisch.android.taskerm)
2. 创建 Task → HTTP Request
3. 配置 URL、方法和标头
4. 添加到主屏幕或通过手势触发

### 自定义客户端

任何支持 HTTP 的工具都可以访问，例如：

**cURL (命令行):**
```bash
# 获取
curl -H "X-Auth-Token: your-token" http://192.168.1.100:5678/pull

# 设置
curl -X POST -H "X-Auth-Token: your-token" \
  -d "text=Hello from terminal" \
  http://192.168.1.100:5678/push
```

**Python 脚本:**
```python
import requests

headers = {'X-Auth-Token': 'your-token'}
url = 'http://192.168.1.100:5678'

# 获取
text = requests.get(f'{url}/pull', headers=headers).text

# 设置
requests.post(f'{url}/push', data={'text': 'Hello'}, headers=headers)
```

## ⚙️ 系统托盘菜单

右键点击托盘图标：

| 菜单项 | 说明 |
|--------|------|
| 📡 服务地址 | 显示外部访问地址 |
| 💻 本机地址 | 显示本机测试地址 |
| 🚀 开机自启 | 切换开机自启状态 |
| ▶️ 启动/停止服务 | 手动控制服务 |
| 📄 打开日志文件 | 查看运行日志 |
| ❌ 退出 | 退出程序 |

## 🔒 安全建议

1. **设置 Token** - 在 `config.json` 中设置 `token`，避免未授权访问
2. **局域网使用** - 建议仅在可信网络（家庭/办公室）使用
3. **防火墙配置**
   - Windows：程序会自动尝试添加规则
   - Linux：`sudo ufw allow 5678/tcp`
   - macOS：在系统设置中允许入站连接
4. **HTTPS** - 如需加密传输，建议使用 Nginx/Caddy 反向代理

## 📝 日志文件

日志位置：`clipboard_bridge.log`（程序同目录）

**日志级别：**
- `error` - 仅错误
- `info` - 关键操作和错误（默认）
- `debug` - 所有详细信息

**示例日志：**
```
[2024-12-05 10:30:15] [INFO] 程序启动
[2024-12-05 10:30:15] [INFO] 剪贴板监听已启动
[2024-12-05 10:30:15] [INFO] 🚀 剪贴板服务已启动
[2024-12-05 10:30:15] [INFO]    外部访问: http://192.168.1.100:5678
[2024-12-05 10:31:20] [INFO] 收到 Push 请求 (来自 192.168.1.200:54321)
[2024-12-05 10:31:20] [INFO] ✓ 成功写入剪贴板，内容长度: 15 字节
```

## 🛠️ 从源码编译

### 环境要求

- Go 1.20+
- GCC（Windows 需要 MinGW）

### 编译步骤

```bash
# 1. 克隆仓库
git clone https://github.com/copypasteengine/clipboard-bridge.git
cd clipboard-bridge

# 2. 安装依赖
go mod download

# 3. 编译

# Windows (无窗口)
go build -ldflags="-H windowsgui" -o clipboard-bridge.exe

# Linux
go build -o clipboard-bridge

# macOS
go build -o clipboard-bridge
```

### Linux 依赖

```bash
# Ubuntu/Debian - X11
sudo apt-get install xclip libgtk-3-dev

# Ubuntu/Debian - Wayland
sudo apt-get install wl-clipboard

# Fedora
sudo dnf install xclip gtk3-devel

# Arch
sudo pacman -S xclip gtk3
```

## 🐛 常见问题

### Q1: 手机无法连接

**A:** 检查：
- 手机和电脑在同一局域网
- 电脑防火墙允许 5678 端口
- 使用电脑的局域网 IP（不是 127.0.0.1）
- 程序正在运行（查看托盘图标）

### Q2: Token 验证失败

**A:** 确保：
- `config.json` 中的 `token` 与请求中的一致
- Token 区分大小写
- HTTP 头名称正确：`X-Auth-Token`

### Q3: iOS 快捷指令报错

**A:** 检查：
- URL 格式正确，包含 `http://`
- IP 地址正确（电脑的局域网 IP）
- Token 正确（如果设置了）
- 电脑服务正在运行

### Q4: Linux 剪贴板不工作

**A:** 安装依赖：
```bash
# X11
sudo apt-get install xclip

# Wayland
sudo apt-get install wl-clipboard
```

## 🔧 技术栈

- **语言**: Go 1.20
- **GUI**: getlantern/systray
- **剪贴板**: atotto/clipboard
- **Windows C API**: CGo
- **HTTP**: 标准库 net/http

## 📄 许可证

MIT License

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📧 联系方式

- **GitHub**: https://github.com/copypasteengine/clipboard-bridge
- **Issues**: https://github.com/copypasteengine/clipboard-bridge/issues

---

**享受跨设备剪贴板同步的便利！** 🎉
