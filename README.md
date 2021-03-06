# XaaS: X as a Service

<div>
    <img src="https://img.shields.io/badge/Go-1.17-00ADD8?logo=Go"> 
    <img src="https://img.shields.io/badge/Docker-20.10.7-2496ED?logo=Docker">
    <img src="https://img.shields.io/badge/Echo-3.3.10-ffffff?logo=Echo">
    <img src="https://img.shields.io/badge/TEST Env-ubuntu, macos, windows-008fff">
    <img src="https://github.com/Dev-Beom/XaaS/actions/workflows/CI-TEST.yml/badge.svg">
</div>

It is a data & command(task) delivery and status monitoring service.  
This process sends all content to nodes within an instance, and users can check all the processes.  
The structure was simply constructed based on the Kubernetes structure, and was developed using the golang and echo framework.  
If you have any opinions on **XaaS** project, please enroll for issues and pull requests.  
#### Specifications for APIs and key methods can be found [here](https://dev-beom.github.io/XaaS).

## Architecture
![XaaS drawio](https://user-images.githubusercontent.com/66074802/148428900-aa5c780a-222e-4d99-9da3-e2fce1fed47d.png)  
1. The flow of `commands from outside to instances` and `information from instances to outside`.  
2. The flow in which `File i/o operations are delivered by accessing the storage` inside the instance.  
3. `Deliver commands(tasks), data, and status` to nodes inside the instance. It also `binds the storage` of the instance.  

## How to use
```shell
# XaaS run.
go run main.go

# All file tests
cd apiserver
go test -v ./...

cd controlmanager
go test -v ./... 
```
