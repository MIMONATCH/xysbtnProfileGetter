# xysbtn-profile-getter
虚研社按钮头像Gettet-获取b站对应up的头像并转换为webp格式

### 所属项目

- [MIMONATCH/xuyanshe-voice-button: 虚研社按钮 - 来自虚研社 小希🤖小桃🍑、小柔💚、兰音🐇、艾露露🐻的声音 (github.com)](https://github.com/MIMONATCH/xuyanshe-voice-button)

## Feature
- 并发下载
- 失败重试
- 支持GitHub Action部署

## 所需依赖
- [webp](https://developers.google.com/speed/webp/download)

## 配置
1. 不同环境下的配置
在`Macos`环境变量中添加下面两个变量、或者也仅在本次执行的命令前添加
```shell
  CGO_CFLAGS: -I/opt/homebrew/include
  CGO_LDFLAGS: -L/opt/homebrew/lib
```
在`Linux`和`Windows`请参考上面安装依赖的部分
2. `assets`文件夹下请根据要求的格式填写
```yaml
  supports:
    -
      uid: '209730937'
      name: '艾露露干脆面厂'
    -
    # 更多的配置对象
```

## 说明

1. 每七天会执行一次
2. GitHub Action的示例可以查看`.github`文件夹
3. 远程部署需要自己服务器的 public_key
