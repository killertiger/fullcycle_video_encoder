FROM golang:1.19-alpine3.18
ENV PATH="$PATH:/bin/bash" \
    BENTO4_BIN="/opt/bento4/bin" \
    PATH="$PATH:/opt/bento4/bin"

# FFMPEG
RUN apk add --update ffmpeg bash curl make

# Install Bento
# WORKDIR /tmp/bento4
# ENV BENTO4_BASE_URL="http://zebulon.bok.net/Bento4/source/" \
#     BENTO4_VERSION="1-5-0-615" \
#     BENTO4_CHECKSUM="5378dbb374343bc274981d6e2ef93bce0851bda1" \
#     BENTO4_TARGET="" \
#     BENTO4_PATH="/opt/bento4" \
#     BENTO4_TYPE="SRC"
    # download and unzip bento4
# RUN apk add --update --upgrade curl python2 unzip bash gcc g++ scons zip && \
#     wget -q ${BENTO4_BASE_URL}/Bento4-${BENTO4_TYPE}-${BENTO4_VERSION}${BENTO4_TARGET}.zip && \
#     sha1sum -b Bento4-${BENTO4_TYPE}-${BENTO4_VERSION}${BENTO4_TARGET}.zip | grep -o "^$BENTO4_CHECKSUM " && \
#     mkdir -p ${BENTO4_PATH} && \
#     unzip Bento4-${BENTO4_TYPE}-${BENTO4_VERSION}${BENTO4_TARGET}.zip -d ${BENTO4_PATH} && \
#     rm -rf Bento4-${BENTO4_TYPE}-${BENTO4_VERSION}${BENTO4_TARGET}.zip && \
#     apk del unzip && \
    # don't do these steps if using binary install
    # cd ${BENTO4_PATH} && scons -u build_config=Release target=x86_64-unknown-linux 
    # && \
    # cp -R ${BENTO4_PATH}/Build/Targets/x86_64-unknown-linux/Release ${BENTO4_PATH}/bin && \
    # cp -R ${BENTO4_PATH}/Source/Python/utils ${BENTO4_PATH}/utils && \
    # cp -a ${BENTO4_PATH}/Source/Python/wrappers/. ${BENTO4_PATH}/bin
ENV BENTO4_VERSION=1.6.0-637
ENV BENTO4_INSTALL_DIR=/opt/bento4
ENV PATH=/opt/bento4/bin:${PATH}
 
# Install dependencies.
RUN apk update \
  && apk add --no-cache \
  ca-certificates bash python3 wget libgcc make gcc g++

# Fetch source.
RUN cd /tmp/ \
  && wget -O Bento4-${BENTO4_VERSION}.tar.gz https://github.com/axiomatic-systems/Bento4/archive/v${BENTO4_VERSION}.tar.gz \
  && tar -xzvf Bento4-${BENTO4_VERSION}.tar.gz && rm Bento4-${BENTO4_VERSION}.tar.gz

# Create installation directories.
RUN mkdir -p \
  ${BENTO4_INSTALL_DIR}/bin \
  ${BENTO4_INSTALL_DIR}/scripts \
  ${BENTO4_INSTALL_DIR}/include

# Build.
RUN cd /tmp/Bento4-${BENTO4_VERSION}/Build/Targets/x86-unknown-linux \
  && make AP4_BUILD_CONFIG=Release

# Install.
RUN cd /tmp \
  && cp -r Bento4-${BENTO4_VERSION}/Build/Targets/x86-unknown-linux/Release/. ${BENTO4_INSTALL_DIR}/bin \
  && cp -r Bento4-${BENTO4_VERSION}/Source/Python/utils/. ${BENTO4_INSTALL_DIR}/utils \
  && cp -r Bento4-${BENTO4_VERSION}/Source/Python/wrappers/. ${BENTO4_INSTALL_DIR}/bin \
  && cp -r Bento4-${BENTO4_VERSION}/Source/C++/**/*.h . ${BENTO4_INSTALL_DIR}/include

# Cleanup.
RUN rm -rf /var/cache/apk/* /tmp/*



WORKDIR /go/src

#vamos mudar para o endpoint correto. Usando top apenas para segurar o processo rodando
ENTRYPOINT [ "top" ]