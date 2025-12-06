package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

// 支持的语言
const (
	LangEN = "en" // English
	LangZH = "zh" // 中文
	LangJA = "ja" // 日本語
)

var (
	currentLang = LangEN
	i18nMu      sync.RWMutex
)

// 翻译字典
var translations = map[string]map[string]string{
	// 英文
	LangEN: {
		"app_title":            "Clipboard Bridge",
		"service_address":      "📡 Service Address: %s",
		"local_address":        "💻 Local Address: %s",
		"auto_start":           "🚀 Auto-Start",
		"start_service":        "▶️  Start Service",
		"stop_service":         "⏸️  Stop Service",
		"open_log":             "📄 Open Log File",
		"quit":                 "❌ Quit",
		"ext_access":           "External access via this address",
		"local_test":           "For local testing",
		
		// 日志消息
		"log_separator":        "========================================",
		"program_start":        "Program started, log file: %s",
		"config_loaded":        "Config loaded: Port=%d, Token=%s, AutoStart=%v, LogLevel=%s",
		"clipboard_listener":   "Clipboard listener started",
		"service_started":      "🚀 Clipboard service started",
		"external_access":      "   External: http://%s:%d",
		"local_access":         "   本机访问: http://localhost:%d",
		"service_stopped":      "Clipboard service stopped",
		"clipboard_updated":    "Local clipboard updated, length: %d bytes",
		"clipboard_preview":    "Clipboard preview: %s",
		"push_request":         "Push request from %s, length: %d bytes",
		"pull_request":         "Pull request from %s",
		"meta_request":         "Meta request from %s",
		"ping_request":         "Ping request from %s",
		"token_failed":         "%s request: Token verification failed (from %s)",
		"write_success":        "✓ Clipboard written, length: %d bytes",
		"read_success":         "✓ Clipboard read, length: %d bytes",
		"auto_start_enabled":   "✓ Auto-start enabled",
		"auto_start_disabled":  "✓ Auto-start disabled",
		"firewall_added":       "Firewall rule added (may need admin privileges or already exists)",
		"exit_signal":          "Exit signal received, program will exit",
		"log_file_opened":      "Log file opened",
		
		// 错误消息
		"error_network":        "Network interface error",
		"error_read_body":      "Failed to read request body",
		"error_write_clipboard":"Failed to write clipboard",
		"error_read_clipboard": "Failed to read clipboard",
		"error_config_save":    "Failed to save config",
		"error_server":         "Server error",
		"error_server_close":   "Server close error",
		"error_open_log":       "Failed to open log file",
		"error_exe_path":       "Failed to get executable path",
		"error_registry":       "Failed to open Run registry",
		"error_auto_start":     "Failed to write auto-start",
	},
	
	// 中文
	LangZH: {
		"app_title":            "剪贴板桥接",
		"service_address":      "📡 服务地址: %s",
		"local_address":        "💻 本机地址: %s",
		"auto_start":           "🚀 开机自启",
		"start_service":        "▶️  启动服务",
		"stop_service":         "⏸️  停止服务",
		"open_log":             "📄 打开日志文件",
		"quit":                 "❌ 退出",
		"ext_access":           "外部设备通过此地址访问",
		"local_test":           "本机测试使用",
		
		"log_separator":        "========================================",
		"program_start":        "程序启动，日志文件: %s",
		"config_loaded":        "配置加载完成: Port=%d, Token=%s, AutoStart=%v, LogLevel=%s",
		"clipboard_listener":   "剪贴板监听已启动",
		"service_started":      "🚀 剪贴板服务已启动",
		"external_access":      "   外部访问: http://%s:%d",
		"local_access":         "   本机访问: http://localhost:%d",
		"service_stopped":      "剪贴板服务已停止",
		"clipboard_updated":    "检测到本地剪贴板更新，内容长度: %d 字节",
		"clipboard_preview":    "剪贴板内容详情: %s",
		"push_request":         "收到 Push 请求 (来自 %s)，内容长度: %d 字节",
		"pull_request":         "收到 Pull 请求 (来自 %s)",
		"meta_request":         "收到 Meta 请求 (来自 %s)",
		"ping_request":         "收到 Ping 请求 (来自 %s)",
		"token_failed":         "%s 请求: Token 验证失败 (来自 %s)",
		"write_success":        "✓ 成功写入剪贴板，内容长度: %d 字节",
		"read_success":         "✓ 成功读取剪贴板，内容长度: %d 字节",
		"auto_start_enabled":   "✓ 已设置开机自启",
		"auto_start_disabled":  "✓ 已取消开机自启",
		"firewall_added":       "已尝试添加防火墙规则（可能需要管理员权限或规则已存在）",
		"exit_signal":          "收到退出信号，程序即将退出",
		"log_file_opened":      "已打开日志文件",
		
		"error_network":        "获取网络接口失败",
		"error_read_body":      "读取请求体失败",
		"error_write_clipboard":"写入剪贴板失败",
		"error_read_clipboard": "读取剪贴板失败",
		"error_config_save":    "保存配置文件失败",
		"error_server":         "服务错误",
		"error_server_close":   "服务器关闭错误",
		"error_open_log":       "打开日志文件失败",
		"error_exe_path":       "获取 exe 路径失败",
		"error_registry":       "打开 Run 注册表失败",
		"error_auto_start":     "写入开机自启失败",
	},
	
	// 日语
	LangJA: {
		"app_title":            "クリップボードブリッジ",
		"service_address":      "📡 サービスアドレス: %s",
		"local_address":        "💻 ローカルアドレス: %s",
		"auto_start":           "🚀 自動起動",
		"start_service":        "▶️  サービス開始",
		"stop_service":         "⏸️  サービス停止",
		"open_log":             "📄 ログファイルを開く",
		"quit":                 "❌ 終了",
		"ext_access":           "外部デバイスからこのアドレスでアクセス",
		"local_test":           "ローカルテスト用",
		
		"log_separator":        "========================================",
		"program_start":        "プログラム起動、ログファイル: %s",
		"config_loaded":        "設定読み込み完了: Port=%d, Token=%s, AutoStart=%v, LogLevel=%s",
		"clipboard_listener":   "クリップボードリスナー起動",
		"service_started":      "🚀 クリップボードサービス起動",
		"external_access":      "   外部アクセス: http://%s:%d",
		"local_access":         "   ローカルアクセス: http://localhost:%d",
		"service_stopped":      "クリップボードサービス停止",
		"clipboard_updated":    "ローカルクリップボード更新検出、長さ: %d バイト",
		"clipboard_preview":    "クリップボード詳細: %s",
		"push_request":         "Push リクエスト受信 (%s から)、長さ: %d バイト",
		"pull_request":         "Pull リクエスト受信 (%s から)",
		"meta_request":         "Meta リクエスト受信 (%s から)",
		"ping_request":         "Ping リクエスト受信 (%s から)",
		"token_failed":         "%s リクエスト: トークン検証失敗 (%s から)",
		"write_success":        "✓ クリップボード書き込み成功、長さ: %d バイト",
		"read_success":         "✓ クリップボード読み取り成功、長さ: %d バイト",
		"auto_start_enabled":   "✓ 自動起動設定完了",
		"auto_start_disabled":  "✓ 自動起動解除完了",
		"firewall_added":       "ファイアウォールルール追加試行（管理者権限が必要またはすでに存在）",
		"exit_signal":          "終了シグナル受信、プログラムを終了します",
		"log_file_opened":      "ログファイルを開きました",
		
		"error_network":        "ネットワークインターフェースエラー",
		"error_read_body":      "リクエストボディの読み取り失敗",
		"error_write_clipboard":"クリップボード書き込み失敗",
		"error_read_clipboard": "クリップボード読み取り失敗",
		"error_config_save":    "設定ファイル保存失敗",
		"error_server":         "サーバーエラー",
		"error_server_close":   "サーバークローズエラー",
		"error_open_log":       "ログファイルオープン失敗",
		"error_exe_path":       "実行ファイルパス取得失敗",
		"error_registry":       "レジストリオープン失敗",
		"error_auto_start":     "自動起動設定失敗",
	},
}

// 初始化语言设置
func initLanguage() {
	// 从环境变量或配置读取语言设置
	lang := os.Getenv("LANG")
	
	if lang == "" {
		// Windows: 尝试从系统语言检测
		lang = detectSystemLanguage()
	}
	
	// 解析语言代码
	lang = strings.ToLower(lang)
	if strings.HasPrefix(lang, "zh") {
		currentLang = LangZH
	} else if strings.HasPrefix(lang, "ja") {
		currentLang = LangJA
	} else {
		currentLang = LangEN
	}
}

// 检测系统语言
func detectSystemLanguage() string {
	// 简单的语言检测，可以根据需要扩展
	lang := os.Getenv("LANGUAGE")
	if lang == "" {
		lang = os.Getenv("LC_ALL")
	}
	if lang == "" {
		lang = os.Getenv("LC_MESSAGES")
	}
	return lang
}

// 获取翻译文本
func t(key string, args ...interface{}) string {
	i18nMu.RLock()
	defer i18nMu.RUnlock()
	
	// 尝试获取当前语言的翻译
	if langMap, ok := translations[currentLang]; ok {
		if text, ok := langMap[key]; ok {
			if len(args) > 0 {
				return fmt.Sprintf(text, args...)
			}
			return text
		}
	}
	
	// 回退到英文
	if langMap, ok := translations[LangEN]; ok {
		if text, ok := langMap[key]; ok {
			if len(args) > 0 {
				return fmt.Sprintf(text, args...)
			}
			return text
		}
	}
	
	// 如果都没有，返回key本身
	return key
}

// 设置语言
func setLanguage(lang string) {
	i18nMu.Lock()
	defer i18nMu.Unlock()
	
	switch lang {
	case LangEN, LangZH, LangJA:
		currentLang = lang
	default:
		currentLang = LangEN
	}
}

// 获取当前语言
func getLanguage() string {
	i18nMu.RLock()
	defer i18nMu.RUnlock()
	return currentLang
}

