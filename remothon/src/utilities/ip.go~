package utilities

import (
        "net"
        "os"
)

func GetIPs()([]string, error){
  
   var IPs []string
   IPs = make([]string, 0, 0)

   addrs, err := net.InterfaceAddrs()
   if err != nil {
      os.Stderr.WriteString("Oops: " + err.Error() + "\n")
      os.Exit(1)
   }

   for _, a := range addrs {
	if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
		 IPs = append(IPs, ipnet.IP.String())
	}
   }
   
   return IPs, err
}

