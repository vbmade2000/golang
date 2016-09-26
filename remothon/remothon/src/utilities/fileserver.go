package utilities

import (
        "net/http"
        "os/user"
)

/*
   Function to start file server and points it to home directory 
   of logged in user
*/
func StartFileServer() {
    port := "9451"
    currentUser, _ := user.Current()    
    http.ListenAndServe(":" + port, http.FileServer(http.Dir(currentUser.HomeDir)))
}

