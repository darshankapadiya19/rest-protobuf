**Testing RESTful APIs with protobuf encoding**

Create binary data file
```shell
echo -n "binary Darshan" | xxd -r -p > data.bin
```

Add data.bin file as binary input

add following header
```http request
Content-Type:application/x-binary
```

Make Http Request
```http request
http://localhost:8080/hello
```

**Postman setup example:**

ayHello using binary/protobuf over http/1.x
![Example sayHello using binary/protobuf over http/1.x](img/sayHello-grpc.png)