# law-assignment2

## Task Description

1. File compression service: accepts input file and output from compression

2. File storage service: accepts input files to be saved and can be accessed at
one URL

   - Upload file

   - Download the URL to download

3. CRUD metadata storage service (at least there should be 4 attributes for each
metadata record)

   - Create, Read, Update, Delete

## Communication Details

Each service request require OAuth's Access Token as `Authorization` header.

Example for getting access token using Infralabs Oauth (http://oauth.infralabs.cs.ui.ac.id/oauth/token):

- Request:

  This require `username`, `password`, `grant_type`, `client_id`, and `client_secret`

  ```http
  POST /oauth/token HTTP/1.1
  Host: oauth.infralabs.cs.ui.ac.id
  Content-Type: application/x-www-form-urlencoded
  
  username=dummy&password=thepassword&grant_type=password&client_id=7kw6hrzffaw41k6j286l1zjt13fwk322366dwv0j&client_secret=ql0s53emfoy9ckyzkwi3u2jg9fhgsm8gsi9yj2pv
  ```

- Response

  ```json
  {
      "access_token": "66n7yp9bljrutuh75scjogidrwyyppkd5whhl0k0",
      "expires_in": 3600,
      "token_type": "Bearer",
      "scope": null,
      "refresh_token": "k3og0py5jua0sbk4clrtshye6imfsp8hyxcgiwdl"
  }
  ```
  
Using the `access_token`, the following requests should have this header:

```
Authorization: Bearer 66n7yp9bljrutuh75scjogidrwyyppkd5whhl0k0
```

All invalid authentication token will be replied with `401 Unauthorized`

### Compression Sevice

- Request

  The file should be uploaded as `myFile` field.

  ```http
  POST /compress/ HTTP/1.1
  Host: localhost:8887
  Authorization: Bearer 66n7yp9bljrutuh75scjogidrwyyppkd5whhl0k0
  Content-Type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW

  ------WebKitFormBoundary7MA4YWxkTrZu0gW
  Content-Disposition: form-data; name="myFile"; filename="42W23RWXW6.txt"
  Content-Type: text/plain


  ------WebKitFormBoundary7MA4YWxkTrZu0gW--
  ```

- Response

  The response is the gzip compressed file. You can uncompress it using any gzip tools to view the original data

### File Storage Service

There are two endpoints in this service: one for uploading the file, and one for downloading it. Both are depending
on the [Metadata Service](#metadata-service) for mantaining the files.

#### Upload

This will store the file (by key `myFile`) with the authorized user as its owner.

- Request

  ```http
  POST /upload/ HTTP/1.1
  Host: localhost:8889
  Authorization: Bearer 66n7yp9bljrutuh75scjogidrwyyppkd5whhl0k0
  Content-Type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW

  ------WebKitFormBoundary7MA4YWxkTrZu0gW
  Content-Disposition: form-data; name="myFile"; filename="Juan.png"
  Content-Type: image/png


  ------WebKitFormBoundary7MA4YWxkTrZu0gW--
  ```

#### Download

Before a user can download a file, it checks if he/she owns the file from [Metadata List File Service](#list-file).

- Request

  ```http
  GET /download/?id=31 HTTP/1.1
  Host: localhost:8889
  Authorization: Bearer 66n7yp9bljrutuh75scjogidrwyyppkd5whhl0k0
  ```

### Metadata Service

#### Upload Data

This endpoint should be called by [File Storage Upload Service](#upload), each time an user upload a file.

- Request

  ```http
  POST /upload HTTP/1.1
  Host: localhost:8888
  Authorization: Bearer 66n7yp9bljrutuh75scjogidrwyyppkd5whhl0k0
  Content-Type: application/json

  {
    "FileName": "pic1",
    "Owner": "ashlah",
    "Path": "blablabla/blehbleh"
  }
  ```

#### List File

Return the lists of files that owned by the authorized user.

- Request

  ```http
  GET /files/ HTTP/1.1
  Host: localhost:8888
  Authorization: Bearer 66n7yp9bljrutuh75scjogidrwyyppkd5whhl0k0
  ```

- Response

  ```json
  [
      {
          "Id": 31,
          "FileName": "Ash_SM.png",
          "Owner": "1606895884",
          "Path": "files/524317775-Ash_SM.png",
          "Timestamp": "2020-02-28T15:16:15.913481Z"
      },
      {
          "Id": 32,
          "FileName": "Juan.png",
          "Owner": "1606895884",
          "Path": "files/998134922-Juan.png",
          "Timestamp": "2020-03-04T23:07:03.358688Z"
      }
  ]
  ```

It can also directly get a file metadata based on its ID

- Request

  ```http
  GET /files/?id=31 HTTP/1.1
  Host: localhost:8888
  Authorization: Bearer 66n7yp9bljrutuh75scjogidrwyyppkd5whhl0k0
  ```

- Response

  ```json
  {
      "Id": 31,
      "FileName": "Ash_SM.png",
      "Owner": "1606895884",
      "Path": "files/524317775-Ash_SM.png",
      "Timestamp": "2020-02-28T15:16:15.913481Z"
  }
  ```
