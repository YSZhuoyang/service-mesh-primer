FROM mcr.microsoft.com/dotnet/core/sdk:3.1
# AS build-env
WORKDIR /app

# Build src
ADD . .
RUN dotnet build
# RUN dotnet restore

# Build release package
# RUN dotnet publish -c Release -o out

# Build runtime image
# FROM mcr.microsoft.com/dotnet/core/aspnet:3.0

# Install envoy
RUN apt update && apt -y install \
  software-properties-common \
  curl \
  ca-certificates \
  apt-transport-https \
  gnupg2
RUN curl -sL 'https://getenvoy.io/gpg' | apt-key add -
RUN add-apt-repository "deb [arch=amd64] https://dl.bintray.com/tetrate/getenvoy-deb \
  $(lsb_release -cs) stable"
RUN apt update && apt install -y getenvoy-envoy

# WORKDIR /app
# COPY --from=build-env /app/out .

# Start service with envoy sidecar
ADD ./start_service.sh ./start_service.sh
RUN chmod u+x ./start_service.sh
ENTRYPOINT ./start_service.sh
