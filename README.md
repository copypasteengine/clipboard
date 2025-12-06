# 📋 Clipboard Bridge - Cross-Device Clipboard Sync

[中文文档](./README.zh-CN.md) | English

A lightweight clipboard synchronization solution that enables seamless clipboard sharing between computers and mobile devices via HTTP API.

## 🎯 Use Cases

### ✅ Recommended Scenarios

**📱 Android + 🖥️ Windows/Linux/macOS**
- ✅ **Sync Android phone with any OS computer**
- ✅ Full-featured native Android App
- ✅ Modern Material Design 3 interface
- ✅ One-tap sync, smooth experience

**📱 iPhone + 🪟 Windows / 🐧 Linux**
- ✅ **Sync iOS with Windows/Linux computers**
- ✅ Implemented via iOS Shortcuts
- ✅ Siri voice control support
- ✅ Home screen widget support

### 💡 Not Recommended

**📱 iPhone + 🍎 Mac**
- ⚠️ Apple ecosystem has built-in Universal Clipboard
- ⚠️ iCloud auto-sync, better experience
- ⚠️ No third-party tools needed

> **Note**: macOS and iOS already have native clipboard sync via iCloud. This project primarily solves **cross-ecosystem** clipboard sync needs (Android ↔ Computer, iOS ↔ Windows/Linux).

## ✨ Key Features

- 🌐 **HTTP REST API** - Simple and accessible from any device
- 📱 **Native Android App** - Beautiful Material Design 3 UI
- 🌍 **Multi-Language** - English, 简体中文, 日本語
- 🔄 **Smart Sync** - Auto-detect sync direction
- 🔒 **Token Authentication** - Optional access token protection
- 📊 **System Tray** - User-friendly tray icon and menu
- ⚡ **Lightweight** - CPU <0.1%, Memory ~15MB
- 📝 **Complete Logging** - Easy troubleshooting

## 🖥️ Supported Platforms

### Desktop Service

| Platform | Architecture | Clipboard Monitoring | Auto-Start |
|----------|--------------|---------------------|------------|
| Windows | x64 | ✅ System-level (Real-time) | ✅ Auto-config |
| Linux | x64 | ⚡ Polling (1s interval) | 📝 Manual |
| macOS | Apple Silicon | ⚡ Polling (1s interval) | 📝 Manual |

### Mobile Clients

| Platform | Type | Features |
|----------|------|----------|
| Android | Native App (APK) | Smart sync, Config save, Real-time preview |
| iOS | Shortcuts | Smart sync, Siri control, Automation |

## 📥 Quick Start

### Step 1: Install Desktop Service

Visit [Releases](https://github.com/copypasteengine/clipboard-bridge/releases) to download:

**Windows:**
```powershell
# 1. Download clipboard-bridge-windows-amd64.zip
# 2. Extract to any folder
# 3. Double-click clipboard-bridge.exe
# 4. Check system tray icon ✓
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

### Step 2: Install Mobile Client

#### Android - Native App (Recommended ⭐)

1. **Download and Install APK**
   - Download `clipboard-bridge-android-v1.0.0-debug.apk`
   - Enable "Install unknown apps" in Settings
   - Install the APK

2. **Configure Server**
   - Open the App
   - Tap ⚙️ icon in top-right corner
   - Enter your PC IP: `http://192.168.1.100:5678`
   - Save

3. **Start Syncing**
   - Tap "Smart Sync" button
   - ✓ Done!

**App Features:**
- ✅ One-tap pull from PC
- ✅ One-tap push to PC
- ✅ Smart sync (auto-detect direction)
- ✅ Real-time clipboard preview on both sides
- ✅ Auto-save configuration
- ✅ Material Design 3 interface

#### iOS - Shortcuts

**Applicable:** iOS with Windows/Linux computers

> **Tip**: If you use iPhone + Mac, use Apple's Universal Clipboard instead (requires same iCloud account).

**Quick Setup - Pull from PC:**

1. Open "Shortcuts" App → Tap "+"
2. Add these actions:

```
Get Contents of URL
  URL: http://192.168.1.100:5678/pull
  Method: GET
  (If token is set, add header: X-Auth-Token)

Set Clipboard
  Content: [Get Contents of URL result]

Show Notification
  Content: ✓ Synced from PC
```

**Quick Setup - Push to PC:**

```
Get Clipboard

Get Contents of URL
  URL: http://192.168.1.100:5678/push
  Method: POST
  Request Body: Form
  Field: text = [Clipboard]
  (If token is set, add header: X-Auth-Token)

Show Notification
  Content: ✓ Sent to PC
```

See [iOS Shortcuts Configuration](#ios-shortcuts-configuration) at the end of this document for details.

## 🎯 Usage Examples

### Example 1: Open PC's Copied URL on Phone

```
1. Copy a URL in PC browser
2. Tap "Pull from PC" on phone App
3. Paste and open in phone browser
```

### Example 2: Send Phone's Verification Code to PC

```
1. Receive SMS code on phone, copy it
2. Tap "Push to PC"
3. Paste the code on PC
```

### Example 3: Cross-Device Content Editing

```
1. Find useful text on phone, copy it
2. Smart sync to PC
3. Paste into Word document on PC
```

## 🔌 API Reference

The desktop service provides REST API accessible from any client:

### Get Clipboard

```http
GET http://PC_IP:5678/pull
X-Auth-Token: your-token
```

**Response:** Clipboard text content

### Set Clipboard

```http
POST http://PC_IP:5678/push
X-Auth-Token: your-token
Content-Type: application/x-www-form-urlencoded

text=Hello World
```

**Response:** `OK`

### Get Metadata

```http
GET http://PC_IP:5678/meta
X-Auth-Token: your-token
```

**Response:**
```json
{
  "text": "Hello World",
  "updated": 1733400000
}
```

### Health Check

```http
GET http://PC_IP:5678/ping
X-Auth-Token: your-token
```

**Response:** `PONG`

## ⚙️ Configuration

First run creates `config.json`:

```json
{
  "port": 5678,
  "token": "",
  "auto_start": true,
  "auto_firewall": true,
  "log_level": "info"
}
```

| Option | Default | Description |
|--------|---------|-------------|
| `port` | 5678 | Service port (1024-65535) |
| `token` | "" | API token, empty to disable auth |
| `auto_start` | true | Auto-start on boot (Windows auto-config) |
| `auto_firewall` | true | Auto firewall rules (Windows only) |
| `log_level` | "info" | Log level: debug/info/error |

## 📊 System Tray Menu

Right-click the tray icon:

| Menu Item | Function |
|-----------|----------|
| 📡 Service Address | Show external access address |
| 💻 Local Address | Show localhost address |
| 🚀 Auto-Start | Toggle auto-start (Windows) |
| ▶️ Start/Stop Service | Manual control |
| 📄 Open Log File | View logs |
| ❌ Exit | Quit program |

## 🔒 Security

### Basic Security

1. **Set Token** - Configure `token` in `config.json` to prevent unauthorized access
2. **LAN Only** - Use only in trusted networks (home/office)
3. **Firewall**
   - Windows: Auto-configured by the program
   - Linux: `sudo ufw allow 5678/tcp`
   - macOS: System Settings → Firewall

### Advanced Security

For encrypted transmission, configure reverse proxy with Nginx/Caddy:

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

## 🐛 Troubleshooting

### Q1: Phone Cannot Connect

**Checklist:**
- [ ] Phone and PC on same WiFi
- [ ] PC firewall allows port 5678
- [ ] Using PC's LAN IP (e.g., 192.168.1.100, not 127.0.0.1)
- [ ] Desktop service is running (check tray icon)

**Test Method:**  
Visit `http://PC_IP:5678/ping` in phone browser. If it shows `PONG`, connection is OK.

### Q2: Token Verification Failed

- Ensure tokens match exactly (case-sensitive)
- Android: Enter token in Settings dialog
- iOS: Add HTTP header `X-Auth-Token` in Shortcuts

### Q3: Chinese Text Garbled

Service uses UTF-8 encoding, should work fine. Please submit an Issue if you encounter this.

### Q4: Linux Clipboard Not Working

Install dependencies:
```bash
# Ubuntu/Debian (X11)
sudo apt-get install xclip

# Ubuntu/Debian (Wayland)
sudo apt-get install wl-clipboard
```

## 🛠️ Build from Source

### Desktop Service

```bash
# Clone repository
git clone https://github.com/copypasteengine/clipboard-bridge.git
cd clipboard-bridge

# Install dependencies
go mod download

# Build
# Windows
go build -ldflags="-H windowsgui" -o clipboard-bridge.exe

# Linux/macOS
go build -o clipboard-bridge
```

**Linux Dependencies:**
```bash
sudo apt-get install xclip libgtk-3-dev  # Ubuntu/Debian
```

### Android App

```bash
cd android-app

# Open with Android Studio
# Or build via command line:
./gradlew assembleDebug

# APK location: app/build/outputs/apk/debug/app-debug.apk
```

See [android-app/BUILDING.md](./android-app/BUILDING.md) for details.

## 📝 Log Files

Log location: `clipboard_bridge.log` (same directory as executable)

**View Logs:**
- Windows: Tray menu → "📄 Open Log File"
- Linux/macOS: `tail -f clipboard_bridge.log`

**Example:**
```
[2024-12-05 10:30:15] [INFO] Program started
[2024-12-05 10:30:15] [INFO] 🚀 Clipboard service started
[2024-12-05 10:30:15] [INFO]    External: http://192.168.1.100:5678
[2024-12-05 10:31:20] [INFO] Push request from 192.168.1.200
[2024-12-05 10:31:20] [INFO] ✓ Clipboard updated, 15 bytes
```

## 🔧 Tech Stack

**Desktop Service:**
- Go 1.20 + CGo
- getlantern/systray (System tray)
- atotto/clipboard (Clipboard access)
- net/http (HTTP server)

**Android App:**
- Kotlin + Jetpack Compose
- Material Design 3
- OkHttp (HTTP client)
- Coroutines (Async)
- DataStore (Config storage)

## 📚 Documentation

- **[QUICKSTART.md](./QUICKSTART.md)** - 5-minute setup guide ⭐
- **[ANDROID.md](./ANDROID.md)** - Android HTTP Shortcuts guide
- **[android-app/README.md](./android-app/README.md)** - Android App manual
- **[android-app/BUILDING.md](./android-app/BUILDING.md)** - Android development guide
- **[ARCHITECTURE.md](./ARCHITECTURE.md)** - System architecture

## 📄 License

MIT License

## 🤝 Contributing

Issues and Pull Requests are welcome!

- **GitHub**: https://github.com/copypasteengine/clipboard-bridge
- **Issues**: https://github.com/copypasteengine/clipboard-bridge/issues
- **Releases**: https://github.com/copypasteengine/clipboard-bridge/releases

---

## 📱 iOS Shortcuts Configuration

### Basic Setup - Pull from PC

1. Open "Shortcuts" App
2. Tap "+" in top-right
3. Add these actions:

```
"Get Contents of URL"
  URL: http://192.168.1.100:5678/pull  ← Replace with your PC IP
  Method: GET
  
  If token is set, tap "Show More" → Enable "Headers"
    Add header: X-Auth-Token = your-token

"Set Clipboard"
  Content: [Get Contents of URL]

"Show Notification"
  Content: ✓ Synced from PC
```

### Basic Setup - Push to PC

```
"Get Clipboard"

"Get Contents of URL"
  URL: http://192.168.1.100:5678/push  ← Replace with your PC IP
  Method: POST
  Request Body: Form
  Add field: text = [Clipboard]
  
  If token is set:
    Add header: X-Auth-Token = your-token

"Show Notification"
  Content: ✓ Sent to PC
```

### Advanced - Smart Sync

Create an intelligent shortcut that auto-detects sync direction:

1. Get PC clipboard (`/meta` endpoint)
2. Get iPhone clipboard
3. Compare and auto-handle:
   - Same → Show "Already synced"
   - iPhone empty → Pull from PC
   - PC empty → Push to PC
   - Both different → Show selection menu

### Usage Tips

**Add to Home Screen:**
- Long-press home screen → Widgets → Shortcuts
- Select your clipboard sync shortcut

**Siri Voice:**
- "Hey Siri, run clipboard sync"

**Automation:**
- Auto-run when connecting to specific WiFi
- Auto-run when opening specific apps

---

**Enjoy seamless cross-device clipboard sync!** 🎉

**Primary Use Cases:**
- ✅ **Android ↔ Windows/Linux/macOS** - Use Android App
- ✅ **iOS ↔ Windows/Linux** - Use Shortcuts
- ⚠️ **iOS ↔ macOS** - Use Apple Universal Clipboard instead

