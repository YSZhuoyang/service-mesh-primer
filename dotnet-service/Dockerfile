# Note: build dotnet-service with this file from outer(root) directory

FROM mcr.microsoft.com/dotnet/sdk:8.0-jammy AS build-env
WORKDIR /app

COPY ./dotnet-service .
COPY ./contracts ../contracts
# Build release package
RUN dotnet publish -c Release -o out

# Build runtime image
FROM mcr.microsoft.com/dotnet/aspnet:8.0-jammy
WORKDIR /app

COPY --from=build-env /app/out ./

ENTRYPOINT dotnet server.dll
