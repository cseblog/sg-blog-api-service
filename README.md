Running from local: Assume you have docker installation in local
-------------------
git clone <repo url> 
cd sg-blog-api-service

docker build . -t sg-blog-api-service
docker run --rm -it -p 8080:8080 sg-blog-api-service




Running service from docker image:
----------------------------------
docker pull andrewhq/sg-blog-api-service:latest
docker images
docker run andrewhq/sg-blog-api-service:latest
