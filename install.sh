#!/bin/bash

echo -e "\n
Welcome to Boiler package developed by descholar. This package was developed to help you make your life very easy 
while working on the next project\n"

# CHECKING IF THE .boiler DIRECTORY IS PRESENT IN HOME DIR
if [ -d "$HOME/.boiler" ]
then
    # boiler is already installed, so remove it and recreate it
    rm -r $HOME/.boiler
    mkdir $HOME/.boiler $HOME/.boiler/boiler 
else
    # boiler is not installed yet so create the dir
    mkdir $HOME/.boiler $HOME/.boiler/boiler
fi

# cd $HOME/.boiler/boiler

# curl -L -O https://raw.githubusercontent.com/descholar-ceo/boiler/develop/main.go
# copy all to the installation folder
echo "Copying installtion files ... "
# cp -r . $HOME/.boiler/boiler
cp -r `ls | egrep -v '^.git$'` $HOME/.boiler/boiler
# curl -L -O https://raw.githubusercontent.com/descholar-ceo/boiler/develop/Makefile

# mkdir lib lib/.ruby

# cd lib/.ruby

# curl -L -O https://raw.githubusercontent.com/descholar-ceo/boiler/develop/lib/.ruby/.rubocop.yml
# curl -L -O https://raw.githubusercontent.com/descholar-ceo/boiler/develop/lib/.ruby/README.md

# mkdir .github .github/workflows

# cd .github/workflows

# curl -L -O https://raw.githubusercontent.com/descholar-ceo/boiler/develop/lib/.ruby/.github/workflows/linters.yml
# curl -L -O https://raw.githubusercontent.com/descholar-ceo/boiler/develop/lib/.ruby/.github/workflows/tests.yml

cd $HOME/.boiler/boiler

make build

sudo echo PATH=$PATH:$HOME/.boiler/boiler/bin>>$HOME/.bashrc

echo -e "\nEverything is set now\n\nTo use boiler, type boiler in your terminal and follow the instructions."
