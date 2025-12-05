# 📋 Clipboard Bridge - 跨设备剪贴板同步

一个轻量级的剪贴板同步解决方案，通过 HTTP API 实现电脑与手机之间的无缝剪贴板共享。

## 🎯 适用场景

### ✅ 推荐使用场景

**📱 Android + 🖥️ Windows/Linux/macOS**
- ✅ **Android 手机与任意系统电脑同步**
- ✅ 提供完整的原生 Android App
- ✅ Material Design 3 现代化界面
- ✅ 一键同步，体验流畅

**📱 iPhone + 🪟 Windows / 🐧 Linux**  
- ✅ **iOS 与 Windows/Linux 电脑同步**
- ✅ 通过 iOS 快捷指令实现
- ✅ 支持 Siri 语音控制
- ✅ 可添加到主屏幕小组件

### 💡 不推荐场景

**📱 iPhone + 🍎 Mac**
- ⚠️ Apple 生态内建 Universal Clipboard（通用剪贴板）
- ⚠️ iCloud 自动同步，体验更好
- ⚠️ 无需第三方工具

> **说明**: macOS 和 iOS 之间已有苹果原生的剪贴板同步功能，建议直接使用系统功能。本项目主要解决**跨生态**（Android ↔ 电脑，iOS ↔ Windows/Linux）的剪贴板同步需求。

## ✨ 主要特性

- 🌐 **HTTP REST API** - 简单易用，任何设备都能访问
- 📱 **Android 原生 App** - Material Design 3 精美界面
- 🔄 **智能同步** - 自动判断同步方向
- 🔒 **Token 认证** - 可选的访问令牌保护
- 📊 **系统托盘** - 友好的托盘图标和菜单
- ⚡ **轻量高效** - CPU <0.1%，内存 ~15MB
- 📝 **完整日志** - 便于故障排查

## 🖥️ 支持的平台

### 桌面服务

| 平台 | 架构 | 剪贴板监听 | 开机自启 |
|------|------|------------|----------|
| Windows | x64 | ✅ 系统级（实时） | ✅ 自动配置 |
| Linux | x64 | ⚡ 轮询（1秒） | 📝 手动配置 |
| macOS | Apple Silicon | ⚡ 轮询（1秒） | 📝 手动配置 |

### 移动客户端

| 平台 | 类型 | 功能 |
|------|------|------|
| Android | 原生 App (APK) | 智能同步、配置保存、实时预览 |
| iOS | 快捷指令 | 智能同步、Siri 控制、自动化 |

## 📥 快速开始

### 第一步：安装桌面服务

访问 [Releases 页面](https://github.com/copypasteengine/clipboard-bridge/releases) 下载：

**Windows:**
```powershell
# 1. 下载 clipboard-bridge-windows-amd64.zip
# 2. 解压到任意文件夹
# 3. 双击 clipboard-bridge.exe
# 4. 查看系统托盘图标 ✓
```

**Linux:**
```bash
wget https://github.com/copypasteengine/clipboard-bridge/releases/latest/download/clipboard-bridge-linux-amd64.tar.gz
tar -xzf clipboard-bridge-linux-amd64.tar.gz
chmod +x clipboard-bridge
./clipboard-bridge
```

**macOS:**
```bash
curl -LO https://github.com/copypasteengine/clipboard-bridge/releases/latest/download/clipboard-bridge-macos-arm64.tar.gz
tar -xzf clipboard-bridge-macos-arm64.tar.gz
chmod +x clipboard-bridge
./clipboard-bridge
```

### 第二步：安装手机客户端

#### Android - 原生 App（推荐 ⭐）

1. **下载安装 APK**
   - 下载 `clipboard-bridge-android-v1.0.0-debug.apk`
   - 在设置中允许安装未知应用
   - 安装 APK

2. **配置服务器**
   - 打开 App
   - 点击右上角 ⚙️ 图标
   - 输入电脑 IP：`http://192.168.1.100:5678`
   - 保存

3. **开始同步**
   - 点击"智能同步"按钮
   - ✓ 完成！

**App 功能：**
- ✅ 一键获取电脑剪贴板
- ✅ 一键发送到电脑
- ✅ 智能同步（自动判断方向）
- ✅ 实时显示两端剪贴板内容
- ✅ 自动保存配置
- ✅ Material Design 3 界面

#### iOS - 快捷指令

**基础配置（从电脑获取）:**
```
1. 打开"快捷指令" App
2. 创建新快捷指令
3. 添加动作：
   - "获取 URL 的内容"
     URL: http://你的电脑IP:5678/pull
     方法: GET
   - "设定剪贴板"
   - "显示通知": ✓ 已同步
```

详细配置参见文档末尾的 [iOS 快捷指令配置](#ios---快捷指令配置)

## 🎯 使用场景示例

### 场景 1: 在手机上打开电脑复制的网址

```
1. 在电脑浏览器复制网址
2. 在手机 App 点击"从电脑获取"
3. 在手机浏览器粘贴打开
```

### 场景 2: 将手机上的验证码发送到电脑

```
1. 手机收到短信验证码，复制
2. 点击"发送到电脑"
3. 在电脑上粘贴验证码
```

### 场景 3: 编辑文档时跨设备复制内容

```
1. 在手机上看到有用的文字，复制
2. 智能同步到电脑
3. 在电脑 Word 文档中粘贴
```

## 🔌 API 接口

桌面服务提供 REST API，可以被任何客户端访问：

### 获取剪贴板

```http
GET http://电脑IP:5678/pull
X-Auth-Token: your-token
```

**响应：** 剪贴板文本内容

### 设置剪贴板

```http
POST http://电脑IP:5678/push
X-Auth-Token: your-token
Content-Type: application/x-www-form-urlencoded

text=Hello World
```

**响应：** `OK`

### 获取元数据

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

**响应：** `PONG`

## ⚙️ 配置文件

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

| 配置项 | 默认值 | 说明 |
|--------|--------|------|
| `port` | 5678 | 服务端口（1024-65535） |
| `token` | "" | API 令牌，空则不验证 |
| `auto_start` | true | 开机自启（Windows 自动配置） |
| `auto_firewall` | true | 自动防火墙规则（仅 Windows） |
| `log_level` | "info" | 日志级别：debug/info/error |

## 📱 移动端详细说明

### Android App（推荐使用 ⭐⭐⭐）

**下载安装：**
1. 从 [Releases](https://github.com/copypasteengine/clipboard-bridge/releases) 下载 APK
2. 在手机上安装（需允许未知来源）
3. 打开 App 并配置服务器地址

**功能特性：**
- 🎨 **Material Design 3** - 现代化界面
- 🔄 **智能同步** - 自动判断同步方向
- 👀 **实时预览** - 同时显示手机和电脑剪贴板
- 💾 **配置保存** - 自动记住服务器设置
- 🌓 **深色模式** - 跟随系统主题
- ⚡ **一键操作** - 简单快捷

**界面预览：**
```
┌──────────────────────────┐
│ Clipboard Bridge     ⚙️  │
├──────────────────────────┤
│ ● 已连接              🔄 │
│ http://192.168.1.100:5678│
├──────────────────────────┤
│ 📱 手机剪贴板         🔄 │
│ Hello from phone         │
├──────────────────────────┤
│ 💻 电脑剪贴板         🔄 │
│ Hello from PC            │
│ 更新: 2分钟前            │
├──────────────────────────┤
│  [ ⬇️  从电脑获取 ]      │
│  [ ⬆️  发送到电脑 ]      │
│  [ 🔄  智能同步 ]        │
└──────────────────────────┘
```

**使用步骤：**
1. 复制内容到手机或电脑
2. 打开 App，点击对应按钮
3. 内容已同步 ✓

详细说明：[android-app/README.md](./android-app/README.md)

### iOS 快捷指令配置

**适用场景：** iOS 与 Windows/Linux 电脑同步

> **提示**: 如果你使用 iPhone + Mac 组合，建议直接使用 Apple 的通用剪贴板功能（需登录同一 iCloud 账号），无需本工具。

**快速配置 - 从电脑获取：**

1. 打开"快捷指令" App → 点击 "+"
2. 添加以下动作：

```
获取 URL 的内容
  URL: http://192.168.1.100:5678/pull
  方法: GET
  (如果设置了 Token，添加标头: X-Auth-Token)

设定剪贴板
  内容: [获取 URL 的内容结果]

显示通知
  内容: ✓ 已从电脑同步
```

**快速配置 - 发送到电脑：**

```
获取剪贴板

获取 URL 的内容
  URL: http://192.168.1.100:5678/push
  方法: POST
  请求体: 表单
  字段: text = [剪贴板]
  (如果设置了 Token，添加标头: X-Auth-Token)

显示通知
  内容: ✓ 已发送到电脑
```

**使用方法：**
- 直接运行快捷指令
- 对 Siri 说"运行剪贴板同步"
- 添加到主屏幕小组件
- 设置自动化触发

## 📊 系统托盘菜单

右键点击托盘图标：

| 菜单项 | 功能 |
|--------|------|
| 📡 服务地址 | 显示外部访问地址（如 `http://192.168.1.100:5678`） |
| 💻 本机地址 | 显示本机测试地址 |
| 🚀 开机自启 | 切换开机自启状态（Windows） |
| ▶️ 启动/停止服务 | 手动控制服务 |
| 📄 打开日志文件 | 查看运行日志 |
| ❌ 退出 | 退出程序 |

## 🔒 安全建议

### 基础安全

1. **设置 Token** - 在 `config.json` 中设置 `token`，防止未授权访问
2. **局域网使用** - 仅在家庭/办公室可信网络使用
3. **防火墙配置**
   - Windows: 程序自动添加规则
   - Linux: `sudo ufw allow 5678/tcp`
   - macOS: 系统设置 → 防火墙

### 高级安全

如需加密传输，可配置 Nginx/Caddy 反向代理：

```nginx
server {
    listen 443 ssl;
    ssl_certificate /path/to/cert.pem;
    ssl_certificate_key /path/to/key.pem;
    
    location / {
        proxy_pass http://localhost:5678;
    }
}
```

## 🐛 常见问题

### Q1: 手机无法连接

**检查清单：**
- [ ] 手机和电脑在同一 WiFi 网络
- [ ] 电脑防火墙允许 5678 端口
- [ ] 使用电脑的局域网 IP（如 192.168.1.100）
- [ ] 桌面服务正在运行（查看托盘图标）

**测试方法：**
在手机浏览器访问 `http://电脑IP:5678/ping`，如果显示 `PONG`，说明连接正常。

### Q2: Token 验证失败

- 确保 Token 完全一致（区分大小写）
- Android: 在设置对话框中输入 Token
- iOS: 在快捷指令中添加 HTTP 头 `X-Auth-Token`

### Q3: 中文乱码

服务端使用 UTF-8 编码，不应该有乱码。如遇到请提交 Issue。

### Q4: Linux 剪贴板不工作

安装依赖：
```bash
# Ubuntu/Debian (X11)
sudo apt-get install xclip

# Ubuntu/Debian (Wayland)
sudo apt-get install wl-clipboard
```

## 🛠️ 从源码编译

### 桌面服务

```bash
# 克隆仓库
git clone https://github.com/copypasteengine/clipboard-bridge.git
cd clipboard-bridge

# 安装依赖
go mod download

# 编译
# Windows
go build -ldflags="-H windowsgui" -o clipboard-bridge.exe

# Linux/macOS
go build -o clipboard-bridge
```

**Linux 额外依赖：**
```bash
sudo apt-get install xclip libgtk-3-dev  # Ubuntu/Debian
```

### Android App

```bash
cd android-app

# 使用 Android Studio 打开项目
# 或命令行构建：
./gradlew assembleDebug

# APK 位置：app/build/outputs/apk/debug/app-debug.apk
```

详细构建说明：[android-app/BUILDING.md](./android-app/BUILDING.md)

## 📝 日志文件

日志位置：`clipboard_bridge.log`（程序同目录）

**查看日志：**
- Windows: 托盘菜单 → "📄 打开日志文件"
- Linux/macOS: `tail -f clipboard_bridge.log`

**示例日志：**
```
[2024-12-05 10:30:15] [INFO] 程序启动
[2024-12-05 10:30:15] [INFO] 🚀 剪贴板服务已启动
[2024-12-05 10:30:15] [INFO]    外部访问: http://192.168.1.100:5678
[2024-12-05 10:31:20] [INFO] 收到 Push 请求 (来自 192.168.1.200)
[2024-12-05 10:31:20] [INFO] ✓ 成功写入剪贴板，15 字节
```

## 🔧 技术栈

**桌面服务：**
- Go 1.20 + CGo
- getlantern/systray（系统托盘）
- atotto/clipboard（剪贴板）
- net/http（HTTP 服务器）

**Android App：**
- Kotlin + Jetpack Compose
- Material Design 3
- OkHttp（HTTP 客户端）
- Coroutines（异步）
- DataStore（配置）

## 📚 相关文档

- **[QUICKSTART.md](./QUICKSTART.md)** - 5 分钟快速上手指南 ⭐
- **[ANDROID.md](./ANDROID.md)** - Android HTTP Shortcuts 详细配置
- **[android-app/README.md](./android-app/README.md)** - Android App 使用说明
- **[android-app/BUILDING.md](./android-app/BUILDING.md)** - Android App 开发指南
- **[ARCHITECTURE.md](./ARCHITECTURE.md)** - 系统架构设计文档

## 📄 许可证

MIT License

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

- **GitHub**: https://github.com/copypasteengine/clipboard-bridge
- **Issues**: https://github.com/copypasteengine/clipboard-bridge/issues
- **Releases**: https://github.com/copypasteengine/clipboard-bridge/releases

---

## 📱 iOS 快捷指令配置

### 基本配置 - 从电脑获取

1. 打开"快捷指令" App
2. 点击右上角 "+"
3. 添加以下动作：

```
"获取 URL 的内容"
  URL: http://192.168.1.100:5678/pull  ← 改为你的电脑 IP
  方法: GET
  
  如果设置了 Token，点击"显示更多" → 打开"标头"
    添加标头: X-Auth-Token = your-token

"设定剪贴板"
  内容: [获取 URL 的内容]

"显示通知"
  内容: ✓ 已从电脑同步
```

### 基本配置 - 发送到电脑

```
"获取剪贴板"

"获取 URL 的内容"
  URL: http://192.168.1.100:5678/push  ← 改为你的电脑 IP
  方法: POST
  请求体: 表单
  添加字段: text = [剪贴板]
  
  如果设置了 Token:
    添加标头: X-Auth-Token = your-token

"显示通知"
  内容: ✓ 已发送到电脑
```

### 进阶配置 - 智能同步

创建一个能自动判断方向的快捷指令：

1. 获取电脑剪贴板（`/meta` 接口）
2. 获取 iPhone 剪贴板
3. 比较内容自动处理：
   - 相同 → 提示"已同步"
   - iPhone 为空 → 从电脑同步
   - 电脑为空 → 发送到电脑
   - 都有但不同 → 弹出选择菜单

### 使用技巧

**添加到主屏幕：**
- 长按主屏幕 → 小组件 → 快捷指令
- 选择你的剪贴板同步快捷指令

**Siri 语音：**
- "嘿 Siri，运行剪贴板同步"

**自动化触发：**
- 连接特定 WiFi 时自动运行
- 打开特定 App 时自动运行

---

**享受跨设备剪贴板同步的便利！** 🎉

**主要使用场景：**
- ✅ **Android ↔ Windows/Linux/macOS** - 使用 Android App
- ✅ **iOS ↔ Windows/Linux** - 使用快捷指令
- ⚠️ **iOS ↔ macOS** - 建议使用 Apple 通用剪贴板
