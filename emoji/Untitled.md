各版本区别
版本	说明	示例
alpine	Alpine Linux 操作系统，它是一个独立发行版本，相比较 Debian 操作系统来说 Alpine，更加轻巧
https://alpinelinux.org/	docker pull node:alpine3.14
squeeze	Debian 6	
wheezy	Debian 7	
jessie	Debian 8，更旧的稳定（oldoldstable）版，马上就淘汰了 2015
https://wiki.debian.org/DebianJessie	docker pull node:10.22.0-jessie
stretch	Debian 9，旧的稳定（oldstable）版，现有长期支持 2017，比较老，目前除了 LTS 其他版本已经不再提供技术支持了，所以我们非必要情况下还是不要选择它比较好
https://wiki.debian.org/DebianStretch	docker pull node:lts-stretch
buster	Debian 10，当前的稳定（stable）版 2019，比较新，支持比较全面，受广大Debian爱好者的好评！像 PHP、Python 之类的语言、应用都会使用这个版本的 Debian 搭建 Docker 基础镜像
https://wiki.debian.org/DebianBuster	docker pull node:lts-buster
bullseye	Debian 11，即将上位的稳定（stable）版 2021
https://wiki.debian.org/DebianBullseye	docker pull node:lts-bullseye
xxx-slim	一般都基于 Debian 和 glibc，删除了许多非必需的软件包，优化了体积	docker pull node:stretch-slim
docker pull node:buster-slim
参考