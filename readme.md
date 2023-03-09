>minikube start --ports=8080:30080
> 
>minikube docker-env   
> 
>eval $(minikube -p minikube docker-env)
> 
>docker build . -t grpc_k8s/grpc_k8s
> 
> kubectl create -f compute-service.yaml
> 
>kubectl create -f api.yaml
> 
> minikube service api-service

