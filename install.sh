#!/bin/bash
echo "
Welcome to Boiler package developed by descholar. This package was developed to help you make your life very easy 
while working on the next project"

mkdir $HOME/.boiler $HOME/.boiler/boiler  && cd $HOME/.boiler/boiler

curl -L -O https://raw.githubusercontent.com/descholar-ceo/boiler/develop/main.go
curl -L -O https://raw.githubusercontent.com/descholar-ceo/boiler/develop/Makefile

mkdir lib lib/.ruby && cd lib/.ruby

curl -L -O https://raw.githubusercontent.com/descholar-ceo/boiler/develop/lib/.ruby/.rubocop.yml
curl -L -O https://raw.githubusercontent.com/descholar-ceo/boiler/develop/lib/.ruby/README.md

mkdir .github .github/workflows && cd .github/workflows

curl -L -O https://raw.githubusercontent.com/descholar-ceo/boiler/develop/lib/.ruby/.github/workflows/linters.yml
curl -L -O https://raw.githubusercontent.com/descholar-ceo/boiler/develop/lib/.ruby/.github/workflows/tests.yml

cd $HOME/.boiler/boiler

make build

sudo echo PATH=$PATH:$HOME/.boiler/boiler/bin>>$HOME/.bashrc

echo "Everything is set now"
