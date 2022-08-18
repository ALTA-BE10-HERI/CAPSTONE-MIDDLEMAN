# Middleman-Apps

<!-- ABOUT THE PROJECT -->

## 💻 About The Project

Middleman built to help between stores and wholesalers.

Feature in Middleman

  <!--- feature USER
   --->
<div>
      <details>
<summary>🙎 Users</summary>

In users, there is a feature to login either user or admin, we also create Create, Read, Update, Delete for users here

<div>
  
| Feature User | Endpoint | Param | JWT Token | Function |
| --- | --- | --- | --- | --- |
| GET | /users/products  | - | YES | get all data product user |
| POST | /users/products | - | YES | add product (not available in distributtor) |
| GET | /users/products/search | productname | YES | serach product |
| PUT | /users/products| idproduct | YES | update product (not available in distributtor) |
| DELETE | /users/products | idproduct | YES | delete product (users) |

</details>
<div>
      <details>
<summary>📋 User Product</summary>

In User Product, there is a feature to Create, Read, Update, Delete product but not available in product Admin

<div>
  
| Feature User Product | Endpoint | Param | JWT Token | Function |
| --- | --- | --- | --- | --- |
| POST | /admins/products  | - | YES | new product by Admin |
| GET | /admins/products | - | NO | get all product sell  |
| GET | /admins/products/search | productname | NO | serach product |
| PUT | /admins/products | idproduct | YES | update product by Admin |
| DELETE | /admins/products | idproduct | YES | delete product by id |

</details>

<div>
      <details>
<summary>🛒 Carts</summary>

In Carts for user to create Cart before order

<div>
  
| Feature Cart | Endpoint | Param | JWT Token | Function |
| --- | --- | --- | --- | --- |
| GET | /carts | - | YES | get all product in cart by user  |
| POST | /carts  | - | YES | add product in cart |
| PUT | /carts | idproduct | YES | update cart |
| DELETE | /carts | idproduct | YES | delete cart |

</details>

<div>
      <details>
<summary>👨‍💻 Admin Product</summary>

In Admin, there is a feature to Create, Read, Update, Delete product to shell in application

<div>
  
| Feature Admin | Endpoint | Param | JWT Token | Function |
| --- | --- | --- | --- | --- |
| POST | /admins/products  | - | YES | new product by Admin |
| GET | /admins/products | - | NO | get all product sell  |
| GET | /admins/products/search | idproduct | NO | serach product |
| PUT | /admins/products | idproduct | YES | update product by Admin |
| DELETE | /admins/products | idproduct | YES | delete product by id |

</details>

<div>
      <details>
<summary>🛍️ Order</summary>

In Order, feature to transaction order

<div>
  
| Feature Order | Endpoint | Param | JWT Token | Function |
| --- | --- | --- | --- | --- |
| POST | /orders/users  | - | YES | create new order user |
| GET | /orders/users | - | YES | get all history order  |
| GET | /orders/ | idorder | YES | get detail order user and admin |
| GET | /orders/admins | - | YES | get all history order admin |
| GET | /orders/admins/incoming | - | YES | get incomming order from user (ADMIN) |
| PUT | /orders/confrim/ | idorder | YES | confrim order by id(ADMIN) |
| PUT | /orders/done/ | idorder | YES | finish order by id(ADMIN) |

</details>

<div>
      <details>
<summary>📜 Inoutbounds</summary>

In Inoutbounds feature handle realation stok in admin and user. if admin out == user in

<div>
  
| Feature Order | Endpoint | Param | JWT Token | Function |
| --- | --- | --- | --- | --- |
| GET | /inoutbounds | - | YES | get cart for stock user (out) and stok admin (in)  |
| POST | /inoutbounds  | - | YES | create new cart for stock user (out) and stok admin(in) |
| PUT | /inoutbounds/ | idproducts | YES | update quantity product in carts for stock user (out) and admin (in) |
| DELETE | /inoutbounds/ | idproducts | YES | delete product by id in carts for stock user (out) and admin (in) |

</details>

<div>
      <details>
<summary>📊 Inventories</summary>

In Inventories feature to record stok in and out from inventory user admin

<div>
  
| Feature Order | Endpoint | Param | JWT Token | Function |
| --- | --- | --- | --- | --- |
| POST | /users/inventory  | - | YES | create a form to list product (OUT) |
| GET | /users/inventory | - | YES | get all form product inventory (OUT)  |
| GET | /users/inventory/ | idinventory | YES | get detail form product inventory (outbound)  | 
| POST | /admins/inventory | - | YES | create a form to list product (IN) |
| GET | /admins/inventory | - | YES | get all form product inventory (IN)  |
| GET | /admins/inventory/ | idinventory | YES | get detail form product inventory (inbound)  |

</details>

### 🛠 &nbsp;Build App & Database

![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)
![Visual Studio Code](https://img.shields.io/badge/Visual%20Studio%20Code-0078d7.svg?style=for-the-badge&logo=visual-studio-code&logoColor=white)
![MySQL](https://img.shields.io/badge/mysql-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![AWS](https://img.shields.io/badge/AWS-%23FF9900.svg?style=for-the-badge&logo=amazon-aws&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Ubuntu](https://img.shields.io/badge/Ubuntu-E95420?style=for-the-badge&logo=ubuntu&logoColor=white)
![Cloudflare](https://img.shields.io/badge/Cloudflare-F38020?style=for-the-badge&logo=Cloudflare&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens)
![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)
![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)

## 🗃️ ERD

<img src="ERD.png">

## Run Locally

Clone the project

```bash
      https://github.com/ALTA-BE10-HERI/CAPSTONE-MIDDLEMAN
```

Go to the project directory

```bash
      cd CAPSTONE-MIDDLEMAN
```

## Open Api

if you want to consume our api,
here's the way !

```bash
https://app.swaggerhub.com/apis-docs/vaniliacahya/capstone/1.0.0#/
```

## Authors

[![GitHub Hilmi](https://img.shields.io/badge/-Heri-white?style=flat&logo=github&logoColor=black)](https://github.com/darmon17)
[![GitHub Hilmi](https://img.shields.io/badge/-Ivan-white?style=flat&logo=github&logoColor=black)](https://github.com/ivands26)
[![GitHub Hilmi](https://img.shields.io/badge/-Vanilia-white?style=flat&logo=github&logoColor=black)](https://github.com/vaniliacahya)

<h3>
 <p align="right">(<a href="#top">back to top</a>)</p>
<p align="center">:copyright: 2022  </p>
</h3>
