# xysbtnProfileGetter
 虚研社按钮头像getter

### 所属项目

- [MIMONATCH/xuyanshe-voice-button: 虚研社按钮 - 来自虚研社 小希🤖小桃🍑、小柔💚、兰音🐇、艾露露🐻的声音 (github.com)](https://github.com/MIMONATCH/xuyanshe-voice-button)

## Feature
- 并发下载
- 失败重试
- 支持GitHub Action部署

## 配置

`assets`文件夹下请根据要求的格式填写
```
  supports:
    -
      uid: '209730937'
      name: '艾露露干脆面厂'
    -
    # 更多的配置对象
```

## 说明

1. 每7天会执行一次
2. GitHub Action的示例可以查看`.github`文件夹
3. 远程部署需要自己服务器的 public_key
