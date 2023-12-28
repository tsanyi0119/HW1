import React, { useEffect, useState } from "react";
import ProductList from "./ProductList";

export default function ProductContent() {
  const [products, setProducts] = useState([]);
  const [newProduct, setNewProduct] = useState({
    name: "",
    price: "",
    category_id: "",
  });

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch("http://localhost:8080/products");
        const data = await response.json();
        console.log(data);
        setProducts(data);
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    };

    fetchData();
  }, []);

  const handleInputChange = (event) => {
    const { name, value } = event.target;
    setNewProduct((prevProduct) => ({
      ...prevProduct,
      [name]: value,
    }));
  };

  const addProduct = async () => {
    const requestData = {
      name: newProduct.name,
      price: parseFloat(newProduct.price),
      category_id: parseInt(newProduct.category_id, 10),
    };
  
    try {
      const apiUrl = 'http://localhost:8080/products/add';
  
      const response = await fetch(apiUrl, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(requestData),
      });
  
      if (!response.ok) {
        throw new Error('No Network');
      }
      window.alert('Product added successfully');
      setNewProduct({
        name: "",
        price: "",
        category_id: "",
      });
    } catch (error) {
      console.error('Error during POST request:', error);
    }
  };

  return (
    <div>
      <div style={{ display: "flex", alignItems: "center" }}>
      <div style={{ marginRight: "10px" }}>
        <label>Name:</label>{' '}
        <input
          type="text"
          name="name"
          value={newProduct.name}
          onChange={handleInputChange}
        />
      </div>
        <div style={{ marginRight: "10px" }}>
          <label>Price:</label>{' '}
          <input
            type="text"
            name="price"
            value={newProduct.price}
            onChange={handleInputChange}
          />
        </div>
        <div style={{ marginRight: "10px" }}>
          <label>Category ID:</label>{' '}
          <input
            type="text"
            name="category_id"
            value={newProduct.category_id}
            onChange={handleInputChange}
          />
        </div>
        <button onClick={addProduct}>新增產品</button>
      </div>

      <ProductList products={products} setProducts={setProducts} />
    </div>
  );
}
