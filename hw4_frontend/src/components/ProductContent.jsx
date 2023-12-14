import React, { useEffect, useState } from "react";
import ProductList from "./ProductList";
import BusinessNavbar from "./BusinessNavbar";

export default function ProductContent() {
  const [products, setProducts] = useState([]);

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

  return (
    <div>
      <ProductList products={products} />
    </div>
  );
}
