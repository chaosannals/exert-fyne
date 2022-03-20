# [exert-fyne](https://github.com/chaosannals/exert-fyne)

```bash
# 新项目建立, 引入并执行包整理
go get fyne.io/fyne/v2
go mod tidy

# 安装命令行工具
go get fyne.io/fyne/cmd/fyne

# 打包指定平台
fyne package -os darwin -icon demo.png
fyne package -os linux -icon demo.png
fyne package -os windows -icon demo.png
fyne package -os android -appID com.example.demo -icon demo.png
fyne package -os ios - appID com.example.demo -icon demo.png
```
