FROM mcr.microsoft.com/dotnet/sdk:8.0-alpine AS build
WORKDIR /source

COPY src/. .
RUN dotnet publish ./WebApi.csproj -c release -o out

FROM mcr.microsoft.com/dotnet/aspnet:8.0-alpine
WORKDIR /app
RUN apk --no-cache add curl
COPY --from=build /source/out .

ENV ASPNETCORE_HTTP_PORTS=8080
EXPOSE 8080

RUN echo "dotnet WebApi.dll" > entypoint.sh && chmod a+x entypoint.sh

ENTRYPOINT ["sh", "entypoint.sh"]