# Project Description

Our team has made the project dedicated to Animal Shelters. In this project we have implemented Frontend (HTML, CSS, TS) and Backend (Go, PostgreSQL).

 This is the website where people can look for pets they can adopt (photos and other data) and also give up a pet.

There are 3 creatures: Admin, User and Animal. 

Our team have made a database for pets, containing description data, RestAPI and functional website.  

## Team Members
 
Polina Stelmakh 22B030588

Anel Tulepbergen 22B030602

Alina Amreyeva 22B031240

Dossym Ibray  22B030545


## Build

Run `ng build` to build the project. The build artifacts will be stored in the `dist/` directory.


## API Structure

Animal Shelter REST API

POST /animals

GET /animals/:id

PUT /animals/:id

DELETE /animals/:id


## Creatures' Structure

Users
• UserID (AUTO_INCREMENT PRIMARY KEY)
• User_Email(TEXT)
• Username (VARCHAR(30))
• Password (encrypted password)
• Number_of_phone_user(VARCHAR(50))
• Profile_Picture_User(TEXT)
• Role(Default “User”)
Animal
• ID(AUTO_INCREMENT PRIMARY KEY)
• Kind_Of_Animal(VARCHAR(255))
• Breed_Of_Animal(VARCHAR(255))
• Name(VARCHAR(255))
• Age(INTEGER)
• Description(TEXT)
• Animal_Picture(TEXT)
Admins
• AdminID(AUTO_INCREMENT PRIMARY KEY)
• Admin_Email(TEXT)
• Adminame (VARCHAR(30))
• Password (encrypted password)
• Number_of_phone_Admin(VARCHAR(50))
• Profile_Picture_Admin(TEXT)
• Role(VARCHAR(255), Type:Back_end, Front_end, G_admin, etc)
Role
• Permissions(TEXT)