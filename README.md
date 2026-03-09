# Read Me

__This is work in progress__

Small util to find executables. A modern version of an ancient (1995) util made with C++ in my school years.  
  
I aim for this to build and work on most platforms, but it's mostly tested on MacOs and Linux.  
  
The main purpose of this project is to learn more about Go in general, and also dust off the old Makefile skills i used to have back in the day :)
  
## Build
make build - build for current os/platform

Set GOOS and GOARCH to build for another. Eg.  
- make build GOOS=windows GOARCH=amd64  
- make build GOOS=linux GOARCH=amd64  

Check for compatibility for your platform:
- make prereq (also run by make build)