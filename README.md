# Origin Take Home

## Start
To help the interaction with the source code a makefile was made, so you can run and test
the application more easily, as if it were hard with go :-) . The makefile is a common practice
in Linux and can be used without any effort, in macOS it will ask you to install the Apple 
Developer Tools if you haven't yet. If you are on Windows, deserves to be punished and run
all commands by hand. 

A Dockerfile was created to not force you to install Go on your machine, so you can run only the
container if you want, but the unit tests runs only locally and will need the Go installation.

To install Go: https://golang.org/doc/install

To use the makefile you need just to run ```make <command>``` in the project root. Example ```make run-docker```.

### Commands

```run```: runs the application on your machine

```test```: runs the unit tests on your machine

```build```: builds the application Docker image

```run-docker```: runs the application using the built Docker image (will always build the image)

### Windows
If you are a sad person you can run the make commands equivalent.

```run```: ```go run src/cmd/api/main.go```

```test```: ```go test ./...```

```build```: ```docker build --tag origin-take-home .```

```run-docker```: ```docker run --publish 8080:8080 origin-take-home``` (remember to build the image before)

### Quick cURL
```
curl --request POST \
--url http://localhost:8080/user/risk \
--header 'Content-Type: application/json' \
--data '{
"age": 35,
"dependents": 2,
"house": {"ownership_status": "owned"},
"income": 0,
"marital_status": "married",
"risk_questions": [0, 1, 0],
"vehicle": {"year": 2018}
}'
```

## Technical Stuff

If you are not used to Go principles (https://golang.org/doc/faq#principles) and comes from more 
corporate languages ( C# / Java ) you may find the code a little strange. Go is heavily oriented
by simplicity and removes a lot of the abstraction tools to achieve that goal. Besides the 
community conventions, maybe the more difficult concept to understand is the duck typing behavior
(https://golang.org/doc/faq#implements_interface).

I tried my best to find a middle ground implementation between a go-standard application and a more usual 
corporate approach to make the code organization more friendly (or less strange) to people that could come 
from other languages.

The engine was created with a deterministic idea behind. So, does not matter which order you run the rules
the result should always be the same. The extensibility is easily achieved by adding a new rule to the engine
according to the risk rule interface. To help keep the deterministic behavior I followed some of the Functional
Programming principles, trying to usually create pure functions and thus working with an anemic model.

The unit tests were done only on the core layer, where the business is really contained. 

### Source Organization

The source code was organized taking into consideration mostly the ideas of decoupling from the
Clean Architecture (https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

.```adapter```: Conversions between the infra world and the core world.  
.```cmd```: All executable entrypoint and where de dependency injection is usually done (Go standard)   
.```controller```: Exposes a core service.  
.```core```: Where the business model is centered.  
&nbsp;&nbsp;&nbsp;&nbsp; .```entity```: All entities of the application. Following the anemic model.  
&nbsp;&nbsp;&nbsp;&nbsp; .```riskengine```: A parallel between the Use Cases layer but concentrating "risk rules" cases.  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; .```rule```: All the "risk rules" that can be used in the risk engine.  
.```infra```: Where external communications and technologies are centered.  

## Final Comments

Unfortunately I didn't have all the time that I wanted and needed to apply in the project so some "details" are missing.
I didn't add a middleware or validator layer to guarantee the perfect consistency of the API inputs, so only
the static contract binding is being verified. And I couldn't add integration tests following BDD. Some parameters
(port for example) should have being mapped to be brought from the environment. 
