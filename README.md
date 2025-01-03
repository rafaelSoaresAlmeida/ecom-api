# ecom-api

This is just a tutorial that I have followed to learn Go, I didn't created that...(the original link: https://www.youtube.com/watch?v=7VLmLOiQ3ck)

The paylads:

http://localhost:8080/api/v1/register

{
"email": "lucifer@demon.com",
"password": "demon666",
"firstName": "lucifer",
"lastName": "capeta"
}

http://localhost:8080/api/v1/login

{
"email": "Joao@deus.com",
"password": "deusGood33"
}

http://localhost:8080/api/v1/product

{
"name": "bone",
"description": "bone basico",
"image": "image.png",
"price": 42.50,
"quantity": 100
}

http://localhost:8080/api/v1/cart/checkout

{
"Items": [
{
"ProductID": 1,
"Quantity": 10
}
]  
}
