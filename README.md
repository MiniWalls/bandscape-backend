# BandScape-backend
The backend server for BandScape project, the "social media" platform with LastFM api for discussing music!

Front-end repo: https://github.com/MiniWalls/bandscape-frontend

## Introduction
Main goal of this project was to deepen my knowledge in web development technologies and learn new and essential ones. The back-end app utilizes Golang with Gin webframework and MySQL database on Google Cloud. Back-end handles most of the HTTP calls made by front-end. Both LastFM API authorization and routes used to get info on music are routed on back-end with all the required utils to make using LastFM API convenient for front-end. The back-end server also handles the creation of new users on logging in aswell as updating old credentials. Finally back-end handles making new posts and retrieving existing ones on demand. The back-end is hosted on Google App Engine.

### The API routes
![Capute5](https://github.com/MiniWalls/bandscape-backend/assets/69449726/a57c84a7-8726-4158-9a6d-3d5d1a4f900e)
