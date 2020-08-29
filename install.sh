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

echo "Copying installtion files ... "

cp -r `ls | egrep -v '^.git$'` $HOME/.boiler/boiler

cd $HOME/.boiler/boiler

sudo echo PATH=$PATH:$HOME/.boiler/boiler/bin>>$HOME/.bashrc

echo -e "\nEverything is set now\n\nTo use boiler, type boiler in your terminal and follow the instructions.\n"

source ~/.bashrc

sleep 1s

exit
