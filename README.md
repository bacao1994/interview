# Interview Restful API 
This is an application built with gin-gonic, jwt, gorm, logrus, postgresql, redis, docker, docker-compose. 
As Backend requirements: "This project include many issues from code to DB strucutre, feel free to optimize them". 
I have changed a lot of things. My structure base on Awesome project lists using Gin web framework: https://github.com/photoprism/photoprism
##### New changed
 - I have chosen new structure source code and used Gin-gonic framework.
 - Changed SQLite to PostgreSQL and Gorm.
 - Improved testing by Postman with Swagger UI and UnitTest
 - Support to run source code with docker and docker-compose and wrote some script for starting, clean up docker images.
 - Improved JWT algorithm: HS256 -> RS256. Then, add 1 more step for hashing token(I chose AES encrypted and decrypted because it's fast). Thus, the hacker can't decrypt and try to read the body of JWT. 
 - Provided Redis for caching. If the daily limit is reached, we can reject quickly and reduce query records from DB. 
 - Provided Logrus. Log anything into log file, separate by day.  
 - I am aware that we can use Interfaces and mock database calls, but I decided to be replicate the exact same process that happens in real life for testing.
   We can define a test database in our .env so, everyone can create a new database for the tests alone.
##### Improvement later
 - Try to support CICD and deploy source code with no downtime: kubernetes, docker swarm,CircleCI,...
 - Hide swagger UI (run as production mode)
##### Run project
 - With docker:
   + [Docker](https://www.docker.com/)  installed (**version 17.12.0+ is required**)
   + [Docker compose](https://docs.docker.com/compose/install/) installed (**version 1.21.0+ is required**)
   + Please see docker-compose.yml and change .env file. 
   + 