FROM mcr.microsoft.com/dotnet/sdk:5.0
WORKDIR /app

ADD . .
RUN dotnet build

ENTRYPOINT dotnet run
