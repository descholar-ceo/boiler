# boiler
[![Maintainability](https://api.codeclimate.com/v1/badges/0ef5f6e9398a22c4b5ee/maintainability)](https://codeclimate.com/github/descholar-ceo/boiler/maintainability) [![Build Status](https://travis-ci.org/descholar-ceo/boiler.svg?branch=develop)](https://travis-ci.org/descholar-ceo/boiler)  [![Coverage Status](https://coveralls.io/repos/github/descholar-ceo/boiler/badge.svg?branch=develop)](https://coveralls.io/github/descholar-ceo/boiler?branch=develop) [![Join the chat at https://gitter.im/boiler-community/community](https://badges.gitter.im/boiler-community/community.svg)](https://gitter.im/boiler-community/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)


## Project vision
Nowadays, programming is taking over the world of computers, almost every programmer thinks of the next project while still working on the current project. Since programming has occupied almost all of the technology we use these days; setting up a new project take us some minutes, or hours.

This project `boiler` is here to make your life very easy while you are planning to work on your new project. We want to make it easy while you bootstrap your project, by running few commands, it will give you the project skeleton set, ready to continue working on your project.

## Project description

This project is a command line utility, easy to be intalled, and after the installation, you will be able to do the following with it:
- Setup a basic ruby project
- Setup Ruby on Rails project

## Built with
- [GoLang](https://golang.org/)

## Features
- Start ruby project
- Start ruby on rails project

## Installation
- [Click here](https://github.com/descholar-ceo/boiler/releases/tag/v1.0.0-beta.3) and download a zip file
- Extract it
- Open terminal in the extracted folder, and run `./install.sh`
- Close your current terminal and open another one and run `boiler` ===> there you go!

### Usage
#### Pre-requisites
##### For ruby projects
- [ ] Ruby should be installed already on your computer
- [ ] Bundler

##### For Ruby on Rails project
To get your project initiated, you should have the following tools already installed on your computer:
- [ ] Ruby (Installed and well working)
- [ ] Ruby on rails (Installed and well working)
- [ ] Yarn
- [ ] Bundler

 # NOTE: DATABASES REQUIREMENTS
1. MySQL: requires you to have pre-installed `libmysqlclient-dev`, on your system, if you are not sure run `sudo apt-get install libmysqlclient-dev` which will install it if you don't have it or update it if the version you have is obsolete

1. Oracle: requires you to have ruby-oci8 installed, and to install it, requires you to have oracle environment installed properly on your computer

1. [Frontbase](http://www.frontbase.com/cgi-bin/WebObjects/FBWebSite): requires you to be a MacOS user

1. [ibm_db](https://www.ibm.com/support/knowledgecenter/hr/SSEPGG_9.7.0/com.ibm.db2.luw.qb.server.doc/doc/t0008875.html): requires you to have `db2` server installed on your computer

1. Sqlserver: requires you to have pre-installed `freeTDS` on your system, 
if you are not sure run `sudo apt-get install freetds-dev` which will install it if you don't have it or update it if the version you have is obsolete

1. jdbcmysql, jdbcsqlite3, jdbcpostgresql,: are for use with [JRuby](https://www.jruby.org/) only

#### Running
After you install successfully, run `boiler` on your terminal and follow the instructions

## Contributions
### NOTE: Please follow [this contribution guide](https://github.com/descholar-ceo/boiler/blob/make-a-new-release/CONTRIBUTING.md)
