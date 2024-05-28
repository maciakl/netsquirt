package main

import (
    "os"
    "fmt"
    "net"
    "flag"
    "net/http"
    "path/filepath"
)

const version = "0.1.1"

func main() {

    // command line args
    var file string
    flag.StringVar(&file, "file", "", "path to file to send")
    var ver bool
    flag.BoolVar(&ver, "version", false, "display version number")
    flag.Parse()

    if ver {
        fmt.Println("netsquirt version", version)
        os.Exit(0)
    }

    ip := getIP()
    fmt.Println("\nServer running on: http://"+ ip)
    fmt.Println("Press Ctrl+C to stop the server")

    filename := filepath.Base(file)

    // check if the file exists and is not a directory
    if fileinfo, err := os.Stat(file); err == nil && !fileinfo.IsDir(){
    
        fmt.Println("\nServing file: ", filename, "on port 80")
        fmt.Println("\nTo get the file:")
        fmt.Println("\t - Paste this address into your browser: http://"+ ip)
        fmt.Println("\t - Use the following command: wget http://"+ ip)
        fmt.Println("\t - Or on Windows (Powershell): wget http://"+ ip, "-O", filename)

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
            fmt.Fprint(w, "ip address: ", ip, "\n")
        }
    })


    if err := http.ListenAndServe(":80", nil); err != nil {
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

    // iterate through the interfaces and return the first non-loopback address
    for _, address := range addrs {
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }

    return ""
}
