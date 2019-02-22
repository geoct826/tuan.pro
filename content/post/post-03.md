---
title: Setting up Angular 6 with Zurb Foundation
draft: false 
date: 2018-07-21
publishdate: 2018-07-21
image: https://images.unsplash.com/photo-1495435798646-a289417ece44?auto=format&fit=crop&w=1200&h=600
---

# Building a Project Starter

*This and the next few blog posts will cover my process of creating a project starter using an angular webapp along with a go lang based back end sever.*

## Setting up a workspace

When setting up a proejct, it is important to have a proper development workspace setup. This typically involves installing an Integrated Development Enviornment (IDE) on your computer along with all the programming packages required depending on what coding language and framework the project is built in. However, instead of using a typical desktop workstation, I'm going to be using a cloud IDE instead. Cloud IDE is essentially a workstation on the cloud. The advantages of such is that it reduces the compexity of having to setup your own workstation, allowing the user to quickly spin up their development environment in seconds no matter what device they're on, from a chromebook to macbook pro. While having played used [CodeAnywhere](https://www.codeanywhere.com/) in the past, for this project, I'm going to be using [Codenvy](https://codenvy.com/). The main reason for that is that my project setup requires a bit more resources then what the free tier of CodeAnywhere supports. So I took this as an opportunity to play around with Codenvy.

For this project I'm planning on developing a [Go](https://golang.org/doc/) backend service with an [Angular](https://angular.io/) front-end application on [Zurb Foundation](https://foundation.zurb.com/) framework. The project will be deployed to a [Google App Engine](https://cloud.google.com/appengine/) platform on Google Cloud. Google App Engine (GAE) is a quick and easy way to build and deploy an application without having to worry about setting up a cloud platform, and is fairly cost effective for a small or prototype project.

## Creating a Stack

The first thing when getting started with Codenvy is to add a workspace. Workspace are created from Stacks, which is a docker container that can range from an empty ubuntu server with JDK or a server with various packages built in. In my case, I want to create a stack using Google Cloud with GAE, along with GoLang and Angular, which requires Node.JS and NPM. Since there is so stack with this configuration, I am going to create one from a Dockerfile Recipe. So in the Stacks menu, I'm going to Build Stack From Recipe and then put in the following Docker recipe.


```
# Starting from a Google Cloud SDK docker image. 
FROM google/cloud-sdk

# Installing Go 1.9 (supported by GAE standard enviornment)
RUN curl -O https://storage.googleapis.com/golang/go1.9.7.linux-amd64.tar.gz && \
    tar -xf go1.9.7.linux-amd64.tar.gz && \
    mv go /usr/local && \
    rm go1.9.7.linux-amd64.tar.gz

# Installing Node.JS and NPM
RUN curl -sL https://deb.nodesource.com/setup_9.x | bash - && \
    apt-get update && apt-get install -y \
    nodejs build-essential

# Setting system path
ENV GOPATH /projects/.che/go_modules
ENV PATH $PATH:/usr/local/go/bin

# Installing Typescript and Angular
RUN npm i -g typescript && \
    npm i -g @angular/cli
```

From that, you can run a Test build of the recipe to make sure that the dockerfile works properly. After it's build, verify all the proper components are installed.

```
root@...:/# gcloud components list
Your current Cloud SDK version is: 215.0.0
The latest available version is: 215.0.0

root@...:/# node -v
v9.11.2

root@...:/# npm -v
5.6.0

root@...:/# go version
go version go1.9.7 linux/amd64

root@...:/# ng version
Angular CLI: 6.2.1
Node: 9.11.2
OS: linux x64
Angular:
...
```

Once verified, update the Stack Name, Description, Components, and Tag to your liking. This is just for organizational purpose and doesn't impact the workspace.

## Creating a Workspace

Save the Stack, and then you can create a Workspace from it. To create the workspace, select Add Workspace and select the Stack that was just created. With the free Codenvy account, you can use up to 3GB of RAM at once. The stack that was created requires at least 1GB of RAM to be stable. So if you select 1GB of RAM, you can have up to 3 of this workspace running at the same time. Once Create is selected, Codenvy will then again run the docker recipe to build the workspace.

Now the workspace is setup, I'm going to create 2 projects using the Project Explorer in this workspace. One is for the Go backend server, and the other is the angular webapp. Using Create Project, I create 2 blank projects with "/" as the parent directory. Then under the Commands Explore, I'm going to create 2 commands, one for build, which will build the angular webapp, and one for run, to run the go server.

Create 2 different commands with the following command line, replacing the project directories with your project directory.

Build command line
```
cd projects/<angular_project_directory> && npm install && ng build --watch
```
Run command line
```
cd projects/<go_server_project_directory> && dev_appserver.py app.yaml --host 0.0.0.0 --admin_host 0.0.0.0 --enable_host_checking false
```
Run preview URL
```
http://${server.port.8080}
```

Also create 2 server URL. The two server url is for the local development server and for the google app engine admin console server


You can try running the commands now, but as there are no files in any of the projects yet, it'll just fail.

So let's get started creating the angular project scaffolding first. The easiest way to do this is through the angular command line. 

So in the terminal, go into the projects folder
> cd /projects

And then run the angular CLI command to create new project, using the routing flag to create a good starting point for application pagination and routing and making sure that project name is the same as your project directory, so this way the CLI creates the files within the proper directory.
> ng new <project name> --routing

You should now see a bunch of generated files within the angular prject directory.

Before we build it, we're going to make a change in angular.json file to specify a different distribution directory.

```
"projects": {
  "<project name>": {
    "targets": {
      "build": {
        "options": {
          "outputPath": "../<go server project>/dist",
          ...
        }
      }
    }
  }
}
```

Now if you run the build command, you should see a /dist folder created in the go server project. You may need to refresh the Projects Explorer view or you can use the terminal 

```
root@...:/# ls /projects/<angular server>/
README  dist
```

Now to create a simple go server.

Within the go server project create a app.yaml file.
```
runtime: go
api_version: go1

handlers:

# All URLs are handled by the Go application script
- url: /.*
  script: _go_app
```

Now create a main.go file.
```
package main

import (
        "net/http"
        "google.golang.org/appengine"
)

func main() {
    angularFileServer := http.FileServer(http.Dir("dist"))    
    http.Handle("/", angularFileServer)
    appengine.Main()
}

```

Install the appengine package
> go get google.golang.org/appengine

Once that's done, run the server command. That will bring up the server run window with a preview link, which will be something like http://nodeX.codenvy.io:port. Click on the link and you should get the default angular splash screen. 

Last thing is to now push this onto the google cloud platform. Assuming you already created a google cloud project, you run 
 >gcloud init
 select 1. Re-initialize this configuration [default] with new setting
 
 Log in. which will direct you to copy an oAuth2 link into your browser to get a verification code.
 
 If validated successfully, you can link folder to your existing project, and that's it.
 
 Then to deploy the app
 
 > gcloud app deploy
 
Select your region and then google cloud will do it's magic. Once deployed, you'll get a link that something like https://your-project-name.appspot.com. 
