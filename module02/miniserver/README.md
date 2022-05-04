# miniserver
<a href="http://www.wtfpl.net"><img src="https://img.shields.io/badge/license-WTFPL-blue" alt="LICENSE"></a>
<a href="https://go.dev"><img src="https://img.shields.io/badge/language-golang-blue" alt="language"></a>

# Overview
miniserver基于golang实现了轻量级http server 

## 基础功能

- [x] 接收客户端 request，并将 request 中带的 header 写入 response header 
- [x] 读取当前系统的环境变量中的 VERSION 配置，并写入 response header 
- [x] Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出 
- [x] 当访问 localhost/healthz 时，应返回 200

## 工程化实践

- [x] 规范化目录结构
    - 符合: https://github.com/golang-standards/project-layout
- [ ] 参数配置化
- [ ] Makefile
- [ ] 应用容器化
- [ ] 应用部署K8S
- [ ] 接口测试
- [ ] CI/CD
- [ ] 自动化测试