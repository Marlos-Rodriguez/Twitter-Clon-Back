# Twitter Clon with Go
This is a Twitter Clon Back-End with Go and MongoDB. Make Users, Tweets and follow others users.
### Instalation
Download the code and add a `.env` file with:
- `PORT` Port for API
- `DATABASE_URL` MongoDB URL
- `SECRECT_KEY` Secret password for JWT

execute `go mod download` to download the packages and run `go run main.go`

### Paths
##### User Routes
- `POST` /registro - _Register User_
Body:
```
{
    "email": "coreo@coreo.com",
    "password": "*******",
    "nombre": "Nombre",
    "apellidos": "Apellido",
    "fechaNacimiento": "1970-06-30T00:00:00Z"
}
```
- `POST` /login - _Login User_
Body:
```
{
    "email": "coreo@coreo.com",
    "password": "*******"
}
```

*All the next routes require a JWT in the header, like this:*  `Authorization`: Bearer `token`

- `GET` /profile?id=`ID` - _Get Profile Info_
- `PUT` /modifyProfile - _Modify Profile_
Body: 
```
{
    "nombre": "nombre",
    "apellidos": "apellido",
    "biografia": "Biografia",
    "ubicacion": "Ciudad, Pais",
    "sitioWeb": "www.sitioweb.com",
    "fechaNacimiento": "1970-06-30T00:00:00Z"
}
```

##### Relation Routes

- `POST` /createRelation?id=`ID` - _Create relation with other User_
- `DELETE` /deleteRelation?id=`ID` - _Deleted relation with other User_
- `GET` /requestRelation?id=`ID` - _Return if exits relation with that user_
- `GET` /listUsers?`type`=follow&`page`=1&`search`=*something* - _Get all users_
    - `type`: `follow` return all the users related. `new` return will not
    - `page`: is the batch of users that returns (batch of 20)
    - `search`: This is optional search for something with that character

#### Tweet Route

- `POST` /tweet - _Create a Tweet_
Body: 
```
{
    "message": "message"
}
```
- `GET` /readTweet?id=`ID`&page=1 - _return all tweets of one user_
    - `page` is the batch of tweets that returns (batch of 20)
- `DELETE` /deleteTweet?id=`ID` - _Delete a tweet_
- `GET` /mainTweets?page=1 - _returns the main tweets of the following people_
    - `page` is the batch of tweets that returns (batch of 20)  
#### Image Route
- `POST` /uploadAvatar - _Upload the user avatar_
must be send a image in a form data with name `avatar`
- `POST` /uploadBanner - _Upload the user avatar_
must be send a image in a form data with name `banner`

    
