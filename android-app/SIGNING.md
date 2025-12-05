# 🔐 Android APK 签名配置指南

## 📝 签名类型对比

| 特性 | Debug 签名 | Release 签名 |
|------|-----------|-------------|
| **可以安装** | ✅ 是 | ✅ 是 |
| **用户体验** | ⚠️ "测试版"警告 | ✅ 正常安装 |
| **发布到 Google Play** | ❌ 否 | ✅ 是 |
| **需要配置** | ✅ 自动 | ❌ 需要 |
| **密钥安全** | ⚠️ 公开的 | ✅ 私有的 |
| **应用更新** | ⚠️ 任何人都能覆盖 | ✅ 只有相同密钥能更新 |
| **APK 大小** | 较大（~5MB） | 较小（~3MB） |
| **代码混淆** | ❌ 否 | ✅ 是 |
| **GitHub Actions** | ✅ 自动构建 | 📝 需要配置密钥 |
| **适用场景** | GitHub/开源分发 | Google Play 商店 |

## 🎯 推荐方案

### 方案 A: 仅使用 Debug 签名（当前方案）⭐

**适合：**
- ✅ 开源项目
- ✅ GitHub Releases 分发
- ✅ 个人/内部使用
- ✅ 快速迭代

**优点：**
- 零配置，自动构建
- 用户可以直接安装
- 完全透明和开源

**缺点：**
- 系统显示"测试版"提示
- 不能发布到 Google Play

### 方案 B: 同时提供 Debug 和 Release（推荐）⭐⭐⭐

**配置步骤：**

#### 1. 生成签名密钥

```bash
cd android-app

# 生成密钥（一次性操作）
keytool -genkey -v -keystore release.jks \
  -alias clipboard-bridge \
  -keyalg RSA \
  -keysize 2048 \
  -validity 10000

# 会提示输入：
# - 密钥库口令（记住！）
# - 密钥口令（记住！）
# - 姓名、组织等信息
```

#### 2. 配置 GitHub Secrets

访问仓库设置：
```
https://github.com/copypasteengine/clipboard-bridge/settings/secrets/actions
```

添加以下 Secrets（点击 "New repository secret"）：

| Name | Value | 说明 |
|------|-------|------|
| `ANDROID_KEYSTORE_BASE64` | (见下方) | 密钥库的 Base64 编码 |
| `ANDROID_KEYSTORE_PASSWORD` | `你的密钥库口令` | 密钥库密码 |
| `ANDROID_KEY_ALIAS` | `clipboard-bridge` | 密钥别名 |
| `ANDROID_KEY_PASSWORD` | `你的密钥口令` | 密钥密码 |

**生成 Base64 编码：**

```bash
# Linux/macOS
base64 -i release.jks | tr -d '\n' > keystore.base64.txt
cat keystore.base64.txt

# Windows (PowerShell)
[Convert]::ToBase64String([IO.File]::ReadAllBytes("release.jks")) > keystore.base64.txt
type keystore.base64.txt
```

将输出的内容复制到 `ANDROID_KEYSTORE_BASE64` secret 中。

⚠️ **重要：** 
- 密钥库文件（.jks）不要提交到 Git
- GitHub Secrets 是加密的，安全
- 密码妥善保管，丢失无法恢复

#### 3. 更新 GitHub Actions

编辑 `.github/workflows/build-android.yml`，添加签名配置：

```yaml
- name: Setup Signing Config
  if: ${{ secrets.ANDROID_KEYSTORE_BASE64 != '' }}
  run: |
    # 解码密钥库
    echo "${{ secrets.ANDROID_KEYSTORE_BASE64 }}" | base64 -d > android-app/release.jks
    
    # 创建签名配置文件
    cat > android-app/keystore.properties << EOF
    storePassword=${{ secrets.ANDROID_KEYSTORE_PASSWORD }}
    keyPassword=${{ secrets.ANDROID_KEY_PASSWORD }}
    keyAlias=${{ secrets.ANDROID_KEY_ALIAS }}
    storeFile=../release.jks
    EOF

- name: Build Release APK (signed)
  if: ${{ secrets.ANDROID_KEYSTORE_BASE64 != '' }}
  run: |
    cd android-app
    ./gradlew assembleRelease --stacktrace --no-daemon
```

#### 4. 更新 app/build.gradle.kts

在 `android-app/app/build.gradle.kts` 中添加：

```kotlin
// 在 android 块之前
val keystorePropertiesFile = rootProject.file("keystore.properties")
val keystoreProperties = java.util.Properties()

if (keystorePropertiesFile.exists()) {
    keystoreProperties.load(java.io.FileInputStream(keystorePropertiesFile))
}

android {
    // ... 其他配置 ...
    
    signingConfigs {
        create("release") {
            if (keystorePropertiesFile.exists()) {
                keyAlias = keystoreProperties["keyAlias"] as String
                keyPassword = keystoreProperties["keyPassword"] as String
                storeFile = file(keystoreProperties["storeFile"] as String)
                storePassword = keystoreProperties["storePassword"] as String
            }
        }
    }

    buildTypes {
        release {
            signingConfig = if (keystorePropertiesFile.exists()) {
                signingConfigs.getByName("release")
            } else {
                null  // 如果没有配置，生成未签名的 APK
            }
            isMinifyEnabled = true
            proguardFiles(
                getDefaultProguardFile("proguard-android-optimize.txt"),
                "proguard-rules.pro"
            )
        }
    }
}
```

#### 5. 推送并测试

```bash
git add .
git commit -m "feat: add Release APK signing support"
git push origin main

# 重新触发构建
git push origin :refs/tags/v1.0.0
git tag -d v1.0.0
git tag v1.0.0
git push origin v1.0.0
```

**构建成功后，Release 将包含：**
- `clipboard-bridge-android-v1.0.0-debug.apk` - Debug 版本
- `clipboard-bridge-android-v1.0.0-release.apk` - Release 版本（已签名）

### 方案 C: 仅使用 Release 签名

如果想只发布 Release 版本，配置好签名后，在 workflow 中移除 Debug 构建即可。

---

## 🔒 安全建议

### 保护你的密钥

1. **永远不要**将 `.jks` 文件提交到 Git
2. **永远不要**在代码中硬编码密码
3. **使用 GitHub Secrets** 存储敏感信息
4. **备份密钥文件** 到安全位置（丢失无法恢复）
5. **定期更换密码**

### 如果密钥泄露

1. 立即撤销 GitHub Secrets
2. 生成新的密钥
3. 用新密钥发布新版本
4. 旧版本无法升级到新版本（签名不同）

---

## 💡 我的推荐

**对于你的项目：**

### 现阶段：使用 Debug 签名（当前方案）✅

**理由：**
- 项目刚起步，主要通过 GitHub 分发
- 用户是技术人员，不介意"测试版"提示
- 简单、透明、完全开源
- 零配置，自动构建

**用户安装体验：**
```
1. 下载 APK
2. 系统提示"此应用未经 Google Play 验证"
3. 点击"仍要安装"
4. 安装成功 ✓
```

### 未来：添加 Release 签名（可选）

**什么时候需要：**
- 准备发布到 Google Play 商店
- 用户数量增长，需要更专业的形象
- 需要应用内购买等商店功能
- 需要防止他人篡改你的应用

---

## 📊 对比总结

### 当前方案（Debug 签名）

```
优点：
✅ 自动构建
✅ 零配置
✅ 用户可以安装
✅ 完全开源
✅ 快速迭代

缺点：
⚠️ 系统提示"测试版"
⚠️ 不能上架 Google Play
⚠️ 任何人都能用相同签名
```

### 升级方案（Release 签名）

```
优点：
✅ 专业形象
✅ 可以上架 Google Play
✅ 唯一的签名密钥
✅ 代码混淆
✅ APK 更小

缺点：
⚠️ 需要配置密钥
⚠️ 需要管理 Secrets
⚠️ 密钥丢失无法恢复
```

---

## 🎯 结论

**当前的 Debug 签名 APK 完全可以使用！**

用户安装步骤：
1. 下载 APK 到手机
2. 设置 → 安全 → 允许安装未知来源应用
3. 点击 APK 文件安装
4. 系统可能提示"测试版应用"
5. 点击"仍要安装"
6. 安装成功 ✓

**没有任何功能限制，只是系统会有个"测试版"的提示而已。**

如果未来需要配置 Release 签名，告诉我，我可以帮你配置！
