#FROM mongo:4.0
#
#RUN service mongod start
#RUN mongo
#RUN use user
#RUN db.createCollection("mycollection")



# https://medium.com/@pablo_ezequiel/creating-a-docker-image-with-mongodb-4c8aa3f828f2
#FROM ubuntu:18.04
#
#MAINTAINER smet_k
#
## Update the repository sources list
#RUN apt-get update && apt-get install -y gnupg2
#
################### BEGIN INSTALLATION ######################
## Install MongoDB Following the Instructions at MongoDB Docs
## Ref: http://docs.mongodb.org/manual/tutorial/install-mongodb-on-ubuntu/
#
## Add the package verification key
#RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 7F0CEB10
#
## Add MongoDB to the repository sources list
#RUN echo 'deb http://downloads-distro.mongodb.org/repo/ubuntu-upstart dist 10gen' | tee /etc/apt/sources.list.d/mongodb.list
#
## Update the repository sources list once more
#RUN apt-get update
#
## Install MongoDB package (.deb)
#RUN apt-get install -y mongodb-10gen
#
## Create the default data directory
#RUN mkdir -p /data/db
#
###################### INSTALLATION END #####################
#
## Expose the default port
#EXPOSE 27017
#
## Default port to execute the entrypoint (MongoDB)
## CMD ["--port 27017"]
#CMD ["--port 27017", "--smallfiles"]
#
## Set default container command
#ENTRYPOINT usr/bin/mongod

# off site
#https://github.com/docker/docker.github.io/blob/master/engine/examples/mongodb/Dockerfile

FROM ubuntu:18.04

#RUN apt remove --autoremove mongodb-org
#RUN rm /etc/apt/sources.list.d/mongodb*.list

ENV TZ=Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apt-get update && apt-get install -y gnupg

# Installation:
# Import MongoDB public GPG key AND create a MongoDB list file
RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv E52529D4
RUN bash -c 'echo "deb [arch=amd64] http://repo.mongodb.org/apt/ubuntu bionic/mongodb-org/4.0 multiverse" > /etc/apt/sources.list.d/mongodb-org-4.0.list'

#RUN apt-get install --reinstall ca-certificates

# Update apt-get sources AND install MongoDB
RUN apt-get update && apt-get install -y mongodb-org

# Create the MongoDB data directory
RUN mkdir -p /data/db

# Expose port #27017 from the container to the host
EXPOSE 27017

# Set /usr/bin/mongod as the dockerized entry-point application
ENTRYPOINT ["/usr/bin/mongod"]
