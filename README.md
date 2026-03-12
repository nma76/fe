[![CI](https://github.com/nma76/fe/actions/workflows/ci.yml/badge.svg)](https://github.com/nma76/fe/actions/workflows/ci.yml)

# Read Me

__This is work in progress__

Small util to find executables. A modern version of an ancient (1995) util made with C++ in my school years.  
  
I aim for this to build and work on most platforms, but it's mostly tested on MacOs and Linux.  
  
The main purpose of this project is to learn more about Go in general, and also dust off the old Makefile skills i used to have back in the day :)
  
## Build
make build - build for current os/platform  
make release - build for all supported os/platforms  

### Build for a specific supported os
make build-linux  
make build-darwin  
make build-windows  

## Install / uninstall
make install - install to /usr/local/bin  
make uninstall

## How to use:
Usage: __fe [options] [directory]__
  
Options:  
 -h, --help              Show this help message  
 -i, --icon              Show icon for each file  
 -d, --debug             Enable debug output  
