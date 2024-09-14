package main

import (
	pb "github.com/darshankapadiya19/rest-protobuf/proto/gen"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type HaloHandler struct{}

func (h *HaloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Unable to read message from request : %v", err)
	}

	// Notice the declaration of the variable here
	// Alternatively we can do the following
	// 		req := &pb.HelloRequest{}
	// 		proto.Unmarshal(data, req)
	var req pb.HelloRequest
	err = proto.Unmarshal(data, &req)
	if err != nil {
		log.Fatalf("Unable to unmarshal message from request : %v", err)
	}
	log.Printf("Halo request form %s", req.Name)

	resp := &pb.HelloResponse{
		Message: "Halo " + req.Name + ", Majama?",
	}

	response, err := proto.Marshal(resp)
	if err != nil {
		log.Fatalf("Unable to marshal response : %v", err)
	}

	w.Write(response)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received.")
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Unable to read message from request : %v", err)
	}

	req := &pb.HelloRequest{}
	err = proto.Unmarshal(data, req)
	if err != nil {
		log.Fatalf("Unable to unmarshal message from request : %v", err)
	}
	log.Printf("Hello request form %s", req.Name)

	resp := &pb.HelloResponse{
		Message: "Hello " + req.Name + ", How's it going?",
	}
	response, err := proto.Marshal(resp)
	if err != nil {
		log.Fatalf("Unable to marshal response : %v", err)
	}
	w.Write(response)
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/hello", helloHandler).Methods("POST")

	haloHandler := &HaloHandler{}
	router.Handle("/halo", haloHandler).Methods("POST")

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Server is running and listening on port 8080")
	log.Fatal(server.ListenAndServe())
}
