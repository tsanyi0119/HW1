import React from 'react';
import Table from 'react-bootstrap/Table';

function OrderList({ order }) {
    return (
      <div>
        <Table responsive="sm">
          <thead>
            <tr>
              <th>ID</th>
              <th>訂單ID</th>
              <th>產品ID</th>
              <th>是否已發貨</th>
            </tr>
          </thead>
          <tbody>
            {order.map((item, index) => (
              <tr key={item.id}>
                <td>{item.id}</td>
                <td>{item.order_id}</td>
                <td>{item.product_id}</td>
                <td>{item.is_shipped ? '是' : '否'}</td>
              </tr>
            ))}
          </tbody>
        </Table>
      </div>
    );
  }
  
  export default OrderList;
  