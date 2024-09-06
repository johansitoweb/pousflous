package main  

import (  
    "encoding/json"  
    "net/http"  
)  

type Response struct {  
    Mensaje string `json:"mensaje"`  
}  

func mensajeHandler(w http.ResponseWriter, r *http.Request) {  
    response := Response{Mensaje: "¡Hola desde el servidor Go!"}  
    w.Header().Set("Content-Type", "application/json")  
    json.NewEncoder(w).Encode(response)  
}  

func main() {  
    http.HandleFunc("/api/mensaje", mensajeHandler)  
    http.ListenAndServe(":8080", nil)  
}