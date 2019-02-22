---
title: Creating a blog.
draft: false 
date: 2018-05-05
publishdate: 2018-05-02
image: https://images.unsplash.com/photo-1489533119213-66a5cd877091?auto=format&fit=crop&w=1200&h=600
---
Domain names are pretty straightforward to come by these days. There are lots of domain sites like Hover, Namecheap, GoDaddy, HostGator, are just some examples. Moreover, new generic top-level domains like .me,  .beer, .cafe, are ensuring that there is no shortage of creative domain names going forward. However, what do you do once you find and purchase the perfect domain name? 

There are a lot of services out there that help people create their own personal website. But to have full control of your own site and layout, nothing can compare with the ultimate flexibility you get when you create the site from code. True, building a website from scratch is not for the faint of heart, however, it is a rewarding experience.

This site is built from scratch using [Hugo](https://gohugo.io) as a static website generator and [GitHub](https://github.com) as hosting solution. In this first series of blogs, I intend to go through the process I used to create, manage, and publish this site. However, before I get into the development details, the very first thing that any developer needs is a development environment. Setting up a development environment can be a tedious process, but with the advent of some great cloud-based Integrated Development Environment (IDE),  this process is greatly simplified, all you need is a web browser.

For this development, I'm using [Codeanywhere](https://codeanywhere.com) as my cloud-based IDE. It simplifies the development environment setup and provides the flexibility to be able to work on my site regardless of which computer I am on. They now they even support iOS and Android devices so I can edit files on my iPhone.

Setting up a Codeanywhere account is pretty straightforward; using my GitHub account, it quickly created the account and even automatically setup SSH keys between the GitHub account and the Codeanywhere environment. 

When first signing into Codeanywhere, it automatically creates a default project and takes the user directly to creating a container in the project. Containers are virtual development environments and come in many flavors with either Ubuntu 14.04 or Centos 6.5 as the base operating system and a variety of development setup. Since I'm planning on using Hugo as my static website generator, I am using the Go Development Stack.

And that's it. With this development environment setup, I am now ready to start building my site from scratch.
