# PicLiveCode

PicLiveCode 是一个简单的图片上传和二维码生成服务。用户可以通过网页界面轻松上传图片，并获得相应的二维码。

## 功能特点

- 用户友好的网页界面
- 点击选择图片上传
- 自动生成并显示包含图片链接的二维码
- 支持多种图片格式（JPG、JPEG、PNG、GIF）

## 操作步骤

### 1. 从源码运行

1. 确保你已安装 Go 1.21.3 或更高版本。
2. 克隆仓库：
   ```
   git clone https://github.com/yourusername/piclivecode.git
   cd piclivecode
   ```
3. 安装依赖：
   ```
   go mod download
   ```
4. 运行应用：
   ```
   go run main.go
   ```
5. 在浏览器中访问 `http://localhost:8080` 打开主页。

### 2. 打包成可执行文件

1. 在项目根目录下执行：
   ```
   # 对于 Windows 用户
   go build -o piclivecode.exe
   
   # 对于 macOS 和 Linux 用户
   go build -o piclivecode
   ```
2. 这将生成一个名为 `piclivecode` 的可执行文件。
3. 运行可执行文件：
   ```
   ./piclivecode
   ```
4. 在浏览器中访问 `http://localhost:8080` 打开主页。

### 3. 构建 Docker 镜像

1. 确保你已安装 Docker。
2. 在项目根目录下构建 Docker 镜像：
   ```
   docker build -t piclivecode .
   ```
3. 运行 Docker 容器：
   ```
   docker run -p 8080:8080 piclivecode
   ```
4. 在浏览器中访问 `http://localhost:8080` 打开主页。

## 使用方法

1. 打开主页后，点击"选择图片"按钮。
2. 在弹出的文件选择器中选择要上传的图片。
3. 图片会自动上传，上传成功后会显示一个包含图片链接的二维码。
4. 您可以使用手机扫描二维码来查看上传的图片。

注意：确保 `uploads` 目录存在并具有适当的写入权限。
