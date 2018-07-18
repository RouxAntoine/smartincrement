FROM golang

LABEL author="antoine roux <antoinroux@hotmail.fr>"

ARG PROJECT_NAME

WORKDIR /go/src/

RUN apt-get -y update && \
    apt-get install -y vim

COPY projects/Makefile .

ENV PROGRAM_NAME $PROJECT_NAME

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# CMD bash -c "make run; sleep infinity"
CMD [ "make", "run" ]
