FROM ubuntu:18.04 AS Server

WORKDIR /app

COPY api-go ./api-go

RUN chmod +x ./api-go

RUN apt update -y && apt install ca-certificates -y

# ENV TZ=America/Recife

# RUN echo $TZ > /etc/timezone && \
#     apt-get update && apt-get install -y tzdata && \
#     rm /etc/localtime && \
#     ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
#     dpkg-reconfigure -f noninteractive tzdata && \
#     apt-get clean
    
CMD ["./api-go"]

EXPOSE 9080