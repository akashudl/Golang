# GO Windows Service Shell

This is based on the [GO Windows service example program](https://godoc.org/golang.org/x/sys/windows/svc/example) provided by the GO Project. 
It is a project shell to create a Windows service.

## Getting Started

The program compiles and runs on GO 1.8.  The generated executable accepts a single parameter.  The parameter values include:

* start


## Installing and Updating a Service

After compiling an executable, the service can be installed from an Administrative command prompt.  Typing

    YourExecutable.EXE install 

will install the service.

To update the service, stop the service, replace the executble and restart the service.

The service can be removed from an Administrative command prompt by typing:

    YourExecutable.EXE remove 

## Customizing

The code exists in two packabages

* cmd/gosvc - Wrapper to control the service
* app - Your application

All service boilerplate code is in the four files in `cmd/gosvc` with a "svc_" prefix.  There should
be no need to modify this code.

The only code you should need to change is in main.go.


You should also rename the `gosvc` directory to the name of your executable.










