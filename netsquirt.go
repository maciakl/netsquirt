package main

import (
    "os"
    "fmt"
    "net"
    "flag"
    "net/http"
    "path/filepath"
)

const version = "0.1.2"

func main() {

    // command line args
    var file string
    flag.StringVar(&file, "file", "", "path to file to send (can be relative or absolute)")

    var port int
    flag.IntVar(&port, "port", 80, "port to run the server on")

    var ver bool
    flag.BoolVar(&ver, "version", false, "display version number and exit")

    flag.Parse()

    if ver {
        fmt.Println("netsquirt version", version)
        os.Exit(0)
    }

    portnum := "" // port number string

    if port != 80 {
        portnum = fmt.Sprintf(":%d", port)
    }

    ip := getIP()
    fmt.Println("\nServer running on: http://"+ ip + portnum)
    fmt.Println("Press Ctrl+C to stop the server")

    filename := filepath.Base(file)

    // check if the file exists and is not a directory
    if fileinfo, err := os.Stat(file); err == nil && !fileinfo.IsDir(){
    
        fmt.Println("\nServing file: ", filename, "on port 80")
        fmt.Println("\nTo get the file:")
        fmt.Println("\t - Paste this address into your browser: http://"+ ip + portnum)
        fmt.Println("\t - Use the following command: wget http://"+ ip + portnum)
        fmt.Println("\t - Or on Windows (Powershell): wget http://"+ ip + portnum, "-O", filename)

    } else {
        if file != "" {
            fmt.Fprintln(os.Stderr, "File", file , "does not exist or is not a file")
            os.Exit(1)
        }
    }


    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        if file != "" {

            // need to set the content type header so that the bowser knows the file name and type
            w.Header().Set("Content-Disposition", "attachment; filename="+filename)
            http.ServeFile(w, r, file)


        } else {
            // this happens when the file is not set
            fmt.Fprint(w, "netsquirt version ", version, "\n")
            fmt.Fprint(w, "server running on: http://"+ ip + portnum, "\n")
        }
    })


    if err := http.ListenAndServe(portnum, nil); err != nil {
        fmt.Fprintln(os.Stderr, "Server error:", err)
    }
}






// get the ip address of the machine
func getIP() string {

    // get list of interfaces and panic if there is an error
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        panic(err)
    }

    // iterate through the interfaces and return the first private address
    for _, address := range addrs {
        if ipnet, ok := address.(*net.IPNet); ok && ipnet.IP.IsPrivate(){
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }

    return ""
}
