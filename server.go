package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/julienschmidt/httprouter"
    emailVerifier "github.com/reacherhq/check-if-email-exists" // Import the email verifier package
)

// VerifyEmail function to handle email verification
func VerifyEmail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    email := ps.ByName("email")
    exists := emailVerifier.CheckIfEmailExists(email) // Use the check function from the repo
    
    if exists {
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "Email %s exists.", email)
    } else {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintf(w, "Email %s does not exist.", email)
    }
}

func main() {
    router := httprouter.New()
    router.GET("/verify/:email", VerifyEmail) // Define the endpoint

    log.Fatal(http.ListenAndServe(":8080", router)) // Start the server
}
