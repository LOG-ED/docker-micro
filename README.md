# Docker and Golang Microservices
### Corso introduttivo su Docker e Microservizi in Golang

Questo è il repository ufficiale del corso. E' possibile consultare la scheda e i materiali del corso a questo indirizzo: 
**https://log-ed.github.io/docker-micro**

### Docente: Federico Minzoni  
- fminzoni@enter.eu  
- https://github.com/f-minzoni  

Sono un programmatore e smanettone della prima ora. Appassionato da sempre di Git, MongoDB, Rich Internet Application, IOT, Cloud Computing, da diversi mesi ho aggiunto alla lista, Docker e le architetture a Microservizi. Due temi che, insieme, stanno rivoluzionando l'intero processo di sviluppo e rilascio delle applicazioni sul Cloud. Ambito in cui lavoro da 3 anni, in Enter, occupandomi della piattaforma [Enter Cloud Suite](http://www.entercloudsuite.com).

### Avvertenze per lo svolgimento e l'utilizzo di questo corso

Questa tipologia di corso, definita e organizzata da [LOG.ED](https://loged.it) è strutturata in modo tale che tutti i materiali e le esercitazioni siano pubblici e accessibili liberamente. Si prega di rispettare le condizioni di utilizzo e/o licenza di qualsiasi codice qui presente. Per ogni algoritmo o codice che si desidera riutilizzare, si richiede di indicare nei crediti la fonte originale, con un commento in linea.

### Setup (Requirements)

1. Download the protol buffers compiler from https://github.com/google/protobuf/releases   

2. Run the following command to install the Go protocol buffers plugin:   
   `$ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}` 

3. If changes are made to the Protocol Buffer file use the following command to regenerate:  
   `$ protoc -I$GOPATH/src --go_out=plugins=micro:$GOPATH/src \ $GOPATH/src/github.com/log-ed/docker-micro/proto/task.proto`


