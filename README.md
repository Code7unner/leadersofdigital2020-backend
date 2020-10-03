## Digital hackathon backand api 

[![Docs Status](https://img.shields.io/badge/Docs-Status-brightgreen.svg?style=flat)](https://github.com/Code7unner/leadersofdigital2020-backend/blob/main/README.md)

<!-- TOC depthFrom:1 depthTo:2 withLinks:1 updateOnSave:1 orderedList:0 -->
- [User](#user-routes)
- [Products](#products-routes)
- [Order](#order-routes)
- [Store](#store-routes)
  - [Register user](#register-user)
<!-- /TOC -->

### User routes:
**`POST`**
+ */api/v1/user/create*
```json5
{
  "id": 1,
  "name": "Name",
  "phone": "+79149647499",
  "password": "1234",
  "address": "test",
  "sex": "male",
  "role": "admin"
}
```

**`GET`**
+ */api/v1/user/{id}* 

**`DELETE`**
+ */api/v1/user/delete/{id}*

### Products routes:
**`POST`**
+ */api/v1/products/create*
```json5
{
  "id": 1,
  "name": "name",
  "type": "meat",
  "description": "1kg of meat",
  "price": 329.0,
  "img_url": "google.com/images",
  "additional_info": "super good meat"
}
```

**`GET`**
+ */api/v1/products/get*
+ */api/v1/products/get/{type}*
+ */api/v1/products/get/{order_id}*

**`DELETE`**
+ */api/v1/products/delete/{id}*

### Order routes:
**`POST`**
+ */api/v1/order/create*
```json5
{
  "id": 1,
  "courier_id": 2,
  "status": true
}
```

**`GET`**
+ */api/v1/order/delete/{id}*
+ */api/v1/order/{id}* 
+ */api/v1/order/{courier_id}* 

**`DELETE`**
+ */api/v1/order/delete/{id}*

### Store routes:
**`POST`**
+ */api/v1/store/create*
```json5
{
  "id": 1,
  "name": "name",
  "address": "test street"
}
```

**`GET`**
+ */api/v1/store/delete/{id}*

**`DELETE`**
+ */api/v1/store/delete/{id}*

### Register user:
**`POST`**
+ */register*
```json5
{
  "phone": "+79149647499",
  "password": "pass"
}
```