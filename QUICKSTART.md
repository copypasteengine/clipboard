# 🚀 快速开始指南

## 📱 完整的剪贴板同步解决方案

5 分钟内在你的设备之间实现剪贴板同步！

## 🎯 场景示例

### 场景 1: 在手机上打开电脑复制的链接

```
1. 在电脑浏览器复制一个网址
2. 在手机上运行同步
3. 在手机浏览器粘贴并打开
```

### 场景 2: 将手机照片的分享链接发送到电脑

```
1. 在手机相册分享照片，复制链接
2. 运行同步到电脑
3. 在电脑上粘贴链接
```

### 场景 3: 电脑上的验证码发送到手机

```
1. 电脑收到短信验证码，复制
2. 手机同步获取
3. 粘贴到手机 App
```

---

## 第一步：安装桌面服务 (3 分钟)

### Windows 用户

1. 访问 [Releases](https://github.com/copypasteengine/clipboard-bridge/releases)
2. 下载 `clipboard-bridge-windows-amd64.zip`
3. 解压到任意文件夹（例如 `C:\clipboard-bridge\`）
4. 双击运行 `clipboard-bridge.exe`
5. 程序会在系统托盘显示图标 📋

**✓ 完成！** 服务已启动，默认端口 5678

### Linux 用户

```bash
# 1. 下载
wget https://github.com/copypasteengine/clipboard-bridge/releases/latest/download/clipboard-bridge-linux-amd64.tar.gz

# 2. 解压
tar -xzf clipboard-bridge-linux-amd64.tar.gz

# 3. 安装依赖
sudo apt-get install xclip  # Ubuntu/Debian

# 4. 运行
chmod +x clipboard-bridge
./clipboard-bridge
```

### macOS 用户

```bash
# 1. 下载
curl -LO https://github.com/copypasteengine/clipboard-bridge/releases/latest/download/clipboard-bridge-macos-arm64.tar.gz

# 2. 解压
tar -xzf clipboard-bridge-macos-arm64.tar.gz

# 3. 运行
chmod +x clipboard-bridge
./clipboard-bridge
```

---

## 第二步：获取电脑 IP 地址 (1 分钟)

### 方法 1: 查看托盘提示

鼠标悬停在系统托盘图标上，会显示：
```
剪贴板同步服务 - 192.168.1.100:5678
```

### 方法 2: 查看托盘菜单

右键点击托盘图标，第一行显示：
```
📡 服务地址: http://192.168.1.100:5678
```

### 方法 3: 命令行查看

**Windows:**
```cmd
ipconfig
# 查找 "IPv4 地址"
```

**Linux/macOS:**
```bash
ip addr  # 或 ifconfig
# 查找 inet 192.168.x.x
```

**记下你的 IP 地址！** 例如：`192.168.1.100`

---

## 第三步：配置手机端 (2 分钟)

### Android 用户 - 原生 App（最简单）

1. **下载安装**
   - 访问 [Releases](https://github.com/copypasteengine/clipboard-bridge/releases)
   - 下载 `clipboard-bridge-android-*-debug.apk`
   - 在手机上安装（需允许未知来源）

2. **配置服务器**
   - 打开 App
   - 点击右上角 ⚙️ 图标
   - 输入：`http://192.168.1.100:5678`（替换为你的 IP）
   - 如果设置了 Token，输入 Token
   - 点击"保存"

3. **开始同步**
   - 点击"智能同步"按钮
   - 完成！✓

### iOS 用户 - 快捷指令

1. **创建快捷指令**
   - 打开"快捷指令" App
   - 点击右上角 "+"
   - 添加以下动作：

2. **从电脑获取**
   ```
   获取 URL 的内容
     URL: http://192.168.1.100:5678/pull
     方法: GET
   
   设定剪贴板
     内容: [上一步结果]
   
   显示通知
     内容: ✓ 已同步
   ```

3. **发送到电脑**
   ```
   获取剪贴板
   
   获取 URL 的内容
     URL: http://192.168.1.100:5678/push
     方法: POST
     请求体: 表单
     字段: text = [剪贴板]
   
   显示通知
     内容: ✓ 已发送
   ```

详细配置参见 [README.md](./README.md)

---

## 🎉 开始使用！

### 测试同步

**从电脑到手机：**
1. 在电脑上复制文本（Ctrl+C）
2. 在手机上运行同步
3. 在手机上粘贴（长按 → 粘贴）
4. ✓ 成功！

**从手机到电脑：**
1. 在手机上复制文本
2. 运行同步到电脑
3. 在电脑上粘贴（Ctrl+V）
4. ✓ 成功！

---

## 🔒 可选：设置 Token 保护

为了安全，建议设置 Token：

1. **修改配置文件**
   
   编辑电脑上的 `config.json`：
   ```json
   {
     "port": 5678,
     "token": "my-secret-123",
     "auto_start": true,
     "auto_firewall": true,
     "log_level": "info"
   }
   ```

2. **重启服务**
   - 右键托盘图标 → 退出
   - 重新运行程序

3. **在手机上配置 Token**
   - Android: 在设置中输入 `my-secret-123`
   - iOS: 在快捷指令中添加 Token 到请求头

---

## 💡 使用技巧

### 技巧 1: 添加到主屏幕

**Android:**
- 长按主屏幕 → 小部件 → HTTP Shortcuts（如果用此方案）
- 或直接使用原生 App

**iOS:**
- 长按主屏幕 → 添加小部件 → 快捷指令
- 选择你的剪贴板同步快捷指令

### 技巧 2: 语音控制（iOS）

对 Siri 说：
```
"嘿 Siri，运行剪贴板同步"
```

### 技巧 3: 开机自启（电脑）

**Windows:**
- 右键托盘图标
- 勾选"🚀 开机自启"
- 下次开机会自动启动

**Linux/macOS:**
- 需要手动配置系统启动项
- 参见 README.md

### 技巧 4: 查看日志

**电脑端:**
- 右键托盘图标 → "📄 打开日志文件"
- 查看所有同步记录

**Android:**
- 暂无日志功能（可在未来版本添加）

---

## 🐛 遇到问题？

### 问题：手机无法连接

**检查清单：**
- [ ] 手机和电脑在同一 WiFi 网络
- [ ] 电脑防火墙允许 5678 端口
- [ ] IP 地址正确（不是 127.0.0.1）
- [ ] 桌面服务正在运行（查看托盘图标）

**测试方法：**

在手机浏览器访问：
```
http://你的电脑IP:5678/ping
```

如果显示 `PONG`，说明连接正常！

### 问题：Token 验证失败

确保：
- Token 完全一致（区分大小写）
- 没有多余的空格
- HTTP 头名称是 `X-Auth-Token`

### 问题：同步后内容不对

检查：
- 确认复制的是最新内容
- 查看电脑日志文件
- 尝试手动刷新

---

## 📚 更多文档

- [README.md](./README.md) - 完整功能文档
- [ANDROID.md](./ANDROID.md) - Android 详细集成指南
- [ARCHITECTURE.md](./ARCHITECTURE.md) - 架构设计文档
- [android-app/README.md](./android-app/README.md) - Android App 说明

---

## 🎊 恭喜！

现在你已经拥有了一个**完整的跨设备剪贴板同步系统**！

**支持的场景：**
- ✅ 电脑 ↔ Android
- ✅ 电脑 ↔ iOS
- ✅ Windows / Linux / macOS 任意组合

**享受无缝的跨设备体验！** 🎉

---

**需要帮助？** 提交 [Issue](https://github.com/copypasteengine/clipboard-bridge/issues)

