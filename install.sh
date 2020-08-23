#!/bin/bash
echo "
Welcome to Boiler package developed by descholar. This package was developed to help you make your life very easy 
while working on the next project"

mkdir $HOME/.boiler && cd $HOME/.boiler && git clone https://github.com/descholar-ceo/boiler

make build

sudo echo PATH=$PATH:$HOME/boiler/boiler/bin>>$HOME/.bashrc

echo "Everything is set now"
