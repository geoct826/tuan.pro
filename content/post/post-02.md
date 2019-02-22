---
title: Setting up Hugo.
draft: false 
date: 2018-05-12
publishdate: 2018-05-12
image: https://images.unsplash.com/photo-1484417894907-623942c8ee29?auto=format&fit=crop&w=1200&h=600
---

After creating an IDE, it needs to be setup with Hugo and the code to create a hugo site.

First, installing Hugo.

While Hugo is not available in Ubuntu's default repositories, packages are available on GitHub for various architectures and distributions. The one that we're looking for is the Linux 64-bit distribution. You can find the latest Hugo packages at https://github.com/gohugoio/hugo/releases/. To install it, first, open up the SSH Terminal, by right-clicking on the container and selecting SSH Terminal.
Downloading the Hugo package using the the wget command to download the file at the link location:

```
cabox@box-codeanywhere:~/workspace$ wget  https://github.com/gohugoio/hugo/releases/download/v0.35/hugo_0.35_Linux-64bit.deb
```

Next, install the package with dpkg by typing:

```
cabox@box-codeanywhere:~/workspace$ sudo dpkg -i hugo*.deb
```

And then verify the install.

```
cabox@box-codeanywhere:~/workspace$ hugo version
Hugo Static Site Generator v0.35 linux/amd64 BuildDate: 2018-01-31T10:44:43Z
```

Now that Hugo is installed, the next step is to setup a development webserver. Hugo being a static website generator, this webserver is used to test the Huo generated site. Setting up a web server in go is relatively simple.

In the root workspace create a file called server.go by right clicking on the container name and select Create File.

In server.go, use the following code

```
package main

import (
  "os"
  "log"
  "net/http"
)

func main() {
  port := os.Getenv("port")

  fs := http.FileServer(http.Dir("public/"))
  http.Handle("/", fs)
	
  log.Println("Starting webserver listening on", port)
  http.ListenAndServe(":"+port, nil)
}
```

This server.go creates a golang webserver serving files in a directory called public on the port specified in the config file. Next, open up the config file to setup commands and parameters necessary for the Stack to run. Right click on the container name to open up the config, and add the port number as a envionment varialbe.

```
"environment": {
    "port": "3000"
},
```

And dd the command to start the go webserver in the config files.

```
"commands": [
    "go run server.go"
],
```

Next, start the webserver by selecting Run, you should see the log line "Starting webserver listening on 3000"

To preview the site, open up the info screen on the container and see the provided custom link to access the website.

Because we haven't added a public directory, you should see "404 page not found".

So lets add an index.html file into the public directory. First create a directory called public and add in the following lines into index.html

```
<body>
  <h1>
    Hello World
  </h1>
</body>
```

Now, if you go to the preview link, you should see Hello World. What the webserver is doing is serving whtever static file is in the public directory. So if you go an edit the index.html file that was just created to

```
<body>
  <h1>
    Hello Bob
  </h1>
</body>
```

And if you go tho the peview link, you should see it now as Hello Bob.