# 🏗️ Clipboard Bridge - 架构设计

## 📊 系统架构

```
┌─────────────────────────────────────────────────────────────┐
│                    移动设备 (客户端)                         │
├─────────────────────────────────────────────────────────────┤
│  📱 Android App          │  📱 iOS 快捷指令                 │
│  - Material Design 3     │  - 智能同步逻辑                  │
│  - 原生 Kotlin          │  - 多网络支持                    │
│  - 双向同步             │  - Siri 集成                     │
└──────────────┬──────────────────────┬─────────────────────────┘
               │                      │
               │   HTTP API (REST)    │
               │                      │
┌──────────────┴──────────────────────┴─────────────────────────┐
│                   桌面服务 (HTTP Server)                       │
├─────────────────────────────────────────────────────────────┤
│  🌐 HTTP API Server                                          │
│  ├── GET  /pull  - 获取剪贴板                                │
│  ├── POST /push  - 设置剪贴板                                │
│  ├── GET  /meta  - 获取元数据                                │
│  └── GET  /ping  - 健康检查                                  │
│                                                               │
│  🔒 认证层 (可选)                                            │
│  └── X-Auth-Token 验证                                       │
│                                                               │
│  📋 剪贴板管理层                                             │
│  ├── Windows: 系统 API (C/CGo)                              │
│  └── Unix: 轮询机制 (Go)                                     │
└─────────────────────────────────────────────────────────────┘
```

## 🔄 数据流

### 场景 1: 从电脑获取剪贴板

```
手机 APP                桌面服务              系统剪贴板
   │                       │                       │
   │─────GET /pull────────>│                       │
   │                       │──────ReadAll()───────>│
   │                       │<──────"Hello"─────────│
   │<──────"Hello"─────────│                       │
   │                       │                       │
 复制到                                          
手机剪贴板                                        
```

### 场景 2: 发送到电脑剪贴板

```
手机 APP                桌面服务              系统剪贴板
   │                       │                       │
读取手机剪贴板                                    
   │                       │                       │
   │────POST /push────────>│                       │
   │  text="World"         │                       │
   │                       │─────WriteAll()───────>│
   │                       │                       │
   │<───────OK─────────────│                       │
```

### 场景 3: 智能同步

```
手机 APP                桌面服务
   │                       │
   │─────GET /meta────────>│
   │<───{text,updated}─────│
   │                       │
比较内容判断:
 - 相同 → 提示"已同步"
 - 手机空 → 执行 /pull
 - 电脑空 → 执行 /push
 - 都有且不同 → 提示选择
```

## 📂 代码架构

### 桌面服务 (Go)

```
main_common.go              # 跨平台通用代码
├── Config                 # 配置管理
├── HTTP Server            # Web 服务
│   ├── /pull             # 获取接口
│   ├── /push             # 设置接口
│   ├── /meta             # 元数据接口
│   └── /ping             # 健康检查
├── Logger                 # 日志系统
├── Systray UI             # 托盘界面
└── Repository             # 数据持久化

clipboard_windows.go        # Windows 特定实现
├── CGo Bindings           # C 语言接口
├── Clipboard Listener     # 系统级监听
├── Registry Management    # 开机自启
└── Firewall Rules         # 防火墙配置

clipboard_unix.go           # Unix 特定实现
├── Polling Listener       # 轮询监听
├── Desktop Integration    # 桌面环境集成
└── Platform Helpers       # 平台辅助函数

clipboard_windows.c         # Windows C 代码
└── WndProc                # 窗口消息处理
```

### Android App (Kotlin)

```
ClipboardBridgeApp         # Application 类
MainActivity               # 入口 Activity

ui/                        # 界面层 (Compose)
├── ClipboardBridgeScreen # 主界面
├── ClipboardViewModel    # 视图模型
├── SettingsDialog        # 设置对话框
└── theme/                # Material 主题

data/                      # 数据层
└── ClipboardRepository   # 数据仓库
    ├── pullFromPC()      # 从电脑获取
    ├── pushToPC()        # 发送到电脑
    ├── getMetadata()     # 获取元数据
    └── ping()            # 测试连接

network/                   # 网络层
└── ClipboardApiService   # HTTP 客户端
    ├── OkHttp Client     # 网络请求
    ├── Token Auth        # 认证处理
    └── Error Handling    # 错误处理
```

## 🔐 安全设计

### 认证机制

```
客户端请求
   │
   ├─ HTTP Header: X-Auth-Token
   │     或
   └─ URL 参数: ?token=xxx
          │
          ▼
    服务端验证
          │
   ┌──────┴──────┐
   │             │
 匹配          不匹配
   │             │
返回数据      401 未授权
```

### 网络安全

- ✅ **局域网限制** - 建议仅在可信网络使用
- ✅ **Token 认证** - 防止未授权访问
- ✅ **请求体限制** - 最大 10MB，防止滥用
- ✅ **超时控制** - 防止资源耗尽
- ⚠️ **HTTP 传输** - 明文传输，可用反向代理加密

## ⚡ 性能特性

### 桌面服务

| 指标 | Windows | Linux/macOS |
|------|---------|-------------|
| CPU 使用 | ~0% | ~0.1% |
| 内存占用 | ~15MB | ~12MB |
| 监听延迟 | <10ms | <1s |
| API 响应 | <50ms | <50ms |

### Android App

| 指标 | 数值 |
|------|------|
| APK 大小 | ~5MB (Debug) / ~3MB (Release) |
| 内存占用 | ~30MB |
| 启动时间 | <1s |
| API 调用 | <500ms (局域网) |

## 🔄 并发安全

### 桌面服务

使用读写锁保护共享资源：

```go
clipboardMu sync.RWMutex    // 保护剪贴板变量
cfgMu sync.RWMutex          // 保护配置变量

// 读取（允许并发）
clipboardMu.RLock()
text := lastClipboard
clipboardMu.RUnlock()

// 写入（独占访问）
clipboardMu.Lock()
lastClipboard = newText
clipboardMu.Unlock()
```

### Android App

使用 Kotlin Flow 和 Coroutines：

```kotlin
// StateFlow 确保 UI 线程安全
private val _uiState = MutableStateFlow(...)
val uiState: StateFlow<...> = _uiState.asStateFlow()

// 网络请求在 IO 线程
withContext(Dispatchers.IO) {
    apiService.pull(...)
}

// UI 更新在 Main 线程
_uiState.update { ... }
```

## 📦 依赖管理

### 桌面服务 (Go Modules)

```
github.com/getlantern/systray  # 系统托盘
github.com/atotto/clipboard     # 跨平台剪贴板
golang.org/x/sys/windows        # Windows API (仅 Windows)
```

### Android App (Gradle)

```
androidx.compose.*              # Jetpack Compose UI
com.squareup.okhttp3           # HTTP 客户端
androidx.lifecycle.*           # ViewModel & LiveData
androidx.datastore             # 配置持久化
```

## 🌐 跨平台策略

### 条件编译

**Go Build Tags:**
```go
//go:build windows
// 文件只在 Windows 编译

//go:build !windows
// 文件在非 Windows 平台编译
```

**文件名约定:**
```
clipboard_windows.go  → 仅 Windows
clipboard_windows.c   → 仅 Windows
clipboard_unix.go     → Linux/macOS
main_common.go        → 所有平台
```

### 平台差异封装

```go
// 接口定义（所有平台相同）
func initClipboardListener()
func openLogFile()
func updateAutostart()

// 实现分离（各平台不同）
// clipboard_windows.go: Windows 实现
// clipboard_unix.go: Unix 实现
```

## 🔮 扩展性设计

### 添加新 API

在 `main_common.go` 中：

```go
func handleNewAPI(w http.ResponseWriter, r *http.Request) {
    // 实现逻辑
}

// 注册路由
mux.HandleFunc("/new-api", handleNewAPI)
```

### 添加新平台

1. 创建 `clipboard_newplatform.go`
2. 实现必要的函数
3. Go 会自动根据 build tags 选择正确的实现

### 添加 Android 新功能

1. 在 `ClipboardApiService.kt` 添加 API 方法
2. 在 `ClipboardRepository.kt` 添加数据层方法
3. 在 `ClipboardViewModel.kt` 添加业务逻辑
4. 在 `ClipboardBridgeScreen.kt` 添加 UI

## 📈 性能优化

### 桌面服务

- ✅ **读写锁** - 最小化锁持有时间
- ✅ **连接池** - HTTP 服务器自动管理
- ✅ **优雅关闭** - 等待请求完成
- ✅ **资源清理** - 及时释放文件句柄

### Android App

- ✅ **协程** - 非阻塞异步操作
- ✅ **缓存策略** - 配置本地缓存
- ✅ **状态管理** - Flow 响应式编程
- ✅ **ProGuard** - Release 版本代码混淆

## 🧪 测试策略

### 单元测试

**桌面服务:**
```bash
go test ./...
```

**Android App:**
```bash
cd android-app
./gradlew test
```

### 集成测试

1. 启动桌面服务
2. 使用 curl 测试 API
3. 运行 Android App 测试同步

### 测试用例

- ✅ Token 认证
- ✅ 空剪贴板处理
- ✅ 大文本传输（>1MB）
- ✅ 特殊字符和 Unicode
- ✅ 并发请求

## 📄 许可证

MIT License

---

**架构设计遵循：简单、可靠、可扩展** 🎯

