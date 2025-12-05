# 📱 Android 端集成指南

本指南详细说明如何在 Android 设备上使用 Clipboard Bridge 服务。

## 🎯 三种集成方案

### 方案 1: HTTP Shortcuts（推荐 ⭐）
- ✅ **免费**，无需编程
- ✅ 简单易用，5 分钟配置
- ✅ 支持主屏幕小部件
- ✅ 可以通过分享菜单调用

### 方案 2: Tasker
- ✅ 功能强大，支持自动化
- ✅ 可以设置触发条件
- ❌ 付费应用（约 $3.49）

### 方案 3: 自定义 App
- ✅ 完全定制化
- ✅ 更好的用户体验
- ❌ 需要开发能力

---

## 方案 1: HTTP Shortcuts（详细教程）

### 📥 第一步：安装 App

在 Google Play 下载 [HTTP Shortcuts](https://play.google.com/store/apps/details?id=ch.rmy.android.http_shortcuts)

或从 [GitHub Releases](https://github.com/Waboodoo/HTTP-Shortcuts/releases) 下载 APK。

### ⚙️ 第二步：配置准备

在配置前，确认以下信息：

```
电脑 IP 地址: ____________ (例如: 192.168.1.100)
服务端口: 5678
Token: ____________ (如果设置了)
```

**如何查看电脑 IP？**
- Windows: 在电脑托盘图标中显示
- 或命令行运行: `ipconfig`（Windows）/ `ifconfig`（Linux/Mac）

### 🔽 第三步：创建"从电脑获取"快捷方式

1. 打开 HTTP Shortcuts App
2. 点击右下角 **"+"** 按钮
3. 选择 **"Regular Shortcut"**
4. 按以下配置：

#### 基本信息
```
Name: 从电脑获取
Description: 从电脑同步剪贴板到手机
```

#### 请求设置
```
URL: http://192.168.1.100:5678/pull
  ⚠️ 替换为你的电脑 IP

Method: GET

Headers: 点击 "Add Header"
  - Name: X-Auth-Token
  - Value: your-token
  ⚠️ 如果没设置 token，删除此头
```

#### 响应处理
```
Response Handling → Success:
  - 点击 "Add Action"
  - 选择 "Copy to Clipboard"
  - Source: Response Body
  
  - 点击 "Add Action" (再添加一个)
  - 选择 "Show Toast"
  - Message: ✓ 已从电脑同步
```

#### 外观
```
Icon: 选择一个下载图标
Color: 蓝色或你喜欢的颜色
```

5. 点击右上角 **"✓"** 保存

### 🔼 第四步：创建"发送到电脑"快捷方式

1. 再次点击 **"+"** 按钮
2. 选择 **"Regular Shortcut"**
3. 按以下配置：

#### 基本信息
```
Name: 发送到电脑
Description: 从手机发送剪贴板到电脑
```

#### 请求设置
```
URL: http://192.168.1.100:5678/push
  ⚠️ 替换为你的电脑 IP

Method: POST

Headers: 点击 "Add Header"
  - Name: X-Auth-Token
  - Value: your-token
  ⚠️ 如果没设置 token，删除此头
  
Request Body:
  - Content Type: application/x-www-form-urlencoded
  - 点击 "Add Parameter"
    - Key: text
    - Value: {clipboard}
    ⚠️ 注意是 {clipboard}，这是内置变量
```

#### 响应处理
```
Response Handling → Success:
  - 点击 "Add Action"
  - 选择 "Show Toast"
  - Message: ✓ 已发送到电脑
```

#### 外观
```
Icon: 选择一个上传图标
Color: 绿色或你喜欢的颜色
```

4. 点击右上角 **"✓"** 保存

### 🏠 第五步：添加到主屏幕

1. 在 HTTP Shortcuts 主界面，长按快捷方式
2. 选择 **"Place on Home Screen"**
3. 调整图标位置

现在你的主屏幕上有两个快捷图标了！

### 📤 第六步：添加到分享菜单（可选）

让你能从任何 App 直接分享文本到电脑：

1. 编辑"发送到电脑"快捷方式
2. 进入 **"Advanced Settings"**
3. 启用 **"Accept text via share"**
4. 保存

现在可以在任何 App 中：
```
选择文字 → 分享 → HTTP Shortcuts → 发送到电脑
```

### 🎯 使用方法

**从电脑获取剪贴板：**
1. 在电脑上复制内容
2. 在手机上点击"从电脑获取"图标
3. 内容已复制到手机剪贴板 ✓

**发送到电脑剪贴板：**
1. 在手机上复制内容
2. 点击"发送到电脑"图标
3. 内容已发送到电脑 ✓

### 🚀 进阶：创建智能同步快捷方式

创建一个快捷方式，自动判断同步方向：

#### 配置步骤

1. 创建新快捷方式 **"智能同步"**
2. URL: `http://192.168.1.100:5678/meta`
3. Method: GET
4. 添加 Token 头（如果需要）
5. 响应处理 → Success:
   ```
   - Extract from Response
     - Key: text
     - Variable: pc_clipboard
   
   - Show Dialog
     - Title: 选择同步方向
     - Message: 电脑: {pc_clipboard}\n手机: {clipboard}
     - Options:
       * 从电脑获取
       * 发送到电脑
       * 取消
   
   - If option = 从电脑获取:
       Copy to Clipboard: {pc_clipboard}
       Show Toast: ✓ 已从电脑同步
   
   - If option = 发送到电脑:
       Trigger Shortcut: 发送到电脑
   ```

---

## 方案 2: Tasker（付费但功能强大）

### 📥 安装

在 Google Play 购买 [Tasker](https://play.google.com/store/apps/details?id=net.dinglisch.android.taskerm)

### 🔽 创建"从电脑获取"任务

1. 打开 Tasker
2. Tasks 标签 → 点击 **"+"**
3. 命名: **Get from PC**
4. 添加 Action:

```
Net → HTTP Request
  - Server:Port: 192.168.1.100:5678
  - Path: /pull
  - Method: GET
  - Headers:
      X-Auth-Token: your-token
  - Output File: (留空)
  
System → Set Clipboard
  - Text: %http_data
  
Alert → Flash
  - Text: ✓ 已从电脑同步
```

### 🔼 创建"发送到电脑"任务

1. Tasks 标签 → 点击 **"+"**
2. 命名: **Send to PC**
3. 添加 Action:

```
Variables → Variable Set
  - Name: %clipboard_text
  - To: %CLIP
  
Net → HTTP Request
  - Server:Port: 192.168.1.100:5678
  - Path: /push
  - Method: POST
  - Headers:
      X-Auth-Token: your-token
  - Body: text=%clipboard_text
  - Content Type: application/x-www-form-urlencoded
  
Alert → Flash
  - Text: ✓ 已发送到电脑
```

### 🎯 创建快捷方式

1. 长按主屏幕
2. Widget → Tasker → Task Shortcut
3. 选择任务

### 🤖 自动化触发

Tasker 的优势在于可以设置自动触发：

**示例：复制内容时自动同步**
```
Profile: Clipboard Changed
  Event → System → Clipboard Changed

Task: Auto Sync
  - 等待 2 秒
  - 执行 Send to PC 任务
```

**示例：连接家庭 WiFi 时自动获取**
```
Profile: Home WiFi
  State → Net → Wifi Connected
  SSID: MyHomeWiFi

Task: Auto Pull
  - 等待 5 秒
  - 执行 Get from PC 任务
```

---

## 方案 3: 开发自定义 Android App

### 📱 简单实现示例（Kotlin）

创建一个简单的 Android App：

#### 1. 添加权限（AndroidManifest.xml）

```xml
<uses-permission android:name="android.permission.INTERNET" />
<uses-permission android:name="android.permission.ACCESS_NETWORK_STATE" />
```

#### 2. 创建 API 服务（ClipboardService.kt）

```kotlin
import okhttp3.*
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
import java.io.IOException

class ClipboardService(
    private val serverUrl: String,
    private val token: String?
) {
    private val client = OkHttpClient()

    // 从电脑获取剪贴板
    fun pull(callback: (String?, Exception?) -> Unit) {
        val request = Request.Builder()
            .url("$serverUrl/pull")
            .apply {
                token?.let { header("X-Auth-Token", it) }
            }
            .build()

        client.newCall(request).enqueue(object : Callback {
            override fun onFailure(call: Call, e: IOException) {
                callback(null, e)
            }

            override fun onResponse(call: Call, response: Response) {
                if (response.isSuccessful) {
                    callback(response.body?.string(), null)
                } else {
                    callback(null, Exception("HTTP ${response.code}"))
                }
            }
        })
    }

    // 发送到电脑剪贴板
    fun push(text: String, callback: (Boolean, Exception?) -> Unit) {
        val formBody = FormBody.Builder()
            .add("text", text)
            .build()

        val request = Request.Builder()
            .url("$serverUrl/push")
            .post(formBody)
            .apply {
                token?.let { header("X-Auth-Token", it) }
            }
            .build()

        client.newCall(request).enqueue(object : Callback {
            override fun onFailure(call: Call, e: IOException) {
                callback(false, e)
            }

            override fun onResponse(call: Call, response: Response) {
                callback(response.isSuccessful, null)
            }
        })
    }
}
```

#### 3. 创建主界面（MainActivity.kt）

```kotlin
import android.content.ClipData
import android.content.ClipboardManager
import android.os.Bundle
import android.widget.Toast
import androidx.appcompat.app.AppCompatActivity
import kotlinx.android.synthetic.main.activity_main.*

class MainActivity : AppCompatActivity() {
    private lateinit var clipboardService: ClipboardService
    private lateinit var clipboardManager: ClipboardManager

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        clipboardService = ClipboardService(
            serverUrl = "http://192.168.1.100:5678",
            token = "your-token" // 或 null
        )

        clipboardManager = getSystemService(CLIPBOARD_SERVICE) as ClipboardManager

        // 从电脑获取按钮
        btnPull.setOnClickListener {
            clipboardService.pull { text, error ->
                runOnUiThread {
                    if (error != null) {
                        Toast.makeText(this, "获取失败: ${error.message}", Toast.LENGTH_SHORT).show()
                    } else {
                        // 复制到剪贴板
                        val clip = ClipData.newPlainText("clipboard", text)
                        clipboardManager.setPrimaryClip(clip)
                        Toast.makeText(this, "✓ 已从电脑同步", Toast.LENGTH_SHORT).show()
                    }
                }
            }
        }

        // 发送到电脑按钮
        btnPush.setOnClickListener {
            val text = clipboardManager.primaryClip
                ?.getItemAt(0)?.text?.toString() ?: ""

            if (text.isEmpty()) {
                Toast.makeText(this, "剪贴板为空", Toast.LENGTH_SHORT).show()
                return@setOnClickListener
            }

            clipboardService.push(text) { success, error ->
                runOnUiThread {
                    if (success) {
                        Toast.makeText(this, "✓ 已发送到电脑", Toast.LENGTH_SHORT).show()
                    } else {
                        Toast.makeText(this, "发送失败: ${error?.message}", Toast.LENGTH_SHORT).show()
                    }
                }
            }
        }
    }
}
```

#### 4. 布局文件（activity_main.xml）

```xml
<?xml version="1.0" encoding="utf-8"?>
<LinearLayout xmlns:android="http://schemas.android.com/apk/res/android"
    android:layout_width="match_parent"
    android:layout_height="match_parent"
    android:orientation="vertical"
    android:padding="16dp"
    android:gravity="center">

    <Button
        android:id="@+id/btnPull"
        android:layout_width="match_parent"
        android:layout_height="wrap_content"
        android:text="⬇️ 从电脑获取"
        android:textSize="18sp"
        android:padding="16dp" />

    <Button
        android:id="@+id/btnPush"
        android:layout_width="match_parent"
        android:layout_height="wrap_content"
        android:layout_marginTop="16dp"
        android:text="⬆️ 发送到电脑"
        android:textSize="18sp"
        android:padding="16dp" />

</LinearLayout>
```

#### 5. 添加依赖（build.gradle）

```gradle
dependencies {
    implementation 'com.squareup.okhttp3:okhttp:4.10.0'
}
```

---

## 🔧 常见问题

### Q1: 连接失败

**检查清单：**
- [ ] 手机和电脑在同一 WiFi
- [ ] 电脑防火墙允许 5678 端口
- [ ] 使用正确的 IP 地址（电脑的局域网 IP）
- [ ] 电脑程序正在运行

### Q2: Token 验证失败

确保：
- HTTP Shortcuts 中的 Header 名称是 `X-Auth-Token`（区分大小写）
- Token 值与电脑 `config.json` 中一致

### Q3: 中文乱码

这个不应该发生，服务端使用 UTF-8。如果遇到，请提交 Issue。

### Q4: 自动化不工作（Tasker）

确保：
- Tasker 有后台运行权限
- 电池优化已关闭
- 相关权限已授予

---

## 💡 使用技巧

### 1. 快速访问

将快捷方式添加到：
- 主屏幕
- 通知栏（HTTP Shortcuts 支持）
- 侧边栏

### 2. 配置多个电脑

如果有多台电脑：
```
创建快捷方式:
- "从台式机获取" → http://192.168.1.100:5678/pull
- "从笔记本获取" → http://192.168.1.101:5678/pull
```

### 3. 安全建议

- ✅ 只在家庭/办公室 WiFi 使用
- ✅ 设置 Token 认证
- ✅ 不要在公共 WiFi 使用
- ✅ 敏感信息不要通过剪贴板传输

---

## 🎯 总结

| 方案 | 难度 | 成本 | 功能 | 推荐度 |
|------|------|------|------|--------|
| HTTP Shortcuts | ⭐ 简单 | 免费 | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| Tasker | ⭐⭐ 中等 | $3.49 | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ |
| 自定义 App | ⭐⭐⭐⭐ 困难 | 免费 | ⭐⭐⭐⭐⭐ | ⭐⭐⭐ |

**新手推荐**: HTTP Shortcuts - 5 分钟搞定！

---

**有问题？** 提交 [Issue](https://github.com/copypasteengine/clipboard-bridge/issues) 或查看 [Wiki](https://github.com/copypasteengine/clipboard-bridge/wiki)

