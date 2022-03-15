# WebStore_Fullstack

## API

### User

| Function    | Methoid | Route       | Description       |
| ----------- | ----------- | ----------- | ----------- |
| Register    | POST     | /api/admin/register/       |-|
| Log In      | POST        | /api/admin/login/       |-|
| Log Out     | POST       | /api/admin/logout/         |-|
| User Information| GET    | /api/admin/user/       | Returns current user information|
| Change Password | PUT    | /api/admin/user/password       |-|

### Products

| Function    | Methoid | Route       | Description       |
| ----------- | ----------- | ----------- | ----------- |
| Products          | GET       | /api/admin/products/       | Returns all products.       |
| Product (ID)      | GET        | /api/admin/products/:id/       | Returns a specific product.       |
| Create Products   | POST        | /api/admin/products/       | Creates a new product.       |
| Delete Products   | DELETE        | /api/admin/products/:id      | Deletes a product.       |
| Update Products   | POST        | /api/admin/products/:id      | Updates a product.       |


### Orders

| Function    | Methoid | Route       | Description       |
| ----------- | ----------- | ----------- | ----------- |
| Orderst          | GET       | /api/admin/orders/       | Returns all orders.       |

See /src/commands/ for populating the database.
