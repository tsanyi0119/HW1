import React, { useState } from 'react';
import axios from 'axios';
import OrderList from './OrderList';
import Button from 'react-bootstrap/Button';

function OrderContent() {
  const [orderId, setOrderId] = useState("");
  const [items, setItems] = useState([]);

  const handleButtonClick = async () => {
    try {
      const response = await axios.get(`http://localhost:8080/orders/${orderId}/items`);
      setItems(response.data);
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  };

  return (
    <div>
      <div style={{ display: 'flex', alignItems: 'center' }}>
        <input
          type="text"
          placeholder="輸入訂單號碼"
          style={{ marginRight: '8px' , marginLeft: '10px'}}
          value={orderId}
          onChange={(e) => setOrderId(e.target.value)}
        />
        <Button variant="primary" onClick={handleButtonClick}>查詢</Button>{' '}
      </div>
      <OrderList order={items} />
    </div>
  );
}

export default OrderContent;
