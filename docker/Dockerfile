FROM nginx
EXPOSE 9000
WORKDIR "/note-gin"
COPY ./note-gin/* ./
RUN ["mkdir","config","log"]
RUN ["mv","file","config/"]
RUN ["touch","./log/log.log"]
CMD ["./app","-c","config/file/BootLoader.yaml"] 

