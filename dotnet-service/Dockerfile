# Note: build dotnet-service with this file from outer(root) directory

FROM mcr.microsoft.com/dotnet/sdk:10.0 AS build-env
# Install dependencies for protoc and the compiler itself to avoid segfaults on some ARM64 environments
RUN apt-get update && apt-get install -y --no-install-recommends \
    libatomic1 \
    protobuf-compiler \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY ./dotnet-service .
COPY ./contracts ../contracts
# Build release package, overriding the bundled protoc with the system one
RUN dotnet publish -c Release -o out /p:Protobuf_ProtocFullPath=$(which protoc)

# Build runtime image
FROM mcr.microsoft.com/dotnet/aspnet:10.0
WORKDIR /app

COPY --from=build-env /app/out ./

ENTRYPOINT ["dotnet", "server.dll"]
