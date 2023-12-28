import React from 'react';
import Table from 'react-bootstrap/Table';

function ProductList({ products, setProducts }) {

  const sortedProducts = [...products].sort((a, b) => a.id - b.id);

  const reviseButtonClick = (productId) => {

    const newPrice = window.prompt('請輸入新的價格：');

  if (newPrice !== null) {
    const apiUrl = `http://localhost:8080/products/revise/${productId}`;

    const requestData = {
      price: parseFloat(newPrice)
    };

    fetch(apiUrl, {
      method: 'PATCH',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(requestData),
    })
      .then(response => {
        if (!response.ok) {
          throw new Error('No Network response');
        }
        setProducts(prevProducts => prevProducts.map(product => {
          if (product.id === productId) {
            return {
              ...product,
              price: parseFloat(newPrice)
            };
          }
          return product;
        }));
        alert('Price revised successfully');
      })
      .catch(error => {
        alert('Failed to revise price.');
      });
  }
};

  
const deleteButtonClick = (productId) => {
  if (window.confirm('確定要刪除這個商品嗎？')) {
    const apiUrl = `http://localhost:8080/products/delete/${productId}`;

    fetch(apiUrl, {
      method: 'DELETE',
    })
      .then(response => {
        if (!response.ok) {
          throw new Error('No Network');
        }
        setProducts(prevProducts => prevProducts.filter(product => product.id !== productId));
        alert('Product deleted successfully'); // 顯示成功訊息
      })
      .catch(error => {
        alert('Failed to delete product.');
      });
  }
};
  
const createButtonClick = (productId) => {
  const apiUrl = 'http://localhost:8080/orders/add';

  const requestData = {
    customer_id: productId,
    is_paid: false,
  };

  fetch(apiUrl, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(requestData),
  })
    .then(response => {
      if (!response.ok) {
        throw new Error('No Network');
      }
      alert('Order created successfully');
    })
    .catch(error => {
      alert('Failed to create order.');
    });
};


  return (
    <div>
      <Table responsive="sm">
        <thead>
          <tr>
            <th>#</th>
            <th>ID</th>
            <th>NAME</th>
            <th>PRICE</th>
            <th>ACTION</th>
          </tr>
        </thead>
        <tbody>
          {sortedProducts.map((product, index) => (
            <tr key={product.id}>
              <td>{index + 1}</td>
              <td>{product.id}</td>
              <td>{product.name}</td>
              <td>{product.price}</td>
              <td>
                <button onClick={() => reviseButtonClick(product.id)}>
                  修改價格
                </button>{' '}
                <button onClick={() => deleteButtonClick(product.id)}>
                  刪除商品
                </button>{' '}
                <button onClick={() => createButtonClick(product.id)}>
                  新增為訂單
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </Table>
    </div>
  );
}

export default ProductList;
