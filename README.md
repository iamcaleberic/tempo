# tempo

Using temporal and client-go to interact with clusters.

# Requirements

- golang 1.23

- temporal-cli

    `brew install temporal`

- minikube

    `brew install minikube`

- podman

    `brew install podman`
- running (local) kubenernetes cluster
    ```
        podman machine init
        podman machine start
        minikube start --driver=podman --container-runtime=cri-o
    ```


# How to run

- start local temporal server

    `temporal server start-dev`

    - with docker-compose
    
        ```
            git clone https://github.com/temporalio/docker-compose.git
            cd  docker-compose
            docker-compose up
        ```

- start workers

    `go run worker/main.go`

- run worklows

    `go run app/main.go`

# Notes 

- By default the app uses `~/.kube/config` and gets info from currectly selected context 
