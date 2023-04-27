Students system ( client-server with go and mysql )


https://user-images.githubusercontent.com/69214737/233971353-898a564b-0084-4967-8b83-aa1f83942f47.mp4



Each supposed to work on separate device
1. master
2. 2 clients
3. 2 slaves

![image](https://user-images.githubusercontent.com/69214737/233971425-27a1eae6-6b0c-4226-ab43-6fd31a48c8a2.png)


- Both of clients send requests to the master with student data they needs, wether it grade1 or grade2
- Master based on requested data will send request to the slave that data stored on
- Slave which is connected to sql database, will send response to the client with data
