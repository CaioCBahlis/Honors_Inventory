Dear Honors IT Team

To run this project:

git clone this repo and
cd into the project root

docker-compose up -d 
(flag ensures the containers run in the background)

Access localhost:8080/ on your browser

To stop running it 

docker-compose down -v;
(-v ensures the psql also stops)

The Server waits for the deployment of the Database
to start, so it might take a few seconds for everything to boot up.

About the Project:

Due to the amount of dependencies and different packages,
I decided to containerize the project, for easier deployment
and ensuring it would work on any machine

It is an Honor to have the opportunity to showcase the challenge to the team
,and I'm extremely grateful for this opportunity.

Thank you :)

Made with passion by

-Caio Bahlis