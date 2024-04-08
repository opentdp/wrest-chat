# Wrest-chat for Docker

此镜像仅用于部署 wrest 主体程序，不包含 wcf.exe 和 wechat，请配置 WCF_ADDRESS 连接外部 wcf 服务使用。

## 快速开始

```shell
docker run -d -p 7600:7600 \
    -e WCF_ADDRESS="192.168.1.2:7601" \
    -v ./storage:/srv/storage \
    rehiy/wrest-chat
```

### 参数说明

- `-e WCF_ADDRESS="192.168.1.2:7601"` 指定 wcf 地址
- `-v ./storage:/srv/storage` 指定数据存储路径
- `-v ./plugin:/srv/plugin` 指定插件存储路径
- `-v ./log:/srv/log` 指定日志存储路径
- `rehiy/wrest-chat /srv/storage/config.yml` 自定义配置文件路径

## 其他信息

官方仓库 <https://github.com/opentdp/wrest-chat>
