# Note: build dotnet-service with this file from outer(root) directory

FROM mcr.microsoft.com/dotnet/sdk:9.0 AS build-env
WORKDIR /app

COPY ./dotnet-service .
COPY ./contracts ../contracts
# Build release package
RUN dotnet publish -c Release -o out

# Build runtime image
FROM mcr.microsoft.com/dotnet/aspnet:9.0
WORKDIR /app

COPY --from=build-env /app/out ./

ENTRYPOINT dotnet server.dll
