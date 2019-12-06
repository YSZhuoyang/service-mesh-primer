### Build Dotnet service

1. Run `dotnet restore` to install dependency packages.
2. Copy `contracts/greeter2/google/api` folder (from googleapi repo) to `build/native/include/google/` dir under where `grpc.tools` package is installed.
3. Run `dotnet build`.
