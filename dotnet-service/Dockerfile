FROM mcr.microsoft.com/dotnet/core/sdk:3.1
WORKDIR /app

# Install envoy
RUN apt update && apt -y install \
  software-properties-common \
  curl apt-utils \
  ca-certificates \
  apt-transport-https \
  gnupg2
RUN curl -sL 'https://getenvoy.io/gpg' | apt-key add -
RUN add-apt-repository "deb [arch=amd64] https://dl.bintray.com/tetrate/getenvoy-deb \
  $(lsb_release -cs) stable"
RUN apt update && apt install -y getenvoy-envoy make

# Build src
ADD . .
RUN dotnet build

# Start service with envoy sidecar
ADD ./start_service.sh ./start_service.sh
RUN chmod u+x ./start_service.sh
ENTRYPOINT ./start_service.sh
