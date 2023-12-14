// import Table from 'react-bootstrap/Table';

// function ProductList() {
//   return (
//     <div>
//       <Table responsive="sm">
//         <thead>
//           <tr>
//             <th>#</th>
//             <th>ID</th>
//             <th>NAME</th>
//             <th>PRICE</th>
//           </tr>
//         </thead>
//         <tbody>
//           <tr>
//             <td>1</td>
//             <td>Table cell</td>
//             <td>Table cell</td>
//             <td>Table cell</td>
//           </tr>
//           <tr>
//             <td>2</td>
//             <td>Table cell</td>
//             <td>Table cell</td>
//             <td>Table cell</td>
//           </tr>
//           <tr>
//             <td>3</td>
//             <td>Table cell</td>
//             <td>Table cell</td>
//             <td>Table cell</td>
//           </tr>
//         </tbody>
//       </Table>
//     </div>
//   );
// }

// export default ProductList;

import React from 'react';
import Table from 'react-bootstrap/Table';

function ProductList({ products }) {
  return (
    <div>
      <Table responsive="sm">
        <thead>
          <tr>
            <th>#</th>
            <th>ID</th>
            <th>NAME</th>
            <th>PRICE</th>
          </tr>
        </thead>
        <tbody>
          {/* 使用map函數來動態渲染每一個產品的行 */}
          {products.map((product, index) => (
            <tr key={product.id}>
              <td>{index + 1}</td>
              <td>{product.id}</td>
              <td>{product.name}</td>
              <td>{product.price}</td>
            </tr>
          ))}
        </tbody>
      </Table>
    </div>
  );
}

export default ProductList;
