/*
------------- HTTP ----------------

- HTTP is the foundation to learning package WebDevelopment.
- HTTP is a stateless , text based  request response protocol that uses
the client server computing model.
- the client is also known as the user-agent
and is often a web browser. The server is often called the web server.
■ GET—Tells the server to return the specified resource.
■ HEAD—The same as GET except that the server must not return a message
body. This method is often used to get the response headers without carrying
the weight of the rest of the message body over the network.
■ POST—Tells the server that the data in the message body should be passed to
the resource identified by the URI. What the server does with the message body
is up to the server.
■ PUT—Tells the server that the data in the message body should be the resource
at the given URI. If data already exists at the resource identified by the URI, that
data is replaced. Otherwise, a new resource is created at the place where the
URI is.
■ DELETE—Tells the server to remove the resource identified by the URI.
■ TRACE—Tells the server to return the request. This way, the client can see what
the intermediate servers did to the request.
■ OPTIONS—Tells the server to return a list of HTTP methods that the server supports.
■ CONNECT—Tells the server to set up a network connection with the client. This
method is used mostly for setting up SSL tunneling (to enable HTTPS).
■ PATCH—Tells the server that the data in the message body modifies the
resource identified by the URI. 
 */