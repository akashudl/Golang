package main

import (
	"fmt"
	"golang.org/x/sys/windows/svc"

	"golang.org/x/sys/windows/svc/mgr"

)

func startService(name string) error {
	m, err := mgr.Connect() //connect with mg package to manage the Services 
	if err != nil {
		return err
	}
	defer m.Disconnect()
	s, err := m.OpenService(name) //then we have to open That Service 
	if err != nil {
		return fmt.Errorf("could not access service: %v", err)
	}
	defer s.Close()
	err = s.Start("is", "manual-started")//this function is used for starting the Services
	if err != nil {
		return fmt.Errorf("could not start service: %v", err)
	}
	return nil
}

func main() {
	svcName:="Akashservice"
	m, err := mgr.Connect() //first we have to conncet with Window Services 
	if err != nil {
		fmt.Println(err)
	}
	defer m.Disconnect()
	s, err := m.OpenService(svcName) //then we have to open That Service 
	if err != nil {
		fmt.Printf("could not access service: %v", err)
	}
	defer s.Close();
	status, err := s.Query() //Checking the status of Service if status is  1 it is Stopped if 4 it is Running 
	if err != nil {
	fmt.Println("Error", err)
	}else{
		fmt.Println("Service Status")
		fmt.Println(status.State);
		if status.State==svc.Stopped{
			startService(svcName); //this Function is used for starting the service
		}
	}
	    


}

