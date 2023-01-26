# fampay-assignment

#Local Setup

clone this repository using command below

"git clone https://github.com/singhbrijesh540/fampay-assignment.git"

Run Postgres server in your local


Create a database in your local postgreys with name: "assignment"



Create a table in the above created "assignment" database using below Query:

CREATE TABLE video_detail (
id serial4 NOT NULL,
created_at timestamp NULL DEFAULT now(),
updated_at timestamp NULL DEFAULT now(),
deleted_at timestamp NULL,
title text NOT NULL,
description text NOT NULL,
published_at timestamp NOT NULL,
thumbnail_url  text NOT NULL,
PRIMARY KEY (id),
Unique (thumbnail_url)
);



Now open the cloned project(fampay-assignment) in some IDE(preferably GoLand).
Update dbConfig "fampay-assignment/config/config.go" set USERNAME And PASSWORD of your Local postgres db and run "main.go" file

OR

open terminal and go the the file where project is cloned using "cd" command: 

eg: cd fampay-assignment

when you reached fampay-assignment folder in terminal, Run the below commands:

go mod download    // This will download all dependency

then run 

"go run main.go"   


The above command will run this application and you can see the message like: "Listening and serving HTTP on :8081" 


Now go to the Postman or browser and hit the below Api's


"http://localhost:8081/fampay-assignment/video-detail?page=0&size=5"                               (get video detail with pagination)  
or
"http://localhost:8081/fampay-assignment/search/video-detail?title=Rohit Sharma&description=ODI"   (search using title and description)            
or
"http://localhost:8081/fampay-assignment/search/v2/video-detail?query=Sharma Rohit steady"         (Optimized Search)

