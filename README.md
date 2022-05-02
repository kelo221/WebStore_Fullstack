# WebStore_Fullstack

Demo: http://35.178.79.226:8000/
Example user: user,user
Technologies: Fiber, Redis, ArangoDB

## API

### User

| Function    | Methoid | Route (user OR admin)      | Description        | Parametres (JSON)      |
| ----------- | ----------- | ----------- | ----------- | ----------- |
| Register    | POST     | /api/X/register/       |-|"first_name","last_name","email","password","password_confirm"|
| Log In      | POST        | /api/X/login/       |-|"email","password"|
| Log Out     | POST       | /api/X/logout/         |-|-|
| User Information| GET    | /api/X/       | Returns current user information|-|
| Change Password | PUT    | /api/X/password       |-|"password","password_confirm"|

### Products

| Function    | Methoid | Route       | Description       |
| ----------- | ----------- | ----------- | ----------- |
| Products          | GET       | /api/admin/products/       | Returns all products.       |
| Product     | GET        | /api/admin/products/:id/       | Returns a specific product.       |
| Create Products   | POST        | /api/admin/products/       | Creates a new product.       |
| Delete Products   | DELETE        | /api/admin/products/:id      | Deletes a product.       |
| Update Products   | POST        | /api/admin/products/:id      | Updates a product.       |


### Orders

| Function    | Methoid | Route       | Description       |
| ----------- | ----------- | ----------- | ----------- |
| Orders          | GET       | /api/admin/orders/       | Returns all orders.       |

See /src/commands/ for populating the database.

## Frontend
https://github.com/kelo221/Webstore_Fullstack_Frontend
