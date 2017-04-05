---
layout: default
---

## Descrizione  
Il corso propone un percorso formativo che mescola due progetti open source di successo: il linguaggio di programmazione Golang e gli strumenti di Docker. Attraverso esercitazioni pratiche e un pizzico di teoria (q.b) vengono introdotti i concetti peculiari di Golang e di Docker per lavorare su architetture a Microservizi, con un linguaggio di sviluppo moderno e un efficace processo di distribuzione del software.  

### Livello: Base

### Durata: 4 sessioni
 
### Svolgimento delle sessioni:  
Le sessioni on site sono previste nei giorni:  

- **mercoledì 22 marzo dalle 18:30 alle 20:30.**  
- **mercoledì 29 marzo dalle 18:30 alle 20:30.** 
- **mercoledì 5 aprile dalle 18:30 alle 20:30.** 
- **mercoledì 12 aprile dalle 18:30 alle 20:30.** 

Per la natura del corso è consigliato frequentare tutte le sessioni.
Ogni sessione, della durata di 2 ore, prevede questi momenti formativi:

- una presentazione in cui vengono introdotti e descritti i nuovi concetti  
- una dimostrazione pratica degli strumenti e dei comandi descritti  
- l'assegnazione di esercitazioni pratiche individuali o a piccoli gruppi  
- Un momento di revisione delle esercitazioni assegnate

### Tools:

- Le esercitazioni pratiche sono svolte utilizzando la piattaforma [GitHub Classroom](https://classroom.github.com){:target="_blank"} che consente di creare un ambiente di lavoro dedicato per ogni studente e semplifica la verifica. **Richiedi l'accesso al Docente!**    
- E' presente una chat dedicata per il corso [![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/LOG-ED/docker-micro?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)  
- L'editor consigliato è [Visual Studio Code](https://code.visualstudio.com/){:target="_blank"}. 

### Requisiti:
 
E' indispensabile una conoscenza base di programmazione.  
Porta il tuo pc, il corso è molto pratico!

## Docente

### Federico Minzoni

- fminzoni@enter.eu 
- https://github.com/f-minzoni

Sono un programmatore e smanettone della prima ora. Appassionato da sempre di Git, MongoDB, Rich Internet Application, IOT, Cloud Computing, da diversi mesi ho aggiunto alla lista, Docker e le architetture a Microservizi. Due temi che, insieme, stanno rivoluzionando l'intero processo di sviluppo e rilascio delle applicazioni sul Cloud. Ambito in cui lavoro da oltre 3 anni, in Enter, occupandomi della piattaforma [Enter Cloud Suite](http://www.entercloudsuite.com){:target="_blank"}.

## Agenda

### Prima sessione
 
**Microservizi.** Dopo un'iniziale presentazione del corso, valuteremo insieme vantaggi e problematiche tipiche delle architetture a microservizi.  
In questa prima sessione vedremo come strutturare una semplice applicazione di esempio. Utilizzeremo il framework [go-micro](https://github.com/micro/go-micro) per sviluppare i microservizi in golang, con l'obiettivo di avere una prima versione funzionante dell'applicazione. 

#### Materiali

- [The Art of Scalability, AKF Scale Cube](http://akfpartners.com/techblog/wp-content/uploads/2008/05/app_cube.png)
- [Play With Docker](http://labs.play-with-docker.com/)
- [The protocol buffer language](https://developers.google.com/protocol-buffers/docs/proto3)
- [Esercitazione su Protobuf e Golang](https://gist.github.com/f-minzoni/8ac69d7193f2f1021743d98e86b67264)
- [Applicazione a Microservizi di esempio](https://gist.github.com/f-minzoni/27d3f9753744bd00150089b88f07c268)

### Seconda sessione

**Docker.** Con un recap iniziale, ripercorreremo tutti i concetti e gli strumenti visti finora, step-by-step, per avviare l'applicazione e valutare il suo funzionamento. Vedremo quindi come utilizzare Docker Compose per definire le componenti software (servizi) dell'applicazione di esempio e renderla facilmente distribuibile.

#### Materiali

- [Introduzione a Docker](https://log-ed.github.io/docker-get-started/sessione1/)
- [Introduzione a Docker Compose](https://log-ed.github.io/docker-get-started/sessione1_1/)
- [Guida alla creazione di un'immagine Docker](https://log-ed.github.io/docker-get-started/sessione2_1/)
- [HTTP/2 Demo](http://www.http2demo.io/)
- [Homepage di gRPC](http://www.grpc.io/)
- [Esercitazione su gRPC e Golang](https://github.com/LOG-ED/go-grpc)

### Terza sessione

**Golang.** In questa terza sessione, dopo il classico recap iniziale, valuteremo meglio la struttura della nostra applicazione e l'implementazione specifica in Golang. Vedremo come creare dei test per verificarne il funzionamento e per strutturare un processo di Continuous Integration.

#### Materiali

- [RPC framework per microservizi](https://github.com/micro/go-micro)
- [Documentazione di Micro](https://micro.mu/docs/go-micro.html#writing-a-service)
- [Homepage di Consul](https://www.consul.io/)
- [Drone](https://github.com/drone/drone) Continuous Delivery platform built on Docker, written in Go
- [Package testing di Golang](https://golang.org/pkg/testing/)

### Quarta sessione

Con l'ultima sessione, avremo tutte le conoscenze minime per creare un nostro progetto di applicazione a microservizi. Lavoreremo a piccoli gruppi per implementare le componenti applicative, creare un processo di verifica e configurare il delivery dei nostri microservizi sul Docker Registry. 
