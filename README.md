## Digital hackathon backand api 

[![Docs Status](https://img.shields.io/badge/Docs-Status-brightgreen.svg?style=flat)](https://github.com/Code7unner/leadersofdigital2020-backend/README.md)

<!-- TOC depthFrom:1 depthTo:2 withLinks:1 updateOnSave:1 orderedList:0 -->
- [User](#user-routes)
- [Products](#products-routes)
- [Order](#order-routes)
- [Store](#store-routes)
<!-- /TOC -->

### User routes:
**`POST`**
+ */user/create*

**`GET`**
+ */user/delete/{id}*
+ */user/{id}* 

### Products routes:
**`POST`**
+ */products/create*

**`GET`**
+ */products/delete/{id}*
+ */products/get*
+ */products/get/{type}*
+ */products/get/{order_id}*

### Order routes:
**`POST`**
+ */order/create*

**`GET`**
+ */order/delete/{id}*
+ */order/{id}* 
+ */order/{courier_id}* 

### Store routes:
**`POST`**
+ */store/create*

**`GET`**
+ */store/delete/{id}*