# Online Shopping API

This is a simple online shopping system project that allows customers to browse and purchase products. The project includes the following features:

## Features

1. **View Product List by Category**
    - Customers can easily browse products by their respective categories, making it convenient to find items of interest.

2. **Add Products to Shopping Cart**
    - Customers can select products they want to purchase and add them to their shopping cart for easy tracking.

3. **View Shopping Cart**
    - Customers can view a list of products that have been added to their shopping cart, including product details and quantities.

4. **Delete Products from Shopping Cart**
    - If customers change their minds or want to remove items from their cart, they can easily delete products from the shopping cart.

5. **Checkout and Make Payments**
    - Customers can proceed to checkout from their shopping cart and complete the purchase.

6. **Login and Register**
    - The system provides user authentication, allowing customers to create accounts (register) or log in to existing accounts.

## Getting Started

To get this project up and running, follow these steps:

1. **Clone the Repository**
    - Clone this repository to your local machine.
   ```bash
   git clone https://github.com/rilgilang/synapsis-assesment
   ```
2. **Installing go dependencies**
    - Run this command to install depedencies.
   ```bash
   go mod tidy
   go mod init
   ```
3. **Create your database using MySQL.**
    - This apps will auto migrate when you start it.
    - since im not providing seeders you need to seed it by yourself, you can run this sql syntax to seed the data in your MySQL
    - Remember to change the database name
   ```bash
   INSERT INTO synapsis_assignment.product_categories
   (id, created_at, updated_at, deleted_at, category_name)
   VALUES('6880f736-2ae5-4efa-b72a-541f880bec49', curdate(), curdate(), NULL, 'gadgets');

   INSERT INTO synapsis_assignment.product_categories
   (id, created_at, updated_at, deleted_at, category_name)
   VALUES('3c1e7c2f-a8c0-4109-8b81-550a1bcaf7f3', curdate(), curdate(), NULL, 'fashion');

   INSERT INTO synapsis_assignment.product_categories
   (id, created_at, updated_at, deleted_at, category_name)
   VALUES('089fbbc6-b1a2-4af1-81e9-3c3304ea3de5', curdate(), curdate(), NULL, 'automotive');


   INSERT INTO synapsis_assignment.products
   (id, created_at, updated_at, deleted_at, product_name, category_id, price)
   VALUES('09576280-6bba-42dc-ab99-e9c7a66eaa02', curdate(), curdate(), NULL, 'Macbook Pro M2 13 Inch 8/16 256GB', '6880f736-2ae5-4efa-b72a-541f880bec49', 18200);

   INSERT INTO synapsis_assignment.products
   (id, created_at, updated_at, deleted_at, product_name, category_id, price)
   VALUES('00d84aed-4fe8-46a8-90ae-20d89bc8b792', curdate(), curdate(), NULL, 'Uniqlo Oversize L', '3c1e7c2f-a8c0-4109-8b81-550a1bcaf7f3', 500);

   INSERT INTO synapsis_assignment.products
   (id, created_at, updated_at, deleted_at, product_name, category_id, price)
   VALUES('52515c7b-1842-47da-b37c-d01fe5cb6247', curdate(), curdate(), NULL, 'Supra bapak mulus like new', '089fbbc6-b1a2-4af1-81e9-3c3304ea3de5', 8000);
   ```

4. **Setup the config file.**
    - go to **_config/yaml_** folder then rename **app.yaml.example** to **app.yaml**
    - change the database name, host, username and password value with your MySQL that you've created before  
    - you can change the value of jwt too

5. **Running the application.**
    - running this app is simple just run this command in your terminal
   ```bash
   go run main.go
   ```
   - migration will auto migrate when you start this app

And that's it


## API Documentation

I published API documentation using postman, feel free to check it out to see the list of endpoint

[API Documentation](https://documenter.getpostman.com/view/29535426/2s9YJaYPX9) 



