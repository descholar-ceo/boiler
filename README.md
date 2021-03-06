# boiler
[![Maintainability](https://api.codeclimate.com/v1/badges/0ef5f6e9398a22c4b5ee/maintainability)](https://codeclimate.com/github/descholar-ceo/boiler/maintainability) [![Build Status](https://travis-ci.org/descholar-ceo/boiler.svg?branch=main)](https://travis-ci.org/descholar-ceo/boiler)  [![Coverage Status](https://coveralls.io/repos/github/descholar-ceo/boiler/badge.svg?branch=main)](https://coveralls.io/github/descholar-ceo/boiler?branch=main) [![Join the chat at https://gitter.im/boiler-community/community](https://badges.gitter.im/boiler-community/community.svg)](https://gitter.im/boiler-community/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

__boiler__ is an application that helps you to generate a project's boilerplate very easily. You don't need to waste your first hours of development in setting up the project. It gives you the first setup of your project so that you can focus on writing your codes.

## Why do we need boiler?
- I have noticed that when you start working on a new project, it can take you between 4 - 8 hours setting up the project boilerplate, (project skeleton), depending on the technologies you are going to use. So, that's why I developed this tool that will save you time.

- Even though some frameworks like [Ruby on Rails](https://rubyonrails.org/) or [Express.js](https://expressjs.com/) have their commands to initialize a project, but I found them as not complete, because there are some cases where you need to use some more technologies than what you are given by running those commands, so this tool will save you in that case as well. It will do everything for you.

## How to use it?

##### Note: Before using it, please make sure that you read all pre-requisites and your system adhere to them

### Pre-requisites
#### For ruby projects
- [ ] Ruby should be installed already on your computer
- [ ] Bundler

##### For Ruby on Rails project
To get your project initiated, you should have the following tools already installed on your computer:
- [ ] Ruby (Installed and working well)
- [ ] Ruby on rails (Installed and working well)
- [ ] Yarn
- [ ] Bundler
- [ ] And checkout the database of your choice requirements below :point_down: :point_down: :point_down:

 ##### NOTE: DATABASES REQUIREMENTS
1. [MySQL](https://www.mysql.com/): requires you to have pre-installed `libmysqlclient-dev`, on your system, if you are not sure run `sudo apt-get install libmysqlclient-dev` which will install it if you don't have it or update it if the version you have is obsolete

1. [Oracle](https://www.oracle.com/database/technologies/): requires you to have ruby-oci8 installed, and to install it, requires you to have oracle environment installed properly on your computer

1. [Frontbase](http://www.frontbase.com/cgi-bin/WebObjects/FBWebSite): requires you to be a MacOS user

1. [ibm_db](https://www.ibm.com/support/knowledgecenter/hr/SSEPGG_9.7.0/com.ibm.db2.luw.qb.server.doc/doc/t0008875.html): requires you to have `db2` server installed on your computer

1. [Sqlserver](https://www.microsoft.com/en-us/sql-server/sql-server-downloads): requires you to have pre-installed `freeTDS` on your system, 
if you are not sure run `sudo apt-get install freetds-dev` which will install it if you don't have it or update it if the version you have is obsolete

1. jdbcmysql, jdbcsqlite3, jdbcpostgresql,: are for use with [JRuby](https://www.jruby.org/) only

## Installation :electric_plug:
- [Click here](https://github.com/descholar-ceo/boiler/releases/tag/v1.0.0) and download a zip file
- Extract it wherever you wish
- Open terminal in the extracted folder, and run `./install.sh`
- Close your current terminal and open another one and run `boiler` :point_right: :point_right: :point_right: there you go!

### How to get a help? :pray:
Please join our community on [Gitter](https://gitter.im/boiler-community/community) and post any question you have about this project.

## Built with
[Golang](https://golang.org/)

## Languages and systems
1. The current version [v1.0.0](https://github.com/descholar-ceo/boiler/releases/tag/v1.0.0) can generate:
   - Basic ruby Projects
   - Ruby on Rails Projects
   - :warning: Other languages are coming soon in later versions

1. The current version [v1.0.0](https://github.com/descholar-ceo/boiler/releases/tag/v1.0.0) support linux systems (Ubuntu/debian) based OS that use [Bash](https://www.gnu.org/software/bash/) as their command line shell. :point_right: We welcome whoever who can add some other system support, your contributions are welcomed!

## Contributions
Please follow [this contribution guide](https://github.com/descholar-ceo/boiler/blob/main/CONTRIBUTING.md)
