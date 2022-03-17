# WebStore_Fullstack

## API

### User

| Function    | Methoid | Route       | Description        | Parametres (JSON)      |
| ----------- | ----------- | ----------- | ----------- | ----------- |
| Register    | POST     | /api/admin/register/       |-|"first_name","last_name","email","password","password_confirm"|
| Log In      | POST        | /api/admin/login/       |-|"email","password"|
| Log Out     | POST       | /api/admin/logout/         |-|-|
| User Information| GET    | /api/admin/user/       | Returns current user information|-|
| Change Password | PUT    | /api/admin/user/password       |-|"password","password_confirm"|

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
