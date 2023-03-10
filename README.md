# blogspot


REST API for blog using golang, gin framework & ent ORM.


## Requirements
- Golang 
- Postgres


## Run Locally

Clone the project

```bash
  git clone https://github.com/roshanikhairnar/blogspot.git
```

Go to the project directory
go ro the main.go file.
change the credentials of postgres db in the main function.

```bash
  go run main.go
```



## API Reference

#### Get all blogs. you can use postman or curl request

```
  GET localhost:8080/blogs
```


#### Create admin account
collect token generated after making the request, might need it for authorisation if you want to perform CRUD operation

```
  POST /api/register
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`      | `string` | **Required**. name as admin name |
| `password`      | `string` | **Required**. password for login |

### Login as admin
```
  POST /api/login
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`      | `string` | **Required**. name as admin name |
| `password`      | `string` | **Required**. password for login |

### Create blog
```
  POST /api/admin/createblog
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `blogTitle`      | `string` | **Required**. Title of blog |
| `blogType`      | `string` | **Required**. Type of blog for e.g tech,fashion etc|
| `blogContent`      | `string` | **Required**. blog content |
| `blogAuthor`      | `string` | **Required**. author of blog |

### Delete blog
```
DELETE /api/admin/deleteblog/{$id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. id of blog |

### Update blog
```
PUT /api/admin/updateblog/{$id}
 ```
 | Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `blogTitle`      | `string` | **Required**. title of blog |


