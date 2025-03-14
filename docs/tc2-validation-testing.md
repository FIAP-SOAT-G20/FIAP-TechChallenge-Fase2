# TC-2 Validation Testing

## Introduction

This document describes how to perform validation testing on the TC-2 API. Validation testing is the process of ensuring that the API meets the requirements specified in the [Software Requirements Specification](./tc2-spec.pdf) document. This document will guide you through the process of testing the API to ensure that it meets the requirements.

After following the steps in the [Readme](./README.md) file, you should have the API running on your local machine.

> [!IMPORTANT]
> Use `http://localhost:8080` as the base URL for the API if you are running it locally via Docker Compose (`make compose-up`).  
> Alternatively, you can use `http://localhost` if you are running it locally via Kubernetes (`make k8s-apply`).  


## Test Cases

The following test cases will be used to validate the TC-2 API:

### 1. **Test Case 1**: Verify that the API health.

```bash
curl --location --request GET 'http://localhost:8080/api/v1/health'
```

### 2. **Test Case 2**: Create a new customer.

> TC-1 2.b.i

```bash
curl --location 'http://localhost:8080/api/v1/customers' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "John Doe 6",
    "email": "john.doe.6@email.com",
    "cpf": "000.000.000-06"
}'
```

### 3. **Test Case 3**: Get a customer by id.

```bash
curl --location 'http://localhost:8080/api/v1/customers/6'
```

### 4. **Test Case 4**: Sign in a customer with CPF.

> TC-1 2.b.ii

```bash
curl --location 'http://localhost:8080/api/v1/sign-in' \
--header 'Content-Type: application/json' \
--data '{
    "cpf": "000.000.000-06"
}'
```

### 5. **Test Case 5**: Get all products.

```bash
curl --location 'http://localhost:8080/api/v1/products
```

### 6. **Test Case 6**: Create a new product.

> TC-1 2.b.iii

```bash
curl --location 'http://localhost:8080/api/v1/products' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Product X",
    "description": "Product X description",
    "price": 13,
    "category_id": 1,
    "active": true
}'
```

### 7. **Test Case 7**: Get a product by id.

```bash
curl --location 'http://localhost:8080/api/v1/products/6'
```

### 8. **Test Case 8**: Update a product.

> TC-1 2.b.iii

```bash
curl --location --request PUT 'http://localhost:8080/api/v1/products/6' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Product X UPDATED",
    "description": "Product X description UPDATED",
    "price": 12.11,
    "category_id": 1,
    "active": true
}'
```

### 9. **Test Case 9**: Delete a product.

> TC-1 2.b.iii

```bash
curl --location --request DELETE 'http://localhost:8080/api/v1/products/6'
```

### 10. **Test Case 10**: Get all products by category.

> TC-1 2.b.iv

```bash
curl --location 'http://localhost:8080/api/v1/products/?category_id=1'
```

### 11. **Test Case 11**: Create a new order.

```bash
curl --location 'http://localhost:8080/api/v1/orders' \
--header 'accept: application/json' \
--header 'Content-Type: application/json' \
--data '{
  "customer_id": 6
}'
```

### 12. **Test Case 12**: Get an order by id.

```bash
curl --location 'http://localhost:8080/api/v1/orders/15' \
--header 'accept: application/json'
```

> The order status should be `OPEN`.

### 13. **Test Case 13**: Add a product to an order.

```bash
curl --location 'http://localhost:8080/api/v1/orders/products/15/2' \
--header 'Content-Type: application/json' \
--data '{
    "quantity": 4
}'
```

### 14. **Test Case 14**: Checkout an order.

> TC-1 2.b.v
> TC-2 1.a.i

```bash
curl --location --request POST 'http://localhost:8080/api/v1/payments/15/checkout'
```

### 15. **Test Case 15**: Get payment status.

> TC-2 1.a.ii

```bash
curl --location 'http://localhost:8080/api/v1/payments/15'
```

### 16. **Test Case 16**: Confirm payment via webhook.

> TC-2 1.a.iii

<details>
<summary>Webhook Request</summary>

```bash
curl --location 'http://localhost:8080/api/v1/payments/callback' \
--header 'Content-Type: application/json' \
--data '{
    "resource": "15",
    "topic": "payment"
}'
```

</details>

> [!IMPORTANT]
> This test case is not meant to be run manually. It is meant to be run by the payment gateway service.

> [!NOTE]
> The payment gateway service will send a POST request to the API with the payment confirmation.  
> The API will then update the order status from `PENDING` to `RECEIVED`.

### 17. **Test Case 17**: Get an order by id.

```bash
curl --location 'http://localhost:8080/api/v1/orders/15' \
--header 'accept: application/json'
```

> The order status should be `RECEIVED`.

### 18. **Test Case 18**: Update an order status with staff.

> TC-2 1.a.v

```bash
curl --location --request PATCH 'http://localhost:8080/api/v1/orders/15' \
--header 'Content-Type: application/json' \
--data '{
    "staff_id": 1,
    "status": "PREPARING"
}'
```

### 19. **Test Case 19**: Get an order by id.


```bash
curl --location 'http://localhost:8080/api/v1/orders/15' \
--header 'accept: application/json'
```

> The order status should be `PREPARING`.

### 20. **Test Case 20**: Update an order status with staff.

```bash
curl --location --request PATCH 'http://localhost:8080/api/v1/orders/15' \
--header 'Content-Type: application/json' \
--data '{
    "staff_id": 1,
    "status": "READY"
}'
```

### 21. **Test Case 21**: Get an order by id.


```bash
curl --location 'http://localhost:8080/api/v1/orders/15' \
--header 'accept: application/json'
```

> The order status should be `READY`.

### 22. **Test Case 22**: Update an order status with staff.

```bash
curl --location --request PATCH 'http://localhost:8080/api/v1/orders/15' \
--header 'Content-Type: application/json' \
--data '{
    "staff_id": 1,
    "status": "COMPLETED"
}'
```

### 23. **Test Case 23**: Get an order by id.


```bash
curl --location 'http://localhost:8080/api/v1/orders/15' \
--header 'accept: application/json'
```

> The order status should be `COMPLETED`.

### 24. **Test Case 24**: Get all orders.

> TC-1 2.b.vi
> TC-2 1.a.iv

```bash
curl --location 'http://localhost:8080/api/v1/orders' \
--header 'accept: application/json'
```

> [!NOTE]
> The list of orders should return them ordered with the following rule:  
> 1. READY > PREPARING > RECEIVED;  
> 2. Older orders first and newer orders last;  
> 3. Orders with status COMPLETED should not appear in the list.  

