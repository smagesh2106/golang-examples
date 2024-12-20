package main

import (
	//"crypto/tls"
	mgo "examples/go-server-user_example/db"
	sw "examples/go-server-user_example/go"
	"log"
	"net/http"
	"time"
)

/*
 * Main function to start server and listen.
 */
func main() {
	label := false

	//Initialize mongodb and start.
	for {
		if label {
			break
		}
		err := mgo.Init_Mongo()
		if err != nil {
			log.Printf("Error setting up mongoDB :%v", err)
			time.Sleep(5 * time.Second)
		} else {
			label = true
		}
	}
	router := sw.NewRouter()
	/*
		cfg := &tls.Config{
			MinVersion: tls.VersionTLS12,
			//CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			//PreferServerCipherSuites: true,
			//CipherSuites: []uint16{
			//	tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			//	tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			//	tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			//	tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			//},
		}

		srv := &http.Server{
			Addr:         ":8443",
			Handler:      router,
			TLSConfig:    cfg,
			TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
		}
		mgo.Log.Printf("Server started")
	*/
	mgo.Log.Fatal(http.ListenAndServe(":8080", router))
	/*
		_certBasePath := "/home/magesh/02_Work/Openssl/"

		key := _certBasePath + "client/cl.key"
		//cacert := _certBasePath + "ca/ca.crt"
		cert := _certBasePath + "client/cl.crt"
		mgo.Log.Fatal(srv.ListenAndServeTLS(cert, key))
	*/
}

//sudo mongod --tlsMode requireTLS --dbpath /var/lib/mongodb --auth --tlsCertificateKeyFile /home/magesh/02_Work/Openssl/server/server.pem --tlsCAFile /home/magesh/02_Work/Openssl/ca/ca.crt
//./mongobasic
