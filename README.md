# Running Redis With Docker Locally

Let’s get started with this tutorial and download the redis docker image and run it using the following 2 docker commands:

    docker pull redis
  
    docker run --name redis-test-instance -p 6379:6379 -d redis

The first pull command does the job of retrieving the redis image from DockerHub so that we can then run it as a container using the second command. 
In the second command we specify the name of our redis container and we also map our local port 6379 to the port that redis is running against within
the container using the -p flag.
