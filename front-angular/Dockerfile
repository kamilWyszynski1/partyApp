FROM node:9.6.1

#set working directory
RUN mkdir /usr/src/app
WORKDIR /usr/src/app

ENV PATH /usr/src/app/node_modules/.bin:$PATH

# install and cache app dependencies
COPY package.json /usr/src/app/package.json
RUN npm install
RUN npm install -g @angular/cli@7.3.8

EXPOSE 4200
COPY . /usr/src/app
ENTRYPOINT ["ng", "serve", "--host", "127.0.0.1"]