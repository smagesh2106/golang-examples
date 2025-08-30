all:
	# Generate the private key for CA
	openssl genpkey -algorithm RSA -out ca.key -aes256

	# Create the self-signed certificate (CA)
	openssl req -new -x509 -key ca.key -out ca.crt -days 3650

	# Generate the server private key
	openssl genpkey -algorithm RSA -out server.key

	# Create the server's CSR
	openssl req -new -key server.key -out server.csr

	# Sign the server CSR with the CA
	openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365

	# Generate the client private key
	openssl genpkey -algorithm RSA -out client.key

	# Create the client CSR
	openssl req -new -key client.key -out client.csr

	# Sign the client CSR with the CA
	openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 365


