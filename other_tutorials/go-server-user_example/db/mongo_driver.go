package swagger

import (
	"bytes"

	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Collection *mongo.Collection
var Log *log.Logger

/*
 * Init function to initialize mongo DB connection.
 */
func Init_Mongo() error {

	Log = log.New(os.Stdout, "iam-policy-administration :", log.LstdFlags)
	Log.Println("mongo basics")
	_certBasePath := "/home/magesh/02_Work/Openssl/"

	tlsConfig := new(tls.Config)
	tlsConfig.InsecureSkipVerify = false
	//<FIXME> replace key, cert path from env .
	key := _certBasePath + "client/cl.key"
	cacert := _certBasePath + "ca/ca.crt"
	cert := _certBasePath + "client/cl.crt"
	sub, err := AddClientCertFromSeparateFiles(tlsConfig, key, cacert, cert, "")

	//sub, err := mgo.AddClientCertFromConcatenatedFile(tlsConfig, keyCert, "")
	if err != nil {
		Log.Printf("Error while reading certificates, %v", err)
		return err
	} else {
		Log.Printf("certificates added...%v\n", sub)
	}
	Log.Printf("Verify Certificate : %s\n", strconv.FormatBool(!tlsConfig.InsecureSkipVerify))

	//<FIXME include replicaset in the production cluster>
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/?ssl=true").
		SetAuth(options.Credential{
			AuthSource: "magesh-mongo-test", Username: "user1", Password: "passw0rd",
		}).SetTLSConfig(tlsConfig) //.SetReplicaSet("rs0")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	Client, _ = mongo.Connect(ctx, clientOptions)
	err = Client.Ping(ctx, nil)
	if err != nil {
		Log.Printf("mongo connection error %v", err)
		return err
	}

	Collection = Client.Database("magesh-mongo-test").Collection("people")
	Log.Println("done init ....")
	return nil
}

func AddClientCertFromConcatenatedFile(cfg *tls.Config, certKeyFile, keyPassword string) (string, error) {
	data, err := ioutil.ReadFile(certKeyFile)
	if err != nil {
		return "", err
	}

	return addClientCertFromBytes(cfg, data, keyPassword)
}
func AddClientCertFromSeparateFiles(cfg *tls.Config, keyFile, cacert, certFile, keyPassword string) (string, error) {

	err := addCACertFromFile(cfg, cacert)
	if err != nil {
		Log.Printf("Error while reading CA cert from file :%v", err)
		return "", err
	}

	keyData, err := ioutil.ReadFile(keyFile)
	if err != nil {
		Log.Printf("Error whiel reading key file: %v", err)
		return "", err
	}
	certData, err := ioutil.ReadFile(certFile)
	if err != nil {
		Log.Printf("Error while reading cert data: %v", err)
		return "", err
	}

	data := append(keyData, '\n')
	data = append(data, certData...)
	return addClientCertFromBytes(cfg, data, keyPassword)
}

func addCACertFromFile(cfg *tls.Config, file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		Log.Printf("Error while reading CA cert from file: %v", err)
		return err
	}

	certBytes, err := loadCACert(data)
	if err != nil {
		Log.Printf("Error while loading CA cert: %v", err)
		return err
	}

	cert, err := x509.ParseCertificate(certBytes)
	if err != nil {
		Log.Printf("Error parsing certificate: %v", err)
		return err
	}

	if cfg.RootCAs == nil {
		cfg.RootCAs = x509.NewCertPool()
	}

	cfg.RootCAs.AddCert(cert)

	return nil
}

func loadCACert(data []byte) ([]byte, error) {
	var certBlock *pem.Block

	for certBlock == nil {
		if data == nil || len(data) == 0 {
			return nil, errors.New("no CERTIFICATE section found")
		}

		block, rest := pem.Decode(data)
		if block == nil {
			return nil, errors.New("invalid .pem file")
		}

		switch block.Type {
		case "CERTIFICATE":
			certBlock = block
		}
		data = rest
	}
	return certBlock.Bytes, nil
}

func addClientCertFromBytes(cfg *tls.Config, data []byte, keyPasswd string) (string, error) {
	var currentBlock *pem.Block
	var certBlock, certDecodedBlock, keyBlock []byte

	remaining := data
	start := 0
	for {
		currentBlock, remaining = pem.Decode(remaining)
		if currentBlock == nil {
			break
		}

		if currentBlock.Type == "CERTIFICATE" {
			certBlock = data[start : len(data)-len(remaining)]
			certDecodedBlock = currentBlock.Bytes
			start += len(certBlock)
		} else if strings.HasSuffix(currentBlock.Type, "PRIVATE KEY") {
			//<FIXME> if the key file has a password, replace "" with password
			if keyPasswd != "" && x509.IsEncryptedPEMBlock(currentBlock) {
				var encoded bytes.Buffer
				buf, err := x509.DecryptPEMBlock(currentBlock, []byte(keyPasswd))
				if err != nil {
					Log.Printf("Error while decrypting a key using password: %v", err)
					return "", err
				}

				pem.Encode(&encoded, &pem.Block{Type: currentBlock.Type, Bytes: buf})
				keyBlock = encoded.Bytes()
				start = len(data) - len(remaining)
			} else {
				keyBlock = data[start : len(data)-len(remaining)]
				start += len(keyBlock)
			}
		}
	}
	if len(certBlock) == 0 {
		return "", fmt.Errorf("failed to find CERTIFICATE")
	}
	if len(keyBlock) == 0 {
		return "", fmt.Errorf("failed to find PRIVATE KEY")
	}

	cert, err := tls.X509KeyPair(certBlock, keyBlock)
	if err != nil {
		Log.Printf("Error while creating a certificate: %v", err)
		return "", err
	}

	cfg.Certificates = append(cfg.Certificates, cert)
	crt, err := x509.ParseCertificate(certDecodedBlock)
	if err != nil {
		Log.Printf("Error Parsing Certificate: %v", err)
		return "", err
	}
	return crt.Subject.String(), nil
}
