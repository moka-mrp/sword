package baseTpl

const TplDockerfile  = `# Base image
FROM  shanbumin/golang:1.14-alpine3.11
# Author
MAINTAINER Sam <shanbumin@qq.com>
# COPY
ADD ./bin/shanbumin.run   .
# Cmd
CMD ["./shanbumin.run","run"]`
